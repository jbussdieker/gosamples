class nginx::remove inherits nginx {

#  file { "/etc/nginx/nginx.conf":
#    ensure => purged,
#	force => true,
#    require => Package["nginx_uninstalled"],
#  }

  package { "nginx_uninstalled":
	name => "nginx",
    ensure => purged,
    require => Service["nginx_stopped"],
  }

  service { "nginx_stopped":
	name => "nginx",
    ensure => stopped,
    hasstatus => true,
  }

}
