package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/BradPreston/blog-backend/pkg/api"
	"golang.org/x/crypto/bcrypt"
)

// ApiStatus handles the check to see if the API is connected properly
func (s *Server) ApiStatus(w http.ResponseWriter, r *http.Request) {
	ResponseJSON(w, "blog post API is running properly", "success", http.StatusOK)
}

// CreatePost handles inserting a post to the database
func (s *Server) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post api.BlogPost

	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		ResponseJSON(w, "could not read post body", "fail", http.StatusBadRequest)
		return
	}

	err = s.postService.New(post)

	if err != nil {
		ResponseJSON(w, "could not create blog post", "fail", http.StatusBadRequest)
		return
	}

	ResponseJSON(w, "post created successfully", "success", http.StatusCreated)
}

// GetAllPosts handles returning every post from the database
func (s *Server) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := s.postService.GetAll()

	if err != nil {
		ResponseJSON(w, "could not get all posts", "fail", http.StatusBadRequest)
		return
	}

	ResponseJSON(w, posts, "success", http.StatusOK)
}

// GetOnePost handles returning one post from the database
func (s *Server) GetOnePost(w http.ResponseWriter, r *http.Request) {
	uri := strings.Split(r.RequestURI, "/")
	id, err := strconv.Atoi(uri[4])
	if err != nil {
		ResponseJSON(w, fmt.Sprintf("could not find id [%d] in URI", id), "fail", http.StatusBadRequest)
		return
	}

	post, err := s.postService.GetOne(id)
	if err != nil {
		ResponseJSON(w, fmt.Sprintf("could not get post by id: %d", id), "fail", http.StatusBadRequest)
		return
	}

	ResponseJSON(w, post, "success", http.StatusOK)
}

// UpdatePost handles updating one post in the database
func (s *Server) UpdatePost(w http.ResponseWriter, r *http.Request) {
	uri := strings.Split(r.RequestURI, "/")
	id, err := strconv.Atoi(uri[4])
	if err != nil {
		ResponseJSON(w, fmt.Sprintf("could not find id [%d] in URI", id), "fail", http.StatusBadRequest)
		return
	}

	var updatedPost api.BlogPost

	err = json.NewDecoder(r.Body).Decode(&updatedPost)
	if err != nil {
		ResponseJSON(w, "could not read post body", "fail", http.StatusBadRequest)
		return
	}

	updatedPost.UpdatedAt = time.Now()
	updatedPost.ID = id

	if err != nil {
		ResponseJSON(w, "could not update post", "fail", http.StatusBadRequest)
	}

	err = s.postService.Update(&updatedPost)
	if err != nil {
		ResponseJSON(w, "could not update post", "fail", http.StatusBadRequest)
		return
	}

	ResponseJSON(w, "post updated successfully", "success", http.StatusOK)
}

// DeletePost handles deleting one post from the database
func (s *Server) DeletePost(w http.ResponseWriter, r *http.Request) {
	uri := strings.Split(r.RequestURI, "/")
	id, err := strconv.Atoi(uri[4])
	if err != nil {
		ResponseJSON(w, fmt.Sprintf("could not find id [%d] in URI", id), "fail", http.StatusBadRequest)
		return
	}

	err = s.postService.Delete(id)
	if err != nil {
		ResponseJSON(w, fmt.Sprintf("could not delete post by id: %d", id), "fail", http.StatusBadRequest)
		return
	}

	ResponseJSON(w, "post deleted successfully", "success", http.StatusOK)
}

// CreateUser handles inserting a user into the database
func (s *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user api.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		ResponseJSON(w, "could not read post body", "fail", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		ResponseJSON(w, "could not hash password", "fail", http.StatusBadRequest)
		return
	}

	user.Password = string(hashedPassword)

	err = s.userService.New(user)

	if err != nil {
		ResponseJSON(w, "could not create user", "fail", http.StatusBadRequest)
		return
	}

	ResponseJSON(w, "user created successfully", "success", http.StatusCreated)
}

// GetAllUsers handles returning every user from the database
func (s *Server) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := s.userService.GetAll()

	if err != nil {
		ResponseJSON(w, "could not get all users", "fail", http.StatusBadRequest)
		return
	}

	ResponseJSON(w, users, "success", http.StatusOK)
}

// GetOneUser handles returning one user from the database
func (s *Server) GetOneUser(w http.ResponseWriter, r *http.Request) {
	uri := strings.Split(r.RequestURI, "/")
	id, err := strconv.Atoi(uri[4])
	if err != nil {
		ResponseJSON(w, fmt.Sprintf("could not find id [%d] in URI", id), "fail", http.StatusBadRequest)
		return
	}

	user, err := s.userService.GetOne(id)
	if err != nil {
		ResponseJSON(w, fmt.Sprintf("could not get user by id: %d", id), "fail", http.StatusBadRequest)
		return
	}

	ResponseJSON(w, user, "success", http.StatusOK)
}

// UpdateUser handles updating a user in the database
func (s *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {
	uri := strings.Split(r.RequestURI, "/")
	id, err := strconv.Atoi(uri[4])
	if err != nil {
		ResponseJSON(w, fmt.Sprintf("could not find id [%d] in URI", id), "fail", http.StatusBadRequest)
		return
	}

	var user api.User

	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		ResponseJSON(w, "could not read post body", "fail", http.StatusBadRequest)
		return
	}

	userFromDB, err := s.userService.GetOne(id)
	if err != nil {
		ResponseJSON(w, "could not get user from database", "fail", http.StatusBadRequest)
		return
	}

	userFromDB.Email = user.Email
	userFromDB.Username = user.Username
	userFromDB.FirstName = user.FirstName
	userFromDB.LastName = user.LastName

	if err != nil {
		ResponseJSON(w, "could not update post", "fail", http.StatusBadRequest)
	}

	err = s.userService.Update(userFromDB)
	if err != nil {
		ResponseJSON(w, "could not update user", "fail", http.StatusBadRequest)
		return
	}

	ResponseJSON(w, "user updated successfully", "success", http.StatusOK)
}

func (s *Server) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	uri := strings.Split(r.RequestURI, "/")
	id, err := strconv.Atoi(uri[4])
	if err != nil {
		ResponseJSON(w, fmt.Sprintf("could not find id [%d] in URI", id), "fail", http.StatusBadRequest)
		return
	}

	var user api.User

	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		ResponseJSON(w, "could not read post body", "fail", http.StatusBadRequest)
		return
	}

	userFromDB, err := s.userService.GetOne(id)
	if err != nil {
		ResponseJSON(w, "could not get user from database", "fail", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		ResponseJSON(w, "could not hash password", "fail", http.StatusBadRequest)
		return
	}

	if userFromDB.Password != string(hashedPassword) {
		userFromDB.Password = string(hashedPassword)
	} else {
		ResponseJSON(w, "new password cannot be the same as the old password", "fail", http.StatusBadRequest)
		return
	}

	err = s.userService.Update(userFromDB)
	if err != nil {
		ResponseJSON(w, "could not update user", "fail", http.StatusBadRequest)
		return
	}

	ResponseJSON(w, "password updated successfully", "success", http.StatusOK)
}

func (s *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {
	uri := strings.Split(r.RequestURI, "/")
	id, err := strconv.Atoi(uri[4])
	if err != nil {
		ResponseJSON(w, fmt.Sprintf("could not find id [%d] in URI", id), "fail", http.StatusBadRequest)
		return
	}

	err = s.userService.Delete(id)
	if err != nil {
		ResponseJSON(w, fmt.Sprintf("could not delete user by id: %d", id), "fail", http.StatusBadRequest)
		return
	}

	ResponseJSON(w, "user deleted successfully", "success", http.StatusOK)
}
