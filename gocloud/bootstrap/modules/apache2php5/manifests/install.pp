class apache2php5::install inherits apache2php5 {

	package { "libapache2-mod-php5":
		ensure => installed,
    	require => Package["apache2", "php5"], 
	}

	exec { "/usr/sbin/a2enmod php5":
		unless => "/bin/readlink -e /etc/apache2/mods-enabled/php5.load",
		require => Package["libapache2-mod-php5"],
		notify => Class["apache2::service"],
	}

}
