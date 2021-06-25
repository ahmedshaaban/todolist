package db

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type DB struct {
	filename string
}

func New(filename string) *DB {
	return &DB{
		filename: filename,
	}
}

// GetAll return all the saved data as an array of an array filtered by active
func (db DB) GetAll() ([][]string, error) {
	f, err := os.Open(db.filename)
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(f)
	lines, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	result := [][]string{}
	for _, l := range lines {
		// checks the active status
		if l[2] == "true" {
			result = append(result, l)
		}
	}

	return result, nil
}

// Add the input to the csv file, also get the last index and increment it (acts as ID).
func (db DB) Add(input []string) error {
	// read the file
	f, err := os.Open(db.filename)
	if err != nil {
		return err
	}

	r := csv.NewReader(f)
	lines, err := r.ReadAll()
	if err != nil {
		return err
	}

	lastIndexString := "0"

	// check if the file is empty
	if len(lines) > 0 {
		lastIndexString = lines[len(lines)-1][0]
	}

	lastIndex, err := strconv.ParseInt(lastIndexString, 0, 32)
	if err != nil {
		return err
	}

	input = append([]string{fmt.Sprintf("%d", lastIndex+1)}, input...)
	lines = append(lines, input)

	// write the file
	f, err = os.Create(db.filename)
	if err != nil {
		return err
	}
	w := csv.NewWriter(f)
	if err = w.WriteAll(lines); err != nil {
		f.Close()
		return err
	}
	return f.Close()
}

// UpdateStatus updates the line based on the index to active or disabled
func (db DB) UpdateStatus(i int, active bool) error {
	f, err := os.Open(db.filename)
	if err != nil {
		return err
	}

	r := csv.NewReader(f)
	lines, err := r.ReadAll()
	if err != nil {
		return err
	}

	if i-1 >= len(lines) {
		return errors.New("index out of range")
	}

	if lines[i-1][3] == "true" {
		return errors.New("index not found")
	}

	lines[i-1][2] = fmt.Sprintf("%t", active)

	// write the file
	f, err = os.Create(db.filename)
	if err != nil {
		return err
	}
	w := csv.NewWriter(f)
	if err = w.WriteAll(lines); err != nil {
		f.Close()
		return err
	}

	return f.Close()
}

// DeleteDisabled removes lines with active = false
func (db DB) DeleteDisabled() error {
	f, err := os.Open(db.filename)
	if err != nil {
		return err
	}

	r := csv.NewReader(f)
	lines, err := r.ReadAll()
	if err != nil {
		return err
	}

	for _, l := range lines {
		// check if active is false then mark as deleted
		if l[2] == "false" {
			l[3] = "true"
		}
	}

	// write the file
	f, err = os.Create(db.filename)
	if err != nil {
		return err
	}
	w := csv.NewWriter(f)
	if err = w.WriteAll(lines); err != nil {
		f.Close()
		return err
	}

	return f.Close()
}
