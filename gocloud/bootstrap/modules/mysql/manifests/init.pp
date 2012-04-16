class mysql {

  file { "/tmp/puppet/mysql":
      ensure => present,
      owner => "root",
      group => "root",
      require => File["/tmp/puppet"],
    }

}
