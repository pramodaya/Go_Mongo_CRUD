package traveluser

type TravelUser struct {
	ID                      string `json:"id" bson:"_id,omitempty"`
	Username                string `json:"username" bson:"username"`
	Email                   string `json:"email" bson:"email"`
	IsEmailVerified         bool   `json:"isEmailVerified" bson:"isEmailVerified"`
	ContactNumber           string `json:"contactNumber" bson:"contactNumber"`
	IsContactNumberVerified bool   `json:"isContactNumberVerified" bson:"isContactNumberVerified"`
	Address                 string `json:"address" bson:"address"`
	Country                 string `json:"country" bson:"country"`
	Age                     int    `json:"age" bson:"age"`
}
