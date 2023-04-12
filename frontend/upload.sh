npm run build 
docker build --platform linux/amd64  -t elasticcommunity/atp:frontend  .
docker push  elasticcommunity/atp:frontend 