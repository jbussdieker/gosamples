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
puppet apply --modulepath=/home/ubuntu/bootstrap/modules -v bootstrap/$RECIPE.pp || exit 1

