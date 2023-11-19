# ORM

## Overview
ORM have been there for a decade or two, however it has never been an **easy task** to integrate and use one.

The **my.simple ORM** is built to completely be agnostic & abstract to:
* **Where the data is persisted?** 
* **How the data is persisted?**
* **How the data is being queries?**

So:

* **The persistence layer is pluggable**, hence **seamless switching** between sql databased, no-sql database, file based yaml/csv/json persistence layers.
* **The usage of the **MDQL** as an interface, hence **The quaries are agnostic** to the persistence layer.

## [Relational Data](https://github.com/saichler/my.simple/tree/main/go/orm/relational)
The challenge of translating **model instances** to **relational data** or 
back from **relational data** to **model instances**, 
is encapsulated and made extremely simple & seamless with this component.