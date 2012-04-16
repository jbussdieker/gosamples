class nginx {

	file { "/tmp/puppet/nginx":
		ensure => present,
		owner => "root",
		group => "root",
		require => File["/tmp/puppet"],
	}

}
