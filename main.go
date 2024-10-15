package main

import (
	"fmt"
	"strings"
)

func main() {
	// Triangle(5)
	// Triangle2(5)
	// Triangle3(5)
	// Triangle4(5)
	// Triangle5(5)
	Triangle6(5)

	// Diagonal1(5)
	// Diagonal2(5)
	// Diagonal3(5)

	// fmt.Println(CoderByteTokenMixString("1,2,3,4", "abcdefghijklmn"))
	// fmt.Println(CoderByteTokenMixString("1,4,13", "02j9wthc5f"))

	// fmt.Println(CoderByteTokenMixArrOfString([]string{"1", "2", "3", "4"}, "abcdefghijklmn"))
	// fmt.Println(CoderByteTokenMixArrOfString([]string{"1", "4", "13"}, "02j9wthc5f"))

}

//	     *
//		* *
//	   * * *
//	  * * * *
//	 * * * * *
func Triangle(n int) {
	for i := 0; i <= n; i++ {
		for j := n - i; j > 0; j-- {
			fmt.Printf(" ")
		}

		for k := 0; k < i; k++ {
			fmt.Printf("* ")
		}

		fmt.Println()
	}
}

// *
// * *
// * * *
// * * * *
// * * * * *
func Triangle2(n int) {
	for i := 0; i <= n; i++ {

		for k := 0; k < i; k++ {
			fmt.Printf("* ")
		}

		fmt.Println()
	}
}

// * * * * *
// * * * *
// * * *
// * *
// *
func Triangle3(n int) {
	for i := 0; i <= n; i++ {

		for k := n - i; k > 0; k-- {
			fmt.Printf("* ")
		}

		fmt.Println()
	}
}

// ! * * * * *
// !  * * * *
// !   * * *
// !    * *
// !     *
func Triangle4(n int) {
	for i := 0; i <= n; i++ {

		for j := n - i; j < n; j++ {
			fmt.Printf(" ")
		}

		for k := n - i; k > 0; k-- {
			fmt.Printf("* ")
		}

		fmt.Println()
	}
}

// !     *
// !    **
// !   ***
// !  ****
// ! *****
func Triangle5(n int) {
	for i := 0; i < n; i++ {
		for s := i; s < n; s++ {
			fmt.Printf(" ")
		}

		for j := 0; j <= i; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

// ! *****
// !  ****
// !   ***
// !    **
// !     *
func Triangle6(n int) {
	for i := 0; i < n; i++ {
		for s := n; s < n+i; s++ {
			fmt.Printf(" ")
		}

		for j := 0; j < n-i; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

// ! *
// !  *
// !   *
// !    *
// !     *
func Diagonal1(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}

		fmt.Println()
	}
}

// !     *
// !    *
// !   *
// !  *
// ! *
func Diagonal2(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i+j == n-1 {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}

		fmt.Println()
	}
}

// ! *   *
// !  * *
// !   *
// !  * *
// ! *   *
func Diagonal3(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i+j == n-1 || i == j {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}

		fmt.Println()
	}
}

func CoderByteTokenMixString(answer, token string) string {
	result := ""

	// tokenLen := len(token)
	tokenCounter := 0
	for i := 0; i < len(answer); i++ {
		result += string(answer[i])
		result += string(token[tokenCounter])
		tokenCounter++
	}

	// result  = append(result, string(token[tokenCounter:]...))
	// result += string(tokenCounter:)
	result += token[tokenCounter:]

	return result

}

func CoderByteTokenMixArrOfString(answerArrOfString []string, token string) string {
	answer := strings.Join(answerArrOfString, ",")
	result := ""

	// tokenLen := len(token)
	tokenCounter := 0
	for i := 0; i < len(answer); i++ {
		result += string(answer[i])
		result += string(token[tokenCounter])
		tokenCounter++
	}

	// result  = append(result, string(token[tokenCounter:]...))
	// result += string(tokenCounter:)
	result += token[tokenCounter:]

	return result

}
