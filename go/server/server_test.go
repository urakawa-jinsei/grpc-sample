package main

import (
	"context"
	"testing"
	"time"

	"github.com/ymmt2005/grpc-tutorial/go/deepthought"
	"google.golang.org/grpc"
)

type mockComputeServer struct {
	bootResponse *deepthought.BootResponse
	bootErr      error
}

func (m *mockComputeServer) Boot(req *deepthought.BootRequest, stream deepthought.Compute_BootServer) error {
	if m.bootErr != nil {
		return m.bootErr
	}

	for {
		select {
		case <-stream.Context().Done():
			return nil
		case <-time.After(1 * time.Second):
		}

		if err := stream.Send(m.bootResponse); err != nil {
			return err
		}
	}
}

type mockBootStream struct {
	grpc.ServerStream
}

func (m *mockBootStream) Send(resp *deepthought.BootResponse) error {
	return nil
}

func (m *mockBootStream) Context() context.Context {
	return context.TODO()
}

func (m *mockBootStream) SendMsg(m_ interface{}) error {
	return nil
}

func (m *mockBootStream) RecvMsg(m_ interface{}) error {
	return nil
}

func TestBoot(t *testing.T) {
	mockServer := &mockComputeServer{
		bootResponse: &deepthought.BootResponse{
			Message:   "Mock Boot Response",
			Timestamp: nil,
		},
		bootErr: nil,
	}

	req := &deepthought.BootRequest{}
	stream := &mockBootStream{}

	err := mockServer.Boot(req, stream)
	if err != nil {
		t.Errorf("Error in Boot: %v", err)
	}
}
