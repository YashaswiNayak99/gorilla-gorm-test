package services

import (
	"log"
	"strconv"
	"net/http"
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/gorilla/mux"
	"github.com/YashaswiNayak99/gorilla-gorm-test/models"
)

var dbconn *gorm.DB

type Response struct {
	Data []models.Post `json:"data"`
	Message string `json:"message"`
}

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var posts = models.GetPosts()
	var resp Response
	err := dbconn.Find(&posts).Error
	if err == nil {
		log.Println(posts)
		resp.Data = posts
		resp.Message = "SUCCESS"
		json.NewEncoder(w).Encode(&resp)
	} else {
		log.Println(err)
		http.Error(w, err.Error(), 400)
	}
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var resp Response
	var post = models.GetPost()
	err := dbconn.Where("id = ?",  id).Find(&post).Error
	if err == nil {
		log.Println(post)
		resp.Data = append(resp.Data, post)
		resp.Message = "SUCCESS"
		json.NewEncoder(w).Encode(&resp)
	} else {
		log.Println(err)
		http.Error(w, err.Error(), 400)
	}
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var resp Response
	var post = models.GetPost()
	_ = json.NewDecoder(r.Body).Decode(&post)
	log.Println(post)

	err := dbconn.Create(&post).Error
	if err != nil {
		http.Error(w, "Error Creating Record", 400)
		return
	}
	resp.Message = "CREATED"
	json.NewEncoder(w).Encode(resp)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var resp Response
	var post = models.GetPost()
	_ = json.NewDecoder(r.Body).Decode(&post)
	id, _ := strconv.Atoi(params["id"])
	
	err := dbconn.Model(&post).Where("id = ?",id).Update(&post).Error
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	resp.Message = "UPDATED"
	json.NewEncoder(w).Encode(resp)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var resp Response
	var post = models.GetPost()
	err := dbconn.Delete(&post, params["id"]).Error
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	resp.Message = "DELETED"
	json.NewEncoder(w).Encode(resp)
}

func SetDB(db *gorm.DB) {
	dbconn = db
	var post = models.GetPost()
	dbconn.AutoMigrate(&post)
}