# gogogadget
A sample app for playing with microservices in PWS

## PWS Deploy

```sh
go get github.com/tools/godep
godep save
echo 'web: gogogadget' > Procfile
cf push $APP_NAME
```

## When problems arise

Notice that if you don't have a service bound to the app
which provides a DATABASE_URL, the app will fail to boot and
produce no output in the logs beyond "app crashed".

Once you bind a DATABASE_URL service, the app will boot and
output will show up in the logs.
