package main

func conversionfunc() {
	var value32 int32 = 123
	var value64 int64 = int64(value32)
	var value16 int16 = int16(value32)

	println(value16)
	println(value32)
	println(value64)

	var test = "sdsdsd"
	println(test[2])
	println(string(test[2]))
}
