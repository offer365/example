./mongo
# 建库
# 键集合
非auth 模式下创建用户
use admin
db.createUser({user:"admin",pwd:"666666",roles:["root"],db:"admin"})

配置文件添加 auth=true 重启mongodb
use admin
db.auth("admin","666666")
use example
db.createUser({user:"example",pwd:"example",roles:[{role:"dbOwner",db:"example"}]})
use example
db.auth("example","example") # 返回1


# 删除某个集合
db.principals.drop()
# 创建唯一索引
db.principals.ensureIndex({"id":1},{"unique":true})
# 数据聚合
db.products.aggregate([{$match:{"_id":ObjectId("5d5d0a3a306ba203ca7447a1")}},{$lookup:{from:"projects",localField:"projects",foreignField:"_id",as:"projects"}},{$lookup:{from:"users",localField:"authors",foreignField:"_id",as:"authors"}}]).pretty()