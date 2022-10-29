* Back to **[main branch](https://github.com/cuongpiger/golang/tree/main)**.
* Packt source code: [https://github.com/PacktPublishing/Hands-On-Software-Architecture-with-Golang](https://github.com/PacktPublishing/Hands-On-Software-Architecture-with-Golang)

# [Chapter 3. Design patterns](https://github.com/cuongpiger/golang/tree/hands-on-software-architecture-with-golang/chap03)
## 3.1. Creational design pattern
* Is design patterns that deal with object creation mechanisms in a safe and efficient way.
* With these design patterns, the code using an object **need not** know details about **how the object is created**.

### 3.1.1. [Factory method](https://github.com/cuongpiger/golang/tree/hands-on-software-architecture-with-golang/chap03/factory)
* *References*
  * [https://refactoring.guru/design-patterns/factory-method/go/example](https://refactoring.guru/design-patterns/factory-method/go/example)
* Is used to create **other objects**.
* **For example**, we have the `Vehicle` object, and we want to create `Car` and `Motorbike` objects. We can create a `VehicleFactory` object that through this object, it will create the `Car` and `Motorbike` objects for us.

### 3.1.2. [Builder](https://github.com/cuongpiger/golang/tree/hands-on-software-architecture-with-golang/chap03/builder)
* *References*
  * [https://refactoring.guru/design-patterns/builder](https://refactoring.guru/design-patterns/builder) 
* This pattern lets you **construct complex objects step by step**.
* **For example**: You have a `House` object, sometime you want this **house** has a `swimming_pool`, but later - you want this **house** has a `garden` and then you continuously want this **house** has a `garage`. You can use the `Builder` pattern to **create this house step by step**.