
### Production Mode (docker)

```
docker run --detach --publish=3000:3000 --name=whatsapp --restart=always --volume=$(docker volume create --name=whatsapp):/app/storages <image> --autoreply="Dont't reply this message please" --webhook="http://yourwebhook.site/handler" --basic-auth=kemal:secret,toni:password,userName:secretPassword
```
