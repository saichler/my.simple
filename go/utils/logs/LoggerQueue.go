package logs

import (
	"errors"
	"github.com/saichler/my.simple/go/utils/queues"
	"time"
)

type LoggerQueue struct {
	queue      *queues.Queue
	loggerImpl LoggerImpl
}

type LogLevel int

const (
	Trace_Level   LogLevel = 1
	Debug_Level   LogLevel = 2
	Info_Level    LogLevel = 3
	Warning_Level LogLevel = 4
	Error_Level   LogLevel = 5
)

type LoggerEntry struct {
	any  interface{}
	anys []interface{}
	t    int64
	l    LogLevel
}

func NewLoggerQueue(loggerImpl LoggerImpl) *LoggerQueue {
	lq := &LoggerQueue{}
	lq.loggerImpl = loggerImpl
	lq.queue = queues.NewQueue("Logger Queue", 50000)
	go lq.processQueue()
	return lq
}

func (q *LoggerQueue) SetLoggerImpl(impl LoggerImpl) {
	q.loggerImpl = impl
}

func (q *LoggerQueue) Empty() bool {
	return q.queue.Size() == 0
}

func (q *LoggerQueue) processQueue() {
	for {
		entry := q.queue.Next().(*LoggerEntry)
		var str string
		switch entry.l {
		case Trace_Level:
			str = TraceToString(entry.any, entry.anys...)
		case Debug_Level:
			str = DebugToString(entry.any, entry.anys...)
		case Info_Level:
			str = InfoToString(entry.any, entry.anys...)
		case Warning_Level:
			str = WarningToString(entry.any, entry.anys...)
		case Error_Level:
			str = ErrorToString(entry.any, entry.anys...)
		}
		q.loggerImpl.Print(str)
	}
}

func newEntry(l LogLevel, any interface{}, anys ...interface{}) *LoggerEntry {
	return &LoggerEntry{
		t:    time.Now().Unix(),
		l:    l,
		any:  any,
		anys: anys,
	}
}

func (q *LoggerQueue) Trace(any interface{}, anys ...interface{}) {
	q.queue.Add(newEntry(Trace_Level, any, anys...))
}

func (q *LoggerQueue) Debug(any interface{}, anys ...interface{}) {
	q.queue.Add(newEntry(Debug_Level, any, anys...))
}

func (q *LoggerQueue) Info(any interface{}, anys ...interface{}) {
	q.queue.Add(newEntry(Info_Level, any, anys...))
}

func (q *LoggerQueue) Warning(any interface{}, anys ...interface{}) {
	q.queue.Add(newEntry(Warning_Level, any, anys...))
}

func (q *LoggerQueue) Error(any interface{}, anys ...interface{}) error {
	q.queue.Add(newEntry(Error_Level, any, anys...))
	err := ErrorToString(any, anys...)
	return errors.New(err)
}
