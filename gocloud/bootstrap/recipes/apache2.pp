node default {
	file { "/tmp/puppet":
		ensure => directory,
		owner => "root",
		group => "root",
		purge => true,
	}
	include apache2::setup
}
