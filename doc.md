# 接口文档
## /oauth/authorzie/
```
req:
GET /oauth/authorize?client_id=CLIENT_ID&redirect_url=CALLBACK_URL&scope=SCOPE

resp:
if logged:
302 confirm html
if not logged:
302 login html
```
## /oauth/login/
```
req:
POST /oauth/login/ username=xxx&password=xxx
resp:
302 confirm html
```
## /oauth/confirm/
```
req:
POST /oauth/confirm/  confirm=xxx&scope=xxx
resp:
302 rediret_url=xxx&code=xxx
```
