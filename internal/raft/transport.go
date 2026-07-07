package raft

import (
	context "context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Transport struct {
	peerMap map[int]string
}

func (t *Transport) SendRequestVote(peerID int, args *RequestVoteArgs) (*RequestVoteReply, error) {
	ctx := context.Background()
	address, ok := t.peerMap[peerID]
	if !ok {
		return nil, fmt.Errorf("no address for peer %d", peerID)
	}

	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := NewRaftClient(conn)
	reply, err := client.RequestVote(ctx, args)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
