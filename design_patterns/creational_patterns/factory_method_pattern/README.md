# FACTORY METHOD

The Factory method pattern (or simply, Factory) is probably the second-best known and used design pattern
in the industry. Its purpose is to abstract the user from the knowledge of the struct it needs to acheive a specific purpose
By delegating this decision to a factory, this factory can provide the object that best fits the user needs or the most updated version.
It can also ease the process of downgrading or upgrading of the implementation of an object if needed

## Description

When using the Factory method design pattern, we gain an extra layer of encapsulation so that our program can grow in a controlled environment
with the factory method, we delegate the creation of families of objects to a different package or object to abstract us from the knowledge of the pool
of possible objects we could use. Imagine that you have two ways to access some specific resource: by HTTP or FTP. For us the specific implementation
of this access should be invisible. Maybe, we just know that the resource is in HTTP or in FTP, and we just want a connection that uses onee of these protocols.
Instead of implementing the connection by ourselves, we can use the Factory method to ask for the specific connection. With this approach we can grow easily in the future
if we need to add an HTTPS object

### Objective of the Factory method

The following objectives of the Factory Method design pattern must be clear to you

- Delegating the creation of new instances of structures to a different part of the program
- Working at the interface level instead of with concrete implementations
- Grouping families of objects to obtain a family object creator
