# Package: strng
The package **strng** has a very simple purpose, handle in a simple & efficient way string concatenation & To/**From** string representation of different types.
## Why From?
Later in the journey, when we review the various plugins of the **my.simple.orm**, will see the value of having a string that is reconvertable back to its object form. 
## String type
The String type is a wrapper of bytes.buffer so make it seamless to concatenate string in an efficient way. 
When writing code, specially go code, we would like to avoid concatenating string inside code as this is a very slow operation.
The String type also holds the option to **include/exclude** the embedded string typing from the **ToString** via the **TypesPrefix** attribute.
## ToString
Convert the instance into string representation and adds info about the types/kinds in the prefix of the string so it can be converted back to an instance without any extra info. For example, an **int8** with value of **5** will be converted to '**{3}5**', **3** is the Kind of **int8** and **5** is the value. This is different than the standard String as it also contains the info to convert it back to an **int8** for the FromString.
## FromString
Convert a string containing the typing of the value from the ToString method, back to an **instance**.
