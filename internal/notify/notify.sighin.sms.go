package notify

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/sdk"
	"github.com/lishimeng/passport/internal/db/model"
	"github.com/lishimeng/passport/internal/etc"
)

// TODO 登录验证码(SMS)
func SighInSendSms(code string, to string) (resp sdk.Response, err error) {
	client := sdk.NewClient(etc.Config.Notify.Host)
	template, err := GetSighInSmsTemplate()
	if err != nil {
		log.Info(err)
		return
	}
	params := make(map[string]string)
	params["code"] = code
	var req = sdk.SmsRequest{
		Template:      template.Template,
		TemplateParam: params,
		Receiver:      to,
	}
	resp, err = client.SendSms(req)
	if err != nil {
		log.Info(err)
		return
	}
	return
}

func GetSighInSmsTemplate() (info model.Notify, err error) {
	var qs = app.GetOrm().Context.QueryTable(new(model.Notify))
	cond := orm.NewCondition()
	cond = cond.And("Status__exact", 1)
	cond = cond.And("Category__exact", model.SmsSighIn)
	err = qs.SetCond(cond).One(&info)
	return
}
