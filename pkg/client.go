package pkg

import (
	"api-gateway/config"
	pbi "api-gateway/genproto/item"
	pbu "api-gateway/genproto/user"
	"log"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewUserClient(cfg *config.Config) pbu.UserServiceClient {
	conn, err := grpc.NewClient(cfg.USER_SERVICE_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Println(errors.Wrap(err, "failed to connect to the address"))
		return nil
	}

	return pbu.NewUserServiceClient(conn)
}


func NewItemClient(cfg *config.Config) pbi.ItemServiceClient {
	conn, err := grpc.NewClient(cfg.ITEM_SERVICE_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
		
	if err != nil {
		log.Println(errors.Wrap(err, "failed to connect to the address"))
		return nil
	}
	
	return pbi.NewItemServiceClient(conn)
}
