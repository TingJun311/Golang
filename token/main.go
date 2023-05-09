package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	mathRand "math/rand"
	"net/http"
	"pkg/cache"
	"pkg/token"
	"strings"
	"time"
	"golang.org/x/crypto/sha3"
)

// CustomTokenHandler implements the TokenHandler interface
type CustomTokenHandler struct{}

const APIKEY = "SnN5QCFc5nmiSJ+1lEuIBLiY6RC/ayRNhJ+cFgYooaA=XC68no+gX70uocilrdp1IG97U4=GFcgTRZZdYSh8cXgl"
func main() {
    cache := cache.NewCache()
    newAPIley := APIKeyGenerator()
    fmt.Print("NEW API KEY: ", newAPIley)

    http.HandleFunc("/", Test)
    http.HandleFunc("/gettoken", func(w http.ResponseWriter, r *http.Request) {
        GetToken(w, r, cache)
        for j, i := range cache.DisplayAll() {
            fmt.Println("DISPLAY cache: ", j, i)
        }
    })
    http.HandleFunc("/validate_token", func (w http.ResponseWriter, r *http.Request) {
        ValidateTokensFunc(w, r, cache)
        for j, i := range cache.DisplayAll() {
            fmt.Println("DISPLAY after validate cache: ", j, i)
        }
    })
    http.HandleFunc("/get_token_v2", func (w http.ResponseWriter, r *http.Request) {

        // handle token requests
        //srv.HandleTokenRequest(w, r)
    })

    http.HandleFunc("/webhook", handleWebhook)

    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {
    // Parse the incoming request to extract the data
    // For example, if the webhook sends JSON data:
    // var payload SomePayloadStruct
    // err := json.NewDecoder(r.Body).Decode(&payload)
    // if err != nil {
    //     http.Error(w, err.Error(), http.StatusBadRequest)
    //     return
    // }

    // Handle the webhook data according to your application logic
    // For example:
    // fmt.Println("Received payload:", payload)

    // Respond to the webhook with an appropriate HTTP status code
    w.WriteHeader(http.StatusOK)
}


func Test(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        res, _ := io.ReadAll(r.Body)
        fmt.Println("Body: ", string(res))
        fmt.Println("Header: ", r.Header)
        fmt.Println("Host: ", r.Host)
        fmt.Println("Post form value: ", r.PostFormValue("here"))
        fmt.Println("Link query param: ", r.URL.Query())
    } else {
        json.NewEncoder(w).Encode(map[string]interface{} {
            "code": 404,
            "status": "FAILED",
            "body": nil,
            "message": "request invalid.",
        })
    }
}

func GetToken(w http.ResponseWriter, r *http.Request, cache *cache.Cache) {
    if r.Method == http.MethodPost {
        apiKey := r.Header.Values("Apikey")[0]
        if apiKey != APIKEY {
            json.NewEncoder(w).Encode(map[string]interface{} {
                "code": 404,
                "status": "FAILED",
                "body": nil,
                "message": "Unauthorized user.",
            })
        }
        newToken := token.CreateToken(*cache)
        json.NewEncoder(w).Encode(map[string]interface{} {
            "code": 201,
            "status": "SUCCESS",
            "body": newToken.Value,
            "message": "token generated",
        })
    } else {
        json.NewEncoder(w).Encode(map[string]interface{} {
            "code": 404,
            "status": "FAILED",
            "body": nil,
            "message": "request invalid.",
        })
    }
}

func ValidateTokensFunc(w http.ResponseWriter, r *http.Request, cache *cache.Cache) {
    if r.Method == http.MethodPost {
        var newToken token.Token
        var doneUsed, ok bool
        var res interface{}
        apiKey := r.Header.Values("Authorization")
        tokens := strings.Split(apiKey[0], "Bearer ")
        ok = token.ValidateToken(tokens[1], *cache)
        if !ok {
            json.NewEncoder(w).Encode(map[string]interface{} {
                "code": 500,
                "status": "FAILED",
                "body": nil,
                "message": "Invalid token.",
                "isUsed": doneUsed,
                "ok": ok,
            })
            return
        }
        if res, ok = cache.Get(tokens[1]); !ok {
            log.Fatalln("HERE", ok)
        } 
        if tempByte, err := json.Marshal(res); err == nil {
            if err = json.Unmarshal(tempByte, &newToken); err != nil {
                log.Fatal(err)
            }
        }
        doneUsed = newToken.Used
        if doneUsed {
            json.NewEncoder(w).Encode(map[string]interface{} {
                "code": 500,
                "status": "FAILED",
                "body": nil,
                "message": "Invalid token.",
                "isUsed": doneUsed,
                "ok": ok,
            })
            return
        }
        newToken, _, ok = token.UseToken(&newToken, *cache)
        if !ok {
            log.Fatalln(ok)
            return 
        }
        json.NewEncoder(w).Encode(map[string]interface{} {
            "code": 200,
            "status": "SUCCESS",
            "body": nil,
            "message": "Using token",
            "token": newToken.Value,
            "expiry": newToken.Expiration,
        })
    } else {
        json.NewEncoder(w).Encode(map[string]interface{} {
            "code": 404,
            "status": "FAILED",
            "body": nil,
            "message": "request invalid.",
        })
    }
}


func APIKeyGenerator() (string) {
        // Generate 32 random bytes
        key := make([]byte, 32)
        _, err := rand.Read(key)
        if err != nil {
            fmt.Println("Error generating random bytes:", err)
            return ""
        }
    
        // Encode the random bytes using base64
        apiKey := base64.StdEncoding.EncodeToString(key)
    
        // Get the current Unix timestamp in seconds
        timestamp := time.Now().Unix()
    
        // Encode the timestamp as a string
        timestampStr := fmt.Sprintf("%d", timestamp)
        
        // Combine the API key and timestamp
        apiKeyWithTimestamp := fmt.Sprintf("%s:%s", apiKey, timestampStr)
        hash := sha3.Sum256([]byte(apiKeyWithTimestamp))
        hashStr := base64.StdEncoding.EncodeToString(hash[:])

        shuffleString := func (s string) string {
            // Convert the string to a slice of runes
            r := []rune(s)
        
            // Initialize a new source of random numbers
            src := mathRand.NewSource(time.Now().UnixNano())

            // Create a new random number generator using the source
            rgen := mathRand.New(src)
        
            // Use the Fisher-Yates algorithm to shuffle the slice
            for i := len(r) - 1; i > 0; i-- {
                j := rgen.Intn(i + 1)
                r[i], r[j] = r[j], r[i]
            }
            return string(r)
        } (apiKey)
        return hashStr + shuffleString
}
