# my.simple - Simplifying, Expediting Microservices & Kubernetes Development

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