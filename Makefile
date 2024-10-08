up:
	 cd deployments && docker compose up -d

deploy:
	make test
	cd deployments && docker compose down
	cd deployments && docker compose up --build -d

down:
	cd deployments && docker compose down

test:
	go test ./...
