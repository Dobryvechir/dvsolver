Copyright by Volodymyr Dobryvechir dobrivecher@yahoo.com vdobryvechir@gmail.com

### Installation and Getting Started

Place the file dvserver (linux) or dvserver.exe (Windows) to any folder,
so that it will run from any folder.

For Windows: You can create a folder c:\\Program Files\\DvServer and
place dvserver.exe in this folder. Then add this path to the Path
variable in your environment

For Linux: You can copy the dvserver to /usr/bin folder.

The behaviour of the program completely depends on the configuration you
provide in DvServer.conf. You can have as many configs as you wish:
place them to different directories. When you run dvserver, it starts to
search for DvServer.conf in the current directory. If it cannot find it
in the current directory, it starts to look for it in
{user-folder}/DvServer/{brief name of the current directory}. If the
config is not found, it uses a default config, which means that the
DvServer starts to listen at port 80 and serves as a simple static
server in the current directory.

The directory {user-folder}/DvServer/{brief name of the current
directory} is considered as the current work folder (logs, cruds and so
on), unless you specify the DV\_SERVER\_CURRENT\_NAMESPACE in the
environment variables. If you specify this variable, the current
namespace will be
{user-folder}/DvServer/{DV\_SERVER\_CURRENT\_NAMESPACE}. In addition to
the environment variables you can add some other variables, which can be
used everywhere (in config, templates and other files). The file with
your variables must have the name DvServer.properties The search for
this file is made at first in the current directory, than in the
directory of your current namespace, than in the directory
{DV\_SERVER\_PATH} if you specify this directory in the environment
variables.

To create a config, at first create an initial config (see Command line
options dvserver config) and then extend the config depending on your
tasks. Possible tasks and instructions how to fill the config for those
tasks are described below.

### Command line options

Usually dvserver is run without any parameters in a folder with a
configuration stored in the DvServer.conf file: dvserver

There are some special cases where the command line is necessary

dvserver config or dvserver config .

saves a sample config in the current directory for the current directory

dvserver config \#

saves a sample config for the current directory in the current user's
dvserver folder. It creates all necessary subfolders in the users
directory, starting with {user-folder}/DvServer/{brief name of the
current directory}. When the program finishes, it outputs the name of
the folder where the config has been created.

dvserver finish

finishes the work of the dvserver and removes all hooks from the system.
(Logs and data folders are not removed.)

### Task: start a server listening at a specific port {#task_listen}

For this task, you should provide the port (1-65535) number and
optionally the local IP

For example, for port \#80, you should specify in the config:

"listen": ":80"

This is already provided in the default config

Specifying IP is not necessary in the majority of cases, but if you wish
to specify, for example, 127.0.0.2, you can write in the config as
follows:

"listen": "127.0.0.2:80"

If you specify IP, you will have to use only this IP for local
development. If you do not specify IP, this port will be valid for all
local hosts (local hosts are all hosts whose IP starts with 127
(127.X.X.X))

If the listener is running, it can find data in local file system (so
server as a local server), process crud options or forward requests to
external servers

### Task: set host names for local servers

This task is solved in any system in file /etc/hosts (linux) or
{windows}/System32/drivers/etc/hosts

DvServer can automatically add/remove your host names in these files

If you specify hosts in the config, they will be added when you start
dvserver and will be removed if you run dvserver finish

The format of hosts is as follows in the example

"hosts": [{"ip":"127.0.0.1","urls":"mysite.com
www.mysite.com"},{"ip":"127.0.0.3","urls":"test.com www.test.com"}]\
 If ip is 127.0.0.1, it can be omitted. So, the simplest way to add one
host name is as follows: "hosts": [{"urls":"mysite.com"}]

If you run this task, Windows will question you to provide the
administrative rights

The urls can be separated by either spaces, or commas or semicolons, and
prefixes starting with http:// or https:// or so on are removed
automatically by DvServer, also slashes in the host names with the rest
are removed automatically.

### Task: set DNS server list for your computer

This task is solved in windows in the registry
(HKEY\_LOCAL\_MACHINE\\SYSTEM\\ControlSet001(2)\\services\\Tcpip\\Parameters
ServerList)

DvServer can automatically change your DNS server list in the registry

If you in the list hosts (see above) add {"kind":"DNS","urls":"YOUR DNS
SERVER LIST"} they DNS server list on your computer will be changed when
you start DvServer

The example for combined 2 tasks is as follows in the example

"hosts": [{"ip":"127.0.0.1","urls":"mysite.com
www.mysite.com"},{"kind":"DNS","urls":"svc.cluster.local,my.company.net"}]

If you run this task, Windows will question you to provide the
administrative rights

The urls can be separated by either spaces, or commas or semicolons, and
prefixes starting with http:// or https:// or so on are removed
automatically by DvServer, also slashes in the host names with the rest
are removed automatically.

### Task: log more info

By default, only errors are recorded in the log, but you can tell to log
everything, including all server requests and replies with all headers
in special files

At first choose the logging level from the list as follows:

none (only startup information is logged, the log files do not grow at
all)\
 error (only errors are logged, this is the default logging level)\
 warning (only errors and warnings are logged)\
 info (errors, warnings and additional non-verbous info, for server
module, it includes the bodies of all requests and responses)\
 details (errors, warnings and any verbous info, for server module, it
includes the bodies and headers of all requests and responses)\
 debug (all info as in details plus specific debug info)\
 internal (all debug, but for internal use)\

In addition to debugging level, you should specify areas (one or many)
for which you want to receive the information, as follows:

config (issues related to the configuration)\
 json (issues related to the json processing )\
 hosts (issues related to the processing of hosts)\
 crud (requests and responses thru crud operations processed by this
server)\
 server (issues related to the forwarded server requests and responses)\

When you decide what level and what areas you want to log, you must
place them to the config as in the example below:

"logLevel": "details", \
 "logModules": "crud server"

The items in the list of modules can be separated by either spaces or
commas or semicolons

### Task: serve local files and/or forwarding requests to another server.

First, you must configure [the task to listen](#task_listen)

You have 2 entries in the config: server to be used for any hosts and
hostServers to be used for specific hosts specified in the request.
hostServers is an array of objects, each of which has the same format as
the server.

If you wish to configure this server disregarding the host names in the
requests, then set server

If you wish to run this server as a simple static server, you can
specify a field baseServer inside the server, providing either relative
or absolute paths. To specify relative path, you can use . (one dot) to
refer to the current directory and .. (two dots) to refer to the parent
directory. In the example below we use one dot (.) to have the current
folder as the root folder for the static server

"listen": ":80", \
 "server": {"baseFolder": "."}

To forward a request to another server, you should specify extraServer
(you can specify either baseFolder or extraServer or crud or any their
combination).

"listen": ":80", \
 "server": {"baseFolder": "/home/volodymyr/public","extraServer":
"http://localhost:3000"}, \

In the example above the search at first will be made in the local
folder /home/volodymyr/public, but if not found, it will be forwarded to
the server http://localhost:3000

If you wish to configure this server with regard to the host names in
the requests (in this case you can specify many root folders: one for
each host name), you should use the hostServers=[{},{},{},...]

Inside each {} as well as inside server you can specify the fields as
follows:

hosts (a list of host names comma(or space or semicolon)-separated for
which this block is related. This field is ignored if it is in server,
not in hostServers\
 baseFolder (a root folder to serve local files: should be empty if not
used)\
 rewrites (primary rewrite rules for url: [see rewrite
rules](#task_rewrite))\
 extraServer (server name to forward the request (starting with http://
or https://), should be empty if not used)\
 extraServerSettings (special settings for the extra Server [see
configure extra server settings](#task_extra_server_settings) )\
 serverRewrites (secondary rewrite rules for url applied before sending
to the extra server [see rewrite rules](#task_rewrite))\
 proxyName (the name on behalf of which server forwards requests to the
server specified in extraServer. This name will be used as the Referer
and Origin parameters in request headers)\
 accessControlAllowOrigin (the origin to be allowed in requests from the
browser (combined with extraServer origins if extraServer is used), for
example: \* - for all)\
 accessControlAllowMethod (for OPTIONS requests: the methods to be
allowed in request from the browser (combined with extraServer methods
if extraServer is used), for example: GET,POST,DELETE,PUT\
 accessControlAllowHeaders (for OPTIONS requests: the headers to be
allowed in request from the browser (combined with extraServer headers
if extraServer is used), for example: Authorization\
 accessControlExposeHeaders (the methods to be allowed in reply to the
browser (combined with extraServer headers if extraServer is used), for
example: Request-Id\
 accessControlMaxAge (for OPTIONS requests:the maximum time in seconds
for accessControl info to be kept in the browser (overwrites extraServer
MaxAge if non empty), for example: 3600 - 1 hour\
 accessControlAllowCredentials (=true or false. Indicates whether or not
the response to the request can be exposed when the credentials flag is
true)\
 crud (crud implementation, described later in this document)\
 crudEnabled (=true or false, false is default, enable the crud
functionality specified above in the crud)\
 cacheControl (Cache-Control header for this host, example 1: no-cache,
no-store, must-revalidate (completely no cache) example 2: public,
max-age=31536000 (maximum cache)\
 headersStatic Additional headers for static part of the server (not for
OPTIONS method) {"key1":"value1","key2":"value2",...}\
 headersExtraStatic Additional headers for the extra server (not for
OPTIONS method) {"key1":"value1","key2":"value2",...}\
 headersStaticOptions Additional headers for static part of the server
(for OPTIONS method only) {"key1":"value1","key2":"value2",...}\
 headersExtraStaticOptions Additional headers for the extra server (for
OPTIONS method only) {"key1":"value1","key2":"value2",...}\

"listen": ":80", \
 "hostServers": [{"hosts":"www.example.com
www.example1.com","baseFolder":"C:/Users/volodymyr/DvServer/accounts","extraServer":"http://localhost:8080"},
{"hosts":"api.example.com","extraServer":"http://api.mycompany.com","proxyName":
"www.mycompany.com","accessControlAllowOrigin":"www.example.com
www.example1.com"}], \
 "server": {"baseFolder": "."}

In the exampe above, if a requested host is either www.example.com or
www.example1.com, at first the search will be in the
C:/Users/volodymyr/DvServer/accounts as a root folder, but if not found,
it will be forwarded to http://localhost:8080. \
 If a requested host is api.example.com, the requests will be forwarded
to http://api.mycompany.com, and in the requests the Origin and
Reference will be replaced as if the requests are made from
www.mycompany.com and the browser will also receive the information,
that hosts names www.example.com www.example1.com are valid origins for
cross-domain requests.\
 If the host is not in the above (www.example.com www.example1.com
api.example.com), they will be served as static server requests to the
current folder (disregarding the host).

### Task: rewrite urls {#task_rewrite}

You can specify rewrites in rewrites or serverRewrites in within server
or within hostServers

Rewrites has the format as follows: "rewrites":[{"url":"url to be
replaced","src":"replacing source"},{"url":"url to be
replaced","src":"replacing source"},...]

url (it can be either exact url or url with the asterisk(\*) at the end.
The asterisk at the end means that the rest of this request is replaced
too\
 src (a new url, the source replacing the url)\

Example

"listen": ":80", \
 "hostServers":
[{"hosts":"www.example.com","serverRewrites":[{"url":"/login\*","src":"/"}],"baseFolder":"C:/Users/volodymyr/DvServer/accounts","extraServer":"http://localhost:8080"}],
\
 "server": {"baseFolder": ".", "rewrites":
[{"url":"/login\*","src":"/"},{"url":"/accounts/\*","src":"/"}]}

### Task: configure crud operations {#task_crud}

### Task: configure special settings for the extra server {#task_extra_server_settings}

You can configure special settings for the extraServer by providing the
extraServerSettings field in the server or hostServers

maxIdleConnections (integer number (not quoted), maxIdleConnections
controls the maximum number of idle (keep-alive) connections across all
hosts. Zero (default) means no limit.)\
 idleConnectionTimeout (number of seconds (not quoted), is the maximum
amount of time in seconds an idle (keep-alive) connection will remain
idle before closing itself. Zero (default) means no limit.)\
 disableCompression (true or false (default) (not quoted),
DisableCompression, if true, prevents the server from requesting
compression with an "Accept-Encoding: gzip" request header when the
Request contains no existing Accept-Encoding value. If the server
requests gzip on its own and gets a gzipped response, it's transparently
decoded. \
 disableKeepAlives (true or false (default)(not quoted),
disableKeepAlives, if true, prevents re-use of TCP connections between
different HTTP requests.\
 maxIdleConnectionsPerHost (MaxIdleConnsPerHost, if non-zero, controls
the maximum idle (keep-alive) connections to keep per-host. The default
is 2.)\
 responseHeaderTimeout (integer number of seconds (not quoted),
responseHeaderTimeout, if non-zero, specifies the amount of time in
seconds to wait for a server's response headers after fully writing the
request (including its body, if any). This time does not include the
time to read the response body. )\
 expectContinueTimeout (integer number of seconds (not quoted),
expectContinueTimeout, if non-zero, specifies the amount of time to wait
for a server's first response headers after fully writing the request
headers if the request has an "Expect: 100-continue" header. Zero means
no timeout and causes the body to be sent immediately, without waiting
for the server to approve. This time does not include the time to send
the request header.)\

Example

"listen": ":80", \
 "hostServers":
[{"hosts":"www.example.com","extraServer":"http://localhost:8080","extraServerSettings":{"responseHeaderTimeout":60}}],
\
 "server":
{"extraServer":"http://localhost:3000","extraServerSettings":{"disableCompression":true,"disableKeepAlives":true}}

### Folder management {#folder_management}

DvServer uses variables from the system environment as well as your
custom variables stored in DvServer.properties. Variables can be used in
any configuration files (DvServer.conf) as well as in templates. In
addition, DvServer uses the configuration file DvServer.conf.

At first, system environment variables are loaded and then file
DvServer.properties is looked for. The search for this file is made in
the current directory first. If not found and you set DV\_SERVER\_PATH
variable in your environment, with the name of the path, the search for
this file is made in that directory. If not found, the search is made in
your namespace directory. If you specify DV\_SERVER\_CURRENT\_NAMESPACE,
the directory name will be {user-folder}/DvServer/{namespace}. If you do
not provide DvServer.properties file, only environment variables are
used.

DvServer.properties file is optional. Its structure is as follows:

    #   comment
    key1=value1
    key2=value2
       ...

Character \\ is used as an escape character, so to place \\, you should
use \\\\. If you wish to specify = in the key, you should use \\=. If a
line starts with \#, it is considered as a comment. If a not comment
line does not have = or if the key is empty, it is considered as an
error. You can also you [\# directives(\#include \#if \#else \#endif
\#define...)](documents.html) in your DvServer.properties file.

If you specify DV\_SERVER\_CURRENT\_NAMESPACE in your
DvServer.properties file, it will override your namespace name for the
other use of the namespace.

Your configuration file DvServer.conf is also optional, but if you do
not provide it, the functionality will be very basic: just a static
server will run on the current folder and listening on port 80. So, to
use the full strength of DvServer, it is desirable to create a config,
the structure of this config is described above in this document. At
first, the search for DvServer.conf is made in the current folder. If
not found, in the folder {user-folder}/DvServer/{namespace}. If you
specify to create logs for the work of DvServer, they will be placed in
{user-folder}/DvServer/{namespace}/LOGS folder.

If you are not sure whether your config DvServer.conf is correct after
the application of all variables and \# directives, you can specify
DEBUG\_CONFIG\_SAVE\_FILENAME variable and the resulted config will be
written to this file name, if it is a valid file name.

If you would like to see global variables at the log, specify logging
level at least info, and logging area should include "config".
