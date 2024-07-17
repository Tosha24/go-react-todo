package models

// create user model with the following fields: id, name, email, password
type User struct {
	ID       string `json:"id,omitempty" bson:"_id,omitempty"`
	Email    string `json:"email" bson:"email" unique:"true"`
	Password string `json:"password" bson:"password"`
	Todos    []TODO `json:"todos" bson:"todos"`
}