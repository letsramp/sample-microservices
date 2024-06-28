gh codespace ports visibility 4000:public -c $CODESPACE_NAME
# Start Skyramp
cd /workspaces/sample-microservices/skyramp-server && npm install && npm run start > output.log 2>&1 &
docker compose -f /workspaces/sample-microservices/skyramp/docker/demo/docker-compose.yml up -d --wait || true