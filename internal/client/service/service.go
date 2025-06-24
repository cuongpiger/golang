package service

import (
	"context"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"

	uploadpb "github.com/cuongpiger/golang/proto"
)

type ClientService struct {
	addr      string
	filePath  string
	batchSize int
	client    uploadpb.FileServiceClient
}

func New(addr string, filePath string, batchSize int) *ClientService {
	return &ClientService{
		addr:      addr,
		filePath:  filePath,
		batchSize: batchSize,
	}
}

func (s *ClientService) SendFile(method string) error {
	log.Println(s.addr, s.filePath)
	conn, err := grpc.Dial(s.addr, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()
	s.client = uploadpb.NewFileServiceClient(conn)
	interrupt := make(chan os.Signal, 1)
	shutdownSignals := []os.Signal{
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGQUIT,
	}
	signal.Notify(interrupt, shutdownSignals...)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func(s *ClientService) {
		switch method {
		case "file":
			if err = s.upload(ctx, cancel); err != nil {
				log.Fatal(err)
				cancel()
			}
		case "memory":
			if err = s.uploadMemory(ctx, cancel); err != nil {
				log.Fatal(err)
				cancel()
			}
		default:
			log.Fatalf("Unknown method: %s. Use 'file' or 'memory'.", method)
			cancel()
		}

	}(s)

	select {
	case killSignal := <-interrupt:
		log.Println("Got ", killSignal)
		cancel()
	case <-ctx.Done():
	}
	return nil
}

func (s *ClientService) upload(ctx context.Context, cancel context.CancelFunc) error {
	stream, err := s.client.Upload(ctx)
	if err != nil {
		return err
	}
	file, err := os.Open(s.filePath)
	if err != nil {
		return err
	}
	buf := make([]byte, s.batchSize)
	batchNumber := 1
	for {
		num, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		chunk := buf[:num]

		if err := stream.Send(&uploadpb.FileUploadRequest{FileName: s.filePath, Chunk: chunk}); err != nil {
			return err
		}
		log.Printf("Sent - batch #%v - size - %v\n", batchNumber, len(chunk))
		batchNumber += 1

	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}
	log.Printf("Sent - %v bytes - %s\n", res.GetSize(), res.GetFileName())
	cancel()
	return nil
}

func (s *ClientService) uploadMemory(ctx context.Context, cancel context.CancelFunc) error {
	stream, err := s.client.UploadOnMemory(ctx)
	if err != nil {
		return err
	}
	file, err := os.Open(s.filePath)
	if err != nil {
		return err
	}
	buf := make([]byte, s.batchSize)
	batchNumber := 1
	for {
		num, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		chunk := buf[:num]

		if err := stream.Send(&uploadpb.FileUploadRequest{FileName: s.filePath, Chunk: chunk}); err != nil {
			return err
		}
		log.Printf("Sent - batch #%v - size - %v\n", batchNumber, len(chunk))
		batchNumber += 1

	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}
	log.Printf("Sent - %v bytes - %s\n", res.GetSize(), res.GetFileName())
	cancel()
	return nil
}
