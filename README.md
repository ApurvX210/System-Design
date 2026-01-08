# System Design - Low Level Design Patterns

This repository contains implementations of various system design problems with a focus on **Low Level Design (LLD)** patterns. Each project explores different design principles, patterns, and best practices to strengthen understanding of software architecture and design.

## 📚 Purpose

The goal of this repository is to:
- Practice and master Low Level Design patterns
- Implement real-world system design problems
- Explore different design principles (SOLID, DRY, etc.)
- Understand object-oriented design concepts
- Build scalable and maintainable code structures

## 🗂️ Project Structure

```
System Design/
├── Parking Lot LLD/          # Parking Lot Management System
│   ├── entities/             # Domain entities
│   └── main.go              # Entry point
├── [Future Projects]/        # More system designs to come
└── README.md                # This file
```

## 🚗 Current Project: Parking Lot LLD

### Overview
A parking lot management system that handles vehicle parking, spot allocation, and ticket generation. This project demonstrates:
- **Entity Design**: Core domain objects (Vehicle, ParkingSpot, Floor, Ticket)
- **Design Patterns**: Strategy, Factory, Observer patterns
- **SOLID Principles**: Single Responsibility, Open/Closed, etc.
- **State Management**: Parking spot states (available, occupied, reserved)

### Features
- Multiple vehicle types (Car, Motorcycle, Truck, etc.)
- Multi-floor parking structure
- Different parking spot types per vehicle
- Ticket generation and tracking
- Parking fee calculation
- Spot allocation and deallocation

### Entity Structure
- **Vehicle**: Interface for different vehicle types
- **ParkingSpot**: Represents individual parking spaces
- **Floor**: Contains multiple parking spots
- **Ticket**: Tracks parking sessions

### Design Patterns Used
- **Strategy Pattern**: Different pricing strategies for different vehicle types
- **Factory Pattern**: Creating appropriate parking spots for vehicles
- **Observer Pattern**: Notifications for spot availability

## 🎯 Learning Objectives

Through this project, you'll learn:
1. How to model real-world systems using OOP principles
2. Design patterns and when to apply them
3. Entity relationship modeling
4. State management in complex systems
5. API design for system interactions

## 🚀 Getting Started

### Prerequisites
- Go 1.19+ installed
- Basic understanding of Go programming
- Familiarity with OOP concepts

### Running the Project
```bash
cd "Parking Lot LLD"
go run main.go
```

## 📝 Daily Progress

### Day 1: Parking Lot LLD
- ✅ Project structure setup
- ✅ Entity definitions started
- 🔄 Core functionality implementation (in progress)


## 🤝 Contributing

This is a personal learning repository. Feel free to fork and use it for your own learning journey!

## 📄 License

This project is for educational purposes.

---

**Happy Learning! 🎓**

*Building better software, one design at a time.*



