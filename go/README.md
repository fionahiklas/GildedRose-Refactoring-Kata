# GO Starter

## Overview

The task is described [here](https://github.com/emilybache/GildedRose-Refactoring-Kata/blob/main/GildedRoseRequirements.txt)

Essentially the code is so awful that the only sensible solution is to re-write it from scratch 
using correct patterns/paradigms like TDD.  It would be a good lesson for the client in this 
case to ensure that code is of sufficient quality before accepting it.

Also being held hostage to bad data types because of an angry goblin isn't a good way to implement
software and someone needs to explain this politely to said mythical creature :D 

- Run :

```shell
make run DAYS=<number-of-days>  
```

If the `DAYS` variable is omitted it defaults to 2

- Run tests :

```shell
make test
```

- Run tests and show coverage report :

```shell
make test_report
```