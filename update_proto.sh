#!/usr/bin/bash
#Shell commands to update protobuff files used in the project
protoc -I ./proto --go_out=plugins=grpc:players/proto players.proto
protoc -I ./proto --go_out=plugins=grpc:levels/proto levels.proto