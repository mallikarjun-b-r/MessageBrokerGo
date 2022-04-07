set -ex
export PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/usr/local/go/bin

mkdir -p /tmp/data
umount /tmp/data || true
mount -t ext4 -o defaults,rw,seclabel,noatime,nodiratime,journal_async_commit,nobarrier,data=writeback /dev/nvme1n1 /tmp/data
rm -rf /tmp/data/lost+found

/home/ec2-user/MessageBrokerGo/build/abyss
