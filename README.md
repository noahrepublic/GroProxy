# GroProxy

A small light-weight **private** proxy server to access Roblox's public API.

# Why is this a thing?

While Roblox provides built in API for lots of things, they do not let you call their own web API. Games like Pls Donate have to use something like this.

# Why not RoProxy?

Honestly, for most cases RoProxy is good. **However**

RoProxy is a public proxy,  if you require to use APIs that use private information there may be security issues as it is not open-source. 

Furthermore, you are relying on his server. If his server goes down so does yours. More flaws in public proxies [here.](https://devforum.roblox.com/t/psa-stop-using-roblox-proxies-roproxy-rprxyxyz-rprxy/1573256)


# Features

- Uses "Gin" a HTTP framework 40 times faster then Martini.
- Supports API_KEY validation through .env files
- Port setting and IP address through .env files
# Planned

- Proxy List Support
- Pre-Configured deployment buttons to known providers