ts
==

Simple go Application to get Terminal Size 
Run `github.com/olekukonko/ts` to download, build, and

#### Example

```go
	size , _  := GetSize()
	fmt.Println(size.Col())     // Get Width
	fmt.Println(size.Row())     // Get Height
	fmt.Println(size.PosX())    // Get X Position
	fmt.Println(size.PosY())    // Get Y Position
}
```

