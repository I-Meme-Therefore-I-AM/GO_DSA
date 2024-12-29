# SINGLE DESIGN PATTERN

### Description

Singleton pattern is easy to remember as the name implies, it will provide you a single instance of an
object, and guarantee that there are no duplication. At the first call to use the instance, it is created
and then reused between all the parts in the application that need to use that particular behavior

## Objective of the Singleton Pattern

You'll use singleton pattern in many different situations. For example:

- When You want to use the same connection to a database to make every query
- when you opena secure shell (SSH) connection to a server to do a few tasks, and don't want to reopen the connection for each task
- If you need to limit the access to some variable or space, you use a singleton as the door to this variable
- If you need to limit the number of calls to some places, you create a Singleton instance to make the calls in the accepted window.
