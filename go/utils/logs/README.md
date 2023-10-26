# Logs & Logging

Yet another logging interface, trying to encapsulate visual attributes to logs so mixed debug,info,warning,error messages will be indented differently so the viewer can spot easily the errors over the debugs, info & etc.

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