package weather

import (
	"encoding/json"
	"testing"

	"github.com/353solutions/weather/pb"
	"google.golang.org/protobuf/proto"
)

var (
	rec = db[2]
)

func BenchmarkJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(rec)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// use real data for benchmark for more accurate results
func BenchmarkProtobuf(b *testing.B) {
	for i := 0; i < b.N; i++ {
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
		_, err := proto.Marshal(&pbr)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// result
// ヾ(✿ ◔ ڼ ◔ )ノ go test -bench . -benchmem
// goos: darwin
// goarch: amd64
// pkg: github.com/353solutions/weather/weather
// BenchmarkJSON-12                  432090              2839 ns/op             400 B/op         13 allocs/op
// BenchmarkProtobuf-12             1991380               628 ns/op             336 B/op          5 allocs/op
