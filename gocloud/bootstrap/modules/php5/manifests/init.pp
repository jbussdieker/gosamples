class php5 {

  file { "/tmp/puppet/php5":
      ensure => present,
      owner => "root",
      group => "root",
      require => File["/tmp/puppet"],
    }

}
