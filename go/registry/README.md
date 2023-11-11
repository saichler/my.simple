# Struct name to Type registry

## Overview
Unfortunetaly, Golang does not have a name -> type registry built in, 
e.g. you cannot instantiate a struct based on its name only. 
This limits a component performing a generic task, while being **agnostic** to the data type it handles.

## The Resitry
The registry is a simple name->type repository that is able to instantiate a struct based on its type name via the reflect package.

## Usage
````
//Register a struct type via an instance of this struct
common.Registry.RegisterStruct(<a struct instance>)
````

````
//Instantiate a new struct instance
//v := common.Registry.NewInstance("my stuct name").(<cast to the type)
````