# Always restore before starting
litestream restore -o /bin/todox.db https://f72ffa7fed95c214a8bcc247a3e59805.r2.cloudflarestorage.com/todox

# Migrate the database
/bin/tools migrate

# Start sever and backup
litestream replicate -exec "/bin/app" /bin/todox.db https://f72ffa7fed95c214a8bcc247a3e59805.r2.cloudflarestorage.com/todox