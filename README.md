# wallet-service


export MONGODB_URL="127.0.0.1"
export DB_USER=""
export DB_PASSWORD=""

export HOST="192.168.1.2"
export PORT=9002
export SERVICE_ID="hchain.wallet-service"
export SERVICE_NAME="hchain.wallet-service"
export CONSUL_HOST="192.168.1.2"
export CONSUL_PORT=8500


docker run -d --name=dev-consul -p 8500:8500 -p 8600:8600  -p 8300:8300 -p 8301:8301 -p 8302:8302 -e CONSUL_BIND_INTERFACE=eth0 consul