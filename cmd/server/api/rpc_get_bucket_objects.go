package api

import (
	"context"
	"fmt"
	pb "github.com/zcubbs/power/proto/gen/v1"
)

func (s *Server) GetBucketObjects(_ context.Context, _ *pb.GetBucketObjectListRequest) (*pb.GetBucketObjectListResponse, error) {
	// Get the list of objects in blueprints bucket
	listChan := s.s3Client.ListObjects(s.cfg.S3.BucketName)

	objects := make([]string, 0)
	for object := range listChan {
		if object.Err != nil {
			return nil, fmt.Errorf("failed to list objects: %v", object.Err)
		}
		objects = append(objects, fmt.Sprintf("%s", object.Key))
	}

	return &pb.GetBucketObjectListResponse{Objects: objects}, nil
}
