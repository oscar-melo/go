package business

import (
    "fmt"
//    "log"
    "net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"general/users/rolesapi/forms"
	"io/ioutil"
	"strconv"

	"general/users/rolesapi/config"
    
)

var Articles []forms.Article

func InitData(){

	fmt.Println("Rest API v2.0 - Mux Routers")

	fmt.Println(config.Cfg.Server.Port)


	Articles = []forms.Article{
        forms.Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
        forms.Article{Id: "2", Title: "Hide.llo 2", Desc: "Article Description", Content: "Article Content"},
    }
}

func ReturnAllArticles(w http.ResponseWriter, r *http.Request){
	
    fmt.Println("Endpoint Hit: returnAllArticles")
    json.NewEncoder(w).Encode(Articles)
}

func ReturnSingleArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnSingleArticle")
    vars := mux.Vars(r)
    key := vars["id"]

    // Loop oidv.er all of our Articles
    // if the article.Id equals the key we pass in
    // return the article encoded as JSON
    for _, article := range Articles {
        if article.Id == key {
            json.NewEncoder(w).Encode(article)
        }
    }
}

func CreateNewArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: createNewArticle")
    // get the body of our POST request
    // unmarshal this into a new Article struct
    // append this to our Articles array.    
    reqBody, _ := ioutil.ReadAll(r.Body)
    var article forms.Article 
    json.Unmarshal(reqBody, &article)
    // update our global Articles array to include
    // our new Article
    Articles = append(Articles, article)

    json.NewEncoder(w).Encode(article)
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: deleteArticle")
    // once again, we will need to parse the path parameters
    vars := mux.Vars(r)
    // we will need to extract the `id` of the article we
    // wish to delete
    id := vars["id"]

    // we then need to loop through all our articles
    for index, article := range Articles {
        // if our id path parameter matches one of our
        // articles
        if article.Id == id {
            // updates our Articles array to remove the 
            // article
            Articles = append(Articles[:index], Articles[index+1:]...)
        }
    }

}

func UpdateArticle (w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: updateArticle")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	


	reqBody, _ := ioutil.ReadAll(r.Body)
    var article forms.Article 
    json.Unmarshal(reqBody, &article)

	Articles[id-1].Title = article.Title
	Articles[id-1].Desc = article.Desc
	Articles[id-1].Content = article.Content

	fmt.Print("Error")
	fmt.Println(err)

}