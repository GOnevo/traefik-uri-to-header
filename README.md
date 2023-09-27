[![Go Report Card](https://goreportcard.com/badge/github.com/gonevo/traefik-uri-to-header)](https://goreportcard.com/report/github.com/gonevo/traefik-uri-to-header)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/gonevo/traefik-uri-to-header)
![GitHub](https://img.shields.io/github/license/gonevo/traefik-uri-to-header)

# traefik-uri-to-header

A plugin that writes the requested URI to a custom header.

> Important: URI is stored without leading slash.

### Parameters

- enabled: (boolean) should the plugin work 
- headerName: (string) the name of custom header 

### Usage

Add plugin to your traefik configuration:
```yaml
experimental:
  plugins:
    uri2header:
      moduleName: github.com/gonevo/traefik-uri-to-header
      version: v0.0.1
```

or 

```shell
--experimental.plugins.uri2header.moduleName=github.com/gonevo/traefik-uri-to-header
--experimental.plugins.uri2header.version=v0.0.1
```

Then enable it:
```yaml
http:
  routers:
    my-router:
      rule: host(`domain.com`)
      entryPoints:
        - web
      middlewares:
        - my-plugin
  
  middlewares:
    my-plugin:
      plugin:
        uri2header:
          enabled: true
          headerName: x-custom-uri
          
```
or
```shell
"traefik.http.routers.my-router.middlewares=my-plugin"
"traefik.http.middlewares.my-plugin.plugin.uri2header.enabled=true"
"traefik.http.middlewares.my-plugin.plugin.uri2header.headerName=x-custom-uri"
```

### License

**traefik-uri-to-header** is open-sourced software licensed under the [MIT license](./LICENSE.md).

[Vano Devium](https://github.com/vanodevium/)

---

Made with ❤️ in Ukraine
