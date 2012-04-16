class wordpress::config inherits wordpress {

	file { "/etc/apache2/sites-available/wordpress":
		ensure => present,
		owner => root,
		group => root,
		mode => 644,
		source => ["puppet:///modules/wordpress/wordpress"],
		require => Class["apache2::install"],
		notify => Class["apache2::service"],
	}

	exec { "/usr/sbin/a2ensite wordpress":
		unless => "/bin/readlink -e /etc/apache2/sites-enabled/wordpress",
		require => File["/etc/apache2/sites-available/wordpress"],
		notify => Class["apache2::service"],
	}

    mysql::db::mysqldb { "wordpress":
        user => "root",
        password => "",
    }

}

