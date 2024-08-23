package object

import (
	"context"
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/Mungrel/objects/internal/db/repo"
	"github.com/Mungrel/objects/internal/types"
)

type Server struct {
	store *repo.ObjectStore
}

var _ ObjectServer = Server{}

// CreateObject implements ObjectServer.
func (s Server) CreateObject(ctx context.Context, request *CreateObjectRequest) (*CreateObjectResponse, error) {
	now := time.Now().UTC()

	object := types.Object{
		ID:        uuid.NewString(),
		Name:      request.GetName(),
		Content:   request.GetContent(),
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := s.store.Create(ctx, object); err != nil {
		return nil, err
	}

	return &CreateObjectResponse{
		Object: toPB(object),
	}, nil
}

// ListObjects implements ObjectServer.
func (s Server) ListObjects(ctx context.Context, request *ListObjectsRequest) (*ListObjectsResponse, error) {
	objects, err := s.store.List(ctx)
	if err != nil {
		return nil, err
	}

	resp := &ListObjectsResponse{
		Objects: make([]*ObjectRecord, len(objects)),
	}

	for i := range objects {
		resp.Objects[i] = toPB(objects[i])
	}

	return resp, nil
}

func toPB(object types.Object) *ObjectRecord {
	return &ObjectRecord{
		Id:        object.ID,
		Name:      object.Name,
		Content:   object.Content,
		CreatedAt: timestamppb.New(object.CreatedAt),
		UpdatedAt: timestamppb.New(object.UpdatedAt),
	}
}
