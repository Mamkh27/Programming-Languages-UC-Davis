package smash

import (
	"io"
	"bytes"
	"strings"
)

type word string

func Smash(r io.Reader, smasher func(word) uint32) map[uint32]uint {
	m := make(map[uint32]uint)

	buf:= new(bytes.Buffer)
	buf.ReadFrom(r)
	s:= buf.String()

	fi := strings.Fields(s)

	c := make(chan map[uint32]uint,len(fi))
	go Smashme(m,fi,smasher,c)
	for x := range c{
		m = x
	}

	return m
}

func Smashme(m map[uint32]uint,fi []string,smasher func(word) uint32,c chan map[uint32]uint){
	for i:=0; i<len(fi);i++{
		m[smasher(word(fi[i]))] += 1
		c <- m
	}
	close(c)
}