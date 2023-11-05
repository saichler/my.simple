# my.simple
# Fullstack software building blocks for better work/life & business balance
## Overview
In a nutshell, **Software Development** is being treated & manage as any other industry out there.
However, some of the very basic management & work/life practices, alongside though to be true paradoxes, are simply not true for **Software Development**.
Via **Coherent Analysis**, **Design Decomposition** & **Out-of-the-Box** thinking, Software Development, done right, can solve paradoxes and extremely better work/life balance.

### The Employee Paradox
* The incentive of the Company: The **Employee** will **work more** for the wage.
* The incentive of the Employee: **Work less** for the wage.

### Not True for Software Development! 
### The company incentive is for the engineer to deliver Quality Code
Because quality code is good for the business:
* Lower MTTR on bugs
* Lower maintenance cost
* More deliveries in a shorter time
* Happy Customers

### The employee incentive is to deliver Quality Code
Because quality code is good for work/live balance:
* Less time spent on bugs
* Less time spent on maintenance
* More deliveries in a shorter time
* Happy Company

**E.g. Writing quality code creates a pattern where over time you Work Less & Do More**

### So how to "Work Less & Do More"?
If you wait a moment and think about it... If you follow the following thumb rules:
* Make it Simple
* Make it Encapsulated
* Make it Coherent
* Make it Agnostic
* Make it Optimized

## Engineers way of thinking is mostly concentrated on the **How?**
![alt text](https://github.com/saichler/my.simple/blob/main/girrafe.png)

This question is originally from a line of questioning dedicated to discover if one can simplify a challenge. 
**Immediately**,Engineers will start planning... "How big is the giraffe?" "How big is the refrigerator?" "How...? ..."
## However, the most important, overlooked, question is the **What?**
"What is the ROI in placing a giraffe into a refrigerator?" 
"What business will improve by placing a giraffe into a refrigerator?"

Those line of questions are mostly overlooked by the software engineers/architects, 
which leads to "ventures" that cause companies, writing software, to waste **Billions of $** without any true ROI.

## But life isn't perfect, to say the least, on the **How?** either
**Do not invent the wheel!** As engineers, we all know that... However, as engineers, we tent to, intentionally or subconsciously, identify the **wheel** incorrectly.
### The API as an example
**API** has the most common mistake in identifying the **wheel**... 
Most, if not all, engineers will identify the **wheel** as the **protocol**. **Restful**, **GRPC**, **KAFKA**, **NATS**,... 
However, with an analogy to **Language**, **protocols** are only the **alphabet** of the language.
The **wheel/language** in this analogy is the way **processes/microservices** concurrently query, share & update each other with data, models & updates.

**API definition of a process/microservice is like re-inventing a language over & over again, every time...**
The process of "You send me that, I will reply with this, I will update you with that" is a **huge time & money pit** when developing a microservice base application.
Just imagine how much effort, time & money is spent in that area, and that is without maintenance, versioning & backward compatability...

**Remember, this is just one example... They are many more, throughout the software stack**.

# my.simple 
Throughout the software stack of building a microservice based application, 
there are some **building blocks & challenges** that can be encapsulated into
a **single, agnostic & simple components** that can be used to remove, the money pits, infra challenges
and allow the team to concentrate on the business logic.

##So what is **my.simple**?
A collection of coherent, while agnostic, components that can extremely expedite the building of a microservice based application.
Years of fullstack experience, experimenting & coherence analysis were materialize in this repository.


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
* [Protobuf Object](https://github.com/saichler/my.simple/tree/main/go/utils/protobuf_object)
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