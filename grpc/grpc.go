package grpc

import (
	"context"
	"fmt"
	"io"
	"log"
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
	var fms []entity.FailMail
	for {
		m, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("%v.GetFailMails(_) = _, %v", s.client, err)
		}

		fm := entity.FailMail{
			ID:        m.ID,
			Action:    m.Action,
			Payload:   m.Payload,
			Reason:    m.Reason,
			CreatedAt: m.CreatedAt,
		}

		if m.DeletedAt != "" {
			fm.DeletedAt = &m.DeletedAt
		}

		fms = append(fms, fm)
	}
	return fms, nil
}

func (s server) Count(ctx context.Context) (int64, error) {
	res, err := s.client.CountFailMails(ctx, &api.Empty{})
	if err != nil {
		return 0, fmt.Errorf("could get count on from grpc: %v", err)
	}

	return res.Total, nil
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
	// TODO: is it ok that we don't close the grpc connection?
	//defer conn.Close()
	client := api.NewFailMailClient(conn)
	Server = server{client: client}
}
