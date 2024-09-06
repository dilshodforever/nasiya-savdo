package postgres

// import (
// 	"log"
// 	"testing"

// 	pb "github.com/dilshodforever/nasiya-savdo/genprotos"

// 	"github.com/stretchr/testify/assert"
// )

// func TestRegisterUser(t *testing.T) {
// 	stg, err := NewPostgresStorage()
// 	if err != nil {
// 		log.Fatal("Error while connecting to the db: ", err.Error())
// 	}

// 	user := &pb.UserReq{
// 		UserName:    "Your_Name",
// 		Email:       "e@example.com",
// 		Password:    "helloworld",
// 		PhoneNumber: "Phone Number",
// 	}
// 	result, err := stg.User().Register(user)

// 	assert.NoError(t, err)
// 	assert.NotNil(t, result)
// }

// func TestGetByIdUser(t *testing.T) {
// 	stg, err := NewPostgresStorage()
// 	if err != nil {
// 		log.Fatal("Error while connecting to the db: ", err.Error())
// 	}

// 	id := pb.ById{Id: "80e07ab6-708c-4534-a3b5-9d3e643e78cd"}

// 	user, err := stg.User().GetById(&id)

// 	assert.NoError(t, err)
// 	assert.NotNil(t, user)
// }

// func TestGetAllUser(t *testing.T) {
// 	stg, err := NewPostgresStorage()
// 	if err != nil {
// 		log.Fatal("Error while connecting to the db: ", err.Error())
// 	}

// 	users, err := stg.User().GetAll(&pb.UserFilter{})
// 	assert.NoError(t, err)
// 	assert.NotNil(t, users)
// }

// func TestUpdateUser(t *testing.T) {
// 	stg, err := NewPostgresStorage()
// 	if err != nil {
// 		log.Fatal("Error while connecting to the db: ", err.Error())
// 	}

// 	user := &pb.User{
// 		Id:          "80e07ab6-708c-4534-a3b5-9d3e643e78cd",
// 		UserName:    "New_User_Name",
// 		Email:       "new.e@example.com",
// 		PhoneNumber: "New Phone Number",
// 	}
// 	result, err := stg.User().Update(user)

// 	assert.NoError(t, err)
// 	assert.NotNil(t, result)
// }

// func TestLoginUser(t *testing.T) {
// 	stg, err := NewPostgresStorage()
// 	if err != nil {
// 		log.Fatal("Error while connecting to the db: ", err.Error())
// 	}

// 	user, err := stg.User().Login(&pb.UserLogin{UserName: "New_User_Name", Password: "New_Password"})

// 	assert.NoError(t, err)
// 	assert.NotNil(t, user)
// }

// func TestDeleteUser(t *testing.T) {
// 	stg, err := NewPostgresStorage()
// 	if err != nil {
// 		log.Fatal("Error while connecting to the db: ", err.Error())
// 	}

// 	id := pb.ById{Id: "80e07ab6-708c-4534-a3b5-9d3e643e78cd"}

// 	result, err := stg.User().Delete(&id)

// 	assert.NoError(t, err)
// 	assert.NotNil(t, result)
// }
