package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

//структура для учётной записи пользователя
type Account struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token";sql:"-"`
}

//расширенный мониторинг воркера(успешно ли доавлен пользователь и какой)
type JobStatusAdd struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	Token     string
	Status    bool
	UserName  string
}

// таблица для мониторинга работы воркера
type JobStatusController struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Token     string
	Status    string
}

//таблица хранит пользовательские запросы на стоп
type JobCanceledActions struct {
	ID        uint
	CreatedAt time.Time
	Token     string
}

func GetUser(u uint) *Account {

	acc := &Account{}
	GetDB().Table("accounts").Where("id = ?", u).First(acc)
	if acc.Email == "" { //Пользователь не найден!
		return nil
	}

	acc.Password = ""
	return acc
}

func InsertJobStatus(token string, status string) {
	result := &JobStatusController{}
	result.Status = status
	result.Token = token

	GetDB().Table("job_status_controllers").Create(result)
}
func InsertJobStatusAdd(token string, username string, status bool) {
	result := &JobStatusAdd{}

	result.Token = token
	result.Status = status
	result.UserName = username
	GetDB().Table("job_status_adds").Create(result)
}

func GetJobStatusFromDb(token string) *JobStatusController {
	result := &JobStatusController{}
	GetDB().Table("job_status_controllers").Order("created_at desc").Where("token = ?", token).First(result)
	return result
}

func CheckActionCancelJob(token string) bool {
	result := &JobCanceledActions{}

	GetDB().Table("job_canceled_actions").Where("token = ?", token).First(result)
	var status bool
	fmt.Printf("заходим и проверяем нужно ли останавливать воркер\n")
	fmt.Printf("Токен, который прилетает из базы:\n")
	fmt.Printf(result.Token + "\n")
	if result.Token == "" {
		fmt.Printf("Токена нет, продолжаем работу воркера\n")
		status = false
	} else {
		fmt.Printf("Токена есть, останавливаем работу воркера\n")
		status = true
	}
	if status {
		GetDB().Table("job_canceled_actions").Where("token = ?", token).Delete(result)
	}
	fmt.Printf("Конечный результат")

	return status
}

func InsertCancelAction(token string) {
	result := &JobCanceledActions{}
	result.Token = token
	//GetDB().Table("job_canceled_actions").Update(result)
	GetDB().Model(&JobCanceledActions{}).Where("token = ?", token).FirstOrCreate(result)

}

func AddUser() {

	//GetDB().Model(&acc).Update("Token", 200)
	db.Create(&Account{Email: "jfkddf", Password: "fsdfdsfdsf", Token: "sdfdfdsfddfs"})

}
