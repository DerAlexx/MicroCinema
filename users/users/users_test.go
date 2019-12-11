package users_test

import (
	"context"
	"fmt"
	"testing"

	protoo "github.com/ob-vss-ws19/blatt-4-pwn2own/users/proto"
	"github.com/ob-vss-ws19/blatt-4-pwn2own/users/users"
)

/*
TestAddUser will be a testcase for adding users into the service.
*/
func TestAddUser(t *testing.T) {
	TestName := "Tim"
	service := users.CreateNewUserHandleInstance()
	response := protoo.CreatedUserResponse{User: &protoo.UserMessageResponse{}}
	err := service.CreateUser(context.TODO(), &protoo.CreateUserRequest{Name: TestName}, &response)
	if err == nil {
		switch {
		case response.User.Name != "Tim":
			t.Errorf("Cannot create a user with the name %s", TestName)
		case response.User.Userid < 1:
			t.Fatal("Cannot create a user with a proper ID")
		default:
			t.Log("Creating a User will work.")
		}
	} else {
		fmt.Println(err)
	}
}

/*
TestGetInformationFromMap will be a testcase for adding users into the service and
get all information from him from the map.
*/
func TestGetInformationFromMap(t *testing.T) {
	TestName := "Tim"
	service := users.CreateNewUserHandleInstance()
	responseInsert := protoo.CreatedUserResponse{User: &protoo.UserMessageResponse{}}
	err := service.CreateUser(context.TODO(), &protoo.CreateUserRequest{Name: TestName}, &responseInsert)

	var id int32 = service.GetInformationFromMap("Tim").(int32)
	if err == nil {
		switch {
		case id < 1:
			t.Errorf("Got a wrong id back was smaller then 1! Was: %d", id)
		case responseInsert.User.Userid != id:
			t.Errorf("Cannot find a user with given ID --> Does not match up given with expected ID given %d, wanted %d", responseInsert.User.Userid, id)
		default:
			t.Log("Can get information of a user by his id --> Working fine.")
		}
	} else {
		fmt.Println(err)
	}
}

/*
TestGetInformationFromMap will be a testcase for adding users into the service and
get all information from him from the map.
*/
func TestGetInformationFromMapByName(t *testing.T) {
	TestName := "Tim"
	service := users.CreateNewUserHandleInstance()
	responseInsert := protoo.CreatedUserResponse{User: &protoo.UserMessageResponse{}}
	err := service.CreateUser(context.TODO(), &protoo.CreateUserRequest{Name: TestName}, &responseInsert)

	name := service.GetInformationFromMap(responseInsert.User.Userid).(string)
	if err == nil {
		switch {
		case name != TestName:
			t.Errorf("Got a wrong name back was %s wanted %s", name, TestName)
		case responseInsert.User.Name != TestName:
			t.Errorf("Cannot find a user with given Name --> Non matching: got %s, wanted %s", responseInsert.User.Name, TestName)
			fmt.Println(responseInsert.User.Name)
		default:
			t.Log("Can get information of a user by his name --> Working fine.")
		}
	} else {
		fmt.Println(err)
	}
}

/*
TestAddUserAndFindHim will be a testcase for adding users into the service and try to find him by ID.
*/
func TestAddUserAndFindHim(t *testing.T) {
	TestName := "Tim"
	service := users.CreateNewUserHandleInstance()
	responseInsert := protoo.CreatedUserResponse{User: &protoo.UserMessageResponse{}}
	err := service.CreateUser(context.TODO(), &protoo.CreateUserRequest{Name: TestName}, &responseInsert)

	responseFind := protoo.FindUserResponse{User: &protoo.UserMessageResponse{}}

	vid := responseInsert.User.Userid

	err1 := service.FindUser(context.TODO(), &protoo.FindUserRequest{User: &protoo.UserMessageRequest{Userid: vid}}, &responseFind)

	if err == nil && err1 == nil {
		switch {
		case responseFind.User.Name != "Tim":
			t.Errorf("Cannot find or create a user with the name %s", TestName)
		case responseFind.User.Userid < 1 || responseFind.User.Userid != vid:
			t.Errorf("Cannot find a user with given ID --> Does not match up given with expected ID given %d, wanted %d", vid, responseFind.User.Userid)
		default:
			t.Log("Can create a user and get him by his id.")
		}
	} else {
		fmt.Println(err)
		fmt.Println(err1)
	}
}

/*
TestAddUserAndFindHim will be a testcase for adding users into the service and try to find him by his name.
*/
func TestAddUserAndFindHimByHisName(t *testing.T) {
	TestName := "Tim"
	service := users.CreateNewUserHandleInstance()
	responseInsert := protoo.CreatedUserResponse{User: &protoo.UserMessageResponse{}}
	err := service.CreateUser(context.TODO(), &protoo.CreateUserRequest{Name: TestName}, &responseInsert)

	responseFind := protoo.FindUserResponse{User: &protoo.UserMessageResponse{}}

	vid := responseInsert.User.Userid

	err1 := service.FindUserByName(context.TODO(), &protoo.FindUserByNameRequest{User: &protoo.UserMessageRequestByName{Name: TestName}}, &responseFind)

	if err == nil && err1 == nil {
		switch {
		case responseFind.User.Userid != vid:
			t.Errorf("Cannot find or create a user with the name %s", TestName)
		case responseFind.User.Userid < 1 || responseFind.User.Userid != vid:
			t.Errorf("Cannot find a user with given ID --> Does not match up given with expected ID given %d, wanted %d", vid, responseFind.User.Userid)
		case responseFind.User.Name == "" || responseFind.User.Name != TestName:
			t.Errorf("Cannot find a user with given Name --> Missing match given %s, wanted %s", responseFind.User.Name, TestName)
		default:
			t.Log("Can create a user and get him by his name is fine.")
		}
	} else {
		fmt.Println(err)
		fmt.Println(err1)
	}
}

/*
TestChange will create a user change him an later on call getinformationfrommap in order to see whether or not something
has changed.
*/
func TestAddChangeAndGetInfo(t *testing.T) {
	FirstName := "Tim"
	NewName := "Paulanius"
	service := users.CreateNewUserHandleInstance()
	responseInsert := protoo.CreatedUserResponse{User: &protoo.UserMessageResponse{}}
	err := service.CreateUser(context.TODO(), &protoo.CreateUserRequest{Name: FirstName}, &responseInsert)
	id := responseInsert.User.Userid
	chresponse := protoo.ChangeUserResponse{}
	beforeChange := service.GetInformationFromMap(responseInsert.User.Userid).(string)
	err1 := service.ChangeUser(context.TODO(), &protoo.ChangeUserRequest{Change: &protoo.UserMessageResponse{Userid: id, Name: NewName}}, &chresponse)
	AfterChange := service.GetInformationFromMap(responseInsert.User.Userid).(string)
	if err == nil && err1 == nil {
		switch {
		case beforeChange != FirstName:
			t.Errorf("Beforename is wrong got: %s wanted %s", beforeChange, FirstName)
		case AfterChange != NewName:
			t.Errorf("Aftername is wrong got: %s wanted %s", AfterChange, FirstName)
		case AfterChange == NewName && !chresponse.Change:
			t.Errorf("Name was changed. Found by getinformationfrommap but did not send the correct response.")
		default:
			t.Log("Create a user and change him later on is fine.")
		}
	} else {
		fmt.Println(err)
		fmt.Println(err1)
	}
}

/*
TestChange will create a user and later on delete him.
*/
func TestAddandDeleteAUser(t *testing.T) {
	TestName := "Tim"
	service := users.CreateNewUserHandleInstance()
	responseInsert := protoo.CreatedUserResponse{User: &protoo.UserMessageResponse{}}
	err := service.CreateUser(context.TODO(), &protoo.CreateUserRequest{Name: TestName}, &responseInsert)
	id := responseInsert.User.Userid
	deleteResponse := protoo.DeleteUserResponse{}
	err1 := service.DeleteUser(context.TODO(), &protoo.DeleteUserRequest{User: &protoo.UserMessageRequest{Userid: id}}, &deleteResponse)

	responseFind := protoo.FindUserResponse{User: &protoo.UserMessageResponse{}}

	err2 := service.FindUser(context.TODO(), &protoo.FindUserRequest{User: &protoo.UserMessageRequest{Userid: id}}, &responseFind)
	if err == nil && err1 == nil && err2 == nil {
		if !deleteResponse.IsDeleted && responseFind.User.Userid == -1 {
			t.Errorf("User was deleted. Found by getinformationfrommap but did not send the correct response.")
		} else {
			t.Log("Create a user and change him later on is fine.")
		}
	} else {
		fmt.Println(err)
		fmt.Println(err1)
		fmt.Println(err2)
	}
}

/*
TestChange will create bunch of users and read them later on all from the service.
*/
func TestAddMultipleUsersAndReadAllOfThem(t *testing.T) {
	FirstName := "Tim"
	SecondName := "Paulanius"
	service := users.CreateNewUserHandleInstance()
	responseInsert := protoo.CreatedUserResponse{User: &protoo.UserMessageResponse{}}
	responseInsert2 := protoo.CreatedUserResponse{User: &protoo.UserMessageResponse{}}
	err := service.CreateUser(context.TODO(), &protoo.CreateUserRequest{Name: FirstName}, &responseInsert)
	err1 := service.CreateUser(context.TODO(), &protoo.CreateUserRequest{Name: SecondName}, &responseInsert2)

	all := protoo.AllUsersResponse{}

	err2 := service.ReceiveAndSendAllUsers(context.TODO(), &protoo.AllUsersRequest{}, &all)
	if err == nil && err1 == nil && err2 == nil {
		if len(all.Users) != 2 {
			t.Errorf("The length does not match up. expected %d got %d", 2, len(all.Users))
		}
	} else {
		fmt.Println(err)
		fmt.Println(err1)
		fmt.Println(err2)
	}
}
