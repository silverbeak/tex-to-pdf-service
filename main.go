package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/silverbeak/tex-to-pdf-service/com_sootsafe_latex"
	generator "github.com/silverbeak/tex-to-pdf-service/service"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	serv := grpc.NewServer()
	pb.RegisterTexToPdfServer(serv, &server{})

	reflection.Register(serv)
	if err := serv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

type server struct{}

func (s *server) Convert(ctx context.Context, in *pb.TexMessage) (*pb.PdfResponse, error) {
	pdf, err := generator.Generate(in.GetTex(), "temp", true)

	if err != nil {
		errorMsg := &pb.ErrorMessage{ErrorCode: 500, ErrorMessage: fmt.Sprintf("Error: %v", err)}
		return &pb.PdfResponse{Error: errorMsg, Pdf: nil}, nil
	}

	return &pb.PdfResponse{Error: nil, Pdf: pdf}, nil
}
