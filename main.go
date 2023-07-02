package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/a-16-1-inisialisasi-slice", func(c echo.Context) error {
		fruits := []string{"apple", "grape", "banana", "melon"}
		return c.String(http.StatusOK, fruits[0])
	})

	e.GET("/a-16-2-hubungan-array-slice-dengan-array-&-operasi-slice", func(d echo.Context) error {
		fruits := []string{"apple", "grape", "banana", "melon"}
		newFruits := fruits[0:2]

		return d.String(200, newFruits[0])
	})

	e.GET("/a-16-3-merupakan-tipe-data-refference", func(e echo.Context) error {

		var fruits = []string{"apple", "grape", "banana", "melon"}

		aFruits := fruits[0:3]
		bFruits := fruits[1:4]

		aaFruits := aFruits[1:2]
		baFruits := bFruits[0:1]

		fmt.Println(fruits)   // [apple grape banana melon]
		fmt.Println(aFruits)  // [apple grape banana]
		fmt.Println(bFruits)  // [grape banana melon]
		fmt.Println(aaFruits) // [grape]
		fmt.Println(baFruits) // [grape]

		// Buah "grape" diubah menjadi "pinnaple"
		baFruits[0] = "pinnaple"

		fmt.Println(fruits)   // [apple pinnaple banana melon]
		fmt.Println(aFruits)  // [apple pinnaple banana]
		fmt.Println(bFruits)  // [pinnaple banana melon]
		fmt.Println(aaFruits) // [pinnaple]
		fmt.Println(baFruits) // [pinnaple]

		return e.String(http.StatusOK, baFruits[0])
	})

	e.GET("/a-16-4fungsi-len()", func(f echo.Context) error {

		var fruits = []string{"apple", "grape", "banana", "melon"}

		return f.String(200, fruits[1]) // 4

	})

	e.GET("/a-16-5fungsi-cap()", func(g echo.Context) error {

		var fruits = []string{"apple", "grape", "banana", "melon"}
		fmt.Println(len(fruits)) // len: 4
		fmt.Println(cap(fruits)) // cap: 4

		var aFruits = fruits[0:3]
		fmt.Println(len(aFruits)) // len: 3
		fmt.Println(cap(aFruits)) // cap: 4

		var bFruits = fruits[1:4]
		fmt.Println(len(bFruits))        // len: 3
		fmt.Println(cap(bFruits))        // cap: 3
		return g.String(200, aFruits[1]) // 4
	})

	e.GET("/a-16-6-fungsi-append()", func(f echo.Context) error {

		var fruits = []string{"apple", "grape", "banana"}
		var cFruits = append(fruits, "papaya")

		fmt.Println(fruits)  // ["apple", "grape", "banana"]
		fmt.Println(cFruits) // ["apple", "grape", "banana", "papaya"]}

		return f.String(200, cFruits[0]) // 4

	})

	e.GET("/a-16-7-fungsi-copy()", func(f echo.Context) error {

		dst := []string{"potato", "potato", "potato"}
		src := []string{"watermelon", "pinnaple"}
		n := copy(dst, src)

		fmt.Println(dst) // watermelon pinnaple potato
		fmt.Println(src) // watermelon pinnaple
		fmt.Println(n)   // 2

		return f.String(200, dst[0]) // 4

	})

	e.GET("/a-16-8-pengaksesan-elemen-slice-dengan-3-ndeks", func(f echo.Context) error {

		var fruits = []string{"apple", "grape", "banana"}
		var aFruits = fruits[0:2]
		var bFruits = fruits[0:2:2]

		fmt.Println(fruits)      // ["apple", "grape", "banana"]
		fmt.Println(len(fruits)) // len: 3
		fmt.Println(cap(fruits)) // cap: 3

		fmt.Println(aFruits)      // ["apple", "grape"]
		fmt.Println(len(aFruits)) // len: 2
		fmt.Println(cap(aFruits)) // cap: 3

		fmt.Println(bFruits)      // ["apple", "grape"]
		fmt.Println(len(bFruits)) // len: 2
		fmt.Println(cap(bFruits)) // cap: 2

		return f.String(200, fruits[0]) // 4

	})

	e.Start(":9000")
}
