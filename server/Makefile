include .env

test:
	python ./test/system.py

################### PODMAN #####################
up:
	podman-compose --env-file ./.env up -d

down:
	podman-compose --env-file ./.env down

re: down up

ps:
	podman-compose --env-file ./.env ps

### Logs

logs:
	podman-compose --env-file ./.env logs
	
flogs:
	podman-compose --env-file ./.env logs -f

### Clean

vclean:
	podman-compose down -v

fclean: down
	podman system prune -af --volumes

################### Postgres #####################

psql:
	psql -h 127.0.0.1 -p $(DB_EXT) -U $(DB_USER) -d $(DB_NAME)

.PHONY: test up down re ps logs flogs vclean fclean psql
