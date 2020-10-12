package handler

import (
	"Order_Inventory/db"
	"encoding/json"
	"github.com/op/go-logging"
	"log"
	"net/http"
	"strings"
)

type Handler struct {
	Log logging.Logger
}
type User struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Phone_Number string `json:"number"`
}
type Order struct {
	Order_Id   string `json:"orderid"`
	Order_Name string `json:"ordername"`
	Quantity   string `json:"quantity"`
	UserID     string `json:"id"`
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := User{}
	dbConn, err := db.CreatConnection()
	if err != nil {
		h.Log.Error(err.Error())
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer dbConn.Close()
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		h.Log.Error(err.Error())
		w.Write([]byte("Cannot Decode request"))
		w.WriteHeader(http.StatusBadRequest)
	}
	insForm, err := dbConn.Prepare("INSERT INTO User(id, name, email, number) VALUES(?,?,?,?)")
	if err != nil {
		h.Log.Error(err.Error())
		w.Write([]byte("Cannot insert data"))
		w.WriteHeader(http.StatusInternalServerError)
	}
	_, err = insForm.Exec(user.Id, user.Name, user.Email, user.Phone_Number)
	if err != nil {
		h.Log.Error(err.Error())
		w.Write([]byte("Cannot insert data"))
		w.WriteHeader(http.StatusInternalServerError)
	}

}
func (h *Handler) FetchUser(w http.ResponseWriter, r *http.Request) {
	dbConn, err := db.CreatConnection()
	if err != nil {
		h.Log.Error(err.Error())
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer dbConn.Close()
	value := strings.Split(r.URL.Path, "/")[3]
	result, err := dbConn.Query("SELECT * FROM User WHERE id=?", value)
	if err != nil {
		h.Log.Error(err.Error())
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
	}
	user := User{}
	for result.Next() {
		var id, name, email, number string
		err = result.Scan(&id, &name, &email, &number)
		if err != nil {
			h.Log.Error(err.Error())
			w.Write([]byte("Cannot read user data"))
			w.WriteHeader(http.StatusInternalServerError)
		}
		user.Id = id
		user.Name = name
		user.Email = email
		user.Phone_Number = number
	}
	jData, err := json.Marshal(user)
	if err != nil {
		h.Log.Error(err.Error())
		w.Write([]byte("Cannot unmarshal data fetched for user"))
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)

}
func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		h.Log.Error(err.Error())
		w.Write([]byte("Cannot Decode request"))
		w.WriteHeader(http.StatusBadRequest)
	}
	var queryString = "UPDATE User SET "
	if user.Name != "" {
		queryString += "name=?"
	}
	if user.Email != "" {
		if user.Name != "" {
			queryString += ", "
		}
		queryString += "email=?"
	}
	if user.Phone_Number != "" {
		if user.Name != "" || user.Email != "" {
			queryString += ", "
		}
		queryString += "number=?"
	}
	queryString += " WHERE id=?"
	dbConn, err := db.CreatConnection()
	if err != nil {
		h.Log.Error(err.Error())
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer dbConn.Close()
	value := strings.Split(r.URL.Path, "/")[3]
	insForm, err := dbConn.Prepare(queryString)
	if err != nil {
		h.Log.Error(err.Error())
		w.Write([]byte("Cannot update user data"))
		w.WriteHeader(http.StatusInternalServerError)
	}
	if user.Name != "" && user.Email != "" && user.Phone_Number != "" {
		_, err = insForm.Exec(user.Name, user.Email, user.Phone_Number, value)
	} else if user.Name != "" && user.Email != "" {
		_, err = insForm.Exec(user.Name, user.Email, value)
	} else if user.Name != "" && user.Phone_Number != "" {
		_, err = insForm.Exec(user.Name, user.Phone_Number, value)
	} else if user.Email != "" && user.Phone_Number != "" {
		_, err = insForm.Exec(user.Email, user.Phone_Number, value)
	} else if user.Name != "" {
		_, err = insForm.Exec(user.Name, value)
	} else if user.Email != "" {
		_, err = insForm.Exec(user.Email, value)
	} else {
		_, err = insForm.Exec(user.Phone_Number, value)
	}
	if err != nil {
		h.Log.Error(err.Error())
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
func (h *Handler) FetchAllUsers(w http.ResponseWriter, r *http.Request) {
	user := User{}
	dbConn, err := db.CreatConnection()
	if err != nil {
		h.Log.Error(err.Error())
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer dbConn.Close()
	result, err := dbConn.Query("SELECT * FROM User ")
	if err != nil {
		h.Log.Error(err.Error())
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
	}
	for result.Next() {
		var id, name, email, number string
		err = result.Scan(&id, &name, &email, &number)
		if err != nil {
			h.Log.Error(err.Error())
			log.Fatal(err.Error())
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
		}
		user.Id = id
		user.Name = name
		user.Email = email
		user.Phone_Number = number
		jData, err := json.Marshal(user)
		if err != nil {
			h.Log.Error(err.Error())
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Write(jData)
	}
}

func (h *Handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	order := Order{}
	dbConn, err := db.CreatConnection()
	if err != nil {
		h.Log.Error(err.Error())
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer dbConn.Close()
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		h.Log.Error(err.Error())
		w.Write([]byte("Cannot Decode request"))
		w.WriteHeader(http.StatusBadRequest)
	}
	insForm, err := dbConn.Prepare("INSERT INTO Order_inv(orderid, ordername, quantity, id) VALUES(?,?,?,?)")
	if err != nil {
		h.Log.Error(err.Error())
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
	}
	_, err = insForm.Exec(order.Order_Id, order.Order_Name, order.Quantity, order.UserID)
	if err != nil {
		h.Log.Error(err.Error())
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
	}

}

func (h *Handler) CancelOrder(w http.ResponseWriter, r *http.Request) {
	value := strings.Split(r.URL.Path, "/")[3]
	dbConn, err := db.CreatConnection()
	if err != nil {
		h.Log.Error(err.Error())
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
	}
	_, err = dbConn.Query("Delete FROM Order_inv WHERE orderid=?", value)
	if err != nil {
		h.Log.Error(err.Error())
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
	}
}
func (h *Handler) FetchAllOrders(w http.ResponseWriter, r *http.Request) {
	value := strings.Split(r.URL.Path, "/")[3]
	dbConn, err := db.CreatConnection()
	if err != nil {
		h.Log.Error(err.Error())
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer dbConn.Close()
	result, err := dbConn.Query("SELECT * FROM Order_inv WHERE id=?", value)
	if err != nil {
		h.Log.Error(err.Error())
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
	}
	for result.Next() {
		order := Order{}
		var orderid, ordername, quantity, id string
		err = result.Scan(&orderid, &ordername, &quantity, &id)
		if err != nil {
			panic(err.Error())
		}
		order.Order_Id = orderid
		order.Order_Name = ordername
		order.Quantity = quantity
		order.UserID = id
		jData, err := json.Marshal(order)
		if err != nil {
			h.Log.Error(err.Error())
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Write(jData)
	}
}
