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
| Depth one   |  0.47   |   0.44 |
| Depth two   |  0.63   |   0.46 |
| Depth three |  1.04   |   0.52 |
| Depth Four  |  0.68   |   0.67 |
| Depth Five  |         |        |

### PostgreSQL

| Depth       | Average | Median |
| ----------- | :-----: | -----: |
| Depth one   |         |        |
| Depth two   |         |        |
| Depth three |         |        |
| Depth Four  |         |        |
| Depth Five  |         |        |

### 