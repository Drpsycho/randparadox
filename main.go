package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type t_stage struct {
	rightDoor   int
	choosenDoor int
	doors       [3]bool
}

//колличество дверей
const doorNum int = 3

func GetStage() t_stage {
	var stage t_stage

	s1 := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s1)

	//заполняем рандомно правильную дверь и выбранную игроком
	stage.rightDoor = r.Intn(doorNum)
	stage.choosenDoor = r.Intn(doorNum)

	stage.doors[0] = false
	stage.doors[1] = false
	stage.doors[2] = false

	// делаем одну из дверей с выйгрышом
	stage.doors[stage.rightDoor] = true

	// fmt.Println("Choose door - ", stage.choosenDoor, "\nRight door - ", stage.rightDoor)

	// for index, it := range stage.doors {
	// 	fmt.Println("Door #", index, " ", it)
	// }
	return stage
}

func PickDoor(stage t_stage) (bool, bool) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s1)

	//номер двери которую удаляем 0, 1, 2
	numDoorForRemove := 0

	for {
		// рандомно генерим номер двери и пытаемся удалить
		// если это не правильная дверь и не выбранная игроком, тогда удаляем
		// т.е. выходим из цикла с номером удалённой двери
		numDoorForRemove = r.Intn(doorNum)

		if numDoorForRemove != stage.rightDoor {
			if numDoorForRemove != stage.choosenDoor {
				break
			}
		}
	}
	// fmt.Println("Delete door number ", numDoorForRemove)

	// если isRevers правда меняем выбор
	res_with_repick := 0
	res_without_repick := 0

	for j := 0; j < 3; j++ {
		if j != numDoorForRemove {
			if j != stage.choosenDoor {
				// fmt.Println("Return door number - ", j)
				res_with_repick = j
			}

			if j == stage.choosenDoor {
				// fmt.Println("Return door number - ", j)
				res_without_repick = j
			}
		}
	}

	// fmt.Println("Result is ", stage.doors[res])
	return stage.doors[res_with_repick], stage.doors[res_without_repick]
}

func worker(steps int) {
	result := 0
	result2 := 0
	for j := 0; j < steps; j++ {
		//получаем с чем работать
		stage := GetStage()
		//выбираем дверь, если ок, плюсуем
		res_with_repick, res_without_repick := PickDoor(stage)
		if res_with_repick {
			result++
		}
		if res_without_repick {
			result2++
		}
	}
	fmt.Println("With Repick is bingo -    ", result, " steps = ", steps)
	fmt.Println("Without Repick is bingo - ", result2, " steps = ", steps)
}

func main() {
	if len(os.Args) != 2 {
		s1 := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s1)
		fmt.Println("Randomizer in work. Give you 0 or 1 or 2")
		for i := 1; i < 20; i++ {
			fmt.Println("random result = ", r.Intn(doorNum))
		}
		fmt.Println("For start give as argument amount of steps")
		os.Exit(0)
	}

	steps, _ := strconv.Atoi(os.Args[1])
	worker(steps)
}
