#!/bin/bash
RECIPE=$1

# Install puppet
which puppet &> /dev/null
if [[ "$?" != "0" ]]; then
	APTFLAGS="-q -y"
	apt-get update ${APTFLAGS}
	apt-get install puppet ${APTFLAGS}
fi

# Run local puppet
puppet apply --detailed-exitcodes --modulepath=/home/ubuntu/bootstrap/modules -v bootstrap/recipes/$RECIPE.pp

# There were changes
if [[ "$?" == "2" ]]; then
	exit 0
fi

# There were errors
if [[ "$?" == "4" ]]; then
	exit 1
fi

