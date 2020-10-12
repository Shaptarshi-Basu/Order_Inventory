# Order_Inventory
APIs to create , fetch and update user and also to create , cancel and fetch orders for those users


#### Go get docker dependencies ####
###### go get github.com/go-sql-driver/mysql ,
###### go get github.com/op/go-logging ,
###### go get github.com/gorilla/mux 

#### RUN test ####
  ###### docker-compose up
  ###### go test -v handler_test.go
  ###### docker-compose down
  
#### RUN API ####
 ###### docker-compose up
 ###### go build cmd/main.go
 ###### ./main
 
 
 
