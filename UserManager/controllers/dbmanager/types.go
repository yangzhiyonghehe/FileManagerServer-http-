package dbmanager

type mysql_data struct{
	Mysql_name  string  `json:"mysql_name"`
	Mysql_port  int  	`json:"mysql_port"`
	Mysql_user	string  `json:"mysql_user"`
	Mysql_pass  string  `json:"mysql_pass"`
}