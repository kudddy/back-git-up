package Job

import (
	"git-up-back/MessageTypes"
	"git-up-back/models"
	c "git-up-back/utils"
	"encoding/json"
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	guuid "github.com/google/uuid"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func Workers(in chan MessageTypes.Profile, out chan MessageTypes.Profile) {
	for item := range in {
		fmt.Println("1. Item", item)
		time.Sleep(10 * time.Second)
		item.Name = "Kirill"
		out <- item
	}
}

func StartWorker(token string) {
	// бесконечный цикл
	models.GetMC().Set(&memcache.Item{Key: token, Value: []byte("START")})
	models.InsertJobStatus(token, "START")
	count := 0
	for {
		stop := models.CheckActionCancelJob(token)
		if stop {
			fmt.Printf("Останавливаем работу воркера по требованию пользователя")
			break
		}
		// db.Create(&Account{Email: "jfkddf", Password: "fsdfdsfdsf", Token: "sdfdfdsfddfs"})
		// записываем в базу статус выполнения фоновоый залачи
		models.GetMC().Set(&memcache.Item{Key: token, Value: []byte("RUNNING")})
		models.InsertJobStatus(token, "RUNNING")
		// операция добавления нового пользователя
		fmt.Println("Infinite Loop 1")
		fmt.Println(token)
		username := guuid.New()
		models.InsertJobStatusAdd(token, username.String(), true)
		time.Sleep(20 * time.Second)

		count += 1
		if count == 10 {
			break
			fmt.Printf("Останавливаем работу воркера истечению срока жизни воркера")
		}
	}

	models.InsertJobStatus(token, "FINISH")
	fmt.Println()

}

func ContorellerWorkers(in chan MessageTypes.Profile, out chan MessageTypes.Profile) {
	//метод реализует передачу основному сервису статуса от воркера, а так же их запуск
	// в случае если требуется запуск
	go Workers(in, out)

	// в случае если нужено статус job

	//в случае если нужна остановка работника
}

// страндартный воркер должен иметь следующий функционал: выполнять джоб и писать в базу все действия
// Фронтальный компонент будет забирать статус из базы

func AddFriendWorker(token string) {
	// пытаемся распарсить профиль для получения основного языка програмирования
	// потом берем один из популярных фреймворком и получаем оттуда список юзеров к кому нужно добавиться
	// добавляем раз в несколько минут
	// Определяем язык клиента
	// TODO добавить проверку токена на валидность
	lang := "python"
	rand.Seed(time.Now().UnixNano())
	page := rand.Intn(700)

	fmt.Println(page)

	url := c.MapMainLangToRepo[lang] + "?page=" + strconv.Itoa(page)

	fmt.Println(url)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)

	res, err := client.Do(req)
	fmt.Println(err)
	// TODO нужен еще один цил
	if res.StatusCode == 200 {
		//парсим json
		decoder := json.NewDecoder(res.Body)

		var resp []MessageTypes.StarGazers
		err := decoder.Decode(&resp)
		//fmt.Println(resp)
		if err == nil {
			for _, gazers := range resp {
				client := &http.Client{}
				req, _ := http.NewRequest("PUT", c.UrlAddFriend+gazers.Login, nil)
				req.Header.Set("Accept", c.GitAccept)
				req.Header.Set("Authorization", "token"+" "+token)
				res, err := client.Do(req)

				if err == nil {
					// проверяем все ли успешно
					if res.StatusCode == 204 {
						fmt.Println("Пишем в базу что все успешно")
					} else {
						fmt.Println("Пишем в базу что не получилось")
					}
				} else {
					fmt.Println("Пишем в базу что сетевая ошибка")
				}
				//засыпаем
				time.Sleep(25 * time.Second)
			}
		}
	} else {
		fmt.Println("лалка")
	}
}
