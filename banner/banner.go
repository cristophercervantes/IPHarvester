package banner

import "fmt"

const version = "v1.0"

func PrintVersion() {
	fmt.Printf("IPHarvester %s\n", version)
}

func PrintBanner() {
	banner := `
    ___ ____  __  __    __   __           __
   (_  )__  \/ / / /_  / /  / /_  _______/ /_____  _____
  / / / / / / /_  __ \/ /  / __ \/ __   / __/ __ \/ ___/
 / / / / / / / / /_/ / /  / / / / /_/  / /_/ /_/ / /
/_/_/ / /_/ /_/ .___/_/  /_/ /_/\__,_/\__/\____/_/
             /_/
              V1.0
`
	fmt.Printf("%s\n%56s\n\n", banner, "IPHarvester "+version)
}
