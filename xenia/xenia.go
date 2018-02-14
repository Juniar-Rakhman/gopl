package main

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Data struct {
	Line string
}

type Xenia struct {
	Host    string
	Timeout time.Duration
}

type Pillar struct {
	Host    string
	Timeout time.Duration
}

type Fiqs struct {
	Host    string
	Timeout time.Duration
}

type System struct {
	Puller
	Storer
}

type Puller interface {
	pull(*Data) error
}

type Storer interface {
	store(*Data) error
}

type Oracle struct {
	Host    string
	Timeout time.Duration
}

type PullStorer interface {
	Puller
	Storer
}

func (*Xenia) pull(d *Data) error {
	switch rand.Intn(20) {
	case 1, 9:
		return io.EOF
	case 5:
		return errors.New("Error reading data from xenia")
	default:
		d.Line = "Data from xenia"
		fmt.Println("In: ", d.Line)
		return nil
	}
}

func (*Pillar) store(d *Data) error {
	fmt.Println("Out:", d.Line)
	return nil
}

func (*Fiqs) pull(d *Data) error {
	switch rand.Intn(20) {
	case 1:
		return io.EOF
	case 5:
		return errors.New("Error reading data from Fiqs")
	default:
		d.Line = "Data from fiqs"
		fmt.Println("In: ", d.Line)
		return nil
	}
}

func (*Oracle) store(d *Data) error {
	fmt.Println("Storing data to oracle:", d.Line)
	return nil
}

func pull(p Puller, data []Data) (int, error) {
	for i := range data {
		if err := p.pull(&data[i]); err != nil {
			return i, err
		}
	}
	return len(data), nil
}

func store(s Storer, data []Data) (int, error) {
	for i := range data {
		if err := s.store(&data[i]); err != nil {
			return i, err
		}
	}
	return len(data), nil
}

func copyDB(ps PullStorer, batch int) error {
	data := make([]Data, batch)

	for {
		i, err := pull(ps, data)

		fmt.Println("Number of data to copy : ", i)

		if i > 0 {
			if _, err := store(ps, data[:i]); err != nil {
				return err
			}
		}

		if err != nil {
			return err
		}
	}
}

func main() {
	sys := System{
		Puller: &Xenia{
			Host:    "localhost:8000",
			Timeout: time.Second,
		},
		Storer: &Pillar{
			Host:    "localhost:9000",
			Timeout: time.Second,
		},
	}

	if err := copyDB(&sys, 10); err != io.EOF {
		fmt.Println(err)
	}

	fmt.Println("====== Copy from xenia to pillar done =====")

	sys2 := System{
		Puller: &Fiqs{
			Host:    "localhost:8000",
			Timeout: time.Second,
		},
		Storer: &Oracle{
			Host:    "localhost:9000",
			Timeout: time.Second,
		},
	}

	if err := copyDB(&sys2, 10); err != io.EOF {
		fmt.Println(err)
	}


}
