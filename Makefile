.DEFAULT: protos

protos: generate_proto

generate_proto:
	sh scripts/generate_proto.sh  "go" "./" "./" "github.com/pjoc-team/timer-proto"

deps:
	sh scripts/download_dependencies.sh