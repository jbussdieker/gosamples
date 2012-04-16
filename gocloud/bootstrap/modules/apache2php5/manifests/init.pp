class apache2php5 {

  file { "/tmp/puppet/apache2php5":
      ensure => present,
      owner => "root",
      group => "root",
      require => File["/tmp/puppet"],
    }

}
