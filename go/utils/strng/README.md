# strng
## Overview
Concatenating strings is a bad practice by junior, and not so junior, programmers... When practiced, It creates huge logging delays throughout the code, one that is being noticable only when the project is getting to big and the time & effort of mitigating it becomes not visible. 

What about **From String**? There were points while developing **my.simple** where From String would have make my life much easier, hence embedded **FromString** functionality inside this package. For example, translating a slice of strings into a single string for **ORM** purposes.

## Package Content  
### String type
The String type is a wrapper of bytes.buffer so make it seamless to concatenate string in an efficient way. 
When writing code, specially go code, we would like to avoid concatenating string inside code as this is a very slow operation.
The String type also holds the option to **include/exclude** the embedded string typing from the **ToString** via the **TypesPrefix** attribute.
### ToString
Convert the instance into string representation and adds info about the types/kinds in the prefix of the string so it can be converted back to an instance without any extra info. For example, an **int8** with value of **5** will be converted to '**{3}5**', **3** is the Kind of **int8** and **5** is the value. This is different than the standard String as it also contains the info to convert it back to an **int8** for the FromString.
### FromString
Convert a string containing the typing of the value from the ToString method, back to an **instance**.
