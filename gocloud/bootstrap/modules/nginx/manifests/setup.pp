class nginx::setup inherits nginx {

  include nginx::install
  include nginx::config
  include nginx::service

}
