module msfclient

go 1.18

replace github.com/haochen1204/msf => ../rpc

require github.com/haochen1204/msf v0.0.0

require (
	github.com/golang/protobuf v1.3.1 // indirect
	golang.org/x/net v0.0.0-20190603091049-60506f45cf65 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	gopkg.in/vmihailenco/msgpack.v2 v2.9.2 // indirect
)
