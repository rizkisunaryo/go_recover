package go_recover
import "log"

func Recover(header string) {
	if r := recover(); r != nil {
		log.Println(header, "recover: ", r)
	}
}