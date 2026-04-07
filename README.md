# RiftBattle Demo

A small Go RPG combat codebase designed for SAAYN demos and labs.

## Purpose

This repo is intentionally small enough to understand quickly, but large enough to produce meaningful SAAYN nodes, candidate matches, and graph relationships.

## Suggested SAAYN demo scenarios

- Prevent healing while the target is poisoned
- Increase fire damage against frozen enemies
- Cap critical hit bonus for heavy weapons
- Reduce armor mitigation for magic damage
- Add a stun check before a combat turn resolves
- Ensure burn damage cannot defeat a target below 1 HP in training mode

## Run

```bash
go run .
```

## Test

```bash
go test ./...
```
