class php5mysql::install inherits php5mysql {

	package { "php5-mysql":
		ensure => installed,
    	require => Package["mysql-server", "php5"],
	}

}
