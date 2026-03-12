# GO 튜토리얼

## 다른 모듈에서 코드 호출

```
<홈>/
 |-- hello/
 |-- greetings/
```

```bash
mkdir hello
cd hello

git mod init example.com/hello

```

```go
package main

import (
    "fmt"

    "example.com/greetings"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}
```

```bash
go mod edit -replace example.com/greetings=../greetings
```

```bash
go mod tidy
```

## 오류를 반환하고 처리

```go
// greetings/greetings.go

package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person
func Hello(name string) (string, error) {

	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

 	// If a name was received, return a value that embeds the name
  // in a greeting message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
```

```go
// hello/hello.go

package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
```

## 임의의 인사말을 보내세요

```go
// greetings/greetings.go

package greetings

import (
    "errors"
    "fmt"
    "math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    // If no name was given, return an error with a message.
    if name == "" {
        return name, errors.New("empty name")
    }
    // Create a message using a random format.
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
    // A slice of message formats.
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }

    // Return a randomly selected message format by specifying
    // a random index for the slice of formats.
    return formats[rand.Intn(len(formats))]
}
```

## 여러 사람에게 답례 인사

```go
// greetings/greetings.go

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}
```

```go
// hello/hello.go

	// A slice of names.
	names := []string{"hyunwoo", "eunbi", "jian"}

	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)

	for _, message := range messages {
		fmt.Println(message)
	}
```

## 테스트

```go
package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b`+name+`\b`)
	msg, err := Hello("Gladys")
	if !want.MatchString(msg) || err != nil {
		t.Errorf(`Hello("Gladys")) = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Errorf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
```

> 테스트가 많아지면 함수가 너무 많아짐 그래서 `Table Driven Test` 사용

```go
func TestHello(t *testing.T) {

	tests := []struct {
		name string
		wantErr bool
	} {
		{"Hyunwoo", false},
		{"Eunbi", false},
		{"", true},
	}

	for _, tt := range tests {

		msg, err := Hello(tt.name)

		if tt.wantErr {
			if err == nil {
				t.Errorf("Hello(%q) expected error", tt.name)
			}

			continue
		}

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if !strings.Contains(msg, tt.name) {
			t.Errorf("Hello(%q) = %q", tt.name, msg)
		}
	}
}
```

## `naked` return

```go
package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string)  (length int, uppercase string) {
	defer fmt.Println("I'm done")

	length = len(name)
	uppercase = strings.ToUpper(name)
	return

}

func main() {

	totalLength, upperName := lenAndUpper("hyunwoo")

	fmt.Println(totalLength, upperName)

}
```

## defer

> 함수가 끝난 후에 어떤 동작을 해주게 함

```go
func lenAndUpper(name string)  (length int, uppercase string) {
	defer fmt.Println("I'm done")

	length = len(name)
	uppercase = strings.ToUpper(name)
	return

}
```

## range

```go
func superAdd(numbers ...int) int {

	result := 0

	for _, number := range numbers {
		result += number
	}

	return result
}

superAdd(1, 2, 3, 4, 5, 6) // 21
```

## pointer

`&` 메모리 주소

`*` 탐색

```go
func main() {
	a := 2
	b := &a
	a = 5

	fmt.Println(*b) // 5

  *b = 20

  fmt.Pintln(a) // 20
}
```

> 아주 무거운 object를 여러 곳에서 참조해야한다면 계속 복사본을 만들지 말고 이 개념을 사용하자 !!

## slice

```go
func main() {

	names := []string{"nico", "lynn", "dal", "hyunwoo", "eunbi"}
	names = append(names, "jian")

	fmt.Println(names)

}
```

## map

```go
func main() {

	nico := map[string]string{"name": "nico", "age": "12"}

	for _, value := range nico {

		fmt.Println(value)
    // nico
    // 12
	}

	fmt.Println(nico) // map[age:12 name:nico]
}
```

## struct

```go
type person struct {
	name string
	age int
	favFood []string
}

func main() {

	favFood := []string{"kimchi", "ramen"}

	nico := person{name:"nick", age:18, favFood: favFood}
	fmt.Println(nico.name)
}
```

## Account tutorial

```go
// main.go
package main

import (
	"fmt"

	"github.com/hyunwoomemo/learngo/accounts"
)

func main() {

	account := accounts.NewAccount("hyunwoo")
	account.Deposit(10)
	err := account.Withdraw(20)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(account)
}
```

```go
// account.go
package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner string
	balance int
}

var noMoney = errors.New("Can't withdraw you are poor")

// NewAccount account 만드는 함수
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

func (a Account) Balance() int {
	return a.balance
}

func (a *Account) Withdraw(amount int) error {

	if a.balance < amount {
		return noMoney
	}

	a.balance -= amount
	return nil
}

func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
}
```

## Pointer Receiver VS Value Receiver

**_값 변경 → pointer receiver_**

```go
func (a *Account) Deposit()
func (a *Account) Withdraw()
func (a *Account) ChangeOwner()
```

**_읽기만 → value receiver_**

```go
func (a Account) Balance()
func (a Account) Owner()
```

> 그런데 Go에서는 `struct`면 대부분 `pointer receiver`로 통일
>
> - 1️⃣ 복사 비용 줄이기
> - 2️⃣ 일관성 유지
> - 3️⃣ method set 문제 방지

```go
func (a *Account) Balance() int {
	return a.balance
}
```

---

## Stringer 인터페이스 구현

`Go에서 매우 좋은 패턴`

```go
func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
}

fmt.Println(account)

/// Hyunwoo's account
/// Has: 100
```

---

## Dictionary tutorial

- 타입에 메소드 사용 가능

#### Search Method

```go
// mydict.go
package mydict

import "errors"

type Dictionary map[string]string

var errNotFound =  errors.New("Not Found")

func (d Dictionary) Search(word string) (string, error) {

	value, exists := d[word]

	if exists {
		return value, nil
	}

	return "", errNotFound

}
```

```go
// main.go
package main

import (
	"fmt"

	mydict "github.com/hyunwoomemo/dict/dict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}

	dictionary["hello"] = "hello"
	definition, err := dictionary.Search("first")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
}
```

#### Add Method

```go
func (d Dictionary) Add(word, def string) error {

	_, err := d.Search(word)

	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}

	return nil
}
```

#### Update & Delete Method

```go

var errNotExists = errors.New("단어가 존재하지 않습니다.")

// ...

func (d Dictionary) Update(word, def string) error {

	_, err := d.Search(word)

	switch err {
	case errWordExists:
		return errNotExists
	case nil:
		d[word] = def
	}

	return nil
}

func (d Dictionary) Delete(word string) error {

	_, err := d.Search(word)

	switch err {
	case errNotFound:
		return errNotExists
	case nil:
		delete(d, word)
	}

	return nil
}
```

## 숫자 -> 문자열 변환

```go
import "strconv"

num := 42
str := strconv.Itoa(num)

fmt.Println(str) // "42"
```

`Itoa` 의미

> Integer To ASCII

## URL Checker Tutorial - Slow

```go
package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

var errReqeustFailed = errors.New("Request failed")

func main() {

	var results = make(map[string]string)

	urls := []string{
"https://www.airbnb.com/",
"https://www.google.com/",
"https://www.amazon.com/",
"https://www.reddit.com/",
"https://www.google.com/",
"https://soundcloud.com/",
"https://www.facebook.com/",
"https://www.instagram.com/",
"https://academy.nomadcoders.co/",
}

	for _, url := range urls {

		result := "OK"

		statusCode, err := hitURL(url)
		if err != nil {
			result = strconv.Itoa(statusCode)
		}
		results[url] = result
	}

	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitURL(url string) (statusCode int, error error) {

	fmt.Println("Checking", url)

	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		return resp.StatusCode, errReqeustFailed
	}

	return resp.StatusCode, nil
}
```

## URL Checker Tutorial - Fast (Channel)

```go
package main

import (
	"fmt"
	"net/http"
)

type requestResult struct {
	URL string
	Success bool
	StatusCode int
}

func main() {

	results := make(map[string]int)
	c := make(chan requestResult)


	urls := []string{
"https://www.airbnb.com/",
"https://www.google.com/",
"https://www.amazon.com/",
"https://www.reddit.com/",
"https://www.google.com/",
"https://soundcloud.com/",
"https://www.facebook.com/",
"https://www.instagram.com/",
"https://academy.nomadcoders.co/",
}

	for _, url := range urls {

		go hitURL(url, c)
	}

	for i:=0;i<len(urls);i++ {
		result := <-c

		results[result.URL] = result.StatusCode
	}


	fmt.Println(results)
}


func hitURL(url string, c chan<- requestResult) {

	resp, err := http.Get(url)

	success := true

	if err != nil || resp.StatusCode >= 400 {
		success = false
	}

	c <-requestResult{URL: url, StatusCode: resp.StatusCode, Success: success}
}
```

## Scrapper

```go
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id string
	title string
	location string
	date string
	condition []string
}

var baseURL string = "https://www.saramin.co.kr/zf_user/search/recruit?&searchword=python"

func main() {
	c := make(chan []extractedJob)
	var jobs []extractedJob
	pages := getPages()

	fmt.Println(pages)

	for i:=0;i<pages;i++ {
		go getPage(i, c)
	}

	for i:=0;i<pages;i++ {
		result := <-c
		jobs = append(jobs, result...)
	}

	writeJobs(jobs)
}

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkError(err)

	w := csv.NewWriter(file)

	defer w.Flush()

	headers := []string{"ID", "TITLE", "LOCATION", "DATE", "CONDITION"}

	wErr := w.Write(headers)

	for _, job := range jobs {

		jobSlice := []string{job.id, job.title, job.location, job.date, strings.Join(job.condition, " | ")}
		jwErr := w.Write(jobSlice)
		checkError(jwErr)

	}
	checkError(wErr)

}

func getPage(page int, c chan []extractedJob) {
	var jobs []extractedJob
	pageURL := baseURL + "&recruitPage=" + strconv.Itoa(page)

	res, err := http.Get(pageURL)

	checkError(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkError(err)

		doc.Find(".content").Each(func (i int, s *goquery.Selection) {

			s.Find(".item_recruit").Each(func (i int, card *goquery.Selection) {

				job := extractJob(card)
				jobs = append(jobs, job)
			})

	})

	c <- jobs
}

func getPages() int  {
	pages := 0
	res, err := http.Get(baseURL)
	checkError(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkError(err)

	doc.Find(".pagination").Each(func (i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})

	return pages
}

func checkError(err error)  {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.StatusCode)
	}
}

func extractJob(card *goquery.Selection) extractedJob {
		id, _ := card.Attr("value")
				title := cleanString(card.Find(".area_job > .job_tit").Text())
				location := card.Find(".area_job > .job_condition > span").First().Text()

				date := card.Find(".job_date > .date").Text()

				var conditions []string
				card.Find(".area_job > .job_condition > span").Each(func(i int, s *goquery.Selection) {
					if i > 0 {
						conditions = append(conditions, s.Text())
					}
				})
				// fmt.Println(id, exist, title,location,date, conditions)
				return extractedJob{id: id, title: title, location: location, date: date, condition: conditions}

}

func cleanString(str string) string {
	return strings.TrimSpace(str)
}
```

`strings.Join`

> []string 을 "a | b | c"로 변환

```go
// job.condition []string
strings.Join(job.condition, " | ")
```

`strings.TrimSpace(str)`

> 앞 뒤 공백 제거

## Echo

검색창에 직무 검색하면 검색 결과 csv파일 다운로드

```go
package main

import (
	"os"
	"strings"

	"github.com/hyunwoomemo/scrapper/scrapper"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

const fileName string = "jobs.csv"

func handleHome(c *echo.Context) error {
			return c.File("home.html")
}

func handleScrape(c *echo.Context) error {
	defer os.Remove(fileName)
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(term)

	return c.Attachment(fileName, fileName)
}

func main() {
	// scrapper.Scrape("go")

	e := echo.New()
	e.Use(middleware.RequestLogger())

	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)

	if err := e.Start(":1323"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
```

```html
<!doctype html>
<html lang="ko">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Go Jobs</title>
  </head>
  <body>
    <h1>Go Jobs</h1>
    <h3>Saramin.com scrapper</h3>

    <form method="post" action="/scrape">
      <input type="text" placeholder="키워드를 입력하세요." name="term" />
      <button>Search</button>
    </form>
  </body>
</html>
```
