# hyperloglog-go-kafka
Estimate cardinality for a data stream using kafka and Go.
This solution is a solution to the data-engineer-challenge.
It sends statistics to the output topic and the counting error margin is less than 1%.
It doesn't deal with latency, but it does work with historically data (you simply re-run it on a Kafka topic, since the entire point is for it to be stateless).

## How to run this project:

You can simply use `go run`, or if you want you can use Docker. If you are using Docker, just make sure the
container has access to the host network, so it can use the local Kafka instance. There are a few environment variables
that this project uses, but they have sensible defaults, so you don't have to bother with them. Here's the list of the
used environment variables and their default values:

```
KAFKA_BROKER=localhost:9092
USERS_TOPIC=users
STATS_TOPIC=stats
```

To send messages to a Kafka topic you can use this command:

```
gzcat stream.jsonl.gz | kafka-console-producer --bootstrap-server localhost:9092 --topic users
```

To receive messages from Kafka topic, you can use this command:

```
kafka-console-consumer --bootstrap-server localhost:9092 --topic stats
```

You can find sample data (`stream.jsonl.gz`) [here](https://tda-public.s3.eu-central-1.amazonaws.com/hire-challenge/stream.jsonl.gz).
