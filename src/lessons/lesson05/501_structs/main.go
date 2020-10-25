package main

import "fmt"

type employee struct {
	first  string
	middle string
	last   string
	email  string
}

type management struct {
	employee
	teamSize int
}

func main() {
	e1 := employee{
		first: "johnny",
		last:  "one",
		email: "j1@emp.co",
	}

	e2 := employee{
		first: "jimmy",
		last:  "two",
		email: "j2@emp.co",
	}

	mgr1 := management{
		employee: employee{
			first: "joe",
			last:  "K",
			email: "bojo@jojo.yoyo",
		},
		teamSize: 5,
	}

	// remember that employee and management are different types
	// management is constructed from composition, it includes an
	// employee struct as part of its makeup.

	emps := []employee{}
	mgrs := []management{}

	emps = append(emps, e1, e2)
	mgrs = append(mgrs, mgr1)

	fmt.Println(emps)
	fmt.Println(mgrs)

}
