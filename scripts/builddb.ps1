# "D:\Program Files\PostgreSQL\10\bin\psql.exe" -h localhost -p 5432 -d postgres -U postgres -f D:\Users\jmeixner\Code\goCode\src\TacIt-go\scripts\builddb.sql

# D:\Program Files\PostgreSQL\10\data\pg_hba.conf modify local ipv4 connections to use "trust" instead of md5

#installing mysql on windows was done manually using the following links guides
# https://dev.mysql.com/doc/refman/5.5/en/windows-start-service.html
# https://stackoverflow.com/questions/34448628/after-install-mysql-doesn-start-windows10-source-install
# https://dev.mysql.com/doc/refman/5.7/en/problems-connecting.html
# https://dev.mysql.com/doc/refman/5.7/en/default-privileges.html
# https://dev.mysql.com/doc/refman/5.7/en/data-directory-initialization-mysqld.html

# should be a relative path 
Get-Content D:\Users\jmeixner\Code\goCode\src\TacIt-go\scripts\builddb.sql | .\mysql.exe -u root -padmin mysql