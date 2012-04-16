node default {
	file { "/tmp/puppet":
		ensure => directory,
		owner => "root",
		group => "root",
		purge => true,
	}

	include apache2php5::setup
	include php5mysql::setup
	include mysql::setup
	include nginx::setup
	include apache2::setup
	include php5::setup
	include wordpress::setup
}
