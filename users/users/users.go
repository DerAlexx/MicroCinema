package users

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	resproto "github.com/ob-vss-ws19/blatt-4-pwn2own/reservation/proto"
	proto "github.com/ob-vss-ws19/blatt-4-pwn2own/users/proto"
)

const (
	maxuserid int32 = 987654321
)

/*
containsID will check whether generated ID is already set or not.
@id will be a int32 id to check for.
*/
func (u *UserHandlerService) containsID(id int32) bool {
	_, inMap := (*u.getUserMap())[id]
	return inMap
}

/*
Function to produce a random userID.
@param length will be the length of the number.
@param seed will be a seed in order to produce "really" random numbers.
*/
func (u *UserHandlerService) getRandomUserID(length int32) int32 {
	rand.Seed(time.Now().UnixNano())
	for {
		potantialID := rand.Int31n(length)
		if !u.containsID(potantialID) {
			return potantialID
		}
	}
}

/*
getUserMap will return a pointer to the current user map in order to work in that.
*/
func (u *UserHandlerService) getUserMap() *map[int32]*userS {
	return &u.user
}

/*
setUserMap will set the map of a userhandlerservice instance.
@param users will be the map to set.
func (u *UserHandlerService) setUserMap(users *map[int32]*userS) {
	u.user = *users
}
*/

type userS struct {
	name string
}

/*
getName will return the name of a given userS-instance.
*/
func (us *userS) getName() string {
	return us.name
}

/*
Dependencies are all Dependencies a user-service has.
*/
type Dependencies struct {
	ResService func() resproto.ReservationService
}

/*
UserHandlerService will be the representation of our service.
*/
type UserHandlerService struct {
	user         map[int32]*userS
	dependencies *Dependencies
	mutex        *sync.Mutex
}

/*
CreateNewUserHandleInstance will create a new service for the user management.
*/
func CreateNewUserHandleInstance() *UserHandlerService {
	return &UserHandlerService{
		user:  make(map[int32]*userS),
		mutex: &sync.Mutex{},
	}
}

/*
AddDependency will add a dependency into the service.
*/
func (u *UserHandlerService) AddDependency(dep *Dependencies) {
	u.dependencies = dep
}

/*
appendANewUser will add a new user in the datastructure.
*/
func (u *UserHandlerService) appendANewUser(id int32, user *userS) bool {
	if id != 0 && user != nil {
		(*u.getUserMap())[id] = user
		return true
	}
	return false
}

/*
CreateUser will create a user and append it into the structures. After that this method will
return a response.
*/
func (u *UserHandlerService) CreateUser(context context.Context, request *proto.CreateUserRequest, response *proto.CreatedUserResponse) error {
	if request.GetName() != "" {
		uid := u.getRandomUserID(maxuserid)
		u.mutex.Lock()
		if u.appendANewUser(uid, &userS{name: request.GetName()}) {
			response.User = &proto.UserMessageResponse{Name: request.GetName(), Userid: uid}
			defer u.mutex.Unlock()
			return nil
		}
	}
	return fmt.Errorf("cannot create user with empty name")
}

/*
HasOpenReservations will check for open reservations of a user in the Reservationsservice.
*/
func (u *UserHandlerService) HasOpenReservations(context context.Context, uid int32) bool {
	serv := u.dependencies.ResService()
	in := &resproto.HasReservationsRequest{Res: &resproto.Reservation{User: uid}}
	res, err := serv.HasReservations(context, in)
	if res.Has && err == nil {
		return true
	}
	return false
}

/*
DeleteUser will delete a user from the map.
*/
func (u *UserHandlerService) DeleteUser(context context.Context, request *proto.DeleteUserRequest, response *proto.DeleteUserResponse) error {
	if u.containsID(request.User.Userid) {
		if u.HasOpenReservations(context, request.User.Userid) {
			u.mutex.Lock()
			delete(u.user, request.User.Userid)
			response.IsDeleted = true
			u.mutex.Unlock()
		}
		response.IsDeleted = false
		return nil
	}
	return fmt.Errorf("cannot delete user with the id %d", request.User.Userid)
}

/*
GetInformationFromMap will find a users information by a given parameter for example a string (Name) or his id (int32)
*/
func (u *UserHandlerService) GetInformationFromMap(value interface{}) interface{} {
	switch tp := value.(type) {
	case string:
		u.mutex.Lock()
		for k, v := range *u.getUserMap() {
			if v.getName() == tp {
				u.mutex.Unlock()
				return k
			}
		}
		u.mutex.Unlock()
		return int32(-1)
	case int32:
		u.mutex.Lock()
		name := ((*u.getUserMap())[tp]).getName()
		u.mutex.Unlock()
		if name != "" {
			return name
		}
		return ""
	}
	return nil
}

/*
FindUser will find a user in the datastructure. (For finding a user by his Name see FindUserByName).
*/
func (u *UserHandlerService) FindUser(context context.Context, in *proto.FindUserRequest, out *proto.FindUserResponse) error {
	if u.containsID(in.GetUser().GetUserid()) {
		name := u.GetInformationFromMap(in.GetUser().GetUserid()).(string)
		if name != "" {
			out.User.Userid = in.GetUser().GetUserid()
			out.User.Name = name
			return nil
		}
	}
	return fmt.Errorf("cannot find a user with the id %d", in.GetUser().GetUserid())
}

/*
FindUserByName will find a user by his Name and not by his ID. (For finding a user by his ID see FindUser).
*/
func (u *UserHandlerService) FindUserByName(context context.Context, in *proto.FindUserByNameRequest, out *proto.FindUserResponse) error {
	if in.User.GetName() != "" {
		userID := u.GetInformationFromMap(in.GetUser().GetName()).(int32)
		if userID != -1 {
			out.User.Userid = userID
			out.User.Name = in.User.GetName()
			return nil
		}
		return fmt.Errorf("there is no user with the name %s", out.User.GetName())
	}
	return fmt.Errorf("cannot find an empty username")
}

/*
change will change a user in the datastructure.
*/
func (u *UserHandlerService) change(id int32, pname string) bool {
	if u.containsID(id) {
		u.mutex.Lock()
		(*u.getUserMap())[id] = &userS{name: pname}
		u.mutex.Unlock()
		return true
	}
	return false
}

/*
ChangeUser will change a user in the datastructure to a given value.
*/
func (u *UserHandlerService) ChangeUser(context context.Context, in *proto.ChangeUserRequest, out *proto.ChangeUserResponse) error {
	if in.Change.Name != "" && u.containsID(in.Change.Userid) {
		if u.change(in.Change.Userid, in.Change.Name) {
			out.Change = true
			return nil
		}
	}
	return fmt.Errorf("cannot change the user. The userid or the name are not ok. See: %d %s", in.Change.Userid, in.Change.Name)
}

/*
ReceiveAndSendAllUsers will return a list with all users in our current service (id and name).
*/
func (u *UserHandlerService) ReceiveAndSendAllUsers(context context.Context, in *proto.AllUsersRequest, out *proto.AllUsersResponse) error {
	if len(*u.getUserMap()) > 0 {
		users := []*proto.UserMessageResponse{}
		for k, v := range *u.getUserMap() {
			users = append(users, &proto.UserMessageResponse{Userid: k, Name: v.getName()})
		}
		out.Users = users
	}
	return fmt.Errorf("there are currently no users store (Advice: find some customers)")
}
