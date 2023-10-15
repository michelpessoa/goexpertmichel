package main

type Math struct {
	A int
	B int
}

func (m Math) Add() int {
	return m.A + m.B
}
