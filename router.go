package main

import (
	"Dormitory-Distribution-System/controller"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Results struct {
	Name string `json:"name"`
}

type QuestionnaireInfo struct {
	QID    int    `json:"qid"`
	Title  string `json:"title"`
	State  string `json:"state"`
	Number int    `json:"number"`
	ID     string `json:"id"`
}
type QuestionnaireData struct {
	ID  string `json:"id"`
	Qid struct {
		IsTrusted bool `json:"isTrusted"`
	} `json:"qid"`
	StudentId          interface{} `json:"studentId"`
	Name               string      `json:"name"`
	Sex                string      `json:"sex"`
	Major              string      `json:"major"`
	Age                string      `json:"age"`
	Home               []string    `json:"home"`
	Ethnic             string      `json:"ethnic"`
	SleepTime          interface{} `json:"sleepTime"`
	GetupTime          interface{} `json:"getupTime"`
	SameRoutine        interface{} `json:"sameRoutine"`
	LearnInDorm        interface{} `json:"learnInDorm"`
	NeatExpection      interface{} `json:"neatExpection"`
	CleanPeriod        interface{} `json:"cleanPeriod"`
	BathePeriod        interface{} `json:"bathePeriod"`
	Expense            interface{} `json:"expense"`
	CostType           []string    `json:"costType"`
	OutCost            interface{} `json:"outCost"`
	ShareCost          interface{} `json:"shareCost"`
	Hobby              []string    `json:"hobby"`
	HobbySameExpection interface{} `json:"hobbySameExpection"`
	WantCommunicate    interface{} `json:"wantCommunicate"`
	Smoke              interface{} `json:"smoke"`
	Drink              interface{} `json:"drink"`
	Snore              interface{} `json:"snore"`
	SleepQuality       interface{} `json:"sleepQuality"`
}
type UserBaseInfo struct {
	UID                    uint   `gorm:"column:uid;primaryKey;autoIncrement"`
	Name                   string `gorm:"column:name"`
	Sex                    string `gorm:"column:sex"`
	Major                  string `gorm:"column:major"`
	Age                    string `gorm:"column:age"`
	Homestr                string `gorm:"column:home"`
	SychronizedSchedule    string `gorm:"column:sychronizedSchedule"`
	SpendingResponsibility string `gorm:"column:spendingResponsibility"`
	Interests              string `gorm:"column:interests"`
}
type UserQuestionnaireData struct {
	UID                     uint   `gorm:"column:uid;primaryKey;autoIncrement"`
	BedTime                 string `gorm:"column:bedTime"`
	WakeUpTime              string `gorm:"column:wakeUpTime"`
	SleepQuality            string `gorm:"column:sleepQuality"`
	DomStudy                string `gorm:"column:domStudy"`
	Smoke                   string `gorm:"column:smoke"`
	Drink                   string `gorm:"column:drink"`
	Snore                   string `gorm:"column:snore"`
	ChattingSharinsThoushts string `gorm:"column:chattingSharinsThoushts"`
	Leanliness              string `gorm:"column:leanliness"`
	Cleaningfrsgueney       string `gorm:"column:cleaningfrsgueney"`
	ShowerFrequency         string `gorm:"column:showerkrequency"`
	MonthlyBudget           string `gorm:"column:monthlyBudset"`
	JointOutings            string `gorm:"column:jointOutings"`
	SharedExpenses          string `gorm:"column:sharedExpenses"`
	SharedInterests         string `gorm:"column:sharedInterests"`
}

func InitRouter(r *gin.Engine) {

	g1 := r.Group("/user")
	{
		g1.POST("/login/", controller.Login)
	}
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})
	r.GET("/questionnaireInfo", func(c *gin.Context) {
		questionnaireInfo := []QuestionnaireInfo{
			{0, "2023 Freshman Second Questionnaire", "Enabled", 0, "0daaas"},
		}
		c.JSON(http.StatusOK, questionnaireInfo)
	})

	r.GET("/results", func(c *gin.Context) {
		results := []Results{
			{"林浩"},
			{"冯国强"},
			{"陈国华"},
			{"吴国强"},
			{"李玉英"},
		}
		c.JSON(http.StatusOK, results)
	})

	r.OPTIONS("/questionnaire", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		c.Status(http.StatusOK)
	})
	r.POST("/questionnaire", func(c *gin.Context) {
		var requestData QuestionnaireData

		if err := c.BindJSON(&requestData); err != nil {
			c.JSON(400, gin.H{"error": "Failed to parse JSON"})
			return
		}
		// for i := 0; i < 60; i++ {
		dsn := "gorm.db"
		db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		db.AutoMigrate(&UserBaseInfo{})
		var data2 UserBaseInfo
		data2.Age = requestData.Age
		data2.Name = requestData.Name
		data2.Major = requestData.Major
		data2.Homestr = strings.Join(requestData.Home, ",")
		if requestData.Sex == "男" {
			data2.Sex = "0"
		} else {
			data2.Sex = "1"
		}
		data2.SychronizedSchedule = requestData.SameRoutine.(string)
		data2.SpendingResponsibility = strings.Join(requestData.CostType, ",")
		data2.Interests = strings.Join(requestData.Hobby, ",")
		
			db.Create(&data2)
			// var uu = new(UserBaseInfo)
			// db.First(uu)
			// fmt.Printf("%#v\n", uu)

			db.AutoMigrate(&UserQuestionnaireData{})
			var data UserQuestionnaireData

			data.UID = data2.UID
			data.BedTime = requestData.SleepTime.(string)
			data.WakeUpTime = requestData.GetupTime.(string)
			data.SleepQuality = requestData.SleepQuality.(string)
			data.DomStudy = requestData.LearnInDorm.(string)
			data.Smoke = requestData.Smoke.(string)
			data.Drink = requestData.Drink.(string)
			data.Snore = requestData.Snore.(string)
			data.ChattingSharinsThoushts = requestData.WantCommunicate.(string)
			data.Leanliness = requestData.NeatExpection.(string)
			data.Cleaningfrsgueney = requestData.CleanPeriod.(string)
			data.ShowerFrequency = requestData.BathePeriod.(string)
			data.MonthlyBudget = requestData.Expense.(string)
			data.JointOutings = requestData.OutCost.(string)
			data.SharedExpenses = requestData.ShareCost.(string)
			data.SharedInterests = requestData.HobbySameExpection.(string)

			db.Create(&data)
		// }
		// db.Commit()
		// var u = new(UserQuestionnaireData)
		// db.Last(u)
		// fmt.Printf("%#v\n", u)

	})
}
