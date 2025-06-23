package service

import (
	"fmt"
	"io"
	"path/filepath"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	config "github.com/cuongpiger/golang/config/server"
	"github.com/cuongpiger/golang/pkg/logger"
	uploadpb "github.com/cuongpiger/golang/proto"
)

type FileServiceServer struct {
	uploadpb.UnimplementedFileServiceServer
	l   *logger.Logger
	cfg *config.Config
}

func New(l *logger.Logger, cfg *config.Config) *FileServiceServer {
	return &FileServiceServer{
		l:   l,
		cfg: cfg,
	}
}

func (g *FileServiceServer) Upload(stream uploadpb.FileService_UploadServer) error {
	file := NewFile()
	var fileSize uint32
	fileSize = 0
	defer func() {
		if err := file.OutputFile.Close(); err != nil {
			g.l.Error(err)
		}
	}()
	for {
		req, err := stream.Recv()
		if file.FilePath == "" {
			file.SetFile(req.GetFileName(), g.cfg.FilesStorage.Location)
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return g.logError(status.Error(codes.Internal, err.Error()))
		}
		chunk := req.GetChunk()
		fileSize += uint32(len(chunk))
		g.l.Debug("received a chunk with size: %d", fileSize)
		if err := file.Write(chunk); err != nil {
			return g.logError(status.Error(codes.Internal, err.Error()))
		}
	}

	fmt.Println(file.FilePath, fileSize)
	fileName := filepath.Base(file.FilePath)
	g.l.Debug("saved file: %s, size: %d", fileName, fileSize)
	return stream.SendAndClose(&uploadpb.FileUploadResponse{FileName: fileName, Size: fileSize})
}
