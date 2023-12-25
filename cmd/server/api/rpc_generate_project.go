package api

import (
	"context"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/zcubbs/power/pkg/blueprint"
	"github.com/zcubbs/power/pkg/designer"
	pb "github.com/zcubbs/power/proto/gen/v1"
	"os"
	"path/filepath"
	"time"
)

func (s *Server) GenerateProject(_ context.Context, req *pb.GenerateProjectRequest) (*pb.GenerateProjectResponse, error) {
	// Convert the request options to a map[string]interface{} for the designer package
	options := make(map[string]interface{})
	for k, v := range req.Options {
		options[k] = v
	}

	// Create a ProjectSpec
	spec := designer.ProjectSpec{
		Components: []blueprint.ComponentSpec{
			{
				Type:   req.Blueprint,
				Config: options,
			},
		},
	}

	// Define the output path for the generated project
	// os tmp dir
	outputPath := filepath.Join(os.TempDir(), "power", req.Blueprint, time.Now().Format("20060102150405"))

	// Generate the project
	if err := designer.GenerateProject(&spec, outputPath); err != nil {
		return nil, fmt.Errorf("failed to generate project: %v", err)
	}

	log.Debug("Generated project", "outputPath", outputPath)

	zipFilePath := filepath.Join(outputPath, "project.zip")

	// Upload the generated project to MinIO
	objectName := fmt.Sprintf("%s-%s.zip", req.Blueprint, time.Now().Format("20060102150405"))
	_, err := s.s3Client.UploadFile(s.cfg.S3.BucketName, objectName, zipFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to upload project to MinIO: %v", err)
	}

	downloadUrl, err := s.s3Client.GetDownloadURL(s.cfg.S3.BucketName, objectName, 10*time.Minute, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get download url: %v", err)
	}

	// Generate a download URL for the uploaded file
	log.Debug("Uploaded project to S3", "bucket", s.cfg.S3.BucketName, "object", objectName)
	log.Debug("Generated download URL", "url", downloadUrl)

	// clean up temp dir
	err = os.RemoveAll(outputPath)
	if err != nil {
		log.Error("Failed to remove temp dir",
			"package", "api",
			"function", "GenerateProject",
			"error", err,
			"path", outputPath,
		)

		// return grpc internal error
		return nil, fmt.Errorf(InternalRedactedError)
	}

	return &pb.GenerateProjectResponse{DownloadUrl: downloadUrl.String()}, nil
}
