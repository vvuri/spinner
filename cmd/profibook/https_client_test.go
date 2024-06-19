package profibook

// $ openssl genrsa -out server.key 2048
// $ openssl ecparam -genkey -name secp384r1 -out server.key
// $ openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
// --> server.cr, server.key

// если сертификат является самозаверяющим, для работы HTTPS-клиента нужно использовать
// параметр InsecureSkipVerify: true в структуре http.Transport

// сертификат для клиента
// openssl req -x509 -nodes -newkey rsa:2048 -keyout client.key -out
// --> client.crt, client.key
