######### STARTUP #############

desktop:
	cd desktop && npm start

server:
	$(MAKE) -C ./server/ up

.PHONY: desktop server
