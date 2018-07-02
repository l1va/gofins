package fins

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"sync"
)

// Client Omron FINS client
type Client struct {
	conn net.Conn
	resp []chan response
	sync.Mutex
	dst FinsAddr
	src FinsAddr
	sid byte
}

type response struct {
	sid  byte
	data []byte
}

// NewClient creates a new Omron FINS client
func NewClient(plcAddr string, dst FinsAddr, src FinsAddr) *Client {
	c := new(Client)
	c.dst = dst
	c.src = src
	conn, err := net.Dial("udp", plcAddr)

	if err != nil {
		log.Fatal(err)
		panic(fmt.Sprintf("error resolving UDP port: %s\n", plcAddr))
	}
	c.conn = conn
	c.resp = make([]chan response, 256) //storage for all responses, sid is byte - only 256 values
	go c.listenLoop()

	return c

}

// CloseConnection closes an Omron FINS connection
func (c *Client) CloseConnection() {
	c.conn.Close()
}

func (c *Client) incrementSid() byte {
	c.Lock() //thread-safe sid incrementation
	c.sid++
	sid := c.sid
	c.Unlock()
	c.resp[sid] = make(chan response) //clearing cell of storage for new response
	return sid
}

// ReadD reads from the PLC data area
func (c *Client) ReadData(startAddr uint16, readCount uint16) ([]uint16, error) {
	sid := c.incrementSid()
	cmd := readDCommand(defaultHeader(c.dst, c.src, sid), startAddr, readCount)
	return c.read(sid, cmd)
}

// ReadD reads from the PLC data area
func (c *Client) ReadDataBytes(startAddr uint16, readCount uint16) ([]byte, error) {
	d, err := c.ReadData(startAddr, readCount)
	bs := make([]byte, len(d)*2)
	for i, b := range d {
		binary.BigEndian.PutUint16(bs[i*2:i*2+2], b)
	}
	return bs, err
}

// ReadD reads from the PLC data area
func (c *Client) ReadDataString(startAddr uint16, readCount uint16) (string, error) {
	d, err := c.ReadDataBytes(startAddr, readCount)
	n := bytes.Index(d, []byte{0})
	s := string(d[:n])
	return s, err
}

// ReadW reads from the PLC work area
func (c *Client) ReadWork(startAddr uint16, readCount uint16) ([]uint16, error) {
	sid := c.incrementSid()
	cmd := readWCommand(defaultHeader(c.dst, c.src, sid), startAddr, readCount)
	return c.read(sid, cmd)
}

func (c *Client) read(sid byte, cmd []byte) ([]byte, error) {
	_, err := c.conn.Write(cmd)
	if err != nil {
		return nil, err
	}

	ans := <-c.resp[sid]
	return ans.data, nil
}

// WriteD writes to the PLC data area
func (c *Client) WriteData(startAddr uint16, data []uint16) error {
	sid := c.incrementSid()
	cmd := writeDCommand(defaultHeader(c.dst, c.src, sid), startAddr, data)
	return c.write(sid, cmd)
}

// WriteW writes to the PLC work area
func (c *Client) WriteWork(startAddr uint16, data []uint16) error {
	sid := c.incrementSid()
	cmd := writeWCommand(defaultHeader(c.dst, c.src, sid), startAddr, data)
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

// ReadDAsync reads from the PLC data area asynchronously
func (c *Client) ReadDataAsync(startAddr uint16, readCount uint16, callback func(resp response)) error {
	sid := c.incrementSid()
	cmd := readDCommand(defaultHeader(c.dst, c.src, sid), startAddr, readCount)
	return c.asyncCommand(sid, cmd, callback)
}

// WriteDAsync writes to the PLC data area asynchronously
func (c *Client) WriteDataAsync(startAddr uint16, data []uint16, callback func(resp response)) error {
	sid := c.incrementSid()
	cmd := writeDCommand(defaultHeader(c.dst, c.src, sid), startAddr, data)
	return c.asyncCommand(sid, cmd, callback)
}

func (c *Client) asyncCommand(sid byte, cmd []byte, callback func(resp response)) error {
	_, err := c.conn.Write(cmd)
	if err != nil {
		return err
	}
	asyncResponse(c.resp[sid], callback)
	return nil
}

func asyncResponse(ch chan response, callback func(r response)) {
	if callback != nil {
		go func(ch chan response, callback func(r response)) {
			ans := <-ch
			callback(ans)
		}(ch, callback)
	}
}

// WriteDNoResponse writes to the PLC data area and doesn't request a response
func (c *Client) WriteDataNoResponse(startAddr uint16, data []uint16) error {
	sid := c.incrementSid()
	cmd := writeDCommand(newHeaderNoResponse(c.dst, c.src, sid), startAddr, data)
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
				log.Println("failed to parse response: ", err, " \nresponse: ", buf[0:n])
			} else {
				c.resp[ans.sid] <- *ans
			}
		} else {
			log.Println("cannot read response: ", buf)
		}
	}
}
