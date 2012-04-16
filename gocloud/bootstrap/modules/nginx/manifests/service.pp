class nginx::service inherits nginx {
  service { "nginx":
    ensure => running,
    hasstatus => true,
  }
}
