Super Fast User Agent Detector written in Go. ~10 μs.
# UserAgent Detector:
This package uses "strings.Contains()" to match user agent string against multiple rules
to detect browser name, device type and operating system.

## Benchmark:
```
BenchmarkXMatcher-4       660596              1808 ns/op
BenchmarkXMatcher-4       548271              1847 ns/op
```

## Install
```
go get github.com/superiss/xagent
```

## Example
```go
func main() {
	// browser:safari device:computer os:macintosh
	s := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_3) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.5 Safari/605.1.15"
	s = strings.ToLower(s)
	//
	ua := detector.New(s).Detect()
	//
	fmt.Println("browser:" + ua.Name + " device:" + ua.DeviceType + " os:" + ua.OSType)
}
```

## Important Note:
To match results correctly strings must be converted to lower cases. Use "strings.ToLower"
