@echo off

start java -jar "tools/tika-server-1.13.jar"
start sf -serve "localhost:2040"