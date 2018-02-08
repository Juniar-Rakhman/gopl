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

type System struct {
	Xenia
	Pillar
}

func (*Xenia) pull(d *Data) error {
	switch rand.Intn(10) {
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

func pull(x *Xenia, data []Data) (int, error) {
	for i := range data {
		if err := x.pull(&data[i]); err != nil {
			return i, err
		}
	}
	return len(data), nil
}

func store(p *Pillar, data []Data) (int, error) {
	for i := range data {
		if err := p.store(&data[i]); err != nil {
			return i, err
		}
	}
	return len(data), nil
}

func copy(s *System, batch int) error {
	data := make([]Data, batch)

	for {
		i, err := pull(&s.Xenia, data)
		if err != nil {
			return err
		}
		if i > 0 {
			if store(&s.Pillar, data[:i]); err != nil {
				return err
			}
		}
	}
}

func main() {
	sys := System{
		Xenia: Xenia{
			Host:    "localhost:8000",
			Timeout: time.Second,
		},
		Pillar: Pillar{
			Host:    "localhost:9000",
			Timeout: time.Second,
		},
	}

	if err := copy(&sys, 3); err != io.EOF {
		fmt.Println(err)
	}
}
