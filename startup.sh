#!/bin/sh
set -ex
mkfs.ext4 /dev/nvme1n1 >> /home/ec2-user/startup.log 2>&1
sh /home/ec2-user/MessageBrokerGo/run.sh >> /home/ec2-user/startup.log 2>&1

