package client

//import (
//	"encoding/binary"
//	"encoding/json"
//	"errors"
//	"example.com/m/common/message"
//	"fmt"
//	"net"
//)
//
////ReadPkg 封装读取客户端发送的消息
//func ReadPkg(conn net.Conn) (mes message.Message,err error)  {
//	buf := make([]byte,8096)
//	fmt.Println("conn start read ")
//	//在conn没有被关闭的情况下 才会阻塞
//	_,err = conn.Read(buf[:4])
//	if  err != nil {
//		//err = errors.New("read pakLen fail error")
//		return
//	}
//
//	//根据buf[:4] 转成一个 uint32
//	var pkgLen uint32
//	pkgLen = binary.BigEndian.Uint32(buf[:4])
//
//	//根据 pkgLen读取数据 从conn都pkgLen个字节到buf里面去
//	n,err :=conn.Read(buf[:pkgLen])
//	if uint32(n) != pkgLen || err != nil {
//		//err = errors.New("conn read dataJson error")
//		return
//	}
//
//	//反序列化 -》 message。Message类型
//	err =json.Unmarshal(buf[:pkgLen],&mes)
//	if err != nil {
//		err = errors.New("dataJson unMarshal error")
//		return
//	}
//	return
//}
//
//
////WritePkg 写入信息
//func WritePkg(conn net.Conn, data []byte) (err error) {
//	//先发送一个长度给对方
//	var dataLen uint32
//	dataLen = uint32(len(data))
//	var buf [4]byte
//	binary.BigEndian.PutUint32(buf[0:4],dataLen)
//
//	//发送长度
//	n,err := conn.Write(buf[:4])
//	if n != 4 || err != nil{
//		log.Fatalln("conn.Write err = ",err)
//		return
//	}
//
//	//发送data 本身
//	n,err = conn. Write(data)
//	if uint32(n) != dataLen || err != nil{
//		log.Fatalln("conn.Write err = ",err)
//		return
//	}
//	return
//}

