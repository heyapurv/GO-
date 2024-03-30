package main 
import "fmt"
const LoginToken string = "sjflksjflks"
// var with first character as capital will be a global variable and it can be used anywhere 
func main()  {
	// fmt.Println("Variables")
	var username string = "Apurv"
    fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
    fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255 //uint8 value from 1 to 255 
    fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.23344325345
	// u can use float 64 for more detaild value
    fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	var anotherVar int
	fmt.Println(anotherVar)
	fmt.Printf("Variable is of type: %T \n",anotherVar)

	// implicit type to declare a var
	var website = "xyz.online"
	fmt.Println(website)
	// website = 3  we cannod do it 

	//    no var style 
	// you can use this operator inside a method 
	numberOfUser := 3000000
	fmt.Println(numberOfUser)
    

	fmt.Println(LoginToken)
	fmt.Printf("Type of this var is : %T",LoginToken)

	
}