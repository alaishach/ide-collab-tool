include .env

AIR_TOML=./air.toml

test:
	python ./test/system.py

################### PODMAN #####################

up:
	podman-compose --env-file ./.env up -d
	$(MAKE) flogs

down:
	podman-compose --env-file ./.env down

re: down up

ps:
	podman-compose --env-file ./.env ps

restart:
	podman-compose --env-file ./.env restart $(ARGS)

### Logs

logs:
	podman-compose --env-file ./.env logs
	
flogs:
	podman-compose --env-file ./.env logs -f

### Clean

vclean:
	if [ $(ENV) = dev ]; then \
		echo "Removing volumes"; \
		sleep 2; \
		podman-compose down -v; \
	else \
		echo "WARNING: do not run vclean in prod"; \
	fi

clean: down
	sudo rm -rf .go.cont server/tmp/main

fclean: clean
	podman system prune -af --volumes

################### Postgres #####################

psql:
	psql -h 127.0.0.1 -p $(DB_EXT) -U $(DB_USER) -d $(DB_NAME)

################### STARTUP #####################

air:
	go install github.com/air-verse/air@latest
	air

.PHONY: test up down re ps logs flogs vclean fclean psql air
