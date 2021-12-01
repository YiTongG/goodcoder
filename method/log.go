package method

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"strings"
	"sync"
	"time"
)

var once sync.Once
var lastSendTime time.Time

//true的情况下会额外使用fmt.Printf打印一次
var IsDebugMode = false
var Env = "prod"
var defaultToken = constants.DEFAULT_DINGDING_ACCESS_TOKEN

// panic 打印的日志，
func Panic(format string, v ...interface{}) {
	commonLog.Fatalf(format, v...)
}

// 重大错误报警用
func Fatal(format string, v ...interface{}) {
	commonLog.Fatalf(format, v...)
}

// Legacy functions
// 历史遗留
func Crit(format string, v ...interface{}) {
	commonLog.Fatalf(format, v...)
}

func Debug(format string, v ...interface{}) {
	if Env == constants.ENV_DEVBOX {
		fmt.Printf(time.Now().Format(constants.DATE_TIME_FMT_3)+" "+format+"\n", v...)
	}
	commonLog.Debugf(format, v...)
}

func DebugSql(ctx string, t time.Time, sql string, v ...interface{}) {
	if IsDebugMode {
		sql = strings.ReplaceAll(sql, "%", "%%")
		if ctx != "" {
			sql = ctx + " SQL: " + sql
		} else {
			sql = "SQL: " + sql
		}
		if len(v) >= 0 && len(v) < 50 {
			for _, val := range v {
				if val == nil || reflect.TypeOf(val).Kind() == reflect.String {
					sql = strings.Replace(sql, "?", "'%v'", 1)
				} else {
					sql = strings.Replace(sql, "?", "%v", 1)
				}
			}
			sql = fmt.Sprintf(sql, v...)
		} else {
			firstQuestionMark := strings.Index(sql, "?")
			sql = sql[:utils.MinInt(firstQuestionMark, len(sql))]
		}

		duration := time.Now().Sub(t).Seconds()
		Debug("%s\nSQL execution time: %.4fs", sql, duration)
	}
}

func Err(format string, v ...interface{}) {
	if Env == constants.ENV_DEVBOX {
		fmt.Printf(time.Now().Format(constants.DATE_TIME_FMT_3)+" "+format+"\n", v...)
	}
	commonLog.Errorf(format, v...)
	if metrics.CommonMetrics.Counter != nil {
		metrics.CommonMetrics.Counter.Inc("log.error")
	}
}

func ErrWithMetrics(counter *commonMetrics.CounterVec, label, format string, v ...interface{}) {
	Err(format, v...)
	if counter != nil {
		counter.Inc(label)
	}
}

//打印错误日志并钉钉&短信报警（短信有频次限制）
func ErrWithNotification(format string, v ...interface{}) {
	AlertDingdingWithToken(defaultToken, format, v...)

	//2020-07-29 短信接口已下线，全部改为发钉钉
	//Err(format, v...)
	//if Env == constants.ENV_PROD {
	//	sendNotificationWithLimit(fmt.Sprintf(format, v...))
	//}
}

//打印多个错误日志并报警
func ErrsWithNotification(format string, errs []error) {
	for _, err := range errs {
		ErrWithNotification(format, err.Error())
	}
}

func Info(format string, v ...interface{}) {
	commonLog.Infof(format, v...)
	if Env == constants.ENV_DEVBOX {
		fmt.Printf(time.Now().Format(constants.DATE_TIME_FMT_3)+" "+format+"\n", v...)
	}
}

func Warning(format string, v ...interface{}) {
	commonLog.Warnf(format, v...)
	if metrics.CommonMetrics.Counter != nil {
		metrics.CommonMetrics.Counter.Inc("log.warning")
	}
	if Env == constants.ENV_DEVBOX {
		fmt.Printf(time.Now().Format(constants.DATE_TIME_FMT_3)+" "+format+"\n", v...)
	}
}

//打印错误日志并钉钉报警（默认报警群
func AlertWithDingding(format string, v ...interface{}) {
	AlertDingdingWithToken(defaultToken, format, v...)
}

// 打印错误日志并钉钉报警
func AlertDingdingWithToken(token, format string, v ...interface{}) {
	Err(format, v...)
	if Env == constants.ENV_PROD {
		SendDingdingGroupMessageWithToken(token, fmt.Sprintf(format, v...))
	}
}

//发送钉钉报警，不打印错误日志（默认报警群
func SendDingdingGroupMessage(data string, atMobiles ...string) {
	SendDingdingGroupMessageWithToken(defaultToken, data, atMobiles...)
}

//发送钉钉报警，不打印错误日志
func SendDingdingGroupMessageWithToken(token string, data string, atMobiles ...string) {
	err := dingding.HookDingding(token, data, atMobiles...)
	if err != nil {
		Err("mds_send_msg dingding failed: %s", err.Error())
	}
}

//发送钉钉&短信报警with限频（短信5分钟内只发送1次）
func sendNotificationWithLimit(data string) {
	SendDingdingGroupMessage(data)
	if lastSendTime.Add(smsInterval * time.Minute).Before(time.Now()) {
		lastSendTime = time.Now()
		SendSms(data)
	}
}

//发送钉钉&短信报警
func SendNotification(data string) {
	if Env == constants.ENV_PROD {
		SendDingdingGroupMessage(data)
		//SendSms(data)
	}
}

// 向钉钉国际化报警群发送报警消息
func SendDingdingInter(data string) {
	return
}

//短信报警
func SendSms(data string) {
	if data == "" {
		return
	}
	encodeString := base64.StdEncoding.EncodeToString([]byte(data))
	smsPostBody := make(url.Values)
	smsPostBody.Set("tousername", userNames)
	smsPostBody.Set("content", encodeString)

	_, err := SmsClient.PostForm(smsUrl, smsPostBody)
	if err != nil {
		commonLog.Errorf("mds_send_msg failed: %s", err.Error())
	}
}

func SendPhone(names []string, data string) {
	conf := config.Get().ThirdParty.Infra
	// commonLog.Debug("mds_send_phone config %+v", conf)
	if !conf.Enabled {
		return
	}
	if data == "" || len(names) == 0 {
		return
	}
	once.Do(func() {
		header := make(map[string][]string, 0)
		header["x-access-username"] = []string{conf.UserName}
		header["x-access-token"] = []string{conf.Token}
		PhoneClient.Header = header
	})
	body := map[string]interface{}{
		"types": []string{"phone"},
		"users": names,
		"content": data,
	}
	b, err := json.Marshal(body)
	if err != nil {
		commonLog.Errorf("mds_send_phone failed with Marshal: %+v", err)
		return
	}
	_, err = PhoneClient.Post(conf.WebHook.Url, "application/json", bytes.NewReader(b))
	if err != nil {
		commonLog.Errorf("mds_send_phone failed: %+v", err)
		return
	}
	commonLog.Infof("mds_send_phone TO %+v, content: [%+v] success!", names, data)
}

func Duration(t time.Time, ctx string) {
	Debug("%.4fs for %s", time.Now().Sub(t).Seconds(), ctx)
}

func SetDebugMode(is bool) {
	IsDebugMode = is
	Info("debug mode set to %v", is)
}

func SetEnv(env string) {
	Env = env
	Info("env set to %s", env)
}

func SetDingdingToken(token string) {
	if token != "" {
		defaultToken = token
	}
}