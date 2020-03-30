package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {
	p("Contains:   ", s.Contains("test", "es"))
	p("Count:      ", s.Count("test", "t"))
	p("hasprefix:  ", s.HasPrefix("test", "te"))
	p("hassuffix:  ", s.HasSuffix("test", "st"))
	p("index:      ", s.Index("test", "e"))
	p("join:       ", s.Join([]string{"a", "b"}, "-"))
	p("repeat:     ", s.Repeat("a", 5))
	p("replace:    ", s.Replace("foo", "o", "0", -1))
	p("replace:    ", s.Replace("foo", "o", "0", 1))
	p("split:      ", s.Split("a-b-c-d-e", "-"))
	p("tolower:    ", s.ToLower("TEST"))
	p("toupper:    ", s.ToUpper("test"))
	p()

	p("len:", len("hello"))
	p("char:", "hello"[1])
}
