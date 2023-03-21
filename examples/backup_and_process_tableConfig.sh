#!/bin/bash
mkdir -p backups
for i in `startree table listTables | jq -r '.tables | .[]'`
do
echo | startree table getConfig --tableName $i | jq '.' > $i.json 


if [ `jq -r .realtime.tableType $i.json` = "REALTIME" ]
then echo | jq '.realtime' $i.json > $i.tableconfig.json 
else echo | jq '.offline' $i.json > $i.tableconfig.json 
fi

echo | jq '.schema' $i.json > $i.schema.json

mv $i.json backups
mv $i.*.json backups

printf "$i backup created \n"
done
printf "All backups created \n"
