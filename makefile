export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin

infra:
	docker-compose -f docker/docker-compose.yml up -d

infra-off:
	docker-compose -f docker/docker-compose.yml down
