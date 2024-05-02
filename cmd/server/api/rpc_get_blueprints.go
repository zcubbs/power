package api

import (
	"context"
	"github.com/zcubbs/blueprint"
	pb "github.com/zcubbs/power/proto/gen/v1"
)

func (s *Server) GetBlueprints(_ context.Context, req *pb.GetBlueprintListRequest) (*pb.GetBlueprintListResponse, error) {
	// Get the list of registered blueprints
	blueprints := make([]*pb.Blueprint, 0)
	for _, bpt := range blueprint.GetAllBlueprints() {
		blueprints = append(blueprints, &pb.Blueprint{
			Spec: toSpecPb(bpt.Spec),
			Type: string(bpt.Type),
		})
	}

	return &pb.GetBlueprintListResponse{Blueprints: blueprints}, nil
}

func toSpecPb(spec blueprint.Spec) *pb.Spec {
	options := make([]*pb.Option, 0)
	for _, option := range spec.Options {
		options = append(options, &pb.Option{
			Id:          option.ID,
			Name:        option.Name,
			Type:        option.Type,
			Description: option.Description,
			Choices:     option.Choices,
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
