package tlb

import (
	"encoding/hex"
	"github.com/xssnick/tonutils-go/tvm/cell"
	"testing"
)

func TestBlockMaster(t *testing.T) {
	boc, _ := hex.DecodeString("b5ee9c72e2020152000100002a490000002400cc00ea0180026202fe033003520361037a03940404047404c0056805a8069a06b4075c079c0806087608c30a160a3a0a5e0b0a0b2a0b4a0b6a0b880ba60bc20bde0bfa0c160c320cd80d5c0d800da00dec0e380e580e780e980eb60ed60ef60f160f360f560f76102010a8110e119011ae11cc11ea120612aa132a13761442148f14ae154415621580159e15bc15da15f81616163416521670168e16ac16ca16d816e616f417021710171e172c173a1748175617641772178017cc17da17e817f6180418121820182e187a1888189618a418b218c018ce18dc18ea18f8190619521976199a19e71a921ab21ad21b1f1b6b1b8a1ba81bf51c411c5e1c7a1cc71d131d2e1d4a1d971de31dfe1e4b1e661f0c1f591fdc2029207b20c721122132217f21cb21ea220a22572276229422e12300234d236c238c23d923f82445249124b024fd251c25c62613269a26e7274c279927e5286628b328d0291d293a298729a429f12a3d2a582afc2b492bc82c152c612cad2ccc2cda2d272d442d912dae2dfb2e182e652e822ecf2eec2f392f562fa32fc0300d302a3077309430e130fe314b316831b531d2321f323c325a3308335533a1343634443491349e34eb34f835453552356035ad35f936063653366036ad36ba370737143722376f377c37c937d6382338d838e6399a3a4e3a5c3aa93ab63ac43ad23b1f3b2c3b793b863bd33be03c2d3c3a3c873c943ce13cee3d3b3d483d953da43df13e683f1c3f693f763f843fd1401d402a40774084409240df412b41384185419241df41ec41fa42ae43624416442244284476449e44f244ff4542454c463046484656466546744684472847d0487848844890491649d64a5c4a6e4b124bd34bdc4c624c7e4d2f4dd04ddc4de84e6e4f2e4fb44fc65086510c511e51c3523152f052f7537d538e54325493041011ef55aaffffff11000100020003000401a09bc7a9870000000004010173ed450000000100ffffffff0000000000000000634e93ea00001d3677b8338000001d3677b83384955d862e00058edb0173ed410173bfbec400000003000000000000002e00050211b8e48dfb4a0eebb004000600070a8a040a13051bcbbbdeccd56f979164b7da81b8e49732e7215334e2b8ce57c41888d03190bd932f59e8bcca9b92cd1032c316407ca6099409a8aedf4146f39e95ffec016e016e000b000c14892736daee89910b52d7041a889bf97c864cfc84eeafba291a1b5b2e931cc1b5e800084a33f6fd0be55a2d75c3eae367b5ba338705f0319041b70e23a5c9b65374cf2739898e58f78b372f9292751451def1be4dc7cf494ef17470574d85c253ef77746b8ea127c00123012401250126009800001d3677a8f1440173ed443de180887d5f5a84d44bd19c87cbb664b0561d5eb81da88c5782b0e36e9a07e4d7fd7d801561f54bffc0cb5c4ec4e855deeeeb6fdf26d4c99a086ffafb93580a022581fa7454b05a2ea2ac0fd3a2a5d348d295400800080008001d43b9aca00250775d8011954fc400080201200009000a0015be000003bcb355ab466ad00015bfffffffbcbd0efda563d0245b9023afe2ffffff1100ffffffff00000000000000000173ed4400000001634e93e700001d3677a8f1440173ed4160000d000e000f0010245b9023afe2ffffff1100ffffffff00000000000000000173ed4500000001634e93ea00001d3677b833840173ed416000110012001300142848010124871f46ee0eb1ae00a27d5c29f6cdbcc378c1f4f1380805ff2297c9ed9fcf2200013213a09776db739953220712f110cafb8c5d8dc0fab70e9391a4894fc7cc8706b210f3282d0d6691f0235fdd6b31911b4bd49619c99ab3dbcb80305cea71d75d2eb0016d00128207e9d152c168ba8ab00018009122330000000000000000ffffffffffffffff81fa7454b05a2ea2a8280091001634558d88cb7e0929c9a44ffc1cf3c5a230c0db9b7ba3f7e489ee00f8289579a5c967d13c5fdbab5e261b0b55885aa0a63abbc4487521903d912f3e35c296ad0eb23d001b0010cc26aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaac23305e8350c57b37e003f008e004000410111000000000000000050001532139abefb5447ac011734d52324dd76a05aff0775b45c928601222624dbc2141bf2fa4c3e38ec56db5fc900c7c628d963f6fd5f016fdbf141f3f5ae91c3314e7b7f016d00128207e9d152e9a4694ab00072009122330000000000000000ffffffffffffffff81fa7454ba691a52a828009100162455cc26aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaac23305e83a13cd8b7e0128008e00170041006bb0400000000000000000b9f6a280000e9b3bd478a1ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc028480101de5adf45c03a745fc9d3a418d9e2ba096db9c1aaa77bc0898b6d9ebb611d2b40000532bfe179ee8495cd144a2f6d30950db2d48088fac5f9a778519dc0114219bf54cead2e7212dcdf398af721e9e425782d575b6c4026b7ec52ea1d3fe9dbb2bbe2aecf001a0010000100f59f3900058edb600003a6cef51e28880000e99c6da994200b9dfdf25ea78c41c94e4584ff2623b917b789f5d0e6a859861410542cb94e76fdc6efde77dc19b6aa8de0d4ad05651468c38cb9f64e1f8cf1351de6d5a7d98d56d6ded0be00bb00bc23130103f4e8a960b45d4558001900740091231301022a87a0b197c88778001a001b0091331367865aed64db08164a3138d816a8cacbb8b8fa46bc1fd1023ef5bd4e8fb6299ddc6201a0bfe76f22f36ca8ad0c54aa645c5739752b0a53bac3e57e9dda18febd0027000f01015ec2a32762fd21d800270028009122130100cbc4fd8a34cb65a8001c007822130100596b57d9c1932d880079001d221301003f0bad3989c46848001e007c221100e0b187aea6583a68007d001f221100e0a7528c0ef9512800200080220f00c141a6498c4d0800810021220f00c02225548664a800220084220f00c0221de12e910800850023220f4030085e7768002a00870024220f00c02170f2272c280025008a219dbceaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa818042d76318e19f365e6f7e0780eb2bc9f0c382940ef38b61bb6257ad617eb36fe8c9f56964f2800003a6cef51e28700262277cff55555555555555555555555555555555555555555555555555555555555555554085ac1288e0000000000000074d9dea3c5118042d76318e195d0008c008d2313010068a1a14c598fee380029002a009122130100f62101db096d33a80092002b28480101357b3e386bb95837e17d8fc7dd37f292efdc3301f6e109d8b33eb8a02afbb95d002228480101df3229c929cdae91378fd16242bda5a97246c54da4e9700d7427829825ee7b5d001922130100dd08f96dc461bcc8002c009522130100a6df074312541a08002d002e22130100902873121142a0c80098002f221100f6b6943101117948003900ae2213010090175c7b3161aee8009a0030221301008fb45fdc97d252080031009d221301008f689e5fb80803e80032009f221301008f677b3e09283ac800a00033221301008f677a70024cc70800a20034221301008f6779cbfeb8f188003500a521a1bcd99999999999999999999999999999999999999999999999999999999999982011ecef393af6d61933fe8eada66c79771a5de359b379b70bfe4414022dee90877585c9818f39dda000003a6cef51e2850036227bcff33333333333333333333333333333333333333333333333333333333333333334081ac1664bc000000000000074d9dea3c50e011ecef393af6d6196d000a700372355ec039e4242ff8cc69bf4260c44ddc7b820f838fa85ad1828d2b83ace409d6c02a3b89a505ac592d94a7c4d00a900aa00382179a0634dfa13634f7a130000800006226ee3dc107c1c7d42d68c14695c1d67204eb60151dc4d282d62c96ca53e26c0100ee542c8b882e30ec339334e5ca000ac221100f6a2a63eab3d4e48003a00b0221100ea5905b0b329bd88003b00b2221100ea58fcd0996ec14800b3003c220f00c035987df0cb08003d00b6219bbd62f8f7bea30f8ab5e9f16c3fb8642b118f56ed1bdc49600dbe5220c8b1af9e040c474f803d1a0544cba813425adf3253dd3727a789c9e418b5e788a4cab5df805caee92b00000e9b3bd478a1c0003e236fcff34517c7bdf5187c55af4f8b61fdc321588c7ab768dee24b006df29106458d7cf21881f48000000000000074d9dea3c5110311d3e017f000b800b900ba28480101db29f7a5808e1a673feb2258d777f3005642991b8025b1f92bcd7c498a8bd8ed000222bf000100f59f3900058edb600003a6cef335e0880000e99c6da994200b9dfdf25ea78c41c94e4584ff2623b917b789f5d0e6a859861410542cb94e76fdc6efde77dc19b6aa8de0d4ad05651468c38cb9f64e1f8cf1351de6d5a7d98d56d6ded0be0042004328480101b20e36a3b36a4cdee601106c642e90718b0a58daf200753dbb3189f956b494b600012213c3c000074d9de66bc12000bd004432014645ed4db913f636932b4bebb489b51112e1a415136d57d6c6735ffc4bd556606e33f960111aa97c043f6040c47d1298da52d6a46a0378304ba87f61f3dee9db0010000c20005100522211480000e9b3bccd782400bf00452211200003a6cef335e09000c100462211200003a6cef335e09000c3004722116200003a6cef335e0900c500482211200003a6cef335e09000c700492211200003a6cef335e09000c9004a2211200003a6cef335e09000cb004b2211000003a6cef335e09000cd004c2211400000e9b3bccd782400cf004d2211000003a6cef335e09000d1004e2211400000e9b3bccd782400d3004f2211400000e9b3bccd782400d500502211d000003a6cef335e0900d900da220120005300f822012000dd006722012000540055220120005600fc220120010f005f220120005700fe2201200058010022012000590102220120005a0104220120005b0106220120005c0108220120005d010a220120005e010c28480101f25a1e1d7f11115186543ff6eb95e3d9b98f71d2c959af6b0dad6b63cd1e6d6900012201200060011222012001130061220120011500622201200063011822012001190064220120011b0065220120011d006628480101a96f5d75bc79b8d1640e680704965baaa2245e4b3a5ad8e980ef5a3568409418000222012000df006822012000e10069220120006a00e422012000e5006b22012000e7006c22012000e9006d22012000eb006e220120006f00ee220120007000f0220120007100f2284801016f2780ba9d3cdce8eee34a23d893d90800da0ac1be8a973c513909136b7f636b000223130103f4e8a974d234a558007300740091231301022a87a0c5b59fe77800750076009128480101827773c365eccfd6cb46a3f783a09f1aba77ce1a0a4d62048569e9b845e955f8016b3313e844a4da57b7cf2b0b52b004cc5886c966478e37012ee2d9b473bf083e1e3072beee68024dd8cc7dffdfde9e114e4818dc10824b446d661453963fa117c1fcbf0027000f01015ec2a33b80d481d8008f0090009122130100cbc4fd8a34cb65a80077007822130100596b57d9c1932d880079007a28480101cf20bddca78403c3e40e2ab1b3eeb526c425b5efb531e6c5bd3d52019c25125c002628480101ff7081e66c7f0d6e868021316b0189b9e67b61284c176cd74f78fa6baa18a025001a221301003f0bad3989c46848007b007c221100e0b187aea6583a68007d007e284801011d818de56750d053b2a227f0da85ba37fb1869d0c774629d04e2668e1204c0a3001b28480101174c3878604468b08b7b76f5349c41201ae28b81ca1d94794ab11a692e4668250018221100e0a7528c0ef95128007f0080220f00c141a6498c4d080081008228480101e00721ae4b2be2ed708fa68e6b7bec2759a107504812069b3b8aa9acea346c66001528480101747c06e45f53ca1dc7d23f39db07d12f2e67fa595a36fa41d26683ab42d1a3040014220f00c02225548664a800830084220f00c0221de12e91080085008628480101eb38ef90c590bedb3ce31140d2d4176d43db6b7aab35df685afc4ccf2a383209000b2848010179a2e20b8a926ab2fe83108ff00f2fbced9958047008e5cb5fdf8c798aab63850010220f4030085e7768002a0087008828480101a248b81f22333cc28f6b6744e4298aefcd9b6f2dc5d7c99e1da1b28c37f3aa0c0007220f00c02170f2272c280089008a219dbceaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa818042d76318e1805a8c67376726b2dcf563d47e9c2ed6fa8fd993942535d6c8ed758902593e99400003a6cef706707008b284801010143b3d2dd671b2559543155e003f847022e510b3a57afabbca05d4069c327ef000d2277cff55555555555555555555555555555555555555555555555555555555555555554085ac1288e0000000000000074d9dee0ce118042d76318e195d0008c008d2848010164a43970f2007a1da6d6fc81773cc095d1cc270e81359e471f3b03469abeb7b5000c214900000027cbb9d1062954439a83a91f27835fb9d2e3e798910356650c3c493c946234646840008e28480101374e198a900e08edc634a5f2ad73e388b0a3019d24269fae8046024e437476b1001028480101903aa268fecbed38822a8972ba42eadb53c0972f11b8486f1534210245db4898002322130100f62101ef274493a80092009328480101a5a7d24057d8643b2527709d986cda3846adcb3eddc32d28ec21f69e17dbaaef0001284801012bd772e408a34578028922281a3e5b5384970a6a6dd741b1cfa3b80a3e5ec57d002322130100dd08f981e2391cc80094009522130100a6df0757302b7a080096009728480101b90ed7fc04a4971294b12a078ec8189e8fdba184de6e23043922a774ae403ee2002422130100902873262f1a00c800980099221100f6b694310111794800ad00ae28480101eda54e0b0237690499c3e159ab800469fdbcb3c162d42181c2c298acd4e98f3100152213010090175c8f4f390ee8009a009b28480101cc6ead611f9fa7c0598d8f88d658fe0b91f5f9c9635c872154234c16c722970c0014221301008fb45ff0b5a9b208009c009d221301008f689e73d5df63e8009e009f28480101c7c146bea2ced23475861d11146c0560a46c3d243563fda0e32bf8c34229d2670013221301008f677b5226ff9ac800a000a128480101ef1aa8b2068cf6a8eadef8197235a5d5976865a32a3ad1fe80db069ddb8cc2fe001128480101c2ef35325f62d0b4cc17d1f5d083894100c3c478504d70b6eb8d3cf26e604ff40011221301008f677a842024270800a200a3284801018a51fe69422dbf7e028fb1dcac5a62064eefeb4c080793e78a24ef22334b307c0010221301008f6779e01c90518800a400a521a1bcd99999999999999999999999999999999999999999999999999999999999982011ecef3bbeb1c21823653419f6ebabf10f7978dae7e33a12d3bbe9815917a62eb4fe902f2a96e88600003a6cef70670500a62848010150725eee52e86432f846698a08ac153a67bc9ad9c160130af907c3bef05f29480007227bcff33333333333333333333333333333333333333333333333333333333333333334081ac1664bc000000000000074d9dee0ce0e011ecef3bbeb1c2196d000a700a8284801016217f872c99fafcb870f2c11a362f59339be95095f70d00b9cff2f6dcd69d3dd000e2355ec039e4242ff8cc69bf4260c44ddc7b820f838fa85ad1828d2b83ace409d6c02a3b89a505ac592d94a7c4d00a900aa00ab28480101ff06225996392d9e78d92fef981828f3459892841111b2d352901236d506cb65000b28480101336df3bd068890e3f26c1a8f5e77c4bf7cc3c81fc88006ab614b6db43647262600072179a0634dfa13634f7a130000800006226ee3dc107c1c7d42d68c14695c1d67204eb60151dc4d282d62c96ca53e26c0100ee542c8b882e30ec389aaabdca000ac28480101b8ad45439ed0f9f1ffb12362a0c0a6f522734feed11dda077d5f6067f1305170000b221100f6a2a63eab3d4e4800af00b028480101e2a96bbff9be849635722263833d77a90f0a832b410f8b73bca56041fd7e21970016221100ea5905b0b329bd8800b100b22848010130dd0d5ef5796dc4c101fbf5b4b083599e509d0f738b07a8dbfad6b5ae53aecb0012221100ea58fcd0996ec14800b300b42848010130219e3c8c788af6da8a296da6f3e9925c909eed9821a0ae1911c38f56f7b37e000b28480101e2bc337ece7f3af5171f3265f44c612fc2fcba87f4b4563dc7fdc3285dd6a44d0008220f00c035987df0cb0800b500b6219bbd62f8f7bea30f8ab5e9f16c3fb8642b118f56ed1bdc49600dbe5220c8b1af9e040c474f804c5ac6247ef1e5d11d080c3d8b21135b54598a72e11fbc6ebe1fa0c4b2a7df0a00000e9b3bdc19c1c000b72848010118dd0a8040c21a2cfb6c0acf4ad636dc67ef3ab0a3e102f1b43ad500c55728d00007236fcff34517c7bdf5187c55af4f8b61fdc321588c7ab768dee24b006df29106458d7cf21881f48000000000000074d9dee0ce110311d3e017f000b800b900ba284801017269fb9feb45d719ebdbc3b0816b987bab06f43378dc84dc84d55727905482140002004811fd096c000000000000000000000000000000000000000000000000000000000000000028480101986c49971b96062e1fba4410e27249c8d73b0a9380f7ffd44640167e68b215e800032213c3c000074d9dea3c512000bd00be22012000db00dc28480101258d602eaa21d621634dcf86692aeae308ff3cf888f3edafc6a5b21848d732f900182211480000e9b3bd478a2400bf00c0284801014b01ebcf5425735461aa8b83bae89e70fa21e95d2ee85e57b05dad26c1d6d53000162211200003a6cef51e289000c100c22848010165b0a85a0fdea0c76a2a98445623ea62427099a6318624794dea416f1bdc6f5c00152211200003a6cef51e289000c300c428480101b5b64686c719580155341cb7347af0405dec7158c283ad30833b07325bdc48a5001422116200003a6cef51e28900c500c628480101fde4f74a9866e3de066d6d27e3b1fe107053ecce8b54d8b05ebf4a3b0789c26b00112211200003a6cef51e289000c700c828480101f7a4391731a8136b142d214311bd2f8c162938f27185d22de576a045a13b1e1600102211200003a6cef51e289000c900ca2848010187c846be2bc06a266ae017ae9a13c66cf156125edd95b8bd4f6cfe3c903e3b35000f2211200003a6cef51e289000cb00cc2848010122da148fcc6a6a317ae3c41ee888034019cbfa89e57f306b85601dd2045d6daa000e2211000003a6cef51e289000cd00ce2848010191c44865f6767ab41750fbf5117df2d8be3110925c7993aa2e03780673c31f32000d2211400000e9b3bd478a2400cf00d0284801016d16afa0d70d41df6abe49636527c0b566bd3b722b731eba03433d7efbcb3908000b2211000003a6cef51e289000d100d228480101d744ca7d3ce6fe4538b3fa6a138971ca129c227d8a6736a9cd1d33c2f1fd06cc000a2211400000e9b3bd478a2400d300d42848010175d211346d824c33aff56800c12e0b320854590aadfd85e3f909502cdb6ec3c100082211400000e9b3bd478a2400d500d628480101322f03bbddf42b900d602199315f5d4befa1a9282a2a6c845f3db6ccd2b6bfc000062211cc00003a6cef51e28900d700d82211000003a6cef335e09000d900da00a9d0000074d9dea3c51000003a6cef51e28802e7da887bc30110fabeb509a897a3390f976cc960ac3abd703b5118af0561c6dd340fc9affafb002ac3ea97ff8196b89d89d0abbdddd6dfbe4da9933410dff5f726b01528480101523e62a3a95932c2a65f2314a8a818f82f48644967cc31dcfda9954109d8b55100012848010177c2748c31a7f78c56862aa9d06df60981de7aaa59e67d4d0360a2903384fe1500013201032a5ac73da06a6b989d158bec539003d36dc087d663eda6337be5667c284f16310ee22bacedde5f1c215edbbbdf7a1c20c98ec248b7893266ebfceeb41817bd000f000c2000f700f822012000dd00de284801019deed5e9cd5995ad6c97a06276c939029a1d05a6de03b6c724a4b5567e9adb7a000e22012000df00e0284801016bc4ad2e5c909f6f452be243edc65694f7e6db5f2fc615f69756954a60a563a2000c22012000e100e2284801016925c827cdb72656785a860c0ed1b94c1ff9f0614b9e2ed1b0aa1ee8fbb395aa000c22012000e300e422012000e500e62848010197d9c97586b5cf9a93f5077cf1e13c91f7a4d5b240601e4d08030ab62cd17707000b28480101f613c63e75ce90bdb3aadf01297ba9a958588392473ea542ef8654f281d2854f000922012000e700e828480101d83f99b6b2deca33e45337ea0fa4788a5590c2a9f88654c24c1e4b5282ec7787000822012000e900ea28480101497deb7f82cc061521c9f6bf58ddd3043ecb1dbaea13352ecb73bb53236a9dd8000622012000eb00ec28480101e86bec3c2e5a0c5b9bad30e9b0efd5c74409fece4efd571f8fe02eccbbd0af1a000422012000ed00ee22012000ef00f02848010178a2f12e152f91343bff8aeda8ca7bab1039578fb6b03832c150f22786d0500c000422012000f100f228480101b26a0cc496805853f303d8a00ae9fc7f7b20dc7cab6d1d1c21f5b86469874a84000202012000f300f428480101a31f27b17ffa79bcaf0e47f55dffa054f825e019e447026255e7e1a8d7488701000200b1bcd91dbefdb40075ad92878e330bb79115bcfc28f3c5b9833df391ce8138514a3199fcd1800000000000006500000002b0d782ed8000003f2a3414d8b199f72d0000000000000040800000038a0ed0708000002d14497b814002014800f500f600afbc6827bcf8957c10b8a5694ecd7f0dd41e6a2cd906c77e5340983b618fb6fb0800000000000000000000000000000000000000000000000000000000c695ef7e000000000000007200000006efab3f64000000450742496300afbc66e5f2524ea28a3bde37c9b8f9e929de2e8e0ae9b0b84a7deeee8d71d424e8c69d27d400000000000000e2000000093197d6c2000000a8632502f4c69d1ba4000000000000008a000000116e3e81da0000006af0f2148122012000f900fa28480101912d60694234d59e4645f5d2ebd90e081979a3f6eaf4124bec3980e4547a5094000e22012000fb00fc220120010f011022012000fd00fe28480101348a81067d100edaf90feeb18db50c3c315a07c6c0944b52368c30e76a6f40df000b22012000ff0100284801018540d2166efad6f7a81289ddf3983d3ed177993dce47ccb150f2fcc287428d53000a220120010101022848010110b3b5e79df7c963efb443120853eb1bf9377e78020993bf79d5aaa9b02d1a6a000822012001030104284801016f610eec3a1e4dc9bdacbda0e586e7a8f6b4734b6599ecc0f8c5d0e9666d0ed3000822012001050106284801019100c451439a1cfdcf444d77bc78d03f19ca5e71b1f8fdae5e9e0ccf3e8214a000072201200107010828480101e0140ab9f7e276e1143af00713243e470dfc2c93c02b124622926fd33551a71b00062201200109010a284801014ef684da255649795b7830d1100f419d8f8a0eeb9ed6ed3610ba20b5d815deed0003220120010b010c2848010155cdb8f72801ef11ba562172ed2626c88208eddcf4a0c8f6d5447a785d02b790000202037820010d010e28480101b6eb72df89b91190ab85640f1ef9817bf00e49c5c11e8fd173b5b382ca4a104700010073dde8c69d27d40000000002e7da88000004baec525bbe000096f6a2fa0e38c69d27d4000000001598a6b2000004da46f3846200009e49d38f1eeb00afbbdc4b61f8041625a15bee3b094ff72034e12e69e8d71521ac6748fe6359832319afc8c8000000000000066800000033ffaa8b0000000419c4a9d68319afccd800000000000003d00000002e72705fc80000026a49c0209c284801018dfe3c99df194f8fec2b5b64b5ef08296b853794a29497c7c425ca62a44695e6000c2201200111011222012001130114284801010c275d6749b7c9102256e4abafdecda16a1697f20d26cd0e711bd9d58ef4a2f2000a28480101c169f7745c95d5f3f6b4e550c15978aaf563631f3a9e6ddaaa361ed4042e4f05000922012001150116284801017a3b4493fefcfd2275fa2f6ab01a8db5d70a443fb48dfb00e152545adfb97cbb0007220120011701182201200119011a28480101b13de2fa76c60764833d05264a2e1081609e2dfa9a04bf8d6c5e1162ac7d47cb000728480101027742b12159d2d1310044b4a94e1eea928905b045871a52b552b4b4841288400006220120011b011c28480101df4611dc79f46dc700809e0c3140796be5ca9572c3f3fab70ddbe6a5460bf4900003220120011d011e28480101d52b65c44fcc1a90bbbf8cc01e8ab9c7b6c51f95c2735d6de72a669c1135a8360003020162011f01200201200121012200b0bc885a77c249fb95f38bb13224853b7944942c4b10b84f1b99aa32892aabd4f46335f78a00000000000000bc00000003fb0146f100000074c1cd03f56335e64b0000000000000058000000040ca9eccb000000387d19d10000afbc60807ffe2b018ea1eb65ddc523f7772fc1e1f16e73ee8905b4b66c12275ca800000000000000000000000000000000000000000000000000000000c67fd5260000000000000096000000081b2d1d7000000064c8dda77100afbc799607d471065ad26b0a721946ed764b15a01cf341e11989f08863962a362800000000000000000000000000000000000000000000000000000000c69d27d4000000000000003e00000002165b2b1c0000002c0b7cc7ad0103802001270001021101918f8df47d89a592d9a8e2220276e210d49d789c174ab2b303917d71c6655837000782012f0317cca5687735940043b9aca00401280129012a0247a00f076afb8843d0d2618df1779691876f9ffd59f9b30f47df60f49496744dec67200610012e013b0103d040012b003fb000000000400000000000000021dcd650010ee6b280087735940043b9aca004010150012d01db500e3a26680b9f6a280000e9b3bd478a000000e9b3bd478a0d73fd8f873243316a5b55a0395be4d1c584d71a7c52116b6379a3e93f9649360375de02eb9865b5786e167c4411e5c415cf6064be75fe04b0f81f0c4b84c7260880002c7d7c00000000000000000b9f6a0b1a749f2a012c001343b9aca0021dcd650020020161012e013b0106460600013f020340400130013102037604013201330297bf955555555555555555555555555555555555555555555555555555555555555502aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaad00000074d9dee0ce0c1014c014e0397beb33333333333333333333333333333333333333333333333333333333333333029999999999999999999999999999999999999999999999999999999999999999cf8000074d9dee0ce00400134013501360397be8517c7bdf5187c55af4f8b61fdc321588c7ab768dee24b006df29106458d7cf029a28be3defa8c3e2ad7a7c5b0fee190ac463d5bb46f71258036f9488322c6be7cf8000074d9dee0ce004001410142014301035040013701034040013b0082722f7566ede0ba3a333ac2ca4e9820a0eb28fa3c675e8c5b7378fbba7d487af6b6d8b330226ee7a4226c9a4e28167203a4dec229d3f51655422b56dd7122352fda03af7333333333333333333333333333333333333333333333333333333333333333300001d3677b8338199ff4756d3363cbb8d2ef1acd9bcdb85ff220a0116f74843bac2e4c0c79ceed000001d3677a8f142634e93ea0001408014d013801390082722f7566ede0ba3a333ac2ca4e9820a0eb28fa3c675e8c5b7378fbba7d487af6b69e73e012c2b93293818802ecda692b6a70c7bc140c3d6ba22159dd15949099300205203024013a015100a0431b9004c4b4000000000000000000960000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003af7333333333333333333333333333333333333333333333333333333333333333300001d3677b8338232a2e0d714f1820922c85912266b270b551f97fd88c9ce96b02fbe0b94fedd3300001d3677b83381634e93ea0001408013c013d013e0101a0013f0082729e73e012c2b93293818802ecda692b6a70c7bc140c3d6ba22159dd1594909930d8b330226ee7a4226c9a4e28167203a4dec229d3f51655422b56dd7122352fda020f0409283baec018110140015100ab69fe00000000000000000000000000000000000000000000000000000000000000013fccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccd283baec000000003a6cef706700c69d27d440009e42614c107ac0000000000000000064000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000103504001440103504001470082723145f857768776495406acbcd9f6451e43b82a7cf2b787bdfcd66f54e8f61eb2f5fc1aa51cd06879f30dac067b3d17d0571b7c8ef15db6c57ce3e90f63e7d80803af734517c7bdf5187c55af4f8b61fdc321588c7ab768dee24b006df29106458d7cf00001d3677b833817a340a8997502684b5be64a7ba6e4f4f1393c8316bcf1149956bbf00b95dd25600001d3677a8f143634e93ea0001408014d014501460082723145f857768776495406acbcd9f6451e43b82a7cf2b787bdfcd66f54e8f61eb20d9c166ab6df5f0d47d18e86fc45c1e5f42681c1184337189ef2e4aa3f2552c10205203034014a014b03af734517c7bdf5187c55af4f8b61fdc321588c7ab768dee24b006df29106458d7cf00001d3677b83383a4dbec8658831b756fd060883f7d013972d9838f66cebcd2e28d66f2b2d6d46900001d3677b83381634e93ea0001408014d014801490082720d9c166ab6df5f0d47d18e86fc45c1e5f42681c1184337189ef2e4aa3f2552c1f5fc1aa51cd06879f30dac067b3d17d0571b7c8ef15db6c57ce3e90f63e7d8080205303034014a014b00a042665004c4b400000000000000000030000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000069600000009600000004000600000000000519ae84f17b8f8b22026a975ff55f1ab19fde4a768744d2178dfa63bb533e107a409026bc03af7555555555555555555555555555555555555555555555555555555555555555500001d3677b83383f9b2f37bf03c07595e4f861c14a0779c5b0ddb12bd6b0bf59b7f464fab4b279400001d3677a8f143634e93ea0001408014d014e014f0001200082720ac47779e474df79ac188caf2308fa7fccf511a8be789a6502f15ca63fba64408669008ce4710e1108a5eee86c282b1d13feaf7634e0c592943ae844ddd4ca0c02053030240150015100a041297004c4b40000000000000000002e00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000005bc00000000000000000000000012d452da449e50b8cf7dd27861f146122afe1b546bb8b70fc8216f0c614139f8e04d7cef969")
	c, _ := cell.FromBOC(boc)

	var block Block
	if err := LoadFromCell(&block, c.BeginParse()); err != nil {
		t.Fatal(err)
	}

	parents, err := block.BlockInfo.GetParentBlocks()
	if err != nil {
		t.Fatal(err)
	}

	println(len(parents))
}

func TestBlockNotMaster(t *testing.T) {
	boc, _ := hex.DecodeString("b5ee9c72e1021c0100040b00001c00c400de0170020402a0033c036a037c0387039e03b6041c048204ce04ea0536055405a005ec060406200700077007bc080908100817041011ef55aaffffff110102030402a09bc7a98700000000840101c745200000000100000000000000000000000000634e94ec00001d367caaae4000001d367caaae419bbc68ac00058fb00173ed920173bfbec400000003000000000000002e05060211b8e48dfb43b9aca00407080a8a04250ec78adc9d082383679c3289edc662b628be0e34e51a8f7c412e98d24c8a5fb59960f376a6ad4dce93f406ce904add5a2aea140c99b877d02f67f1cd1e5f51021902190c0d03894a33f6fdb1c342502d7261843b4a3bfdbfb766c45705b7c4410af03c358431620ff05a79b1be0d76ede085c08726e04bad3c5779d949364eb56540f06c2c49b98d514111401a1b1b009800001d367c9b6c040173ed92b57df82537164b18661e22f620e1a7a15826a73d7402eef9433d55c030232370a7caa150ac8f2f4c74cb5c77e6671edb6f8accd65c683faf6e48a88720b2c72d009800001d367c9b6c0101c7451f78d2820caf6a5f100a444450ddab2f7754bbce7c6027dce5349269227866124a33b3efd318a7ec75c8f26844fd4dce5f581927f670a0087d7fec56658b487d720225826b977bb75290e16c135cbbddba94870b40080909000d0010ee6b2800080201200a0b0013be000003bc91627aea900013bfffffffbc8b96fc9c50235b9023afe2ffffff110000000000000000000000000001c7451f00000001634e94e900001d367c9b6c010173ed91200e0f10235b9023afe2ffffff110000000000000000000000000001c7452000000001634e94ec00001d367caaae410173ed9220141516284801017e49cb3c190a5033a93c907c6631d4459cf4bf71f57f041dd14270fb919423dc000122138209ae5deedd4a4385b011192848010125e39d851243cee82c062dd588cfa4587461b7869f68023bad26988d33bf8a24000223130104d72ef76ea521c2d81213192848010105a0d0f5cf8e9d2d98f032e935e8de2208463332de6c74af0b9d5cfc2bc2802102162848010157c418ac5021e527850e982354ed5a21fd7a0b0ac719e443fcd3c80f496dc4db003401110000000000000000501722138209ae5deedd4a4385b0181921d90000000000000000ffffffffffffffff826b977bb75290e16bb5f5e54ddd448c900001d367c9b6c040173ed92b57df82537164b18661e22f620e1a7a15826a73d7402eef9433d55c030232370a7caa150ac8f2f4c74cb5c77e6671edb6f8accd65c683faf6e48a88720b2c72d819006bb0400000000000000000b9f6c900000e9b3e4db601ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0284801012aa19c773967de4112363f58e8331a68fb2b3fcb1d55daf352b93c497a019ce4021728480101b3e9649d10ccb379368e81a3a7e8e49c8eb53f6acc69b0ba2ffa80082f70ee39000100030020000102b1e6b8f1")
	c, _ := cell.FromBOC(boc)

	var block Block
	if err := LoadFromCell(&block, c.BeginParse()); err != nil {
		t.Fatal(err)
	}

	parents, err := block.BlockInfo.GetParentBlocks()
	if err != nil {
		t.Fatal(err)
	}

	println(len(parents))
}
