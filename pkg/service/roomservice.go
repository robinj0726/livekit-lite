package service

import (
	"context"

	"github.com/livekit/protocol/livekit"
	"github.com/pkg/errors"
)

type RoomService struct {
	roomAllocator RoomAllocator
	roomStore     ServiceStore
}

func NewRoomService(ra RoomAllocator, rs ServiceStore) (svc *RoomService, err error) {
	svc = &RoomService{
		roomAllocator: ra,
		roomStore:     rs,
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
	var names []livekit.RoomName
	if len(req.Names) > 0 {
		names = livekit.StringsAsRoomNames(req.Names)
	}
	rooms, err := s.roomStore.ListRooms(ctx, names)
	if err != nil {
		// TODO: translate error codes to twirp
		return
	}

	res = &livekit.ListRoomsResponse{
		Rooms: rooms,
	}
	return
}

func (s *RoomService) DeleteRoom(ctx context.Context, req *livekit.DeleteRoomRequest) (*livekit.DeleteRoomResponse, error) {
	return &livekit.DeleteRoomResponse{}, nil
}

func (s *RoomService) ListParticipants(ctx context.Context, req *livekit.ListParticipantsRequest) (res *livekit.ListParticipantsResponse, err error) {
	participants, err := s.roomStore.ListParticipants(ctx, livekit.RoomName(req.Room))
	if err != nil {
		return
	}

	res = &livekit.ListParticipantsResponse{
		Participants: participants,
	}
	return
}

func (s *RoomService) GetParticipant(ctx context.Context, req *livekit.RoomParticipantIdentity) (res *livekit.ParticipantInfo, err error) {
	participant, err := s.roomStore.LoadParticipant(ctx, livekit.RoomName(req.Room), livekit.ParticipantIdentity(req.Identity))
	if err != nil {
		return
	}

	res = participant
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
