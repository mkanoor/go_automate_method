package main
import (
	"fmt"
	"log"
  "github.com/user/manageiq/utils"
  "flag"
)

func parseArgs(params *utils.ConnectionParameters_t) {
  guidPtr  := flag.String("guid", "", "Automate Workspace GUID")
  urlPtr  := flag.String("url", "http://localhost:4000/api/", "Automate API URL")
  tokenPtr := flag.String("token", "", "Automate Token")
  userPtr := flag.String("username", "admin", "User")
  passwordPtr := flag.String("password", "smartvm", "Password")

  flag.Parse()
  if len(*guidPtr) == 0 {
    log.Fatal("GUID is a required Parameter")
  }

  params.Username =  *userPtr
  params.Password = *passwordPtr
  params.BaseUrl = *urlPtr
  params.MIQToken = *tokenPtr
  params.GUID     = *guidPtr
}

func updateObjects(workspace *utils.Workspace) {
  obj_names := workspace.GetObjectList()
  last_obj_key := obj_names[len(obj_names)-1]

  my_hash := make(map[string]int)
  my_hash["test"] = 2
  my_hash["prod"] = 5

  obj, _ := workspace.GetObject("root")

  obj.SetAttribute("via_rest_ipaddress", "1.1.1.94")
  obj.SetAttribute("via_rest_port", 3460)
  obj.SetAttribute("via_rest_ae_retry", true)
  obj.SetAttribute("via_rest_reboot", false)
  obj.SetAttribute("via_rest_hash", my_hash)

  obj, _ = workspace.GetObject(last_obj_key)
  obj.SetAttribute("via_rest_name", "Fordo")
  obj.SetAttribute("via_rest_age", 34)

  workspace.SetStateVar("my_ids", [6]int{2, 3, 5, 7, 11, 13})
  workspace.SetStateVar("my_name", "Fred Flintstone")
  workspace.SetStateVar("my_friend", "Barney Flintstone")
  workspace.SetStateVar("my_town", "Bedrock")

  cobj, _ := workspace.GetCurrentObject()
  cobj.SetAttribute("via_rest_in_current_obj", "45")

  fmt.Println("StateVar my_ids", workspace.GetStateVar("my_ids"))
}

func main() {
  var params utils.ConnectionParameters_t
  parseArgs(&params)
  workspace  := utils.NewWorkspace(&params)

  err := workspace.Fetch()
  if err != nil {
		log.Fatal(err)
  }
  // workspace.DumpObject("root")
  root,_ := workspace.GetObject("root")
  vm := root.GetAttribute("vm").(*utils.VMDB_Object)
  vm.AddCustomAttribute(root.GetAttribute("cname").(string), root.GetAttribute("cvalue").(string))


  updateObjects(workspace)

  workspace.Update()

  vmdb := utils.NewVMDB_Object(&params, "vms/25")
  // vmdb.CustomAttributes(true)
  fmt.Println("Got Handle")
	if err := vmdb.Fetch(); err != nil {
    fmt.Println("Error Fetching VMDB Object", err)
		log.Fatal(err)
	}
  // fmt.Println("Dumping VMDB Object")
  // vmdb.Dump()
  // vmdb.DeleteCustomAttribute("Freddy")
}

