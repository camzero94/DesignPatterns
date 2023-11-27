package patterns

import (
  "fmt"
)

// Injecting the SafetyPlace Interface type 
//Type rocksClimbed never change implementation onl Inject SafetyPlace
type RockClimber struct {
  rocksClimbed int 
  kind int 
  // Dependanc but Behaviorial Dependancy
  sp SafetyPlace
}

func NewRockClimber(sp SafetyPlace) *RockClimber{
  return &RockClimber{
    sp: sp,
  }
}




// Class that make together similar algorithms 
type SafetyPlace interface {
  PlaceSafety()
}

//Implementation Fire
type FireSafetyPlace struct{
  //api
  //database
  //...
}
func (i FireSafetyPlace) PlaceSafety(){
  fmt.Println("placing my safeties Fire...")
}

//Implementation for Ice 
type IceSafetyPlace struct{
  //api
  //database
  //...
} 
func (i IceSafetyPlace) PlaceSafety(){
  fmt.Println("placing my safeties Ice ...")
}


// Function that implements the SafetyPlace type
func (rc *RockClimber) ClimbRock(){
  rc.rocksClimbed ++
  if rc.rocksClimbed == 10{
    rc.sp.PlaceSafety()
    
  }
}


// // What if we need multiple types of rocks 
// func (rc *RockClimber) PlaceSafety(){
//   switch rc.kind{
//     //Ice
//     case 1:
//       fmt.Println("placing my safeties Ice ...")
//     //Sand mybe this one needs database access  
//     case 2:
//       fmt.Println("placing my safeties Sand...")
//     //Fire
//     case 3:
//       fmt.Println("placing my safeties Fire...")
//     }
//
// }






