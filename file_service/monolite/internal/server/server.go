package server

import (
	"context"
	"io"
	"log/slog"
	"path/filepath"

	"github.com/Heatdog/FileServer/file_service/monolite/internal/config"
	"github.com/Heatdog/FileServer/file_service/monolite/internal/file"
	"github.com/Heatdog/FileServer/file_service/monolite/proto"
)

type FileServiceServer struct {
	proto.UnimplementedFileServiceServer
	logger *slog.Logger
	cfg    *config.Config
}

func (server *FileServiceServer) SetFile(ctx context.Context, stream proto.FileService_SetFileServer) error {
	file := file.NewFile()
	var fileSize uint64 = 0
	defer func() {
		if err := file.OutputFile.Close(); err != nil {
			server.logger.Error(err.Error())
		}
	}()

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			server.logger.Info(err.Error())
			break
		}
		if err != nil {
			server.logger.Error(err.Error())
			return err
		}

		if file.FilePath == "" {
			err = file.SetFile(req.GetUuid(), server.cfg.FileStorage.Location, server.logger)
			if err != nil {
				server.logger.Error(err.Error())
				return err
			}
		}

		chunk := req.GetChunk()
		fileSize += uint64(len(chunk))

		server.logger.Info("received a chunk with size", slog.Uint64("size", fileSize))
		if err := file.Write(chunk); err != nil {
			server.logger.Error(err.Error())
			return err
		}
	}

	fileName := filepath.Base(file.FilePath)
	server.logger.Info("saved file", slog.String("file_name", fileName), slog.Uint64("size", fileSize))
	return stream.SendAndClose(&proto.SetFileResponse{
		Uuid: fileName,
		Size: fileSize,
	})
}
