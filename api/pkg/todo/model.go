package todo

import "go.mongodb.org/mongo-driver/bson/primitive"

/**
 * Document Schema for Todo table.
 * @type {struct} Todo
 */
type Todo struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title,omitempty"`
	Details     string             `json:"details" bson:"details,omitempty"`
	Date        string             `json:"date" bson:"date,omitempty"`
	IsCompleted bool               `json:"isCompleted" bson:"isCompleted,omitempty" default:"false"`
}
