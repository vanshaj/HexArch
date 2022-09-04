**HEXARCH**
Components are decoupled in the hex architecture so we can add different components to our application and derive the application with that component
We have a hexagon = Domain, Application, Framework like a hexagon (starting from inside to outside)
We will have different adapters like DB, GRPC and HTTP which we can plug in to our outlets(Ports)

We have different adapters 
1. Driving adapters like HTTP which will drive our application , perform actions on our application. (Primary)
2. Driven adapters like DB which will be driven by our application. (Secondary)
Driving placed on left side of the hexagon and Driven placed on the right

**Layers of Hexagon Architecture** (Inside to outside)
1. Domain Layer - contains our domain logic or application logic.
2. Application Layer - will orchestrate the use of our domain code and adapter request from the framework layer by sitting between framework and the domain.
3. Framework Layer - will provide logic for the outside component just as Database and GRPC adapters to interact with our application.

Outside Layers Depend on the inside layers. Dependencies point inward. So we make inject our dependencies into the layers. DB instance is instantiated in the framework layer and injected into the application layer and domain layer. 

**Code Structure**
adapters *will contains the code for all of our layers*
--app  *will contains the application code*
--core  *will contains the core domain code*
--framework  *will contains the framework code*
----left
------grpc
------http
----right
-------db