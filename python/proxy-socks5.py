import requests

username = 'username'
password = 'password'
host = 'gw-us.nstproxy.com'
port = '24125'

proxy = f'socks5://{username}:{password}@{host}:{port}'
proxy_dict = {
    "http": proxy,
    "https": proxy
}

response = requests.get("API_URL", proxies=proxy_dict)
print(response.text)