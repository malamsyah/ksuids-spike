package main

import (
	"fmt"

	"github.com/segmentio/ksuid"
)

func generateKSUIDFirst15() string {
	// Generate a new KSUID
	newKSUID := ksuid.New()
	// Convert KSUID to base62 string and return the first 15 characters
	return newKSUID.String()[:15]
}

func validateUniqueness(count int) {
	ksuidMap := make(map[string]bool)
	duplicates := make([]string, 0)

	for i := 0; i < count; i++ {
		ksuidFirst15 := generateKSUIDFirst15()
		if _, exists := ksuidMap[ksuidFirst15]; exists {
			duplicates = append(duplicates, ksuidFirst15)
		} else {
			ksuidMap[ksuidFirst15] = true
		}
	}

	fmt.Printf("Total KSUIDs generated: %d\n", count)
	fmt.Printf("Unique KSUIDs generated: %d\n", len(ksuidMap))
	if len(duplicates) > 0 {
		fmt.Printf("Number of duplicates found: %d\n", len(duplicates))
	} else {
		fmt.Println("No duplicates found - all KSUIDs are unique!")
	}
}

func main() {
	validateUniqueness(10000000)
}
