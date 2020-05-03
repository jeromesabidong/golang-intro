all: build-env run clean

build-env:
	@docker build -t my-golang-app .
run:
	@docker run -it --rm --name my-running-app my-golang-app
clean: 
	docker container prune -f