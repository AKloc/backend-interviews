package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Destination Schemas

type Product struct {
	Id       int
	Title    string
	Variants []Variant
}

type Variant struct {
	Title    string
	ImageURL string
}

// Schemas that adhere to the JSON schema in the data string.
// GOTCHAS: In order for the JSON unmarshal to work implicitly, the names of the properties have to line up exactly
//	like the below, or they can use Struct Tags, e.g. `json:"Title"`

type RawMessage struct {
	Products []struct {
		Id    int
		Title string
	}
	Variants []struct {
		Id         int
		Product_id int
		Title      string
	}
	Images []struct {
		Variant_id int
		Image_url  string
	}
}

const data = `{"products":[{"id":13,"title":"Adidas Shoe"},{"id":3,"title":"A Bathing Ape Collectible"},{"id":23,"title":"A Bathing Ape Collectible"},{"id":21,"title":"Medicom Painting"},{"id":4,"title":"Adidas Book"},{"id":20,"title":"Nike Shoe"},{"id":7,"title":"FunkoPop Collectible"},{"id":11,"title":"Nike Painting"},{"id":10,"title":"Adidas Painting"},{"id":28,"title":"Medicom Painting"},{"id":5,"title":"Be@rBrick Shoe"},{"id":9,"title":"Be@rBrick Painting"},{"id":6,"title":"Adidas Book"},{"id":14,"title":"Nike Collectible"},{"id":2,"title":"A Bathing Ape Collectible"},{"id":8,"title":"Nike Collectible"},{"id":17,"title":"Nike Shoe"},{"id":29,"title":"Adidas Painting"},{"id":26,"title":"Adidas Book"},{"id":15,"title":"FunkoPop Book"},{"id":12,"title":"Medicom Collectible"},{"id":22,"title":"A Bathing Ape Collectible"},{"id":16,"title":"Adidas Shoe"},{"id":1,"title":"Nike Collectible"},{"id":25,"title":"Nike Book"},{"id":24,"title":"Medicom Painting"},{"id":27,"title":"Medicom Book"},{"id":19,"title":"Medicom Shoe"},{"id":18,"title":"Adidas Shoe"}],"images":[{"variant_id":106,"image_url":"https://cdn.thentwrk.com/images/29"},{"variant_id":107,"image_url":"https://cdn.thentwrk.com/images/7"},{"variant_id":140,"image_url":"https://cdn.thentwrk.com/images/3"},{"variant_id":186,"image_url":"https://cdn.thentwrk.com/images/2"},{"variant_id":118,"image_url":"https://cdn.thentwrk.com/images/8"},{"variant_id":153,"image_url":"https://cdn.thentwrk.com/images/9"},{"variant_id":159,"image_url":"https://cdn.thentwrk.com/images/24"},{"variant_id":143,"image_url":"https://cdn.thentwrk.com/images/19"},{"variant_id":137,"image_url":"https://cdn.thentwrk.com/images/22"},{"variant_id":102,"image_url":"https://cdn.thentwrk.com/images/10"},{"variant_id":162,"image_url":"https://cdn.thentwrk.com/images/11"},{"variant_id":132,"image_url":"https://cdn.thentwrk.com/images/6"},{"variant_id":141,"image_url":"https://cdn.thentwrk.com/images/1"},{"variant_id":180,"image_url":"https://cdn.thentwrk.com/images/20"},{"variant_id":159,"image_url":"https://cdn.thentwrk.com/images/21"},{"variant_id":109,"image_url":"https://cdn.thentwrk.com/images/23"},{"variant_id":166,"image_url":"https://cdn.thentwrk.com/images/4"},{"variant_id":152,"image_url":"https://cdn.thentwrk.com/images/14"},{"variant_id":151,"image_url":"https://cdn.thentwrk.com/images/28"},{"variant_id":147,"image_url":"https://cdn.thentwrk.com/images/25"},{"variant_id":182,"image_url":"https://cdn.thentwrk.com/images/17"},{"variant_id":184,"image_url":"https://cdn.thentwrk.com/images/16"},{"variant_id":162,"image_url":"https://cdn.thentwrk.com/images/27"},{"variant_id":165,"image_url":"https://cdn.thentwrk.com/images/26"},{"variant_id":155,"image_url":"https://cdn.thentwrk.com/images/13"},{"variant_id":116,"image_url":"https://cdn.thentwrk.com/images/18"},{"variant_id":110,"image_url":"https://cdn.thentwrk.com/images/5"},{"variant_id":140,"image_url":"https://cdn.thentwrk.com/images/15"},{"variant_id":117,"image_url":"https://cdn.thentwrk.com/images/12"}],"variants":[{"id":124,"product_id":9,"title":"Small"},{"id":157,"product_id":20,"title":"Small"},{"id":184,"product_id":29,"title":"Small"},{"id":103,"product_id":2,"title":"Small"},{"id":155,"product_id":19,"title":"Medium"},{"id":152,"product_id":18,"title":"Medium"},{"id":181,"product_id":28,"title":"Small"},{"id":164,"product_id":22,"title":"Medium"},{"id":170,"product_id":24,"title":"Medium"},{"id":182,"product_id":28,"title":"Medium"},{"id":173,"product_id":25,"title":"Medium"},{"id":179,"product_id":27,"title":"Medium"},{"id":125,"product_id":9,"title":"Medium"},{"id":137,"product_id":13,"title":"Medium"},{"id":144,"product_id":15,"title":"Large"},{"id":119,"product_id":7,"title":"Medium"},{"id":104,"product_id":2,"title":"Medium"},{"id":116,"product_id":6,"title":"Medium"},{"id":149,"product_id":17,"title":"Medium"},{"id":127,"product_id":10,"title":"Small"},{"id":163,"product_id":22,"title":"Small"},{"id":168,"product_id":23,"title":"Large"},{"id":106,"product_id":3,"title":"Small"},{"id":105,"product_id":2,"title":"Large"},{"id":147,"product_id":16,"title":"Large"},{"id":109,"product_id":4,"title":"Small"},{"id":153,"product_id":18,"title":"Large"},{"id":148,"product_id":17,"title":"Small"},{"id":101,"product_id":1,"title":"Medium"},{"id":131,"product_id":11,"title":"Medium"},{"id":135,"product_id":12,"title":"Large"},{"id":162,"product_id":21,"title":"Large"},{"id":129,"product_id":10,"title":"Large"},{"id":143,"product_id":15,"title":"Medium"},{"id":118,"product_id":7,"title":"Small"},{"id":183,"product_id":28,"title":"Large"},{"id":120,"product_id":7,"title":"Large"},{"id":141,"product_id":14,"title":"Large"},{"id":156,"product_id":19,"title":"Large"},{"id":185,"product_id":29,"title":"Medium"},{"id":140,"product_id":14,"title":"Medium"},{"id":176,"product_id":26,"title":"Medium"},{"id":122,"product_id":8,"title":"Medium"},{"id":171,"product_id":24,"title":"Large"},{"id":178,"product_id":27,"title":"Small"},{"id":107,"product_id":3,"title":"Medium"},{"id":150,"product_id":17,"title":"Large"},{"id":166,"product_id":23,"title":"Small"},{"id":112,"product_id":5,"title":"Small"},{"id":165,"product_id":22,"title":"Large"},{"id":180,"product_id":27,"title":"Large"},{"id":133,"product_id":12,"title":"Small"},{"id":139,"product_id":14,"title":"Small"},{"id":102,"product_id":1,"title":"Large"},{"id":154,"product_id":19,"title":"Small"},{"id":121,"product_id":8,"title":"Small"},{"id":145,"product_id":16,"title":"Small"},{"id":108,"product_id":3,"title":"Large"},{"id":114,"product_id":5,"title":"Large"},{"id":175,"product_id":26,"title":"Small"},{"id":146,"product_id":16,"title":"Medium"},{"id":113,"product_id":5,"title":"Medium"},{"id":142,"product_id":15,"title":"Small"},{"id":115,"product_id":6,"title":"Small"},{"id":126,"product_id":9,"title":"Large"},{"id":117,"product_id":6,"title":"Large"},{"id":160,"product_id":21,"title":"Small"},{"id":158,"product_id":20,"title":"Medium"},{"id":159,"product_id":20,"title":"Large"},{"id":177,"product_id":26,"title":"Large"},{"id":132,"product_id":11,"title":"Large"},{"id":169,"product_id":24,"title":"Small"},{"id":151,"product_id":18,"title":"Small"},{"id":138,"product_id":13,"title":"Large"},{"id":186,"product_id":29,"title":"Large"},{"id":130,"product_id":11,"title":"Small"},{"id":123,"product_id":8,"title":"Large"},{"id":136,"product_id":13,"title":"Small"},{"id":111,"product_id":4,"title":"Large"},{"id":134,"product_id":12,"title":"Medium"},{"id":128,"product_id":10,"title":"Medium"},{"id":167,"product_id":23,"title":"Medium"},{"id":110,"product_id":4,"title":"Medium"},{"id":100,"product_id":1,"title":"Small"},{"id":161,"product_id":21,"title":"Medium"},{"id":172,"product_id":25,"title":"Small"},{"id":174,"product_id":25,"title":"Large"}]}`

func main() {
	var oldJson RawMessage

	// COMMON TRIP-UP: Have to know to convert the string to a byte array.
	err := json.Unmarshal([]byte(data), &oldJson)
	if err != nil {
		return
	}

	// Convert the data. Can start top down or bottom up.
	// Important bit: Interviewee should be using a map for fast lookups when putting the data together, rather than
	//   nesting for loops. If they start with for loops, it's worth asking if they know how they might be able to
	//   make the runtime more efficient.

	products := make(map[int]*Product)
	for i, product := range oldJson.Products {
		products[i] = &Product{
			Id:    product.Id,
			Title: product.Title,
		}
	}

	// Put the images in maps - key is the variant id, value is the string url.
	images := make(map[int]string)
	for _, image := range oldJson.Images {
		images[image.Variant_id] = image.Image_url
	}

	// Key: the product ID. Value: array of variants.
	for _, variant := range oldJson.Variants {
		variantToAdd := Variant{
			Title:    variant.Title,
			ImageURL: images[variant.Id],
		}

		v, ok := products[variant.Product_id]
		if !ok {
			log.Println("Couldn't find specified Product ID.")
			continue
		}
		v.Variants = append(v.Variants, variantToAdd)
	}

	for _, product := range products {
		fmt.Printf("%v\n", *product)
	}
}

/*
We frequently have the need to coalesce related data from different sources into a
single comprehensive structure.

We want our clients to receive this data in a way that's easier to handle.
Using the provided JSON data as a guide, we would like you to convert it
to something resembling the following:
  [
    type:Product
    └── [
          property: Variants
          └── type: Variant
              └── property: ImageURL
        ]
  ]

  - There are 0...N Products
  - A Product can contain 0...N Variants
  - A Variant can have 0 or 1 images

!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
You will use the constant string value `data`
that is defined at the end of this file
!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

`data` is a JSON string with the following format:
   {
     products: [{product type}],
     variants: [{variant type}],
     images:   [{imageobject}],
   }

...where the product type is structured similar to:
  {
    "id":13,
    "title":"Adidas Shoe"
  }

...and the variant type is structured similar to:
  {
    "id":124,
    "product_id":13,
    "title":"Small"
  }

...and the image type is structured similar to:
  {
    "variant_id":124,
    "image_url":"https://cdn.thentwrk.com/images/25"
  }

*/
