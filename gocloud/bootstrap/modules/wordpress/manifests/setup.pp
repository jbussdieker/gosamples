class wordpress::setup inherits wordpress {

	include wordpress::install
	include wordpress::config

}
