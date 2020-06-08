REM https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/echo#:~:text=If%20used%20in%20a%20batch,the%20beginning%20of%20the%20file. To prevent all commands in a batch file (including the echo off command) from displaying on the screen, on the first line of the batch file type^
@echo off

REM Output to help validate behaviour.^
echo "Hello World"

REM Delay to help humans validate behaviour.^
timeout /t 3