# Neo4j vs PostgreSQL

## Dependencies

Neo4j Bolt driver:

```bash
go get github.com/johnnadratowski/golang-neo4j-bolt-driver
```

pq:

```
go get github.com/lib/pq
```

## Execution times

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

### ##Evaluation

Indexing on both dbs.





