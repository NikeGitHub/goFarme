#!/bin/bash
# api
etcdctl put  /smartgy/conf/common/api/host 106.75.154.221
etcdctl put  /smartgy/conf/common/api/port 8347
# database
etcdctl put  /smartgy/conf/common/mysql/smartgy/0/database smartgy
etcdctl put  /smartgy/conf/common/mysql/smartgy/0/host 106.75.138.97
etcdctl put  /smartgy/conf/common/mysql/smartgy/0/password epuser@123-New
etcdctl put  /smartgy/conf/common/mysql/smartgy/0/port 3306
etcdctl put  /smartgy/conf/common/mysql/smartgy/0/username epuser
etcdctl put  /smartgy/conf/common/mysql/smartgy/1/database smartgy
etcdctl put  /smartgy/conf/common/mysql/smartgy/1/host 106.75.138.97
etcdctl put  /smartgy/conf/common/mysql/smartgy/1/password epuser@123-New
etcdctl put  /smartgy/conf/common/mysql/smartgy/1/port 3306
etcdctl put  /smartgy/conf/common/mysql/smartgy/1/username epuser
# database
etcdctl put  /smartgy/conf/common/mysql/smartgy/0/database smartgy
etcdctl put  /smartgy/conf/common/mysql/smartgy/0/host rm-2vcwdvj58u4oyn9u1no.mysql.cn-chengdu.rds.aliyuncs.com
etcdctl put  /smartgy/conf/common/mysql/smartgy/0/password O90yH4U^gdC0YTqJ
etcdctl put  /smartgy/conf/common/mysql/smartgy/0/port 3306
etcdctl put  /smartgy/conf/common/mysql/smartgy/0/username consume
etcdctl put  /smartgy/conf/common/mysql/smartgy/1/database smartgy
etcdctl put  /smartgy/conf/common/mysql/smartgy/1/host rm-2vcwdvj58u4oyn9u1no.mysql.cn-chengdu.rds.aliyuncs.com
etcdctl put  /smartgy/conf/common/mysql/smartgy/1/password O90yH4U^gdC0YTqJ
etcdctl put  /smartgy/conf/common/mysql/smartgy/1/port 3306
etcdctl put  /smartgy/conf/common/mysql/smartgy/1/username consume
# grpc
etcdctl put  /smartgy/conf/common/grpc/host 127.0.0.1
etcdctl put  /smartgy/conf/common/grpc/port 12314
# minio
etcdctl put  /smartgy/conf/common/file/host http://106.75.154.221:8321/minio/file