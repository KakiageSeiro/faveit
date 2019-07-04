#!/bin/bash -e

echo waiting db for 30 sec
sleep 1

echo db init start
mysql -u root -h db -e 'create database if not exists faveit character set utf8mb4 collate utf8mb4_bin'
mysql -u root -h db faveit < init-tables.sql

echo db init finish
