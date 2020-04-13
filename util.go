package main

type Int int

func (x  Int)Less(than Item) bool  {
	return  x < than.(Int)
}

type UInt32 uint32

func (x UInt32)Less (than Item)bool  {
	return x < than.(UInt32)
}

type String string

func (x String)Less (than Item) bool  {
	return  x < than.(String)
}
