protoc --proto_path=grpc/proto --go_out=grpc/build --go_opt=paths=source_relative --go-grpc_out=grpc/build --go-grpc_opt=paths=source_relative grpc/proto/\*.proto
export $(cat .env | xargs)

```mermaid
flowchart LR
    Client --> Router
    Router --> Controller
    subgraph Api
        Controller --> UseCase
        UseCase --> Repository
    end
    Database --> Repository
    Domain --> Api
```

```mermaid
sequenceDiagram
    create actor C as Client
    PublicRouter ->> C: Response
    C ->> PublicRouter: Request
    PublicRouter ->> Controller: GET
```

```mermaid
sequenceDiagram
    create actor C as Client
    ProtectedRouter ->> C: Response
    C ->> ProtectedRouter: Request
    ProtectedRouter ->> AuthenticationMiddleware: I am authenticated?
    alt is not_authenticated
        AuthenticationMiddleware->> ProtectedRouter: Error
    else is well
        AuthenticationMiddleware ->> Controller: GET
        Controller ->> ProtectedRouter: Response
    end
    ProtectedRouter ->> C: Response
```
