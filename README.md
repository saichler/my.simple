# Model Driven Programing Framework for Micro Services

* **"80%-90% of the code ever written is repetitive and could have been avoided."**
* **"Work Less & Do more... The impact of correctly designing & writing software results in better work/life balance"**
* **"I am so good, that if I get hit by a bus, the team can still go on."** 

## Overview
**Model Driven Development**, or **MDD** is a development methodology, purposed in **extremely expediting** software development **without** compromising on quality.
**However**, there isn't any **coherent, secure, simple, scalable, abstract, maintainable, high available, agnostic & fullstack** framework that delivers on the concept.

### About the Author...
**Decades** of trial & error in the repository [saichler](https://github.com/saichler), alongside of **decades** of software architecting, developing, testing & **maintaining** experience, alongside deep observation & insights, 
were consolidated & purified into this project.

## Help from a Giraffe
![alt text](https://github.com/saichler/my.simple/blob/main/giraffe.png)

## Can Software Development be that Simple?
# Yes! Yes! Yes!

**my.simple** is a full implementation, providing fullstack, coherent, components for **Model Driven Programing**, without compromising on **quality or security**.
Its purpose is to **extremely** reduce & simplify **Time to Market** when developing a MicroServices based applications/services, 
while being completely **agnostic** to the **underline infrastructure**. 

**Infrastructure agnostic** as to:
* **Machine/VM Phisical location** - The code is the same, regardless of process location. 
* **Running On Bare Metal** - Can be run on bare metal, no need for dockers or K8s.
* **Being Dockerized** - Can be dockerized, no need for K8s.
* **Deployed on Kubernetes** Can be deployed on K8s.
* **Database/Datastore** The code is the same, regardless of persistence layer you choose.

With:

* **No Install** - Just coherent & agnostic libraries to use.
* **Seamless Config** - Minimum to no config is needed.
* **Seamless Security** - Integrated Security to the Bone.
* **Built in health & monitoring** - Health made simple.
* **Single Point of Maintenance** - The model control it all.

Deep dive into my.simple [story](https://github.com/saichler/my.simple/blob/main/docs) or continue on to the components & building blocks of Model Driven Programing.

## Example Application
**my.simple** contains vast amount of components. 
To emphasize, how easy it is to build, test & deploy a microservice base application with **my.simple**, 
allow us to show off with an [example application](@TODO - example application).


# So let's Make It Simple... 

And deliver on the promise, as well as explaining the **"Magic"**.

## Basic Building Blocks
The basic building blocks are a must-have for any application. 
They implement the most basic design patterns, used inside any microservice development.
* [String conversions & concatenation](https://github.com/saichler/my.simple/tree/main/go/utils/strng)
* [Logs & Logging](https://github.com/saichler/my.simple/tree/main/go/utils/logs)
* [Blocking Sync Queues](https://github.com/saichler/my.simple/tree/main/go/utils/queues)
* [Sync Maps](https://github.com/saichler/my.simple/tree/main/go/utils/maps)

## [Model Driven Security](https://github.com/saichler/my.simple/tree/main/go/security)
**Security** & AAA is embedded from day 1 throughout the components, hence avoiding later challenges, 
trying to embed it when the project is well ongoing.

## [Struct Type Registry](https://github.com/saichler/my.simple/tree/main/go/registry)
**Unfortunetaly**, Golang does not have a registry of **Name2Type** fashion. 
This is an implementation of such a mechanism so instances can be instantiated by type name. 

## [Secure Networking](https://github.com/saichler/my.simple/tree/main/go/net)
One of the biggest challenges of MicroServices, one that usually pose the biggest pain point, "Troubles"/Challenges is the Process 2 Process communication.
In other words, "How will my MicroServices **Communicate, Interact & Share data** with each other?"
A single mistake in this area will crete a very big headache, pose challenges that should not have been there, worsen dramatically the engineers work/life balance and will cause the company Hundreds of Millions of dollars in maintenance.
**my.simple Secure Networking** is giving a seamless & simple, secure communication between your MicroServices.

## [Model Introspection & Deep Cloning](https://github.com/saichler/my.simple/tree/main/go/introspect)
The process of deep analysis & cloning a tree/graph models at runtime without having the model at compile time is called Introspection.
This is extremely usable in agnostic handling of **Delta Notifications**, **ORM (Object Relation Mapping)** &**Generic Distributed Cache**.

## [MDQL (Model Driven Query Language)](https://github.com/saichler/my.simple/tree/main/go/mdql)
Process 2 Process interaction is de-facto a language. 
To avoid reinventing the language, we converged to English, it is time to converge to the same query language in software.
**MDQL**, so the **"Wheel"** will not be invented everytime on every project.

## [Protocol Buffers Object](https://github.com/saichler/my.simple/tree/main/go/protobuf_object) 
**An Object** like implementation for **Protocol Buffers**, so delta updating model attributes will be generic and seamless.

## [Metadata driven getter/setter for nested tree/graph models](https://github.com/saichler/my.simple/tree/main/go/instance)
**Instance** is a metadata driven utility for updating **Delta** pieces of data in a nested tree/graph model instances. 
With a few lines of code, one can generically receive and update values inside an internal distribute cache copy.

## [Model Driven Updater](https://github.com/saichler/my.simple/tree/main/go/updater)
**Extremely** simplifying, **Updater** is a **nested model agnostic** utility to update model instances nested attributes, track them and react on explicit nested model attribute values changes. 

## [Seamless & Pluggable, Model Driven, ORM](https://github.com/saichler/my.simple/tree/main/go/orm)
**ORM (Object Relation Mapping)**, has been there for a few decade or two, **so what is new?!** 

How about **No Configuration, No Annotations & Pluggable Persistence Layer**. 

## API Gateways
Dynamic, Secured, [API Gateways](https://github.com/saichler/my.simple/tree/main/go/api_gateways) that avoids multiple implementations, bugs & maintenance.

## Model Driven Service Points

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