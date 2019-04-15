
package rbac

import (
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"rabbit/models/admin"
)

type NodeController struct {
	CommonController
}

func (this *NodeController) Rsp(status bool, str string) {
	this.Data["json"] = &map[string]interface{}{"status": status, "info": str}
	this.ServeJSON()
}

func (this *NodeController) Index() {
	if this.IsAjax() {
		groupid, _ := this.GetInt64("group_id")
		nodes, count := admin.GetNodelistByGroupid(groupid)

		for i := 0; i < len(nodes); i++ {
			nodes[i]["_parentId"] = nodes[i]["Pid"]
		}
		if len(nodes) < 1 {
			nodes = []orm.Params{}
		}
		// beego.Trace("%v", nodes)
		this.Data["json"] = &map[string]interface{}{"total": count, "rows": &nodes}
		this.ServeJSON()
		return
	} else {
		grouplist := admin.GroupList()
		b, _ := json.Marshal(grouplist)
		this.Data["grouplist"] = string(b)
		this.Layout = this.GetTemplate() + "/public/layout.html"
		this.TplName = this.GetTemplate() + "/rbac/node.html"
	}

}
func (this *NodeController) AddAndEdit() {
	n := admin.Node{}
	if err := this.ParseForm(&n); err != nil {
		//handle error
		this.Rsp(false, err.Error())
		return
	}
	var id int64
	var err error
	Nid, _ := this.GetInt64("Id")
	if Nid > 0 {
		id, err = admin.UpdateNode(&n)
	} else {
		group_id, _ := this.GetInt64("Group_id")
		group := new(admin.Group)
		group.Id = group_id
		n.Group = group
		if n.Pid != 0 {
			n1, _ := admin.ReadNode(n.Pid)
			n.Level = n1.Level + 1
		} else {
			n.Level = 1
		}
		id, err = admin.AddNode(&n)
	}
	if err == nil && id > 0 {
		this.Rsp(true, "Success")
		return
	} else {
		this.Rsp(false, err.Error())
		return
	}

}

func (this *NodeController) DelNode() {
	Id, _ := this.GetInt64("Id")
	status, err := admin.DelNodeById(Id)
	if err == nil && status > 0 {
		this.Rsp(true, "Success")
		return
	} else {
		this.Rsp(false, err.Error())
		return
	}
}
