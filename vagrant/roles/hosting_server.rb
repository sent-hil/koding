name "hosting_server"
description "The  role for hosting (cloudlinux) servers"
run_list ["recipe[nodejs]","recipe[golang]","recipe[git]","recipe[authconfig]","recipe[hosting]"]
