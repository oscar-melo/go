package main

import (
    "general/users/rolesapi/controllers"
    "general/users/rolesapi/business"
    "general/users/rolesapi/config"
    
)


func main() {
    config.LoadConfig();

  /*  
    fmt.Println(config.Cfg.Server.Port)
	fmt.Println(config.Cfg.Database.Name)
*/

	business.InitData();
    controllers.HandleRequests()
}

