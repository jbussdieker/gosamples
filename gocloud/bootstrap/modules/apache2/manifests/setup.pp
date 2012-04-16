class apache2::setup inherits apache2 {

  include apache2::install
  include apache2::config
  include apache2::service

}
