package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type ByteSize float64

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

const (
	checkSize    = 2 * time.Minute
	checkExpired = 2 * time.Hour
)

var (
	stdFatal = log.New(os.Stderr, "\033[0;33mFATAL:\033[0m ", log.LstdFlags|log.Lshortfile)
	stdError = log.New(os.Stderr, "\033[0;31mERROR:\033[0m ", log.LstdFlags|log.Lshortfile)
	stdWarn  = log.New(os.Stderr, "\033[0;35mWARN:\033[0m ", log.LstdFlags|log.Lshortfile)
	ll       *Logger
)

type Logger struct {
	errCount int32
	rotate   Rotate
	level    chan int

	rwm                                  sync.RWMutex
	file                                 *os.File
	debug, info, notic, warn, err, fatal *log.Logger
}

type Rotate struct {
	Size              ByteSize
	Expired, Interval time.Duration
}

func New(fp string, lvl int, rotate Rotate) *Logger {
	f, err := os.OpenFile(fp, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		stdFatal.Fatalln(err.Error())
	}
	ll = new(Logger)
	ll.rotate = rotate
	ll.file = f
	ll.level = make(chan int)
	go ll.loop()
	ll.SetLevel(lvl)
	return ll
}

func (l *Logger) SetLevel(lvl int) {
	l.level <- lvl
}

func (l *Logger) setLevel(f *os.File, lvl int) {
	switch {
	case lvl > 5:
		l.debug = log.New(f, "\033[0;36mDEBUG:\033[0m ", log.LstdFlags|log.Lshortfile)
		fallthrough
	case lvl > 4:
		l.info = log.New(f, "INFO: ", log.LstdFlags|log.Lshortfile)
		fallthrough
	case lvl > 3:
		l.notic = log.New(f, "\033[0;32mNOTIC:\033[0m ", log.LstdFlags|log.Lshortfile)
		fallthrough
	case lvl > 2:
		l.warn = log.New(f, "\033[0;35mWARN:\033[0m ", log.LstdFlags|log.Lshortfile)
		fallthrough
	case lvl > 1:
		l.err = log.New(f, "\033[0;31mERROR:\033[0m ", log.LstdFlags|log.Lshortfile)
		fallthrough
	case lvl > 0:
		l.fatal = log.New(f, "\033[0;33mFATAL:\033[0m ", log.LstdFlags|log.Lshortfile)
	}
	switch {
	case lvl < 1:
		l.fatal = nil
		fallthrough
	case lvl < 2:
		l.err = nil
		fallthrough
	case lvl < 3:
		l.warn = nil
		fallthrough
	case lvl < 4:
		l.notic = nil
		fallthrough
	case lvl < 5:
		l.info = nil
		fallthrough
	case lvl < 6:
		l.debug = nil
	}
}

func (l *Logger) getFileSize() ByteSize {
	l.rwm.RLock()
	defer l.rwm.RUnlock()
	fi, err := l.file.Stat()
	if err != nil {
		Warn("get log file size failed, no trunc %s", err.Error())
		return 0.0
	}
	return ByteSize(fi.Size())
}

func (l *Logger) trunc(fp, ext string, lvl int) {
	l.rwm.Lock()
	defer l.rwm.Unlock()
	err := l.file.Close()
	if err != nil {
		stdWarn.Println("fail to close log file", err.Error())
		return
	}
	err = os.Rename(fp, fp+ext)
	if err != nil {
		stdWarn.Println("fail to rename log file, no trunc", err.Error())
	}
	f, err := os.OpenFile(fp, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		Error("create log file failed %s", err.Error())
		return
	}
	l.setLevel(f, lvl)
	l.file = f
}

func suffix(t time.Time) string {
	y, m, d := t.Date()
	return "-" + fmt.Sprintf("%04d%02d%02d%02d", y, m, d, t.Hour())
}

func toNextBound(d time.Duration) time.Duration {
	return time.Now().Truncate(d).Add(d).Sub(time.Now())
}

func (l *Logger) loop() error {
	interval := time.After(toNextBound(l.rotate.Interval))
	expired := time.After(checkExpired)
	var sizeExt, lvl int = 1, 4
	fp, err := filepath.Abs(l.file.Name())
	if err != nil {
		stdFatal.Fatalln("get log filepath failed %s", err.Error())
	}
	for {
		var size <-chan time.Time
		if toNextBound(l.rotate.Interval) != checkSize {
			size = time.After(checkSize)
		}
		select {
		case lvl = <-l.level:
			l.rwm.Lock()
			l.setLevel(l.file, lvl)
			l.rwm.Unlock()
			Notic("log level change to %d", lvl)
		case t := <-interval:
			interval = time.After(l.rotate.Interval)
			l.trunc(fp, suffix(t), lvl)
			sizeExt = 1
			Notic("log truncated by time interval")
		case <-expired:
			expired = time.After(checkExpired)
			err := filepath.Walk(filepath.Dir(fp),
				func(path string, info os.FileInfo, err error) error {
					if err != nil {
						return nil
					}
					isLog := strings.Contains(info.Name(), ".log")

					//log.Println("strings.Contains(", info.Name(), " log') isLog = ", isLog)
					if time.Since(info.ModTime()) > l.rotate.Expired && isLog && info.IsDir() == false {
						if err := os.Remove(path); err != nil {
							return err
						}
						Notic("remove expired log files %s", filepath.Base(path))
					}
					return nil
				})
			if err != nil {
				Warn("remove expired logs failed %s", err.Error())
			}
		case t := <-size:
			if l.getFileSize() < l.rotate.Size {
				break
			}
			l.trunc(fp, suffix(t)+"."+strconv.Itoa(sizeExt), lvl)
			sizeExt++
			Notic("log over size, truncated")
		}
	}
}

// Debug log debug message with cyan color.
func Debug(format string, v ...interface{}) {
	ll.rwm.RLock()
	if ll.debug != nil {
		ll.debug.Output(2, fmt.Sprintf(format, v...))
	}
	ll.rwm.RUnlock()
}

// Info log normal message.
func Info(format string, v ...interface{}) {
	ll.rwm.RLock()
	if ll.info != nil {
		ll.info.Output(2, fmt.Sprintf(format, v...))
	}
	ll.rwm.RUnlock()
}

// Notice log notice message with blue color.
func Notic(format string, v ...interface{}) {
	ll.rwm.RLock()
	if ll.notic != nil {
		ll.notic.Output(2, fmt.Sprintf(format, v...))
	}
	ll.rwm.RUnlock()
}

// Error log error message with red color.
func Error(format string, v ...interface{}) {
	atomic.AddInt32(&ll.errCount, 1)
	stdError.Output(2, fmt.Sprintf(format, v...))
	ll.rwm.RLock()
	if ll.err != nil {
		ll.err.Output(2, fmt.Sprintf(format, v...))
	}
	ll.rwm.RUnlock()
}

func ErrCount() int32 {
	ec := atomic.LoadInt32(&ll.errCount)
	if ec < 0 {
		Warn("error count overflow")
		return -1
	}
	return ec
}

func Fatal(format string, v ...interface{}) {
	stdFatal.Output(2, fmt.Sprintf(format, v...))
	ll.rwm.RLock()
	if ll.fatal != nil {
		ll.fatal.Output(2, fmt.Sprintf(format, v...))
	}
	ll.rwm.RUnlock()
	os.Exit(1)
}

func Warn(format string, v ...interface{}) {
	stdWarn.Output(2, fmt.Sprintf(format, v...))
	ll.rwm.RLock()
	if ll.warn != nil {
		ll.warn.Output(2, fmt.Sprintf(format, v...))
	}
	ll.rwm.RUnlock()
}

func GetLogger() *Logger {
	return ll
}
