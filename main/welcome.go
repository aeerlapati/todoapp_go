package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/boltdb/bolt"
	"github.com/gorilla/mux"
)

type todoJSON struct {
	TaskName         string `json:"taskname"`
	TaskDescription  string `json:"taskdescription"`
	TaskCompleteFlag bool   `json:"taskcompleteflag"`
	TaskCompleteDate string `json:"taskcompletedate"`
}

var tdJSON *todoJSON = &todoJSON{}
var db1 *bolt.DB

//Canary is main function
func Canary(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Inside Canary\n")

	fmt.Fprint(w, "tweet")
}

//Canary is main function
func addTask(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Inside addTask\n")

	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
		}

		//printing the body
		fmt.Printf("Input Value ---- %s\n", body)
		var data todoJSON = todoJSON{}

		if err := json.Unmarshal(body, &data); err != nil {

			fmt.Printf(" Error :%s", err) // 5

			http.Error(w, "Failed to UnMarshal Data",
				http.StatusInternalServerError)
		} else {
			fmt.Printf("Input Value ---- %s\n", data.TaskName)        // 5
			fmt.Printf("Input Value ---- %s\n", data.TaskDescription) // 5

			if data.TaskName != "" && len(data.TaskName) != 0 {

				startBolt()
				//pushing the data into the DB
				err := settasks(db1, data)
				if err != nil {
					http.Error(w, "Insertion of data failed",
						http.StatusInternalServerError)

					fmt.Printf(" Error :%s\n", err) // 5
				}
				//
				//END of insertion
				//

				fmt.Fprint(w, "POST done, added task")
			} else {
				fmt.Fprint(w, "Invalid JSON")
			}
		}
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

//settasks is used for inserting Data
func settasks(db *bolt.DB, result todoJSON) error {
	confBytes, err := json.Marshal(result)
	if err != nil {
		return fmt.Errorf("could not marshal config json: %v", err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		err = tx.Bucket([]byte("DB")).Put([]byte("CONFIG"), confBytes)
		if err != nil {
			return fmt.Errorf("could not set config: %v", err)
		}
		return nil
	})
	fmt.Println("Set Config")
	return err
}

func startBolt() {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	db1 = db
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("DB"))
		if err != nil {
			return fmt.Errorf("could not create root bucket: %v", err)
		}
		return nil
	})
}

func main() {
	fmt.Printf("Starting server at port 9199\n")

	//setting up conn to boltdb
	//startBolt()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/canary", func(w http.ResponseWriter, r *http.Request) {
		Canary(w, r)
	})
	router.HandleFunc("/addTask", func(w http.ResponseWriter, r *http.Request) {
		addTask(w, r)
	})

	log.Fatal(http.ListenAndServe(":9199", router))

}
