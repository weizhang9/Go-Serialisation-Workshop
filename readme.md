### Go Serialization Workshop

This repo corresponds to Go Serialisation Workshop by Miki Tebeka which can be found at https://www.353solutions.com/c/go-serialize-ggc/

The main serialisation workshop code is in this repo. You can find the code on the website above too; I just reorganised the file structure to better reflect the relationship in my opinion, along with some extra notes taken during the workshop session.

There are extra bits of code snippets that's not part of the main workshop task, but can be found below:
#### Anonymous sturct for parsing JSON
[code snippet](https://www.353solutions.com/c/go-serialize-ggc/html/stocks.go.html)

#### Missing vs zero value
[code snippet](https://www.353solutions.com/c/go-serialize-ggc/html/job.go.html)

#### Streaming JSON
[code snippet](https://www.353solutions.com/c/go-serialize-ggc/html/jser.go.html)


### Run the project
```
go run ./cmd/weatherd
```
The auto-generated protobuf code is committed, but can be removed (`weather/weather.pb.go`) and re-generated. You would need to 
- install protoc `brew install protobuf`
- and proto-go-gen `go install google.golang.org/protobuf/cmd/protoc-gen-go`
- and add it to your $PATH `export PATH="${PATH}:${HOME}/go/bin"`
- then generate with command `protoc --go_out pb --go_opt=paths=source_relative weather.proto`

### Run the benchmark test
```
cd weather && go test -bench . -benchmem
```