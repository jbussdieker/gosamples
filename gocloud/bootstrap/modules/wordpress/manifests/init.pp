class wordpress {

	file { "/tmp/puppet/wordpress":
		ensure => present,
		owner => "root",
		group => "root",
		require => File["/tmp/puppet"],
	}

}

