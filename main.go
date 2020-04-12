package main

import (
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

//Memory trainer
//
//1. Build the numbers to persons table according to the "Quantum memory power" book.
//2. Train this table.
//3. Put the persons to your memory palace.
//4. Recall the persons from the palace and translate them to numbers.

func main() {
	args := os.Args[1:]

	if len(args) > 0 {

		fmt.Println(args)

		if args[0] == "1" {
			fmt.Println("\nPersons training...\n")
			time.Sleep(time.Second * time.Duration(3))
			//start with default parameters
			trainPersons(52, 3, 5, true)
			return
		} else if args[0] == "2" {
			fmt.Println("\nMemory palace training...\n")
			time.Sleep(time.Second * time.Duration(3))
			//start with default parameters
			trainPalace(2, false, false, false, true, true)
			return
		}
	}

	var mode string
	for mode != "1" && mode != "2" {
		fmt.Println("Choose the mode: [1] Persons training, [2] Memory palace training")
		fmt.Scan(&mode)
	}

	if mode == "1" {
		//persons training

		fmt.Println("\nPersons training...\n")
		time.Sleep(time.Second * time.Duration(3))

		showTags := true
		sec := 4
		itemsPerLine := 8
		trainPersons(52, sec, itemsPerLine, showTags)
	} else if mode == "2" {
		//memory palace training

		fmt.Println("\nMemory palace training...\n")
		time.Sleep(time.Second * time.Duration(3))

		reverse := false
		random := false
		showTag := false
		showNumber := true
		oneLine := true
		trainPalace(2, reverse, random, showTag, showNumber, oneLine)
	}
}

func trainPersons(cardsAmount, intervalInSeconds, itemsPerLine int, showTags bool) {
	seed := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(seed)
	var cards map[int][]string = map[int][]string{
		0:  {"папаха    ", "b"},
		1:  {"ракета    ", "b"},
		2:  {"блюдо     ", "b"},
		3:  {"трость    ", "b"},
		4:  {"шарф      ", "b"},
		5:  {"протезы   ", "b"},
		6:  {"шпага     ", "b"},
		7:  {"коктейль  ", "b"},
		8:  {"монетка   ", "b"},
		9:  {"фут. мяч  ", "b"},
		10: {"зуб       ", "b"},
		11: {"робот     ", "b"},
		12: {"бас. мяч  ", "b"},
		13: {"часы      ", "b"},
		14: {"помада    ", "b"},
		15: {"перо      ", "b"},
		16: {"тату      ", "b"},
		17: {"леденец   ", "b"},
		18: {"пачка     ", "b"},
		19: {"листья    ", "b"},
		20: {"кресло    ", "b"},
		21: {"шест      ", "b"},
		22: {"сигарета  ", "b"},
		23: {"небоскрёб ", "b"},
		24: {"доспехи   ", "b"},
		25: {"нож       ", "b"},
		26: {"кросовки  ", "b"},
		27: {"гантеля   ", "b"},
		28: {"лопата    ", "b"},
		29: {"ударные   ", "b"},
		30: {"семечки   ", "b"},
		31: {"таблетки  ", "b"},
		32: {"молоток   ", "b"},
		33: {"винтовка  ", "b"},
		34: {"вилка     ", "b"},
		35: {"паяльник  ", "b"},
		36: {"молния    ", "b"},
		37: {"кисть     ", "b"},
		38: {"г. клюшка ", "b"},
		39: {"виски     ", "b"},
		40: {"лапша     ", "b"},
		41: {"эл. гитара", "b"},
		42: {"лимон     ", "b"},
		43: {"бейсболка ", "b"},
		44: {"котелок   ", "b"},
		45: {"паутина   ", "b"},
		46: {"автомат   ", "b"},
		47: {"колонка   ", "b"},
		48: {"вагон м.  ", "b"},
		49: {"будда     ", "b"},
		50: {"конь      ", "b"},
		51: {"спас. круг", "b"},
		52: {"колода    ", "b"},
		53: {"костер    ", "b"},
		54: {"пиво      ", "b"},
		55: {"красн. куб", "b"},
		56: {"венок     ", "b"},
		57: {"бассейн   ", "b"},
		58: {"микрофон  ", "b"},
		59: {"пудель    ", "b"},
		60: {"скрипка   ", "b"},
		61: {"ребенок   ", "b"},
		62: {"кубик руб.", "b"},
		63: {"коньки    ", "b"},
		64: {"мобильник ", "b"},
		65: {"расчёска  ", "b"},
		66: {"горшок    ", "b"},
		67: {"лет. мышь ", "b"},
		68: {"скалка    ", "b"},
		69: {"диадема   ", "b"},
		70: {"ворота    ", "b"},
		71: {"скафандр  ", "b"},
		72: {"поб. слон ", "b"},
		73: {"танк      ", "b"},
		74: {"газета    ", "b"},
		75: {"волга     ", "b"},
		76: {"корова    ", "b"},
		77: {"радио     ", "b"},
		78: {"икебана   ", "b"},
		79: {"питон     ", "b"},
		80: {"приставка ", "b"},
		81: {"брусья    ", "b"},
		82: {"машина    ", "b"},
		83: {"груша     ", "b"},
		84: {"буденовка ", "b"},
		85: {"флаг      ", "b"},
		86: {"тайд      ", "b"},
		87: {"шахматы   ", "b"},
		88: {"гитара    ", "b"},
		89: {"рентген   ", "b"},
		90: {"обои      ", "b"},
		91: {"яхта      ", "b"},
		92: {"ж. человек", "b"},
		93: {"оборода   ", "b"},
		94: {"скальпель ", "b"},
		95: {"лыжи      ", "b"},
		96: {"кокос     ", "b"},
		97: {"доллары   ", "b"},
		98: {"метла     ", "b"},
		99: {"топор     ", "b"}}

	for i := 0; i < cardsAmount; i++ {
		n := rnd.Int31n(100)
		if n < 10 {
			fmt.Print("0")
		}
		fmt.Print(n, " ")
		time.Sleep(time.Second)

		if i != cardsAmount-1 {
			if ((i+1)%itemsPerLine != 0 && !showTags) || showTags {
				for j := 0; j < intervalInSeconds-1; j++ {
					fmt.Print("-")
					time.Sleep(time.Second)
				}
				fmt.Print(" ")
			}
		}
		if showTags {
			fmt.Print(cards[int(n)][0], " | ")
			time.Sleep(time.Second)
		}
		if (i+1)%itemsPerLine == 0 {
			fmt.Print("\n\n")
		}
	}
	fmt.Println()
}

//52 cells for cards memorising
func trainPalace(intervalInSeconds int, reverse, random, showTag, showNumber, oneLine bool) {
	var stops map[int]string = map[int]string{
		1:  "[Часть] КПП",
		2:  "[Часть] Бухгалтерия",
		3:  "[Часть] Актовый зал",
		4:  "[Часть] Библиотека",
		5:  "[Часть] Медсанчасть",
		6:  "[Часть] Самоволка",
		7:  "[Часть] Дела",
		8:  "[Часть] Плац",
		9:  "[Часть] Холм",
		10: "[Часть] РЛС",

		11: "[КП] Вход в бункер",
		12: "[КП] Секретка",
		13: "[КП] Коммутатор",
		14: "[КП ЛАЗ] Дежурный",
		15: "[КП ЛАЗ] Громы",
		16: "[КП ЛАЗ] Койка",
		17: "[КП ЛАЗ] АТС",
		18: "[КП ЛАЗ] Ноутбук",
		19: "[КП ЛАЗ] Паяльник",
		20: "[КП ЛАЗ] Стол",

		21: "[КП Зал] Майор",
		22: "[КП Зал] Планшеты",
		23: "[КП ПРЦ] Приемники",
		24: "[КП ПРЦ] Новый год",
		25: "[КП] Телеграф",
		26: "[КП] Кладовка", //the half of the deck
		27: "[Часть] Вышка",
		28: "[Часть] Релейка",
		29: "[Часть] Окоп",
		30: "[Часть] Автопарк",

		31: "[Казарма] Тумба",
		32: "[Казарма] Коптерка",
		33: "[Казарма] Сейф",
		34: "[Казарма] Штанги",
		35: "[Казарма] Кровать",
		36: "[Казарма] Бытовка",
		37: "[Казарма] Сушилка",
		38: "[Казарма] Канцелярия",
		39: "[Казарма] Комната досуга",
		40: "[Казарма] Оружейка",

		41: "[Часть] Курилка",
		42: "[Часть] Теплица",
		43: "[Часть] Спорт городок",
		44: "[Часть] Штаб",
		45: "[Вторая] Столовая",
		46: "[Вторая] Гостиница",
		47: "[Вторая] Турники",
		48: "[Вторая] Плац 2",
		49: "[Вторая] Склад",
		50: "[Вторая] Чепок",

		51: "[Дембель] Уазик",
		52: "[Дембель] Вокзал"}

	rx := regexp.MustCompile("\\[.*\\]\\s")

	var k int
	for i := 1; i <= 52; i++ {
		k = i
		if reverse {
			k = 53 - i
		}

		//arrow
		if i != 1 {
			fmt.Print(" -> ")
			time.Sleep(time.Second * time.Duration(intervalInSeconds))
		}

		//tag & name
		tag := rx.FindString(stops[k])
		if showTag {
			fmt.Print(stops[k])
		} else {
			fmt.Print(strings.Replace(stops[k], tag, "", -1))
		}

		//number
		if showNumber {
			fmt.Print("(", k, ")")
		}
		time.Sleep(time.Second)

		if !oneLine {
			fmt.Println()
		} else {
			if !reverse && k%10 == 0 {
				fmt.Print("\n\n")
			} else if reverse && k%10 == 1 {
				fmt.Print("\n\n")
			}
		}
	}
}
