package query

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"math/rand"
	"net"
	"strconv"
	"time"
)

type Request struct {
	con         *net.UDPConn
	readTimeout time.Duration
	sessionID   [4]byte

	challengeTokenCache   [4]byte
	challengeTokenExpires time.Time
}

func NewRequest() *Request {
	req := &Request{}
	return req
}

func (req *Request) Connect(hostaddr string) error {
	addr, err := net.ResolveUDPAddr("udp4", hostaddr)

	if err != nil {
		return errors.New("error resolving host: " + err.Error())
	}
	req.con, err = net.DialUDP("udp4", nil, addr)

	if err != nil {
		return errors.New("error dialing udp4: " + err.Error())
	}
	req.readTimeout = 5000
	req.generateSessionID()

	return nil
}

func (req *Request) getChallengeToken() ([]byte, error) {
	if req.con == nil {
		return nil, errors.New("no connection, call Request.Connect first")
	}

	reqBuf := [7]byte{0xFE, 0xFD, 0x09}
	copy(reqBuf[3:], req.sessionID[0:])
	req.con.Write(reqBuf[:])

	resBuf, err := req.readWithDeadline()

	if err != nil {
		return nil, err
	}
	err = req.verifyResponseHeader(resBuf)

	if err != nil {
		return nil, err
	}
	res, err := resBuf.ReadBytes(0x00)

	if err != nil {
		return nil, errors.New("malformed challenge response")
	}
	tokenString := string(res[:len(res)-1])
	tokenInt, err := strconv.ParseInt(tokenString, 10, 32)

	if err != nil {
		return nil, errors.New("malformed challenge response")
	}

	tokenBuf := &bytes.Buffer{}
	binary.Write(tokenBuf, binary.BigEndian, tokenInt)
	tokenBytes := tokenBuf.Bytes()
	return tokenBytes[len(tokenBytes)-4:], nil
}

func (req *Request) verifyResponseHeader(input *bytes.Buffer) error {
	var buf [5]byte
	var err error
	var bytesRead int

	bytesRead, err = input.Read(buf[:1])

	if err != nil || bytesRead != 1 || (buf[0] != 0x00 && buf[0] != 0x09) {
		return errors.New("invalid response header")
	}
	bytesRead, err = input.Read(buf[1:])

	if err != nil || bytesRead != 4 {
		return errors.New("invalid response header")
	}

	if bytes.Compare(buf[1:], req.sessionID[0:]) != 0 {
		return errors.New("invalid server sessionID")
	}

	return nil
}

func (req *Request) SetReadTimeout(timeout time.Duration) {
	req.readTimeout = timeout
}

func (req *Request) readWithDeadline() (*bytes.Buffer, error) {
	var buf [2048]byte
	var res = &bytes.Buffer{}
	defer req.con.SetDeadline(time.Time{})

	for {
		req.con.SetDeadline(time.Now().Add(req.readTimeout * time.Millisecond))
		bytes, err := req.con.Read(buf[0:])

		if bytes > 0 {
			res.Write(buf[:bytes])
		}

		if err == io.EOF || bytes < 2048 {
			break
		}

		if bytes == 0 && err != io.EOF {
			return nil, errors.New("timeout exceeded when reading from server (" + err.Error() + ")")
		}
	}
	return res, nil
}

func (req *Request) generateSessionID() {
	var buf [4]byte

	rand.Seed(time.Now().UTC().UnixNano())
	rand.Read(buf[0:])

	for i := 0; i < 4; i++ {
		buf[i] = buf[i] & 0x0F
	}

	req.sessionID = buf
}

func scanDelimittedResponse(input []byte, eof bool) (adv int, token []byte, err error) {
	if len(input) == 0 {
		return 0, nil, errors.New("end of input")
	}

	i := bytes.Index(input, []byte{0x00})
	return i + 1, input[:i], nil
}
