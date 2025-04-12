package core

import (
	"log/slog"
	"math/rand"

	"github.com/ltlaitoff/2048/pkg/assert"
)

/*

RULES

2 2 2 2 => 4 4 0 0

4 2 2 0 => 4 4 0 0

4 0 0 4 => 8 0 0 0

4 2 2 4 => 4 4 4 0

Тобто, нам потрібно зробити так, щоб всі однакові елементи ОДИН раз сумувались, але при цьому не їх результат

Якщо це можна зробити?

### Ідея 1

Спочатку ми здвигаємо все в купу:

4 0 0 4 => 4 4 0 0

А далі запускаємо алгоритм який вираховує пари та сумує їх та додає в новий масив, не змінюючи поточний

4 4 0 0 => 8 0 0 0

Ну і далі знову треба прибрати пробіли...

### Ідея 2

Коли ми здвигаємо елементи ми можемо відразу знаходити їх суму

Але тут є проблема - коли ми здвигаємо наступний елемент, то він може сумуватись з тим, який зараз
А на це не треба

Як це можна вирішити?
Ну, по суті ніяк...


---

Як буде працювати закінчення гри?
...
*/

/*

 2 0 0 0 			4 0 0 0
 2 0 0 0  =>  4 0 0 0
 2 0 0 0  =>  0 0 0 0
 2 0 0 0    	0 0 0 0

Коли ми свайпаємо вверх нам потрібно перевіряти всі елементи зверху вниз, оскільки тільки в такому варіанті ми можемо гарантувати, що все буде ок...

І тоді у нас є проблема з 3-ма циклами, бо їх не повинно бути...

Як це можна переписати правильно?

Ми можемо за раз розраховувати якусь статичну кількість значень, наприклад одну строку або ж рядок

Взагалі цей розрахунок буде однаковий для всіх строк, просто вони йдуть в різному порядку...

Тобто в теоріііїї ми можемо написати функцію яка приймає масив значень, наприклад, 4, і розраховує їх?

А потім з board витягувати значення в деякому порядку та викликати функцію для їх перезаписування

Типу, спочатку першу колонку. Взяти значення всіх рядків в правильному порядку
Розрахувати та перезаписати
Після другу і так далі

Можливо це справді буде гарно...

---

Як буде працювати перевірка одної строки/стовпця:

2		2		2		4
0->	0->	2->	0
0		2		0		0
2		0		0		0

Тобто, нам потрібно для кожного набору значень перевіряти перше значення
Якщо над ним є якесь - то мерджити його
Якщо над ним пустота - то не робити нічого
*/

const SIZE = 4

var board [SIZE][SIZE]int64

func Reset() {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			board[i][j] = 0
		}
	}
	addRandomCell()
}

func Up() {
	SomeOperation(-1, 0, 1, 0, 0, 0)
	addRandomCell()
}

func Left() {
	SomeOperation(0, -1, 0, 1, 0, 0)
	addRandomCell()
}

func Right() {
	SomeOperation(0, 1, 0, 0, 0, -1)
	addRandomCell()
}

func Down() {
	SomeOperation(1, 0, 0, 0, -1, 0)
	addRandomCell()
}

func Map(callback func(value int64)) {
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			callback(board[i][j])
		}
	}
}

func SomeOperation(di int, dj int, startI int, startJ int, baseDiffI int, baseDiffJ int) {
	for k := 1; k < 4; k++ {
		/*
			up: si = 1, di = -1
			down: ei = 3, di = 1

			left: sj = 1, dj = -1
			right: ej = 3, dj = 1
		*/

		for i := startI; i < SIZE+baseDiffI; i++ {
			for j := startJ; j < SIZE+baseDiffJ; j++ {
				if board[i][j] == 0 {
					continue
				}

				newI := i + di
				newJ := j + dj

				if board[newI][newJ] == 0 {
					board[newI][newJ] = board[i][j]
					board[i][j] = 0
				}

				if board[newI][newJ] == board[i][j] {
					board[newI][newJ] = board[i][j] + board[newI][newJ]
					board[i][j] = 0
				}
			}
		}
	}
}

func getEmptyIndexes() [][2]int {
	res := make([][2]int, 0, 16)

	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			if board[i][j] != 0 {
				continue
			}

			res = append(res, [2]int{i, j})
		}
	}

	return res
}

func isBoardFull() bool {
	sum := 0

	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			if board[i][j] != 0 {
				sum += 1
			}
		}
	}

	return sum == SIZE * SIZE
}

func addRandomCell() {
	emptyIndexes := getEmptyIndexes()

	if len(emptyIndexes) == 0 {
		slog.Debug("Not add random cell because all cells is not empty")
		return
	}

	index := rand.Intn(len(emptyIndexes))
	i := emptyIndexes[index][0]
	j := emptyIndexes[index][1]

	assert.Assert(board[i][j] == 0, "Board cell on random add not equal to 0")

	isFour := rand.Intn(10)

	if isFour == 9 {
		board[i][j] = 4
	} else {
		board[i][j] = 2
	}
}
