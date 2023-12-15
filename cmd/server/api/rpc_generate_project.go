package api

import (
	"context"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/zcubbs/power/blueprint"
	"github.com/zcubbs/power/designer"
	pb "github.com/zcubbs/power/proto/gen/v1"
	"os"
	"path/filepath"
	"time"
)

func (s *Server) GenerateProject(ctx context.Context, req *pb.GenerateProjectRequest) (*pb.GenerateProjectResponse, error) {
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

	// TODO: Create a URL or some mechanism to access the generated project
	downloadURL := "http://example.com/download/generated/project.zip" // Update this URL as needed

	return &pb.GenerateProjectResponse{DownloadUrl: downloadURL}, nil
}
