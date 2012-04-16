class nginx::config inherits nginx {

  file { "/etc/nginx/nginx.conf":
    ensure => present,
    owner => root,
    group => root,
    mode => 644,
    source => ["puppet:///modules/nginx/nginx.conf"],
    require => Class["nginx::install"],
    notify => Class["nginx::service"],
  }

}
