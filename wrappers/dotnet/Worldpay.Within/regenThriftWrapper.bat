@echo off
setlocal

set _WPWithinHome=%GOPATH%\src\github.com\wptechinnovation\worldpay-within-sdk

rem Edit this to point to your Thrift installation
set _ThriftHome=c:\apps\thrift-0.9.3

set _OutputDirectory=%_WPWithinHome%\wrappers\dotnet\Worldpay.Within\Worldpay.Within.Rpc
echo Regenerating Thrift RPC classes in to %_OutputDirectory%

%_ThriftHome%\thrift-0.9.3.exe -r -out %_OutputDirectory% --gen csharp:nullable,union %_WPWithinHome%\rpc\wpwithin.thrift

endlocal