# SelectQ
SelectQ adalah project open source yang di bangun atas landasan fitur query eloquent laravel yang sangat elegan dan selectQ mencoba mengikuti keindahan itu dan menerapkanya di golang

## how to install
 - kalian bisa ketik perintah di bawah di project kalian
 - ``` go get -u github.com/aryadiahmad4689/selectq-go ```
 
## How To Use
```
selectQ := selectq.Init(context.Background(),db) // init project
selectQ.Read.SetTable("test") // set table yang di tuju
```
## Support DB
 - Postgres


## Fitur Yang Tersedia
- Read
    - select
   ```
    selectQ.Read.Select("username,email").Get(); // select username,email from table
   ```
   - where
   ```
    selectQ.Read.Select("username,email").Where("username =?","ariadi").Get(); // select username,email from table where username ="ariadi"
    
    selectQ.Read.Select("username,email").Where("username =?","ariadi").WhereOr("email =?","ariadi@com").Get(); // select username,email from table where username ="ariadi" or email="riadi@.com"
   ```
   - group by
   ```
    selectQ.Read.Select("username,email").GroupBy("username,email").Get(); // select username,email from table group by username,email
   ```
   - offset and limit
   ```
    selectQ.Read.Select("username,email").Offset(1).Limit(10).Get(); // select username,email from table offset 1 LIMIT 10
   ```
    - order by
   ```
    selectQ.Read.Select("username,email").OrderBy("id asc")Offset(1).Limit(10).Get(); // select username,email from table order by id asc offset 1 LIMIT 10
   ```
     - join
   ```
    selectQ.Read.Select("chanel").InnerJoin("order_detail", "id", "order_header.id").Where("order_header.id =$1", "4").WhereOr("order_header.id !=$2", "10").GroupBy("order_header.id").Limit(10).OrderBy("order_header.id asc").Get()

    selectQ.Read.Select("chanel").LeftJoin("order_detail", "id", "order_header.id").Where("order_header.id =$1", "4").WhereOr("order_header.id !=$2", "10").GroupBy("order_header.id").Limit(10).OrderBy("order_header.id asc").Get()

    selectQ.Read.Select("chanel").RightJoin("order_detail", "id", "order_header.id").Where("order_header.id =$1", "4").WhereOr("order_header.id !=$2", "10").GroupBy("order_header.id").Limit(10).OrderBy("order_header.id asc").Get()
   ```

   - save data
   ```
    selectQ.Create.AddFied("username","aryadiahmad").AddFied("email","ariadi@gmail.com").Save(ctx)
   ```
    
    - Full Combine
    ```
     selectQ.Read.Select("id").Where("id =$1", "4").WhereOr("id !=$2", "10").GroupBy("id").Limit(10).OrderBy("id asc").Get()
    ```

   Untuk Melihat Contoh Kasusnya Kalian Bisa Buka Folder Example

 ## Fitur Yang Akan Di Bangun
  - update data
  - delete data
  - first
   
   ### Licence
   MIT Licence
  