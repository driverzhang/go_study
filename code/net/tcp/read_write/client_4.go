package main
import(
	"log"
	"net"
	"time"
)

func main(){
	log.Println("begin dial...")
	conn, err := net.Dial("tcp", ":8899")
	if err != nil {
		log.Println("dial error:", err)
		return
	}

	defer conn.Close()
	log.Println("dial ok")
	
	data := make([]byte, 65536)
	conn.Write(data)
	time.Sleep(time.Second * 1000) 
}