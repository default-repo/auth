package user

import desc "github.com/default-repo/auth/pkg/proto/auth_v1"

type Implementation interface {
	desc.UnimplementedAuthV1Server
}
