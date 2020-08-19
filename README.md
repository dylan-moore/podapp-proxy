# podapp-proxy

Experimental http proxy / ingress controller written in go.

Built by the [podapp](https://getpod.app) team.

Targets:

- Authentication
  - JWT verify
  - Session verify
- Kubernetes support
- Small codebase
- SSL
  - kubernetes cert-manager support
- Websockets
- GeoIP
- Rate limiting
- Load balancing
  - Round robin
  - Sticky sessions
- Statistics

This project uses the MIT license, please read the 'LICENSE' file for more information.
