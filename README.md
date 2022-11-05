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

### 3.1.5. [Prototype](https://github.com/cuongpiger/golang/tree/hands-on-software-architecture-with-golang/chap03/prototype)
* *References*
  * [https://refactoring.guru/design-patterns/prototype/go/example](https://refactoring.guru/design-patterns/prototype/go/example)
* This design pattern lets you **copy existing objects without making your code dependent on their classes**.
* **For example**: You have a `Car` object, and you want to create a `Car` object with the same properties as the `Car` object. You can use the `Prototype` pattern to **create a new `Car` object**.

## 3.2. Structural Design Patterns
* Help delineate **clean relationships** between objects.
* Explain how to **assemble** objects and classes **into larger structures**, while kleeping these structures flexible and efficient.


### 3.2.1. [Adapter](https://github.com/cuongpiger/golang/tree/hands-on-software-architecture-with-golang/chap03/adapter)
* This design pattern help to transform the **old objects** into **new objects** which satisfy the current requirements.
* **For example**: Supposed that you get **data in XML format** from **your legacy company server**, but now you want to **combine this old data with your generated data from your team** and send this combined data to clients **in JSON format**. In these situations like that, you can think about this design pattern.

### 3.2.2. [Bridge](https://github.com/cuongpiger/golang/tree/hands-on-software-architecture-with-golang/chap03/bridge)
* *References*:
  * [https://refactoring.guru/design-patterns/bridge/go/example](https://refactoring.guru/design-patterns/bridge/go/example)
* This design pattern will **split** the **original business logic** or **original huge class** into **seperate class hierarchies** that can be developed **independently**.
* **For example**: You have **two types of `computer`s**, `Mac` and `Windows`. And you have **two types of printers**, `HP` and `Epson`. You can use this design pattern to **create these computers and printers** and make them able to communicate with each other.

### 3.2.3. [Composite](https://github.com/cuongpiger/golang/tree/hands-on-software-architecture-with-golang/chap03/composite)
* *References*
  * [https://refactoring.guru/design-patterns/composite/go/example](https://refactoring.guru/design-patterns/composite/go/example)
* Composite allows composing objects into a tree like structure.
* **For example**: You are developing a **file system**. In the file system, there are two types of objects, `File` and `Folder`. You can use this design pattern to **create these files and folders** because these objects are organized in **tree structure**. Now imagine you need to run **search** for a **particular keyword** in the file system. You can use this design pattern to **search for this keyword**.

### 3.2.4. [Decorator](https://github.com/cuongpiger/golang/tree/hands-on-software-architecture-with-golang/chap03/decorator)
* *References*
  * [https://refactoring.guru/design-patterns/decorator/go/example](https://refactoring.guru/design-patterns/decorator/go/example)
* It allows adding new behaviors to existing objects without modifying their code.
* We can dynamically place them inside special wrapper objects, called **decorator**.
* **For example**: You have a `Car` object, and you want to add a `GPS` feature to this `Car` object. You can use this design pattern to **add this GPS feature** to this `Car` object.

### 3.2.5. [Facade](https://github.com/cuongpiger/golang/tree/hands-on-software-architecture-with-golang/chap03/facade)
* *References*
  * [https://refactoring.guru/design-patterns/facade/go/example](https://refactoring.guru/design-patterns/facade/go/example)
* This design pattern provides a **simplified interface** to a **complex system**.
* **For example**: The methods of specific object are using so many lib, package,... and you want to **simplify the interface** of this object. You can use this design pattern to **simplify the interface** of this object.

### 3.2.6. [Flyweight](https://github.com/cuongpiger/golang/tree/hands-on-software-architecture-with-golang/chap03/flyweight)
* *References*
  * [https://refactoring.guru/design-patterns/flyweight/go/example](https://refactoring.guru/design-patterns/flyweight/go/example)
* Flyweight is a structural design pattern that lets you **fit more objects into the available amount of RAM** by **sharing common parts** of state **between multiple objects** instead of keeping all of the data in each object.
* **For example**: In a game of Counter-Strike, Terrorist and Counter-Terrorist have a different type of dress. For simplicity, let’s assume that both Terrorist and Counter-Terrorists have one dress type each. The dress object is embedded in the player object as below.

### 3.2.7. [Proxy](https://github.com/cuongpiger/golang/tree/hands-on-software-architecture-with-golang/chap03/proxy)
* *References*
  * [https://refactoring.guru/design-patterns/proxy/go/example](https://refactoring.guru/design-patterns/proxy/go/example)
* Proxy is a structural design pattern that provides an object that acts as a substitute for a real service object used by a client. A proxy receives client requests, does some work (access control, caching, etc.) and then passes the request to a service object.
* **For example**: A web server such as Nginx can act as a proxy for your application server:
  * It provides controlled access to your application server.
  * It can do rate limiting.
  * It can do request caching.

## 3.3. Behavioral Design Patterns
* Behavioral design patterns are design patterns that identify communication patterns among objects and provide solution templates for specific situations.

### 3.3.1. [Command](https://github.com/cuongpiger/golang/tree/hands-on-software-architecture-with-golang/chap03/command)
* *References*
  * [https://refactoring.guru/design-patterns/command/go/example](https://refactoring.guru/design-patterns/command/go/example)
* when you need to work with **multiple objects** but they are **not related to each other** but have **similar behavior**, such as tv, microwave, etc. You can use this design pattern to create an **command interface** containing the **similar behaviors** of these objects and **implement this interface** for each object.
* **For example**: You have `Tv` object and this object has **on/off** behaviors, you can use this design pattern to **create an interface** containing **on/off** behaviors and **implement this interface** for `Tv` object. Furthermore, you can **create a `Microwave` object** and **implement this interface** for `Microwave` object. And you also create a `Button` object that reponsible for **calling the behaviors** of `Tv` and `Microwave` objects.