#!/bin/bash
mkdir -p backups
for i in `startree table listTables | jq -r '.tables | .[]'`
do
echo | startree table getConfig --tableName $i | jq 'del(..|select(. == null))' > $i.json 
mv $i.json backups
printf "$i backup created \n"
done
printf "All backups created \n"
