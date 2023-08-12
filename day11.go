
package main


import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type WorryOperation int

const (
	PLUS WorryOperation = iota
	MULTIPLY
	SQUARED
)

type Monkey struct {
	monkeyId           uint64
	items              []uint64
	worryOperation     WorryOperation
	worryValueModifier uint64
	testValue          uint64
	throwToMonkeyTrue  uint64
	throwToMonkeyFalse uint64
	inspectCounter     uint64
}

func (m *Monkey) String() string {
	return fmt.Sprintf("MonkeyId: %v items: %v inspectCounter: %v", m.monkeyId, m.items, m.inspectCounter)
}

func getMonkeyIdFromString(row string) uint64 {
	re := regexp.MustCompile(`Monkey (\d+):`)
	match := re.FindStringSubmatch(row)
	if len(match) != 2 {
		log.Fatal("Could not parse monkey id: ", row)
	}
	monkeyId, err := strconv.ParseUint(match[1], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return monkeyId
}

func getItemsFromString(row string) []uint64 {
	prefixString := "  Starting items: "
	items := strings.Split(row[len(prefixString):], ", ")
	itemsInt := make([]uint64, len(items))
	for idx, item := range items {
		itemInt, err := strconv.ParseUint(item, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		itemsInt[idx] = itemInt
	}
	return itemsInt
}

func getOperationsFromString(row string) (WorryOperation, uint64) {
	parsedString := strings.Split(row, " ")

	var value uint64
	var operation WorryOperation
	var old bool

	if parsedString[len(parsedString)-1] == "old" {
		old = true
	} else {
		var err error
		value, err = strconv.ParseUint(parsedString[len(parsedString)-1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
	}

	if parsedString[len(parsedString)-2] == "+" {
		if old {
			value = 2
			operation = MULTIPLY
		} else {
			operation = PLUS
		}
	} else if parsedString[len(parsedString)-2] == "*" {
		if old {
			operation = SQUARED
		} else {
			operation = MULTIPLY
		}
	} else {
		log.Fatal("Unknown operation")
	}

	return operation, value
}

func getTestValueFromString(row string) uint64 {
	re := regexp.MustCompile(`  Test: divisible by (\d+)`)
	match := re.FindStringSubmatch(row)
	if len(match) != 2 {
		log.Fatal("Could not parse test value: ", row)
	}

	value, err := strconv.ParseUint(match[1], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return value
}

func getThrowToMonkeyFromString(row string) uint64 {
	re := regexp.MustCompile(`throw to monkey (\d+)`)
	match := re.FindStringSubmatch(row)
	if len(match) != 2 {
		log.Fatal("Could not parse throw to monkey: ", row)
	}

	value, err := strconv.ParseUint(match[1], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return value
}

func parseStringsForMonkeys(rows []string) *[]Monkey {
	monkeys := make([]Monkey, 0)

	idx := 0
	for {
		if idx >= len(rows) {
			break
		}

		monkeyId := getMonkeyIdFromString(rows[idx])
		items := getItemsFromString(rows[idx+1])
		worryOperation, worryValueModifier := getOperationsFromString(rows[idx+2])
		testValue := getTestValueFromString(rows[idx+3])
		throwToMonkeyTrue := getThrowToMonkeyFromString(rows[idx+4])
		throwToMonkeyFalse := getThrowToMonkeyFromString(rows[idx+5])

		monkey := Monkey{monkeyId: monkeyId,
			items:              items,
			worryOperation:     worryOperation,
			worryValueModifier: worryValueModifier,
			testValue:          testValue,
			throwToMonkeyTrue:  throwToMonkeyTrue,
			throwToMonkeyFalse: throwToMonkeyFalse,
			inspectCounter:     0}

		monkeys = append(monkeys, monkey)
		idx += 7
	}

	return &monkeys
}

func (m *Monkey) throwItem(superMod uint64) (uint64, uint64) {

	item := m.items[0]
	m.items = m.items[1:]

	switch m.worryOperation {
	case MULTIPLY:
		item = item * m.worryValueModifier
	case PLUS:
		item = item + m.worryValueModifier
	case SQUARED:
		item = item * item
	default:
		log.Fatalln("Not a valid worry operation")
	}

	if superMod == 0 {
		item = item / 3
	} else {
		item = item % superMod
	}

	m.inspectCounter++

	testPassed := item%m.testValue == 0
	var toThrowMonkeyId uint64
	if testPassed {
		toThrowMonkeyId = m.throwToMonkeyTrue
	} else {
		toThrowMonkeyId = m.throwToMonkeyFalse
	}

	return item, toThrowMonkeyId
}

func (m *Monkey) hasNoItemLeft() bool {
	return len(m.items) == 0
}

func runMonkeyRounds(monkeys *[]Monkey, rounds int, superMod uint64) {
	for round := 0; round < rounds; round++ {
		for idx := 0; idx < len(*monkeys); idx++ {
			monkey := &(*monkeys)[idx]
			for {
				if monkey.hasNoItemLeft() {
					break
				}
				item, toThrowMonkeyId := monkey.throwItem(superMod)

				recevingMonkey := &(*monkeys)[toThrowMonkeyId]
				if recevingMonkey.monkeyId != toThrowMonkeyId {
					log.Fatalf("Wrong monkey id")
				}

				recevingMonkey.items = append(recevingMonkey.items, item)
			}
		}
	}
}

func day11_part1() {
	rows, err := getRowsFromFile("input11.txt")

	if err != nil {
		log.Fatal(err)
	}

	monkeys := parseStringsForMonkeys(rows[:len(rows)-1])
	runMonkeyRounds(monkeys, 20, 0)

	sort.Slice(*monkeys, func(i, j int) bool {
		return (*monkeys)[i].inspectCounter > (*monkeys)[j].inspectCounter
	})

	monkeyBusiness := (*monkeys)[0].inspectCounter * (*monkeys)[1].inspectCounter
	fmt.Println("Monkey business: ", monkeyBusiness)
}

func getSuperMod(monkeys *[]Monkey) uint64 {
	var superMod uint64
	superMod = 1
	for _, monkey := range *monkeys {
		superMod *= monkey.testValue
	}
	return superMod
}

func day11_part2() {
	rows, err := getRowsFromFile("input11.txt")

	if err != nil {
		log.Fatal(err)
	}

	monkeys := parseStringsForMonkeys(rows[:len(rows)-1])

	runMonkeyRounds(monkeys, 10000, getSuperMod(monkeys))

	sort.Slice(*monkeys, func(i, j int) bool {
		return (*monkeys)[i].inspectCounter > (*monkeys)[j].inspectCounter
	})

	monkeyBusiness := (*monkeys)[0].inspectCounter * (*monkeys)[1].inspectCounter
	fmt.Println("Monkey business: ", monkeyBusiness)
}
