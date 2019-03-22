package dbmanager

import (
	"fmt"
		"log"
		"flag"
		"strconv"
		"github.com/larspensjo/config"
		"database/sql"
		_ "github.com/go-sql-driver/mysql"
  )

//D:\GO_LuaProjecy\src\FileManagerServer\UserManager\conf
  var filepath ="D:/GO_LuaProjecy/src/FileManagerServer/UserManager/conf/mysql.ini"


  var(
	  configFile = flag.String("configfile","../../conf/mysql.ini","General configuration file")
		TOPIC = make(map[string]string)
		mysql_conf = mysql_data{"", 0, "", ""}
	)



  func ReadConfig() error {
		flag.Parse()
		cfg,err := config.ReadDefault(*configFile)

		if err != nil {
			log.Fatalf("Fail to find %v,%v",*configFile,err)
			}

			if	cfg.HasSection("mysql") {   //判断配置文件中是否有section（一级标签）
				options,err := cfg.SectionOptions("mysql")    //获取一级标签的所有子标签options（只有标签没有值）
				if err == nil {
					for _,v := range options{
						optionValue,err := cfg.String("mysql",v)  //根据一级标签section和option获取对应的值
						if err == nil {
							TOPIC[v] =optionValue
						}else{
							log.Fatalf("err: %s" , err.Error())
							return err
						}
					}
				}else{
					log.Fatalf("err: %s", err.Error())
					return err
				}
		}else{
			log.Fatal("no section mysql")
			return err
		}
		//log.Println(TOPIC)

		mysql_conf.Mysql_name = TOPIC["mysql_name"]
		mysql_conf.Mysql_port, _ = strconv.Atoi( TOPIC["mysql_port"]) 
		mysql_conf.Mysql_user = TOPIC["mysq_usr"]
		mysql_conf.Mysql_pass = TOPIC["mysql_pass"]

		log.Printf("%v", mysql_conf)

		return nil
}



func insert(db *sql.DB, sqlstr string){
		stmt, err := db.Prepare(sqlstr)
		defer stmt.Close()

		if err != nil {

			log.Println(err)
			
			return
			
		}
}


func initmysql()(*sql.DB, error){
		err := ReadConfig()
		if err != nil {
				return nil, err
		}

		sqlinit := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", mysql_conf.Mysql_user, mysql_conf.Mysql_pass, mysql_conf.Mysql_name, mysql_conf.Mysql_port, "test")
		

		db, err := sql.Open("mysql", sqlinit)
		if err != nil{
			log.Fatalf("err in open mysql: %s", err.Error())
			return nil, err
		}


		defer db.Close()
 
		err = db.Ping()
		if err != nil {
			log.Fatalf("err : %s", err.Error())
			return nil, err
		}



		return db, nil
	}


