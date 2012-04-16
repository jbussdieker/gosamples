class mysql::install inherits mysql {

	package { "mysql-server":
		ensure => installed,
	}

}
