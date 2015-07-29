package main

import (
	"fmt"
	"os"
	"sync"
)

type ReOpeningWriter struct {
	fp *os.File

	lock     sync.Mutex
	filename string // should be set to the actual filename
}

// Make a new ReOpeningWriter. Return nil if error occurs during setup.
func NewReOpeningWriter(filename string) *ReOpeningWriter {
	w := ReOpeningWriter{filename: filename}
	err := w.ReOpen()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &w
}

// Write satisfies the io.Writer interface.
func (w *ReOpeningWriter) Write(output []byte) (int, error) {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.CheckReOpen()
	return w.fp.Write(output)
}

// Perform the actual act of rotating and reopening file.
func (w *ReOpeningWriter) ReOpen() (err error) {
	// Close existing file if open
	if w.fp != nil {
		err = w.fp.Close()
		w.fp = nil
		if err != nil {
			return
		}
	}

	// Open file for appending
	w.fp, err = os.OpenFile(w.filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	return
}

func (w *ReOpeningWriter) CheckReOpen() (err error) {
	// Reopen dest file if it has been rotated
	if _, err := os.Stat(w.filename); os.IsNotExist(err) {
		err = w.ReOpen()
	}
	return err
}
