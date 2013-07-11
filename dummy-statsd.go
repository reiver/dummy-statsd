package main



import (
	"fmt"
	"net"
	"log"
)



func main() {

	// Output intro message.
		fmt.Println("Welcome to the Dummy StatsD server.")
		fmt.Println()
		fmt.Println("To exit, press: [CTRL]+[C]")
		fmt.Println()



	// Specify UDP port to listen to. (The real StatsD listens to UDP port 8125.)
		sAddr, err := net.ResolveUDPAddr("udp", ":8125")
		if err != nil {
			log.Fatalln(err)
		}



	// Listen to UDP port 8125. (This is what the real StatsD server listens to.)	
		udpConn, err := net.ListenUDP("udp", sAddr)
		if nil != err {
			log.Fatalln(err)
		}
		defer udpConn.Close()


		//DEBUG
		fmt.Printf("Dummy StatsD server is listening on UDP %s\n", udpConn.LocalAddr().String())
		fmt.Println()



	// Handle incoming UDP packets.


		//DEBUG
		fmt.Println("Dummy StatsD server is waiting for incoming messages....")
		fmt.Println()

		buffer := make([]byte, 65536)

		for {
			n, err := udpConn.Read(buffer)
			if err != nil {
				log.Fatalln(err)
			}


			//DEBUG
			log.Printf("Received: %s", string(buffer[0:n]))

		} // for

}
