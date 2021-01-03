package models

import (
	"BitcoinProject/db_mysql"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

type User struct {
	Id       int    `form:"id"`
	UserName string `form:"username"`
	Password string `form:"password"`
}

//保存用户信息的方法：保存用户信息到数据库中

func (u User) SaverUser() (int64, error) {
	//	1.密码脱敏处理
	hashMd5 := md5.New()
	hashMd5.Write([]byte(u.Password))
	byte := hashMd5.Sum(nil)
	u.Password = hex.EncodeToString(byte)

	//	2。执行数据库操作
	//row,err:=db_mysql.Db.Exec("insert into user ")
	row, err := db_mysql.Db.Exec("insert into users (username,password)"+"values(?,?)", u.UserName, u.Password)
	if err != nil {
		return -1, err
	}
	id, err := row.RowsAffected()
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (u User) QueryUser() (*User, error) {
	//fmt.Println(u.Password)
	hashMd5 := md5.New()
	hashMd5.Write([]byte(u.Password))
	//hashMd5.Write([]byte(u.Password))
	byte := hashMd5.Sum(nil)
	u.Password = hex.EncodeToString(byte)
	//fmt.Println(u.Password)
	row, err := db_mysql.Db.Query("select username from users where username = ? and  password=?", u.UserName,u.Password)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for row.Next() {
		err := row.Scan(&u.UserName)
		//err := row.Scan(&u.Id,&u.UserName,&u.Password,&u.Phone)
		if err != nil {
			fmt.Println("----------",err)
			return nil, err
		}
	}
	//fmt.Println("000000000",u)
	return &u, nil
}

/**
* 根据数据库存储信息查询
*/
//func (u User) QueryUser11(phone string, pwd string) (*User, error) {
//	fmt.Println(u.Password)
//	hashMd5 := md5.New()
//	hashMd5.Write([]byte(pwd))
//	//hashMd5.Write([]byte(u.Password))
//	byte := hashMd5.Sum(nil)
//	u.Password = hex.EncodeToString(byte)
//	fmt.Println(u.Password)
//	row, err := db_mysql.Db.Query("select username from users where username = ? and  password=?",phone, u.Password)
//	if err != nil {
//		fmt.Println(err)
//		return nil, err
//	}
//	for row.Next() {
//		err := row.Scan(&u.UserName)
//		//err := row.Scan(&u.Id,&u.UserName,&u.Password,&u.Phone)
//		if err != nil {
//			fmt.Println("----------",err)
//			return nil, err
//		}
//	}
//	//fmt.Println("000000000",u)
//	return &u, nil
//}
