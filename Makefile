infra-up:
	docker-compose -f docker-compose.db.yml up -d

infra-down:
	docker-compose -f docker-compose.db.yml down

run-auth:
	go run main.go ./env/dev/.env.auth

run-player:
	go run main.go ./env/dev/.env.player

run-item:
	go run main.go ./env/dev/.env.item

run-inventory:
	go run main.go ./env/dev/.env.inventory

run-payment:
	go run main.go ./env/dev/.env.payment

protogen-auth:
	protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        ./modules/auth/authPb/authPb.proto

protogen-player:
	protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        ./modules/player/playerPb/playerPb.proto

protogen-item:
	protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        ./modules/item/itemPb/itemPb.proto

protogen-inventory:
	protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        ./modules/inventory/inventoryPb/inventoryPb.proto