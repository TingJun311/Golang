package main

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	//"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
    Id int
    Name string
    Contact int
}


func main() {
    
    for {
        db, err := sql.Open("mysql", "forGoTest:z_Nb3EWC4hlLwhuO@tcp(localhost:3306)/mvc_golang")
        if err != nil {
            log.Fatal(err)
        }
        func(sb *sql.DB) {
            defer db.Close()
            //fmt.Println(time.Now())
            //WithoutConcurrent(db)
            //here := "appKey=eb794f63a7bc47b2a0aeb8491013b577&data=\"{\"shopId\":\"1100000001\",\"pageNo\":\"1\",\"pageSize\":\"10\"}\"&pampasCall=item.paging"
            secret := "8fea63348db42082a42fa89fa1461b66"

            param := map[string]interface{} {
                "appKey": "4ea6b9127b8e4e4d9757e6ccf5cac353",
                "data": "{\"email\":\"chewtingjun311@gmail.com\",\"eventId\":\"3002580074\",\"fullName\":\"Chew TingJun\",\"mobile\":\"60182092288\",\"senhengId\":\"3218350\"}",
                "pampasCall": "senHeng.member.voucher",
            }
            param2 := map[string]interface{} {
                "appKey": "4ea6b9127b8e4e4d9757e6ccf5cac353",
                "data": "{\"fullName\":\"NUR SHAFIKA IZANI\", \"mobile\":\"60145470835\", \"email\":\"izanieyshafika@gmail.com\", \"senhengId\":\"3400598\", \"eventId\":\"3000600012\"}",
                "pampasCall": "senHeng.member.voucher",
            }
            fmt.Println(param["data"])
            fmt.Println("SIGN HERE:::: ", signatureAlgo(secret, param2))

           // var currencyData map[string]interface{}
           // res, err := ConvertCurrency("USD", "MYR", "50000.7")
           // if err != nil {
           //     fmt.Println(err)
           // }

           // if err := json.Unmarshal(res, &currencyData); err != nil {
           //     fmt.Println(err)
           // }
           // fmt.Println("Before convert: ", currencyData["result"])
           // fmt.Println("Coverted result here: ", fmt.Sprintf("%4.2f", currencyData["result"].(float64)))

            //s := "627231680168310411,PG-SRCRM40,hi_PY_you_miss_me_?,2023-03-01,2023-06-01,Campaign_test"
          //  s2 := "627231680168310411,PG-SRCRM10,hi PY you miss me ?? v2,2023-03-01,2023-06-01,Campaign_test"
          //  arr := strings.Split(s2, ",")
          //  for _, i := range arr {
          //      fmt.Println(i)
          //  }
            //Concurrent(db)
            //fmt.Println(time.Now())
            //InsertRandom(db)
        }(db)
        time.Sleep(5 * time.Second)
    }
}

func ConvertCurrency(from, to, amount string) ([]byte, error) {
    url := fmt.Sprintf("https://api.apilayer.com/exchangerates_data/convert?to=%s&from=%s&amount=%s", to, from, amount)

    client := &http.Client {}
    req, err := http.NewRequest("GET", url, nil)
    req.Header.Set("apikey", "RueopsViA0NjnAHLfxmfzj1kwdvXlKBc")
    if err != nil {
        return nil, err
    }

    res, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer res.Body.Close()

    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        return nil, err
    }
    return body, nil
}

func ReadLocalFile() (string, error) {
    // Open the file for reading
    cwd, err := os.Getwd()
    if err != nil {
        fmt.Println("Error getting current directory:", err)
        return "", err
    }
    fmt.Println(cwd)
    file, err := os.Open("./file.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return "", err
    }
    defer file.Close()

    // Read the file into a byte slice
    data := make([]byte, 1024)
    count, err := file.Read(data)
    if err != nil {
        fmt.Println("Error reading file:", err)
        return "", err
    }

    // Print the contents of the file
    return string(data[:count]), nil
}

func WithoutConcurrent(db *sql.DB) {
    
    data := getUserV2(db)
    fmt.Print("data detc: ", len(data))
    for _, i := range data {
        v := fibonacci(i.Contact)
        fmt.Println("Result: ", v)
    }
}

func Concurrent(db *sql.DB) {
    res := make(chan User)

    go getAllUser(res, db)
    for each := range res {
        ids := make(chan int)
        fmt.Println("Channel receive: ", each)
        go func(n int, c chan int) {
            data := fibonacci(n)
            fmt.Println("Fibonacci: ", data)
            c <- data
        }(each.Contact, ids)

    //    go func(id int, r chan int) {
    //        select {
    //        case data := <-r:
    //            UpdateCustomer(db, id)
    //            fmt.Println(data)
    //        default:
    //            return
    //        }
    //    }(each.Id, ids)
    }
}


func getAllUser(res chan User, db *sql.DB) {
    
    query := ` SELECT 
                id,
                name,
                contact
               FROM goroutine_get 
               ORDER BY id DESC LIMIT 5000`


    cusInfo, err := db.Query(query)
	if err != nil {
		fmt.Println("[ERROR] getCustomerInfo -> Query")
        close(res)
		return
	}
    for cusInfo.Next() {
        var info User
        err := cusInfo.Scan(&info.Id, &info.Name, &info.Contact)
        if err != nil {
            fmt.Println("ERROR")
        }
        res <- info
    }
    close(res)
}

func fibonacci(n int) int {

    if n <= 1 {
        return n
    }

    // Create a slice to store previously computed Fibonacci numbers
    fib := make([]int, n + 1)
    fib[0] = 0
    fib[1] = 1

    for i := 2; i <= n; i++ {
        fib[i] = fib[i - 1] + fib[i - 2]
    }

    return fib[n]
}

func getUserV2(db *sql.DB) ([]User) {
    var data []User
    query := ` SELECT 
                    name,
                    contact
                FROM goroutine_get 
                ORDER BY id DESC LIMIT 5000`


    cusInfo, err := db.Query(query)
    if err != nil {
        fmt.Println("[ERROR] getCustomerInfo -> Query")
        return nil
    }
    for cusInfo.Next() {
        var info User
        err := cusInfo.Scan(&info.Name, &info.Contact)
        if err != nil {
            fmt.Println("ERROR")
        }
        data = append(data, info)
    }
    return data
}

func UpdateCustomer(db *sql.DB, id int) {

    query := ` UPDATE goroutine_get
                SET updated = NOW(), name = ?
                WHERE id = ?`

    stmt, err := db.Prepare(query)
    if err != nil {
        fmt.Println("		[*] Error Exec Prepare query -> UpdateCustomer/UpdateCustomer", err)
    }
    defer stmt.Close()
        
    insertRes, err := stmt.Exec("TingJuns", id)
    if err != nil {
        fmt.Println("       [*] ERROR EXEC", err)
    }
    isInsert, err := insertRes.RowsAffected()
    if err != nil {
        fmt.Println("       [*] ERROR last id")
    }
    if isInsert > 0 {
        fmt.Println("Updated. ")
    }
}

func InsertRandom(db *sql.DB) {
    
    query := ` INSERT INTO goroutine_get (name, contact, created, updated)
    VALUES (?, ?, NOW(), '0000-00-00 00:00:00')`


    for i := 0; i < 20000; i++ {
        stmt, err := db.Prepare(query)
        if err != nil {
            fmt.Println("		[*] Error Exec Prepare query -> merchant_share_cronjob/InsertSGame", err)
        }
        defer stmt.Close()

        insertRes, err := stmt.Exec("TingJun", fmt.Sprintf("%d", i))
        if err != nil {
            fmt.Println("       [*] ERROR EXEC", err)
        }
        isInsert, err := insertRes.LastInsertId()
        if err != nil {
            fmt.Println("       [*] ERROR last id")
        }
        fmt.Println(isInsert)
    }
}


func sign(secret string, params map[string]interface{}) string {
    // Step 1: Create the toVerify string by joining the params map
    // using & and = as separators
    var b strings.Builder
    for k, v := range params {
        if b.Len() > 0 {
            b.WriteByte('&')
        }
        b.WriteString(k)
        b.WriteByte('=')
        b.WriteString(fmt.Sprint(v))
    }
    toVerify := b.String()

    // Step 2: Compute the MD5 hash of the toVerify string concatenated with the secret
    h := md5.New()
    h.Write([]byte(toVerify))
    h.Write([]byte(secret))
    signature := hex.EncodeToString(h.Sum(nil))

    return signature
}

func signatureAlgo(secret string, params map[string]interface{}) string {
    // Step 1: Sort the keys of the params map alphabetically
    var keys []string
    for k := range params {
        keys = append(keys, k)
    }
    sort.Strings(keys)

    // Step 2: Create the toVerify string by joining the sorted params map
    // using & and = as separators
    var b strings.Builder
    for i, k := range keys {
        if i > 0 {
            b.WriteByte('&')
        }
        b.WriteString(k)
        b.WriteByte('=')
        b.WriteString(fmt.Sprint(params[k]))
    }
    toVerify := b.String()

    // Step 3: Compute the MD5 hash of the toVerify string concatenated with the secret
    h := md5.New()
    h.Write([]byte(toVerify))
    h.Write([]byte(secret))
    signature := hex.EncodeToString(h.Sum(nil))

    return signature
}