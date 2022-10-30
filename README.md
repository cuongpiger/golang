![](./images/book_cover.png)

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
  * [https://refactoring.guru/design-patterns/builder/go/example](https://refactoring.guru/design-patterns/builder/go/example) 
* This pattern lets you **construct complex objects step by step**.
* **For example**: You have a `House` object, sometime you want this **house** has a `swimming_pool`, but later - you want this **house** has a `garden` and then you continuously want this **house** has a `garage`. You can use the `Builder` pattern to **create this house step by step**.

### 3.1.3. [Singleton](https://github.com/cuongpiger/golang/tree/hands-on-software-architecture-with-golang/chap03/singleton)
* *References*
  * [https://refactoring.guru/design-patterns/singleton/go/example](https://refactoring.guru/design-patterns/singleton/go/example)
* Is the design pattern that **restricts the creation of objects** to **one single instance**.
* **For example**: You want to create a `Mongo` object which represents a connection to the `MongoDB` database. You want to make sure that you only have **one single instance** of this `Mongo` object and you do not want to create **multiple instances** of this `Mongo` object. You easily you this `Singleton` pattern to **make sure that you only have one single instance** of this `Mongo` object.

### 3.1.4. [Abstract Factory](https://github.com/cuongpiger/golang/tree/hands-on-software-architecture-with-golang/chap03/abstract_factory)
* *References*
  * [https://refactoring.guru/design-patterns/abstract-factory/go/example](https://refactoring.guru/design-patterns/abstract-factory/go/example)
* This design pattern defines **an interface** for creating **all distinct products** but leaves the actual product creation to **concreye factory classes**.
* The client code calls the creation methods of a factory object instead of creating products directly with a constructor call.
* **For example**: Your clothes are selling **two different products** `Shoe` and `Shirt`. But in each product, you have **two different types** `Nike` and `Adidas`. You can use the `Abstract Factory` pattern to **create these products**.

### 3.1.5. [Prototype]()
* *References*
  * [https://refactoring.guru/design-patterns/prototype/go/example](https://refactoring.guru/design-patterns/prototype/go/example)
* This design pattern lets you **copy existing objects without making your code dependent on their classes**.
* **For example**: You have a `Car` object, and you want to create a `Car` object with the same properties as the `Car` object. You can use the `Prototype` pattern to **create a new `Car` object**.