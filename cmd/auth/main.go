package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/default-repo/auth/internal/model"
	"log/slog"
	"os"

	"github.com/default-repo/auth/internal/config"
	"github.com/default-repo/auth/internal/config/env"
	"github.com/default-repo/auth/internal/repository/pg_db"
	g "github.com/default-repo/auth/internal/server/grpc"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	ctx := context.Background()

	flag.Parse()

	if err := config.Load(configPath); err != nil {
		logger.Error("config loading failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

	pgConfig, err := env.NewPGConfig()
	if err != nil {
		logger.Error("pgDB config configuring failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

	db, err := pg_db.NewPGStore(*logger, pgConfig.DSN())
	if err != nil {
		logger.Error(
			"PG store creating failed",
			slog.String("DSN", pgConfig.DSN()),
			slog.String("error", err.Error()),
		)
		os.Exit(1)
	}

	if err := basicDBInteraction(ctx, db); err != nil {
		logger.Error("basicDB interaction failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

	grpcConfig, err := env.NewGRPCConfig()
	if err != nil {
		logger.Error("grpc config configuring failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

	listener, err := g.NewListener(grpcConfig)
	if err != nil {
		logger.Error(
			"grpc listener creating failed",
			slog.String("address", grpcConfig.Address()),
			slog.String("error", err.Error()),
		)
		os.Exit(1)
	}

	grpcServer, err := g.NewGRPCServer()
	if err != nil {
		logger.Error("grpc server starting failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

	defer func() {
		grpcServer.S.Stop()
		db.Close()
	}()

	fmt.Printf("grpc server started on [ %s ]\n", grpcConfig.Address())

	if err := grpcServer.S.Serve(*listener.NetListener); err != nil {
		logger.Error("serve finished with error", slog.String("error", err.Error()))
		os.Exit(1)
	}
}

func basicDBInteraction(ctx context.Context, db *pg_db.PGStore) error {
	customer := model.Customer{
		UUID:     gofakeit.UUID(),
		Name:     gofakeit.Name(),
		Password: gofakeit.Password(true, true, true, true, true, 8),
		Email:    gofakeit.Email(),
	}

	lastID, err := db.InsertData(ctx, customer)
	if err != nil {
		return errors.New("inserting by ID failed: " + err.Error())
	}

	rows, err := db.List(ctx, 2)
	if err != nil {
		return errors.New("executing list failed: " + err.Error())
	}

	defer rows.Close()

	fmt.Println("\nList of customers:")
	for rows.Next() {
		var id int
		var name, uuid, password, email string

		if err := rows.Scan(&id, &uuid, &name, &password, &email); err != nil {
			return errors.New("scanning row failed: " + err.Error())
		}

		fmt.Printf("ID: %d, UUID: %s, Name: %s, Password: %s, Email: %s\n", id, uuid, name, password, email)
	}

	result, err := db.UpdateByID(ctx, lastID)
	if err != nil {
		return errors.New("update by ID failed: " + err.Error())
	}

	if result > 0 {
		fmt.Printf("\nSuccessfully rows updated: [ %d ]\n", result)
	}

	c, err := db.GetCustomerByUID(ctx, lastID)
	if err != nil {

		return errors.New("get by ID failed: " + err.Error())
	}

	fmt.Printf("\nLast created/updated customer: \nID:\t%d\nName:\t%s\nPas.:\t%s\nEmail:\t%s\n\n", c.ID, c.Name, c.Password, c.Email)

	return nil
}
