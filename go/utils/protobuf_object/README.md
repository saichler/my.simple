# ProtoBuf Object
## Overview
Protobuf is a very nice modeling & optimized serialization method, supported by multiple languages.
However, it is missing a key feature, the **Protobuf Object**. You cannot define an attribute of a protobuf message as **generic/object**, e.g. **can be anything**.
### Why is Protobuf Object important?
Simple, without a Protobuf Object a MicroService cannot update its peers on **Delta Changes** within its model and must share the entire model instance. This is a **major drawback** when the microservice model is a complex tree structure, with millions of instances.

### Generic/Object implementation
In preparation to support **Delta Notifications** between MicroServices, **ProtobufObject** was implemented as a vessel/infra to allow the capability of one MicroService to **Delta Notify** another MicroService with model changes in an optimal and optimized way & extremely reduce, over the wire, network utilization.