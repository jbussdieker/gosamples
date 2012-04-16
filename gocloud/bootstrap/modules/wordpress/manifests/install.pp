class wordpress::install inherits wordpress {

	file { "/srv/wordpress-3.3.1.tar.gz":
		ensure => present,
		owner => root,
		group => root,
		mode => 644,
		source => ["puppet:///modules/wordpress/wordpress-3.3.1.tar.gz"],
	}

	exec { "untar wordpress": 
		command => "tar -zxvf wordpress-3.3.1.tar.gz", 
		cwd     => "/srv", 
		path    => "/bin",
		creates => "/srv/wordpress",
		require => File["/srv/wordpress-3.3.1.tar.gz"],
	}

}

