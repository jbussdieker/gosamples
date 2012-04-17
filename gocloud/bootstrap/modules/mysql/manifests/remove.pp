class mysql::remove inherits mysql {

  package { "mysql_uninstalled":
	name => "mysql-server",
    ensure => purged,
    require => Service["mysql_stopped"],
  }

  service { "mysql_stopped":
	name => "mysql",
    ensure => stopped,
    hasstatus => true,
  }

}

