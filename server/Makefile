test:
	python ./test/system.py

up:
	podman-compose --env-file ./.env up -d

down:
	podman-compose --env-file ./.env down

re: down up

ps:
	podman-compose --env-file ./.env ps

logs:
	podman-compose --env-file ./.env logs
	
flogs:
	podman-compose --env-file ./.env logs -f

.PHONY: test up down re ps logs flogs
