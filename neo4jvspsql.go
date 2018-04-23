package main

import (
	"./neo4jdb"
	"./psqldb"
	"log"
	"fmt"
	"time"
	"database/sql"
	"github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

func main() {
	conn, err := psqldb.Open()
	if err != nil{
		log.Panic(err)
	}
	defer conn.Close()

	//Doesn't matter what db we get the random names from
	names := psqldb.GetRandomName(20)
	//names := neo4jdb.GetRandomName(20)

	//fmt.Println("Starting Neo4j benchmarks...")
	//neo4jBenchmark(1)
	fmt.Println("Starting PSQL benchmarks...")
	psqlBenchmark(conn,3, names)
}

func neo4jBenchmark(conn golangNeo4jBoltDriver.Conn, depth int, names []string) {
	durations := make([]float64, 0)

	for _, name := range names {
		startTime := time.Now()
		neo4jdb.GetDepthCount(name, depth)
		fmt.Print(".")
		elapsed := time.Since(startTime)
		durations = append(durations, elapsed.Seconds())
	}

	log.Println(durations)
	fmt.Printf("\nAverage time for depth %d:\n", depth)
	average := calcAverage(durations)
	fmt.Println(average)

	fmt.Printf("Median time for depth %d:\n", depth)
	median := calcMedian(durations)
	fmt.Println(median)
}

func psqlBenchmark(conn *sql.DB, depth int, names []string){
	durations := make([]float64, 0)

	for _, name := range names {
		startTime := time.Now()
		psqldb.GetDepthCount(name, depth)
		fmt.Print(".")
		elapsed := time.Since(startTime)
		durations = append(durations, elapsed.Seconds())
	}

	log.Println(durations)
	fmt.Printf("\nAverage time for depth %d:\n", depth)
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
