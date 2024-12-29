# Go | Functional Builder Pattern

This project demonstrates a functional builder pattern in Go by implementing a simple server system. 
The example showcases how to use the functional builder pattern to configure a server with options like protocol, host, port, and origins.

## Overview
The server:

- Listens on a specified address (protocol, host, and port).
- Accepts connections from clients.
- Logs incoming client connections.

While the `origins` feature is included in the configuration, it is not used in this demo. 
It just showed as an example.

## Functional Builder Pattern
The functional builder pattern is a design approach that uses functions to construct and configure objects in a modular, extensible, and readable way.

### Key Benefits of the Pattern:

1. **Flexibility**: Allows you to add or modify configuration options without altering existing code or constructors.
2. **Readability**: Code using the builder pattern is often easier to read, especially when configuring objects with many optional parameters.
3. **Scalability**: Adding new options or features is straightforward, as each new option is encapsulated in its own function.
4. **Default Values**: The pattern provides a clean way to handle default values while allowing them to be overwritten.

### How It Works
1. Start with a base configuration containing default values (e.g., `defaultConfig` in this demo).
2. Use functional options (functions that accept a pointer to the configuration and modify it) to build the final configuration.
3. Pass these functional options to a constructor, which applies them in sequence to generate the desired configuration.

## Practical Use Cases of the Builder Pattern
The functional builder pattern can be applied in many scenarios:

1. **Database Connections**: Configuring clients with options like timeouts, max connections, and credentials.
2. **HTTP Clients**: Setting headers, retries, and connection pooling.
3. **Microservices**: Building complex initialization pipelines for services with customizable settings.
