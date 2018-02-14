# Oppgave 1
| Binære tall | Hexadesimaltall | Desimaltall |
| ------ | ------ | ------ |
| 1101 | 0xD | 13 |
| 110111101010 | 0xDEA | 3562 |
| 1010111100110100 | 0xAF34 | 44852 |
| 11111111111111111 | 0xFFFF | 65535 |
| 10001011110001010 | 0x1178A | 71562 |
------------------------------------------

Oppgave A
Beskriv kort metode for å gå fra binære tall til hexadesimale tall og motsatt. Beskriv kort metoden for å gå fra binære tall til 
desimaltall og motsatt.

Vår fremgangsmåte for å gå fra binæretall til hexadesimaletall og motsatt, var å sette opp en tallrekke med 1, 2, 4 og 8. Så gruperte vi 
binærtallet i fire, fra høyre til venstre, og stilte hver gruppe under tallrekken. Vi plusset så sammen alle tallene over 1-erene, og 
ignorerte det som stod over 0. Eksempel: 110111101010. Forste gruppe (1010) blir til A, neste (1110) blir til E, mens site (1101) blir 
til D. Altså er 110111101010 i hexadesimale tall 0xDEA. 

Metoden vi bruker for å gå fra binære tall til desimaltall er veldig lik metoden til vi brukte til å omregne til hexadesimal. Det vi 
gjorde var å forlenge tallrekken videre med å doble desimaltallene, eks. 16,32,64... Vi satte hele binæretall under desimalene og 
plusset alle enerene og ignorerte nullene. Svaret vi fikk var desimaltallet.
F.eks: 110111101010 blir til 3562.

Oppgave B
Beskriv kort metoden for å gå fra hexadesimale tall til desimaltall og motsatt.

Heksadesimaltall er et system basert på nummeret 16. Så for å gå fra heksadesimaltall til et desimaltall så må hvert heksadesimal 
bli ganget med 16^x avhenging av hvilken "plass" tallet står på. La ogg ta et eksempel. Vi gjør om 7DF. Da begynner vi med F. Siden den er
helt til høgre tar vi Fx16^0, som er 15x1, som blir 15. Videre tar vi Dx16^1, som er 13x16, som blir 208. Siste er da 7x16^2, som er 7x256, 
som blir 1792. Til slutt plusser vi alle sammen, 15 + 208 + 1792 = 2015.

Vår metode på å gå fra desimaltall til hexadesimale tall er å først dele tallet på 16. Hvis det er noen desimaltall i rest så går de bort.
Så ganget vi svaret med 16, for å få differansen av det opprinnelige tallet. Også fortsetter vi til vi ender med null. 

2015/ 16 = 125
125*16 = 2000 
2015-2000= 15
Differansen er 15(F)

125/16 = 7
7*16= 112 
125-112= 13
Differansen er 13(D)

7/16 = 0
0*16= 0 
7-0=7 
Differansen er 7

Dette gjør at svaret  blir 7DF


