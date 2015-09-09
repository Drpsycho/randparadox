package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

type t_stage struct {
	rightDoor   int
	choosenDoor int
	doors       [3]bool
}

//колличество дверей
const doorNum int = 3

// для ожидания конца работы горутин
var wg sync.WaitGroup

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

func PickDoor(stage t_stage, isRevers bool) bool {
	s1 := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s1)

	//номер двери которую удаляем 0, 1, 2
	numDoorForRemove := 0

	for {
		// рандомно генерим номер двери и пытаемся удалить
		// если это не правильная дверь и не выбранная игроком, тогда удаляем
		// т.е. выходим из цикла с номером удалённой двери
		numDoorForRemove := r.Intn(doorNum)

		if numDoorForRemove != stage.rightDoor {
			if numDoorForRemove != stage.choosenDoor {
				break
			}
		}
	}
	// fmt.Println("Delete door number ", numDoorForRemove)

	// если isRevers правда меняем выбор
	res := 0

	for j := 0; j < 3; j++ {
		if j != numDoorForRemove {
			if isRevers {
				if j != stage.choosenDoor {
					// fmt.Println("Return door number - ", j)
					res = j
				}
			} else {
				if j == stage.choosenDoor {
					// fmt.Println("Return door number - ", j)
					res = j
				}
			}
		}
	}

	// fmt.Println("Result is ", stage.doors[res])
	return stage.doors[res]
}

func worker(steps int, repick bool) {
	result := 0

	for j := 0; j < steps; j++ {
		//получаем с чем работать
		stage := GetStage()
		//выбираем дверь, если ок, плюсуем
		if PickDoor(stage, repick) {
			result++
		}
	}
	fmt.Println("Repick is ", repick, " bingo - ", result, " steps = ", steps)

	wg.Done()
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
	wg.Add(2)
	//запускаем две горутины с перевыбором двери(true) и без(false)
	go worker(steps, true)
	go worker(steps, false)

	wg.Wait()
}
