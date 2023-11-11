# Logs & Logging

## Overview
Logging, although seems as a straight forward component, has its challenges:
* How to reduce/take away the impact of logging from your running code?
* How to log to multiple persistence layers with no impact to your running code?
* How to format the log, so it will be visualizing enough to distinguish between Debug & Error?
* Flexible log line format.

## Logger Queue & Format
The Logger queue is offloading the formatting and persisting of a log line to 
a different thread/go routine task so the impact of logging will now be observed by the running thread.

## Default Log Line Format

formatting is like so:
````
Tr > Trace message
 Db > Debug message
  In > Info message
   Wr > Warning message
    Er > Error message
````
So errors will be highlighted like so
````
 Db > Debug message
    Er > Error message
 Db > Debug message
````