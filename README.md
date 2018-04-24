# Neo4j vs PostgreSQL

Author: Danny Nielsen.

Assignment description can be found [here](https://github.com/datsoftlyngby/soft2018spring-databases-teaching-material/blob/master/assignments/Neo4J%20Exercise.ipynb?short_path=55f0528).

## Setup

I've chosen to run this experiment in a small program written in [Go](https://golang.org/).

In order to replicate the setup and run my program, there are a few steps to follow first.

### Dependencies

#### The databases

As mentioned, this is a Go program. That means you have to have Go installed on the machine you want to run the program on. 

Also, the program pulls its data from a **Neo4j** database instance and a **PostgreSQL** database instance. Both these instances are required to contain the large "node and edges" dataset provided [here](https://github.com/datsoftlyngby/soft2018spring-databases-teaching-material/raw/master/data/archive_graph.tar.gz).

**Neo4j**: The program is assuming that a Neo4j instance is running on localhost with the port 7687 with a username called "neo4j" and the password "class". If that doesn't match your setup, you can change the connection string in the `neo4jdb.go` file:

```go
const connString = "bolt://neo4j:class@localhost:7687"
```

**PostgreSQL**: The program is also assuming that a Postgres instance is running on your localhost with the port 5432 with a username called "postgres" and the database name "postgres". If that doesn't match your setup, it can be changed from the connection string in the `psqldb.go` file:

```go
const connString = "user=postgres dbname=postgres sslmode=disable"
```

Note that my table names in Postgres are called `person` and `endorsement` like so:

```sql
			List of relations
 Schema |    Name     | Type  |  Owner   
--------+-------------+-------+----------
 public | endorsement | table | postgres
 public | person      | table | postgres
```

Their attributes and relations look like this, respectively:

Endorsement:

```sql
                Table "public.endorsement"
    Column     |  Type   | Collation | Nullable | Default 
---------------+---------+-----------+----------+---------
 person_one_id | integer |           |          | 
 person_two_id | integer |           |          | 
Foreign-key constraints:
    "endorsement_person_one_id_fkey" FOREIGN KEY (person_one_id) REFERENCES pers
on(id)
    "endorsement_person_two_id_fkey" FOREIGN KEY (person_two_id) REFERENCES pers
on(id)
```

Person:

```sql
                     Table "public.person"
  Column  |       Type        | Collation | Nullable | Default 
----------+-------------------+-----------+----------+---------
 id       | integer           |           | not null | 
 name     | character varying |           |          | 
 job      | character varying |           |          | 
 birthday | date              |           |          | 
Indexes:
    "person_pkey" PRIMARY KEY, btree (id)
    "id_index" btree (id)
Referenced by:
    TABLE "endorsement" CONSTRAINT "endorsement_person_one_id_fkey" FOREIGN KEY 
(person_one_id) REFERENCES person(id)
    TABLE "endorsement" CONSTRAINT "endorsement_person_two_id_fkey" FOREIGN KEY 
(person_two_id) REFERENCES person(id)
```

There is a big chance that your setup and attribute names aren't the same. It's possible to fix the query details in the program to match your setup. 

#### Go packages

Go also needs a couple of dependency packages in order to run the Neo4j and Postgres drivers.

To install the Neo4j Bolt driver, open your terminal and run:

```bash
$ go get github.com/johnnadratowski/golang-neo4j-bolt-driver
```

To install the pq (Postgres) driver, open your terminal and run:

```bash
$ go get github.com/lib/pq
```

### Run the program

With all of the above dependencies above settled, it's very simple to run the program.

1. Open your terminal and simply clone my project with git:

   ```bash
   $ git clone https://github.com/dbdness/neo4j-vs-psql.git
   ```

2. Change working directory to the root of the project:

   ```bash
   $ cd neo4j-vs-psql/
   ```

3. Build the Go program:

   ```bash
   $ go build
   ```

4. Execute the desired query on the built Go file. Simply state what database you want to benchmark, and what depth of endorsement you want to test:

   ```bash
   $ ./neo4jvspsql --neo4j 3
   ```

   The above query will launch the experiment with the Neo4j database, and with depth three of endorsement.

   Likewise, this query will launch the experiment with the PostgreSQL database, and with depth 5 of endorsement:

   ```bash
   $ ./neo4jvspsql --psql 5
   ```

   You can choose any depth you want. 

   An example output looks like this:

   ```bash
   $ ./neo4jvspsql --psql 2

   Opening PSQL connection...
   Starting PSQL benchmarks:
   ....................
   Average time for depth 2:
   1.7430000000000003
   Median time for depth 2:
   1.7650000000000001
   ```

## My results: execution times

All execution times are in seconds.

### Neo4j

| Depth       | Average | Median |
| ----------- | :-----: | -----: |
| Depth one   |  0.38   |   0.37 |
| Depth two   |  0.37   |   0.37 |
| Depth three |  0.40   |   0.41 |
| Depth Four  |  0.70   |   0.82 |
| Depth Five  |  7.33   |  10.45 |

###PostgreSQL

| Depth       | Average | Median |
| ----------- | :-----: | -----: |
| Depth one   |  0.75   |   0.67 |
| Depth two   |  1.75   |   1.75 |
| Depth three |  2.90   |   3.02 |
| Depth Four  |  5.30   |   5.64 |
| Depth Five  |   8.9   |   9.76 |

## Evaluation

As seen above, the query times aren't really that long in general. This is most likely due to the 20 random names I fetched from the dataset. Some nodes doesn't have a lot of data in their endorsement depths. The results could look completely different with 20 new random nodes.

If we look at the results, Neo4j wins with the average execution times. This is expected, because Neo4j is designed to handle these types of relationships between nodes in a very unique way. Relational DBMS's such as PostgreSQL are not built for this kind of node interaction. That said, Postgres did yield some very impressive results in this test, being only about 1½ seconds slower than Neo4j on average in depth 5. To be fair, PostgreSQL is one of the most popular relational DBMS out there, due to it's great performance and general usage.

It's also important to mention that I added indexes on both databases. This is also very important for performance, and especially on large datasets and complicated queries such as these ones. Maybe this is the reason why they both perform so well. A lot of different variables can have an impact on the performance times, so the above results should not be taken too seriously.

Either way, this was a very interesting experiment all in all.
