* Source code of Packt: [https://github.com/PacktPublishing/Go-Design-Patterns](https://github.com/PacktPublishing/Go-Design-Patterns)
* All design pattern in Golang: [https://golangbyexample.com/all-design-patterns-golang/](https://golangbyexample.com/all-design-patterns-golang/)

# 2. Structural Design Patterns
## 2.1. [Proxy](https://golangbyexample.com/proxy-design-pattern-in-golang/)
* Personal source code: [https://github.com/cuongpiger/golang/tree/go-design-patterns/chap04/proxy](https://github.com/cuongpiger/golang/tree/go-design-patterns/chap04/proxy)
* Use to wrap an object to hide some of its characteristics. These characteristics could be a heavy object, cache,...
## 2.2. [Decorator](https://golangbyexample.com/decorator-pattern-golang/)
* Personal source code: [https://github.com/cuongpiger/golang/tree/go-design-patterns/chap04/decorator](https://github.com/cuongpiger/golang/tree/go-design-patterns/chap04/decorator)
* Imagine you have a cup of bubble tea, you can add so many toppings to it. This is the decorator pattern.

# 3. Behavioural Design Patterns
## 3.1. [Interpreter design patter](https://viblo.asia/p/interpreter-design-pattern-tro-thu-dac-luc-cua-developers-djeZ1d43KWz)
* Personal design pattern: [https://github.com/cuongpiger/golang/tree/go-design-patterns/chap06/interpreter](https://github.com/cuongpiger/golang/tree/go-design-patterns/chap06/interpreter)
* This design pattern is used to define function which reading a sentence from command line, text,...
## 3.2. [Visitor](https://viblo.asia/p/visitor-design-pattern-tro-thu-dac-luc-cua-developers-gDVK2oGeZLj)
* Personal source code: [https://github.com/cuongpiger/golang/tree/go-design-patterns/chap07/visitor](https://github.com/cuongpiger/golang/tree/go-design-patterns/chap07/visitor)
* Imagine that you are developing the function **payment orders** of an e-commerce company. Users can pay by either credit cards or cash. So you need to go through entire the products inside your order and implement identified payment method for each one and then add all at the end. You can create 2 object `cash` and `credit card` and just implement them as params of function when looping via the orders.

## 3.3. [State](https://viblo.asia/p/state-design-pattern-07LKXjPDlV4)
* Personal source code: [https://github.com/cuongpiger/golang/tree/go-design-patterns/chap07/state](https://github.com/cuongpiger/golang/tree/go-design-patterns/chap07/state)
* This pattern is used when your program have a **finite state** and need to **switch** between them. For example, you are developing a game, your game will transfer these step: "Start", "Play", "Pause", "End". You can use this pattern to switch between these state.
## 3.4. [Mediator](https://golangbyexample.com/mediator-design-pattern-golang/)
* Personal source code: [https://github.com/cuongpiger/golang/tree/go-design-patterns/chap07/mediator](https://github.com/cuongpiger/golang/tree/go-design-patterns/chap07/mediator)
* This pattern suggests creating a **mediator object** to **prevent direct communication** among **objects** so that **direct dependencies between them is avoided**.

## 3.5. [Observer](https://golangbyexample.com/observer-design-pattern-golang/)
* Personal source code: [https://github.com/cuongpiger/golang/tree/go-design-patterns/chap07/observer](https://github.com/cuongpiger/golang/tree/go-design-patterns/chap07/observer)
* This pattern allows **an instance** *(called **SUBJECT**)* to **publish events** to **other multiple instances** *(called **OBSERVERS**)*. These **observers** subcribe to the **subject** and hence get notified by events in case of **any change happening in the SUBJECT**.

## 3.6. [Memento](https://golangbyexample.com/memento-design-pattern-go/)
* Personal source code: [https://github.com/cuongpiger/golang/tree/go-design-patterns/chap06/memento](https://github.com/cuongpiger/golang/tree/go-design-patterns/chap06/memento)
* Imagine that you are developing a program that you can **UNDO** and **REDO** actions. This design pattern help you to do that.

## 3.7. [Template](https://golangbyexample.com/template-method-design-pattern-golang/)
* Personal source code: [https://github.com/cuongpiger/golang/tree/go-design-patterns/chap06/template](https://github.com/cuongpiger/golang/tree/go-design-patterns/chap06/template)
* If your app have notification services, including of SMS message and Email, two methods are different but if they have the same process, preparing data, build message, and send to customer, you can think about this pattern.

## 3.8. [Chain of Responsibility](https://golangbyexample.com/all-design-patterns-golang/)
* Personal source code: [https://github.com/cuongpiger/golang/tree/go-design-patterns/chap05/chain](https://github.com/cuongpiger/golang/tree/go-design-patterns/chap05/chain)
* If you have a long process and they are related to each other, the previous step passes the result to the next step, you can use this pattern.

## 3.9. [Command](https://golangbyexample.com/command-design-pattern-in-golang/)
* Personal source code: [https://github.com/cuongpiger/golang/tree/go-design-patterns/chap05/command](https://github.com/cuongpiger/golang/tree/go-design-patterns/chap05/command)
* Example that you create help function to show the description about a command in your console app. You can use this pattern to do that, it will create a `help` object to reference to the all function which implemented the `Command` interface.

## 3.10. [Strategy](https://golangbyexample.com/strategy-design-pattern-golang/)
* Personal source code: [https://github.com/cuongpiger/golang/tree/go-design-patterns/chap05/strategy](https://github.com/cuongpiger/golang/tree/go-design-patterns/chap05/strategy)
* Personal source code: [https://github.com/cuongpiger/golang/tree/go-design-patterns/chap05/strategy](https://github.com/cuongpiger/golang/tree/go-design-patterns/chap05/strategy)
* Example you are coding the hashing password function with algorithm A and B, you can think about this patterm to implement it, create a hasher object and then put the hashed algorithm into it.

# [4. Go concurrency](https://github.com/cuongpiger/golang/tree/go-design-patterns/chap08)
* Anonymous function
* Callback function
* Channel
* Buffered channel
* Concurrent Singleton
* Mutex
* Ranging
* Select statement
* WaitGroup
