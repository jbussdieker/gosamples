node default {
	file { "/tmp/puppet":
		ensure => directory,
		owner => "root",
		group => "root",
		purge => true,
	}

	#include build_essential::setup
	include git::setup
	include apache2php5::setup
	include php5mysql::setup
	include mysql::setup
	include nginx::setup
	include imagemagick::setup
	include apache2::setup
	include php5::setup
	#include rvm::setup
	#include ruby192::setup

	include users::jbussdieker
	include ssh::setup

	include wiki::setup
	include wiki::jbussdieker
	include wordpress::setup
	#include wordpress::jbussdieker

	#include jenkins::setup
}
