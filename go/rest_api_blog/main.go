package main

import (
	. "blog/model"

	"encoding/json"
	"net/http"
	"path"
	"strconv"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/posts", handleGetList)
	http.HandleFunc("/posts/", handleRequest)

	server.ListenAndServe()
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	case "POST":
		err = handlePost(w, r)
	case "PUT":
		err = handlePut(w, r)
	case "DELETE":
		err = handleDelete(w, r)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleGetList(w http.ResponseWriter, r *http.Request) {
	var err error
	posts, err := GetPosts(100)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	output, err := json.MarshalIndent(&posts, "", "\t")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	_, err = w.Write(output)
	return
}

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}

	post, err := Retrieve(id)
	if err != nil {
		return
	}

	output, err := json.MarshalIndent(&post, "", "\t")
	if err != nil {
		return
	}

	w.WriteHeader(200)
	_, err = w.Write(output)
	return
}

func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	contentLength := r.ContentLength
	contentBody := make([]byte, contentLength)
	_, err = r.Body.Read(contentBody)
	if err != nil {
		return
	}

	var post Post
	err = json.Unmarshal(contentBody, &post)
	if err != nil {
		return
	}
	err = post.Create()
	if err != nil {
		return
	}

	output, err := json.MarshalIndent(&post, "", "\t")
	if err != nil {
		return
	}

	w.WriteHeader(200)
	_, err = w.Write(output)
	return
}

func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}

	post, err := Retrieve(id)
	if err != nil {
		return
	}

	contentLength := r.ContentLength
	contentBody := make([]byte, contentLength)
	_, err = r.Body.Read(contentBody)
	if err != nil {
		return
	}

	err = json.Unmarshal(contentBody, &post)
	if err != nil {
		return
	}

	err = post.Update()
	if err != nil {
		return
	}

	output, err := json.MarshalIndent(&post, "", "\t")
	if err != nil {
		return
	}

	w.WriteHeader(200)
	_, err = w.Write(output)
	return
}

func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}

	post, err := Retrieve(id)
	if err != nil {
		return
	}

	err = post.Delete()
	if err != nil {
		return
	}

	output, err := json.MarshalIndent(&post, "", "\t")
	if err != nil {
		return
	}

	w.WriteHeader(200)
	_, err = w.Write(output)
	return
}
