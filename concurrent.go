package main

import (
	"fmt"
	"time"
)

func main() {
	//go say("world")
	//say("hello")

	//theMine := [5]string{"ore1", "ore2", "ore3"}
	//oreChan := make(chan string)
	//
	//// Finder
	//go func(mine [5]string) {
	//	for _, item := range mine {
	//		oreChan <- item //send
	//	}
	//}(theMine)
	//
	//// Ore Breaker
	//go func() {
	//	for i := 0; i < 3; i++ {
	//		foundOre := <-oreChan //receive
	//		fmt.Println("Miner: Received " + foundOre + " from finder")
	//	}
	//}()
	//<-time.After(time.Second * 5) // Again, ignore this for now

	//bufferedChan := make(chan string, 3)
	//
	//go func() {
	//	bufferedChan <-"first"
	//	fmt.Println("Sent 1st")
	//	bufferedChan <-"second"
	//	fmt.Println("Sent 2nd")
	//	bufferedChan <-"third"
	//	fmt.Println("Sent 3rd")
	//}()
	//
	//go func() {
	//	firstRead := <- bufferedChan
	//	fmt.Println("Receiving..")
	//	fmt.Println(firstRead)
	//	secondRead := <- bufferedChan
	//	fmt.Println(secondRead)
	//	thirdRead := <- bufferedChan
	//	fmt.Println(thirdRead)
	//}()
	//
	//<-time.After(time.Second * 1)

	//theMine := [5]string{"rock", "ore", "ore", "rock", "ore"}
	//oreChannel := make(chan string)
	//minedOreChan := make(chan string)
	//
	//// Finder
	//go func(mine [5]string) {
	//	for _, item := range mine {
	//		if item == "ore" {
	//			oreChannel <- item //send item on oreChannel
	//		}
	//	}
	//}(theMine)
	//
	//// Ore Breaker
	//go func() {
	//	for i := 0; i < 3; i++ {
	//		foundOre := <-oreChannel //read from oreChannel
	//		fmt.Println("From Finder:", foundOre)
	//		minedOreChan <-"minedOre" //send to minedOreChan
	//	}
	//}()
	//
	//// Smelter
	//go func() {
	//	for i := 0; i < 3; i++ {
	//		minedOre := <-minedOreChan //read from minedOreChan
	//		fmt.Println("From Miner:", minedOre)
	//		fmt.Println("From Smelter: Ore is smelted")
	//	}
	//}()
	//
	//<-time.After(time.Second * 5) // Again, you can ignore this

	//myChan := make(chan string)
	//
	//go func(){
	//	myChan <- "Message!"
	//}()
	//
	//select {
	//	case msg := <- myChan:
	//		fmt.Println(msg)
	//	default:
	//		fmt.Println("No Msg")
	//}
	//<-time.After(time.Second * 1)
	//
	//select {
	//	case msg := <- myChan:
	//		fmt.Println(msg)
	//	default:
	//		fmt.Println("No Msg")
	//}
}

func say(s string) {
	for i := 0; i< 5; i++ {
		time.Sleep(100 * time.Microsecond)
		fmt.Println(s)
	}
}
