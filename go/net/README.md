# Process 2 Process Secure Networking
Architecting & Implementing the way your MicroServices interact with each other, is a critical stage. 
Engineers tent to concentrate on the **How?**, e.g. Kafka, NATS, Rabbit MQ & etc., overlooking the real challenge of the **What?**, the API & Models.
Imagine the **How?** as the alphabet, and the **What?** as the language. 
Putting it like this, one will agree that having the same **alphabet** is a minor challenge vs. having the same **language**. 

The **my.simple**, although extremely simplifying secure networking between processes, is also concentrated around building a unified language between the processes so the **Language** will not be re-invented over & over again.
The networking module of **my.simple** is laying the foundation of unified API & Models sharing, minimizing extremely the current & future investment in infra & infra maintenance.

To provide seamless secure networking, the networking module was designed to be **Agnostic** to Kubernetes & Containers. You can **With** and you can **Without**...
Here is a slide showing how this was accomplished:

![alt text](https://github.com/saichler/my.simple/blob/main/go/net/SecureNetworking.png)