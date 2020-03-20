package model

//
// Representation of Link domain
//
type Link struct {
	LinkID  string `bson:"_id"`
	UserID  string `bson:"userID"`
	Name    string `bson:"name"`
	Address string `bson:"address"`
}
