package users

import (
	"math/rand"

	proto "github.com/ob-vss-ws19/blatt-4-pwn2own/users/proto"
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
@param length will be the length of the number
@param seed will be a seed in order to produce "really" random numbers
*/
func (u *UserHandlerService) getRandomUserID(length int, seed int64) int {
	for {
		potantialID := rand.Intn(length)
		if !u.containsID(int32(potantialID)) {
			return potantialID
		}
	}
}

/*
getUserMap will return a pointer to the current user map in order to work in that.
*/
func (u *UserHandlerService) getUserMap() *map[int32]*proto.UserMessageResponse {
	return &u.user
}

/*
setUserMap will set the map of a userhandlerservice instance.
@users will be the map to set.
*/
func (u *UserHandlerService) setUserMap(users *map[int32]*proto.UserMessageResponse) {
	u.user = *users
}

/*
UserHandlerService will be the representation of our service.
*/
type UserHandlerService struct {
	user         map[int32]*proto.UserMessageResponse
	dependencies []interface{}
}
