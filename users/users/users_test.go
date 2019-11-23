package users_test

import (
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
	service.CreateUser(nil, &protoo.CreateUserRequest{Name: TestName}, &response)

	if response.User.Name != "Tim" {
		t.Errorf("Cannot create a user with the name %s", TestName)
	} else if response.User.Userid < 1 {
		t.Fatal("Cannot create a user with a proper ID")
	} else {
		t.Log("Creating a User will work.")
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
	service.CreateUser(nil, &protoo.CreateUserRequest{Name: TestName}, &responseInsert)

	var id int32 = service.GetInformationFromMap("Tim").(int32)

	if id < 1 {
		t.Errorf("Got a wrong id back was smaller then 1! Was: %d", id)
	} else if responseInsert.User.Userid != id {
		t.Errorf("Cannot find a user with given ID --> Does not match up given with expected ID given %d, wanted %d", responseInsert.User.Userid, id)
	} else {
		t.Log("Can get information of a user by his id --> Working fine.")
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
	service.CreateUser(nil, &protoo.CreateUserRequest{Name: TestName}, &responseInsert)

	name := service.GetInformationFromMap(responseInsert.User.Userid).(string)

	if name != TestName {
		t.Errorf("Got a wrong name back was %s wanted %s", name, TestName)
	} else if responseInsert.User.Name != TestName {
		t.Errorf("Cannot find a user with given Name --> Non matching: got %s, wanted %s", responseInsert.User.Name, TestName)
		fmt.Println(responseInsert.User.Name)
	} else {
		t.Log("Can get information of a user by his name --> Working fine.")
	}

}

/*
TestAddUserAndFindHim will be a testcase for adding users into the service and try to find him by ID.
*/
func TestAddUserAndFindHim(t *testing.T) {
	TestName := "Tim"
	service := users.CreateNewUserHandleInstance()
	responseInsert := protoo.CreatedUserResponse{User: &protoo.UserMessageResponse{}}
	service.CreateUser(nil, &protoo.CreateUserRequest{Name: TestName}, &responseInsert)

	responseFind := protoo.FindUserResponse{User: &protoo.UserMessageResponse{}}

	vid := responseInsert.User.Userid

	service.FindUser(nil, &protoo.FindUserRequest{User: &protoo.UserMessageRequest{Userid: vid}}, &responseFind)

	if responseFind.User.Name != "Tim" {
		t.Errorf("Cannot find or create a user with the name %s", TestName)
	} else if responseFind.User.Userid < 1 || responseFind.User.Userid != vid {
		t.Errorf("Cannot find a user with given ID --> Does not match up given with expected ID given %d, wanted %d", vid, responseFind.User.Userid)
	} else {
		t.Log("Can create a user and get him by his id.")
	}
}

/*
TestAddUserAndFindHim will be a testcase for adding users into the service and try to find him by his name.
*/
func TestAddUserAndFindHimByHisName(t *testing.T) {
	TestName := "Tim"
	service := users.CreateNewUserHandleInstance()
	responseInsert := protoo.CreatedUserResponse{User: &protoo.UserMessageResponse{}}
	service.CreateUser(nil, &protoo.CreateUserRequest{Name: TestName}, &responseInsert)

	responseFind := protoo.FindUserResponse{User: &protoo.UserMessageResponse{}}

	vid := responseInsert.User.Userid

	service.FindUserByName(nil, &protoo.FindUserByNameRequest{User: &protoo.UserMessageRequestByName{Name: TestName}}, &responseFind)

	if responseFind.User.Userid != vid {
		t.Errorf("Cannot find or create a user with the name %s", TestName)
	} else if responseFind.User.Userid < 1 || responseFind.User.Userid != vid {
		t.Errorf("Cannot find a user with given ID --> Does not match up given with expected ID given %d, wanted %d", vid, responseFind.User.Userid)
	} else if responseFind.User.Name == "" || responseFind.User.Name != TestName {
		t.Errorf("Cannot find a user with given Name --> Missing match given %s, wanted %s", responseFind.User.Name, TestName)
	} else {
		t.Log("Can create a user and get him by his name is fine.")
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
	service.CreateUser(nil, &protoo.CreateUserRequest{Name: FirstName}, &responseInsert)
	id := responseInsert.User.Userid
	chresponse := protoo.ChangeUserResponse{}
	beforeChange := service.GetInformationFromMap(responseInsert.User.Userid).(string)
	service.ChangeUser(nil, &protoo.ChangeUserRequest{Change: &protoo.UserMessageResponse{Userid: id, Name: NewName}}, &chresponse)
	AfterChange := service.GetInformationFromMap(responseInsert.User.Userid).(string)

	if beforeChange != FirstName {
		t.Errorf("Beforename is wrong got: %s wanted %s", beforeChange, FirstName)
	} else if AfterChange != NewName {
		t.Errorf("Aftername is wrong got: %s wanted %s", AfterChange, FirstName)
	} else if AfterChange == NewName && !chresponse.Change {
		t.Errorf("Name was changed. Found by getinformationfrommap but did not send the correct response.")
	} else {
		t.Log("Create a user and change him later on is fine.")
	}
}

/*
TestChange will create a user and later on delete him.
*/
func TestAddandDeleteAUser(t *testing.T) {
	TestName := "Tim"
	service := users.CreateNewUserHandleInstance()
	responseInsert := protoo.CreatedUserResponse{User: &protoo.UserMessageResponse{}}
	service.CreateUser(nil, &protoo.CreateUserRequest{Name: TestName}, &responseInsert)
	id := responseInsert.User.Userid
	deleteResponse := protoo.DeleteUserResponse{}
	beforeChange := service.GetInformationFromMap(responseInsert.User.Userid).(string)
	service.DeleteUser(nil, &protoo.DeleteUserRequest{User: &protoo.UserMessageRequest{Userid: id}}, &deleteResponse)
	AfterChange := service.GetInformationFromMap(responseInsert.User.Userid).(string)

	if beforeChange != TestName {
		t.Errorf("Cant even set a user -> got: %s wanted %s", beforeChange, TestName)
	} else if AfterChange == "" && !deleteResponse.IsDeleted {
		t.Errorf("User was deleted. Found by getinformationfrommap but did not send the correct response.")
	} else {
		t.Log("Create a user and change him later on is fine.")
	}
}

/*
TestChange will create bunch of users and read them later on all from the service.
*/
func TestAddMultipleUsersAndReadAllOfThem(t *testing.T) {
	FirstName := "Tim"
	NewName := "Paulanius"
	service := users.CreateNewUserHandleInstance()
	responseInsert := protoo.CreatedUserResponse{User: &protoo.UserMessageResponse{}}
	service.CreateUser(nil, &protoo.CreateUserRequest{Name: FirstName}, &responseInsert)
	id := responseInsert.User.Userid
	chresponse := protoo.ChangeUserResponse{}
	beforeChange := service.GetInformationFromMap(responseInsert.User.Userid).(string)
	service.ChangeUser(nil, &protoo.ChangeUserRequest{Change: &protoo.UserMessageResponse{Userid: id, Name: NewName}}, &chresponse)
	AfterChange := service.GetInformationFromMap(responseInsert.User.Userid).(string)

	if beforeChange != FirstName {
		t.Errorf("Beforename is wrong got: %s wanted %s", beforeChange, FirstName)
	} else if AfterChange != NewName {
		t.Errorf("Aftername is wrong got: %s wanted %s", beforeChange, FirstName)
	} else if AfterChange == NewName && !chresponse.Change {
		t.Errorf("Name was changed. Found by getinformationfrommap but did not send the correct response.")
	} else {
		t.Log("Create a user and change him later on is fine.")
	}
}
