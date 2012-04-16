class mysql::db inherits mysql {

	define mysqldb( $user, $password ) {
		exec { "create-${name}-db":
			unless => "/usr/bin/mysql -u${user} -p${password} ${name}",
			command => "/usr/bin/mysql -uroot -e \"create database ${name}; grant all on ${name}.* to ${user}@localhost identified by '$password';\"",
			require => Service["mysql"],
		}
	}

}
