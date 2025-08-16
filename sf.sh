#!/bin/bash

cd backend

tree

echo -e "Server Is Running ON: \e[4m\e[36m http://localhost:3000; \e[0m"

go run server.go



#    \e[0;30m Black
#     \e[0;31m Red
#     \e[0;32m Green
#     \e[0;33m Yellow
#     \e[0;34m Blue
#     \e[0;35m Purple
#     \e[0;36m Cyan
#     \e[0;37m White 

# Background Colors:

#     \e[40m Black
#     \e[41m Red
#     \e[42m Green
#     \e[43m Yellow
#     \e[44m Blue
#     \e[45m Purple
#     \e[46m Cyan
#     \e[47m White 

# Other Styles:

#     \e[1m Bold
#     \e[4m Underline
#     \e[0m Reset (resets all attribut