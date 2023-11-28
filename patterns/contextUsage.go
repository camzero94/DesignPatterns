package patterns

import (
	"context"
	"fmt"
	"time"
  "log"
)

type Response struct{
  mssg string 
  err error
}

func ContextUsage(){

  // Initialize variables 
  start := time.Now()
  ctx := context.Background()
  userId := 10 
  message , err := fetchData(ctx,userId)

  if err != nil{
    log.Fatal(err)
  }

  fmt.Printf("Resulting Message is: %s took %s", message,time.Since(start))

}


// This function will not take longer than the context time of 200 milliseconds 
func fetchData(ctx context.Context, idUser int )(string,error){

  // We establish a deadline so the fetch Function cannot take longer than 200 milliseconds
  ctx, cancel := context.WithTimeout(ctx, time.Millisecond * 200)
  defer cancel()
  respch := make(chan Response)

  // Call inside go rutine the third party fetch function that takes too long
  go func(){
    mssg, err := fetchLongThirdPartyCanBeSlow()
    respch <- Response{
      mssg: mssg,
      err: err,
    }
  }()

  // Loop until Context is done or recieve message   from third party function
  for {
    select{
    // When context finish enter inside this block
    case <- ctx.Done():
      return " ", fmt.Errorf("Takes too long")
    case resp := <- respch:
      return resp.mssg,resp.err
    }
  }
}

// This fetch function takes to long
func fetchLongThirdPartyCanBeSlow()(string,error) {
  time.Sleep(100*time.Millisecond)
  return fmt.Sprintf("Info Returned"),nil
}
