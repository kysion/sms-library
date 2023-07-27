package sms_tencent

import (
	"context"
	"fmt"
	"github.com/kysion/sms-library/sms_interface"
	"github.com/kysion/sms-library/sms_model"
	"github.com/kysion/sms-library/sms_model/sms_dao"
	"github.com/kysion/sms-library/sms_model/sms_enum"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	"os"
)

// SmsTencent 腾讯云短信平台

type sSmsTencent struct {
	modules sms_interface.IModules
	dao     *sms_dao.XDao
}

func NewSmsTencent(modules sms_interface.IModules) sms_interface.ISmsTencent {
	return &sSmsTencent{
		modules: modules,
		dao:     modules.Dao(),
	}
}

// SendSms 发送短信 (上下文、渠道商, 短信模版,请求内容)
func (s *sSmsTencent) SendSms(ctx context.Context, provider sms_model.SmsServiceProviderConfig, template sms_model.SmsTemplateConfig, req sms_model.SmsSendMessageReq) (*sms_model.SmsResponse, error) {

	// 必要步骤：
	// 实例化一个认证对象，入参需要传入腾讯云账户密钥对 SecretId，SecretKey。
	// 硬编码密钥到代码中有可能随代码泄露而暴露，有安全隐患，并不推荐。
	// 为了保护密钥安全，建议将密钥设置在环境变量中或者配置文件中，请参考本文凭证管理章节。
	// credential := common.NewCredential("SecretId", "SecretKey")
	credential := common.NewCredential(
		os.Getenv(provider.AccessKeyId),
		os.Getenv(provider.AccessKeySecret),
	)

	// 非必要步骤
	// 实例化一个客户端配置对象，可以指定超时时间等配置
	cpf := profile.NewClientProfile()
	// SDK默认使用POST方法。
	// 如果你一定要使用GET方法，可以在这里设置。GET方法无法处理一些较大的请求。
	// 如非必要请不要修改默认设置。
	cpf.HttpProfile.ReqMethod = "POST"
	// SDK有默认的超时时间，如非必要请不要修改默认设置。
	// 如有需要请在代码中查阅以获取最新的默认值。
	cpf.HttpProfile.ReqTimeout = 30
	// SDK会自动指定域名。通常是不需要特地指定域名的，但是如果你访问的是金融区的服务，
	// 则必须手动指定域名，例如云服务器的上海金融区域名： cvm.ap-shanghai-fsi.tencentcloudapi.com
	cpf.HttpProfile.Endpoint = "cvm.tencentcloudapi.com"
	// SDK默认用TC3-HMAC-SHA256进行签名，它更安全但是会轻微降低性能。
	// 如非必要请不要修改默认设置。
	cpf.SignMethod = "TC3-HMAC-SHA256"
	// SDK 默认用 zh-CN 调用返回中文。此外还可以设置 en-US 返回全英文。
	// 但大部分产品或接口并不支持全英文的返回。
	// 如非必要请不要修改默认设置。
	cpf.Language = "en-US"
	//打印日志，默认是false
	// cpf.Debug = true

	// 实例化要请求产品(以cvm为例)的client对象
	// 第二个参数是地域信息，可以直接填写字符串ap-guangzhou，或者引用预设的常量
	client, _ := cvm.NewClient(credential, regions.Guangzhou, cpf)
	// 实例化一个请求对象，根据调用的接口和实际情况，可以进一步设置请求参数
	// 你可以直接查询SDK源码确定DescribeInstancesRequest有哪些属性可以设置，
	// 属性可能是基本类型，也可能引用了另一个数据结构。
	// 推荐使用IDE进行开发，可以方便的跳转查阅各个接口和数据结构的文档说明。
	request := cvm.NewDescribeInstancesRequest()

	// 基本类型的设置。
	// 此接口允许设置返回的实例数量。此处指定为只返回一个。
	// SDK采用的是指针风格指定参数，即使对于基本类型你也需要用指针来对参数赋值。
	// SDK提供对基本类型的指针引用封装函数
	request.Limit = common.Int64Ptr(1)

	// 数组类型的设置。
	// 此接口允许指定实例 ID 进行过滤，但是由于和接下来要演示的 Filter 参数冲突，先注释掉。
	// request.InstanceIds = common.StringPtrs([]string{"ins-r8hr2upy"})

	// 复杂对象的设置。
	// 在这个接口中，Filters是数组，数组的元素是复杂对象Filter，Filter的成员Values是string数组。
	request.Filters = []*cvm.Filter{
		&cvm.Filter{
			Name:   common.StringPtr("zone"),
			Values: common.StringPtrs([]string{"ap-guangzhou-1"}),
		},
	}

	// 通过client对象调用想要访问的接口，需要传入请求对象
	response, err := client.DescribeInstances(request)
	// 处理异常
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return nil, err
	}
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		panic(err)
	}

	// 打印返回的json字符串
	fmt.Printf("%s\n", response.ToJsonString())

	// 遍历结果数据
	// var statuses []sms_model.SmsSendStatus
	//for _, status := range response.Response.InstanceSet {
	//	statuses = append(statuses, sms_model.SmsSendStatus{
	//		Fee:     *status.Fee,
	//		Message: *status.Message,
	//		Code:    *status.Code,
	//		Phone:   *status.Placement,
	//	})
	//}
	resp := sms_model.SmsResponse{
		RequestId:     *response.Response.RequestId,
		SmsSendStatus: nil,
	}
	return &resp, err
}

// ReceiveSms 接收短信 (上下文, 短信请求体) (bool,err)
func (s *sSmsTencent) ReceiveSms(ctx context.Context, req sms_model.SmsReceiveSmsReq) (bool, error) {
	return true, nil
}

// RegisterTemplate 添加短信模版
func (s *sSmsTencent) RegisterTemplate(ctx context.Context, info *sms_model.SmsTemplateConfig) (*sms_model.SmsTemplateConfig, error) {
	info.ProviderNo = sms_enum.Sms.Type.Tencent.Code() // tencent

	result, err := s.modules.SmsTemplateConfig().CreateTemplate(ctx, info)

	return result, err
}

// RegisterSign 添加短信签名
func (s *sSmsTencent) RegisterSign(ctx context.Context, signInfo *sms_model.SmsSignConfig) (*sms_model.SmsSignConfig, error) {
	// 编号 名称 从程序内部获取
	signInfo.ProviderNo = sms_enum.Sms.Type.Tencent.Code()          // tencent
	signInfo.ProviderName = sms_enum.Sms.Type.Tencent.Description() // 腾讯云
	sign, err := s.modules.SmsSignConfig().CreateSign(ctx, signInfo)

	return sign, err
}

// CreateProvider 添加渠道商
func (s *sSmsTencent) CreateProvider(ctx context.Context, info *sms_model.SmsServiceProviderConfig) (*sms_model.SmsServiceProviderConfig, error) {
	info.ProviderNo = sms_enum.Sms.Type.Tencent.Code()          // tencent
	info.ProviderName = sms_enum.Sms.Type.Tencent.Description() // 腾讯云

	provider, err := s.modules.SmsServiceProviderConfig().CreateProvider(ctx, info)

	return provider, err
}

//// QueryProviderList 获取渠道商列表
//func (s *sSmsTencent) QueryProviderList(ctx context.Context, search *sys_model.SearchParams, isExport bool) (*sms_model.ServiceProviderConfigListRes, error) {
//	// 只获取腾讯云渠道商列表
//	search.Filter = append(search.Filter, sys_model.FilterInfo{
//		Field:       sms_dao.SmsServiceProviderConfig.Columns().ProviderNo,
//		Where:       "=",
//		Value:       sms_enum.Sms.Type.Tencent.Code(),
//		IsOrWhere:   true,
//		IsNullValue: false,
//	})
//
//	list, err := s.modules.SmsServiceProviderConfig().QueryProviderList(ctx, search, false)
//	if err != nil {
//		return nil, err
//	}
//
//	return list, nil
//}
