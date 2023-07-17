@echo off

cls

:: Check for admin privileges
NET SESSION >nul 2>&1
if %errorLevel% == 0 (
    goto :admin
) else (
    echo Requesting admin privileges...
    powershell Start-Process -Verb runAs -FilePath "%0"
    exit
)

:admin

:: Set variables
set "exePath=C:\Anon\Anon.exe"
set "repo=https://github.com/NotGabry/anon"
set "iconPath=C:\Anon\icons\icon.png"

:: Create the folder
echo Creating 'C:\Anon' folder...
mkdir "C:\Anon"
mkdir "C:\Anon\icons"

:: Download the file
echo Downloading 'Anon.exe' from GitHub...
curl -L -s -o "%exePath%" "%repo%/releases/latest/download/Anon.exe"

:: Download the icon
echo Downloading icon from "%repo%/blob/main/icons/anon.png?raw=true"...
curl -L -s -o "%iconPath%" "%repo%/blob/main/icons/anon.png?raw=true"

:: Modify registry keys
echo Modifying registry keys...
reg add HKEY_CLASSES_ROOT\*\shell\Anon /ve /d "Upload to Anonfiles" /f >nul
reg add HKEY_CLASSES_ROOT\*\shell\Anon /v "icon" /d "%exePath%" /f >nul
reg add HKEY_CLASSES_ROOT\*\shell\Anon\command /ve /d "\"%exePath%\" \"%%1\"" /f >nul

echo Done!

pause
