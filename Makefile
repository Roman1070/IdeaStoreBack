gen_ideas:
	protoc -I proto proto/ideas.proto --go_out=./gen/go/idea/ --go_opt=paths=source_relative --go-grpc_out=./gen/go/idea/ --go-grpc_opt=paths=source_relative
gen_boards:
	protoc -I proto proto/boards.proto --go_out=./gen/go/boards/ --go_opt=paths=source_relative --go-grpc_out=./gen/go/boards/ --go-grpc_opt=paths=source_relative
gen_profiles:
	protoc -I proto proto/profiles.proto --go_out=./gen/go/profiles/ --go_opt=paths=source_relative --go-grpc_out=./gen/go/profiles/ --go-grpc_opt=paths=source_relative
gen_comments:
	protoc -I proto proto/comments.proto --go_out=./gen/go/comments/ --go_opt=paths=source_relative --go-grpc_out=./gen/go/comments/ --go-grpc_opt=paths=source_relative
gen_chats:
	protoc -I proto proto/chats.proto --go_out=./gen/go/chats/ --go_opt=paths=source_relative --go-grpc_out=./gen/go/chats/ --go-grpc_opt=paths=source_relative
auth:
	go run cmd/auth/main.go
ideas:
	go run cmd/ideas/main.go
boards:
	go run cmd/boards/main.go
profiles:
	go run cmd/profiles/main.go
comments:
	go run cmd/comments/main.go
chats:
	go run cmd/chats/main.go
client:
	go run internal/clients/main.go internal/clients/auth.go internal/clients/ideas.go internal/clients/boards.go internal/clients/profiles.go internal/clients/comments.go internal/clients/chats.go

migrate_auth:
	 migrate -path=migrations/auth -database "postgresql://postgres:postgres@localhost:5432/ideastore?sslmode=disable" -verbose up
migrate_ideas:
	 migrate -path=migrations/ideas -database "postgresql://postgres:postgres@localhost:5432/ideastore?sslmode=disable" -verbose up
migrate_profiles:
	 migrate -path=migrations/profiles -database "postgresql://postgres:postgres@localhost:5432/ideastore?sslmode=disable" -verbose up
migrate_boards:
	migrate -path=migrations/boards -database "postgresql://postgres:postgres@localhost:5432/ideastore?sslmode=disable" -verbose up
migrate_comments:
	migrate -path=migrations/comments -database "postgresql://postgres:postgres@localhost:5432/ideastore?sslmode=disable" -verbose up
migrate_chats:
	migrate -path=migrations/chats -database "postgresql://postgres:postgres@localhost:5432/ideastore?sslmode=disable" -verbose up
migrate_init:
	migrate -path=migrations/ -database "postgresql://postgres:postgres@localhost:5432/ideastore?sslmode=disable" -verbose up
	
migrate_test:
	go build ./cmd/migrator/main.go
	go run ./cmd/migrator/main.go --storage_path=./storage/auth.db --migrations_path=./tests/migrations --migrations_table=migrations_test