package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Define command-line flags for source symbol, target symbol, and ratio
	sourceSymbol := flag.String("source", "VIX", "The source symbol to be replaced")
	targetSymbol := flag.String("target", "SPX", "The target symbol to replace with")
	ratio := flag.Float64("ratio", 408.79, "The conversion ratio")
	original := flag.String("original", "$VIX: 1D Exp Move Min, 12.27, HVL 1WTE & Put Support 1WTE & GEX Level 2, 12.5, GEX Level 0 & GEX Level 3 & Put Support, 13, GEX Level 1 & GEX Level 4, 13.5, 1D Exp Move Max, 13.57, HVL, 14.5, Call Resistance 1WTE & Gamma Wall 1WTE, 16, Call Resistance, 20", "The original quote to be modified")
	println(ratio)
	// Parse the command-line flags
	flag.Parse()

	// Split the original string by ":"
	parts := strings.SplitN(*original, ":", 2)

	// Replace the initial part of the string
	if strings.HasPrefix(parts[0], fmt.Sprintf("$%s", *sourceSymbol)) {
		parts[0] = strings.Replace(parts[0], fmt.Sprintf("$%s", *sourceSymbol), fmt.Sprintf("$%s", *targetSymbol), 1)
	}

	// Print the modified string
	fmt.Println(strings.Join(parts, ":"))

	// Split the second part of the string into pairs using comma as a delimiter
	pairs := strings.Split(parts[1], ",")

	//fmt.Printf("%v", pairs)
	// Create a map to store the key-value pairs and a slice to maintain the order of keys
	keyValuePairs := make(map[string]string)
	orderedKeys := make([]string, 0)

	// Process each pair separately
	i := 0
	key := ""
	for _, pair := range pairs {
		i++
		fmt.Printf("%d: %s\n", i, pair)
		if i%2 != 0 {
			pair = strings.TrimSpace(pair)
			key = pair
			//println(key)
		} else {
			pair = strings.TrimSpace(pair)
			keyValuePairs[key] = pair
			orderedKeys = append(orderedKeys, key)
		}
	}

	// Print the key-value pairs in the order of insertion
	for _, key := range orderedKeys {
		fmt.Printf("Key: %s, Value: %s\n", key, keyValuePairs[key])
	}

	// Create a new map to store the modified key-value pairs
	modifiedKeyValuePairs := make(map[string]float64)

	// Iterate over the ordered keys
	for _, key := range orderedKeys {
		// Get the original value and convert it to float64
		originalValue, err := strconv.ParseFloat(keyValuePairs[key], 64)
		if err != nil {
			fmt.Printf("Error converting value to float: %v\n", err)
			continue
		}

		// Calculate the modified value
		modifiedValue := originalValue * *ratio

		// Create the modified key
		modifiedKey := fmt.Sprintf("%s %s", *sourceSymbol, key)

		// Insert the modified key-value pair into the map
		modifiedKeyValuePairs[modifiedKey] = modifiedValue
	}

	// Print the modified key-value pairs
	for key, value := range modifiedKeyValuePairs {
		fmt.Printf("Key: %s, Value: %f\n", key, value)
	}

	// Create a slice to store the key-value pairs as strings
	keyValueStrings := make([]string, 0)

	// Iterate over the modified key-value pairs
	for key, value := range modifiedKeyValuePairs {
		// Create a string with the key and value separated by a comma
		keyValueString := fmt.Sprintf("%s, %.2f", key, value)

		// Append the string to the slice
		keyValueStrings = append(keyValueStrings, keyValueString)
	}

	// Join the strings in the slice with a comma to create a single string
	mapString := strings.Join(keyValueStrings, ", ")

	// Print the string
	fmt.Printf("$%s: %s\n", *targetSymbol, mapString)

}
