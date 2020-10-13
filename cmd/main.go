package main


import (
	"encoding/json"
	"fmt"
	"net/http"

	// docs is generated by Swag CLI, you have to import it.
	"github.com/punithsr27/caching-system"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {

	//handleRequests()
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/punith", homePage).Methods("POST")
	// log.Fatal(http.ListenAndServe(":8000", nil))
	myRouter.PathPrefix("pkg/serviceImpl/docs").Handler(httpSwagger.WrapHandler)
	http.ListenAndServe(":8000", myRouter)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	var cache caching-system.pkg.Cache
	fmt.Fprintf(w, "Welcome to the HomePage!")

	_ = json.NewDecoder(r.Body).Decode(&cache)
	emps = append(emps, cache)
	json.NewEncoder(w).Encode(cache)

}

func GetEmps(w http.ResponseWriter, r *http.Request) {
}

func GetEmp(w http.ResponseWriter, r *http.Request) {
}

func CreateEmp(w http.ResponseWriter, r *http.Request) {
}

func DeleteEmp(w http.ResponseWriter, r *http.Request) {
}

//	http.Handle("/", router)
// srv := &http.Server{
// 	Handler: router,
// 	Addr:    "127.0.0.1:8000",
// 	// Good practice: enforce timeouts for servers you create!
// 	WriteTimeout: 15 * time.Second,
// 	ReadTimeout:  15 * time.Second,
// }
//	http.ListenAndServe(":8080", router)
//	log.Fatal(http.ListenAndServe(":8080", router))
//}

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Welcome to the HomePage!")
// 	fmt.Println("Endpoint Hit: homePage")
// }
//
// func handleRequests() {
// 	http.HandleFunc("/", homePage)
// 	log.Fatal(http.ListenAndServe(":10000", nil))
// }
//
// func handleRequestss() {
// 	http.HandleFunc("/punith", homePage)
// 	log.Fatal(http.ListenAndServe(":10000", nil))
// }
>>>>>>> e9f0958669060fe8e3f1e1fe57c4c277cc707b4a
