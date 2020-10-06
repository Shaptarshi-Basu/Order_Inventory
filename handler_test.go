package handler_test

import (
	"TechNo_Tree_API/db"
	"TechNo_Tree_API/handler"
	"TechNo_Tree_API/routes"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

type testCase struct {
	name string
	path string
	setupPath string
	expectedVal string
}
func TestEndpoints (t *testing.T) {
	defer testCleanup()
	testCases := []testCase{
		{
			"create_user",
			"/create/customer",
			"",
			"",
		},
		{
			"fetch_user",
			"/fetch/customer/basu_11",
			"",
			`{"id":"Shapta","name":"basu_11","email":"shapbasu@gmail.com","number":"0406820430"}`,
		},
		{
			"fetch_all_users",
			"/fetch/customers",
			"/create/customer",
			`{"id":"Shapta","name":"basu_11","email":"shapbasu@gmail.com","number":"0406820430"}{"id":"Basu","name":"basu_12","email":"shapbasu@gmail.com","number":"0406820430"}`,
		},
		{
			"update_user",
			"/update/customer/basu_11",
			"",
			`{"id":"Shapta","name":"basu_11","email":"shapbasu@gmail.com","number":"0406820430"}`,
		},
		{
			"create_order",
			"/create/order",
			"",
			"",
		},
	}
	go startServer()
	for _, testcase := range testCases {
		t.Run(testcase.name, func(t *testing.T) {
			if testcase.name == "create_user" {
				file, _ := os.Open("resources/user1.json")
				reqPath := "http://localhost:7070" + testcase.path
				resp, _ := http.Post(reqPath, "application/json", file)
				if resp.StatusCode != http.StatusOK {
					t.Errorf("User cant be created")
				}

			} else if testcase.name == "fetch_user" {
				reqPath := "http://localhost:7070" + testcase.path
				resp, _ := http.Get(reqPath)
				if resp.StatusCode == http.StatusOK {
					body, _ := ioutil.ReadAll(resp.Body)
					bodStr := string(body)
					if bodStr != testcase.expectedVal {
						t.Errorf("Wrong data fetched")
					}
				} else {
					t.Errorf("User data could not be fetched")
				}
			} else if testcase.name == "fetch_all_users" {
				file, _ := os.Open("resources/user2.json")
				reqPathPost := "http://localhost:7070" + testcase.setupPath
				inserResp, _ := http.Post(reqPathPost, "application/json", file)
				if inserResp.StatusCode != http.StatusOK {
					t.Errorf("User cant be created")
				}
				reqPath := "http://localhost:7070" + testcase.path
				resp, _ := http.Get(reqPath)
				if resp.StatusCode == http.StatusOK {
					body, _ := ioutil.ReadAll(resp.Body)
					bodStr := string(body)
					if bodStr != testcase.expectedVal {
						t.Errorf("Wrong data fetched")
					}
				} else {
					t.Errorf("User data could not be fetched")
				}
			}else if testcase.name == "update_user" {
				file, _ := os.Open("resources/userUpdate1.json")
				reqPath := "http://localhost:7070" + testcase.path
				resp, _ := http.Post(reqPath, "application/json", file)
				if resp.StatusCode != http.StatusOK {
					t.Errorf("User cant be updated")
				}

			}else if testcase.name == "create_order" {
				file, _ := os.Open("resources/order.json")
				reqPath := "http://localhost:7070" + testcase.path
				resp, _ := http.Post(reqPath, "application/json", file)
				if resp.StatusCode != http.StatusOK {
					t.Errorf("Order cant be created")
				}
			}

		})
	}

}
func testCleanup(){
	conn :=db.CreatConnection()
	conn.Query("TRUNCATE TABLE User")
	conn.Query("TRUNCATE TABLE Order_inv")
}
func startServer() {
	r := mux.NewRouter()
	s := &routes.Server{
		Handler: handler.Handler{},
	}
	routes.SetRoutes(r, s)
}