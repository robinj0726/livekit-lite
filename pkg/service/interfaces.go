package service

import (
	"context"

	"github.com/livekit/protocol/livekit"
)

type RoomAllocator interface {
	CreateRoom(ctx context.Context, req *livekit.CreateRoomRequest) (*livekit.Room, error)
}
