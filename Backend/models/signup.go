package models

// User represents a user in the system
type SignUp struct {
	ID                  string    `json:"_id" bson:"_id,omitempty"`
	Username            string    `json:"Username" bson:"Username"`
	Password            string    `json:"Password" bson:"Password"`
	Email			   string    `json:"Email" bson:"Email"`
	Role			   string    `json:"Role" bson:"Role"`
	UnencryptedPassword  string    `json:"un" bson:"un"`
}

