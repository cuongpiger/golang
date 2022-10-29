* Packt source code: [https://github.com/PacktPublishing/Hands-On-Software-Architecture-with-Golang](https://github.com/PacktPublishing/Hands-On-Software-Architecture-with-Golang)

# [Chapter 3. Design patterns](https://github.com/cuongpiger/golang/tree/hands-on-software-architecture-with-golang/chap03)
## 3.1. Creational design pattern
* Is design patterns that deal with object creation mechanisms in a safe and efficient way.
* With these design patterns, the code using an object **need not** know details about **how the object is created**.

### 3.1.1. [Factory method](https://github.com/cuongpiger/golang/tree/hands-on-software-architecture-with-golang/chap03/factory)
* Is used to create **other objects**.
* **For example**, we have the `Vehicle` object, and we want to create `Car` and `Motorbike` objects. We can create a `VehicleFactory` object that through this object, it will create the `Car` and `Motorbike` objects for us.