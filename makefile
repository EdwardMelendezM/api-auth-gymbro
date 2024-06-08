generate_seed_mysql:
	docker-compose up -d

build-docker:
	docker build -t api-auth .

docker-tag:
	docker tag api-auth us-east4-docker.pkg.dev/pragmatic-zoo-425418-a3/infocode1901/api-auth:latest

docker-push:
	docker push us-east4-docker.pkg.dev/pragmatic-zoo-425418-a3/infocode1901/api-auth:latest
