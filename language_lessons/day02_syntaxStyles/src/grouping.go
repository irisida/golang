package main

func groupedVariables() {
	/*
		tidy multiple declaration this syntactic sugar might visually
		imply sme sort of relationship between variables but there is
		not. It is true that there may be some but a developer would
		need to read and understand the program to see the binding.
		This syntax simply allows for less boilerplate.
	*/
	var (
		speed        int
		doors        int
		transmission string
		gears        int
		model        string
		manufacturer string
	)

}
