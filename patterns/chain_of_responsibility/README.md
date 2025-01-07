# Chain Of Responsibility Pattern

The Chain of Responsibility design pattern is useful when multiple objects might handle a request, 
but the exact object that handles the request is not determined until runtime. 
It allows you to pass the request along a chain of potential handlers.

## Use cases

### Authentication and Authorization Middleware:

- A request is passed through a chain of middleware objects.
- Each middleware object checks if the request is authenticated and then checks if the request is authorized.
- Each handler has its own logic and can decide whether to pass the request to the next handler in the chain.

