# git-code-statistics
Golang语言版，git代码量统计

项目clone到GoPath/src下面
```
git clone https://github.com/cheneylew/git-code-statistics.git
```

代码编译：
```
cd gitStat
go build main.go
./main
```

用法，main.go中更改要统计的项目，这些项目都要拉到本地，切换到要统计的主分支：
```
func calcRateOfContribution2()  {
	startDate := "2021-04-01"
	endDate := "2021-07-01"
	var users []User
	//git.Shell("/Users/apple/Desktop/ehsy/opc", "git checkout master")
	users = append(users, start2("/Users/apple/Desktop/ehsy/opc", startDate, endDate)...)
	users = append(users, start2("/Users/apple/Desktop/ehsy/eis", startDate, endDate)...)
	users = append(users, start2("/Users/apple/Desktop/ehsy/uniapp", startDate, endDate)...)
	users = append(users, start2("/Users/apple/Desktop/ehsy/uni-spc", startDate, endDate)...)
	users = append(users, start2("/Users/apple/Desktop/ehsy/crm", startDate, endDate)...)
	users = append(users, start2("/Users/apple/Desktop/ehsy/raxwell-front-html", startDate, endDate)...)
	users = append(users, start2("/Users/apple/Desktop/ehsy/raxwell-eis", startDate, endDate)...)
	users = append(users, start2("/Users/apple/Desktop/ehsy/raxwell-opc", startDate, endDate)...)
	users = append(users, start2("/Users/apple/go/src/gitlab.ehsy.com/ehsy/frontend/go-service", startDate, endDate)...)
	users = append(users, start2("/Users/apple/Desktop/ehsy/flutter-warehouse", startDate, endDate)...)
	users = append(users, start2("/Users/apple/Desktop/ehsy/sso-api", startDate, endDate)...)
	users = append(users, start2("/Users/apple/Desktop/ehsy/questionare", startDate, endDate)...)
	merge(users)
}
```

输出结果为：
```
用户名：Alan_liu        添加行数:839766                 删除行数:314191                 有效行数:525575         alan_liu@ehsy.com
用户名：lily_lv         添加行数:116880                 删除行数:3389           有效行数:113491         lily_lv@ehsy.com
用户名：Claris Lu       添加行数:31772          删除行数:7104           有效行数:24668          claris_lu@ehsy.com
用户名：SH-EHSY\Millie_Zhang    添加行数:15974          删除行数:7868           有效行数:8106   millie_zhang@ehsy.com
用户名：abel_ding       添加行数:13516          删除行数:1632           有效行数:11884          abel_ding@ehsy.com
用户名：tomas           添加行数:12186          删除行数:5745           有效行数:6441   tomas@ehsy.com
用户名：eric_he         添加行数:11729          删除行数:2423           有效行数:9306   eric_he@ehsy.com
用户名：sam_han         添加行数:8946           删除行数:3143           有效行数:5803   sam_han@ehsy.com
用户名：rory_ran@ehsy.com       添加行数:8050           删除行数:939            有效行数:7111   ehsy@rory.com
用户名：tolstoy_gong    添加行数:5267           删除行数:1297           有效行数:3970   tolstoy_gong@ehsy.com
用户名：dan_han         添加行数:3296           删除行数:895            有效行数:2401   dan_han@ehsy.com
用户名：hdy             添加行数:2102           删除行数:735            有效行数:1367   dingyi_huan@ehsy.com
用户名：kinsen_zhu      添加行数:1532           删除行数:493            有效行数:1039   kinsen_zhu@ehsy.com
用户名：SH-EHSY\Jackie_Xu       添加行数:1306           删除行数:302            有效行数:1004   xucong754802892@gmail.com
用户名：color           添加行数:763            删除行数:475            有效行数:288    color_zheng
用户名：jetty           添加行数:84             删除行数:103            有效行数:-19    gitlab@ehsy.com
用户名：Mengjuan        添加行数:72             删除行数:9213           有效行数:-9141          lily_lv@ehsy.com
用户名：CheneyLew       添加行数:6              删除行数:0              有效行数:6      cheneylew@163.com
用户名：陆茹茹  添加行数:6              删除行数:52             有效行数:-46    claris_lu@ehsy.com
提交或删除超过500行的异常节点:
a638d2ff38fe23c8169637f130bc732aa6b46619         2345    0       alan_liu@ehsy.com       Alan_liu        /Users/apple/Desktop/ehsy/sso-api
5894c94941317779874759641e44d7b856570f45         2602    0       alan_liu@ehsy.com       Alan_liu        /Users/apple/Desktop/ehsy/sso-api
12f035a43a3a6f5e962648404389c38c409ccb90         210199          0       alan_liu@ehsy.com       Alan_liu        /Users/apple/Desktop/ehsy/sso-api
9bd6ef170b56879ba8fd5f6b162323de47a2c0f9         0       303615          alan_liu@ehsy.com       Alan_liu        /Users/apple/Desktop/ehsy/sso-api
e56f9bbea5d759ee66193fc62e28df0d3ea22e7d         303612          0       alan_liu@ehsy.com       Alan_liu        /Users/apple/Desktop/ehsy/sso-api
72fc34345cb161e0c999b0aee510abcca4e87e65         2134    0       alan_liu@ehsy.com       Alan_liu        /Users/apple/go/src/gitlab.ehsy.com/ehsy/frontend/go-service
e7692811d902e76e4f2f4d4654110627826ac8a9         578     5       alan_liu@ehsy.com       Alan_liu        /Users/apple/go/src/gitlab.ehsy.com/ehsy/frontend/go-service
68dd399c7cec0bf6c31785e11e149073b2f33c29         1547    2       alan_liu@ehsy.com       Alan_liu        /Users/apple/go/src/gitlab.ehsy.com/ehsy/frontend/go-service
816d97c626505287dce27526a67bc69383b8a5a4         10974   1       alan_liu@ehsy.com       Alan_liu        /Users/apple/go/src/gitlab.ehsy.com/ehsy/frontend/go-service
02693fb04f2ad52d8606f0a6cfbdf2348b1a4c5b         298097          7759    alan_liu@ehsy.com       Alan_liu        /Users/apple/go/src/gitlab.ehsy.com/ehsy/frontend/go-service
ad175e0bb82677b7086bb859d518908cf617bb6e         702     1       alan_liu@ehsy.com       Alan_liu        /Users/apple/go/src/gitlab.ehsy.com/ehsy/frontend/go-service
01391d22a5af88df81a3334be0f4aece2111b93e         515     320     alan_liu@ehsy.com       Alan_liu        /Users/apple/Desktop/ehsy/opc
f921ad98845c7e9dd93d22e2e01979b0f7392a5d         515     320     alan_liu@ehsy.com       Alan_liu        /Users/apple/Desktop/ehsy/raxwell-eis
2a73fbeb966fd790c015df25968c1aedc8c84ea1         515     320     alan_liu@ehsy.com       Alan_liu        /Users/apple/Desktop/ehsy/eis
ec8d1eb72f075bbf736377768e2f2bc98616ab65         515     320     alan_liu@ehsy.com       Alan_liu        /Users/apple/Desktop/ehsy/raxwell-opc
83986cc98ee21d3e2ec700315aed8507941c3c9d         107902          0       lily_lv@ehsy.com        lily_lv         /Users/apple/Desktop/ehsy/questionare
c5f62b6cf79598d80646ded60a95df8eb7325862         535     356     lily_lv@ehsy.com        lily_lv         /Users/apple/Desktop/ehsy/opc
55f862a11e2cf5b72608e6e63b54a859c17d73f7         1950    2       claris_lu@ehsy.com      Claris Lu       /Users/apple/Desktop/ehsy/opc
629fcdf0594faa9c13f6df4cc5fba7b2262ea4f4         1490    0       claris_lu@ehsy.com      Claris Lu       /Users/apple/Desktop/ehsy/opc
56db673d6db199798a8d513e36aa23ba401c2b47         1388    4       claris_lu@ehsy.com      Claris Lu       /Users/apple/Desktop/ehsy/opc
3df747446d1daddfb6d8135e57faa7cdb3848af7         801     274     claris_lu@ehsy.com      Claris Lu       /Users/apple/Desktop/ehsy/opc
75ac2c767c93d42d8deec1d0f9a8fc9db2ad9a87         717     0       claris_lu@ehsy.com      Claris Lu       /Users/apple/Desktop/ehsy/opc
690cf5c2c39d5714a438f2cefb45f759b32e0559         829     42      claris_lu@ehsy.com      Claris Lu       /Users/apple/Desktop/ehsy/opc
ed02fa1be8a1bcd51503164f3f1311da6a7db1c8         716     33      claris_lu@ehsy.com      Claris Lu       /Users/apple/Desktop/ehsy/opc
29114cc9dea7c1553c16820566e168b250896192         1372    0       claris_lu@ehsy.com      Claris Lu       /Users/apple/Desktop/ehsy/raxwell-opc
75ac2c767c93d42d8deec1d0f9a8fc9db2ad9a87         717     0       claris_lu@ehsy.com      Claris Lu       /Users/apple/Desktop/ehsy/raxwell-opc
690cf5c2c39d5714a438f2cefb45f759b32e0559         829     42      claris_lu@ehsy.com      Claris Lu       /Users/apple/Desktop/ehsy/raxwell-opc
ed02fa1be8a1bcd51503164f3f1311da6a7db1c8         716     33      claris_lu@ehsy.com      Claris Lu       /Users/apple/Desktop/ehsy/raxwell-opc
1ca12688328db014633d8feca71a3bf3bd81f201         640     171     claris_lu@ehsy.com      Claris Lu       /Users/apple/Desktop/ehsy/raxwell-eis
14edb31e0b79624299c8228b198977ae2fb2e3fb         1096    404     claris_lu@ehsy.com      Claris Lu       /Users/apple/Desktop/ehsy/raxwell-eis
3ab63e2b8488e76a5a13663a6fbacdfda984583d         2381    0       claris_lu@ehsy.com      Claris Lu       /Users/apple/Desktop/ehsy/raxwell-eis
72b16f9c25893089fcea807c5b6f8cf764e3c636         1822    885     claris_lu@ehsy.com      Claris Lu       /Users/apple/Desktop/ehsy/uniapp
0bac99f69bfcfc434bd11c222fc32be26a9e51ea         1023    1       claris_lu@ehsy.com      Claris Lu       /Users/apple/Desktop/ehsy/uniapp
ae87ff1b15f0beff5cbc8efe3e7db585be0ab553         1077    4       claris_lu@ehsy.com      Claris Lu       /Users/apple/Desktop/ehsy/sso-api
d433187cec3fad9a650e26fd04fc22c02310516f         6366    3560    millie_zhang@ehsy.com   SH-EHSY\Millie_Zhang    /Users/apple/Desktop/ehsy/raxwell-front-html
b744c4472306120c569183a60e75f08f5768b419         4718    3236    millie_zhang@ehsy.com   SH-EHSY\Millie_Zhang    /Users/apple/Desktop/ehsy/raxwell-front-html
ce5d2df92dc0a9154ec5ef894a30ca517aca14d0         1456    543     millie_zhang@ehsy.com   SH-EHSY\Millie_Zhang    /Users/apple/Desktop/ehsy/raxwell-front-html
b18035af91c5aea4d81171cd7e4e29d20ab09bdf         706     479     millie_zhang@ehsy.com   SH-EHSY\Millie_Zhang    /Users/apple/Desktop/ehsy/raxwell-front-html
a50946841db6adaff65907722a39699217bd46ea         2696    16      millie_zhang@ehsy.com   SH-EHSY\Millie_Zhang    /Users/apple/Desktop/ehsy/raxwell-front-html
48dfe4f0da1406e1e9b5b979b0d8b1e494eb38a2         781     4       abel_ding@ehsy.com      abel_ding       /Users/apple/Desktop/ehsy/opc
4eb0298e6f8dfd80d72a56482b1081142040f61b         1466    68      abel_ding@ehsy.com      abel_ding       /Users/apple/Desktop/ehsy/opc
ab4f80533bc5dd7a0947922d42d0e7f5a86ad4c4         2173    62      abel_ding@ehsy.com      abel_ding       /Users/apple/Desktop/ehsy/opc
ab4f80533bc5dd7a0947922d42d0e7f5a86ad4c4         2173    62      abel_ding@ehsy.com      abel_ding       /Users/apple/Desktop/ehsy/raxwell-opc
e1c2317d3b2fc7f38f5a2750908cabb072f5040d         682     650     tomas@ehsy.com          tomas   /Users/apple/Desktop/ehsy/opc
e51bbe0c418d86450c3fe2f0eaf81cb8863e4f23         812     3       tomas@ehsy.com          tomas   /Users/apple/Desktop/ehsy/opc
6f2e5079bbd67da80a899108650169caeca92ef1         425     2563    tomas@ehsy.com          tomas   /Users/apple/Desktop/ehsy/opc
6665a1d94806d41d67ca44c1d5182b005adeec04         777     575     tomas@ehsy.com          tomas   /Users/apple/Desktop/ehsy/opc
98b626df10fda0a01efc2ac248fd3160ff0be50d         3257    0       tomas@ehsy.com          tomas   /Users/apple/Desktop/ehsy/opc
f731a5ce2b5943f80c17039b2360aee85e405cec         595     144     tomas@ehsy.com          tomas   /Users/apple/Desktop/ehsy/opc
c724cd046a4bced677b4915586e69758c05757d5         617     11      eric_he@ehsy.com        eric_he         /Users/apple/Desktop/ehsy/opc
f2031788346285a7bbe54963fdab9bbdbc3d4edf         778     5       eric_he@ehsy.com        eric_he         /Users/apple/Desktop/ehsy/opc
35a03aab9227fe1cb22917dc52d4201c34199acd         518     1       eric_he@ehsy.com        eric_he         /Users/apple/Desktop/ehsy/opc
c37e51cad913aa1053fc17113b96a7c3a64bb08f         1631    309     eric_he@ehsy.com        贺键    /Users/apple/Desktop/ehsy/opc
317cc40b7d37242a04f2b867edf4c96300c569a4         1316    0       sam_han@ehsy.com        sam_han         /Users/apple/Desktop/ehsy/opc
8daa656bd1f9d407cc62e8f3677fd15acbb0d83b         11      598     sam_han@ehsy.com        韩庆满          /Users/apple/Desktop/ehsy/opc
6999af352837c17ed636089a56f3f5b436f49112         4929    164     ehsy@rory.com   rory_ran@ehsy.com       /Users/apple/Desktop/ehsy/raxwell-eis
ed52dd9104d5d06bdc0bd2e0021f2e7c2e2fae1d         594     10      ehsy@rory.com   rory_ran@ehsy.com       /Users/apple/Desktop/ehsy/raxwell-opc
d253ed7e3e670e37d3722f062edde3340f56389e         562     0       ehsy@rory.com   rory_ran@ehsy.com       /Users/apple/Desktop/ehsy/raxwell-opc
ed52dd9104d5d06bdc0bd2e0021f2e7c2e2fae1d         594     10      ehsy@rory.com   rory_ran@ehsy.com       /Users/apple/Desktop/ehsy/opc
d253ed7e3e670e37d3722f062edde3340f56389e         562     0       ehsy@rory.com   rory_ran@ehsy.com       /Users/apple/Desktop/ehsy/opc
4c452e5b56f4d47fce6ac6fde0d9cef1bdb6a7a7         816     0       tolstoy_gong@ehsy.com   tolstoy_gong    /Users/apple/Desktop/ehsy/opc
eb5b34f51325e333986b370d0132a1651e7ee495         572     0       tolstoy_gong@ehsy.com   tolstoy_gong    /Users/apple/Desktop/ehsy/opc
4aa426922d4a9c8dc46e1680c5ac9bc41fbdd0c3         1375    54      tolstoy_gong@ehsy.com   tolstoy_gong    /Users/apple/Desktop/ehsy/opc
f9eb571c027b6815aecb3610661e0a995d304afa         586     28      kinsen_zhu@ehsy.com     kinsen_zhu      /Users/apple/Desktop/ehsy/flutter-warehouse
c0093c3cdacd0602abd45fbddd9e681792125c3a         831     110     xucong754802892@gmail.com       SH-EHSY\Jackie_Xu       /Users/apple/Desktop/ehsy/flutter-warehouse
b6d7c922eda81ef3df2aaabc464cc806550f46eb         0       9169    lily_lv@ehsy.com        Mengjuan        /Users/apple/Desktop/ehsy/questionare
```