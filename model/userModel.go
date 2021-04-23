package model

import "time"

type (
	User struct {
		ID        string    `json:"id" bson:"_id,omitempty"`
		Username  string    `json:"username" bson:"username,omitempty"`
		Password  string    `json:"password,omitempty" bson:"password,omitempty"`
		Email     string    `json:"email" bson:"email,omitempty"`
		Role      string    `json:"role" bson:"role"`
		Name      string    `json:"name" bson:"name"`
		CreatedAt time.Time `json:"created_at" bson:"created_at,omitempty"`
		UpdatedAt time.Time `json:"updated_at" bson:"updated_at,omitempty"`
	}

	GetAll struct {
		From     int    `json:"from" form:"from" example:""`
		Limit    int    `json:"limit" form:"limit" example:""`
		SortType string `json:"sort_type" form:"sort_type" example:""`
		SortBy   string `json:"sort_by" form:"sort_by" example:""`
		Search   string `json:"search" form:"search" example:""`
	}
)
