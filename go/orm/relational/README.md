# Relation Data

## Overview
To abstract and make it easy for the different relational stores plugins to interact & persist model instances, the **relational** component was created.

## Usages
````
// Create a new relational data to contains the translated data
relationalData := relational.NewRelationalData("<transaction ref>")

// Translate the instance/s to relational data. 
// data, is an instance or a list of instances.
// inspect, is an instance of introspect, can be common.Inspect of a separate instance.
err := relationalData.AddInstances(data, inspect)
``