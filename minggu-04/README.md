# Komunikasi Antar Container pada Docker


## Komunikasi menggunakan Link

1. Menjalankan Redis

docker run -d --name redis-server redis

2. Membuat link

	docker run --link redis-server:redis alpine env

	docker run --link redis-server:redis alpine cat /etc/hosts

	docker run --link redis-server:redis alpine ping -c 1 redis

3. Mengkoneksikan ke App

	docker run -d -p 3000:3000 --link redis-server:redis katacoda/redis-node-docker-example

	curl docker:3000

4. Mengkoneksikan ke Redis CLI

docker run -it --link redis-server:redis redis redis-cli -h redis


## KOmunikasi menggunakan Network

1. Membuat network

docker run -it --link redis-server:redis redis redis-cli -h redis

docker run -it --link redis-server:redis redis redis-cli -h redis


2. Komunikasi Network

docker run --net=backend-network alpine env

docker run --net=backend-network alpine cat /etc/hosts

docker run --net=backend-network alpine cat /etc/resolv.conf

docker run --net=backend-network alpine ping -c1 redis


3. Mengkoneksikan 2 container

docker network create frontend-network

docker network connect frontend-network redis

docker run -d -p 3000:3000 --net=frontend-network katacoda/redis-node-docker-example

curl docker:3000


