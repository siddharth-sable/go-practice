package main
import(
	"fmt"
	"time"
)

func main(){
	start := time.Now()
	evilNinjas := []string{"tommy","Andy","bobby","Johnny"}

	defer func(){
		fmt.Println(time.Since(start))
	}()

	for _, evilNinja := range(evilNinjas){
		go attack(evilNinja)
	}

	time.Sleep(time.Second)
	
}

func attack(target string){
	fmt.Println(target)
}