package models

import "go.mongodb.org/mongo-driver/bson/primitive"

/*User Modelo de la base de datos de Usuarios*/
type User struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Name        string               `bson:"name" json:"name,omitempty"`
	LName       string               `bson:"lname" json:"lname,omitempty"`
	Email       string               `bson:"email" json:"email"`
	Password    string               `bson:"password" json:"password"`
	Type        string               `bson:"type" json:"type,omitempty"` //Cliente o
	Phone       []string             `bson:"phone" json:"phone,omitempty"`
	ConvPhone   []string             `bson:"convPhone" json:"convPhone,omitempty"`
	Permissions []primitive.ObjectID `bson:"permissions" json:"permissions,omitempty"`
	State       int                  `bson:"state" json:"state"`
}
