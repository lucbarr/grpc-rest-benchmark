package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/roboime/grpc-rest-benchmark/protos/ssl"
)

// Start entrypoint
func Start() {
	http.HandleFunc("/", GetRefereePacket)
	log.Println(http.ListenAndServe(":10006", nil))
}

// GetRefereePacket handler
func GetRefereePacket(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req ssl.RefereeRequest
	decoder.Decode(&req)
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	responseData, _ := json.Marshal(
		&ssl.SSL_Referee{
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
		})
	fmt.Fprintf(w, string(responseData))
}
