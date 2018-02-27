package grpc

import (
	"log"

	"context"
	"io"

	"fmt"

	"time"

	"os"

	"github.com/Sharykhin/gl-mail-api/entity"
	"github.com/Sharykhin/gl-mail-grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Server is a reference to a private struct that represent
var Server server

type server struct {
	client api.FailMailClient
}

func (s server) GetList(ctx context.Context, limit, offset int64) ([]entity.FailMail, error) {
	filter := &api.FailMailFilter{Limit: limit, Offset: offset}

	stream, err := s.client.GetFailMails(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("could not stream fail mails: %v", err)
	}
	var fm []entity.FailMail
	for {
		m, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("%v.GetFailMails(_) = _, %v", s.client, err)
		}
		t, err := time.Parse("2006-01-02 15:04:05", m.CreatedAt)
		if err != nil {
			fmt.Println(err)
		}
		fm = append(fm, entity.FailMail{
			ID:        m.ID,
			Action:    m.Action,
			Payload:   entity.Payload(m.Payload),
			Reason:    m.Reason,
			CreatedAt: entity.JSONTime(t),
		})
	}
	return fm, nil
}

func init() {
	cert := os.Getenv("GRPC_PUBLIC_KEY")
	cred, err := credentials.NewClientTLSFromFile(cert, "")
	if err != nil {
		log.Fatalf("Could not load tls cert: %s", err)
	}

	address := os.Getenv("GRPC_SERVER_ADDRESS")

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(cred))
	if err != nil {
		log.Fatalf("Could not connet to a grpc server: %v", err)
	}
	//defer conn.Close()
	client := api.NewFailMailClient(conn)
	Server = server{client: client}
}
