class php5mysql {

  file { "/tmp/puppet/php5mysql":
      ensure => present,
      owner => "root",
      group => "root",
      require => File["/tmp/puppet"],
    }

}
