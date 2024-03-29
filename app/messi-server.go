package main

import (
	"fmt"
	"log"
	"net/http"
	// toml "github.com/pelletier/go-toml"
	// "github.com/gorilla/mux"
	// "github.com/BurntSushi/toml"
	"encoding/json"
	"github.com/aws/aws-lambda-go/lambda"
	// "github.com/aws/aws-lambda-go/events"

	"os"
	"bytes"
	"context"
)

const (
	fbAPI = "https://graph.facebook.com/v2.6/me/messages?access_token=%s"
	image = "http://37.media.tumblr.com/e705e901302b5925ffb2bcf3cacb5bcd/tumblr_n6vxziSQD11slv6upo3_500.gif"
)

// Callback from 
type Callback struct {
	Object string `json:"object,omitempty"`
	Entry  []struct {
		ID        string      `json:"id,omitempty"`
		Time      int         `json:"time,omitempty"`
		Messaging []Messaging `json:"messaging,omitempty"`
	} `json:"entry,omitempty"`
}

// Messaging from
type Messaging struct {
	Sender    User    `json:"sender,omitempty"`
	Recipient User    `json:"recipient,omitempty"`
	Timestamp int     `json:"timestamp,omitempty"`
	Message   Message `json:"message,omitempty"`
}

// User from
type User struct {
	ID string `json:"id,omitempty"`
}

// Message from
type Message struct {
	MID        string `json:"mid,omitempty"`
	Text       string `json:"text,omitempty"`
	QuickReply *struct {
		Payload string `json:"payload,omitempty"`
	} `json:"quick_reply,omitempty"`
	Attachments *[]Attachment `json:"attachments,omitempty"`
	Attachment  *Attachment   `json:"attachment,omitempty"`
}

// Attachment from
type Attachment struct {
	Type    string  `json:"type,omitempty"`
	Payload Payload `json:"payload,omitempty"`
}

// Response from
type Response struct {
	Recipient User    `json:"recipient,omitempty"`
	Message   Message `json:"message,omitempty"`
}

// Payload from
type Payload struct {
	URL string `json:"url,omitempty"`
}

type tomlConfig struct {
	Title string
	TomlVars tomlVars 
}

type tomlVars struct {
	PageAccessToken string
	VerifyToken string
}

// TopIndex returns general message
func TopIndex(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello from messi-bot")
}

// MyEvent is
type MyEvent struct {
	Name string `json:"name"`
}

type verify struct{
	Challenge string `json:"key1"`
	Token string      `json:"key2"`
}

// HandleRequest takes in request
func HandleRequest(ctx context.Context, v map[string]interface{}) (string, error) {

for k,v := range(v){
	fmt.Printf("key/value: %s/%s\n", k, v)
}


	// fmt.Printf("Event: %T" , v)
// VerificationEndpoint(v)
	// var environ tomlConfig

// _, err := toml.DecodeFile("config.toml", &environ)
// if err != nil{
// 	log.Fatalf("Threw error %v\n", err)
// 	return
// }
return  "huh", nil
// r := mux.NewRouter()
// r.HandleFunc("/", TopIndex)
// if err := http.ListenAndServe(":8080", r); err != nil {
// 	log.Fatal(err)
// }
// 	return fmt.Sprintf("Hello %s!", 
// 	name.Name ), nil
}

func main() {
	lambda.Start(HandleRequest)
}

// Todo: make other GET and Post endpoints as necessary

// VerificationEndpoint : (GET /webhook)
// sent to bot, verifies the request is coming from Facebook.
func VerificationEndpoint( v verify) (verify){
	// challenge := v.Challenge
	//token := v.Token
	return v
}

// MessagesEndpoint : (POST /webhook)
func MessagesEndpoint(w http.ResponseWriter, r *http.Request) {
	var callback Callback
	json.NewDecoder(r.Body).Decode(&callback)
	if callback.Object == "page" {
		for _, entry := range callback.Entry {
			for _, event := range entry.Messaging {
				ProcessMessage(event)
			}
		}
		w.WriteHeader(200)
		w.Write([]byte("Got your message"))
	} else {
		w.WriteHeader(404)
		w.Write([]byte("Message not supported"))
	}
}

// ProcessMessage for
func ProcessMessage(event Messaging) {
	client := &http.Client{}
	response := Response{
		Recipient: User{
			ID: event.Sender.ID,
		},
		Message: Message{
			Attachment: &Attachment{
				Type: "image",
				Payload: Payload{
					URL: image,
				},
			},
		},
	}
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(&response)
	url := fmt.Sprintf(fbAPI, os.Getenv("PAGE_ACCESS_TOKEN"))
	req, err := http.NewRequest("POST", url, body)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
}