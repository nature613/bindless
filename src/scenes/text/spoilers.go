package text

import "github.com/tinne26/bindless/src/lang"
import "github.com/tinne26/bindless/src/ui"

type pageKey int
const (
	Preamble      pageKey = 0
	TutorialEnd   pageKey = 1
	ToBeContinued pageKey = 2
	Afterword     pageKey = 3
)

var gameTexts = []*lang.Text {
	lang.NewText(
		"\x08Preamble\x07\n" +
		"\x0BIn 3589, life was good for the average citizen: housing and food were guaranteed by the state, work was voluntary, streets were clean and most freedoms were respected.\n" +
		"\x0B\x09MSP Units\x07, implanted in the spine of every citizen, allowed \x09Headsteer\x07 to maintain the social stability.\n" +
		"\x0BDeveloped back during the Primal Wars, \x09MSP Units\x07 were once the key to end the violent outbursts that plagued the country. Nowadays, these magnetic devices had evolved to allow even greater control over individuals, with certain monitoring functions mandated to remain always active.\n" +
		"\x0BMost people had little reason to complain, but for Mirko... things changed when his brother abandoned the love of his life and his family to go work for Marunka Machart, a member of \x09Headsteer\x07.\x03\n" +
		"\x0BIt didn't make any sense. What if the rumors were true? He... had no choice.",
		"\x08Preámbulo\x07\n" +
		"\x0BLa vida en 3589 no trataba nada mal al ciudadano promedio: el habitaje y la comida estaban garantizados por el estado, el trabajo era voluntario, las calles estaban limpias y las libertades eran mayoritariamente respetadas.\n" +
		"\x0BLas \x09unidades MSP\x07, implantadas en la parte alta de la columna de cada ciudadano, permitían a \x09Cabecera\x07 mantener la estabilidad social.\n" +
		"\x0BDesarrolladas durante las Guerras Primales, en algún momento del pasado las \x09unidades MSP\x07 fueron claves para acabar con las olas de violencia que plagaban el país. Hoy en día, estos dispositivos magnéticos habían evolucionado para permitir un control sobre los individuos incluso mayor, con ciertas funciones de monitoraje permanentemente activas bajo imperativo legal.\n" +
		"\x0BLa mayoría de la población no tenía motivo de queja, pero para Mirko... las cosas cambiaron cuando su hermano repentinamente abandonó al amor de su vida y su familia para ir a trabajar para Marunka Machart, una integrante de \x09Cabecera\x07.\x03\n" +
		"\x0BNo tenía ningún sentido. Y si los rumores eran ciertos..? Ahora no le quedaba más alternativa.",
		"\x08Preàmbul\x07\n" +
		"\x0BL'any 3589, la vida no era gens dolenta pel ciutadà mitjà: l'habitatge i el menjar estaven garantitzats per l'estat, la feina era voluntària, els carrers estaven nets i les llibertats eren majoritàriament respectades.\n" +
		"\x0BLes \x09unitats MSP\x07, implantades a la part alta de la columna de cada ciutadà, permetien a \x09Capçalera\x07 mantenir l'estabilitat social.\n" +
		"\x0BDesenvolupades durant les Guerres Primals, les \x09unitats MSP\x07 van ser claus en el passat per acabar amb les onades de violència que arrasaven el país. Avui en dia, aquests dispositius magnètics havien evolucionat per permetre un control fins i tot major sobre les persones, amb certes funcions de monitoratge decretades a romandre permanentment actives.\n" +
		"\x0BLa majoria de la població no tenia motius per queixar-se, però per a en Mirko... les coses van canviar quan el seu germà sobtadament va abandonar l'amor de la seva vida i la seva família per anar a treballar per Marunka Machart, una integrant de \x09Capçalera\x07.\x03\n" +
		"\x0BNo tenia cap sentit. I si els rumors eren certs..? Ara no li quedava cap altra alternativa.",
	),
	lang.NewText(
		"You managed to clear the tutorial!\n" +
		"\x0BThere's still a lot you will have to figure out on your own, but you seem now better equipped to take on the challenge... \x0Eare you ready for it?\x0C",
		"Has conseguido superar el tutorial!\n" +
		"\x0BTodavía tienes mucho que aprender por tu cuenta, pero ahora pareces mejor preparado para enfrentarte al desafío... \x0Eestás listo para ello?\x0C",
		"Has aconseguit superar el tutorial!\n" +
		"\x0BEncara et queda bastant per aprendre pel teu compte, però ara sembles més ben preparat per encarar el repte... \x0Eestàs a punt?\x0C",
	),
	lang.NewText( // this part has to be modified anyway, so no rush translating it
		"\x0BMirko managed to release two of the slaves, but got caught while trying to disable the MSP unit of the third one.\x04\n" +
		"\x0BSeeing the deployment of public safety units around the zone and the lack of contact from Mirko, Jana started fearing the worst.\x04\n" +
		"\x0BShe was a loose end, and they would surely come for her next. But... she hadn't given up just yet.\x04\n" +
		"\x0B\x08>> to be continued",
		"\x0B(no traducido) Mirko managed to release two of the slaves, but got caught while trying to disable the MSP unit of the third one.\x04\n" +
		"\x0BSeeing the deployment of public safety units around the zone and the lack of contact from Mirko, Jana started fearing the worst.\x04\n" +
		"\x0BShe was a loose end, and they would surely come for her next. But... she hadn't given up just yet.\x04\n" +
		"\x0B\x08>> to be continued",
		"\x0B(no traduït) Mirko managed to release two of the slaves, but got caught while trying to disable the MSP unit of the third one.\x04\n" +
		"\x0BSeeing the deployment of public safety units around the zone and the lack of contact from Mirko, Jana started fearing the worst.\x04\n" +
		"\x0BShe was a loose end, and they would surely come for her next. But... she hadn't given up just yet.\x04\n" +
		"\x0B\x08>> to be continued",
	),
	lang.NewText(
		"\x08Afterword\x07\n" +
		"\x0BThanks to Hajime Hoshi, the developer of Ebitengine, for being the true magnet for the Ebitengine community.\x04\n" +
		"\x0BThanks to all the people making cool games and libraries for Ebitengine.\x04 Or just hanging around!\x04\n" +
		"\x0B...oh, and thank you for playing!\x04\n" +
		"\x0B\x08>>",
		"\x08Agradecimientos\x07\n" +
		"\x0BGracias a Hajime Hoshi, el desarrollador de Ebitengine, por ser el verdadero imán para la comunidad de Ebitengine.\x04\n" +
		"\x0BGracias a toda la gente haciendo juegos y librerías para Ebitengine. \x04O simplemente pasando el tiempo con nosotros!\x04\n" +
		"\x0B...ah, y gracias a ti por jugar!\x04\n" +
		"\x0B\x08>>",
		"\x08Agraïments\x07\n" +
		"\x0BGràcies a Hajime Hoshi, el desenvolupador d'Ebitengine, per ser el veritable imant per a la comunitat d'Ebitengine.\x04\n" +
		"\x0BGràcies a tothom que està fent jocs i llibreries per Ebitengine. \x04O simplement passant el temps amb nosaltres!\x04\n" +
		"\x0B...ah, i gràcies a tu per jugar!\x04\n" +
		"\x0B\x08>>",
	),	
}

var optTutorial = &ui.HChoice {
	Text: lang.NewText("[ Play the tutorial ]", "[ Jugar el tutorial ]", "[ Jugar el tutorial ]"),
}
var optRepeatTutorial = &ui.HChoice {
	Text: lang.NewText("[ Hmm.. can I repeat the tutorial? ]", "[ Mmm.. puedo repetir el tutorial? ]", "[ Mmm.. puc repetir el tutorial? ]"),
}
var optStartStory = &ui.HChoice {
	Text: lang.NewText("[ Skip the tutorial ]", "[ Saltarse el tutorial ]", "[ Saltar-se el tutorial ]"),
}
var optStoryNow = &ui.HChoice {
	Text: lang.NewText("[ LET'S GET TO IT! ]", "[ VAMOS ALLÁ! ]", "[ VINGA, SOM-HI! ]"),
}

var gameChoices = [][]*ui.HChoice {
	[]*ui.HChoice{ optStartStory, optTutorial },
	[]*ui.HChoice{ optRepeatTutorial, optStoryNow },
	nil,
	nil,
}
