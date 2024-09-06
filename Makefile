generate:
	protoc -I ./a-study-in-history-spec/api/proto/v1 \
	--go_out=./a-study-in-history-spec \
	--go-grpc_out=./a-study-in-history-spec \
	./a-study-in-history-spec/api/proto/v1/events.proto
	protoc -I ./a-study-in-history-spec/api/proto/v1 \
	--go_out=./a-study-in-history-spec \
	--go-grpc_out=./a-study-in-history-spec \
	./a-study-in-history-spec/api/proto/v1/a_study_in_history.proto
lint:
	gofumpt -w .
	golangci-lint cache clean
	golangci-lint run ./a-study-in-history-daemon/... ./a-study-in-history-api/... ./a-study-in-history-spec/... ./a-study-in-history-bot/....