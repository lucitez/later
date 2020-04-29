dropdb later_test
createdb later_test

for i in ../sql/[0-9]*; do
    psql -f $i later_test
done

psql -f ../sql/setup_test.sql later_test

echo 'finished setting up test db later_test'