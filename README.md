# ChefBook Backend Service Template

This module is the template for new ChefBook backend services.

## Intended Shape

- `cmd/app` for the service runtime entrypoint.
- `cmd/migrations` for migration runtime.
- `internal/app` for process bootstrap.
- `internal/config` for runtime configuration.
- `internal/transport/grpc` for gRPC server adapters.
- `internal/service` for domain use cases.
- `internal/repository/postgres` for persistence.
- `internal/repository/grpc` for remote service clients when needed.
- `api/proto/contract/v1` for proto sources.
- `api/proto/implementation/v1` for generated gRPC code.
- `migrations/sql` for schema changes.
- `deployments/helm` for Kubernetes deployment.

## Usage Guidance

- Copy this shape when adding a new service.
- Replace template naming before publishing a new module.
- Keep domain-specific behavior in the new service, not in this template or `common`.
