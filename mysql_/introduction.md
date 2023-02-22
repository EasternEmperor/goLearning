### Go连接mysql
1. type DB {...}: DB是一个数据库句柄，代表一个具有零到多个底层连接的连接池。可以**安全地
   被多个go协程同时使用**。sql包会**自动**创建和释放连接，它会维护一个闲置连接的连接池。
   连接池的大小由SetMaxIdleConns方法控制

2. `func Open(driverName, dataSourceName string) (*DB, error)`: Open打开一个driverName
   指定的数据库，dataSourceName指定数据源（至少包括数据库文件名和连接信息）
    - Db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/test")
    - dataSourceName参考格式：**数据库用户名:数据库密码@[tcp(localhost:3306)]/数据库名**
    - Open函数可能只是验证其参数，**不创建与数据库的连接**

3. `func (db *DB) Ping() error`: Ping检查与数据库的连接是否仍有效，如果需要会创建连接

4. `func (db *DB) Exec(query string, args ...interface{}) (Result, error)`: 执行一次
   命令**（包括查询、删除、更新、插入等）**，不返回任何执行结果。参数args表示query中的占位
   参数。
    - type Rows struct {...}: Rows是查询的结果，它的游标指向结果集的第0行，使用Next方法
    来遍历各行结果

5. `func (db *DB) Query(query string, args ...interface{}) (*Rows, error)`: 执行一次
   查询，**返回所有结果(Rows)**，用于执行select命令。参数args表示query中的占位参数。
    - func (rs *Rows) Next() bool: Next**准备**用于Scan方法的下一个**结果**。如果成功返回真，如果
    没有下一行或者错误返回假。每次调用Scan方法前（包括第一次），都必须调用Next用于准备结果
    - func (rs *Rows) Scan(dest ...interface{}) error: 将当前行各列结果填入dest指定的各值
    中

6. `func (db *DB) QueryRow(query string, args ...interface{}) *Row`: 执行一次查询，并
   期望**最多返回一行结果(Row)**. QueryRow**总是返回非nil的值**，直到返回值的**Scan方法**被调用时，
   才会返回被延迟的错误（如：未找到结果）
    - func (r *Row) Scan(dest ...interface{}) error: Scan将查询结果分别保存进dest参数指定
    的值中。如果没有匹配查询的行，会返回**ErrNoRows错误**
    
7. `func (db *DB) Prepare(query string) (*Stmt, error)`: 创建一个**准备好的状态**用于之后的
   用于之后的查询和命令。返回的*Stmt可以同时执行多个命令
    - Stmt有类似的Exec(), Query(), QueryRow()方法，其中参数为args ...interface{}，为Prepare
    的**query参数中的占位符**
