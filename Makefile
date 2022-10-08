upload:
	docker build -t molel/cloud-drive:latest .
	docker push molel/cloud-drive:latest

up:
	docker-compose up

down:
	docker-compose down