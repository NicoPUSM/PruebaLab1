hello:
	echo "hello world"

docker-latinoamerica:
	sudo docker build -f /LatinoAmerica/Dockerfile.LatinoAmerica . -t latinoamerica:latest
	sudo docker run --rm --name LatinoAmerica --network="host" latinoamerica:latest

docker-asia:
	sudo docker build -f /Asia/Dockerfile.Asia . -t asia:latest
	sudo docker run --rm --name Asia --network="host" asia:latest

docker-europa:
	sudo docker build -f /Europa/Dockerfile.Europa . -t europa:latest
	sudo docker run --rm --name Europa --network="host" europa:latest

docker-australia:
	sudo docker build -f /Australia/Dockerfile.Australia . -t australia:latest
	sudo docker run --rm --name Australia --network="host" australia:latest

docker-ONU:
	sudo docker build -f /onu/Dockerfile.Onu . -t onu:latest
	sudo docker run -i --rm --name Onu --network="host" onu:latest

docker-datanode1:
	sudo docker build -f /datanode-1/Dockerfile.Datanode-1 . -t datanode-1:latest
	sudo docker run --rm --name Datanode-1 -p 50052:50052 --network="host" datanode-1:latest

docker-datanode2:
	sudo docker build -f /datanode-2/Dockerfile.Datanode-2 . -t datanode2:latest
	sudo docker run --rm --name Datanode-2 -p 50053:50053 --network="host"datanode2:latest

docker-OMS:
	sudo docker build -f /namenode/Dockerfile.Oms . -t oms:latest
	sudo docker run --rm --name Oms -p 50051:50051 -p 50052:50052 -p 50053:50053 --network="host" oms:latest

clean:
	rm -f ./datanode-1/DATA.txt
	rm -f ./datanode-2/DATA.txt
	rm -f ./namenode/DATA.txt