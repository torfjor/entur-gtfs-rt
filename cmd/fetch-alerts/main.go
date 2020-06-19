package main

import (
	pb "entur-gtfs-rt/proto"
	"fmt"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	alerts := pb.FeedMessage{}
	res, err := http.Get("https://api.entur.io/realtime/v1/gtfs-rt/alerts")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = res.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	body, err := ioutil.ReadAll(res.Body)
	err = proto.Unmarshal(body, &alerts)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range alerts.GetEntity() {
		fmt.Println(v.Alert)
	}
}
