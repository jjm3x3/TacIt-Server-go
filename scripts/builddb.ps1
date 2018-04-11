# "D:\Program Files\PostgreSQL\10\bin\psql.exe" -h localhost -p 5432 -d postgres -U postgres -f D:\Users\jmeixner\Code\goCode\src\TacIt-go\scripts\builddb.sql

# D:\Program Files\PostgreSQL\10\data\pg_hba.conf modify local ipv4 connections to use "trust" instead of md5

#installing mysql on windows was done manually using the following links guides
# https://dev.mysql.com/doc/refman/5.5/en/windows-start-service.html
# https://stackoverflow.com/questions/34448628/after-install-mysql-doesn-start-windows10-source-install
# https://dev.mysql.com/doc/refman/5.7/en/problems-connecting.html
# https://dev.mysql.com/doc/refman/5.7/en/default-privileges.html
# https://dev.mysql.com/doc/refman/5.7/en/data-directory-initialization-mysqld.html

# should be a relative path 
$providerRoots = get-psdrive | where { $_.Root -ne "" -and $_.Root -ne "\"} | Select-Object -Property Root
$mysqlPath = foreach($a in $providerRoots) { Get-ChildItem -Path $a.Root -Filter mysql.exe -Recurse -ErrorAction SilentlyContinue | select-object -first 1 |  %{ $_.FullName } } 
# TODO:: check that there is only one path


Get-Content $($(Get-Location).path + '\scripts\builddb.sql') | & $mysqlPath -P3306 -u root -padmin mysql