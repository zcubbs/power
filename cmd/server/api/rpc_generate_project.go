package api

import (
	"context"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/zcubbs/power/pkg/designer"
	pb "github.com/zcubbs/power/proto/gen/v1"
	"path/filepath"
)

func (s *Server) GenerateProject(_ context.Context, req *pb.GenerateProjectRequest) (*pb.GenerateProjectResponse, error) {
	var downloadUrl string

	log.Debug("Generating blueprint",
		"blueprint_id", req.BlueprintId,
		"values", req.Values,
	)

	// Generate the project
	err := designer.Generate(req.BlueprintId, req.Values, func(archivePath string) error {
		// Upload & generate download URL
		url, err := s.UploadFileWithPreSignedUrl(s.cfg.S3.BucketName, filepath.Base(archivePath), archivePath)
		if err != nil {
			return fmt.Errorf("failed to upload project to S3 bucket: %v", err)
		}

		// Generate a download URL for the uploaded file
		log.Debug("Uploaded project to S3",
			"bucket", s.cfg.S3.BucketName,
			"object", archivePath,
			"url", downloadUrl,
		)

		downloadUrl = url
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("error generating project: %v", err)
	}

	return &pb.GenerateProjectResponse{DownloadUrl: downloadUrl}, nil
}
