package handler

import (
	"google.golang.org/protobuf/proto"
	"log"
	"net/http"

	"github.com/353solutions/weather/pb"
	"github.com/353solutions/weather/weather"
)

func protoSend(w http.ResponseWriter, rec weather.Record) {
	pbr := pb.Record{
		Time:    pb.Timestamp(rec.Time),
		Station: rec.Station,
		Temperature: &pb.Value{
			Value: rec.Temperature.Value,
			Unit:  string(rec.Temperature.Unit),
		},
		Rain: &pb.Value{
			Value: rec.Rain.Value,
			Unit:  string(rec.Rain.Unit),
		},
	}

	data, err := proto.Marshal(&pbr)
	if err != nil {
		log.Printf("can't marshal protobuf: %#v - %s", &pbr, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("json data size: %d", len(data))
	w.Header().Set("Content-Type", protoCtype)
	_, err = w.Write(data)
	if err != nil {
		log.Printf("protobuf can't write: %#v - %s", data, err)
	}
}
