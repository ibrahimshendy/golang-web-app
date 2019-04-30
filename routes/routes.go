package routes

import (
	"fmt"
	"net/http"
	"site/middleware"
	"site/models"
	"site/sessions"
	"site/utils"

	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", middleware.AuthMiddleware(hom)).Methods("GET")
	r.HandleFunc("/home", middleware.AuthMiddleware(hom)).Methods("GET")
	r.HandleFunc("/login", login).Methods("GET")
	r.HandleFunc("/auth/login", postLogin).Methods("POST")
	r.HandleFunc("/register", register).Methods("GET")
	r.HandleFunc("/auth/register", postRegister).Methods("POST")
	r.HandleFunc("/about", abo).Methods("GET")
	r.HandleFunc("/comments", middleware.AuthMiddleware(comments)).Methods("GET")
	r.HandleFunc("/send-comment", middleware.AuthMiddleware(sendComment)).Methods("POST")
	r.PathPrefix("/stuff/").Handler(http.StripPrefix("/stuff", http.FileServer(http.Dir("assets"))))

	return r
}

func hom(w http.ResponseWriter, r *http.Request) {
	comments, err := models.AllComments()

	if err != nil {
		fmt.Println(err)
	}

	utils.LoadView(w, "index.gohtml", comments)
}

func abo(w http.ResponseWriter, r *http.Request) {
	utils.LoadView(w, "about.gohtml", nil)
}

func comments(w http.ResponseWriter, r *http.Request) {
	utils.LoadView(w, "comments.gohtml", nil)
}

func sendComment(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	comment := r.PostForm.Get("comment")
	models.Create(comment)
	http.Redirect(w, r, "/", 302)
}

func login(w http.ResponseWriter, r *http.Request) {
	utils.LoadView(w, "login.gohtml", nil)
}

func postLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")

	err := models.Login(username, password)
	if err != nil {
		switch err {
		case models.ErrUserNotFound:
			utils.LoadView(w, "login.gohtml", "User Not Found")
		case models.ErrInvalidAuth:
			utils.LoadView(w, "login.gohtml", "Invalid Authentication")
		default:
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
		}
		return
	}

	session, _ := sessions.Store.Get(r, "session")
	session.Values["username"] = username
	session.Save(r, w)
	http.Redirect(w, r, "/", 302)
}

func register(w http.ResponseWriter, r *http.Request) {
	utils.LoadView(w, "register.gohtml", nil)
}

func postRegister(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")

	err := models.Register(username, password)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	http.Redirect(w, r, "/login", 302)
}
