package exemplo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type GetData interface {
	getArticle() []Article
}

type article Article

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func (a article) getArticle() []Article {
	return []Article{
		{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
}

func HandleRequest() {

	myRoute := mux.NewRouter().StrictSlash(true)
	myRoute.HandleFunc("/", homePage)
	myRoute.HandleFunc("/article", createArticle).Methods("POST")
	myRoute.HandleFunc("/article", returnAllArticles)

	log.Fatal(http.ListenAndServe(":8080", myRoute))

}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: returnAllArticles")

	var getData GetData
	var a Article
	getData = article(a)

	var Articles = getData.getArticle()
	json.NewEncoder(w).Encode(Articles)

}

func createArticle(w http.ResponseWriter, r *http.Request) {
	reqbody, _ := ioutil.ReadAll(r.Body)

	var getData GetData
	var a Article
	getData = article(a)

	var articles = getData.getArticle()
	json.Unmarshal(reqbody, &articles)

	articles = append(getData.getArticle(), a)

	json.NewEncoder(w).Encode(a)
}
