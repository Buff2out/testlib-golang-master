# Testlib for testing module for Code.Hedgehog platform

## Description
This library contains some methods to check user solutions on Code.Hedgehog platform
## [Checker example](example/checker/README.md)
## How to run
All checkers write testing result to `3 fd`.
Example of running command for local debug:  
`./main 3>&1`  
Require `checker.conf` file in root directory with data. [Example](example/checker/checker.conf)  
**Or** if you want, you can pass configuration in JSON to `4 fd`.
Example of running command:  
`./main 3>&1 4<conf.json`