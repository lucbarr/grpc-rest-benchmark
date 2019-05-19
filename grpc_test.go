package benchmark

import (
	"context"
	"testing"
	"time"

	"github.com/roboime/grpc-rest-benchmark/protos/ssl"

	"github.com/roboime/grpc-rest-benchmark/grpc"
	g "google.golang.org/grpc"
)

func init() {
	go grpc.StartServer()
	time.Sleep(time.Second / 2)
}

func BenchmarkGRPC(b *testing.B) {
	conn, _ := g.Dial("127.0.0.1:10005", g.WithInsecure())

	client := ssl.NewRefereeClient(conn)

	for i := 0; i < b.N; i++ {
		client.GetRefereePacket(context.Background(), &ssl.RefereeRequest{
			LastPacketTimestamp: 100,
		})
	}
}
