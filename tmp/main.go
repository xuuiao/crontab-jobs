package main

import "sync"

func mains() {
}

type QueueData interface {
	SetValue(interface{})
	GetValue() interface{}
}

type Queue struct {
	mutex  *sync.Mutex
	Length int
	Data   []*QueueData
}

type StringQueueData struct {
	value string
}

func (s StringQueueData) SetValue(i interface{}) {
	s.value = i.(string)
}

func (s StringQueueData) GetValue() interface{} {
	return s.value
}

type QueueError struct {
}

func (q QueueError) Error() string {
	return "too many"
}

func (q *Queue) Producer(data *QueueData) error {
	if len(q.Data) >= q.Length {
		var err QueueError
		return err
	}
	q.Data = append(q.Data, data)
	return nil
}

func (q *Queue) Consumer() *QueueData {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if len(q.Data) == 0 {
		return nil
	}
	data := q.Data[0]
	q.Data = q.Data[1:]
	return data
}
