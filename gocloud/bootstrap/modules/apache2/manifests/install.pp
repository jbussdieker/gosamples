class apache2::install inherits apache2 {

  package { "apache2":
    ensure => installed,
  }

}
