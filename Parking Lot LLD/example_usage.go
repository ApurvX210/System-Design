package main

import (
	"fmt"
	"time"

	"./entities"
	"./service"
)

func main() {
	// Example: Create a parking lot with 3 floors
	// Each floor configuration: [carCapacity, motorcycleCapacity, truckCapacity]
	floorsConfig := [][3]int{
		{10, 20, 5},  // Floor 1: 10 cars, 20 motorcycles, 5 trucks
		{15, 25, 3},  // Floor 2: 15 cars, 25 motorcycles, 3 trucks
		{20, 30, 2},  // Floor 3: 20 cars, 30 motorcycles, 2 trucks
	}

	// Pricing configuration (price per hour in cents)
	pricing := map[entities.VehicleType]int{
		entities.MOTORCYCLE: 10, // $0.10 per hour
		entities.CAR:        20, // $0.20 per hour
		entities.TRUCK:      50, // $0.50 per hour
	}

	// Create parking lot service
	parkingLot := service.NewParkingLotService(floorsConfig, pricing)

	// Example 1: Park a car
	fmt.Println("=== Example 1: Parking a Car ===")
	car := entities.NewCar("ABC-1234")
	ticket1, err := parkingLot.ParkVehicle(car)
	if err != nil {
		fmt.Printf("Error parking car: %v\n", err)
	} else {
		fmt.Printf("Car parked successfully!\n")
		fmt.Printf("Ticket ID: %s\n", ticket1.ID)
		fmt.Printf("Floor: %d, Spot: %d\n", ticket1.FloorID, ticket1.SpotID)
		fmt.Printf("Entry Time: %s\n", ticket1.EntryTime.Format(time.RFC3339))
	}

	// Example 2: Park a motorcycle
	fmt.Println("\n=== Example 2: Parking a Motorcycle ===")
	motorcycle := entities.NewMotorCycle("XYZ-5678")
	ticket2, err := parkingLot.ParkVehicle(motorcycle)
	if err != nil {
		fmt.Printf("Error parking motorcycle: %v\n", err)
	} else {
		fmt.Printf("Motorcycle parked successfully!\n")
		fmt.Printf("Ticket ID: %s\n", ticket2.ID)
		fmt.Printf("Floor: %d, Spot: %d\n", ticket2.FloorID, ticket2.SpotID)
	}

	// Example 3: Park a truck
	fmt.Println("\n=== Example 3: Parking a Truck ===")
	truck := entities.NewTruck("TRK-9999")
	ticket3, err := parkingLot.ParkVehicle(truck)
	if err != nil {
		fmt.Printf("Error parking truck: %v\n", err)
	} else {
		fmt.Printf("Truck parked successfully!\n")
		fmt.Printf("Ticket ID: %s\n", ticket3.ID)
		fmt.Printf("Floor: %d, Spot: %d\n", ticket3.FloorID, ticket3.SpotID)
	}

	// Example 4: Check parking lot status
	fmt.Println("\n=== Example 4: Parking Lot Status ===")
	status := parkingLot.GetParkingLotStatus()
	fmt.Printf("Total Active Tickets: %d\n", status.TotalActiveTickets)
	for _, floor := range status.Floors {
		fmt.Printf("\nFloor %d:\n", floor.FloorID)
		fmt.Printf("  Car Spots: %d occupied, %d vacant (Total: %d)\n",
			floor.CarSpots.Occupied, floor.CarSpots.Vacant, floor.CarSpots.Total)
		fmt.Printf("  Motorcycle Spots: %d occupied, %d vacant (Total: %d)\n",
			floor.MotorcycleSpots.Occupied, floor.MotorcycleSpots.Vacant, floor.MotorcycleSpots.Total)
		fmt.Printf("  Truck Spots: %d occupied, %d vacant (Total: %d)\n",
			floor.TruckSpots.Occupied, floor.TruckSpots.Vacant, floor.TruckSpots.Total)
	}

	// Example 5: Try to park same vehicle twice (should fail)
	fmt.Println("\n=== Example 5: Attempting to Park Same Vehicle Twice ===")
	_, err = parkingLot.ParkVehicle(car)
	if err != nil {
		fmt.Printf("Expected error: %v\n", err)
	}

	// Example 6: Wait a bit and check current price
	fmt.Println("\n=== Example 6: Checking Current Price (Vehicle Still Parked) ===")
	time.Sleep(2 * time.Second) // Simulate time passing
	ticket1, _ = parkingLot.GetTicket(ticket1.ID)
	currentPrice := ticket1.CalculatePrice()
	fmt.Printf("Ticket %s - Current Price: %d cents (%.2f hours parked)\n",
		ticket1.ID, currentPrice, ticket1.GetDuration().Hours())

	// Example 7: Unpark the car
	fmt.Println("\n=== Example 7: Unparking the Car ===")
	finalTicket, finalPrice, err := parkingLot.UnparkVehicle(ticket1.ID)
	if err != nil {
		fmt.Printf("Error unparking: %v\n", err)
	} else {
		fmt.Printf("Vehicle unparked successfully!\n")
		fmt.Printf("Final Price: %d cents ($%.2f)\n", finalPrice, float64(finalPrice)/100)
		fmt.Printf("Total Duration: %s\n", finalTicket.GetDuration().Round(time.Second))
		fmt.Printf("Exit Time: %s\n", finalTicket.ExitTime.Format(time.RFC3339))
	}

	// Example 8: Check status after unparking
	fmt.Println("\n=== Example 8: Status After Unparking ===")
	status = parkingLot.GetParkingLotStatus()
	fmt.Printf("Total Active Tickets: %d\n", status.TotalActiveTickets)
	fmt.Printf("Floor 1 Car Spots: %d occupied, %d vacant\n",
		status.Floors[0].CarSpots.Occupied, status.Floors[0].CarSpots.Vacant)

	// Example 9: Try to unpark with invalid ticket
	fmt.Println("\n=== Example 9: Attempting to Unpark with Invalid Ticket ===")
	_, _, err = parkingLot.UnparkVehicle("INVALID-TICKET")
	if err != nil {
		fmt.Printf("Expected error: %v\n", err)
	}

	// Example 10: Get active ticket by vehicle
	fmt.Println("\n=== Example 10: Get Active Ticket by Vehicle Number ===")
	activeTicket, err := parkingLot.GetActiveTicketByVehicle("XYZ-5678")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Found active ticket for XYZ-5678: %s\n", activeTicket.ID)
		fmt.Printf("Parked at Floor %d, Spot %d\n", activeTicket.FloorID, activeTicket.SpotID)
	}
}

