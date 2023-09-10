package dto

type StudentResponse struct {
	ID       uint   `json:"id"`
	NIM      string `json:"nim"`
	Fullname string `json:"fullname"`
	Majority string `json:"majority"`
	Address  string `json:"address"`
	Image    string `json:"image"`
}

type StudentRequest struct {
	NIM      string `json:"nim" form:"nim" binding:"required"`
	Fullname string `json:"fullname" form:"fullname" binding:"required"`
	Majority string `json:"majority" form:"majority" binding:"required"`
	Address  string `json:"address" form:"address" binding:"required"`
}

type UpdateStudentRequest struct {
	NIM      string `json:"nim" form:"nim"`
	Fullname string `json:"fullname" form:"fullname"`
	Majority string `json:"majority" form:"majority"`
	Address  string `json:"address" form:"address"`
}
