# ripe-dns-decode
A simple tool to show ripe atlas's DNS measurement result

# How to build?

Please create or find a DNS measurement from ripe atlas and you can retrieve the measurement result url from the Result tab.
And you need to download the ASN (./GeoLite2-ASN.mmdb) and Country (./GeoLite2-Country.mmdb) database from maxmind(https://dev.maxmind.com/geoip/geoip2/geolite2/) and put it in this directory

```
go build ripe-dns-decode.go
```

# How to use?

```
 #./ripe-dns-decode https://atlas.ripe.net/api/v2/measurements/9202394/results/\?start\=1500422400\&stop\=1501113599\&format\=json |head -n 50
 ----------------------
 IP: 83.163.117.153 Country: Netherlands ISP: AS3265 Xs4all Internet BV
 ;; opcode: QUERY, status: NOERROR, id: 30778
 ;; flags: qr tc rd ra; QUERY: 1, ANSWER: 10, AUTHORITY: 8, ADDITIONAL: 0

 ;; QUESTION SECTION:
 ;36-courier.push.apple.com.     IN       A

 ;; ANSWER SECTION:
 36-courier.push.apple.com.      28800   IN      CNAME   36.courier-push-apple.com.akadns.net.
 36.courier-push-apple.com.akadns.net.   60      IN      CNAME   pop-eur-benelux-courier.push-apple.com.akadns.net.
 pop-eur-benelux-courier.push-apple.com.akadns.Net.      17      IN      A       17.252.44.10
 pop-eur-benelux-courier.push-apple.com.akadns.Net.      17      IN      A       17.252.44.80
 pop-eur-benelux-courier.push-apple.com.akadns.Net.      17      IN      A       17.252.44.76
 pop-eur-benelux-courier.push-apple.com.akadns.Net.      17      IN      A       17.252.44.99
 pop-eur-benelux-courier.push-apple.com.akadns.Net.      17      IN      A       17.252.44.30
 pop-eur-benelux-courier.push-apple.com.akadns.Net.      17      IN      A       17.252.44.88
 pop-eur-benelux-courier.push-apple.com.akadns.Net.      17      IN      A       17.252.44.16
 pop-eur-benelux-courier.push-apple.com.akadns.Net.      17      IN      A       17.252.44.28

 ;; AUTHORITY SECTION:
 akadns.Net.     32982   IN      NS      a7-131.akadns.net.
 akadns.Net.     32982   IN      NS      a18-128.akadns.org.
 akadns.Net.     32982   IN      NS      a28-129.akadns.org.
 akadns.Net.     32982   IN      NS      a5-130.akadns.org.
 akadns.Net.     32982   IN      NS      a12-131.akadns.org.
 akadns.Net.     32982   IN      NS      a1-128.akadns.net.
 akadns.Net.     32982   IN      NS      a11-129.akadns.net.
 akadns.Net.     32982   IN      NS      a13-130.akadns.org.
 0
 ----------------------
 IP: 83.68.21.139 Country: Netherlands ISP: AS3265 Xs4all Internet BV
 ;; opcode: QUERY, status: NOERROR, id: 475
 ;; flags: qr rd ra; QUERY: 1, ANSWER: 10, AUTHORITY: 0, ADDITIONAL: 0

 ;; QUESTION SECTION:
 ;36-courier.push.apple.com.     IN       A

 ;; ANSWER SECTION:
 36-courier.push.apple.com.      20624   IN      CNAME   36.courier-push-apple.com.akadns.net.
 36.courier-push-apple.com.akadns.net.   60      IN      CNAME   pop-eur-benelux-courier.push-apple.com.akadns.net.
 pop-eur-benelux-courier.push-apple.com.akadns.net.      57      IN      A       17.252.44.30
 pop-eur-benelux-courier.push-apple.com.akadns.net.      57      IN      A       17.252.44.95
 pop-eur-benelux-courier.push-apple.com.akadns.net.      57      IN      A       17.252.44.82
 pop-eur-benelux-courier.push-apple.com.akadns.net.      57      IN      A       17.252.44.72
 pop-eur-benelux-courier.push-apple.com.akadns.net.      57      IN      A       17.252.44.91
 pop-eur-benelux-courier.push-apple.com.akadns.net.      57      IN      A       17.252.44.22
 pop-eur-benelux-courier.push-apple.com.akadns.net.      57      IN      A       17.252.44.31
 pop-eur-benelux-courier.push-apple.com.akadns.net.      57      IN      A       17.252.44.98
 ```
