# API Gateways

## Overview
Wiring the API, either RestFul or GRPC, has always been a challenge in a MicroServices environment.
Usually, those **End Points** are exposed, statically, via some configuration file and usually the payload is just forward to the MicroService API handler. E.g. you have **double implementation** of web services, bugs & maintenance.

## Gateway & Security that learns
In **my.simple**, the gateway implementation dynamically learns the End Points from the **Service Point Network**. On invocation, validate the **End Point** (and action) with the **SecurityProvider**. Deserialize, if its the RestFul gateway, the payload to Protocol Buffers. And use a Request/Reply method to send over the wire the request to the **Service Point Topic**.
Needless to say, the **API Gateway** is **Active/Active** out-of-the box by implementing **Session Affinity** via the **Distributed Cache** mechanism.

