package processing

import (
	"../../protocol"
	"bufio"
	"net"
)

func Stream(conn net.Conn) {
	for {
		buff, err := bufio.NewReader(conn).ReadBytes(0x00)
		if err != nil {
			return
		}

		packet                := protocol.Unpack(buff)
		PacketsChannelIO[conn]  <- packet

		// Каждое подключение работает в отдельном потоке
		// Каждый новый пакет из потока подключения попадает в общую очередь пакетов
		// На каждый поток подключения приходитя свой поток обработки пакетов
	}
}