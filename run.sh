# Always restore before starting
litestream restore -config /etc/litestream.yml

# Migrate the database
/bin/tools migrate

# Start sever and backup
litestream replicate -exec "/bin/app" -config /etc/litestream.yml