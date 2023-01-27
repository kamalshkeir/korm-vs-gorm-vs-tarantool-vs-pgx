# Korm vs Gorm vs Tarantool vs Pgx 
# Postgres database

### These benchmarks could not be done without the help of [kokizzu](https://github.com/kokizzu) who started the project [gorm-vs-korm](https://github.com/kokizzu/gorm-vs-korm)

```
goos: linux
goarch: amd64
```

## Usage

```shell
./clean-start.sh
go test -bench=Korm -benchmem .
go test -bench=Gorm -benchmem .
go test -bench=Taran -benchmem .
go test -bench=Pgx -benchmem .
```


## 2023-01-27  Result 100K rows, GetAll limited to 1000 rows unordered, concurrency: 32
```
## korm 1.4.8
## pgx 5.2.0
## go-tarantool 1.10.0

## GetAll
BenchmarkGetAllS_Taran_Raw-4                 525           2112783 ns/op          936663 B/op       5736 allocs/op
BenchmarkGetAllS_Taran_ORM-4                1579            721029 ns/op          233939 B/op       4714 allocs/op
BenchmarkGetAllS_Postgres_Korm-4         1570945               732.8 ns/op           240 B/op          3 allocs/op <-- fastest
BenchmarkGetAllS_Postgres_Gorm-4             729           2052911 ns/op          221538 B/op       7809 allocs/op
BenchmarkGetAllS_Postgres_Pgx-4             1573            779752 ns/op           59560 B/op       2968 allocs/op

## GetRow second iteration, first one doesn't make sense, because it only save to cache and never fetched
BenchmarkGetRowS_Taran_ORM2-4              37310             34005 ns/op            1115 B/op         27 allocs/op
BenchmarkGetRowS_Postgres_Korm2-4         850898             28326 ns/op             491 B/op         13 allocs/op <-- fastest
BenchmarkGetRowS_Postgres_Gorm-4            1321            766825 ns/op           16393 B/op        150 allocs/op
BenchmarkGetRowS_Postgres_Pgx-4             3710            277754 ns/op             785 B/op         18 allocs/op

## INSERT
BenchmarkInsertS_Taran_ORM-4              100000            567109 ns/op            1078 B/op         26 allocs/op
BenchmarkInsertS_Postgres_Korm-4          100000            402000 ns/op            3057 B/op         58 allocs/op  <-- fastest
BenchmarkInsertS_Postgres_Gorm-4          100000            978457 ns/op           17284 B/op        172 allocs/op
BenchmarkInsert_Postgres_Pgx-4            100000            913335 ns/op             483 B/op         15 allocs/op

## Update
BenchmarkUpdate_Taran_ORM-32              Failed           
BenchmarkUpdate_Postgres_Korm-4           200000            251661 ns/op            1474 B/op         31 allocs/op <-- fastest
BenchmarkUpdate_Postgres_Gorm-32          Failed            
BenchmarkUpdate_Postgres_Pgx-4            200000            949615 ns/op             447 B/op         15 allocs/op 
```

## Conclusion

Korm fastest for everything, no matter the size of the database and capacity of RAM, because cache is limited and do not need any intervention, and cache is flushed every 10 minutes or programmaticaly using `korm.FlushCache()`.
Also worth to note that korm.Query and korm.QueryS are both 3 times faster than GetAll, in another word, without using the builder you can get more performance like so : `korm.QueryS[User]("dbName","select * from users where id = ?",1)` return `[]User,error`
