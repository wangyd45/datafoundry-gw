# datafoundry-gw hawkular API使用说明
## hawkular API
### 获取CPU信息
**请求方法：** POST 在请求头中加入token

**URL：** host:port/hawkular/cpu

**参数：** ?bucketDuration=12mn&start=-6h,start可不传默认-8h

**请求体：** JSON格式 {"namespace":["service-brokers","service-brokers1",..."service-brokersN"]}，可传入多个namespace

**调用示例：** curl -i -X POST "127.0.0.1:10012/hawkular/cpu?bucketDuration=20mn&start=-1h" -d '{"namespace":["service-brokers"]}' -H "Authorization:Bearer EtfU_QKrFU0jI9YqYzVAQhYy56AmK-ov477eDO4pqL8" -H "Hawkular-Tenant: service-brokers"

**返回结果：** JSON格式 

[{"namespace":"service-brokers","info":[{"start":1531101492994,"end":1531102692994,"min":509112000,"avg":746960060000,"median":653054300000,"max":1770177600000,"sum":657324800000000,"samples":880,"empty":false},{"start":1531102692994,"end":1531103892994,"min":513039200,"avg":748482900000,"median":662315700000,"max":1771908400000,"sum":658665000000000,"samples":880,"empty":false},{"start":1531103892994,"end":1531105092994,"min":547256800,"avg":750007400000,"median":497430700000,"max":1773671000000,"sum":660006550000000,"samples":880,"empty":false}]}]

### 获取内存信息

**请求方法：** POST 在请求头中加入token

**URL：** host:port/hawkular/memory

**参数：** ?bucketDuration=12mn&start=-6h,start可不传默认-8h

**请求体：** JSON格式 {"namespace":["service-brokers","service-brokers1",..."service-brokersN"]}，可传入多个namespace

**调用示例：** curl -i -X POST "127.0.0.1:10012/hawkular/memory?bucketDuration=20mn&start=-1h" -d '{"namespace":["service-brokers"]}' -H "Authorization:Bearer EtfU_QKrFU0jI9YqYzVAQhYy56AmK-ov477eDO4pqL8" -H "Hawkular-Tenant: service-brokers"

**返回结果：** JSON格式 

[{"namespace":"service-brokers","info":[{"start":1531101825447,"end":1531103025447,"min":3956736,"avg":8575995,"median":7647123.5,"max":24588288,"sum":15093752000,"samples":1760,"empty":false},{"start":1531103025447,"end":1531104225447,"min":3956736,"avg":8577718,"median":7704274,"max":24580096,"sum":15096783000,"samples":1760,"empty":false},{"start":1531104225447,"end":1531105425447,"min":3956736,"avg":8573189,"median":7664170.5,"max":24580096,"sum":15088812000,"samples":1760,"empty":false}]}]

### 获取网络信息

**请求方法：** POST 在请求头中加入token

**URL：** host:port/hawkular/network/:sigin

sigin的值为rx或tx

**参数：** ?bucketDuration=12mn&start=-6h,start可不传默认-8h

**请求体：** JSON格式 {"namespace":["service-brokers","service-brokers1",..."service-brokersN"]}，可传入多个namespace

**调用示例：** curl -i -X POST "127.0.0.1:10012/hawkular/network/rx?bucketDuration=20mn&start=-1h" -d '{"namespace":["service-brokers"]}' -H "Authorization:Bearer EtfU_QKrFU0jI9YqYzVAQhYy56AmK-ov477eDO4pqL8" -H "Hawkular-Tenant: service-brokers"

curl -i -X POST "127.0.0.1:10012/hawkular/network/tx?bucketDuration=20mn&start=-1h" -d '{"namespace":["service-brokers"]}' -H "Authorization:Bearer EtfU_QKrFU0jI9YqYzVAQhYy56AmK-ov477eDO4pqL8" -H "Hawkular-Tenant: service-brokers"

**返回结果：** JSON格式 

[{"namespace":"service-brokers","info":[{"start":1531102014739,"end":1531103214739,"min":32.333336,"avg":4192.689,"median":5149.722,"max":11961.158,"sum":3689566.2,"samples":880,"empty":false},{"start":1531103214739,"end":1531104414739,"min":35.47059,"avg":4199.792,"median":5213.5522,"max":11820.824,"sum":3695817,"samples":880,"empty":false},{"start":1531104414739,"end":1531105614739,"min":33.448277,"avg":4200.712,"median":5200.791,"max":11102.5,"sum":3696626.5,"samples":880,"empty":false}]}]

[{"namespace":"service-brokers","info":[{"start":1531102100834,"end":1531103300834,"min":24.666668,"avg":4060.8228,"median":5170.8315,"max":11155.053,"sum":3573524,"samples":880,"empty":false},{"start":1531103300834,"end":1531104500834,"min":27.352942,"avg":4077.348,"median":5187.832,"max":11064.588,"sum":3588066.2,"samples":880,"empty":false},{"start":1531104500834,"end":1531105700834,"min":19.733332,"avg":4069.1575,"median":5203.4224,"max":10371.89,"sum":3580858.8,"samples":880,"empty":false}]}]

###错误码
Error Code  | Error Describe
------------- | -------------
200  | ok
400  | bad request
415  | server body error
417  | request body error
500  | server error
	


