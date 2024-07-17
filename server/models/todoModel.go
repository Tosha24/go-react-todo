package models

// create todo model with the following fields: id, title, description, completed
type TODO struct {
	ID          string `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string `json:"title" bson:"title"`
	Completed   bool   `json:"completed" bson:"completed"`
}
