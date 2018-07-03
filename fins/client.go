package fins

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"
)

// Client Omron FINS client
type Client struct {
	conn net.Conn
	resp []chan Frame
	sync.Mutex
	header *Header
}

// NewClient creates a new Omron FINS client
func NewClient(plcAddr string, dst Address, src Address) *Client {
	c := new(Client)
	c.header = defaultHeader(dst, src, 0)
	conn, err := net.Dial("udp", plcAddr)

	if err != nil {
		log.Fatal(err)
		panic(fmt.Sprintf("error resolving UDP port: %s\n", plcAddr))
	}
	c.conn = conn
	c.resp = make([]chan Frame, 256) //storage for all responses, sid is byte - only 256 values
	go c.listenLoop()

	return c
}

// CloseConnection closes an Omron FINS connection
func (c *Client) CloseConnection() {
	c.conn.Close()
}

// ReadBytes Reads from the PLC data area
// func (c *Client) ReadBytes(memoryArea byte, address uint16, readCount uint16) ([]byte, error) {
// 	sid := c.incrementSid()
// 	cmd := readCommand(IOAddress{
// 		MemoryArea: memoryArea,
// 		Address:    address,
// 		BitOffset:  0x00,
// 	}, readCount)
// 	return c.read(c.header.sid, cmd)
// }

// ReadWord reads from the PLC data area
// func (c *Client) ReadWord(memoryArea byte, address uint16) (uint16, error) {
// 	w, e := c.ReadWords(memoryArea, address, 1)
// 	return w[0], e
// }

// ReadWords reads from the PLC data area
func (c *Client) ReadWords(memoryArea byte, address uint16, readCount uint16) (*Payload, error) {
	c.incrementSid()
	cmd := readCommand(IOAddress{
		MemoryArea: memoryArea,
		Address:    address,
		BitOffset:  0x00,
	}, readCount)
	return c.sendCommand(cmd)
}

// ReadString reads from the PLC data area
// func (c *Client) ReadString(memoryArea byte, address uint16, readCount uint16) (string, error) {
// 	d, err := c.ReadDataBytes(startAddr, readCount)
// 	n := bytes.Index(d, []byte{0})
// 	s := string(d[:n])
// 	return s, err
// }

func (c *Client) sendCommand(payload *Payload) (*Payload, error) {
	bytes := encodeFrame(NewFrame(c.header, payload))
	_, err := c.conn.Write(bytes)
	if err != nil {
		return nil, err
	}

	ans := <-c.resp[c.header.sid]
	return ans.Payload, nil
}

// WriteData writes to the PLC data area
// func (c *Client) WriteData(startAddr uint16, data []uint16) error {
// 	sid := c.incrementSid()
// 	cmd := writeDCommand(defaultHeader(c.dst, c.src, sid), startAddr, data)
// 	return c.write(sid, cmd)
// }

// WriteWork writes to the PLC work area
// func (c *Client) WriteWork(startAddr uint16, data []uint16) error {
// 	sid := c.incrementSid()
// 	cmd := writeWCommand(defaultHeader(c.dst, c.src, sid), startAddr, data)
// 	return c.write(sid, cmd)
// }

// func (c *Client) write(sid byte, cmd []byte) error {
// 	_, err := c.conn.Write(cmd)
// 	if err != nil {
// 		return err
// 	}

// 	<-c.resp[sid]
// 	return nil
// }

// ReadDataAsync reads from the PLC data area asynchronously
// func (c *Client) ReadDataAsync(startAddr uint16, readCount uint16, callback func(resp response)) error {
// 	sid := c.incrementSid()
// 	cmd := readDCommand(defaultHeader(c.dst, c.src, sid), startAddr, readCount)
// 	return c.asyncCommand(sid, cmd, callback)
// }

// WriteDataAsync writes to the PLC data area asynchronously
// func (c *Client) WriteDataAsync(startAddr uint16, data []uint16, callback func(resp response)) error {
// 	sid := c.incrementSid()
// 	cmd := writeDCommand(defaultHeader(c.dst, c.src, sid), startAddr, data)
// 	return c.asyncCommand(sid, cmd, callback)
// }

// WriteDataNoResponse writes to the PLC data area and doesn't request a response
// func (c *Client) WriteDataNoResponse(startAddr uint16, data []uint16) error {
// 	sid := c.incrementSid()
// 	cmd := writeDCommand(newHeaderNoResponse(c.dst, c.src, sid), startAddr, data)
// 	return c.asyncCommand(sid, cmd, nil)
// }

func (c *Client) incrementSid() {
	c.Lock() //thread-safe sid incrementation
	c.header.sid++
	c.Unlock()
	c.resp[c.header.sid] = make(chan Frame) //clearing cell of storage for new response
}

func (c *Client) listenLoop() {
	for {
		buf := make([]byte, 2048)
		n, err := bufio.NewReader(c.conn).Read(buf)
		if err != nil {
			log.Fatal(err)
		}

		if n > 0 {
			ans := decodeFrame(buf[0:n])
			if err != nil {
				log.Println("failed to parse response: ", err, " \nresponse: ", buf[0:n])
			} else {
				c.resp[ans.Header.sid] <- *ans
			}
		} else {
			log.Println("Cannot read response: ", buf)
		}
	}
}

// func (c *Client) asyncCommand(sid byte, cmd []byte, callback func(resp response)) error {
// 	_, err := c.conn.Write(cmd)
// 	if err != nil {
// 		return err
// 	}
// 	asyncResponse(c.resp[sid], callback)
// 	return nil
// }

// func asyncResponse(ch chan response, callback func(r response)) {
// 	if callback != nil {
// 		go func(ch chan response, callback func(r response)) {
// 			ans := <-ch
// 			callback(ans)
// 		}(ch, callback)
// 	}
// }
