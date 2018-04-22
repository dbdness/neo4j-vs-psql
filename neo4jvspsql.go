package main

import (
	neo4jdb "./neo4jdb"
	"log"
	"fmt"
)

func main() {
	conn, err := neo4jdb.Open()
	if err != nil{
		log.Panic(err)
	}
	defer conn.Close()

	count := neo4jdb.GetDepthCount("Genia Crist", "1")

	fmt.Println(count)
	
}
