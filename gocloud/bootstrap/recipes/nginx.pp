node default {
	file { "/tmp/puppet":
		ensure => directory,
		owner => "root",
		group => "root",
		purge => true,
	}
/*	file { "/tmp/puppet/sdf/qwer/sdef":
		ensure => present,
		owner => "root",
		group => "root",
		purge => true,
	}*/
	include nginx::setup
}
