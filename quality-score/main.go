package main

import "log"

func init() {
	log.SetFlags(0)
	log.SetPrefix("Error: ")
}

func main() {
	// starts the lambda
	//fmt.Println("Lambda started!")
	//lambda.Start(handleSchedule)

	// testing purposes
	handleSchedule()
}

// handleSchedule function to call once every week
func handleSchedule() {

}
