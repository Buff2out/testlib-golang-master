# Ints comparator checker
*CheckerVersion: 1.0*  
This checker sequentially compare ints from output stream and the reference answer stream.  
## Verdicts:
- **CF** - Throw when checker can't get one of the stream or can't write to one of the stream 
(see STDERR and answer pipe for more info)
- **WA** - Throw when integer value from output stream not equal with value from answer stream.
- **PE** - Throw if one of the symbols from output stream cannot be converted to int.  
Can be thrown when user answer contains space at the end of the output (//TODO: Need to be fixed in next version)
- **EF** - Throw if one of the symbols from answer stream cannot be converted to int
- **OK** - When output and answer streams are equal