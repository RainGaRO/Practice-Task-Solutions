package queue

type Queue struct {
	Collection []interface{}
}

func (q *Queue) Enqueue(value interface{}) {
	q.Collection = append(q.Collection, value)
}

func (q *Queue) Dequeue() {
	q.Collection = append(q.Collection[1:])
}

func (q *Queue) Peek() interface{} {
	if q.IsEmpty() {
		return nil
	}
	return q.Collection[len(q.Collection)-1]
}

func (q *Queue) Front() interface{} {
	if q.IsEmpty() {
		return nil
	}
	return q.Collection[0]
}

func (q *Queue) Size() int {
	return len(q.Collection)
}

func (q *Queue) IsEmpty() bool {
	if len(q.Collection) != 0 {
		return false
	}
	return true
}

// Пример использования очереди
// func main() {
//     queue := &Queue{}

//     // Добавляем элементы в очередь
//     for i := 0; i < 5; i++ {
//         queue.Enqueue("Элемент " + fmt.Sprint(i))
//     }

//     // Выводим элементы очереди в обратном порядке
//     for queue.Size() > 0 {
//         element := queue.Dequeue().(string)
//         fmt.Println(element)
//     }
// }
