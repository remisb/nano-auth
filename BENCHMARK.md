# Benchmark results

## Package main file test_web.go

/Users/remis/Library/Caches/JetBrains/GoLand2023.3/tmp/GoLand/___BenchmarkStructToJson_in_github_com_remisb_nano_auth.test -test.v -test.paniconexit0 -test.bench ^\QBenchmarkStructToJson\E$ -test.run ^$
goos: darwin
goarch: arm64
pkg: github.com/remisb/nano-auth
BenchmarkStructToJson
BenchmarkStructToJson-16    	 9698664	       123.2 ns/op
PASS

pkg: github.com/remisb/nano-auth
BenchmarkStructToJson
BenchmarkStructToJson-16    	 9509714	       125.5 ns/op
PASS

pkg: github.com/remisb/nano-auth
BenchmarkStructToJson
BenchmarkStructToJson-16    	 9488487	       126.2 ns/op
PASS

----

/Users/remis/Library/Caches/JetBrains/GoLand2023.3/tmp/GoLand/___BenchmarkMapToJson_in_github_com_remisb_nano_auth.test -test.v -test.paniconexit0 -test.bench ^\QBenchmarkMapToJson\E$ -test.run ^$
goos: darwin
goarch: arm64
pkg: github.com/remisb/nano-auth
BenchmarkMapToJson
BenchmarkMapToJson-16    	 3061627	       395.6 ns/op
PASS

pkg: github.com/remisb/nano-auth
BenchmarkMapToJson
BenchmarkMapToJson-16    	 3036990	       391.8 ns/op
PASS

pkg: github.com/remisb/nano-auth
BenchmarkMapToJson
BenchmarkMapToJson-16    	 3013972	       394.5 ns/op
PASS

| BenchmarkMapToJson | BenchmarkStructToJson |
|--------------------|-----------------------|
| 3013972            | 9698664               |
| 394.5 ns/op        | 123.2 ns/op           |       