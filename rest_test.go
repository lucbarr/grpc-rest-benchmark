package benchmark

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/roboime/grpc-rest-benchmark/protos/ssl"

	"github.com/roboime/grpc-rest-benchmark/rest"
)

func init() {
	go rest.Start()
	time.Sleep(time.Second / 2)
}

func BenchmarkREST(b *testing.B) {
	client := &http.Client{}

	for n := 0; n < b.N; n++ {
		doPost(client, b)
	}
}

func doPost(client *http.Client, b *testing.B) {
	buf := new(bytes.Buffer)
	bytes, _ := json.Marshal(&ssl.RefereeRequest{
		LastPacketTimestamp: 100,
	})
	buf.Write(bytes)

	resp, err := client.Post("http://127.0.0.1:10006/", "application/json", buf)
	if err != nil {
		b.Fatalf("http request failed: %v", err)
	}
	defer resp.Body.Close()

	ref := &ssl.SSL_Referee{}
	bytes, _ = ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bytes, ref)

	if err != nil {
		b.Fatalf("unable to decode json: %v", err)
	}
}
