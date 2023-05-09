package websocketfunc

import (
	"app/ds/trie"
	"encoding/json"
	"fmt"
	"log"
)

func SearchBar(data string) (res []byte) {
	newTrie := trie.NewTrie()

	var data2 string
	if err := Unmarshal([]byte(data), &data2); err != nil {
		fmt.Println("Error", err)
	}
	var brands = []string{
		"Nike",
		"Adidas",
		"Apple",
		"Samsung",
		"Microsoft",
		"HP",
		"Canon",
		"Sony",
		"LG",
		"Panasonic",
		"Gucci",
		"Prada",
		"Louis Vuitton",
		"Chanel",
		"Versace",
		"Hermes",
		"Burberry",
		"Coach",
		"Michael Kors",
		"Tiffany & Co.",
		"Rolex",
		"Omega",
		"Cartier",
		"Breitling",
		"Girard-Perregaux",
		"Piaget",
		"Montblanc",
		"Parker",
		"Cross",
		"Sharpie",
		"Crayola",
		"Lego",
		"Hasbro",
		"Mattel",
		"Fisher-Price",
		"Hot Wheels",
		"Barbie",
		"Nerf",
		"Play-Doh",
		"Monopoly",
		"Clarks",
		"Timberland",
		"Dr. Martens",
		"Vans",
		"Converse",
		"Under Armour",
		"Puma",
		"Reebok",
		"New Balance",
		"Skechers",
		"Crocs",
		"Fila",
		"Levi's",
		"Diesel",
		"Wrangler",
		"Calvin Klein",
		"Tommy Hilfiger",
		"Ralph Lauren",
		"H&M",
		"Zara",
		"Uniqlo",
		"GAP",
		"Forever 21",
		"American Eagle",
		"Abercrombie & Fitch",
		"Hollister",
		"Lululemon",
		"Nike SB",
		"Adidas Originals",
		"Polo Ralph Lauren",
		"The North Face",
		"Columbia",
		"Patagonia",
		"Canada Goose",
		"UGG",
		"Timberland Pro",
		"Carhartt",
		"Dickies",
		"Red Wing Shoes",
		"Caterpillar",
		"DeWalt",
		"Black & Decker",
		"Makita",
		"Bosch",
		"Milwaukee",
		"Dremel",
		"Karcher",
		"Fujifilm",
		"Canon",
		"Nikon",
		"Sony",
		"GoPro",
		"DJI",
		"Razer",
		"Logitech",
		"Corsair",
		"HyperX",
		"Audio-Technica",
		"Bose",
		"JBL",
		"Beats by Dre",
		"Sennheiser",
		"Philips",
	}

	for _, i := range brands {
		newTrie.Insert(i)
	}

	completions := newTrie.AutoComplete(data2)

	res = Marshal(completions)
	return res
}

func Marshal(v interface{}) []byte {
	results, err := json.Marshal(v)
	if err != nil {
		log.Println("	[ERROR] Failed to Marshal ", err)
		return nil
	}
	return results
}

func Unmarshal(data []byte, v interface{}) error {
	if err := json.Unmarshal([]byte(data), v); err != nil {
		log.Println("	[ERROR] Failed to Unmarshal ", err)
		return err
	}
	return nil
}
