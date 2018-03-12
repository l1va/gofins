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
	resp []chan Response
	sync.Mutex
	sid  byte
}

type Response struct {
	sid  byte
	Data []uint16
}

func NewClient(plcAddr string) *Client {
	c := new(Client)
	conn, err := net.Dial("udp", plcAddr)
	if err != nil {
		log.Fatal(err)
		panic(fmt.Sprintf("error resolving UDP port: %s\n", plcAddr))
	}
	c.conn = conn
	c.resp = make([]chan Response, 256) //storage for all responses, sid is byte - only 256 values
	go c.listenLoop()

	return c

}

func (c *Client) CloseConnection() {
	c.conn.Close()
}

func (c *Client) incrementSid() byte {
	c.Lock() //thread-safe sid incrementation
	c.sid += 1
	sid := c.sid
	c.Unlock()
	c.resp[sid] = make(chan Response) //clear storage for new response
	return sid
}

func (c *Client) ReadD(startAddr uint16, readCount uint16) ([]uint16, error) {
	sid := c.incrementSid()
	cmd := readDCommand(defaultHeader(sid), startAddr, readCount)
	return c.read(sid, cmd)
}

func (c *Client) ReadW(startAddr uint16, readCount uint16) ([]uint16, error) {
	sid := c.incrementSid()
	cmd := readWCommand(defaultHeader(sid), startAddr, readCount)
	return c.read(sid, cmd)
}

func (c *Client) read(sid byte, cmd []byte) ([]uint16, error) {
	_, err := c.conn.Write(cmd)
	if err != nil {
		return nil, err
	}

	ans := <-c.resp[sid]
	return ans.Data, nil
}

func (c *Client) WriteD(startAddr uint16, data []uint16) error {
	sid := c.incrementSid()
	cmd := writeDCommand(defaultHeader(sid), startAddr, data)
	return c.write(sid, cmd)
}

func (c *Client) WriteW(startAddr uint16, data []uint16) error {
	sid := c.incrementSid()
	cmd := writeWCommand(defaultHeader(sid), startAddr, data)
	return c.write(sid, cmd)
}

func (c *Client) write(sid byte, cmd []byte) error {
	_, err := c.conn.Write(cmd)
	if err != nil {
		return err
	}

	<-c.resp[sid]
	return nil
}

func (c *Client) ReadDAsync(startAddr uint16, readCount uint16, callback func(resp Response)) error {
	sid := c.incrementSid()
	cmd := readDCommand(defaultHeader(sid), startAddr, readCount)
	return c.asyncCommand(sid, cmd, callback)
}

func (c *Client) WriteDAsync(startAddr uint16, data []uint16, callback func(resp Response)) error {
	sid := c.incrementSid()
	cmd := writeDCommand(defaultHeader(sid), startAddr, data)
	return c.asyncCommand(sid, cmd, callback)
}

func (c *Client) asyncCommand(sid byte, cmd []byte, callback func(resp Response)) error {
	_, err := c.conn.Write(cmd)
	if err != nil {
		return err
	}
	asyncResponse(c.resp[sid], callback)
	return nil
}

func asyncResponse(ch chan Response, callback func(r Response)) {
	if callback != nil {
		go func(ch chan Response, callback func(r Response)) {
			ans := <-ch
			callback(ans)
		}(ch, callback)
	}
}

func (c *Client) WriteDNoResponse(startAddr uint16, data []uint16) error {
	sid := c.incrementSid()
	cmd := writeDCommand(newHeaderNoResponse(sid), startAddr, data)
	return c.asyncCommand(sid, cmd, nil)
}

func (c *Client) listenLoop() {
	for {
		buf := make([]byte, 2048)
		n, err := bufio.NewReader(c.conn).Read(buf)
		if err != nil {
			log.Fatal(err)
		}

		if n > 0 {
			ans, err := parseResponse(buf[0:n])
			if err != nil {
				log.Println("failed to parse response: ", err, " \nresponse:", buf[0:n])
			} else {
				c.resp[ans.sid] <- *ans
			}
		} else {
			log.Println("cannot read response: ", buf)
		}
	}
}
