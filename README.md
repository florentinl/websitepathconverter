# Traefik Website path converter middleware

This middleware is used to serve Static Websites from a service that exposes static files on the route "/<directory_reserved_for_my_website>/<files_for_my_website>".

It achieves this by tampering the path according to the following rules:

- /path/to/exact/resource.png -> /mywebsite/path/to/exact/resource.png
- /path/to/virtual/route -> /mywebsite/index.html

The Ingress/Ingressroute is then responsible to redirect this request to the service that exposes static files

Sample Middleware

```yaml

```
