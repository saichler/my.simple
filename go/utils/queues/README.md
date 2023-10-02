# Package queues
The package **queues** provides a set of blocking queues to use in a multiple go routing environment.

# Queue type
The Queue type is a base for all the queues, it is generic and accepts all type of data in its cell. When the Next method is called, and if the queue is empty, it is blocking until a new element is inserted into the queue. It is using the sync.Cond to broadcast a new element inserted to the queue. It is purposly not using channels to be able to be persisted in the future. 