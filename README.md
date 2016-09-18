# safeharborcmdclient
The command line client enables one to access a SafeHarbor server from a command line.
The syntax is,
  safeharborcmdclient [-option]* command [arg]*
The options are,
  -help
  -s The URL protocol scheme: one of "http", "https", or "unix".
  -h The host IP address or DNS name.
  -p The TCP port.
  -u The user Id (if the requested command requires authentication).
  -w The password ( " " )
For example,
  safeharborcmdclient -h 123.456.234.567 -p 6000 ping
calls the SafeHarbor ping REST function on the server at 123.456.234.567:6000.
As another example,
  safeharborcmdclient -h 123.456.234.567 -p 6000 getGroupDesc 6789
calls the getGroupDesc SafeHarbor REST function, supplying 6789 for the requried Groupid parameter.
The result that is returned is always a JSON structure containing the result, or an error.
