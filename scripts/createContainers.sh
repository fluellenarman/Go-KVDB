echo "createContainers.sh"

PORTS=(8080 8081 8082)

# for port in "${ports[@]}"; do
#     docker run -d -p "$port:8080" -e PORT=8080 testkvdb
# done

docker run -d -p 8080:8080 -e PORT=8080 testkvdb
docker run -d -p 8081:8080 -e PORT=8080 testkvdb
docker run -d -p 8082:8080 -e PORT=8080 testkvdb

echo "Containers created"