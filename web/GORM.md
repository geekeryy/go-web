[中文文档](https://learnku.com/docs/gorm/2018)

### 数据库
```go
import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
  db, err := gorm.Open("mysql", "user:password@127.0.0.1:3306/dbname?charset=utf8&parseTime=True&loc=Local")
  defer db.Close()
    
  // 自动迁移
  // 自动迁移仅仅会创建表，缺少列和索引，并且不会改变现有列的类型或删除未使用的列以保护数据。
  db.AutoMigrate(&User{})
  // 创建表时添加表后缀
  db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})

}
```
### 模型
```go
type User struct {
  gorm.Model
  Name         string
  Age          sql.NullInt64
  Birthday     *time.Time
  Email        string  `gorm:"type:varchar(100);unique_index"`
  Role         string  `gorm:"size:255"` //设置字段的大小为255个字节
  MemberNumber *string `gorm:"unique;not null"` // 设置 memberNumber 字段唯一且不为空
  Num          int     `gorm:"AUTO_INCREMENT"` // 设置 Num字段自增
  Address      string  `gorm:"index:addr"` // 给Address 创建一个名字是  `addr`的索引
  IgnoreMe     int     `gorm:"-"` //忽略这个字段
}
```

### CRUD 
- 创建
```go

user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}

db.NewRecord(user) // => 返回 `true` ，因为主键为空

db.Create(&user) // 在记录插入数据库之后，gorm 将从数据库中加载这些字段的值

db.NewRecord(user) // => 在 `user` 之后创建返回 `false`

```
- 查询
```go
// 获取第一条记录，按主键排序
db.First(&user)
//// SELECT * FROM users ORDER BY id LIMIT 1;

// 获取一条记录，不指定排序
db.Take(&user)
//// SELECT * FROM users LIMIT 1;

// 获取最后一条记录，按主键排序
db.Last(&user)
//// SELECT * FROM users ORDER BY id DESC LIMIT 1;

// 获取所有的记录
db.Find(&users)
//// SELECT * FROM users;

// 通过主键进行查询 (仅适用于主键是数字类型)
db.First(&user, 10)
//// SELECT * FROM users WHERE id = 10;

// 获取所有匹配的记录
db.Where("name = ?", "jinzhu").Find(&users)
//// SELECT * FROM users WHERE name = 'jinzhu';

// <>
db.Where("name <> ?", "jinzhu").Find(&users)

// IN
db.Where("name in (?)", []string{"jinzhu", "jinzhu 2"}).Find(&users)

// LIKE
db.Where("name LIKE ?", "%jin%").Find(&users)

// AND
db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)

// Time
db.Where("updated_at > ?", lastWeek).Find(&users)

// BETWEEN
db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&users)


// Struct
// 当通过 struct 进行查询的时候，GORM 将会查询这些字段的非零值， 意味着你的字段包含 0， ''， false 或者其他 零值 , 将不会出现在查询语句中
db.Where(&User{Name: "jinzhu", Age: 0}).First(&user)
//// SELECT * FROM users WHERE name = "jinzhu"  LIMIT 1;

// Map
db.Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users)
//// SELECT * FROM users WHERE name = "jinzhu" AND age = 20;

// 多主键 slice 查询
db.Where([]int64{20, 21, 22}).Find(&users)
//// SELECT * FROM users WHERE id IN (20, 21, 22);

db.Not("name", "jinzhu").First(&user)
//// SELECT * FROM users WHERE name <> "jinzhu" LIMIT 1;

// 不包含
db.Not("name", []string{"jinzhu", "jinzhu 2"}).Find(&users)
//// SELECT * FROM users WHERE name NOT IN ("jinzhu", "jinzhu 2");

//不在主键 slice 中
db.Not([]int64{1,2,3}).First(&user)
//// SELECT * FROM users WHERE id NOT IN (1,2,3);

db.Not([]int64{}).First(&user)
//// SELECT * FROM users;

// 原生 SQL
db.Not("name = ?", "jinzhu").First(&user)
//// SELECT * FROM users WHERE NOT(name = "jinzhu");

// Struct
db.Not(User{Name: "jinzhu"}).First(&user)
//// SELECT * FROM users WHERE name <> "jinzhu";

db.Where("role = ?", "admin").Or("role = ?", "super_admin").Find(&users)
//// SELECT * FROM users WHERE role = 'admin' OR role = 'super_admin';

// Struct
db.Where("name = 'jinzhu'").Or(User{Name: "jinzhu 2"}).Find(&users)
//// SELECT * FROM users WHERE name = 'jinzhu' OR name = 'jinzhu 2';

// Map
db.Where("name = 'jinzhu'").Or(map[string]interface{}{"name": "jinzhu 2"}).Find(&users)
//// SELECT * FROM users WHERE name = 'jinzhu' OR name = 'jinzhu 2';

// 通过主键进行查询 (仅适用于主键是数字类型)
db.First(&user, 23)
//// SELECT * FROM users WHERE id = 23 LIMIT 1;
// 非数字类型的主键查询
db.First(&user, "id = ?", "string_primary_key")
//// SELECT * FROM users WHERE id = 'string_primary_key' LIMIT 1;

// 原生 SQL
db.Find(&user, "name = ?", "jinzhu")
//// SELECT * FROM users WHERE name = "jinzhu";

db.Find(&users, "name <> ? AND age > ?", "jinzhu", 20)
//// SELECT * FROM users WHERE name <> "jinzhu" AND age > 20;

// Struct
db.Find(&users, User{Age: 20})
//// SELECT * FROM users WHERE age = 20;

// Map
db.Find(&users, map[string]interface{}{"age": 20})
//// SELECT * FROM users WHERE age = 20;

// 创建或更新
db.Where(User{Name: "jinzhu"}).Assign(User{Age: 30}).FirstOrCreate(&user)

db.Where("amount > ?", DB.Table("orders").Select("AVG(amount)").Where("state = ?", "paid").QueryExpr()).Find(&orders)
// SELECT * FROM "orders"  WHERE "orders"."deleted_at" IS NULL AND (amount > (SELECT AVG(amount) FROM "orders"  WHERE (state = 'paid')));

db.Select("name, age").Find(&users)
//// SELECT name, age FROM users;

db.Select([]string{"name", "age"}).Find(&users)
//// SELECT name, age FROM users;

db.Table("users").Select("COALESCE(age,?)", 42).Rows()
//// SELECT COALESCE(age,'42') FROM users;

db.Order("age desc, name").Find(&users)
//// SELECT * FROM users ORDER BY age desc, name;

// 多个排序条件
db.Order("age desc").Order("name").Find(&users)
//// SELECT * FROM users ORDER BY age desc, name;

// 重新排序
db.Order("age desc").Find(&users1).Order("age", true).Find(&users2)
//// SELECT * FROM users ORDER BY age desc; (users1)
//// SELECT * FROM users ORDER BY age; (users2)

db.Limit(3).Find(&users)
//// SELECT * FROM users LIMIT 3;

// 用 -1 取消 LIMIT 限制条件
db.Limit(10).Find(&users1).Limit(-1).Find(&users2)
//// SELECT * FROM users LIMIT 10; (users1)
//// SELECT * FROM users; (users2)

db.Offset(3).Find(&users)
//// SELECT * FROM users OFFSET 3;

// 用 -1 取消 OFFSET 限制条件
db.Offset(10).Find(&users1).Offset(-1).Find(&users2)
//// SELECT * FROM users OFFSET 10; (users1)
//// SELECT * FROM users; (users2)

db.Where("name = ?", "jinzhu").Or("name = ?", "jinzhu 2").Find(&users).Count(&count)
//// SELECT * from USERS WHERE name = 'jinzhu' OR name = 'jinzhu 2'; (users)
//// SELECT count(*) FROM users WHERE name = 'jinzhu' OR name = 'jinzhu 2'; (count)

db.Model(&User{}).Where("name = ?", "jinzhu").Count(&count)
//// SELECT count(*) FROM users WHERE name = 'jinzhu'; (count)

db.Table("deleted_users").Count(&count)
//// SELECT count(*) FROM deleted_users;

var ages []int64
db.Find(&users).Pluck("age", &ages)

var names []string
db.Model(&User{}).Pluck("name", &names)

db.Table("deleted_users").Pluck("name", &names)

// Requesting more than one column? Do it like this:
db.Select("name, age").Find(&users)

type Result struct {
    Name string
    Age  int
}

var result Result
db.Table("users").Select("name, age").Where("name = ?", 3).Scan(&result)

// Raw SQL
db.Raw("SELECT name, age FROM users WHERE name = ?", 3).Scan(&result)


```

### 更新
```go
db.First(&user)

user.Name = "jinzhu 2"
user.Age = 100
//Save 方法在执行 SQL 更新操作时将包含所有字段，即使这些字段没有被修改。
db.Save(&user)

//// UPDATE users SET name='jinzhu 2', age=100, birthday='2016-01-01', updated_at = '2013-11-17 21:34:10' WHERE id=111;

// 如果单个属性被更改了，更新它
db.Model(&user).Update("name", "hello")
//// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111;

// 使用组合条件更新单个属性
db.Model(&user).Where("active = ?", true).Update("name", "hello")
//// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111 AND active=true;

// 使用 `map` 更新多个属性，只会更新那些被更改了的字段
db.Model(&user).Updates(map[string]interface{}{"name": "hello", "age": 18, "actived": false})
//// UPDATE users SET name='hello', age=18, actived=false, updated_at='2013-11-17 21:34:10' WHERE id=111;

// 使用 `struct` 更新多个属性，只会更新那些被修改了的和非空的字段
db.Model(&user).Updates(User{Name: "hello", Age: 18})
//// UPDATE users SET name='hello', age=18, updated_at = '2013-11-17 21:34:10' WHERE id = 111;

// 警告： 当使用结构体更新的时候, GORM 只会更新那些非空的字段
// 例如下面的更新，没有东西会被更新，因为像 "", 0, false 是这些字段类型的空值
db.Model(&user).Updates(User{Name: "", Age: 0, Actived: false})

db.Model(&user).Select("name").Updates(map[string]interface{}{"name": "hello", "age": 18, "actived": false})
//// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111;

db.Model(&user).Omit("name").Updates(map[string]interface{}{"name": "hello", "age": 18, "actived": false})
//// UPDATE users SET age=18, actived=false, updated_at='2013-11-17 21:34:10' WHERE id=111;

DB.Model(&product).Update("price", gorm.Expr("price * ? + ?", 2, 100))
//// UPDATE "products" SET "price" = price * '2' + '100', "updated_at" = '2013-11-17 21:34:10' WHERE "id" = '2';

DB.Model(&product).Updates(map[string]interface{}{"price": gorm.Expr("price * ? + ?", 2, 100)})
//// UPDATE "products" SET "price" = price * '2' + '100', "updated_at" = '2013-11-17 21:34:10' WHERE "id" = '2';

DB.Model(&product).UpdateColumn("quantity", gorm.Expr("quantity - ?", 1))
//// UPDATE "products" SET "quantity" = quantity - 1 WHERE "id" = '2';

DB.Model(&product).Where("quantity > 1").UpdateColumn("quantity", gorm.Expr("quantity - ?", 1))
//// UPDATE "products" SET "quantity" = quantity - 1 WHERE "id" = '2' AND quantity > 1;

```

### 删除
```go

db.Delete(&user)
//// UPDATE users SET deleted_at="2013-10-29 10:23" WHERE id = 111;

// 批量删除
db.Where("age = ?", 20).Delete(&User{})
//// UPDATE users SET deleted_at="2013-10-29 10:23" WHERE age = 20;

// 在查询记录时，软删除记录会被忽略
db.Where("age = 20").Find(&user)
//// SELECT * FROM users WHERE age = 20 AND deleted_at IS NULL;

// 使用 Unscoped 方法查找软删除记录
db.Unscoped().Where("age = 20").Find(&users)
//// SELECT * FROM users WHERE age = 20;

// 使用 Unscoped 方法永久删除记录
db.Unscoped().Delete(&order)
//// DELETE FROM orders WHERE id=10;
```

### 钩子函数

### SQL 选项