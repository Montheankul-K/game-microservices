infra-up:
	docker-compose -f docker-compose.db.yml up -d

infra-down:
	docker-compose -f docker-compose.db.yml down

protogen-auth:
	protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        ./modules/auth/authPb/authPb.proto

protogen-player:
	protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        ./modules/player/playerPb/playerPb.proto