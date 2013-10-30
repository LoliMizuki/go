// HOLD, 因為 sqlite3 的 driver 有 gcc 的 bug 所以無法執行

package main

import (
	"database/sql"

	// "fmt"
)

func main() {
	_, err := sql.Open("sqlite3", "gameDB.sqlite")
	if err != nil {
		panic(err)
	}
}

/*
CREATE TABLE userinfo (
    uid INT(10) NOT NULL,
    username VARCHAR(64) NULL DEFAULT NULL,
    departname VARCHAR(64) NULL DEFAULT NULL,
    created DATE NULL DEFAULT NULL,
    PRIMARY KEY (uid)
);

CREATE TABLE `userdetail` (
    `uid` INT(10) NOT NULL DEFAULT '0',
    `intro` TEXT NULL,
    `profile` TEXT NULL,
    PRIMARY KEY (`uid`)
)
*/
