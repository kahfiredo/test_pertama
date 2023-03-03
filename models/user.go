package models

/*
	@Author: DevProblems(Sarang Kumar)
	@YTChannel: https://www.youtube.com/channel/UCVno4tMHEXietE3aUTodaZQ
*/

type User struct {
	Nama    string  `json:"nama" bson:"nama"`
	Umur     int     `json:"umur" bson:"umur"`
	Kelas string `json:"kelas" bson:"kelas"`
}
