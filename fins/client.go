package fins

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

// Client Omron FINS client
type Client struct {
	conn net.Conn
	resp []chan Frame
	sync.Mutex
	dst Address
	src Address
	sid byte
}

// NewClient creates a new Omron FINS client
func NewClient(plcAddr string, dst Address, src Address) *Client {
	c := new(Client)
	c.dst = dst
	c.src = src
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

// ReadBits reads from the PLC data area
// func (c *Client) ReadBits(memoryArea byte, address uint16, bitOffset byte, readCount uint16) ([]bool, error) {
// 	header := c.nextHeader()
// 	command := readCommand(IOAddress{
// 		MemoryArea: memoryArea,
// 		Address:    address,
// 		BitOffset:  bitOffset,
// 	}, readCount)
// 	r, e := c.sendCommand(header, command)
// 	if e != nil {
// 		return nil, e
// 	}

// 	data := make([]uint16, readCount, readCount)
// 	for i := 0; i < int(readCount); i++ {
// 		data[i] = binary.BigEndian.Uint16(r.Data[i*2 : i*2+2])
// 	}

// 	return data, nil
// }

// ReadWords reads from the PLC data area
func (c *Client) ReadWords(memoryArea byte, address uint16, readCount uint16) ([]uint16, error) {
	header := c.nextHeader()
	command := readCommand(IOAddress{
		MemoryArea: memoryArea,
		Address:    address,
		BitOffset:  0x00,
	}, readCount)
	r, e := c.sendCommand(header, command)
	if e != nil {
		return nil, e
	}

	data := make([]uint16, readCount, readCount)
	for i := 0; i < int(readCount); i++ {
		data[i] = binary.BigEndian.Uint16(r.Data[i*2 : i*2+2])
	}

	return data, nil
}

// ReadString reads from the PLC data area
func (c *Client) ReadString(memoryArea byte, address uint16, readCount uint16) (*string, error) {
	header := c.nextHeader()
	command := readCommand(IOAddress{
		MemoryArea: memoryArea,
		Address:    address,
		BitOffset:  0x00,
	}, readCount)
	r, e := c.sendCommand(header, command)
	if e != nil {
		return nil, e
	}

	n := bytes.Index(r.Data, []byte{0})
	s := string(r.Data[:n])
	return &s, nil
}

// ReadClock reads the PLC clock
func (c *Client) ReadClock() (*time.Time, error) {
	header := c.nextHeader()
	command := new(Payload)
	command.CommandCode = CommandCodeClockRead
	command.Data = []byte{}
	r, e := c.sendCommand(header, command)
	if e != nil {
		return nil, e
	}
	year, _ := decodeBCD(r.Data[0:1])
	if year < 50 {
		year += 2000
	} else {
		year += 1900
	}
	month, _ := decodeBCD(r.Data[1:2])
	day, _ := decodeBCD(r.Data[2:3])
	hour, _ := decodeBCD(r.Data[3:4])
	minute, _ := decodeBCD(r.Data[4:5])
	second, _ := decodeBCD(r.Data[5:6])

	t := time.Date(
		int(year), time.Month(month), int(day), int(hour), int(minute), int(second),
		0, // nanosecond
		time.Local,
	)

	return &t, nil
}

func (c *Client) sendCommand(header *Header, payload *Payload) (*Response, error) {
	bytes := encodeFrame(NewFrame(header, payload))
	_, err := c.conn.Write(bytes)
	if err != nil {
		return nil, err
	}

	r := <-c.resp[header.sid]
	response := &Response{
		CommandCode: r.Payload.CommandCode,
		EndCode:     binary.BigEndian.Uint16(r.Payload.Data[:2]),
		Data:        r.Payload.Data[2:],
	}
	return response, nil
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

func (c *Client) nextHeader() *Header {
	sid := c.incrementSid()
	header := defaultHeader(c.dst, c.src, sid)
	return header
}

func (c *Client) incrementSid() byte {
	c.Lock() //thread-safe sid incrementation
	c.sid++
	sid := c.sid
	c.Unlock()
	c.resp[sid] = make(chan Frame) //clearing cell of storage for new response
	return sid
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
