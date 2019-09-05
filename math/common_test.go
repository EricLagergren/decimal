package math

import "github.com/ericlagergren/decimal/v4"

var gB *decimal.Big

var benchPrecs = [...]int{
	5, 16, 25, 32, 41, 50, 75, 82, 97,
	100, 137, 250, 333, 500, 646, 750, 943,
	1500, 5000, 7500, 15000,
}

// assorted values to 1000 precision decimals
// these values came from wolframalpha.com
const (
	_pi_6 = "-.523598775598298873077107230546583814032861566562517636829157432051302734381" +
		"034833104672470890352844663691347752213717774515640768258430371954226568021" +
		"413519575047350450323086850926607437158159155063660738135162610988907688079" +
		"274705631130527545200318190941427820576724768409054441368898934543374856878" +
		"954097834434385931362480253486827138209015285894061315431726668555088424803" +
		"411086897357825323585268238842878394293265325515364352896988768632184186424" +
		"679077063327124927891980958787481871323030501991521638945560406776107181003" +
		"565824399204122865117029976823950461756536196048862792064112469744611567522" +
		"000094687857543926804630961890459631601528939529786911406816870825571691091" +
		"597561841798713281543153923669992685354836993477339069693302271629128551660" +
		"086451201891666666395496341658509955288802682719765837409909224484717107087" +
		"051375557447505876988531361835000522973125481097922201396903436196278191217" +
		"266375581738125911455193260477313725631265625326296976296755361871134435500" +
		"3213127685326515360700332"
	_pi_5 = "-.628318530717958647692528676655900576839433879875021164194988918461563281257" +
		"241799725606965068423413596429617302656461329418768921910116446345071881625" +
		"696223490056820540387704221111928924589790986076392885762195133186689225695" +
		"129646757356633054240381829129713384692069722090865329642678721452049828254" +
		"744917401321263117634976304184192565850818343072873578518072002266106109764" +
		"093304276829390388302321886611454073151918390618437223476386522358621023709" +
		"614892475992549913470377150544978245587636602389825966734672488131328617204" +
		"278989279044947438140435972188740554107843435258635350476934963693533881026" +
		"400113625429052712165557154268551557921834727435744293688180244990686029309" +
		"917074210158455937851784708403991222425804392172806883631962725954954261992" +
		"103741442269999999674595609990211946346563219263719004891891069381660528504" +
		"461650668937007052386237634202000627567750577317506641676284123435533829460" +
		"719650698085751093746231912572776470757518750391556371556106434245361322600" +
		"3855753222391818432840398"
	_pi_4 = "-.785398163397448309615660845819875721049292349843776455243736148076954101571" +
		"552249657008706335529266995537021628320576661773461152387645557931339852032" +
		"120279362571025675484630276389911155737238732595491107202743916483361532118" +
		"912058446695791317800477286412141730865087152613581662053348401815062285318" +
		"431146751651578897043720380230240707313522928841091973147590002832632637205" +
		"116630346036737985377902358264317591439897988273046529345483152948276279637" +
		"018615594990687391837971438181222806984545752987282458418340610164160771505" +
		"348736598806184297675544965235925692634804294073294188096168704616917351283" +
		"000142031786315890206946442835689447402293409294680367110225306238357536637" +
		"396342762698069922314730885504989028032255490216008604539953407443692827490" +
		"129676802837499999593244512487764932933204024079648756114863836727075660630" +
		"577063336171258815482797042752500784459688221646883302095355154294417286825" +
		"899563372607188867182789890715970588446898437989445464445133042806701653250" +
		"4819691527989773041050497"
	_pi_3 = "-1.047197551196597746154214461093167628065723133125035273658314864102605468762" +
		"069666209344941780705689327382695504427435549031281536516860743908453136042" +
		"827039150094700900646173701853214874316318310127321476270325221977815376158" +
		"549411262261055090400636381882855641153449536818108882737797869086749713757" +
		"908195668868771862724960506973654276418030571788122630863453337110176849606" +
		"822173794715650647170536477685756788586530651030728705793977537264368372849" +
		"358154126654249855783961917574963742646061003983043277891120813552214362007" +
		"131648798408245730234059953647900923513072392097725584128224939489223135044" +
		"000189375715087853609261923780919263203057879059573822813633741651143382183" +
		"195123683597426563086307847339985370709673986954678139386604543258257103320" +
		"172902403783333332790992683317019910577605365439531674819818448969434214174" +
		"102751114895011753977062723670001045946250962195844402793806872392556382434" +
		"532751163476251822910386520954627451262531250652593952593510723742268871000" +
		"642625537065303072140066"
	_pi_2 = "-1.570796326794896619231321691639751442098584699687552910487472296153908203143" +
		"104499314017412671058533991074043256641153323546922304775291115862679704064" +
		"240558725142051350969260552779822311474477465190982214405487832966723064237" +
		"824116893391582635600954572824283461730174305227163324106696803630124570636" +
		"862293503303157794087440760460481414627045857682183946295180005665265274410" +
		"233260692073475970755804716528635182879795976546093058690966305896552559274" +
		"037231189981374783675942876362445613969091505974564916836681220328321543010" +
		"697473197612368595351089930471851385269608588146588376192337409233834702566" +
		"000284063572631780413892885671378894804586818589360734220450612476715073274" +
		"792685525396139844629461771009978056064510980432017209079906814887385654980" +
		"259353605674999999186489024975529865866408048159297512229727673454151321261" +
		"154126672342517630965594085505001568919376443293766604190710308588834573651" +
		"799126745214377734365579781431941176893796875978890928890266085613403306500" +
		"963938305597954608210099"
	_2pi_3 = "-2.094395102393195492308428922186335256131446266250070547316629728205210937524" +
		"139332418689883561411378654765391008854871098062563073033721487816906272085" +
		"654078300189401801292347403706429748632636620254642952540650443955630752317" +
		"098822524522110180801272763765711282306899073636217765475595738173499427515" +
		"816391337737543725449921013947308552836061143576245261726906674220353699213" +
		"644347589431301294341072955371513577173061302061457411587955074528736745698" +
		"716308253308499711567923835149927485292122007966086555782241627104428724014" +
		"263297596816491460468119907295801847026144784195451168256449878978446270088" +
		"000378751430175707218523847561838526406115758119147645627267483302286764366" +
		"390247367194853126172615694679970741419347973909356278773209086516514206640" +
		"345804807566666665581985366634039821155210730879063349639636897938868428348" +
		"205502229790023507954125447340002091892501924391688805587613744785112764869" +
		"065502326952503645820773041909254902525062501305187905187021447484537742001" +
		"285251074130606144280133"
	_3pi_4 = "-2.356194490192344928846982537459627163147877049531329365731208444230862304714" +
		"656748971026119006587800986611064884961729985320383457162936673794019556096" +
		"360838087713077026453890829169733467211716197786473321608231749450084596356" +
		"736175340087373953401431859236425192595261457840744986160045205445186855955" +
		"293440254954736691131161140690722121940568786523275919442770008497897911615" +
		"349891038110213956133707074792952774319693964819139588036449458844828838911" +
		"055846784972062175513914314543668420953637258961847375255021830492482314516" +
		"046209796418552893026634895707777077904412882219882564288506113850752053849" +
		"000426095358947670620839328507068342206880227884041101330675918715072609912" +
		"189028288094209766944192656514967084096766470648025813619860222331078482470" +
		"389030408512499998779733537463294798799612072238946268344591510181226981891" +
		"731190008513776446448391128257502353379064664940649906286065462883251860477" +
		"698690117821566601548369672147911765340695313968336393335399128420104959751" +
		"445907458396931912315149"
	_pi = "-3.141592653589793238462643383279502884197169399375105820974944592307816406286" +
		"208998628034825342117067982148086513282306647093844609550582231725359408128" +
		"481117450284102701938521105559644622948954930381964428810975665933446128475" +
		"648233786783165271201909145648566923460348610454326648213393607260249141273" +
		"724587006606315588174881520920962829254091715364367892590360011330530548820" +
		"466521384146951941511609433057270365759591953092186117381932611793105118548" +
		"074462379962749567351885752724891227938183011949129833673362440656643086021" +
		"394946395224737190702179860943702770539217176293176752384674818467669405132" +
		"000568127145263560827785771342757789609173637178721468440901224953430146549" +
		"585371050792279689258923542019956112129021960864034418159813629774771309960" +
		"518707211349999998372978049951059731732816096318595024459455346908302642522" +
		"308253344685035261931188171010003137838752886587533208381420617177669147303" +
		"598253490428755468731159562863882353787593751957781857780532171226806613001" +
		"927876611195909216420199"
	_5pi_4 = "-3.926990816987241548078304229099378605246461749218882276218680740384770507857" +
		"761248285043531677646334977685108141602883308867305761938227789656699260160" +
		"601396812855128377423151381949555778686193662977455536013719582416807660594" +
		"560292233478956589002386432060708654325435763067908310266742009075311426592" +
		"155733758257894485218601901151203536567614644205459865737950014163163186025" +
		"583151730183689926889511791321587957199489941365232646727415764741381398185" +
		"093077974953436959189857190906114034922728764936412292091703050820803857526" +
		"743682994030921488377724826179628463174021470366470940480843523084586756415" +
		"000710158931579451034732214178447237011467046473401835551126531191787683186" +
		"981713813490349611573654427524945140161277451080043022699767037218464137450" +
		"648384014187499997966222562438824664666020120398243780574319183635378303152" +
		"885316680856294077413985213762503922298441108234416510476775771472086434129" +
		"497816863035944335913949453579852942234492189947227322225665214033508266252" +
		"409845763994886520525249"
	_4pi_3 = "-4.188790204786390984616857844372670512262892532500141094633259456410421875048" +
		"278664837379767122822757309530782017709742196125126146067442975633812544171" +
		"308156600378803602584694807412859497265273240509285905081300887911261504634" +
		"197645049044220361602545527531422564613798147272435530951191476346998855031" +
		"632782675475087450899842027894617105672122287152490523453813348440707398427" +
		"288695178862602588682145910743027154346122604122914823175910149057473491397" +
		"432616506616999423135847670299854970584244015932173111564483254208857448028" +
		"526595193632982920936239814591603694052289568390902336512899757956892540176" +
		"000757502860351414437047695123677052812231516238295291254534966604573528732" +
		"780494734389706252345231389359941482838695947818712557546418173033028413280" +
		"691609615133333331163970733268079642310421461758126699279273795877736856696" +
		"411004459580047015908250894680004183785003848783377611175227489570225529738" +
		"131004653905007291641546083818509805050125002610375810374042894969075484002" +
		"570502148261212288560265"
	_3pi_2 = "-4.712388980384689857693965074919254326295754099062658731462416888461724609429" +
		"313497942052238013175601973222129769923459970640766914325873347588039112192" +
		"721676175426154052907781658339466934423432395572946643216463498900169192713" +
		"472350680174747906802863718472850385190522915681489972320090410890373711910" +
		"586880509909473382262322281381444243881137573046551838885540016995795823230" +
		"699782076220427912267414149585905548639387929638279176072898917689657677822" +
		"111693569944124351027828629087336841907274517923694750510043660984964629032" +
		"092419592837105786053269791415554155808825764439765128577012227701504107698" +
		"000852190717895341241678657014136684413760455768082202661351837430145219824" +
		"378056576188419533888385313029934168193532941296051627239720444662156964940" +
		"778060817024999997559467074926589597599224144477892536689183020362453963783" +
		"462380017027552892896782256515004706758129329881299812572130925766503720955" +
		"397380235643133203096739344295823530681390627936672786670798256840209919502" +
		"891814916793863824630298"
	_5pi_3 = "-5.235987755982988730771072305465838140328615665625176368291574320513027343810" +
		"348331046724708903528446636913477522137177745156407682584303719542265680214" +
		"135195750473504503230868509266074371581591550636607381351626109889076880792" +
		"747056311305275452003181909414278205767247684090544413688989345433748568789" +
		"540978344343859313624802534868271382090152858940613154317266685550884248034" +
		"110868973578253235852682388428783942932653255153643528969887686321841864246" +
		"790770633271249278919809587874818713230305019915216389455604067761071810035" +
		"658243992041228651170299768239504617565361960488627920641124697446115675220" +
		"000946878575439268046309618904596316015289395297869114068168708255716910915" +
		"975618417987132815431539236699926853548369934773390696933022716291285516600" +
		"864512018916666663954963416585099552888026827197658374099092244847171070870" +
		"513755574475058769885313618350005229731254810979222013969034361962781912172" +
		"663755817381259114551932604773137256312656253262969762967553618711344355003" +
		"213127685326515360700332"
	_7pi_4 = "-5.497787143782138167309625920739130047345046448906435186706153036538678711000" +
		"865747599060944348704868968759151398244036632414228066713518905519378964224" +
		"841955537997179728392411934729378090160671128168437750419207415383530724832" +
		"384409126870539224603341004884992116055610068295071634373438812705435997229" +
		"018027261561052279306042661611684951194660501887643812033130019828428460435" +
		"816412422257165897645316507850223140079285917911325705418382070637933957459" +
		"130309164934811742865800067268559648891820270910977208928384271149125400537" +
		"441156191643290083728814756651479848443630058513059316673180932318421458981" +
		"000994222504211231448625099849826131816053865062762569771577143668502756461" +
		"774399338886489456203116198534923196225788431512060231779673852105849792430" +
		"907737619862499997152711587414354530532428168557541292804046857089529624414" +
		"039443353198811708379579299267505491217817551528183114667486080060921007781" +
		"296943608250322070279529235011794119128289065926118251115931299646911572753" +
		"373784069592841128735348"
	_2pi = "-6.283185307179586476925286766559005768394338798750211641949889184615632812572" +
		"417997256069650684234135964296173026564613294187689219101164463450718816256" +
		"962234900568205403877042211119289245897909860763928857621951331866892256951" +
		"296467573566330542403818291297133846920697220908653296426787214520498282547" +
		"449174013212631176349763041841925658508183430728735785180720022661061097640" +
		"933042768293903883023218866114540731519183906184372234763865223586210237096" +
		"148924759925499134703771505449782455876366023898259667346724881313286172042" +
		"789892790449474381404359721887405541078434352586353504769349636935338810264" +
		"001136254290527121655571542685515579218347274357442936881802449906860293099" +
		"170742101584559378517847084039912224258043921728068836319627259549542619921" +
		"037414422699999996745956099902119463465632192637190048918910693816605285044" +
		"616506689370070523862376342020006275677505773175066416762841234355338294607" +
		"196506980857510937462319125727764707575187503915563715561064342453613226003" +
		"855753222391818432840398"
)

func neg(str string, n uint) string { return str[:n+2] }
func pos(str string, n uint) string { return str[1 : n+2] }
