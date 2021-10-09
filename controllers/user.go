package controllers

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/IamLucif3r/Insta-API/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Hash struct{}
type UserControler struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserControler {
	return &UserControler{s}
}

func (uc UserControler) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}
	oid := bson.ObjectIdHex(id)
	u := models.User{}

	if err := uc.session.DB("Insta-API").C("users").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserControler) GetPost(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}
	oid := bson.ObjectIdHex(id)
	u := models.Post{}

	if err := uc.session.DB("Insta-API").C("posts").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

// this is a createuser struct method
func (uc UserControler) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}
	// Decode values we got from postman, so that Golang can work on it
	json.NewDecoder(r.Body).Decode(&u)
	u.ID = bson.NewObjectId()

	h := sha1.New()
	h.Write([]byte(u.Password))
	sha := h.Sum(nil) // "sha" is uint8 type, encoded in base16

	shaStr := hex.EncodeToString(sha)

	uc.session.DB("Insta-API").C("users").Insert(u)
	uc.session.DB("Insta-API").C("users").Update((u.Password), (shaStr))
	uj, err := json.Marshal(u)

	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}
func (uc UserControler) CreatePost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.Post{}
	json.NewDecoder(r.Body).Decode(&u)
	u.ID = bson.NewObjectId()
	uc.session.DB("Insta-API").C("posts").Insert(u)
	uc.session.DB("Insta-API").C("posts").Update((u.TimeStamp), (time.Now()))
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserControler) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}
	oid := bson.ObjectIdHex(id)
	if err := uc.session.DB("Insta-API").C("users").RemoveId(oid); err != nil {
		w.WriteHeader(404)

	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted User", oid, "\n")
}
