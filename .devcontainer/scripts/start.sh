gh codespace ports visibility 4000:public -c $CODESPACE_NAME
# Start Skyramp
cd $CODESPACE_VSCODE_FOLDER/skyramp-server && npm i && npm run start &
docker compose -f /workspaces/sample-microservices/skyramp/docker/demo/docker-compose.yml up -d --wait || true