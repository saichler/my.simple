# my.simple - Real **fullstack** software **wheels** to better your work/life & business balance  
## Engineers way of thinking is mostly concentrated on the **How?**
![alt text](https://github.com/saichler/my.simple/blob/main/girrafe.png)

This question is originally from a line of questioning dedicated to discover if one can simplify a challenge. 
Engineers will start planning, how big is the giraffe?, how big is the refrigerator? & etc...
## However, the most important, overlooked, question is the **What?**
What is the ROI in placing a giraffe into a refrigerator? 
What, in my end users business, will improve by placing a giraffe into a refrigerator?
Those line of questions are mostly overlooked by the software engineers/architects, 
which leads to ventures that cause companies, which write software, to waste **Billions of $** without any ROI.
## But life isn't perfect, to say the least, on the **How?** either
**Do not invent the wheel!** As engineers, we all know that... However, as engineers, we tent to, intentionally or subconsciously, identify the **wheel** incorrectly.
A simple example for wrongly identifying the **wheel** is **API**. 
Most, if not all, engineers will identify the **wheel** as the protocol... Restful, GRPC, KAFKA. However, with an analogy to **Language**, those are only the **alphabet** of the language.
The real **wheel** here is the way processes query and share their data/models with each other, AKA the **language**.

Engineers tent to reuse the same **alphabet**, while every time re-inventing a different **language**.
It is **unimaginable** how much, wasted, time & money was invested just on this common mistake.
Teams, having long meetings, discussing and implementing, API between their processes over & over again in an agonising, costly & repetitive pattern, re-inventing the **wheel** over & over again.

**Remember, this is just one example... They are many more, throughout the software stack**

## my.simple 

**Stateless/Stateful, Active/Active, Active/Passive, Security, High Availability, Horizontal Scaling, API, Kubernetes,
Microservices.** All those big words usually popup during planning of a distributed application...
**The problem starts during implementation!**

Over-engineering, Over-complexity & trying to re-use & push past, bloated, code of a single process application into a
container is an **epic scale pandemic**,
causing companies and the industry to spend trillions of $, re-inventing a **complex, money pit & unmaintainable "
wheels"** that should have been simple...

## What is my.simple?
**My simple is an abstracted, agnostic & coherent full stack framework with integrated Security.** 
In a nutshell, it means that the challenge was not just to write the implementation for each component in a simple way, it was also all about making them **agnostic**.
Turns out that **Agnostic as a guideline** outputs a simple & scalable solution to each challenge... And there are many!

## So let's do that properly & Make It Simple... 

### Basic Building Blocks
The basic building blocks are a must-have for any application. 
They implement the most basic design patterns used inside any microservice and expedite development.
* [String conversions & concatenation](https://github.com/saichler/my.simple/tree/main/go/utils/strng)
* [Logs & Logging](https://github.com/saichler/my.simple/tree/main/go/utils/logs)
* [Blocking Sync Queues](https://github.com/saichler/my.simple/tree/main/go/utils/queues)
* [Sync Maps](https://github.com/saichler/my.simple/tree/main/go/utils/maps)
* [Struct Type Registry](https://github.com/saichler/my.simple/tree/main/go/utils/registry)
* [Security](https://github.com/saichler/my.simple/tree/main/go/security)

### Process 2 Process Secure Networking
One of the biggest challenges of MicroServices, one that usually pose the biggest pain point, "Troubles"/Challenges is the Process 2 Process communication.
In other words, "How will my MicroServices **Communicate, Interact & Share data** with each other?"
A single mistake in this area will crete a very big headache, pose challenges that should not have been there, worsen dramatically the engineers work/life balance and will cause the company Hundreds of Millions of dollars in maintenance.
**[my.simple Secure Networking](https://github.com/saichler/my.simple/tree/main/go/net)** is giving a seamless & simple, secure communication between your MicroServices.

### Protocol buffers & 

* Protobuf Actions & Types
* Stateles Services

### Pluggable Data Stores & ORM

* Pluggable ORM & Data storing

### Stateful Services

* Leader/Follower elections
* Cache Sync
* Stateful distributed Transactions

### Maintainability

* Automation & Regression Testing

## Example Applications ##

There are two nerratives we need to cover, stateless application & statefull applications.
Will develop both, following the thumb rules for each one and how its challenges are solved within **my.simple**.

Join in this step-by-step journey of building an example application using **my.simple** infra. There is a lot... Will
cover the **"What?s"** and deep dive to **my.simple "How?s"**, that will save you a huge amount of effort, time & money
and allow you to concentrate on what really matters... **Your business.**