package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
)

const (
	// MaxBodySize max proto body size
	MaxBodySize = int32(1 << 12)
)
const (
	// size
	_packSize      = 4
	_headerSize    = 2
	_verSize       = 2
	_opSize        = 4
	_seqSize       = 4
	_heartSize     = 4
	_rawHeaderSize = _packSize + _headerSize + _verSize + _opSize + _seqSize
	_maxPackSize   = MaxBodySize + int32(_rawHeaderSize)
	// offset
	_packOffset   = 0
	_headerOffset = _packOffset + _packSize
	_verOffset    = _headerOffset + _headerSize
	_opOffset     = _verOffset + _verSize
	_seqOffset    = _opOffset + _opSize
	_heartOffset  = _seqOffset + _seqSize
)

var (
	ErrProtoPackLen   = fmt.Errorf("default server codec pack length error")
	ErrProtoHeaderLen = fmt.Errorf("default server codec header length error")
)

func Pop(rd *bufio.Reader, n int) (b []byte, err error) {
	var buf []byte
	//避免拷贝
	if buf, err = rd.Peek(_rawHeaderSize); err != nil {
		return
	}
	//读下标向后挪
	if _, err = rd.Discard(n); err != nil {
		return
	}
	return buf, nil

}

func readTCP(rd *bufio.Reader) (err error) {
	var (
		bodyLen   int
		headerLen int16
		packLen   int32
		buf       []byte
		body      []byte
	)
	if buf, err = Pop(rd, _rawHeaderSize); err != nil {
		return
	}
	packLen = int32(binary.BigEndian.Uint32(buf[_packOffset:]))
	headerLen = int16(binary.BigEndian.Uint16(buf[_headerOffset:]))
	ver := int16(binary.BigEndian.Uint16(buf[_verOffset:]))
	op := int32(binary.BigEndian.Uint32(buf[_opOffset:]))
	seq := int32(binary.BigEndian.Uint32(buf[_seqOffset:]))
	fmt.Printf("var:%d, op:%d, seq:%d\n", ver, op, seq)
	if packLen > _maxPackSize {
		return ErrProtoPackLen
	}
	if headerLen != _rawHeaderSize {
		return ErrProtoHeaderLen
	}
	if bodyLen = int(packLen - int32(headerLen)); bodyLen > 0 {
		body, err = Pop(rd, bodyLen)
	} else {
		body = nil
	}
	//handle body
	_ = body
	return
}

func handleConnection(conn net.Conn) {
	rd := bufio.NewReader(conn)
	for {
		if err := readTCP(rd); err != nil {
			continue
		}
	}

}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		fmt.Println("Listen tcp server failed,err:", err)
		return
	}
	fmt.Println("server running ...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Listen.Accept failed,err:", err)
			continue
		}
		//每个连接创建一个读协程
		go handleConnection(conn)
	}

}
