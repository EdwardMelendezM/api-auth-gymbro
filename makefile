generate_seed_mysql:
	docker-compose up -d

build-api:
	docker build -t api-auth .

run-api:
	docker run -d -p 9001:9001 \
      --name api-auth-container \
      -e SERVER_PORT=9001 \
      api-auth

dev-docker-tag:
	docker tag api-auth us-east4-docker.pkg.dev/pragmatic-zoo-425418-a3/infocode1901/api-auth:latest

dev-docker-push:
	docker push us-east4-docker.pkg.dev/pragmatic-zoo-425418-a3/infocode1901/api-auth:latest
