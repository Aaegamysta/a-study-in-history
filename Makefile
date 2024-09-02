generate:
	protoc -I ./a-study-in-history-spec/api \
	--go_out=./a-study-in-history-spec/gen \
	--go-grpc_out=./a-study-in-history-spec/gen \
	./a-study-in-history-spec/api/events.proto
	protoc -I ./a-study-in-history-spec/api \
	--go_out=./a-study-in-history-spec/gen \
	--go-grpc_out=./a-study-in-history-spec/gen \
	./a-study-in-history-spec/api/a_study_in_history.proto
	