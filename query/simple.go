package query

import (
	"bufio"
	"encoding/binary"
	"strconv"
)

type SimpleResponse struct {
	Hostname   string `json:"hostname"`
	GameType   string `json:"gametype"`
	Map        string `json:"map"`
	NumPlayers int    `json:"numplayers"`
	MaxPlayers int    `json:"maxplayers"`
	HostPort   int16  `json:"hostport"`
	HostIP     string `json:"hostip"`
}

func (req *Request) Simple() (*SimpleResponse, error) {
	response := &SimpleResponse{}
	challengeToken, err := req.getChallengeToken()

	if err != nil {
		return nil, err
	}

	reqBuf := [11]byte{0xFE, 0xFD}
	copy(reqBuf[3:], req.sessionID[0:])
	copy(reqBuf[7:], challengeToken)
	req.con.Write(reqBuf[:])

	resBuf, err := req.readWithDeadline()

	if err != nil {
		return response, err
	}
	err = req.verifyResponseHeader(resBuf)

	if err != nil {
		return response, err
	}

	scan := bufio.NewScanner(resBuf)
	scan.Split(scanDelimittedResponse)

	scan.Scan()
	response.Hostname = scan.Text()

	scan.Scan()
	response.GameType = scan.Text()
	
	scan.Scan()
	response.Map = scan.Text()

	scan.Scan()
	response.NumPlayers, _ = strconv.Atoi(scan.Text())

	scan.Scan()
	response.MaxPlayers, _ = strconv.Atoi(scan.Text())

	scan.Scan()
	portAndIP := scan.Bytes()
	response.HostPort = int16(binary.LittleEndian.Uint16(portAndIP[:2]))
	response.HostIP = string(portAndIP[2:])

	return response, nil
}