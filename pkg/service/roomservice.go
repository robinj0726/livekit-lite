package service

import (
	"context"

	"github.com/livekit/protocol/livekit"
	"github.com/pkg/errors"
)

type RoomService struct {
	roomAllocator RoomAllocator
}

func NewRoomService(ra RoomAllocator) (svc *RoomService, err error) {
	svc = &RoomService{
		roomAllocator: ra,
	}

	return
}

func (s *RoomService) CreateRoom(ctx context.Context, req *livekit.CreateRoomRequest) (rm *livekit.Room, err error) {
	rm, err = s.roomAllocator.CreateRoom(ctx, req)
	if err != nil {
		err = errors.Wrap(err, "could not create room")
	}

	return
}

func (s *RoomService) ListRooms(ctx context.Context, req *livekit.ListRoomsRequest) (res *livekit.ListRoomsResponse, err error) {
	return
}

func (s *RoomService) DeleteRoom(ctx context.Context, req *livekit.DeleteRoomRequest) (*livekit.DeleteRoomResponse, error) {
	return &livekit.DeleteRoomResponse{}, nil
}

func (s *RoomService) ListParticipants(ctx context.Context, req *livekit.ListParticipantsRequest) (res *livekit.ListParticipantsResponse, err error) {
	return
}

func (s *RoomService) GetParticipant(ctx context.Context, req *livekit.RoomParticipantIdentity) (res *livekit.ParticipantInfo, err error) {
	return
}

func (s *RoomService) RemoveParticipant(ctx context.Context, req *livekit.RoomParticipantIdentity) (res *livekit.RemoveParticipantResponse, err error) {
	return
}

func (s *RoomService) MutePublishedTrack(ctx context.Context, req *livekit.MuteRoomTrackRequest) (res *livekit.MuteRoomTrackResponse, err error) {
	return
}

func (s *RoomService) UpdateParticipant(ctx context.Context, req *livekit.UpdateParticipantRequest) (*livekit.ParticipantInfo, error) {
	return &livekit.ParticipantInfo{}, nil
}

func (s *RoomService) UpdateSubscriptions(ctx context.Context, req *livekit.UpdateSubscriptionsRequest) (*livekit.UpdateSubscriptionsResponse, error) {
	return &livekit.UpdateSubscriptionsResponse{}, nil
}

func (s *RoomService) SendData(ctx context.Context, req *livekit.SendDataRequest) (*livekit.SendDataResponse, error) {
	return &livekit.SendDataResponse{}, nil
}

func (s *RoomService) UpdateRoomMetadata(ctx context.Context, req *livekit.UpdateRoomMetadataRequest) (*livekit.Room, error) {
	return &livekit.Room{}, nil
}
