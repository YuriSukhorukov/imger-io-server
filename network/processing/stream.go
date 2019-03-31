package processing

import (
	"bufio"
	"fmt"
	"net"
	"../../protocol"
)

func Stream(conn net.Conn) {
	for {
		buff, err := bufio.NewReader(conn).ReadBytes(0x00)
		if err != nil {
			return
		}
		fmt.Printf("%v\n", buff)

		packet := protocol.Unpack(buff)
		PacketsQueueIO[&conn] = append(PacketsQueueIO[&conn], packet)

		// Каждое подключение работает в отдельном потоке
		// Каждый новый пакет из потока подключения попадает в общую очередь пакетов
		// На каждый поток подключения приходитя свой поток обработки пакетов

		// conn.Write(protocol.Pack(packet))
	}
}