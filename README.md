# Traefik method whitelist

Traefik plugin to allow only specific HTTP methods.


## Config example

### Static configuration

YAML:
```yml
experimental:
  plugins:
    method-whitelist:
        moduleName: github.com/antoniomacri/traefik-method-whitelist
        version: v1.0.0
```

CLI:
```
--experimental.plugins.method-whitelist.modulename=github.com/antoniomacri/traefik-method-whitelist
--experimental.plugins.method-whitelist.version=v1.0.0
```

### Dynamic configuration

```yml
http:
  routers:
    my-router:
      rule: host(`demo.localhost`)
      service: service-foo
      entryPoints:
        - web
      middlewares:
        - my-plugin

  services:
   service-foo:
      loadBalancer:
        servers:
          - url: http://127.0.0.1:5000
  
  middlewares:
    my-plugin:
      plugin:
        method-whitelist:
          Message: "Method Not Allowed"
          Methods:
            - GET
            - POST
```
