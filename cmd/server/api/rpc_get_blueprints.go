package api

import (
	"context"
	"github.com/zcubbs/power/blueprint"
	pb "github.com/zcubbs/power/proto/gen/v1"
)

func (s *Server) GetBlueprints(_ context.Context, req *pb.GetBlueprintListRequest) (*pb.GetBlueprintListResponse, error) {
	// Get the list of registered blueprints
	blueprints := make([]*pb.Blueprint, 0)
	for _, spec := range blueprint.GetAllBlueprintSpecs() {
		blueprints = append(blueprints, &pb.Blueprint{
			Spec: toSpecPb(spec),
		})
	}

	return &pb.GetBlueprintListResponse{Blueprints: blueprints}, nil
}

func toSpecPb(spec *blueprint.Spec) *pb.Spec {
	options := make([]*pb.Option, 0)
	for _, option := range spec.Options {
		options = append(options, &pb.Option{
			Id:          option.ID,
			Name:        option.Name,
			Type:        option.Type,
			Description: option.Description,
			Options:     option.Choices,
			Default:     option.Default,
		})
	}

	return &pb.Spec{
		Id:          spec.ID,
		Name:        spec.Name,
		Description: spec.Description,
		Options:     options,
	}
}
