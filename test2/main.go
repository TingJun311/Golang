package main

import (
	"crypto/sha1"
	"encoding/base32"
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
    key   string
    value string
    next  *Node
}

type HashTable struct {
    table []*Node
    size  int
}

func hash(key string, size int) int {
    var hashVal int
    for i := 0; i < len(key); i++ {
        hashVal = 31 * hashVal + int(key[i])
    }
    return hashVal % size
}

func (ht *HashTable) Put(key string, value string) {
    index := hash(key, ht.size)
    newNode := &Node{key, value, nil}
    if ht.table[index] == nil {
        ht.table[index] = newNode
    } else {
        curr := ht.table[index]
        for curr.next != nil {
            if curr.key == key {
                curr.value = value
                return
            }
            curr = curr.next
        }
        if curr.key == key {
            curr.value = value
        } else {
            curr.next = newNode
        }
    }
}

func (ht *HashTable) Get(key string) (string, bool) {
    index := hash(key, ht.size)
    curr := ht.table[index]
    for curr != nil {
        if curr.key == key {
            return curr.value, true
        }
        curr = curr.next
    }
    return "", false
}

func (ht *HashTable) Contains(key string) bool {
    index := hash(key, ht.size)
    curr := ht.table[index]
    for curr != nil {
        if curr.key == key {
            return true
        }
        curr = curr.next
    }
    return false
}

func main() {
	var test []string
	collusion := 0
	now := time.Now()
	for i := 0; i < 10; i++ {
		newBarcode, _ := GenerateUniqueBarcode(10)
		previousNum := collusion
		for j := 0; j < len(test); j++ {
			if newBarcode == test[j] {
				fmt.Println("SAME ", newBarcode, test[j])
				collusion++
			}	
		}
		if previousNum == collusion {
			test = append(test, newBarcode)
			fmt.Println("A unqiue barcode appending... ", newBarcode, i)
		} else {
			fmt.Println("A collusion found: ", newBarcode)
		}
	//	time.Sleep(time.Second * 1)
	}
	fmt.Println(now)	
	fmt.Println("Total collusion: ", collusion)
	fmt.Println("Time: ", time.Now())
	//fmt.Println(GenerateUniqueBarcode(10))
	//fmt.Println(GenerateRandomString(10))
	//fmt.Println(GenerateUniqueHash())


	const size = 10
	ht := HashTable{make([]*Node, size), size}
	newBarcode, _ := GenerateUniqueBarcode(10)
	for i := 0; i < size; i++ {
		if _, ok := ht.Get(newBarcode); !ok {
			fmt.Println("Appending a new barcode", newBarcode, i)
			ht.Put(newBarcode, newBarcode)
		} else {
			fmt.Println("Collusion found! ", newBarcode)
			collusion++
		}
	}
    //fmt.Println(ht)
	fmt.Println("Total colluison", collusion)
}


func GenerateUniqueBarcode(length int) (string, error) {

	var Barcode []rune
	alphanumeric := []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	
	
	for i := 0; i < length; i++ {
		var convertedNum int
		nums := func (maxNum int) (int) {
			if maxNum < 0 {
				return 0
			}
			rand.Seed(time.Now().UnixNano())
			return rand.Intn(maxNum + 1)
		}  (time.Now().Nanosecond())
			rand.Seed(time.Now().Unix())
		  
			// Shuffling the string
			rand.Shuffle(len(alphanumeric), func(i, j int) {
				alphanumeric[i], alphanumeric[j] = alphanumeric[j], alphanumeric[i]
			})
			if nums < 0 {
				convertedNum = nums * -1
			} else {
				convertedNum = nums
			}
			Barcode = append(Barcode, alphanumeric[(convertedNum * i * time.Now().Day() + time.Now().Nanosecond()) % len(alphanumeric)])
		}
	return string(Barcode), nil
}

func GenerateRandomString(n int) (string, error) {
    // define the character set for the random string
    const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

    // seed the random number generator
    rand.Seed(time.Now().UnixNano())

    // create a byte slice of the specified length
    b := make([]byte, n)

    // fill the byte slice with random characters
    for i := range b {
        b[i] = charset[rand.Intn(len(charset))]
    }

    // encode the byte slice using base32 encoding
    encoded := base32.StdEncoding.EncodeToString(b)

    // return the first n characters of the encoded string
    return encoded[:n], nil
}

func GenerateUniqueHash() (string, error) {
    // generate a random 8-byte array using the crypto/rand package
    randomBytes := make([]byte, 8)
    if _, err := rand.Read(randomBytes); err != nil {
        return "", err
    }

    // create a SHA1 hash of the random bytes
    hash := sha1.Sum(randomBytes)

    // encode the hash using base32 encoding
    encoded := base32.StdEncoding.EncodeToString(hash[:])

    // return the first 10 characters of the encoded string
    return encoded[:10], nil
}

