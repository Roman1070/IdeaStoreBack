gen_ideas:
	protoc -I proto proto/ideas.proto --go_out=./gen/go/idea/ --go_opt=paths=source_relative --go-grpc_out=./gen/go/idea/ --go-grpc_opt=paths=source_relative
gen_boards:
	protoc -I proto proto/boards.proto --go_out=./gen/go/boards/ --go_opt=paths=source_relative --go-grpc_out=./gen/go/boards/ --go-grpc_opt=paths=source_relative
gen_profiles:
	protoc -I proto proto/profiles.proto --go_out=./gen/go/profiles/ --go_opt=paths=source_relative --go-grpc_out=./gen/go/profiles/ --go-grpc_opt=paths=source_relative
auth_server:
	go run cmd/auth/main.go
ideas_server:
	go run cmd/ideas/main.go
boards_server:
	go run cmd/boards/main.go
profiles_server:
	go run cmd/profiles/main.go
client:
	go run internal/clients/main.go internal/clients/auth.go internal/clients/ideas.go internal/clients/boards.go
migrate_ideas:
	go build ./cmd/migrator/main.go
	go run ./cmd/migrator/main.go --storage_path=./storage/ideas.db --migrations_path=./migrations/ideas
migrate_profiles:
	go build ./cmd/migrator/main.go
	go run ./cmd/migrator/main.go --storage_path=./storage/profiles.db --migrations_path=./migrations/profiles
	
migrate_boards:
	go build ./cmd/migrator/main.go
	go run ./cmd/migrator/main.go --storage_path=./storage/boards.db --migrations_path=./migrations/boards

	
migrate_test:
	go build ./cmd/migrator/main.go
	go run ./cmd/migrator/main.go --storage_path=./storage/auth.db --migrations_path=./tests/migrations --migrations_table=migrations_test