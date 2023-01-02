@echo off
set thisFilePath=%~dp0

cd /d %thisFilePath%
FOR /F "eol=# tokens=*" %%i IN (.env) DO SET %%i

pause
cmd /k