package patterns

import (
  "fmt"
  "time"
  "os"
)

// Program that start counting for launching Rocket, if cancelled with ctrl-c will abort

func Multiplexing(){

  //...create Abort Channel
  fmt.Println("Commencing countdown Press return to Abort")

  //Function that tries to read a single byte from standard input
  abort := make(chan struct{})
  tick := time.Tick(1 * time.Second)

  go func (){
    os.Stdin.Read(make([]byte,1))
    abort <- struct{} {}
  }()

  for countdown := 10; countdown> 0; countdown--{
    fmt.Printf("Launch in: %d",countdown)
    select {
      case <- tick:
        //Do nothing
      //If Abort signal is recieved
      case <- abort:
        fmt.Println("Launch aborted")
        return
      }
  } 
  fmt.Println("Launch")

}

