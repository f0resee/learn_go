package bidirectional

import (
	"context"
	"io"
	"log"
	"net"
	"strconv"
	"testing"
	"time"

	"gitee.com/souththree/learn_go/src/grpc/bidirectional/proto"
	"google.golang.org/grpc"
)

// https://blog.51cto.com/u_15289640/5840571

const (
	Address = ":8000"
	Network = "tcp"
)

type StreamService struct {
	*proto.UnimplementedStreamServiceServer
}

func (s *StreamService) Record(srv proto.StreamService_RecordServer) error {
	n := 1
	for {
		req, err := srv.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("stream get from client err: %v", err)
			return err
		}

		err = srv.Send(&proto.StreamResponse{
			Code:  int32(n),
			Value: "This is the " + strconv.Itoa(n) + " message",
		})
		if err != nil {
			log.Fatalf("stream send to client err:%v", err)
			return err
		}
		n++
		log.Println("stream get from client: ", req.Data)
		time.Sleep(1 * time.Second)
	}
	return nil
}

func Test_biServer(t *testing.T) {
	listener, err := net.Listen(Network, Address)
	if err != nil {
		log.Fatalf("listener err: %v", err)
	}
	log.Println(Address + " net.Listening...")

	grpcServer := grpc.NewServer()

	proto.RegisterStreamServiceServer(grpcServer, &StreamService{})

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpc server error: %v", err)
	}
}

func Test_biClient(t *testing.T) {
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc conn err: %v", err)
	}
	defer conn.Close()

	grpcClient := proto.NewStreamServiceClient(conn)

	stream, err := grpcClient.Record(context.Background())
	if err != nil {
		log.Fatalf("call record err: %v", err)
	}

	for i := 0; i < 5; i++ {
		err := stream.Send(&proto.StreamRequest{
			Data: strconv.Itoa(i),
		})
		if err != nil {
			log.Fatalf("stream send to server err: %v", err)
		}

		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("stream get from server err: %v", err)
		}
		log.Printf("stream get from server, code: %v.value：%v", resp.GetCode(), resp.GetValue())
		time.Sleep(1 * time.Second)
	}
	err = stream.CloseSend()
	if err != nil {
		log.Fatalf("close stream error:%v", err)
	}
}
