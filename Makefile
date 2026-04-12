######### STARTUP #############

desktop:
	cd desktop && npm start

server:
	cd server && go run main.go

########## DOCKER #############

.PHONY: desktop server
