class apache2 {

  file { "/tmp/puppet/apache2":
      ensure => present,
      owner => "root",
      group => "root",
      require => File["/tmp/puppet"],
    }

}
