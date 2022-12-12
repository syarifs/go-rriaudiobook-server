#!/usr/bin/sh

DEFAULT_PROJECT_NAME=$(sed -n '1p' go.mod)
DEFAULT_PROJECT_NAME=$(echo $DEFAULT_PROJECT_NAME | sed -e "s/module go-\(.*\)-server/\1/")

echo -n "Current Project Name: "
echo $DEFAULT_PROJECT_NAME
echo -n "Set Project Name: "
read  PROJECT_NAME

if [[ -n $PROJECT_NAME ]]; then
	echo "Setting Project Name from $DEFAULT_PROJECT_NAME to $PROJECT_NAME"
	LS=$(find * -type f)
	for ls in $LS; do
		sed -i "s/$DEFAULT_PROJECT_NAME/$PROJECT_NAME/g" $ls
	done
else
	echo ""
	echo "No \$PROJECT_NAME variable passed, Cancelled"
fi
