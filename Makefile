all: build

build:
	sudo docker compose up -d --build

stop:
	sudo docker compose stop

start:
	sudo docker compose start

down:
	sudo docker compose down

create_table:
	sudo docker exec -i issue_tracker_db psql -U task-service -d mydb < migrations/0001_create_issues_table.up.sql

delete_table:
	sudo docker exec -i issue_tracker_db psql -U task-service -d mydb < migrations/0001_create_issues_table.down.sql

check_docker:
	sudo docker ps -a