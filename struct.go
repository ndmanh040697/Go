package main

type student struct {
	FirstName string
	LastName  string
	Age       int
	Grade     float32
	ClassName string
}

func main() {
	manh := student{
		FirstName: "Manh",
		LastName:  "Nguyen",
		Age:       26,
		Grade:     9,
		ClassName: "HHM",
	}
	println(getInfo(manh.FirstName, manh.LastName, manh))

}

func getInfo(a, b string, c student) string {
	return c.FirstName + c.LastName

}
