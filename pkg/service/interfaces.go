package service

import (
	"context"
	"time"

	"github.com/livekit/protocol/livekit"
)

type ObjectStore interface {
	ServiceStore

	// enable locking on a specific room to prevent race
	// returns a (lock uuid, error)
	LockRoom(ctx context.Context, name livekit.RoomName, duration time.Duration) (string, error)
	UnlockRoom(ctx context.Context, name livekit.RoomName, uid string) error

	StoreRoom(ctx context.Context, room *livekit.Room) error
	DeleteRoom(ctx context.Context, name livekit.RoomName) error

	StoreParticipant(ctx context.Context, roomName livekit.RoomName, participant *livekit.ParticipantInfo) error
	DeleteParticipant(ctx context.Context, roomName livekit.RoomName, identity livekit.ParticipantIdentity) error
}

type ServiceStore interface {
	LoadRoom(ctx context.Context, name livekit.RoomName) (*livekit.Room, error)
	// ListRooms returns currently active rooms. if names is not nil, it'll filter and return
	// only rooms that match
	ListRooms(ctx context.Context, names []livekit.RoomName) ([]*livekit.Room, error)

	LoadParticipant(ctx context.Context, roomName livekit.RoomName, identity livekit.ParticipantIdentity) (*livekit.ParticipantInfo, error)
	ListParticipants(ctx context.Context, roomName livekit.RoomName) ([]*livekit.ParticipantInfo, error)
}

type RoomAllocator interface {
	CreateRoom(ctx context.Context, req *livekit.CreateRoomRequest) (*livekit.Room, error)
}
