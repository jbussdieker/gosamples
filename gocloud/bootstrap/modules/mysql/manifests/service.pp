class mysql::service inherits mysql {

	service { "mysql":
		enable => true,
		ensure => running,
		hasstatus => true,
		require => Package["mysql-server"],
	}

}

