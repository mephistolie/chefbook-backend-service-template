# Service Template Agents Guide

This directory is the baseline template for backend microservices.

## Purpose

- Use it as the reference structure when creating or reshaping services.
- Keep the template generic; do not add domain-specific behavior here.

## Expected Structure

- `api` for inter-service contracts
- `cmd` for entrypoints
- `internal` for service implementation
- `migrations` for storage migrations
- `deployments` for deployment manifests
- `scripts` for local automation

## Working Rules

- Changes here should improve the default service blueprint, not patch one specific service indirectly.
- If a template change should also be applied to existing services, update them explicitly rather than assuming inheritance.
