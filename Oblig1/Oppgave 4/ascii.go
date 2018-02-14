package main


import (
	"fmt"

)

func main() {

	for i := 128; i <= 255; i++ {
		fmt.Printf("%X %c %b \n", i, i, i)
	}
}

func ExtendedASCIIText() {

}



// Utskriften vil variere i hvor mange ASCII-symboler som representeres. Hvilke tegn som representeres varierer med programmet som kjører prossessen, og programmiljøet det tilhører. 
// I normale brukerprogrammer skal det ikke være noe problem med å skrive ut heksadesimalverdiene eller binærverdiene, da disse tilhørerer

// Ved tre ulike pc'er som kjører Windows10 PC representeres alle tegn [i ASCII] fra {A0(heks), ..., FF} i (umodifisert) Command Promt. Tegnene før, {80, ..., 9F}, kan ikke representeres, så hvert tegn erstattets med  (som i CP er et inbokset spørsmålstegn). I vanilla goland representeres flere tegn - bare 80, 81, 8D, 8E, 8F, 90} eksluderes. Disse representeres også med , selv om symbolet i goland er visuelt likt et blankt space.
