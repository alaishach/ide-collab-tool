######### STARTUP #############

web:
	cd web && npm run dev

server:
	$(MAKE) -C ./server/ up

.PHONY: web server
