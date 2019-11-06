#!/bin/bash
home=$(
  cd $(dirname $0)
  pwd
)

wget http://downloads.mongodb.org/linux/mongodb-linux-x86_64-rhel70-v4.2-latest.tgz
mkdir -p /home/admin
tar xf mongodb-linux-x86_64-rhel70-*.tgz -C /home/admin/
mv /home/admin/mongodb-linux* /home/admin/mongodb
mkdir -p /home/admin/mongodb/{conf,db,logs}
cp mongodb.conf /home/admin/mongodb/conf/
cp mongodb.service /usr/lib/systemd/system/
echo "never" >/sys/kernel/mm/transparent_hugepage/enabled
echo "never" >/sys/kernel/mm/transparent_hugepage/defrag
systemctl enable mongodb
systemctl start mongodb
