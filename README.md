ts
==


[![Views in the last 24 hours](https://sourcegraph.com/api/repos/github.com/olekukonko/ts/counters/views-24h.png)]]


Simple go Application to get Terminal Size 
Run `go get github.com/olekukonko/ts` to download, build, and

#### Example

```go
	size , _  := GetSize()
	fmt.Println(size.Col())     // Get Width
	fmt.Println(size.Row())     // Get Height
	fmt.Println(size.PosX())    // Get X Position
	fmt.Println(size.PosY())    // Get Y Position

```

