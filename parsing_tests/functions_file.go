package parsingtests

var global_var somethingA = somethingA{a: 40}

type somethingA struct {
	a int
}

//no args
func yolo(){
}

/* has args and is pure because it takes 
copies only as arguments and doesn't use access
any global variables to mutate them
*/
func yolo2(name string) string{
	return name+"cool"
}

/* has args and manipulates the value, hence 
	a pointer is passed in, so it's impure
*/
func yolo3(s *somethingA) bool{
	return false
}

/* has no args but manipulates a global value
	internally so it's impure, also returns a bool */
func yolo4() bool {
	global_var.a = 55
	return true
}
