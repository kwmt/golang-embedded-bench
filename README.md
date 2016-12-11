ginでEngine構造体に埋め込まれたRouterGroupを`engine.RouterGroup`のように使っていて、可読性のためであることはすぐわかるのですが、速度も気にしているのかな？とちょっと気になったのでベンチを取ってみた

https://github.com/gin-gonic/gin/blob/master/gin.go#L166

```
% go test -bench .
BenchmarkEmbeded-4  	2000000000	         0.82 ns/op
BenchmarkEmbededNo-4	2000000000	         0.77 ns/op
ok  	3.364s
% go test -bench .
BenchmarkEmbeded-4  	2000000000	         0.81 ns/op
BenchmarkEmbededNo-4	2000000000	         0.77 ns/op
ok  	3.352s
% go test -bench .
BenchmarkEmbeded-4  	2000000000	         0.83 ns/op
BenchmarkEmbededNo-4	2000000000	         0.80 ns/op
ok  		3.460s
% go test -bench .
BenchmarkEmbeded-4  	2000000000	         0.78 ns/op
BenchmarkEmbededNo-4	2000000000	         0.80 ns/op
ok  		3.351s
% go test -bench .
BenchmarkEmbeded-4  	2000000000	         0.82 ns/op
BenchmarkEmbededNo-4	2000000000	         0.81 ns/op
ok  		3.470s
% go test -bench .   
BenchmarkEmbeded-4  	2000000000	         0.80 ns/op
BenchmarkEmbededNo-4	2000000000	         0.78 ns/op
ok  		3.352s
```
