class apache2::config inherits apache2 {

  file { "/etc/apache2/ports.conf":
    ensure => present,
    owner => root,
    group => root,
	recurse => true,
    mode => 644,
    source => ["puppet:///modules/apache2/ports.conf"],
    require => Class["apache2::install"],
    notify => Class["apache2::service"],
  }

}

