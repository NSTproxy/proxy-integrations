<?php
$username = 'USERNAME';
$password = 'PASSWORD';
$proxy = 'gw-us.nstproxy.com:24125';
$query = curl_init('API_URL');
curl_setopt($query, CURLOPT_RETURNTRANSFER, 1);
curl_setopt($query, CURLOPT_PROXY, "http://$proxy");
curl_setopt($query, CURLOPT_PROXYUSERPWD, "$username:$password");
$output = curl_exec($query);
curl_close($query);
if ($output)
    echo $output;
?>