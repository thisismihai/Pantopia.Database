package gapi

import (
	"fmt"

	db "github.com/mbaxamb3/pantopia/db/sqlc"
	"github.com/mbaxamb3/pantopia/pb"
	"github.com/mbaxamb3/pantopia/token"
	"github.com/mbaxamb3/pantopia/util"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	pb.UnimplementedPantopiaServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewServer creates a new GRPC server and setup routing
func NewServer(config util.Config, store db.Store) (*Server, error) {

	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create tokenmaker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil

}
