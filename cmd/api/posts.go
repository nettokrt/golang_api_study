package main

import (
	"net/http"
	"errors"
	"strconv"
	"github.com/go-chi/chi/v5"
	"github.com/nettokrt/golang_api_study/internal/store"
)

type CreatePostPayload struct {
	Title 	string `json:"title"`
	Content string `json:"content"`
	Tags 	[]string `json:"tags"`
} 


func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreatePostPayload 

	if err := readJSON(w, r, &payload); err!= nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}
	
	post := &store.Post{
		Title:     payload.Title,
    Content:   payload.Content,
		Tags: 		 payload.Tags,
		UserID:    2, // Hardcoded for now
	}

	ctx := r.Context()

	if err := app.store.Posts.Create(ctx, post); err!= nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := writejSON(w, http.StatusCreated, post); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
	}
}

func (app *application) getPostHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "postID")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid post id")
		return
	}
	ctx := r.Context()

	post, err := app.store.Posts.GetByID(ctx, id);
		if err != nil {
			switch {
			case errors.Is(err, store.ErrPostNotFound):	
				writeJSONError(w, http.StatusNotFound, err.Error())
			default:
				writeJSONError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}
	if err := writejSON(w, http.StatusCreated, post); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
	}
}
