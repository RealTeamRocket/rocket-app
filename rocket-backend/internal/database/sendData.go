//That could be the basic insert into for
//new users into the database table


var username string = "username" //here data from dart
var email string = "email"     //here data from dart
var firstname string = "firstname" //here data from dart
var lastname string = "lastname" //here data from dart

sqlStatement := `INSERT INTO users VALUES ($1, $2, $3, $4)`

err = db.QueryRow(sqlStatement, username, email, firstname, lastname).Scan(&id)