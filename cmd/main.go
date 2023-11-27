package main
import(
"github.com/camzero94/GO/patterns"
)


// Global Type 

func main(){

  // Inject IceSafetyPlace  of type SafetyPlace
  // rc := patterns.NewRockClimber(patterns.IceSafetyPlace{})
  rc := patterns.NewRockClimber(patterns.FireSafetyPlace{})
  for  i := 0; i< 11; i++{
    rc.ClimbRock()
  }
}
