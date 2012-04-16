class nginx::install inherits nginx {

  package { "nginx":
    ensure => installed,
  }

}
