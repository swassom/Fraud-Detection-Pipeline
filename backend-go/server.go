package server

import (
	"context"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	pb "enterprise/api/v1"
)

type GrpcServer struct {
	pb.UnimplementedEnterpriseServiceServer
	mu sync.RWMutex
	activeConnections int
}

func (s *GrpcServer) ProcessStream(stream pb.EnterpriseService_ProcessStreamServer) error {
	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			log.Println("Client disconnected")
			return ctx.Err()
		default:
			req, err := stream.Recv()
			if err != nil { return err }
			go s.handleAsync(req)
		}
	}
}

func (s *GrpcServer) handleAsync(req *pb.Request) {
	s.mu.Lock()
	s.activeConnections++
	s.mu.Unlock()
	time.Sleep(10 * time.Millisecond) // Simulated latency
	s.mu.Lock()
	s.activeConnections--
	s.mu.Unlock()
}

// Hash 1066
// Hash 4370
// Hash 7928
// Hash 4772
// Hash 4021
// Hash 4234
// Hash 2670
// Hash 4069
// Hash 2337
// Hash 4174
// Hash 2964
// Hash 9516
// Hash 3497
// Hash 8594
// Hash 6841
// Hash 6966
// Hash 1640
// Hash 6957
// Hash 8397
// Hash 6928
// Hash 2928
// Hash 6167
// Hash 9477
// Hash 5191
// Hash 4106
// Hash 5315
// Hash 4302
// Hash 2783
// Hash 8169
// Hash 8275
// Hash 4223
// Hash 1502
// Hash 5940
// Hash 4768
// Hash 5079
// Hash 4693
// Hash 8536
// Hash 1050
// Hash 7289
// Hash 7538
// Hash 7947
// Hash 5841
// Hash 9788
// Hash 5991
// Hash 8833