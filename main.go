package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"text/template"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Privilege string `json:"privilege"`
}

type JWTClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

type Article struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Thumbnail   string `json:"thumbnail"`
	Image       string `json:"image"`
	CreatedAt   string `json:"created_at"`
	IsPublished *bool  `json:"is_published"`
}

type Feedback struct {
	Sender  string `json:"sender"`
	Subject string `json:"subject"`
	Text    string `json:"text"`
}

var signingKey, encryptionKey []byte
var DBName, DBDriver, DBUsername, DBPassword string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	signingKey = []byte(os.Getenv("JWT_SECRET_KEY"))
	encryptionKey = []byte(os.Getenv("ENC_KEY"))
	DBName = os.Getenv("DB_NAME")
	DBDriver = os.Getenv("DB_DRIVER")
	DBUsername = os.Getenv("DB_USERNAME")
	DBPassword = os.Getenv("DB_PASSWORD")
}

func GenerateJWT(user *User) (string, error) {
	claims := JWTClaims{
		user.Name,
		jwt.StandardClaims{
			Id:        user.ID,
			ExpiresAt: time.Now().Add(10 * time.Minute).Unix(),
		},
	}

	stringToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := stringToken.SignedString(signingKey)
	if err != nil {
		return "", err
	}

	return token, nil
}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	message := make(map[string]interface{})

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return signingKey, nil
			})

			if err != nil {
				log.Println("Not authorized")
				w.Header().Set("Content-Type", "application/json")
				message["success"] = false
				message["message"] = err.Error()

				json.NewEncoder(w).Encode(&message)
			}

			if token.Valid {
				endpoint(w, r)
			}
		} else {
			log.Println("Not authorized")
			w.Header().Set("Content-Type", "application/json")
			message["success"] = false
			message["message"] = "Not authorized"

			json.NewEncoder(w).Encode(&message)
		}
	})
}

func signOut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	message := make(map[string]interface{})
	message["success"] = true
	message["message"] = "Your account has been signed out successfully"

	json.NewEncoder(w).Encode(&message)
}

func signIn(w http.ResponseWriter, r *http.Request) {
	var user User
	var dbPassword []byte
	message := make(map[string]interface{})

	w.Header().Set("Content-Type", "application/json")

	_ = json.NewDecoder(r.Body).Decode(&user)

	db, err := connectDB()
	checkErr(w, err)
	defer db.Close()

	row, err := db.Query("SELECT id, name, email, password, privilege FROM `user` WHERE email = ? LIMIT 0, 1", user.Email)
	checkErr(w, err)

	for row.Next() {
		err = row.Scan(&user.ID, &user.Name, &user.Email, &dbPassword, &user.Privilege)
		checkErr(w, err)
	}

	if err := bcrypt.CompareHashAndPassword(dbPassword, []byte(user.Password)); err != nil {
		message["success"] = true
		message["message"] = "Wrong username / password"
	} else {
		token, err := GenerateJWT(&user)
		if err != nil {
			message["success"] = true
			message["message"] = "There was an error while generating JWT"
		} else {
			message["success"] = true
			message["message"] = "Authenticated Successfully"
			message["jwt"] = token
			message["email"] = user.Email
			message["role"] = user.Privilege
		}
	}

	json.NewEncoder(w).Encode(&message)
}

func postFeedback(w http.ResponseWriter, r *http.Request) {
	var feedback Feedback
	message := make(map[string]interface{})

	w.Header().Set("Content-Type", "application/json")

	_ = json.NewDecoder(r.Body).Decode(&feedback)

	if len(feedback.Sender) == 0 || len(feedback.Subject) == 0 || len(feedback.Text) == 0 {
		message["success"] = false
		message["message"] = "Email or Subject or Text must be filled"

		json.NewEncoder(w).Encode(&message)
		return
	}

	if m, _ := regexp.MatchString(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`, feedback.Sender); !m {
		message["success"] = false
		message["message"] = "Email format does not match"

		json.NewEncoder(w).Encode(&message)
		return
	}

	db, err := connectDB()
	checkErr(w, err)
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO `feedback` VALUES (null, ?, ?, ?)")
	checkErr(w, err)

	_, err = stmt.Exec(feedback.Sender, feedback.Subject, feedback.Text)
	checkErr(w, err)

	message["success"] = true
	message["message"] = "Your feedback has been sent"

	json.NewEncoder(w).Encode(&message)
}

func getArticles(w http.ResponseWriter, r *http.Request) {
	message := make(map[string]interface{})
	w.Header().Set("Content-Type", "application/json")

	db, err := connectDB()
	checkErr(w, err)
	defer db.Close()

	rows, err := db.Query("SELECT id, title, content, thumbnail, image, created_at, is_published FROM `article` WHERE deleted_at IS NULL ORDER BY created_at DESC")
	checkErr(w, err)

	var articles []Article

	for rows.Next() {
		var id, title, content, thumbnail, image, createdAt string
		var isPublished *bool
		err = rows.Scan(&id, &title, &content, &thumbnail, &image, &createdAt, &isPublished)
		checkErr(w, err)

		articles = append(articles, Article{id, title, content, thumbnail, image, createdAt, isPublished})
	}

	message["success"] = true
	message["message"] = "Success fetching data"
	message["data"] = articles

	json.NewEncoder(w).Encode(&message)
}

func getPublishedArticles(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
	w.Header().Set("Content-Type", "application/json")
	message := make(map[string]interface{})

	db, err := connectDB()
	checkErr(w, err)
	defer db.Close()

	rows, err := db.Query("SELECT id, title, content, thumbnail, image, created_at, is_published FROM `article` WHERE deleted_at IS NULL AND is_published = 1 ORDER BY created_at DESC")
	checkErr(w, err)

	var articles []Article

	for rows.Next() {
		var id, title, content, thumbnail, image, createdAt string
		var isPublished *bool
		err = rows.Scan(&id, &title, &content, &thumbnail, &image, &createdAt, &isPublished)
		checkErr(w, err)

		articles = append(articles, Article{id, title, content, thumbnail, image, createdAt, isPublished})
	}

	message["success"] = true
	message["message"] = "Success fetching data"
	message["data"] = articles

	json.NewEncoder(w).Encode(&message)
}

func getFeedbacks(w http.ResponseWriter, r *http.Request) {
	var feedbacks []Feedback
	message := make(map[string]interface{})
	w.Header().Set("Content-Type", "application/json")

	db, err := connectDB()
	checkErr(w, err)
	defer db.Close()

	rows, err := db.Query("SELECT sender, subject, text FROM `feedback` WHERE deleted_at IS NULL ORDER BY created_at DESC")
	checkErr(w, err)

	for rows.Next() {
		var sender, subject, text string
		err = rows.Scan(&sender, &subject, &text)
		checkErr(w, err)

		feedbacks = append(feedbacks, Feedback{sender, subject, text})
	}

	message["success"] = true
	message["message"] = "Success fetching data"
	message["data"] = feedbacks

	json.NewEncoder(w).Encode(&message)
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	message := make(map[string]interface{})
	w.Header().Set("Content-Type", "application/json")

	db, err := connectDB()
	checkErr(w, err)
	defer db.Close()

	rows, err := db.Query("SELECT id, title, content, thumbnail, image, created_at, is_published FROM `article` WHERE id = ? AND deleted_at IS NULL", id)
	checkErr(w, err)

	var articles []Article

	for rows.Next() {
		var id, title, content, thumbnail, image, createdAt string
		var isPublished *bool
		err = rows.Scan(&id, &title, &content, &thumbnail, &image, &createdAt, &isPublished)
		checkErr(w, err)

		articles = append(articles, Article{id, title, content, thumbnail, image, createdAt, isPublished})
	}

	message["success"] = true
	message["message"] = "Success fetching data"
	message["data"] = articles

	json.NewEncoder(w).Encode(&message)
}

func postArticle(w http.ResponseWriter, r *http.Request) {
	var article Article
	w.Header().Set("Content-Type", "application/json")
	message := make(map[string]interface{})

	_ = json.NewDecoder(r.Body).Decode(&article)

	if len(article.Title) == 0 || len(article.Content) == 0 || len(article.Thumbnail) == 0 || len(article.Image) == 0 {
		message["success"] = true
		message["message"] = "Title, Content, Thumbnail and Image must be filled"

		json.NewEncoder(w).Encode(&message)
		return
	}

	db, err := connectDB()
	checkErr(w, err)
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO `article` VALUES (null, ?, ?, ?, ?, null, null, null, ?)")
	checkErr(w, err)

	_, err = stmt.Exec(article.Title, article.Content, article.Thumbnail, article.Image, false)
	checkErr(w, err)

	message["success"] = true
	message["message"] = "An article has been created successfully"
	json.NewEncoder(w).Encode(&message)
}

func putArticle(w http.ResponseWriter, r *http.Request) {
	var article Article
	message := make(map[string]interface{})
	params := mux.Vars(r)
	id := params["id"]
	t := time.Now()

	w.Header().Set("Content-Type", "application/json")

	_ = json.NewDecoder(r.Body).Decode(&article)

	if len(article.Title) == 0 || len(article.Content) == 0 || len(article.Thumbnail) == 0 || len(article.Image) == 0 {
		message["success"] = false
		message["message"] = "Title, Content, Thumbnail and Image must be filled"

		json.NewEncoder(w).Encode(&message)
		return
	}

	db, err := connectDB()
	checkErr(w, err)
	defer db.Close()

	stmt, err := db.Prepare("UPDATE `article` SET title = ?, content = ?, thumbnail = ?, image = ?, updated_at = ? WHERE id = ?")
	checkErr(w, err)

	_, err = stmt.Exec(article.Title, article.Content, article.Thumbnail, article.Image, t.Format("2006-01-02 15:04:05"), id)
	checkErr(w, err)

	message["success"] = true
	message["message"] = "Selected article has been updated successfully"
	json.NewEncoder(w).Encode(&message)
}

func patchArticle(w http.ResponseWriter, r *http.Request) {
	var article Article
	params := mux.Vars(r)
	id := params["id"]
	t := time.Now()

	message := make(map[string]interface{})
	w.Header().Set("Content-Type", "application/json")

	_ = json.NewDecoder(r.Body).Decode(&article)

	if article.IsPublished == nil {
		message["success"] = false
		message["message"] = "Is published must be filled in"

		json.NewEncoder(w).Encode(&message)
		return
	}

	db, err := connectDB()
	checkErr(w, err)
	defer db.Close()

	stmt, err := db.Prepare("UPDATE `article` SET is_published = ?, updated_at = ? WHERE id = ?")
	checkErr(w, err)

	_, err = stmt.Exec(article.IsPublished, t.Format("2006-01-02 15:04:05"), id)
	checkErr(w, err)

	message["success"] = true
	message["message"] = "Selected article has been updated successfully"
	json.NewEncoder(w).Encode(&message)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	t := time.Now()

	message := make(map[string]interface{})
	w.Header().Set("Content-Type", "application/json")

	db, err := connectDB()
	checkErr(w, err)
	defer db.Close()

	stmt, err := db.Prepare("UPDATE `article` SET deleted_at = ? WHERE id = ?")
	checkErr(w, err)

	_, err = stmt.Exec(t.Format("2006-01-02 15:04:05"), id)
	checkErr(w, err)

	message["success"] = true
	message["message"] = "Selected article has been deleted successfully"
	json.NewEncoder(w).Encode(&message)
}

func pageIndex(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles(
		"views/index.html",
		"views/_header.html",
		"views/_navigation.html",
		"views/_footer.html",
	))

	err := tmpl.ExecuteTemplate(w, "index", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func pageAbout(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles(
		"views/about.html",
		"views/_header.html",
		"views/_navigation.html",
		"views/_footer.html",
	))

	err := tmpl.ExecuteTemplate(w, "about", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func pageContact(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles(
		"views/contact.html",
		"views/_header.html",
		"views/_navigation.html",
		"views/_footer.html",
	))

	err := tmpl.ExecuteTemplate(w, "contact", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func pageSignIn(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles(
		"views/signin.html",
		"views/_header.html",
		"views/_navigation.html",
		"views/_footer.html",
	))

	err := tmpl.ExecuteTemplate(w, "signin", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func pageDetail(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	db, err := connectDB()
	checkErr(w, err)
	defer db.Close()

	rows, err := db.Query("SELECT id, title, content, thumbnail, image, created_at, is_published FROM `article` WHERE id = ? AND deleted_at IS NULL", id)
	checkErr(w, err)

	var article Article

	for rows.Next() {
		var id, title, content, thumbnail, image, created_at string
		var is_published *bool
		err = rows.Scan(&id, &title, &content, &thumbnail, &image, &created_at, &is_published)
		checkErr(w, err)

		article = Article{id, title, content, thumbnail, image, created_at, is_published}
	}

	var tmpl = template.Must(template.ParseFiles(
		"views/detail.html",
		"views/_header.html",
		"views/_navigation.html",
		"views/_footer.html",
	))

	err = tmpl.ExecuteTemplate(w, "detail", article)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func pageFeedback(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles(
		"views/see-feedbacks.html",
		"views/_header.html",
		"views/_navigation.html",
		"views/_footer.html",
	))

	err := tmpl.ExecuteTemplate(w, "see-feedbacks", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func pageArticles(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles(
		"views/manage-articles.html",
		"views/_header.html",
		"views/_navigation.html",
		"views/_footer.html",
	))

	err := tmpl.ExecuteTemplate(w, "manage-articles", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/feedback", postFeedback).Methods("POST")
	router.HandleFunc("/feedbacks", getFeedbacks).Methods("GET")
	router.HandleFunc("/articles", getArticles).Methods("GET")
	router.HandleFunc("/published-articles", getPublishedArticles).Methods("GET")
	router.HandleFunc("/article/{id}", getArticle).Methods("GET")
	router.Handle("/article", isAuthorized(postArticle)).Methods("POST")
	router.Handle("/article/{id}", isAuthorized(putArticle)).Methods("PUT")
	router.Handle("/article/{id}", isAuthorized(patchArticle)).Methods("PATCH") //publish
	router.Handle("/article/{id}", isAuthorized(deleteArticle)).Methods("DELETE")
	router.HandleFunc("/sign-in", signIn).Methods("POST")
	router.Handle("/sign-out", isAuthorized(signOut)).Methods("POST")

	/* Views */
	router.HandleFunc("/", pageIndex).Methods("GET")
	router.HandleFunc("/detail/{id}", pageDetail).Methods("GET")
	router.HandleFunc("/about", pageAbout).Methods("GET")
	router.HandleFunc("/contact", pageContact).Methods("GET")
	router.HandleFunc("/manage-articles", pageArticles).Methods("GET")
	router.HandleFunc("/see-feedbacks", pageFeedback).Methods("GET")
	router.HandleFunc("/sign-in", pageSignIn).Methods("GET")
	router.HandleFunc("/sign-in", pageSignIn).Methods("GET")

	/* Assets */
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./assets/")))

	fmt.Println("Web Service Started")

	err := http.ListenAndServe(":9090", router)
	if err != nil {
		log.Fatal("Error running service: ", err)
	}
}

func checkErr(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}

func connectDB() (*sql.DB, error) {
	db, err := sql.Open(DBDriver, DBUsername+":"+DBPassword+"@tcp(database-1.chghgvrwjatk.ap-southeast-1.rds.amazonaws.com)/"+DBName+"?charset=utf8")
	if err != nil {
		return nil, err
	}

	return db, nil
}
