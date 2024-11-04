package main

import (
	"fmt"
	"sort"
)

// Event represents a scheduled event with start and end times.
type Event struct {
	StartTime int
	EndTime   int
}

// Scheduler manages a list of scheduled events.
type Scheduler struct {
	events []Event
}

// AddEvent adds a new event if it does not overlap with existing events.
func (s *Scheduler) AddEvent(event Event) bool {
	for _, existingEvent := range s.events {
		if event.StartTime < existingEvent.EndTime && event.EndTime > existingEvent.StartTime {
			return false // Overlap detected
		}
	}
	s.events = append(s.events, event)
	return true
}

// GetEvents returns all scheduled events sorted by start time.
func (s *Scheduler) GetEvents() []Event {
	sort.Slice(s.events, func(i, j int) bool {
		return s.events[i].StartTime < s.events[j].StartTime
	})
	return s.events
}

func main() {
	scheduler := Scheduler{}

	for {
		fmt.Println("\nScheduler Menu:")
		fmt.Println("1. Add Event")
		fmt.Println("2. View Events")
		fmt.Println("3. Exit")

		var choice int
		fmt.Print("Choose an option (1-3): ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var startTime, endTime int
			fmt.Print("Enter start time (0-23): ")
			fmt.Scan(&startTime)
			fmt.Print("Enter end time (0-23): ")
			fmt.Scan(&endTime)

			// Input validation
			if startTime < 0 || startTime > 23 || endTime < 0 || endTime > 23 {
				fmt.Println("Error: Times must be between 0 and 23.")
				continue
			}
			if startTime >= endTime {
				fmt.Println("Error: Start time must be less than end time.")
				continue
			}

			event := Event{StartTime: startTime, EndTime: endTime}
			if scheduler.AddEvent(event) {
				fmt.Println("Event added successfully.")
			} else {
				fmt.Println("Error: Event overlaps with an existing event.")
			}

		case 2:
			events := scheduler.GetEvents()
			if len(events) == 0 {
				fmt.Println("No scheduled events.")
			} else {
				fmt.Println("Scheduled Events:")
				for _, event := range events {
					fmt.Printf("%d - %d\n", event.StartTime, event.EndTime)
				}
			}

		case 3:
			fmt.Println("Exiting the scheduler.")
			return

		default:
			fmt.Println("Error: Invalid option. Please choose again.")
		}
	}
}
