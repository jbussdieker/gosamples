class apache2::service inherits apache2 {
  service { "apache2":
    ensure => running,
    hasstatus => true,
  }
}
