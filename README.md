# ChefBook Backend Service Template

This module is the template for new ChefBook backend services.

## Intended Shape

- `cmd/app` for the service runtime entrypoint.
- `cmd/migrations` for migration runtime.
- `internal/app` for process bootstrap.
- `internal/config` for runtime configuration.
- `internal/entity` for domain data structures, domain-specific enums, and service-local errors.
- `internal/transport/grpc` for gRPC server adapters.
- `internal/transport/dependencies/service` for the service-facing interfaces required by transports.
- `internal/service` for domain use cases.
- `internal/service/<domain>` for cohesive use-case groups when the service has several domains.
- `internal/service/mq` for inbound asynchronous message handling.
- `internal/service/dependencies/repository` for repository interfaces required by services.
- `internal/repository/postgres` for persistence.
- `internal/repository/grpc` for remote service clients when needed.
- `internal/repository/amqp` or shared `common/mq` adapters for outbound RabbitMQ publishing when needed.
- `internal/repository/s3` or other infrastructure-specific repositories when needed.
- `api/proto/contract/v1` for proto sources.
- `api/proto/implementation/v1` for generated gRPC code.
- `migrations/sql` for schema changes.
- `deployments/helm` for Kubernetes deployment.
- `scripts` for local generation and deployment helpers.

## Layering

New services should follow the same dependency direction as existing ChefBook services:

```text
cmd
  -> internal/app
    -> transport
      -> service
        -> repository interfaces
          <- repository implementations
```

`internal/app` wires concrete dependencies and owns process lifecycle. It should not contain business rules beyond startup, health checks, server registration, and graceful shutdown.

`internal/transport` translates external contracts to domain calls. gRPC and AMQP handlers should validate request shape, parse identifiers, map DTOs, pass `context.Context`, and delegate behavior to `internal/service`. Transport code should not contain domain decisions or SQL.

`internal/service` owns use cases and domain orchestration. Keep business decisions here: permissions, cross-repository workflows, outbox publication decisions, calls to remote services, and storage-specific error interpretation only after repositories have mapped raw infrastructure errors.

`internal/repository` owns infrastructure access. Postgres repositories should contain SQL, transactions, DTO scanning, and persistence-specific helpers. gRPC/S3/AMQP repositories should adapt external systems to service-level interfaces.

`internal/entity` is the shared domain language inside one service. It should stay free of transport protobuf types, database DTOs, and infrastructure client types.

## Interfaces

Prefer consumer-owned interfaces:

- transport-facing service interfaces live in `internal/transport/dependencies/service`;
- service-facing repository interfaces live in `internal/service/dependencies/repository`;
- concrete implementations live in their own packages and are wired in `internal/app` or `internal/service.New`.

Keep interfaces grouped by use-case boundary rather than by storage table. For example, a service with recipes and collections can expose separate `Recipe` and `Collection` interfaces, while one small service can keep a single `User` interface.

Avoid package-level "god" interfaces when a dependency is used by only one cohesive service component. Split when a component only needs a smaller capability set, or when tests become forced to mock unrelated methods.

## Context

Every request or message handling path should accept and propagate `context.Context`:

- gRPC handlers use the incoming method context;
- HTTP handlers use `Request.Context()`;
- AMQP/background handlers create a bounded root context for one message or one loop iteration;
- repositories use `QueryContext`, `QueryRowContext`, `ExecContext`, `PrepareContext`, and `BeginTx`.

Use `context.Background()` only at process boundaries, compatibility wrappers, or explicit background jobs. Prefer `context.WithTimeout` for message processing and external calls.

## Naming

- Use package names for the role: `grpc`, `postgres`, `mq`, `recipe`, `collection`, `user`.
- Keep one aggregate `internal/service.Service` that embeds or exposes use-case interfaces used by transports.
- Name concrete service structs simply as `Service` inside their domain package.
- Name concrete repositories as `Repository` inside their infrastructure package.
- Keep DTO conversion in `transport/.../dto` or `repository/.../dto`; do not put protobuf/database conversion into domain services.

## Usage Guidance

- Copy this shape when adding a new service.
- Replace template naming before publishing a new module.
- Keep domain-specific behavior in the new service, not in this template or `common`.

## Current Assessment

The current ChefBook services mostly follow this structure well:

- business workflows are separated from gRPC handlers;
- Postgres access is isolated in repository packages;
- service code depends on small repository interfaces rather than concrete SQL packages;
- larger domains are split into cohesive service packages such as `recipe`, `collection`, `session`, and `password`.

The main improvement areas are:

- make AMQP and outbox background loops cancellable through lifecycle contexts;
- use bounded contexts for each consumed message;
- keep shrinking broad repository interfaces where one service component only needs a small subset;
- avoid direct concrete infrastructure dependencies in services when an interface would make ownership clearer;
- keep the template synchronized with mature services when new patterns become standard.
