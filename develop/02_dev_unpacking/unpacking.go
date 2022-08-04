package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func UnpackString(str string) (string, error) {
	// Инициализация слайса рун для работы с каждым символом по отдельности
	var unpacked []rune
	// Инициализация флага обратного слеша
	var isSlash bool
	// Преобразование переданной строки в слайс рун
	runes := []rune(str)
	// Цикл по слайсу рун
	for i, val := range runes {
		// Если встретилась цифра и она была первой: возвращаем пустую строку и ошибку
		if i == 0 && unicode.IsDigit(val) {
			return "", fmt.Errorf("invalid string")
		}
		// Если уже встретчался обратный слеш и текущая руна - буква: возвращаем пустую строку и ошибку
		if isSlash && unicode.IsLetter(val) {
			return "", fmt.Errorf("invalid string")
		}
		// Если обратный слеш не встречался и текущая руна - обратный слеш: устанавливаем флаг слеша в истину
		if !isSlash && val == '\\' {
			isSlash = true
			continue
		}
		// Если \ уже встретчался обратный слеш: добавляем текущую руну в результат и сбрасываем флаг
		if isSlash {
			unpacked = append(unpacked, val)
			isSlash = false
			continue
		}
		// Если текущая руна - цифра: преобразовываем ее в интовое значение для подсчета количества добавлений
		if unicode.IsDigit(val) {
			valS := string(val)
			times, _ := strconv.Atoi(valS)
			// Если это количество нулевое: удаляем предыдущую руну из результата
			if times == 0 {
				unpacked = unpacked[:len(unpacked)-1]
				continue
			}
			times--
			// Пока количество добавлений не станет равным нулю: добавляем предыдущую руну в результат
			for times != 0 {
				unpacked = append(unpacked, runes[i-1])
				times--
			}
			continue
		}
		// Если не попали ни в одно из условий, значит это были буквы, которые добавляем в результат
		unpacked = append(unpacked, val)
	}
	return string(unpacked), nil
}
