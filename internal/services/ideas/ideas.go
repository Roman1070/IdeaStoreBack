package ideas

import (
	"idea-store-auth/internal/grpc/ideas"
	"log/slog"
)

type Ideas struct {
	log *slog.Logger
	Api ideas.IdeasAPI
}


// New returns a new instance of the Auth service.
func New(log *slog.Logger, ideasApi ideas.IdeasAPI) *Ideas {
	return &Ideas{
		log:         log,
		Api:ideasApi,
	}
}