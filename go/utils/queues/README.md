# Package queues

## Overview
**Blocking Queues** are similar to channels and are **Thread Safe**, 
however instead of using channels they are using the **sync.Cond** with a future implementation of persistence queue in mind. 

## Queue
**Abstract**, base, Queue implementation, wrapped by all **Type** implementations.

## ByteSliceQueue
A wrapper over queue to have a typed []byte elements.

## Usage
````
max_capacity:=50000
myqueue:=queues.NewQueue("my queue name",max_capacity)
````