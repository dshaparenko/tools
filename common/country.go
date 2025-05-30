package common

type Country struct {
	Short   string
	Aliases []string
}

var countries = make(map[string]*Country)

func CountryShort(country string) string {
	c := countries[country]
	if c != nil {
		return c.Short
	}
	return ""
}

func CountryByShort(short string) string {
	for k, v := range countries {
		if v.Short == short {
			return k
		}
	}
	return ""
}

func NewCountry(short string) *Country {
	return &Country{
		Short: short,
	}
}

func NewCountryWithAliases(short string, aliases []string) *Country {
	c := NewCountry(short)
	c.Aliases = aliases
	return c
}

func init() {

	countries["Afghanistan"] = NewCountryWithAliases("AF", []string{"AFG", "004"})
	countries["Åland Islands"] = NewCountryWithAliases("AX", []string{"ALA", "248"})
	countries["Albania"] = NewCountryWithAliases("AL", []string{"ALB", "008"})
	countries["Algeria"] = NewCountryWithAliases("DZ", []string{"DZA", "012"})
	countries["American Samoa"] = NewCountryWithAliases("AS", []string{"ASM", "016"})
	countries["Andorra"] = NewCountryWithAliases("AD", []string{"AND", "020"})
	countries["Angola"] = NewCountryWithAliases("AO", []string{"AGO", "024"})
	countries["Anguilla"] = NewCountryWithAliases("AI", []string{"AIA", "660"})
	countries["Antarctica"] = NewCountryWithAliases("AQ", []string{"ATA", "010"})
	countries["Antigua and Barbuda"] = NewCountryWithAliases("AG", []string{"ATG", "028"})
	countries["Argentina"] = NewCountryWithAliases("AR", []string{"ARG", "032"})
	countries["Armenia"] = NewCountryWithAliases("AM", []string{"ARM", "051"})
	countries["Aruba"] = NewCountryWithAliases("AW", []string{"ABW", "533"})
	countries["Australia"] = NewCountryWithAliases("AU", []string{"AUS", "036"})
	countries["Austria"] = NewCountryWithAliases("AT", []string{"AUT", "040"})
	countries["Azerbaijan"] = NewCountryWithAliases("AZ", []string{"AZE", "031"})
	countries["Bahamas"] = NewCountryWithAliases("BS", []string{"BHS", "044"})
	countries["Bahrain"] = NewCountryWithAliases("BH", []string{"BHR", "048"})
	countries["Bangladesh"] = NewCountryWithAliases("BD", []string{"BGD", "050"})
	countries["Barbados"] = NewCountryWithAliases("BB", []string{"BRB", "052"})
	countries["Belarus"] = NewCountryWithAliases("BY", []string{"BLR", "112"})
	countries["Belgium"] = NewCountryWithAliases("BE", []string{"BEL", "056"})
	countries["Belize"] = NewCountryWithAliases("BZ", []string{"BLZ", "084"})
	countries["Benin"] = NewCountryWithAliases("BJ", []string{"BEN", "204"})
	countries["Bermuda"] = NewCountryWithAliases("BM", []string{"BMU", "060"})
	countries["Bhutan"] = NewCountryWithAliases("BT", []string{"BTN", "064"})
	countries["Bolivia, Plurinational State of"] = NewCountryWithAliases("BO", []string{"BOL", "068"})
	countries["Bosnia and Herzegovina"] = NewCountryWithAliases("BA", []string{"BIH", "070"})
	countries["Botswana"] = NewCountryWithAliases("BW", []string{"BWA", "072"})
	countries["Bouvet Island"] = NewCountryWithAliases("BV", []string{"BVT", "074"})
	countries["Brazil"] = NewCountryWithAliases("BR", []string{"BRA", "076"})
	countries["British Indian Ocean Territory"] = NewCountryWithAliases("IO", []string{"IOT", "086"})
	countries["Brunei Darussalam"] = NewCountryWithAliases("BN", []string{"BRN", "096"})
	countries["Bulgaria"] = NewCountryWithAliases("BG", []string{"BGR", "100"})
	countries["Burkina Faso"] = NewCountryWithAliases("BF", []string{"BFA", "854"})
	countries["Burundi"] = NewCountryWithAliases("BI", []string{"BDI", "108"})
	countries["Cambodia"] = NewCountryWithAliases("KH", []string{"KHM", "116"})
	countries["Cameroon"] = NewCountryWithAliases("CM", []string{"CMR", "120"})
	countries["Canada"] = NewCountryWithAliases("CA", []string{"CAN", "124"})
	countries["Cape Verde"] = NewCountryWithAliases("CV", []string{"CPV", "132"})
	countries["Cayman Islands"] = NewCountryWithAliases("KY", []string{"CYM", "136"})
	countries["Central African Republic"] = NewCountryWithAliases("CF", []string{"CAF", "140"})
	countries["Chad"] = NewCountryWithAliases("TD", []string{"TCD", "148"})
	countries["Chile"] = NewCountryWithAliases("CL", []string{"CHL", "152"})
	countries["China"] = NewCountryWithAliases("CN", []string{"CHN", "156"})
	countries["Christmas Island"] = NewCountryWithAliases("CX", []string{"CXR", "162"})
	countries["Cocos (Keeling) Islands"] = NewCountryWithAliases("CC", []string{"CCK", "166"})
	countries["Colombia"] = NewCountryWithAliases("CO", []string{"COL", "170"})
	countries["Comoros"] = NewCountryWithAliases("KM", []string{"COM", "174"})
	countries["Congo"] = NewCountryWithAliases("CG", []string{"COG", "178"})
	countries["Congo, the Democratic Republic of the"] = NewCountryWithAliases("CD", []string{"COD", "180"})
	countries["Cook Islands"] = NewCountryWithAliases("CK", []string{"COK", "184"})
	countries["Costa Rica"] = NewCountryWithAliases("CR", []string{"CRI", "188"})
	countries["Côte dIvoire"] = NewCountryWithAliases("CI", []string{"CIV", "384"})
	countries["Croatia"] = NewCountryWithAliases("HR", []string{"HRV", "191"})
	countries["Cuba"] = NewCountryWithAliases("CU", []string{"CUB", "192"})
	countries["Cyprus"] = NewCountryWithAliases("CY", []string{"CYP", "196"})
	countries["Czech Republic"] = NewCountryWithAliases("CZ", []string{"CZE", "203"})
	countries["Denmark"] = NewCountryWithAliases("DK", []string{"DNK", "208"})
	countries["Djibouti"] = NewCountryWithAliases("DJ", []string{"DJI", "262"})
	countries["Dominica"] = NewCountryWithAliases("DM", []string{"DMA", "212"})
	countries["Dominican Republic"] = NewCountryWithAliases("DO", []string{"DOM", "214"})
	countries["Ecuador"] = NewCountryWithAliases("EC", []string{"ECU", "218"})
	countries["Egypt"] = NewCountryWithAliases("EG", []string{"EGY", "818"})
	countries["El Salvador"] = NewCountryWithAliases("SV", []string{"SLV", "222"})
	countries["Equatorial Guinea"] = NewCountryWithAliases("GQ", []string{"GNQ", "226"})
	countries["Eritrea"] = NewCountryWithAliases("ER", []string{"ERI", "232"})
	countries["Estonia"] = NewCountryWithAliases("EE", []string{"EST", "233"})
	countries["Ethiopia"] = NewCountryWithAliases("ET", []string{"ETH", "231"})
	countries["Falkland Islands (Malvinas)"] = NewCountryWithAliases("FK", []string{"FLK", "238"})
	countries["Faroe Islands"] = NewCountryWithAliases("FO", []string{"FRO", "234"})
	countries["Fiji"] = NewCountryWithAliases("FJ", []string{"FJI", "242"})
	countries["Finland"] = NewCountryWithAliases("FI", []string{"FIN", "246"})
	countries["France"] = NewCountryWithAliases("FR", []string{"FRA", "250"})
	countries["French Guiana"] = NewCountryWithAliases("GF", []string{"GUF", "254"})
	countries["French Polynesia"] = NewCountryWithAliases("PF", []string{"PYF", "258"})
	countries["French Southern Territories"] = NewCountryWithAliases("TF", []string{"ATF", "260"})
	countries["Gabon"] = NewCountryWithAliases("GA", []string{"GAB", "266"})
	countries["Gambia"] = NewCountryWithAliases("GM", []string{"GMB", "270"})
	countries["Georgia"] = NewCountryWithAliases("GE", []string{"GEO", "268"})
	countries["Germany"] = NewCountryWithAliases("DE", []string{"DEU", "276"})
	countries["Ghana"] = NewCountryWithAliases("GH", []string{"GHA", "288"})
	countries["Gibraltar"] = NewCountryWithAliases("GI", []string{"GIB", "292"})
	countries["Greece"] = NewCountryWithAliases("GR", []string{"GRC", "300"})
	countries["Greenland"] = NewCountryWithAliases("GL", []string{"GRL", "304"})
	countries["Grenada"] = NewCountryWithAliases("GD", []string{"GRD", "308"})
	countries["Guadeloupe"] = NewCountryWithAliases("GP", []string{"GLP", "312"})
	countries["Guam"] = NewCountryWithAliases("GU", []string{"GUM", "316"})
	countries["Guatemala"] = NewCountryWithAliases("GT", []string{"GTM", "320"})
	countries["Guernsey"] = NewCountryWithAliases("GG", []string{"GGY", "831"})
	countries["Guinea"] = NewCountryWithAliases("GN", []string{"GIN", "324"})
	countries["Guinea-Bissau"] = NewCountryWithAliases("GW", []string{"GNB", "624"})
	countries["Guyana"] = NewCountryWithAliases("GY", []string{"GUY", "328"})
	countries["Haiti"] = NewCountryWithAliases("HT", []string{"HTI", "332"})
	countries["Heard Island and McDonald Islands"] = NewCountryWithAliases("HM", []string{"HMD", "334"})
	countries["Holy See (Vatican City State)"] = NewCountryWithAliases("VA", []string{"VAT", "336"})
	countries["Honduras"] = NewCountryWithAliases("HN", []string{"HND", "340"})
	countries["Hong Kong"] = NewCountryWithAliases("HK", []string{"HKG", "344"})
	countries["Hungary"] = NewCountryWithAliases("HU", []string{"HUN", "348"})
	countries["Iceland"] = NewCountryWithAliases("IS", []string{"ISL", "352"})
	countries["India"] = NewCountryWithAliases("IN", []string{"IND", "356"})
	countries["Indonesia"] = NewCountryWithAliases("ID", []string{"IDN", "360"})
	countries["Iran, Islamic Republic of"] = NewCountryWithAliases("IR", []string{"IRN", "364"})
	countries["Iraq"] = NewCountryWithAliases("IQ", []string{"IRQ", "368"})
	countries["Ireland"] = NewCountryWithAliases("IE", []string{"IRL", "372"})
	countries["Isle of Man"] = NewCountryWithAliases("IM", []string{"IMN", "833"})
	countries["Israel"] = NewCountryWithAliases("IL", []string{"ISR", "376"})
	countries["Italy"] = NewCountryWithAliases("IT", []string{"ITA", "380"})
	countries["Jamaica"] = NewCountryWithAliases("JM", []string{"JAM", "388"})
	countries["Japan"] = NewCountryWithAliases("JP", []string{"JPN", "392"})
	countries["Jersey"] = NewCountryWithAliases("JE", []string{"JEY", "832"})
	countries["Jordan"] = NewCountryWithAliases("JO", []string{"JOR", "400"})
	countries["Kazakhstan"] = NewCountryWithAliases("KZ", []string{"KAZ", "398"})
	countries["Kenya"] = NewCountryWithAliases("KE", []string{"KEN", "404"})
	countries["Kiribati"] = NewCountryWithAliases("KI", []string{"KIR", "296"})
	countries["Korea, Democratic People's Republic of"] = NewCountryWithAliases("KP", []string{"PRK", "408"})
	countries["Korea, Republic of"] = NewCountryWithAliases("KR", []string{"KOR", "410"})
	countries["Kuwait"] = NewCountryWithAliases("KW", []string{"KWT", "414"})
	countries["Kyrgyzstan"] = NewCountryWithAliases("KG", []string{"KGZ", "417"})
	countries["Lao People's Democratic Republic"] = NewCountryWithAliases("LA", []string{"LAO", "418"})
	countries["Latvia"] = NewCountryWithAliases("LV", []string{"LVA", "428"})
	countries["Lebanon"] = NewCountryWithAliases("LB", []string{"LBN", "422"})
	countries["Lesotho"] = NewCountryWithAliases("LS", []string{"LSO", "426"})
	countries["Liberia"] = NewCountryWithAliases("LR", []string{"LBR", "430"})
	countries["Libyan Arab Jamahiriya"] = NewCountryWithAliases("LY", []string{"LBY", "434"})
	countries["Liechtenstein"] = NewCountryWithAliases("LI", []string{"LIE", "438"})
	countries["Lithuania"] = NewCountryWithAliases("LT", []string{"LTU", "440"})
	countries["Luxembourg"] = NewCountryWithAliases("LU", []string{"LUX", "442"})
	countries["Macao"] = NewCountryWithAliases("MO", []string{"MAC", "446"})
	countries["Macedonia, the former Yugoslav Republic of"] = NewCountryWithAliases("MK", []string{"MKD", "807"})
	countries["Madagascar"] = NewCountryWithAliases("MG", []string{"MDG", "450"})
	countries["Malawi"] = NewCountryWithAliases("MW", []string{"MWI", "454"})
	countries["Malaysia"] = NewCountryWithAliases("MY", []string{"MYS", "458"})
	countries["Maldives"] = NewCountryWithAliases("MV", []string{"MDV", "462"})
	countries["Mali"] = NewCountryWithAliases("ML", []string{"MLI", "466"})
	countries["Malta"] = NewCountryWithAliases("MT", []string{"MLT", "470"})
	countries["Marshall Islands"] = NewCountryWithAliases("MH", []string{"MHL", "584"})
	countries["Martinique"] = NewCountryWithAliases("MQ", []string{"MTQ", "474"})
	countries["Mauritania"] = NewCountryWithAliases("MR", []string{"MRT", "478"})
	countries["Mauritius"] = NewCountryWithAliases("MU", []string{"MUS", "480"})
	countries["Mayotte"] = NewCountryWithAliases("YT", []string{"MYT", "175"})
	countries["Mexico"] = NewCountryWithAliases("MX", []string{"MEX", "484"})
	countries["Micronesia, Federated States of"] = NewCountryWithAliases("FM", []string{"FSM", "583"})
	countries["Moldova, Republic of"] = NewCountryWithAliases("MD", []string{"MDA", "498"})
	countries["Monaco"] = NewCountryWithAliases("MC", []string{"MCO", "492"})
	countries["Mongolia"] = NewCountryWithAliases("MN", []string{"MNG", "496"})
	countries["Montenegro"] = NewCountryWithAliases("ME", []string{"MNE", "499"})
	countries["Montserrat"] = NewCountryWithAliases("MS", []string{"MSR", "500"})
	countries["Morocco"] = NewCountryWithAliases("MA", []string{"MAR", "504"})
	countries["Mozambique"] = NewCountryWithAliases("MZ", []string{"MOZ", "508"})
	countries["Myanmar"] = NewCountryWithAliases("MM", []string{"MMR", "104"})
	countries["Namibia"] = NewCountryWithAliases("NA", []string{"NAM", "516"})
	countries["Nauru"] = NewCountryWithAliases("NR", []string{"NRU", "520"})
	countries["Nepal"] = NewCountryWithAliases("NP", []string{"NPL", "524"})
	countries["Netherlands"] = NewCountryWithAliases("NL", []string{"NLD", "528"})
	countries["Netherlands Antilles"] = NewCountryWithAliases("AN", []string{"ANT", "530"})
	countries["New Caledonia"] = NewCountryWithAliases("NC", []string{"NCL", "540"})
	countries["New Zealand"] = NewCountryWithAliases("NZ", []string{"NZL", "554"})
	countries["Nicaragua"] = NewCountryWithAliases("NI", []string{"NIC", "558"})
	countries["Niger"] = NewCountryWithAliases("NE", []string{"NER", "562"})
	countries["Nigeria"] = NewCountryWithAliases("NG", []string{"NGA", "566"})
	countries["Niue"] = NewCountryWithAliases("NU", []string{"NIU", "570"})
	countries["Norfolk Island"] = NewCountryWithAliases("NF", []string{"NFK", "574"})
	countries["Northern Mariana Islands"] = NewCountryWithAliases("MP", []string{"MNP", "580"})
	countries["Norway"] = NewCountryWithAliases("NO", []string{"NOR", "578"})
	countries["Oman"] = NewCountryWithAliases("OM", []string{"OMN", "512"})
	countries["Pakistan"] = NewCountryWithAliases("PK", []string{"PAK", "586"})
	countries["Palau"] = NewCountryWithAliases("PW", []string{"PLW", "585"})
	countries["Palestinian Territory, Occupied"] = NewCountryWithAliases("PS", []string{"PSE", "275"})
	countries["Panama"] = NewCountryWithAliases("PA", []string{"PAN", "591"})
	countries["Papua New Guinea"] = NewCountryWithAliases("PG", []string{"PNG", "598"})
	countries["Paraguay"] = NewCountryWithAliases("PY", []string{"PRY", "600"})
	countries["Peru"] = NewCountryWithAliases("PE", []string{"PER", "604"})
	countries["Philippines"] = NewCountryWithAliases("PH", []string{"PHL", "608"})
	countries["Pitcairn"] = NewCountryWithAliases("PN", []string{"PCN", "612"})
	countries["Poland"] = NewCountryWithAliases("PL", []string{"POL", "616"})
	countries["Portugal"] = NewCountryWithAliases("PT", []string{"PRT", "620"})
	countries["Puerto Rico"] = NewCountryWithAliases("PR", []string{"PRI", "630"})
	countries["Qatar"] = NewCountryWithAliases("QA", []string{"QAT", "634"})
	countries["Réunion"] = NewCountryWithAliases("RE", []string{"REU", "638"})
	countries["Romania"] = NewCountryWithAliases("RO", []string{"ROU", "642"})
	countries["Russian Federation"] = NewCountryWithAliases("RU", []string{"RUS", "643"})
	countries["Rwanda"] = NewCountryWithAliases("RW", []string{"RWA", "646"})
	countries["Saint Barthélemy"] = NewCountryWithAliases("BL", []string{"BLM", "652"})
	countries["Saint Helena, Ascension and Tristan da Cunha"] = NewCountryWithAliases("SH", []string{"SHN", "654"})
	countries["Saint Kitts and Nevis"] = NewCountryWithAliases("KN", []string{"KNA", "659"})
	countries["Saint Lucia"] = NewCountryWithAliases("LC", []string{"LCA", "662"})
	countries["Saint Martin (French part)"] = NewCountryWithAliases("MF", []string{"MAF", "663"})
	countries["Saint Pierre and Miquelon"] = NewCountryWithAliases("PM", []string{"SPM", "666"})
	countries["Saint Vincent and the Grenadines"] = NewCountryWithAliases("VC", []string{"VCT", "670"})
	countries["Samoa"] = NewCountryWithAliases("WS", []string{"WSM", "882"})
	countries["San Marino"] = NewCountryWithAliases("SM", []string{"SMR", "674"})
	countries["Sao Tome and Principe"] = NewCountryWithAliases("ST", []string{"STP", "678"})
	countries["Saudi Arabia"] = NewCountryWithAliases("SA", []string{"SAU", "682"})
	countries["Senegal"] = NewCountryWithAliases("SN", []string{"SEN", "686"})
	countries["Serbia"] = NewCountryWithAliases("RS", []string{"SRB", "688"})
	countries["Seychelles"] = NewCountryWithAliases("SC", []string{"SYC", "690"})
	countries["Sierra Leone"] = NewCountryWithAliases("SL", []string{"SLE", "694"})
	countries["Singapore"] = NewCountryWithAliases("SG", []string{"SGP", "702"})
	countries["Slovakia"] = NewCountryWithAliases("SK", []string{"SVK", "703"})
	countries["Slovenia"] = NewCountryWithAliases("SI", []string{"SVN", "705"})
	countries["Solomon Islands"] = NewCountryWithAliases("SB", []string{"SLB", "090"})
	countries["Somalia"] = NewCountryWithAliases("SO", []string{"SOM", "706"})
	countries["South Africa"] = NewCountryWithAliases("ZA", []string{"ZAF", "710"})
	countries["South Georgia and the South Sandwich Islands"] = NewCountryWithAliases("GS", []string{"SGS", "239"})
	countries["Spain"] = NewCountryWithAliases("ES", []string{"ESP", "724"})
	countries["Sri Lanka"] = NewCountryWithAliases("LK", []string{"LKA", "144"})
	countries["Sudan"] = NewCountryWithAliases("SD", []string{"SDN", "736"})
	countries["Suriname"] = NewCountryWithAliases("SR", []string{"SUR", "740"})
	countries["Svalbard and Jan Mayen"] = NewCountryWithAliases("SJ", []string{"SJM", "744"})
	countries["Swaziland"] = NewCountryWithAliases("SZ", []string{"SWZ", "748"})
	countries["Sweden"] = NewCountryWithAliases("SE", []string{"SWE", "752"})
	countries["Switzerland"] = NewCountryWithAliases("CH", []string{"CHE", "756"})
	countries["Syrian Arab Republic"] = NewCountryWithAliases("SY", []string{"SYR", "760"})
	countries["Taiwan, Province of China"] = NewCountryWithAliases("TW", []string{"TWN", "158"})
	countries["Tajikistan"] = NewCountryWithAliases("TJ", []string{"TJK", "762"})
	countries["Tanzania, United Republic of"] = NewCountryWithAliases("TZ", []string{"TZA", "834"})
	countries["Thailand"] = NewCountryWithAliases("TH", []string{"THA", "764"})
	countries["Timor-Leste"] = NewCountryWithAliases("TL", []string{"TLS", "626"})
	countries["Togo"] = NewCountryWithAliases("TG", []string{"TGO", "768"})
	countries["Tokelau"] = NewCountryWithAliases("TK", []string{"TKL", "772"})
	countries["Tonga"] = NewCountryWithAliases("TO", []string{"TON", "776"})
	countries["Trinidad and Tobago"] = NewCountryWithAliases("TT", []string{"TTO", "780"})
	countries["Tunisia"] = NewCountryWithAliases("TN", []string{"TUN", "788"})
	countries["Turkey"] = NewCountryWithAliases("TR", []string{"TUR", "792"})
	countries["Turkmenistan"] = NewCountryWithAliases("TM", []string{"TKM", "795"})
	countries["Turks and Caicos Islands"] = NewCountryWithAliases("TC", []string{"TCA", "796"})
	countries["Tuvalu"] = NewCountryWithAliases("TV", []string{"TUV", "798"})
	countries["Uganda"] = NewCountryWithAliases("UG", []string{"UGA", "800"})
	countries["Ukraine"] = NewCountryWithAliases("UA", []string{"UKR", "804"})
	countries["United Arab Emirates"] = NewCountryWithAliases("AE", []string{"ARE", "784"})
	countries["United Kingdom"] = NewCountryWithAliases("GB", []string{"GBR", "826"})
	countries["United States"] = NewCountryWithAliases("US", []string{"USA", "840"})
	countries["United States Minor Outlying Islands"] = NewCountryWithAliases("UM", []string{"UMI", "581"})
	countries["Uruguay"] = NewCountryWithAliases("UY", []string{"URY", "858"})
	countries["Uzbekistan"] = NewCountryWithAliases("UZ", []string{"UZB", "860"})
	countries["Vanuatu"] = NewCountryWithAliases("VU", []string{"VUT", "548"})
	countries["Venezuela, Bolivarian Republic of"] = NewCountryWithAliases("VE", []string{"VEN", "862"})
	countries["Vietnam"] = NewCountryWithAliases("VN", []string{"VNM", "704"})
	countries["Virgin Islands, British"] = NewCountryWithAliases("VG", []string{"VGB", "092"})
	countries["Virgin Islands, U.S."] = NewCountryWithAliases("VI", []string{"VIR", "850"})
	countries["Wallis and Futuna"] = NewCountryWithAliases("WF", []string{"WLF", "876"})
	countries["Western Sahara"] = NewCountryWithAliases("EH", []string{"ESH", "732"})
	countries["Yemen"] = NewCountryWithAliases("YE", []string{"YEM", "887"})
	countries["Zambia"] = NewCountryWithAliases("ZM", []string{"ZMB", "894"})
	countries["Zimbabwe"] = NewCountryWithAliases("ZW", []string{"ZWE", "716"})
}
