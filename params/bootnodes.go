// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import "github.com/TerraVirtuaCo/vanarchain-blockchain/common"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// Ethereum Foundation Go Bootnodes
	"enode://d860a01f9722d78051619d1e2351aba3f43f943f6f00718d1b9baa4101932a1f5011f16bb2b1bb35db20d6fe28fa0bf09636d26a87d31de9ec6203eeedb1f666@18.138.108.67:30303", // bootnode-aws-ap-southeast-1-001
	"enode://22a8232c3abc76a16ae9d6c3b164f98775fe226f0917b0ca871128a74a8e9630b458460865bab457221f1d448dd9791d24c4e5d88786180ac185df813a68d4de@3.209.45.79:30303",   // bootnode-aws-us-east-1-001
	"enode://2b252ab6a1d0f971d9722cb839a42cb81db019ba44c08754628ab4a823487071b5695317c8ccd085219c3a03af063495b2f1da8d18218da2d6a82981b45e6ffc@65.108.70.101:30303", // bootnode-hetzner-hel
	"enode://4aeb4ab6c14b23e2c4cfdce879c04b0748a20d8e9b59e25ded2a08143e265c6c25936e74cbc8e641e3312ca288673d91f2f93f8e277de3cfa444ecdaaf982052@157.90.35.166:30303", // bootnode-hetzner-fsn
}

// HoleskyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Holesky test network.
var HoleskyBootnodes = []string{
	// EF DevOps
	"enode://ac906289e4b7f12df423d654c5a962b6ebe5b3a74cc9e06292a85221f9a64a6f1cfdd6b714ed6dacef51578f92b34c60ee91e9ede9c7f8fadc4d347326d95e2b@146.190.13.128:30303",
	"enode://a3435a0155a3e837c02f5e7f5662a2f1fbc25b48e4dc232016e1c51b544cb5b4510ef633ea3278c0e970fa8ad8141e2d4d0f9f95456c537ff05fdf9b31c15072@178.128.136.233:30303",
}

// SepoliaBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Sepolia test network.
var SepoliaBootnodes = []string{
	// EF DevOps
	"enode://4e5e92199ee224a01932a377160aa432f31d0b351f84ab413a8e0a42f4f36476f8fb1cbe914af0d9aef0d51665c214cf653c651c4bbd9d5550a934f241f1682b@138.197.51.181:30303", // sepolia-bootnode-1-nyc3
	"enode://143e11fb766781d22d92a2e33f8f104cddae4411a122295ed1fdb6638de96a6ce65f5b7c964ba3763bba27961738fef7d3ecc739268f3e5e771fb4c87b6234ba@146.190.1.103:30303",  // sepolia-bootnode-1-sfo3
	"enode://8b61dc2d06c3f96fddcbebb0efb29d60d3598650275dc469c22229d3e5620369b0d3dedafd929835fe7f489618f19f456fe7c0df572bf2d914a9f4e006f783a9@170.64.250.88:30303",  // sepolia-bootnode-1-syd1
	"enode://10d62eff032205fcef19497f35ca8477bea0eadfff6d769a147e895d8b2b8f8ae6341630c645c30f5df6e67547c03494ced3d9c5764e8622a26587b083b028e8@139.59.49.206:30303",  // sepolia-bootnode-1-blr1
	"enode://9e9492e2e8836114cc75f5b929784f4f46c324ad01daf87d956f98b3b6c5fcba95524d6e5cf9861dc96a2c8a171ea7105bb554a197455058de185fa870970c7c@138.68.123.152:30303", // sepolia-bootnode-1-ams3
}

var VanarBootnodes = []string{
	"enode://2ad32cead20f1835c0f21fef2f972a77ca952b60bad727ed2ccf80de7b6a12b6bdd70ad9f75feff199a323b23da56bb727b3037769cefdcc75631b44a631d868@bootnode1.vanarchain.com:30311",
	"enode://91f122a78d00d06f8020bfb9295b02466d90fb53c97bbea44908101ac66daa6a682849af036709758360e86a32134750212e45261377b8e7cd5546386cb46155@35.234.128.51:30311",
	"enode://5fad56ee9169e72b2c715045052267de385580a328dfafefde33c4641f7711e25e8e72727f57c0b68f8d29a3c378e1fd89bcd2a90cee806f327373624e84d0bd@34.105.180.132:30311",
	"enode://527d2861ea2db2787638efc5899ee1ed5537df53bcc94901a3a3cc00fa96dfb2531ac2f877352cd03b9a389ed1212591726e821f6437d8073440abad6c433a00@bootnode2.vanarchain.com:30311",
	"enode://5aed5f6172b952cdde3315dec5237ea376bd35b198a4ee4c6718a689e2c18d7a8d3a2591c4d8df9d3de0de98a21cdeef5a3be7104d1a379a593646a956406e03@bootnode3.vanarchain.com:30311",
        "enode://feb45a116199cc62e3e32ee15aedec17509a0dea03ae5cb6b190c9f6248b2f4bd4f6d365a279159a43778edb49d5ce79128b6cf02183c190ae104dbfd26195f7@bootnode4.vanarchain.com:30311",
        "enode://e75ac9225b3ee21fa15cfba371fa6f146108c4fd002138122fd26e756d8fded7338de2e007e0bb908250c7b7cf46236307086056cb3b1c0832eeda300e020505@bootnode5.vanarchain.com:30311",
        "enode://e192f318514301d9968108cd14e5ab55d97ede0f70b25ac1b83761fe9bf260b5f3542622e655733dd95d097123e5686b00bc0d6c891cf9412eda5c731e119a28@bootnode6.vanarchain.com:30311",
}

var VanguardBootnodes = []string{
	//"enode://bfe747af336911ac305ce58abfe5e3da2d3f4bc0fa46d2eccc0b51f3cc43f8cc28b27b463072db8d71c2d5c09e7ecd87a81d24299b0730d951e6c0f2fae684b3@bootnode1-vanguard.vanarchain.com:30303",
	//"enode://a0887bf2f2e2dda814c710767dae7aee30695e90a8de54f5655b9f2e3d389257959b65991a5900246a3d422acc8bc96f7ac1188e350f8569df473f68b03d2006@bootnode2-vanguard.vanarchain.com:30303",
	"enode://c5590525e6fa5688c6ac13b0a7e0cbdb42edae64063ac097ae4261864f4c59654324aa049a0568307a29b25bf4b2842421b8a157f2dbf0d70cc8bedeb184e1bd@bootnode1-vanguard.vanarchain.com:30311",
	"enode://e27b211391ad0d3c71651fa36555a13defa812e39b78cfb859b3f1e3496126328e47fcd864fa450ff030ecda53287c1e582e18b6feae2278961a17897fe97857@bootnode2-vanguard.vanarchain.com:30311",
	"enode://b1256c669a2d43c42145293625ffeff7bfe124543d0ca87307a18a0ad380727ec31c7666c320ca5bcbcfeebef17200e8409652318c22cc8123834a6d058ddd5f@bootnode3-vanguard.vanarchain.com:30311",
	"enode://c5088916737bdba6c05297e937950d76ca8d4af70635789b180703fa063cdaf771fd7c130946feb067b19013a86cb3f6079fee83d2e1bb689377e4b4d909f0c5@bootnode4-vanguard.vanarchain.com:30311",
	"enode://873880162dc1cd7319c755244bd73ab40b06c0ebac7b7bdac7550845acc74fd1205150f0e6e99cad025b95ff57158710285812ca7b502640dd7d5e1e4a736a0f@bootnode5-vanguard.vanarchain.com:30311",
	
}

var TestnetBootnodes = []string{
	// "enode://7d92f1024395e6b6edba64683b221f93bd63dc84eaddddfbe9a2a80908cc2106dd081e729e94cd7bbb371467e344c15dd82b22cb8de364a26b1befbd17c32814@bootnode1.bimtvi.com:30311",
	// "enode://70424047d5647192eec5bc1ca4a79a998ca7ecdb30d711695063e8e4be7bc38136644678781ce462ac57cada35edc31d701865e4f7a9877059264a4a195dc7c9@bootnode2.bimtvi.com:30311",

	"enode://027a78587c3a20bb8516fce7591ce14d8e17b5e72d2fd884e68ea9b30cec22ff8a342397c3e4758881854c06b2779ce74b694354f34302bd47ddf0dfcdaa9ad6@bootnode1.bimtvi.com:30311",
	"enode://fa40a8fb72fd0fe386062f93525f2d7d203b3beb8cc59ca5058c1fb18efa861bb926be0d59145ecaea42553c3dc4c961486c1349aeaa42fc7bfa9224cbfe3220@bootnode2.bimtvi.com:30311",
}

var EternalBootnodes = []string{
	"enode://2508e4734141379121aec10d7858477b2f78c30438c390d17ead4c2a0c86175e45c8b07de88404cdbd4e3408e4f7a328f71ddf8a3fb6d288b022e8043e707f52@10.10.10.21:30311",
}

// GoerliBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Görli test network.
var GoerliBootnodes = []string{
	// Upstream bootnodes
	"enode://011f758e6552d105183b1761c5e2dea0111bc20fd5f6422bc7f91e0fabbec9a6595caf6239b37feb773dddd3f87240d99d859431891e4a642cf2a0a9e6cbb98a@51.141.78.53:30303",
	"enode://176b9417f511d05b6b2cf3e34b756cf0a7096b3094572a8f6ef4cdcb9d1f9d00683bf0f83347eebdf3b81c3521c2332086d9592802230bf528eaf606a1d9677b@13.93.54.137:30303",
	"enode://46add44b9f13965f7b9875ac6b85f016f341012d84f975377573800a863526f4da19ae2c620ec73d11591fa9510e992ecc03ad0751f53cc02f7c7ed6d55c7291@94.237.54.114:30313",
	"enode://b5948a2d3e9d486c4d75bf32713221c2bd6cf86463302339299bd227dc2e276cd5a1c7ca4f43a0e9122fe9af884efed563bd2a1fd28661f3b5f5ad7bf1de5949@18.218.250.66:30303",

	// Ethereum Foundation bootnode
	"enode://a61215641fb8714a373c80edbfa0ea8878243193f57c96eeb44d0bc019ef295abd4e044fd619bfc4c59731a73fb79afe84e9ab6da0c743ceb479cbb6d263fa91@3.11.147.67:30303",

	// Goerli Initiative bootnodes
	"enode://d4f764a48ec2a8ecf883735776fdefe0a3949eb0ca476bd7bc8d0954a9defe8fea15ae5da7d40b5d2d59ce9524a99daedadf6da6283fca492cc80b53689fb3b3@46.4.99.122:32109",
	"enode://d2b720352e8216c9efc470091aa91ddafc53e222b32780f505c817ceef69e01d5b0b0797b69db254c586f493872352f5a022b4d8479a00fc92ec55f9ad46a27e@88.99.70.182:30303",
}

var V5Bootnodes = []string{
	// Teku team's bootnode
	"enr:-KG4QOtcP9X1FbIMOe17QNMKqDxCpm14jcX5tiOE4_TyMrFqbmhPZHK_ZPG2Gxb1GE2xdtodOfx9-cgvNtxnRyHEmC0ghGV0aDKQ9aX9QgAAAAD__________4JpZIJ2NIJpcIQDE8KdiXNlY3AyNTZrMaEDhpehBDbZjM_L9ek699Y7vhUJ-eAdMyQW_Fil522Y0fODdGNwgiMog3VkcIIjKA",
	"enr:-KG4QDyytgmE4f7AnvW-ZaUOIi9i79qX4JwjRAiXBZCU65wOfBu-3Nb5I7b_Rmg3KCOcZM_C3y5pg7EBU5XGrcLTduQEhGV0aDKQ9aX9QgAAAAD__________4JpZIJ2NIJpcIQ2_DUbiXNlY3AyNTZrMaEDKnz_-ps3UUOfHWVYaskI5kWYO_vtYMGYCQRAR3gHDouDdGNwgiMog3VkcIIjKA",
	// Prylab team's bootnodes
	"enr:-Ku4QImhMc1z8yCiNJ1TyUxdcfNucje3BGwEHzodEZUan8PherEo4sF7pPHPSIB1NNuSg5fZy7qFsjmUKs2ea1Whi0EBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpD1pf1CAAAAAP__________gmlkgnY0gmlwhBLf22SJc2VjcDI1NmsxoQOVphkDqal4QzPMksc5wnpuC3gvSC8AfbFOnZY_On34wIN1ZHCCIyg",
	"enr:-Ku4QP2xDnEtUXIjzJ_DhlCRN9SN99RYQPJL92TMlSv7U5C1YnYLjwOQHgZIUXw6c-BvRg2Yc2QsZxxoS_pPRVe0yK8Bh2F0dG5ldHOIAAAAAAAAAACEZXRoMpD1pf1CAAAAAP__________gmlkgnY0gmlwhBLf22SJc2VjcDI1NmsxoQMeFF5GrS7UZpAH2Ly84aLK-TyvH-dRo0JM1i8yygH50YN1ZHCCJxA",
	"enr:-Ku4QPp9z1W4tAO8Ber_NQierYaOStqhDqQdOPY3bB3jDgkjcbk6YrEnVYIiCBbTxuar3CzS528d2iE7TdJsrL-dEKoBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpD1pf1CAAAAAP__________gmlkgnY0gmlwhBLf22SJc2VjcDI1NmsxoQMw5fqqkw2hHC4F5HZZDPsNmPdB1Gi8JPQK7pRc9XHh-oN1ZHCCKvg",
	// Lighthouse team's bootnodes
	"enr:-IS4QLkKqDMy_ExrpOEWa59NiClemOnor-krjp4qoeZwIw2QduPC-q7Kz4u1IOWf3DDbdxqQIgC4fejavBOuUPy-HE4BgmlkgnY0gmlwhCLzAHqJc2VjcDI1NmsxoQLQSJfEAHZApkm5edTCZ_4qps_1k_ub2CxHFxi-gr2JMIN1ZHCCIyg",
	"enr:-IS4QDAyibHCzYZmIYZCjXwU9BqpotWmv2BsFlIq1V31BwDDMJPFEbox1ijT5c2Ou3kvieOKejxuaCqIcjxBjJ_3j_cBgmlkgnY0gmlwhAMaHiCJc2VjcDI1NmsxoQJIdpj_foZ02MXz4It8xKD7yUHTBx7lVFn3oeRP21KRV4N1ZHCCIyg",
	// EF bootnodes
	"enr:-Ku4QHqVeJ8PPICcWk1vSn_XcSkjOkNiTg6Fmii5j6vUQgvzMc9L1goFnLKgXqBJspJjIsB91LTOleFmyWWrFVATGngBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpC1MD8qAAAAAP__________gmlkgnY0gmlwhAMRHkWJc2VjcDI1NmsxoQKLVXFOhp2uX6jeT0DvvDpPcU8FWMjQdR4wMuORMhpX24N1ZHCCIyg",
	"enr:-Ku4QG-2_Md3sZIAUebGYT6g0SMskIml77l6yR-M_JXc-UdNHCmHQeOiMLbylPejyJsdAPsTHJyjJB2sYGDLe0dn8uYBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpC1MD8qAAAAAP__________gmlkgnY0gmlwhBLY-NyJc2VjcDI1NmsxoQORcM6e19T1T9gi7jxEZjk_sjVLGFscUNqAY9obgZaxbIN1ZHCCIyg",
	"enr:-Ku4QPn5eVhcoF1opaFEvg1b6JNFD2rqVkHQ8HApOKK61OIcIXD127bKWgAtbwI7pnxx6cDyk_nI88TrZKQaGMZj0q0Bh2F0dG5ldHOIAAAAAAAAAACEZXRoMpC1MD8qAAAAAP__________gmlkgnY0gmlwhDayLMaJc2VjcDI1NmsxoQK2sBOLGcUb4AwuYzFuAVCaNHA-dy24UuEKkeFNgCVCsIN1ZHCCIyg",
	"enr:-Ku4QEWzdnVtXc2Q0ZVigfCGggOVB2Vc1ZCPEc6j21NIFLODSJbvNaef1g4PxhPwl_3kax86YPheFUSLXPRs98vvYsoBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpC1MD8qAAAAAP__________gmlkgnY0gmlwhDZBrP2Jc2VjcDI1NmsxoQM6jr8Rb1ktLEsVcKAPa08wCsKUmvoQ8khiOl_SLozf9IN1ZHCCIyg",
}

const dnsPrefix = "enrtree://AKA3AM6LPBYEUDMVNU3BSVQJ5AD45Y7YPOHJLEF6W26QOE4VTUDPE@"

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	var net string
	switch genesis {
	case MainnetGenesisHash:
		net = "mainnet"
	case GoerliGenesisHash:
		net = "goerli"
	case SepoliaGenesisHash:
		net = "sepolia"
	case HoleskyGenesisHash:
		net = "holesky"
	case VanarGenesisHash:
		net = "vanar"
	case VanguardGenesisHash:
		net = "vanguard"
	case TestnetGenesisHash:
		net = "testnet"
	case EternalGenesisHash:
		net = "eternal"
	default:
		return ""
	}
	return dnsPrefix + protocol + "." + net + ".ethdisco.net"
}
