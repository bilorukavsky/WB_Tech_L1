/*
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
Приведите корректный пример реализации.
*/
var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	copy(justString[:100], v[:100])
}

func main() {
	someFunc()
}
