$mysql_password = "cheese"

class mysql::config inherits mysql {

	exec { "set-mysql-password":
		unless => "/usr/bin/mysqladmin -uroot -p${mysql_password} status",
		command => "/usr/bin/mysqladmin -uroot password ${mysql_password}",
		require => Service["mysql"],
	}

}

