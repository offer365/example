wget http://mirror.bit.edu.cn/apache/kafka/2.3.0/kafka_2.12-2.3.0.tgz
tar xf kafka_2.12-2.3.0.tgz -C /usr/local/
tar xf jdk-8u101-linux-x64.tar.gz -C /usr/local/
msg="export JAVA_HOME=/usr/local/jdk1.8.0_101\nexport PATH=\${JAVA_HOME}/bin:\$PATH\nexport LD_LIBRARY_PATH=/usr/lib:/usr/lib64:/usr/local/lib64:/usr/local/lib:/home/admin/diting/lib:/home/admin/speech-alisr/lib64"
echo -e $msg >>/etc/profile
source /etc/profile
java -version
cd /usr/local/kafka_2.12-2.3.0/
sed -i "/^broker.id=/ s/0/1/g" config/server.properties
# 杩藉姞 advertised.host.name=kafka鏈嶅姟鍣╥p kafka 閰嶇疆鏂囦欢 config/server.properties
echo "advertised.host.name=10.0.0.55" >>config/zookeeper.properties
./bin/zookeeper-server-start.sh -daemon config/zookeeper.properties
./bin/kafka-server-start.sh config/server.properties

# 鍛戒护琛屽紑鍚痥afka娑堣垂鑰呭鎴风鍛戒护
# ./bin/kafka-console-consumer.sh --bootstrap-server 127.0.0.1:9092 --topic asr_log
# Note锛氬湪0.9鐗堟湰鎸囧畾鐨勬槸zookeeper server 0.11鍙樻垚浜哹roker server
# ./bin/kafka-console-consumer.sh -zookeeper localhost:2181 --topic asr_log
