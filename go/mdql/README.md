# MDQL (Model Driven Query Language)

## Overview
Process 2 Process interaction is de-facto a language. 
**A language that is re-invented everytime, by every team on every project**.
* Huge amount of effort!
* Huge amount of time!
* Huge amount of interlock!
* Huge mount of maintenance!

**UNIMAGINABLE amount of money...**

Consuming the proprietary API query from the NBI, translating it to a set of queries to the persistence layer (like SQL DB), is also considered **creating a proprietary language**.

The world has consolidated into **English** when developing software, it is time to **consolidate into a single query language!**

## Model Driven Query Language
**MDQL** is a model driven, model agnostic, API query language, 
enabling the **Querying Entity** to express intent to fetch model instances, filter & scope model instances from the **Querying Provider**, while the **Querying Provider** has minimal to no effort exposing & securing the model instances.

In other words, unified language to query a model provider regardless if it is an SQL database, Microservice, No SQL database & etc. 
**Offloading the translation of the component/persistent layer from the consumer to the provider, giving the consumer a unified language to fetch model instances from any kind of provider.**

## Syntax

An example of fetching data:
````
fetch <model (nested) element> only <scoped attributes (nested)> (optional) criteria <expression referring attributes (nested)>"
````

## Usage