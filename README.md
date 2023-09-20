Books Inventory
================

This is a simple books inventory application. It is a web backed to order books from an inventory

## Notes
* This application has two modes: GraphQL (preferred) and REST API
  * To run this application:
    ```
    $ docker compose -f docker-compose.graphql.yml up -d
    $ docker compose -f docker-compose.rest.yml up -d
    ``` 
    For graphql you can go to http://localhost:4000 and use the graphql playground
    For Rest use http://localhost:4000 and insomina, postman or any other rest client
* There are three microservices. In the Graphql version, there is true separation as Order service does not need to
  know anything about Inventory service to provide `books` on the same query. In the REST API version, the Order service
  must know about the Inventory service to provide the `books` on the order response as it needs to call the inventory 
  service internally.
* All microservices have their own databases. This is to avoid coupling between services. Services should be able to
  manage themeselves and evolve independently.
* You will noticed I pass context through every function, this is to enable usage for other libraries such as for 
  metrics and tracing which gets passed on usually through context. Additionally some libraries my require this to
  pass context so it can pass it back to the parent service. ex: [mfa-lib](https://github.com/honestbank/mfa-lib)


## GraphQL

You can see the schema by either running the application and going to http://localhost:4000 or by looking at the
`sdl.graphql` file. (gateway/sdl.graphql). gqlgen is used to generate boilerplate code for the graphql server.

## REST API

In each microservice there is a folder `/http/htt_routes` and a function `RegisterRoutes` that registers the routes. 
Use this to figure out the routes for the microservices. Microservices are proxied through Traefik to allow easy usage
at a single url.
ex:
  `/api/user/*` -> `http://user-api:3000/api/user/*`
  `/api/orders/*` -> `http://order-api:3000/api/orders/*`
  `/api/inventory/*` -> `http://inventory-api:3000/api/inventory/*`


### Missing features:
[ ] seeding
[ ] response for order doesnt include books in restapi
[ ] JWT verification
[ ] Fix up user microservice Rest version to use `RegisterRoutes`
