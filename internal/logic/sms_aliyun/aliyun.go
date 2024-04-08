package sms_aliyun

import (
	"context"
	"errors"
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/base_model/base_enum"
	"github.com/kysion/base-library/utility/enum"
	"github.com/kysion/base-library/utility/json"
	"github.com/kysion/sms-library/sms_interface"
	"github.com/kysion/sms-library/sms_model"
	"github.com/kysion/sms-library/sms_model/sms_dao"
	"github.com/kysion/sms-library/sms_model/sms_enum"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

// SmsAliyun 阿里云短信平台

type sSmsAliyun struct {
	modules sms_interface.IModules
	dao     *sms_dao.XDao
}

func NewSmsAliyun(modules sms_interface.IModules) sms_interface.ISmsAliyun {
	return &sSmsAliyun{
		modules: modules,
		dao:     modules.Dao(),
	}
}

// VerifySms 校验验证码 (上下文、手机号、验证码、类型标识) (bool ,error) （不属于接口，属于调用者校验）
//func (t *SmsQyxs) VerifySms(ctx context.Context, mobile int64, captcha string, typeIdentifier ...string) (bool, error) {
//	return true, nil
//}

// SendCaptchaBySms 发送验证码

// SendSms 发送短信 (上下文、渠道商, 短信模版,请求内容)
func (s *sSmsAliyun) SendSms(ctx context.Context, provider sms_model.SmsServiceProviderConfig, template sms_model.SmsTemplateConfig, req sms_model.SmsSendMessageReq) (*sms_model.SmsResponse, error) {
	// 1、随机产生验证码
	// 	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	// 验证码序列化
	//bCode, err := json.Marshal(map[string]interface{}{
	//	"code": code,
	//})

	// 2、获取配置信息,创建短信服务客户端
	client, err := s.createClient(tea.String(provider.AccessKeyId), tea.String(provider.AccessKeySecret))
	if err != nil {
		return nil, err
	}

	// 处理所有手机号
	phones := ""
	for _, phone := range req.Phones {
		phones = phones + phone + ","
	}
	// 去除最后一个逗号
	phones = phones[:len(phones)-1]

	// 几个号码产生几个验证码
	splitPhones := strings.Split(phones, ",")
	req.Params = make([]string, len(splitPhones))

	for i, _ := range splitPhones {
		code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
		// 随机的六位数验证码
		req.Params[i] = code
	}

	r, _ := regexp.Compile("\\$\\{\\w+\\}")
	allString := r.FindAllString(template.TemplateContent, -1)
	var mapJson = make(map[string]string)
	for i, s := range allString {
		r1, _ := regexp.Compile("[a-z]")
		findAllString := r1.FindAllString(s, -1)
		key := ""
		for _, s := range findAllString {
			key = key + s
		}
		//  不同的手机相同的验证码
		mapJson[key] = req.Params[i]
	}
	marshal, _ := json.Marshal(mapJson)
	param := string(marshal)
	fmt.Println(param)

	// 3、准备请求结构数据
	sendReq := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  &phones,                           // 所有手机号码：181739642943,13243556309
		TemplateCode:  tea.String(template.TemplateCode), // 模版编号： SMS_246415163
		TemplateParam: &param,                            // 短信参数：$(code)
		SignName:      &template.SignName,                // 签名：菲菲kysion
	}

	// 4、发送并接受结果
	response, err := client.SendSms(sendReq)
	fmt.Println("返回结果：", response)

	var sendStatus []sms_model.SmsSendStatus
	var log []sms_model.SmsSendLog
	// 发送成功
	if *response.Body.Code == "OK" {
		// 遍历所有手机号
		for i, phone := range req.Phones {
			sendStatus = append(sendStatus, sms_model.SmsSendStatus{
				Phone:   phone,
				Code:    *response.Body.Code,
				Message: *response.Body.Message,
				Fee:     gconv.Uint64(len(req.Phones)),
			})

			// 记录发送日志
			log = append(log, sms_model.SmsSendLog{
				BusinessNo:  req.BusinessNo,
				Fee:         1,
				PhoneNumber: phone,
				Message:     *response.Body.Message,
				Code:        *response.Body.Code,
				Content:     gconv.String(sendReq),
				Remark:      "",
				UnionMainId: template.UnionMainId,
				Form:        -1,
				Type:        1,
				Status:      1,
				MetaData:    gconv.String(response),
				SignName:    template.SignName,
			})

			// 一个验证码支持多种业务场景的，那验证码类型就传入复合类型的进来，如：1登录 8找回密码/重置密码，
			captchaTypes := enum.GetTypes[int, base_enum.CaptchaType](req.CaptchaType, base_enum.Captcha.Type)
			cacheTimeLen := 5 * len(captchaTypes)

			for _, value := range captchaTypes {
				// 存储缓存：key = 业务场景 + 邮箱号   register_18170618733@163.com  login_18170618733@163.com
				cacheKey := value.Description() + "_" + phone

				// 方式1：保持验证码到缓存
				err = g.DB().GetCache().Set(ctx, cacheKey, req.Params[i], time.Minute*time.Duration(int64(cacheTimeLen)))
				// 方式2：保持验证码到缓存
				//_, err = g.Redis().Set(ctx, cacheKey, code)
				//if err == nil {
				//_, err = g.Redis().Do(ctx, "EXPIRE", cacheKey, time.Minute*time.Duration(int64(cacheTimeLen)))
				//}
				if err != nil {
					return nil, errors.New("验证码缓存失败")
				}
			}

			// TODO：如下方式只能支持单一的验证码校验业务场景
			// 存储缓存：key = 业务场景 + 手机号   register_18170618700  login_18170618700
			//smsType := base_enum.Captcha.Type.New(req.CaptchaType, "")
			//g.DB().GetCache().Set(ctx, smsType.Description()+"_"+phone, req.Params[i], time.Minute*5)
		}
	} else { // 失败
		// 遍历所有手机号
		for _, phone := range req.Phones {
			sendStatus = append(sendStatus, sms_model.SmsSendStatus{
				Phone:   phone,
				Code:    *response.Body.Code,
				Message: *response.Body.Message,
				Fee:     gconv.Uint64(len(req.Phones)),
			})

			// 记录发送日志
			log = append(log, sms_model.SmsSendLog{
				BusinessNo:  req.BusinessNo,
				Fee:         0,
				PhoneNumber: phone,
				Message:     *response.Body.Message,
				Code:        *response.Body.Code,
				Content:     gconv.String(sendReq),
				Remark:      "",
				UnionMainId: template.UnionMainId,
				Form:        -1,
				Type:        1,
				Status:      0,
				MetaData:    gconv.String(response),
				SignName:    template.SignName,
			})
		}

		// 发送失败换个渠道商进行发送，后续完善。。。。。。。。。

	}

	// 5、成功后统计短信数量
	{
		if *response.Body.Code == "OK" {
			// 根据签名找到appconfig
			smsAppConfig, _ := s.modules.SmsAppConfig().GetAppConfigByName(ctx, template.SignName)

			// 数据库扣除短信余量，appConfig
			_, err = s.modules.SmsAppConfig().UpdateAppNumber(ctx, smsAppConfig.Id, gconv.Uint64(len(req.Phones)))
		}
	}

	// 6、记录日志
	{
		for _, sendLog := range log {
			s.modules.SmsSendLogConfig().SaveSmsLog(ctx, &sendLog)
		}
	}

	// 注意：上诉就不添加事务了，因为短信发送后就算本地日志没写入成功，但是应用余量也还是需要修改的。
	return &sms_model.SmsResponse{
		SmsSendStatus: sendStatus,
		RequestId:     *response.Body.RequestId,
	}, err
}

// CreateClient 创建阿里云客户端
func (s *sSmsAliyun) createClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, err error) {
	config := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
	}

	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")

	// 创建客户端并返回
	client := &dysmsapi20170525.Client{}
	client, err = dysmsapi20170525.NewClient(config)
	return client, err
}

// ReceiveSms 接收短信 (上下文, 短信请求体) (bool,err)
func (s *sSmsAliyun) ReceiveSms(ctx context.Context, req sms_model.SmsReceiveSmsReq) (bool, error) {
	// 判断发送的短信业务类型
	switch req.Type {
	case sms_enum.Sms.Action.Send.Code(): // 普通发送

	case sms_enum.Sms.Action.Refund.Code(): // 退订

	case sms_enum.Sms.Action.Renewal.Code(): // 续费

	case sms_enum.Sms.Action.TopUp.Code(): // 充值

	default:
		return false, errors.New("暂不支持此业务类型")
	}

	return true, nil
}

// RegisterTemplate 添加短信模版
func (s *sSmsAliyun) RegisterTemplate(ctx context.Context, info *sms_model.SmsTemplateConfig) (*sms_model.SmsTemplateConfig, error) {
	info.ProviderNo = sms_enum.Sms.Type.Aliyun.Code() // aliyun

	result, err := s.modules.SmsTemplateConfig().CreateTemplate(ctx, info)

	return result, err
}

// RegisterSign 添加短信签名
func (s *sSmsAliyun) RegisterSign(ctx context.Context, signInfo *sms_model.SmsSignConfig) (*sms_model.SmsSignConfig, error) {
	// 编号 名称 从程序内部获取
	signInfo.ProviderNo = sms_enum.Sms.Type.Aliyun.Code()          // aliyun
	signInfo.ProviderName = sms_enum.Sms.Type.Aliyun.Description() // 阿里云

	sign, err := s.modules.SmsSignConfig().CreateSign(ctx, signInfo)
	return sign, err
}

// CreateProvider 添加渠道商
func (s *sSmsAliyun) CreateProvider(ctx context.Context, info *sms_model.SmsServiceProviderConfig) (*sms_model.SmsServiceProviderConfig, error) {
	info.ProviderNo = sms_enum.Sms.Type.Aliyun.Code()          // aliyun
	info.ProviderName = sms_enum.Sms.Type.Aliyun.Description() // 阿里云

	provider, err := s.modules.SmsServiceProviderConfig().CreateProvider(ctx, info)

	return provider, err
}
