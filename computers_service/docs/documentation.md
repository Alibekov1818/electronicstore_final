Computers Service

This service is used to create, delete, get computers for electronicstore.
Computers Service uses gRPC to handle requests. Also handler for getting 
single computer can be used with rabbitMQ.

gRPC server runs on port 8003 and Rabbit on port 5672

Methods:

1. `GetComputer` needs ComputerId data type as a request. Returns Computer proto type
2. `GetComputers` needs empty GetComputersRequest data type as a request. Retuns ComputersList protoType
3. `CreateComputer` needs Computer data type as a request. Returns Computer proto type
4. `DeleteComputer` needs ComputerId data type as a request. Returns Computer proto type

RabbitMQ:

You can also use RabbitMQ to get single Computer. Service's queue name is `computers`. In order
to get single computer you have to publish to the queue byte array of ComputerId proto type. The
service will publish result to `computers_client` queue. Don't forget to unmarshall the response.