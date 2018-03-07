package fins

import (
	"log"
	"net"
	"fmt"
	"bufio"
	"sync"
)

type Client struct {
	conn net.Conn
	resp []chan Answer
	sync.Mutex
	sid  byte
}

type Answer struct {
	sid  byte
	data []uint16
}

func NewClient(plcAddr string) *Client {
	c := new(Client)
	conn, err := net.Dial("udp", plcAddr)
	if err != nil {
		log.Fatal(err)
		panic(fmt.Sprintf("error resolving UDP port: %s\n", plcAddr))
	}
	c.conn = conn
	c.resp = make([]chan Answer, 256)
	go c.listenLoop()

	return c

}

func (c *Client) CloseConnection() {
	c.conn.Close()
}
func (c *Client) incrementSid()byte{
	c.Lock()
	c.sid+=1
	sid:= c.sid
	c.Unlock()
	return sid
}

func (c *Client) ReadDM(startAddr uint16, readCount uint16) ([]uint16, error) {
	sid := c.incrementSid()
	c.resp[sid] = make(chan Answer)

	_, err := c.conn.Write(readDMCommand(newHeader(sid), startAddr, readCount))
	if err != nil {
		return nil, err
	}

	ans := <-c.resp[sid]
	return ans.data, nil
}

func (c *Client) WriteDM(startAddr uint16, data []uint16) error {
	sid := c.incrementSid()
	c.resp[sid] = make(chan Answer)

	_, err := c.conn.Write(writeDMCommand(newHeader(sid), startAddr, data))
	if err != nil {
		return err
	}

	<-c.resp[sid]
	return nil
}

func (c *Client) listenLoop() {
	for {
		buf := make([]byte, 1024) // is it enough?
		n, err := bufio.NewReader(c.conn).Read(buf)
		if err != nil {
			log.Fatal(err)
		}

		if n > 0 {
			ans, err := parseAnswer(buf[0:n])
			if err != nil {
				log.Println("failed to parse response: ", buf)
			} else {
				c.resp[ans.sid] <- *ans
			}
		} else {
			log.Println("cannot read response: ", buf)
		}
	}
}