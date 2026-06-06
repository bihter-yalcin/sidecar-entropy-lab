# Sidecar Entropy Lab

**Sidecar Entropy Lab** is a small Go project that demonstrates the **sidecar pattern** through a playful developer desk metaphor.

The project shows how a sidecar can reduce **architectural entropy** by moving operational concerns such as caching, request logging, and basic observability outside the main service.

> The desk may still get messy.  
> But the main service stays clean.

---

## Project Idea

This project contains two Go services:

### Desk Mess Service

The `desk-mess-service` represents the main application.

Its responsibility is simple: report the current state of a messy developer desk.

It returns data such as:

- coffee cups
- open browser tabs
- sticky notes
- loose cables
- unread messages
- entropy score

### Entropy Sidecar

The `entropy-sidecar` sits in front of the Desk Mess Service.

Clients call the sidecar instead of calling the main service directly.

The sidecar handles operational concerns such as:

- reverse proxying
- Redis-backed GET response caching
- `X-Cache: HIT / MISS` headers
- cache TTL
- request logging
- basic metrics

This keeps the main service focused on business logic.

---

## Architecture

Client → Entropy Sidecar → Desk Mess Service
             ↓
           Redis