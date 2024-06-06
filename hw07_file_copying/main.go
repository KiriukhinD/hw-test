package main

import (
	"flag"
)

var (
	from, to      string
	limit, offset int64
)

func init() {
	flag.StringVar(&from, "from", "D:\\GoProjectWork\\hw-test\\hw07_file_copying\\testdata\\out_offset0_limit1000.txt", "file to read from")
	flag.StringVar(&to, "to", "D:\\GoProjectWork\\hw-test\\hw07_file_copying\\testdata\\testResult.txt", "file to write to")
	flag.Int64Var(&limit, "limit", 0, "limit of bytes to copy")
	flag.Int64Var(&offset, "offset", 0, "offset in input file")
}

func main() {

	flag.Parse()
	err := Copy(from, to, limit, offset)
	if err != nil {
		return
	}

}
