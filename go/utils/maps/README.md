# Synchronized Maps

## Overview

**Synced Maps** are **Thread Safe** maps,
reducing the complexity of scoping locking/unlocking code sections and avoiding deadlocks in a multithreaded
environment.

## SyncedMap

Abstract implementation of a **synced map** to be wrapped with type2type maps.

## String2StringMap

Wrapper over **SyncedMap** for **string** to **string** Map.

## String2TypeMap

Wrapper over **SyncedMap** for **string** to **reflect.Type** Map.

## String2BoolMap

Wrapper over **SyncedMap** for **string** to **bool** Map.

## Usage
````
mymap:=maps.NewSyncMap()
mymap:=maps.NewString2StringMap()
mymap:=maps.NewString2TypeMap()
mymap:=maps.NewString2BoolMap()
````