package reop_test

type UserData struct {
	Name     string `json:"name"`
	UserId   string `json:"userId"`
	ClassId  string `json:"classId"`
	LessonId string `json:"lessonId"`
	Nums     int    `json:"nums"`
	Score    int    `json:"score"`
	Time     int    `json:"time"`
}

//func SimulateRopeData() [5]UserData {
//	var userListUser [5]UserData
//	store := service.CreateLinkMongoDB("rope_test_data")
//	userBate, err := store.GetCollectAllData()
//	if err != nil {
//		fmt.Println(err)
//		//my_log.WriteLogErrs().Println(err)
//	}
//	ctx := context.Background()
//	defer userBate.Close(ctx)
//	i := 0
//	for userBate.Next(ctx) {
//		if err := userBate.Decode(&userListUser[i]); err != nil{
//			//my_log.WriteLogErrs().Println(err)
//		}
//		fmt.Println(userListUser[i])
//		i++
//	}
//	return userListUser
//}
