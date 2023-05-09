# Always restore before starting
litestream restore

# Migrate the database
/bin/tools migrate

# Start sever and backup
litestream replicate -exec "/bin/app" /bin/todox.db