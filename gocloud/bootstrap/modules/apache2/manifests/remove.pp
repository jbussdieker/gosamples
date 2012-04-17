class apache2::remove inherits apache2 {

  package { "apache2_uninstalled":
	name => "apache2",
    ensure => purged,
    require => Service["apache2_stopped"],
  }

  service { "apache2_stopped":
	name => "apache2",
    ensure => stopped,
    hasstatus => true,
  }

}
