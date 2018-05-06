# Usage

```
Usage:
  main [OPTIONS] [URL]

Application Options:
  -n, --name=  output file name with extension. if not
               provided, rangedownloader will guess a file name
               based on URL
  -p, --procs= number of parallel (default: 1)

Help Options:
  -h, --help   Show this help message
 ```
 
 # Example

```
$ go run main.go http://www.golang-book.com/public/pdf/gobook.pdf -p 4
total length: 2893557 bytes
Range: bytes=723389-1446777, 723389 bytes
Range: bytes=0-723388, 723389 bytes
Range: bytes=1446778-2170166, 723389 bytes
Range: bytes=2170167-2893556, 723390 bytes
```