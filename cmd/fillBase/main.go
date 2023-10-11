package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"runtime"
	"sync"
	"time"
)

const (
	BasePath     string = "http://localhost:8080"
	RegRoute     string = "auth"
	ManagerRoute string = "sign-up-manager"
	ClientRoute  string = "sign-up-client"
	CountEntity  int    = 50000
)

var charset []byte = []byte("abcdefghijklmnopqrstuvwxyz")

type (
	ManagerRegistrationRequest struct {
		Login    string `json:"login"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	ClientRegistrationRequest struct {
		Login     string `json:"login"`
		Password  string `json:"password"`
		Email     string `json:"email"`
		ManagerId int    `json:"managerId"`
	}
	ManagersData []ManagerRegistrationRequest
	ClientsData  []ClientRegistrationRequest
)

func main() {
	var wg sync.WaitGroup

	countWorkers := 1000

	wg.Add(countWorkers)

	client := &http.Client{}

	uri := fmt.Sprintf("%s/%s/%s", BasePath, RegRoute, ManagerRoute)

	managersData := fillManagers()

	go func() {
		defer TimeTrack(time.Now())
		work(managersData, client, uri)
		defer wg.Done()
	}()
	defer TimeTrack(time.Now())
	wg.Wait()

	fmt.Printf("Работа закончена")
}

func fillManagers() ManagersData {
	data := make(ManagersData, CountEntity)
	for i := 0; i < CountEntity; i++ {
		data[i].Login = randomString(10)
		data[i].Password = randomString(10)
		data[i].Email = fmt.Sprintf("%s%s", randomString(10), "@mail.ru")
	}
	return data
}
func randomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}
func printManagers(managers ManagersData) {
	for i, manager := range managers {
		fmt.Printf("%d  Login - %s \n", i, manager.Login)
		fmt.Printf("   Password - %s \n", manager.Password)
		fmt.Printf("   Email - %s \n", manager.Email)
	}
}
func makeReqDataManager(manager ManagerRegistrationRequest) []byte {
	rd, _ := json.Marshal(manager)
	return rd
}
func makeReqPOST(req []byte, uri string) *http.Request {
	r, _ := http.NewRequest(http.MethodPost, uri, bytes.NewBuffer(req))
	return r
}
func work(data ManagersData, client *http.Client, uri string) {
	for i := 0; i < CountEntity; i++ {
		rd := makeReqDataManager(data[i])
		r := makeReqPOST(rd, uri)
		response, err := client.Do(r)
		if err != nil {
			fmt.Printf("ошибка работы клиента - %s \n", err.Error())
		}
		defer response.Body.Close()
	}
}
func TimeTrack(start time.Time) {
	elapsed := time.Since(start)

	pc, _, _, _ := runtime.Caller(1)

	funcObj := runtime.FuncForPC(pc)

	runtimeFunc := regexp.MustCompile(`^.*\.(.*)$`)
	name := runtimeFunc.ReplaceAllString(funcObj.Name(), "$1")

	log.Println(fmt.Sprintf("%s took %s", name, elapsed))
}