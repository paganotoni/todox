# Always restore before starting
litestream restore -o /bin/cach.db s3://f72ffa7fed95c214a8bcc247a3e59805.r2.cloudflarestorage.com/todox

# Migrate the database
/bin/tools migrate

# Start sever and backup
litestream replicate -exec "/bin/app" /bin/cach.db s3://f72ffa7fed95c214a8bcc247a3e59805.r2.cloudflarestorage.com/todox