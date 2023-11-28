package main
import(
"github.com/camzero94/GO/patterns"
)



func main(){
  patterns.ContextUsage()


}

  // // ============== Dependency Injection ==================
  // // Inject IceSafetyPlace or FireSafetyPlace of type SafetyPlace
  // rc := patterns.NewRockClimber(patterns.FireSafetyPlace{})
  // for  i := 0; i< 11; i++{
  //   rc.ClimbRock()
  // }
