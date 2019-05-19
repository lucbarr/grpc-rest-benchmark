package grpc

import (
	"context"
	"log"
	"net"

	"github.com/lucbarr/protos/ssl"
	"google.golang.org/grpc"
)

//StartServer entry
func StartServer() {
	l, _ := net.Listen("tcp", ":10005")

	server := grpc.NewServer()

	ssl.RegisterRefereeServer(server, &RefereeServer{})
	log.Println(server.Serve(l))
}

//RefereeServer type
type RefereeServer struct{}

//GetRefereePacket handler
func (s *RefereeServer) GetRefereePacket(ctx context.Context, req *ssl.RefereeRequest) (*ssl.SSL_Referee, error) {
	return &ssl.SSL_Referee{
		Command:          ssl.SSL_Referee_Command(0),
		CommandCounter:   10,
		CommandTimestamp: req.LastPacketTimestamp + 1,
		PacketTimestamp:  req.LastPacketTimestamp + 1,
		Stage:            ssl.SSL_Referee_Stage(10),
		StageTimeLeft:    10,

		Blue: &ssl.SSL_Referee_TeamInfo{
			Goalie:      0,
			Name:        "Warthog Robotics",
			RedCards:    1,
			Score:       0,
			TimeoutTime: 300,
			Timeouts:    2,
		},
		Yellow: &ssl.SSL_Referee_TeamInfo{
			Goalie:      0,
			Name:        "RoboIME",
			RedCards:    0,
			Score:       10,
			TimeoutTime: 30,
			Timeouts:    3,
		},
	}, nil

}
