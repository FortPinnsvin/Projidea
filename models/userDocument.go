package models

type UserDocument struct {
	Id 		 string `bson:"_id,omitempty"`
	Name 	 string
	Password string
}
