gen_ideas:
	protoc -I proto proto/ideas.proto --go_out=./gen/go/idea/ --go_opt=paths=source_relative --go-grpc_out=./gen/go/idea/ --go-grpc_opt=paths=source_relative
auth_server:
	go run cmd/auth/main.go
ideas_server:
	go run cmd/ideas/main.go
auth_client:
	go run internal/clients/auth/main.go
ideas_client:
	go run internal/clients/ideas/main.go
migrate_ideas:
	go build ./cmd/migrator/main.go
	go run ./cmd/migrator/main.go --storage_path=./storage/ideas.db --migrations_path=./migrations/ideas

	
migrate_test:
	go build ./cmd/migrator/main.go
	go run ./cmd/migrator/main.go --storage_path=./storage/auth.db --migrations_path=./tests/migrations --migrations_table=migrations_test