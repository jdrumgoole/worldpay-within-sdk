# worldpay-within-sdk
Worldpay Within SDK to allow payments within IoT.

The core of this SDK is written in Go with a native Go interface. Along with the native Go interface is an RPC layer (Apache Thrift) to allow communication through other languages. It is intended that we will develop a number of complementary wrapper libraries for other languages which should include C#.NET, Java, Python at a minimum.

<h3>Install</h3>
<ol>
<li>Install Go command line</li>
<li>Set up the environmental variables correctly; you only need to set $GOPATH, and that should be set as <home>/<required_path>/<cloned_repo_structure>, where <home> is wherever you want the code, <required_path> is /src/innovation.worldpay.com</li>
<li>clone the repo to $GOPATH/src/innovation.worldpay.com</li>
<li>Get the dependencies; go get github.com/Sirupsen/logrus</li>
<li>Get the dependencies; go get github.com/gorilla/mux</li>
<li>Get the dependencies; go get github.com/nu7hatch/gouuid</li>
<li>Get the dependencies; go get git.apache.org/thrift.git/lib/go/thrift</li>
</ol>

<h3>Configuration file versus command line flags</h3>
<p>The RPC client takes command line flags e.g. -port 9091 but it can also take the flag -configfile 'conf.json' so you can specify the configuration in a config file. For example</p>

<code>
{<br/><br/>
&nbsp;&nbsp;&nbsp;&nbsp;"WorldpayWithinConfig": {<br/><br/>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"BufferSize" : 100,<br/><br/>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"Buffered": false,<br/><br/>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"Framed": false,<br/><br/>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"Host": "127.0.0.1",<br/><br/>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"Logfile": "worldpayWithin.log",<br/><br/>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"Loglevel": "warn",<br/><br/>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"Port": 9081,<br/><br/>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"Protocol": "binary",<br/><br/>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"Secure": false<br/><br/>
&nbsp;&nbsp;&nbsp;&nbsp;}<br/><br/>
}<br/><br/>
</code>


# Initial pre-alpha release - June 6, 2016

<ol>
<li>Core SDK somewhat complete but not 100%. No service handover (begin/end)</li>
<li>Thrift definition of SDK service and message types</li>
<li>Basic Java program demonstrating RPC function</li>
<li>RPC Agent tool to enable starting the RPC from command line and programmatically. All options exposed via CLI flags. use -h for usage.</li>
<li>C# namespace TBD (A.Brodie)</li>
<li>BUG: There is an issue with the int->price map in the Thrift services definition (Pointer/Value error in Go). This has been disabled for now.</li>
<li>Only binary transport works in Java/Go RPC example. Will investigate others.</li>
<li>Added a semi implemented console application (dev-client) which shows the usage of the SDK in Go. This is probably the best documentation for now :)</li>
</ol>

# Next steps
<ol>
<li>Document, document, document...</li>
<li>Will programatically add feedback Mustafa Kasmani, Andy Brodie to start discussion on features, security concerns etc</li>
<li>Did I already say documentation - need to convert Architecture document to HTML/XML based format. Also need to comment Go core and auto generate via GoDoc.</li>
<li>Andy Brodie will be kindly developing a C# wrapper library via the RPC interface</li>
<li>Conor H to convert the reference Java application into a wrapper library.</li>
</ol>
