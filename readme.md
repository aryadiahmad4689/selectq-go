# SelectQ
SelectQ adalah project open source yang di bangun atas landasan fitur query eloquent laravel yang sangat elegan dan selectQ mencoba mengikuti keindahan itu dan menerapkanya di golang

## how to install
 - kalian bisa ketik perintah di bawah di project kalian
 - ``` go get -u github.com/aryadiahmad4689/selectq-go ```
 
## How To Use
```
selectQ := selectq.Init(context.Background()) // init project
selectQ.NewConnect(db) // masukkan koneksi *sql.DB
SetTable("test") // set table yang di tuju
```
## Fitur Yang Tersedia
- Read
    - select
   ```
    selectQ.Select("username,email").Get(); // select username,email from table
   ```
   - where
   ```
    selectQ.Select("username,email").Where("username =?","ariadi").Get(); // select username,email from table where username ="ariadi"
    
    selectQ.Select("username,email").Where("username =?","ariadi").WhereOr("email =?","ariadi@com").Get(); // select username,email from table where username ="ariadi" or email="riadi@.com"
   ```
   - group by
   ```
    selectQ.Select("username,email").GroupBy("username,email").Get(); // select username,email from table group by username,email
   ```
   - offset and limit
   ```
    selectQ.Select("username,email").Offset(1).Limit(10).Get(); // select username,email from table offset 1 LIMIT 10
   ```
    - order by
   ```
    selectQ.Select("username,email").OrderBy("id asc")Offset(1).Limit(10).Get(); // select username,email from table order by id asc offset 1 LIMIT 10
   ```

   Untuk Melihat Contoh Kasusnya Kalian Bisa Buka Folder Example

 ## Fitur Yang Akan Di Bangun
  - join
  - store data
  - update data
  - delete data
  - first
   
   ### Licence
   MIT Licence
  