package lineus

import (
	"net"
	"strconv"
)

type Client struct {
	conn *net.TCPConn
}

type response struct {
	Message []byte
}

func NewClient(hostname string) (cli *Client, err error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", hostname)
	if err != nil {
		return
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return
	}

	cli = &Client{conn}

	return
}

func (c *Client) LinearInterpolation(x float64, y float64, z float64) (res response, err error) {
	_, err = c.conn.Write(
		[]byte(
			"G01" +
				" X" + strconv.FormatFloat(x, 'f', 4, 64) +
				" Y" + strconv.FormatFloat(y, 'f', 4, 64) +
				" Z" + strconv.FormatFloat(z, 'f', 4, 64) +
				"\u0000",
		),
	)
	if err != nil {
		return
	}

	result, err := c.Read()
	if err != nil {
		return
	}

	res = response{result}

	return
}

func (c *Client) Read() (result []byte, err error) {
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

	return
}
