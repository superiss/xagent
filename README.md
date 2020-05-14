Super Fast User Agent Detector written in Go. ~10 μs.

# UserAgent Detector:
This package uses "strings.Contains()" to match user agent string against multiple rules
to detect browser name, device type and operating system.

## Install
```
go get github.com/superiss/xagent
```

## UserAgent Struct
```go
type UserAgent struct {
	Name       string
	DeviceType string
	OSType     string
}
```
## Example
```go
func main() {
	s := "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:15.0) Gecko/20100101 Firefox/15.0.1"
	ua := xagent.Parser(s)
	//
	fmt.Println(ua)
	// // //
	// output: &{firefox computer linux}
}
```

## Important Note:
To match results correctly strings must be converted to lower cases. Use "strings.ToLower"
