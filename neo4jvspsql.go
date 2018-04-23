package main

import (
	"./neo4jdb"
	"./psqldb"
	"log"
	"fmt"
	"time"
)

func main() {
	//neo4jQuery()
	//psqlQuery()

	/*
	conn, err := psqldb.Open()
	if err != nil{
		log.Panic(err)
	}
	defer conn.Close()

	names := psqldb.GetRandomName(20)

	for _, name := range names{
		count := psqldb.GetDepthCount(name, 1)
		fmt.Println(count)
	}
	*/
	fmt.Println("Starting Neo4j benchmarks...")
	neo4jBenchmark(1)
}

func neo4jBenchmark(depth int) {
	conn, err := neo4jdb.Open()
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	durations := make([]float64, 0)
	names := neo4jdb.GetRandomName(20)

	for _, name := range names {
		startTime := time.Now()
		neo4jdb.GetDepthCount(name, depth)
		elapsed := time.Since(startTime)
		durations = append(durations, elapsed.Seconds())
	}

	log.Println(durations)
	fmt.Printf("Average time for depth %d:\n", depth)
	average := calcAverage(durations)
	fmt.Println(average)

	fmt.Printf("Median time for depth %d:\n", depth)
	median := calcMedian(durations)
	fmt.Println(median)
}

func neo4jQuery() {
	conn, err := neo4jdb.Open()
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	fmt.Println("Starting NEO4J query...")
	count := neo4jdb.GetDepthCount("Genia Crist", 5)

	fmt.Println(count)
}

func psqlQuery() {
	conn, err := psqldb.Open()
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	fmt.Println("Starting PSQL query...")
	count := psqldb.GetDepthCount("Genia Crist", 2)

	fmt.Println(count)
}

func trackTime(start time.Time) float64 {
	elapsed := time.Since(start)
	return elapsed.Seconds()
}

func calcAverage(slice []float64) float64 {
	var total float64
	for _, value := range slice {
		total += value
	}

	average := total / float64(len(slice))
	return average
}

func calcMedian(slice []float64) float64 {
	var median float64
	middle := len(slice) / 2

	if len(slice)%2 == 1 {
		median = slice[middle]
	} else {
		median = (slice[middle-1] + slice[middle]) / 2
	}

	return median
}
