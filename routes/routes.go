package routes

import (
	"log"
	"marcfinserv/utils"
	"net/http"
	"net/smtp"
	"strings"

	"github.com/gorilla/mux"
)

func Init() *mux.Router {

	r := mux.NewRouter()

	return r
}

func Handle(r *mux.Router) {
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/index", indexHandler)
	r.HandleFunc("/products", productsHandler)
	r.HandleFunc("/about", aboutHandler)
	r.HandleFunc("/contact", contactHandler)
	r.HandleFunc("/partner", partnerHandler)
	r.HandleFunc("/submit", submitHandler).Methods("POST")
	r.HandleFunc("/send", sendHandler)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	log.Println(r)
	log.Println("Index page get request")
	p := "vishal"
	utils.ExecuteTemplate(w, "index.html", p)
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	log.Println(r)
	log.Println("products page get request")
	p := "vishal"
	utils.ExecuteTemplate(w, "products.html", p)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	log.Println(r)
	log.Println("Index page get request")
	p := "vishal"
	utils.ExecuteTemplate(w, "about-us.html", p)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	log.Println(r)
	log.Println("Index page get request")
	p := "vishal"
	utils.ExecuteTemplate(w, "contact.html", p)
}

func partnerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	log.Println(r)
	log.Println("Index page get request")
	p := "vishal"
	utils.ExecuteTemplate(w, "partner.html", p)
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("submit request received")
	log.Println(r)
	r.ParseForm()
	body := "Hi you received an enquiry from "
	for k, v := range r.Form {
		log.Println("key ", k)
		log.Println("value", strings.Join(v, ""))
		body = body + " " + k + " " + strings.Join(v, "")
	}
	sendEmail(body)
}

func sendHandler(w http.ResponseWriter, r *http.Request) {
	sendEmail("test")
}

type smtpServer struct {
	host string
	port string
}

// Address URI to smtp server
func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}

func sendEmail(body string) {
	log.Println("sending email")
	from := "marcfinserv@gmail.com"
	password := "9314102495"

	to := []string{
		"marcfinserv@gmail.com",
	}

	smtpServer := smtpServer{host: "smtp.gmail.com", port: "587"}

	message := []byte("Subject: Customer Enquiry Form!\r\n" + body)
	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpServer.host)
	// Sending email.
	err := smtp.SendMail(smtpServer.Address(), auth, from, to, message)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Email Sent!")
}
