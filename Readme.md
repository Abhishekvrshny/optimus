# OptimuS: Optimised Service Bus.
This was a submission for Redis `Beyond Cache` Hackathon 2020 at Bengaluru. **Not to be used in production**

## Inspiration
In the world of microservices, fast and reliable communication between them is a must for any event-driven or async architecture to be successful. Most of the service bus implementations in the market today rely on one or the other type of persistent queue. 
The challenges with persistent queues are:
1. They introduce unnecessary delay when most of the subscribers are fast and can immediately consume the messages.
2. They add to the cost of infra, in terms of storage, cpu etc.
3. The operability and maintainability costs with persistent queues like kafka are also high,

## What it does
OptimuS introduces the idea of `fast path` in the service bus implementation using Redis Pub/Sub i.e if the messages can be consumed by subscribers within the timeout defined by the producer, they will not be persisted. But, if due to some reason some consumer is slow, the messages are persisted to relevant kafka queue for the topic-subscriber which can be consumed by the slow consumer. The slow consumer can also add retryability, replayability etc.

This works because, in practice, most of the subscribers are fast most of the time and hence OptimuS helps to optimise delivery by being fast and saving costs!

## How I built it

This solution has the following components:
1. Redis Pub/Sub, for fast path.
2. Kafka, for slow path.
3. Zookeeper, to persist information of topic, subscribers and consumers listening on specific topics etc.
4. Golang, the language of choice for concurrency, for writing
- the gateway layer, and
- the consumer layer

## Challenges I ran into
No specific challenges as such, but time crunch on implementing the solution end to end.

## Accomplishments that I'm proud of
Able to boot up the MVP in such a short period.

## What I learned
The power of Redis :P

## Design
![High Levl Design](https://user-images.githubusercontent.com/12811812/72902489-8f0f2080-3d51-11ea-8e13-edbd727be5aa.jpg)

## Slide deck
https://docs.google.com/presentation/d/1TNeAIpfF48JhOWw29PNqjyEv9iLZg3YiIaVAoTWeYlA/edit?usp=sharing

## Demo video
https://youtu.be/cOOkdbCiK9k

## What's next for OptimuS: Optimized Service Bus
Make it production grade

