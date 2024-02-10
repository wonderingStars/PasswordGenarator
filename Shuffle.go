package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

func getPassword(c echo.Context) error {
	passwordSize := c.FormValue("passwordSize")
	passwordOption := c.FormValue("passwordOption")
	customPassword := c.FormValue("customPassword")
	
	size, _ := strconv.Atoi(passwordSize)
	option, _ := strconv.Atoi(passwordOption)
	value := shuffle(size, option, customPassword)

	return c.String(http.StatusOK, value)
}

func shuffle(lenth int, option int, custom string) string {
	numbers := []byte("1234561234567890<>?,./:|}{][+_)(7890<1234567890<>?,./:|}{][+_)(>?,./:|}{][+_)(")
	letters := []byte("ABCDEFGeabcdeBCDEfghijklmnopqrJKLMNOPstuvwxyzfghijkHIJKLMNOPQRSTUabcdeBCDEfghijklmnopqrJKLMNOPstuvwxyz")
	customRunes := []byte(custom)
	var sb strings.Builder

	rand.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
		letters[i], letters[j] = letters[j], letters[i]

	})

	if option == 1 {
		for i := 0; i < lenth; i++ {
			sb.WriteString(string(letters[i]))
			sb.WriteString(string(numbers[i]))
		}
		fmt.Println(sb.String())
		return sb.String()
	}
	if option == 2 {
		for i := 0; i < lenth; i++ {

			sb.WriteString(string(letters[i]))

		}
		fmt.Println(sb.String())
		return sb.String()
	}
	if option == 3 {
		for i := 0; i < lenth; i++ {

			sb.WriteString(string(numbers[i]))
		}
		fmt.Println(sb.String())
		return sb.String()
	}

	if option == 4 && custom != "" {
		rand.Shuffle(len(customRunes), func(i, j int) {
			customRunes[i], customRunes[j] = customRunes[j], customRunes[i]
		})

		for i := 0; i < lenth; i++ {

			sb.WriteString(string(customRunes[i]))
		}
		fmt.Println(sb.String())
		return sb.String()
	} else {

		fmt.Println("custom is nil")

	}

	return ""
}
