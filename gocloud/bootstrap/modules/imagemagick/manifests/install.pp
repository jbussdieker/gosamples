class imagemagick::install inherits imagemagick {

	package { "imagemagick":
		ensure => present
	}

}
