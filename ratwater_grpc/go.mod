module ratwater.xyz/grpc/ratwater

go 1.21.6

replace ratwater.xyz/mod/ratwater => ../ratwater_mod

require (
	google.golang.org/grpc v1.64.0
	ratwater.xyz/grpc/proto/ratwater v0.0.0-00010101000000-000000000000
)

require (
	golang.org/x/net v0.22.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240318140521-94a12d6c2237 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
)

replace ratwater.xyz/grpc/proto/ratwater => ./proto
