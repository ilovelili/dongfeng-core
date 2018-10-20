package models

// UserFriend user friend entry
type UserFriend struct {
	UserID   string `dapper:"user_id,primarykey,table=user_friends"`
	FriendID string `dapper:"friend_id"`
}
