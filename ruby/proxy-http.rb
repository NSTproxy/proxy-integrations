require "uri"
require 'net/http'

proxy_host = 'gw-us.nstproxy.com'
proxy_port = 24125
proxy_user = 'username'
proxy_pass = 'password'

uri = URI.parse('API_URL')
proxy = Net::HTTP::Proxy(proxy_host, proxy_port, proxy_user, proxy_pass)

req = Net::HTTP::Get.new(uri.path)

result = proxy.start(uri.host, uri.port) do |http|
   http.request(req)
end

puts result.body