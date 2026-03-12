package scrapper

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


func Scrape(term string) {
	var baseURL string = "https://www.saramin.co.kr/zf_user/search/recruit?&searchword=" + term
	c := make(chan []extractedJob)
	var jobs []extractedJob
	pages := getPages(baseURL)

	fmt.Println(pages)

	for i:=0;i<pages;i++ {
		go getPage(baseURL, i, c)
		// jobs = append(jobs, res...)
	}

	for i:=0;i<pages;i++ {
		result := <-c
		jobs = append(jobs, result...)
	}

	writeJobs(jobs)
}
//https://www.saramin.co.kr/zf_user/search/recruit?=&searchword=python&recruitPage=42&recruitSort=relation&recruitPageCount=40&inner_com_type=&company_cd=0%2C1%2C2%2C3%2C4%2C5%2C6%2C7%2C9%2C10&show_applied=&quick_apply=&except_read=&ai_head_hunting=&mainSearch=n

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

func getPage(url string, page int, c chan []extractedJob) {
	var jobs []extractedJob
	extractJobC := make(chan extractedJob)

	pageURL := url + "&recruitPage=" + strconv.Itoa(page)

	res, err := http.Get(pageURL)

	checkError(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkError(err)

		doc.Find(".content").Each(func (i int, s *goquery.Selection) {
			
			searchCards := s.Find(".item_recruit")
			searchCards.Each(func (i int, card *goquery.Selection) {

				go extractJob(card, extractJobC)
			})

			for i:=0;i<searchCards.Length();i++ {
				job := <- extractJobC
				jobs = append(jobs, job)
			}
	})

	c <- jobs
}

func getPages(baseURL string) int  {
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

func extractJob(card *goquery.Selection, c chan extractedJob) {
		id, _ := card.Attr("value")
				title := CleanString(card.Find(".area_job > .job_tit").Text())
				location := card.Find(".area_job > .job_condition > span").First().Text()

				date := card.Find(".job_date > .date").Text()

				var conditions []string
				card.Find(".area_job > .job_condition > span").Each(func(i int, s *goquery.Selection) {
					if i > 0 {
						conditions = append(conditions, s.Text())
					}
				})
				// fmt.Println(id, exist, title,location,date, conditions)
				c <- extractedJob{id: id, title: title, location: location, date: date, condition: conditions}

}

func CleanString(str string) string {
	return strings.TrimSpace(str)
}