# Builder Design Pattern

Talking about creational design patterns, it looks pretty semantic to have a Builder design pattern.
The Builder patterns helps us construct complex objects without directly instantiating their struct,
or writing the login they require. Imagine an object that could have dozens of fields that are more complex
structs themselves. Now imagine that you have many objects with these characteristics, and you could have more. We
don't want to write the logic to create all these objects in the package that just needs to use the objects.

## Description

Instance creation can be as simple as providing the opening and closing braces {} and leaving the instance with zero values,
or as complex as an object tht need to make some API calls, check state, and create objects for its fields You could also have an object
that is composed of many objects, something that's really idiomatic in Go as it doesn't support inheritance.

At the same time, you could be using the same technique to create many type of objects. For example, you'll use almost the same technique.
to build a car to build a bus, except that they'll be different sizes and number of seats, so why don't we reuse the construction process?
This is where the builder patterns comes to rescue

### Objective of the Builder pattern...

A builder design pattern tries to:

- Abstract complex createions so that object createion is seperated from object user
- Create an object step by step by filling its fields and creating embedded objects
- Reuse the object creation algorithm between many objects...
