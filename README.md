# safeharborcmdclient
The command line client enables one to access a SafeHarbor server from a command line.
The syntax is,
```
  safeharborcmdclient [-option]* command [arg]*
```
The options are,
```
  -help
  -s The URL protocol scheme: one of "http", "https", or "unix". The default is "http".
  -h The host IP address or DNS name.
  -p The TCP port.
  -u The user Id (if the requested command requires authentication).
  -w The password (if a user id is specified).
```
For example,
```
  safeharborcmdclient -h 123.123.123.123 -p 6000 Ping
```
calls the SafeHarbor `Ping` REST function on the server at 123.123.123.123:6000.
As another example,
```
  safeharborcmdclient -h 123.123.123.123 -p 6000 GetGroupDesc 6789
```
calls the `GetGroupDesc` SafeHarbor REST function, supplying 6789 for the requried `Groupid` parameter.
The result that is returned is always a JSON structure containing the result, or an error.
