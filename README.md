# MobsX-MC 🧠🧟

**MobsX-MC** is a standalone, high-performance AI and Pathfinding library for Minecraft Bedrock servers written in Go.

It provides a modular "Brain" system that allows developers to give life to entities through a decoupled architecture of **Sensors**, **Behaviors**, and **Navigation**.

## Features 🚀

- **Independent**: Works with any Minecraft Bedrock server written in Go through a clean API interface.
- **GPS Navigation (A*)**: High-speed pathfinding algorithm that calculates 3D routes avoiding obstacles.
- **Priority Brain**: Entities decide their actions based on importance (e.g., *Attacking* > *Wandering*).
- **Modular Behaviors**: Easy to add new logic like Chase, Flee, Idle, or Breed.
- **Asynchronous**: Designed to run hundreds of mobs simultaneously using Go's concurrency.

## Installation 📦

```bash
go get github.com/AssassinGhostYT/MobsX-MC
```

## How to use 🛠️

### 1. Implement the API
Make your entity and world objects compatible with the `api.Entity` and `api.World` interfaces.

### 2. Connect the Brain
```go
import (
    "github.com/AssassinGhostYT/MobsX-MC"
    "github.com/AssassinGhostYT/MobsX-MC/behavior"
    "github.com/AssassinGhostYT/MobsX-MC/sensor"
)

// Create the AI components
brain := mobsx.NewBrain()
navigator := mobsx.NewNavigator(myZombie, myWorld)
playerScanner := &sensor.PlayerSensor{Range: 16}

// Add life!
brain.AddSensor(playerScanner)
brain.AddBehavior(behavior.NewAttack(playerScanner, navigator))
brain.AddBehavior(behavior.NewWander(navigator, 10))

// In your entity Tick function:
brain.Tick(myZombie, myWorld)
```

Developed with ❤️ by **AssassinGhostYT**.
