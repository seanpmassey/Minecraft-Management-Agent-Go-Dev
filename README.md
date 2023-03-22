# Minecraft-Management-Agent-Go-Dev
 
 To use, download the version for your Minecraft server and the config.yml file. Place both files together somewhere on your server. If you are running on Linux or MacOS, make sure to run chmod +x on the download so it becomes executable.

Edit the config.yml file to match your environment details. You either need to fill in the the server name, port, and password OR the path to the server.properties file. If no hostname is set in the config.yml file, it will assume that the server is localhost.

The following commands are available in this release:

users
-listusers (lists all users currently connected)

Example:
minecraft-agent-- users -listusers

server
-saveall (writes game data to disk)
-getdefaultgamemode (gets the current default game mode for all players)
-setdefaultgamemode (changes the default game mode for all players, accepts survival, creative, adventure, or spectator)
-setweather (used with the -weathertype flag to change the server's weather)
-weathertype (used with the -setweather flag to change the server's weather, accepts clear, rain or thunder as the options)

Examples:
minecraft-agent-- server -saveall
minecraft-agent-- server -setweather -weathertype thunder

ops - Note: these commands have not been tested
-addops (used with the -opsuser flag, grants users the default operator level permissions)
-removeops (used with the -opsuser flag, removes operator rights from a user)
-opsuser (used with the -addops and -removeops flags)

minecraft-agent-- -addops -opsuser 
