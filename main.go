package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	success = 0
	failed  = 1
)

func main() {
	os.Exit(run())
}

func run() int {
	byear, bmonth, bday, err := getArgs()
	if err != nil {
		fmt.Println(err)
		return failed
	}

	toyear, tomonth, today := getToday()

	// 誕生日前後判定
	if bmonth >= tomonth {
		if bmonth > tomonth || bday > today {
			toyear--
		}
	}

	// 数秘術計算
	congenitalNumber := sum(bday)
	acquiredNumber := sum(byear, bmonth, bday)
	respectNumber := sum(bmonth, bday)

	// 運勢計算
	yf := sum(toyear, respectNumber)
	mf := sum(tomonth, yf)
	df := sum(today, mf)

	// 出力
	fmt.Println("先天数：　先天的にある数＿能力・資質＆恋愛中")
	fmt.Println("後天数：　後天的に備わる数＿メインの性格")
	fmt.Println("尊敬数：　好きになりやすい数＿性別関係なく")
	fmt.Println()

	if bday%10 == 0 {
		fmt.Printf("先天数: %2d  %s  ※ 強調\n", congenitalNumber, personality(congenitalNumber))
	} else {
		fmt.Printf("先天数: %2d  %s\n", congenitalNumber, personality(congenitalNumber))
	}
	fmt.Printf("後天数: %2d  %s\n", acquiredNumber, personality(acquiredNumber))
	fmt.Printf("尊敬数: %2d  %s\n", respectNumber, personality(respectNumber))
	fmt.Printf("年運　: %2d  %s\n", yf, fortune(yf))
	fmt.Printf("月運　: %2d  %s\n", mf, fortune(mf))
	fmt.Printf("日運　: %2d  %s\n", df, fortune(df))
	fmt.Printf("%d-%d-%d-%d-%d-%d\n", congenitalNumber, acquiredNumber, respectNumber, yf, mf, df)

	return success
}

func sum(num ...int) int {
	if n := num[0]; len(num) == 1 {
		switch n {
		case 11, 22, 33:
			return n
		}
		if n < 10 {
			return n
		}
	}

	ret := 0
	for _, n := range num {
		ret += n/1000 + (n%1000)/100 + (n%100)/10 + n%10
	}

	return sum(ret)
}

func getArgs() (year, month, day int, err error) {
	args := os.Args
	if len(args) != 2 {
		return 0, 0, 0, errors.New("usage: numerology.exe YYYY/MM/DD")
	}
	num := strings.Split(args[1], "/")
	if len(num) != 3 {
		return 0, 0, 0, errors.New("usage: numerology.exe YYYY/MM/DD")
	}

	year, err = strconv.Atoi(num[0])
	if err != nil {
		return 0, 0, 0, err
	}
	month, err = strconv.Atoi(num[1])
	if err != nil {
		return 0, 0, 0, err
	}
	day, err = strconv.Atoi(num[2])
	if err != nil {
		return 0, 0, 0, err
	}

	return year, month, day, nil
}

func getToday() (year, month, day int) {
	t := time.Now()
	y, m, d := t.Date()
	return y, int(m), d
}

func personality(n int) string {
	switch n {
	case 1:
		return "亭主関白な夫（先導）"
	case 2:
		return "母親（サポート）"
	case 3:
		return "子供"
	case 4:
		return "真面目ちゃん（安定）"
	case 5:
		return "自由人"
	case 6:
		return "愛の人（妊婦さん）"
	case 7:
		return "職人"
	case 8:
		return "バランス調整"
	case 9:
		return "人生2週目の優等生（おじいちゃん）"
	case 11:
		return "感情がわかる人"
	case 22:
		return "カリスマ神"
	case 33:
		return "究極的どＭ"
	default:
		panic(n)
	}
}

func fortune(n int) string {
	switch n {
	case 1:
		return "開始"
	case 2, 11:
		return "関係"
	case 3:
		return "想像"
	case 4, 22:
		return "安定"
	case 5:
		return "変化"
	case 6, 33:
		return "愛情"
	case 7:
		return "内省"
	case 8:
		return "達成"
	case 9:
		return "選択"
	default:
		panic(n)
	}
}