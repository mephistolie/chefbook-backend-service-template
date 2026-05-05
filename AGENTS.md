# Service Template Agents Guide

This directory is the baseline template for backend microservices.

## Purpose

- Use it as the reference structure when creating or reshaping services.
- Keep the template generic; do not add domain-specific behavior here.

## Expected Structure

- `api` for inter-service contracts
- `cmd` for entrypoints
- `internal/app` for composition, runtime startup, health checks, and graceful shutdown
- `internal/config` for runtime configuration
- `internal/entity` for service-local domain data and failures
- `internal/transport` for external adapters such as gRPC or AMQP consumers
- `internal/transport/dependencies/service` for transport-facing service interfaces
- `internal/service` for use cases and domain orchestration
- `internal/service/dependencies/repository` for service-facing repository interfaces
- `internal/repository` for Postgres, gRPC, AMQP, S3, or other infrastructure adapters
- `migrations` for storage migrations
- `deployments` for deployment manifests
- `scripts` for local automation

## Working Rules

- Changes here should improve the default service blueprint, not patch one specific service indirectly.
- If a template change should also be applied to existing services, update them explicitly rather than assuming inheritance.
- Keep dependency direction as `transport -> service -> repository interfaces <- repository implementations`.
- Keep business rules out of transports, SQL out of services, and protobuf/database DTOs out of entities.
- Pass `context.Context` through request and message paths; use context-aware SQL and external calls.
