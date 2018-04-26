package main

import (
	"./neo4jdb"
	"./psqldb"
	"log"
	"fmt"
	"time"
	"database/sql"
	"github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"os"
	"strconv"
	"math"
	"sort"
)

// Randomly extracted names from the database.
var randomNames = []string{
	"Delena Linahan",
	"Claude Beel",
	"Svetlana Pongkhamsing",
	"Pedro Panagakos",
	"Norbert Elstad",
	"Darnell Stlouis",
	"Lenny Hendryx",
	"Gerald Kotcher",
	"Ana Lourence",
	"Rudolph Beaushaw",
	"Oma Gett",
	"Tom Yerger",
	"Terrilyn Steiniger",
	"Allene Lorch",
	"Jody Hyndman",
	"Pablo Gunnerson",
	"Coralie Ciminera",
	"Ambrose Moulder",
	"Danae Schnoke",
	"Kenna Alston",
}

func main() {
	//First command is also an argument
	if len(os.Args) != 3{
		printUsage()
		return
	}

	userCommand := os.Args[1]
	depthCommand := os.Args[2]

	switch userCommand {
	case "--neo4j":
		fmt.Println("Opening Neo4j connection...")
		conn, err := neo4jdb.Open()
		if err != nil {
			log.Panic(err)
		}
		defer conn.Close()

		depth, err := strconv.Atoi(depthCommand)
		if err != nil {
			log.Panic("Error: Wrong depth input")
			fmt.Println("Shutting down...")
		}

		//names := neo4jdb.GetRandomName(20)
		fmt.Println("Starting Neo4j benchmarks:")
		neo4jBenchmark(conn, depth, randomNames)
		break
	case "--psql":
		fmt.Println("Opening PSQL connection...")
		conn, err := psqldb.Open()
		if err != nil {
			log.Panic(err)
		}
		defer conn.Close()

		depth, err := strconv.Atoi(depthCommand)
		if err != nil {
			log.Panic("Error: Wrong depth input")
			fmt.Println("Shutting down...")
		}

		//names := psqldb.GetRandomName(20)
		fmt.Println("Starting PSQL benchmarks:")
		psqlBenchmark(conn, depth, randomNames)
		break
	 default:
	 	printUsage()
		break
	}

}

func neo4jBenchmark(conn golangNeo4jBoltDriver.Conn, depth int, names []string) {
	if conn == nil {
		log.Panic("Neo4j connection null. Did you forget to call Open()?")
	}

	durations := make([]float64, 0)

	for _, name := range names {
		startTime := time.Now()
		neo4jdb.GetDepthCount(name, depth)
		fmt.Print(".") //Activity monitor
		elapsed := time.Since(startTime)
		durations = append(durations, round(elapsed.Seconds()))
	}

	//log.Println(durations)
	fmt.Printf("\nAverage time for depth %d:\n", depth)
	average := calcAverage(durations)
	fmt.Println(average)

	fmt.Printf("Median time for depth %d:\n", depth)
	median := calcMedian(durations)
	fmt.Println(median)
}

func psqlBenchmark(conn *sql.DB, depth int, names []string) {
	if conn == nil {
		log.Panic("PSQL connection null. Did you forget to call Open()?")
	}

	durations := make([]float64, 0)

	for _, name := range names {
		startTime := time.Now()
		psqldb.GetDepthCount(name, depth)
		fmt.Print(".") //Activity monitor
		elapsed := time.Since(startTime)
		durations = append(durations, round(elapsed.Seconds()))
	}

	//log.Println(durations)
	fmt.Printf("\nAverage time for depth %d:\n", depth)
	average := calcAverage(durations)
	fmt.Println(average)

	fmt.Printf("Median time for depth %d:\n", depth)
	median := calcMedian(durations)
	fmt.Println(median)
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

	sort.Float64s(slice)

	middle := len(slice) / 2

	if len(slice)%2 == 1 {
		median = slice[middle]
	} else {
		median = (slice[middle-1] + slice[middle]) / 2
	}

	return median
}

func round(number float64) float64{
	rounded := math.Round(number*100)/100
	return rounded
}

func printUsage(){
	fmt.Println("usage:\n--neo4j [depth]\n--psql [depth]\n\nExample: --neo4j 3")
}
