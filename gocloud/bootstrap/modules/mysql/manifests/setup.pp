class mysql::setup inherits mysql {

	include mysql::install
	include mysql::config
	include mysql::service

}
