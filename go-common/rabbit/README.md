# rabbit

Package that has a convenient management for RabbitMQ stuff.

## Exchanges

Further reading [here](https://www.rabbitmq.com/tutorials/amqp-concepts.html#exchange-default).

We don't use exchanges at all. All queues run in the `direct` RabbitMQ
exchange.

So this parameter is always passed as empty.

| Exchange type    | Default pre-declared names              |
|------------------|-----------------------------------------|
| Direct exchange  | (Empty string) and amq.direct           |
| Fanout exchange  | amq.fanout                              |
| Topic exchange   | amq.topic                               |
| Headers exchange | amq.match (and amq.headers in RabbitMQ) |

### Default Exchange

The default exchange is a direct exchange with no name (empty string) pre-declared by the broker. It has one special property that makes it very useful for simple applications: every queue that is created is automatically bound to it with a routing key which is the same as the queue name.

For example, when you declare a queue with the name of "search-indexing-online", the AMQP 0-9-1 broker will bind it to the default exchange using "search-indexing-online" as the routing key (in this context sometimes referred to as the binding key). Therefore, a message published to the default exchange with the routing key "search-indexing-online" will be routed to the queue "search-indexing-online". In other words, the default exchange makes it seem like it is possible to deliver messages directly to queues, even though that is not technically what is happening.

The default exchange, in RabbitMQ, does not allow bind/unbind operations. Binding operations to the default exchange will result in an error.

### Direct Exchange

A direct exchange delivers messages to queues based on the message routing key. A direct exchange is ideal for the unicast routing of messages. They can be used for multicast routing as well.

Here is how it works:

* A queue binds to the exchange with a routing key K
* When a new message with routing key R arrives at the direct exchange, the exchange routes it to the queue if K = R
* If multiple queues are bound to a direct exchange with the same routing key K, the exchange will route the message to all queues for which K = R

