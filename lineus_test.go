package lineus_test

import (
	"bytes"
	"testing"

	lineus "github.com/youpy/go-lineus"
)

type connection struct {
	*bytes.Buffer
}

func TestLinearInterpolation(t *testing.T) {
	client := setupClient(t)
	response, err := client.LinearInterpolation(1000.0, 1200.0, 0.0)
	if err != nil {
		t.Fatal(err)
	}

	expected := "received: G01 X1000.0000 Y1200.0000 Z0.0000"
	actual := string(response.Message)
	if expected != actual {
		t.Errorf("got \"%v\"\nwant \"%v\"", actual, expected)
	}
}

func TestHome(t *testing.T) {
	client := setupClient(t)
	response, err := client.Home()
	if err != nil {
		t.Fatal(err)
	}

	expected := "received: G28"
	actual := string(response.Message)
	if expected != actual {
		t.Errorf("got \"%v\"\nwant \"%v\"", actual, expected)
	}
}

func setupClient(t *testing.T) *lineus.Client {
	buffer := bytes.NewBufferString("received: ")
	conn := &connection{buffer}
	client := lineus.NewClient(conn)

	return client
}
