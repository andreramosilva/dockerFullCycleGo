package main;

import("fmt")

func main(){
	fmt.Println("Full Cycle Rocks!!")
}


// import ("fmt"
// 		"log"
// 		"net/http")

// func main(){
// 	http.HandleFunc("/helloworld", func(w http.ResponseWriter, r*http.Request){
// 		fmt.Fprint(w,"Full Cycle Rocks!")
// 	})
// 	fmt.Printf("Server runnning(port=8080) , route: http://localhost:8080/helloworld\n")
// 	if err := http.ListenAndServe(":8080",nil); err != nil {
// 		log.Fatal(err)
// 	}
// }