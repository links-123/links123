_NGINX_TLS_PATH="nginx/tls"
_NGINX_TLS_NAME="api.links"

generate-nginx-tls-cert: clean-nginx-tls-cert
	openssl req \
      -x509 -nodes -days 365 -newkey rsa:2048 \
      -subj "/CN=localhost" \
      -keyout "$(_NGINX_TLS_PATH)/$(_NGINX_TLS_NAME).key" \
      -out "$(_NGINX_TLS_PATH)/$(_NGINX_TLS_NAME).crt"

clean-nginx-tls-cert:
	rm -f $(_NGINX_TLS_PATH)/*.key && \
	rm -f $(_NGINX_TLS_PATH)/*.crt

clean: clean-nginx-tls-cert

.SILENT: clean-tls-cert
