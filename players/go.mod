module obscura-players-backend

go 1.12

require (
	google.golang.org/grpc v1.21.1
	obscura-proto v0.0.0
)

replace obscura-proto => ../proto
