package lineus

import (
	"io"
	"strconv"
)

type Client struct {
	conn io.ReadWriteCloser
}

type response struct {
	Message []byte
}

func NewClient(conn io.ReadWriteCloser) *Client {
	return &Client{conn}
}

func (c *Client) LinearInterpolation(x float64, y float64, z float64) (response, error) {
	_, err := c.conn.Write(
		[]byte(
			"G01" +
				" X" + strconv.FormatFloat(x, 'f', 4, 64) +
				" Y" + strconv.FormatFloat(y, 'f', 4, 64) +
				" Z" + strconv.FormatFloat(z, 'f', 4, 64) +
				"\u0000",
		),
	)
	if err != nil {
		return response{}, err
	}

	result, err := c.Read()
	if err != nil {
		return response{}, err
	}

	res := response{result}

	return res, nil
}

func (c *Client) Read() ([]byte, error) {
	var result []byte
	buf := make([]byte, 1)

	for {
		n, err := c.conn.Read(buf)
		if err != nil {
			return result, err
		}
		if n == 0 {
			break
		}

		if string(buf[0]) == "\u0000" {
			break
		}

		result = append(result, buf[0])
	}

	return result, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}
