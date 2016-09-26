package main

// nohup ./export.sh > export.log 2>&1 &

import (
	"code.google.com/p/mahonia"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

var logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)

//go run main.go 01 100 1

func main() {

	pp1 := os.Args[1]
	pp2 := os.Args[2]
	pp3 := os.Args[3]

	fileName1 := pp1 + "_01_" + pp2 + "_" + pp3 + ".txt"
	fileName2 := pp1 + "_02_" + pp2 + "_" + pp3 + ".txt"

	fmt.Println("\n++++++++++++++++++++++++++++++++++++++++++++++\n")

	fmt.Println("filename 1:", fileName1)
	fmt.Println("filename 2:", fileName2)

	fmt.Println("\n++++++++++++++++++++++++++++++++++++++++++++++\n")

	dstFile1, err := os.Create(fileName1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer dstFile1.Close()

	dstFile2, err := os.Create(fileName2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer dstFile2.Close()

	p1 := pp1                  //1  发票类型
	p2, _ := strconv.Atoi(pp2) //2  数量
	p3, _ := strconv.Atoi(pp3) //3  月份

	logger.Println("write data to file begin")
	for i := 1; i <= p2; i++ {
		fphm := fmt.Sprintf("%08d", i+p2*(p3-1))
		kpyf := fmt.Sprintf("2015%02d", p3)
		kprq := "2015-" + fmt.Sprintf("%02d", p3) + "-" + fmt.Sprintf("%02d", rand.Intn(27)+1) + " 00:00:00"
		tslsh := "0101" + kpyf + fmt.Sprintf("%08d", rand.Intn(199)+1)

		enc := mahonia.NewEncoder("GBK")

		// enc.ConvertString(r)

		// fmt.Println(fphm, kpyf, kprq, tslsh)

		if p1 == "01" {

			//01-01
			dstFile1.WriteString(enc.ConvertString(`3400151130~~`+fphm+`~~`+kprq+`~~77595858*32-//1-0619+7<906</-*9><9/<8538-**+<73-69/3264328>1693/50908-79184<11+0<906</25+1698>*/19+2*+<75+49~~上海铁路局合肥电务段~~340102X1035454X~~合肥市瑶海区凤阳路039电务段办公楼 0551-62121237~~工行合肥市长江东路支行 1302010229910000551~~上海铁路物资有限公司合肥分公司~~340102059704117~~安徽省合肥市新站区天水路3000号~~工行合肥长江东路支行 1302010209920017501~~28490~~4843.3~~33333.3~~上海铁路局合肥电务段  F1506CAOJ0003~~曹净~~李晓萍~~李玥~~~~0~~~~1~~N~~~~~~43048332293738858936~~`+kpyf+`~~2015-06-16 09:31:26~~134019600~~MIIBQwYJKoZIhvcNAQcCoIIBNDCCATACAQExCzAJBgUrDgMCGgUAMAsGCSqGSIb3DQEHATGCAQ8wggELAgEBMGkwXjELMAkGA1UEBhMCY24xFTATBgNVBAseDFb9W7Z6DlKhYDtcQDEdMBsGA1UEAx4Ueg5SoXU1W1CLwU5me6F0Bk4tX8MxGTAXBgNVBA0eEABjAGEAMQAwADAAMAAwADICBxgBAAABYEEwCQYFKw4DAhoFADANBgkqhkiG9w0BAQEFAASBgDAr6KJj4NNKNQRuGP1372EDpujLIdk9W2lHXdniux9drYB4kOywN6vRRXQ+FRjM9TPUrmgOIQvysLrWEesknJ63wVHvugy6c/erq72O0Kqbc2Q0TeoATuMCxALAGcQU6IAUYzTjOtu3qkrvrPkoLtgBGntCgxBSiBKc5UJEASvL~~N~~1~~~~~~~~99~~2015-06-17 00:13:40~~`+tslsh+`~~2015-06-17 03:00:17~~~~134000000~~134010000~~134019600~~134000000~~134010000~~134010200~~N~~-~~0~~0~~~~~~0~~0~~~~0~~~~~~0~~0~~~~0~~~~~~~~0~~~~~~0~~0~~01~~0`) + "\n")

			// 01-02
			dstFile2.WriteString(enc.ConvertString(`3400151130~~`+fphm+`~~1~~00~~对讲机电池~~TK31071800mA~~块~~20~~717~~14340~~.17~~2437.8~~`+kpyf+`~~2015-06-17 00:10:09~~`+tslsh+`~~N~~134000000~~134000000~~~~01`) + "\n")
			dstFile2.WriteString(enc.ConvertString(`3400151130~~`+fphm+`~~2~~00~~手持对讲机/电池~~TH-K4AT~~块~~20~~707.5~~14150~~.17~~2405.5~~`+kpyf+`~~2015-06-17 00:10:09~~`+tslsh+`~~N~~134000000~~134000000~~~~03`) + "\n")

			if i%10000 == 0 {
				logger.Println("current", i)
			}

		} else if p1 == "02" {

			// 02-01
			dstFile1.WriteString(enc.ConvertString(`3400124730~~`+fphm+`~~`+kpyf+`~~`+kprq+`~~105437~~~~03<3*+0/4/04>25-7989/139890822/643>22618-*46>44867+816->62786/</8/4317+06->+0-42791+00++9<517771*9>1/04/21-84886-78/9623*799570179/-086>83+93>21~~341222713902514~~太和县中原汽车运输有限公司~~33080077721258X~~浙江明旺乳业有限公司~~440301X18927007~~广州大旺食品有限公司深圳分公司~~33080077721258X~~浙江明旺乳业有限公司~~起运地：衢州  到达地：深圳~~74117.19~~.11~~8152.89~~499902905437~~82270.08~~皖K97833~~17.95~~货物名称：旺仔牛奶  货重：152.279  单价：324.43  旺仔牛奶  货重：121.681  单价：324.43~~运输方式：汽运~~134120200~~太和县国家税务局~~0~~巴赛赛~~巴赛赛~~巴赛赛~~~~~~N~~~~~~~~341222713902514~~太和县中原汽车运输有限公司~~~~99999999999~~网上抄报~~134120200~~太和县国家税务局~~2015-06-16 11:09:10~~9~~1~~MIIBQwYJKoZIhvcNAQcCoIIBNDCCATACAQExCzAJBgUrDgMCGgUAMAsGCSqGSIb3DQEHATGCAQ8wggELAgEBMGkwXjELMAkGA1UEBhMCY24xFTATBgNVBAseDFb9W7Z6DlKhYDtcQDEdMBsGA1UEAx4Ueg5SoXU1W1CLwU5me6F0Bk4tX8MxGTAXBgNVBA0eEABjAGEAMQAwADAAMAAwADICBxgBAAAC9C8wCQYFKw4DAhoFADANBgkqhkiG9w0BAQEFAASBgDUDZlIccZLD5eFpcw/i0YInFgu4BM7M+pPH1yj6lKHCjIEXXEPYBsSXYMagJsD6JNpnnnN5ufKpzxB/W6cZE4pmUWSbJpl0uHiIqIeRh1MBPb5jgCYy31//XiVMViLseKPRQMh6oc+4YuL20oDaJ/4cXR9oHO/IzJ2OdhZzp/Jc~~2015-06-16 11:09:10~~2~~N~~99999999999~~2015-06-16 11:09:10~~2015-06-17 00:05:05~~`+tslsh+`~~2015-06-17 03:00:13~~~~134000000~~134120000~~134120200~~~~~~~~Y~~-~~0~~0~~~~0~~0~~~~~~~~~~0~~~~0~~~~0~~~~~~~~0~~~~~~0~~0~~01~~0`) + "\n")

			//02-02
			dstFile2.WriteString(enc.ConvertString(`3400124730~~`+fphm+`~~1~~运费~~42950.32~~`+kpyf+`~~2015-06-17 00:05:06~~`) + tslsh + `~~Y~~134000000~~~~~~01` + "\n")
			dstFile2.WriteString(enc.ConvertString(`3400124730~~`+fphm+`~~2~~运费~~31166.87~~`+kpyf+`~~2015-06-17 00:05:06~~`) + tslsh + `~~Y~~134000000~~~~~~01` + "\n")

			if i%10000 == 0 {
				logger.Println("current", i)
			}

		} else if p1 == "03" {

			// 03-01
			dstFile1.WriteString(enc.ConvertString(kpyf+`~~134161422242~~`+fphm+`~~661508445447~~`+kprq+`~~000000~~53<>44712<03+924><1*75><1570>1452<-9*+571+/560-*0/<>3*4810<6>8817723>1-2-444/<9>44/+1/62>1452<-9*571+/560-*053<>4+4712<03+924><0/+01*55-*509++</-<97764<56>4706+011210>690/7-/5>7152+9>48175-/~~杨莉~~~~34122419651004132X~~安徽瑞和汽车销售服务有限公司~~341622683639881~~３０７线汽车产业园区~~０５５８－７６０８０００~~农业银行蒙城县支行城南分理处~~675101040002206~~北京现代牌/BH7160QAY~~轿车~~WAJ290006589231~~~~北京~~*~~FB309456~~LBEGCBFC4FX050691~~119658.12~~.17~~20341.88~~140000~~341622683639881~~安徽瑞和汽车销售服务有限公司~~黄金~~1~~~~5~~1~~3~~134160300~~蒙城县国家税务局~~~~~~Y~~~~~~~~13416030099~~鹿少杰~~13416033200~~蒙城县国家税务局纳税服务股~~2015-06-16 15:24:10~~~~000000000000000~~MIIBQwYJKoZIhvcNAQcCoIIBNDCCATACAQExCzAJBgUrDgMCGgUAMAsGCSqGSIb3DQEHATGCAQ8wggELAgEBMGkwXjELMAkGA1UEBhMCY24xFTATBgNVBAseDFb9W7Z6DlKhYDtcQDEdMBsGA1UEAx4Ueg5SoXU1W1CLwU5me6F0Bk4tX8MxGTAXBgNVBA0eEABjAGEAMQAwADAAMAAwADICBxgBAAACCD0wCQYFKw4DAhoFADANBgkqhkiG9w0BAQEFAASBgKGaFcKEokeoQ+UUWwWPAKpVJdrL1s2drfTE6MWWYCRtu6KSATxKVkS41rOWlO9IEKrROVCFR2hOETGcSwUCyZYbGGxpA/ZutBwiDmhgKgLbF6pwqBWpZ9BHkREAVDdz9GGlFTLCcvBIcRvMTI5lfVNZKk4A/qaemJU8X37+AxkU~~2015-06-16 15:24:10~~1~~~~N~~13416030099~~2015-06-16 15:24:10~~2015-06-17 00:05:05~~`+tslsh+`~~2015-06-17 03:00:08~~~~134000000~~134160000~~134160300~~~~~~~~Y~~-~~0~~0~~~~~~0~~0~~~~~~~~~~0~~0~~~~0~~~~~~~~0~~~~~~0~~0~~03~~0`) + "\n")

			if i%10000 == 0 {
				logger.Println("current", i)
			}

		} else if p1 == "04" {

			//04-01
			dstFile1.WriteString(enc.ConvertString(`3400151320~~`+fphm+`~~`+kprq+`~~*5>-7-783+109+>4<42+88146995*97+</814453578/82-87>60>03+*<42*-0<70693717+6<66280146995507+</814453578/82/9>>~~客户~~000000000000000~~~~~~合肥烁星商贸有限公司~~340103793551098~~合肥市杏花村镇刘冲村何大郢1幢 0551-5186299~~交通银行合肥分行三孝口支行 341302000018170081102~~5188.03~~881.97~~6070~~~~~~~~沈华伟~~~~~~~~1~~N~~~~~~65237239593291174147~~`+kpyf+`~~2015-06-16 16:37:21~~~~MIIBQwYJKoZIhvcNAQcCoIIBNDCCATACAQExCzAJBgUrDgMCGgUAMAsGCSqGSIb3DQEHATGCAQ8wggELAgEBMGkwXjELMAkGA1UEBhMCY24xFTATBgNVBAseDFb9W7Z6DlKhYDtcQDEdMBsGA1UEAx4Ueg5SoXU1W1CLwU5me6F0Bk4tX8MxGTAXBgNVBA0eEABjAGEAMQAwADAAMAAwADICBxgBAAAA1ocwCQYFKw4DAhoFADANBgkqhkiG9w0BAQEFAASBgGbX548h884w/Kd/gkzlebNosvVyyLem9XclO7Jh2xLtUcKKkCuU2L+eClHRQ8qeTj8w7kF8irlwqE3LyRITlj+oBFBs39d+QlZ86Q7qkFcpjWhD6Z4TqHPH8HXhgm71ePOa1HcwyX7+JYNlIRtMnYmcIRxQ9xkK9Z/dIhiKbGaV~~N~~1~~~~UNIDEAL~~~~~~P1~~~~~~2015-06-17 00:25:26~~`+tslsh+`~~2015-06-17 03:00:02~~134000000~~134010000~~134010300~~~~~~~~01~~0`) + "\n")

			//04-02
			dstFile2.WriteString(enc.ConvertString(`3400151320~~`+fphm+`~~1~~00~~格力空调KFR-26583FNEa-A3~~~~套~~1~~2471.79487179487~~2471.79~~.17~~420.21~~`+kpyf+`~~2015-06-17 00:12:02~~`+tslsh+`~~2015-06-17 03:00:02~~134000000~~~~~~~~~~~~01`) + "\n")
			dstFile2.WriteString(enc.ConvertString(`3400151320~~`+fphm+`~~2~~00~~格力空调KFR-26583FNBa-A3~~~~套~~1~~2716.23931623932~~2716.24~~.17~~461.76~~`+kpyf+`~~2015-06-17 00:12:02~~`+tslsh+`~~2015-06-17 03:00:02~~134000000~~~~~~~~~~~~01`) + "\n")

			if i%10000 == 0 {
				logger.Println("current", i)
			}

		} else if p1 == "10" {

			// 10-01
			dstFile1.WriteString(enc.ConvertString(tslsh+`~~499099900783~~111001571071~~`+fphm+`~~~~0~~034785344>9994725368++751/6+7117>/>/+49*<3<+749/5>618-/*0<4785344>999472531*258<*725/54*4828-*01>76019<<70+/8>*7~~110192585816506~~江苏圆周电子商务有限公司北京分公司~~北京市北京经济技术开发区荣华中路7号院3号楼十层1015 62648622~~~~000000000000000~~个人~~~~~~`+kprq+`~~154551~~~~~~75.39~~9.81~~85.2~~机器编号:499099900783~~京东商城~~京东商城~~~~~~16743594999278052962~~~~~~MIIBQgYJKoZIhvcNAQcCoIIBMzCCAS8CAQExCzAJBgUrDgMCGgUAMAsGCSqGSIb3DQEHATGCAQ4wggEKAgEBMGgwXjELMAkGA1UEBhMCY24xFTATBgNVBAseDFb9W7Z6DlKhYDtcQDEdMBsGA1UEAx4Ueg5SoXU1W1CLwU5me6F0Bk4tX8MxGTAXBgNVBA0eEABjAGEAMQAwADAAMAAwADICBgEAAAAYqzAJBgUrDgMCGgUAMA0GCSqGSIb3DQEBAQUABIGAjuJr8AHsMuBxgqvBaUyEfr7pdnr9O1l6TU6N8AIzFABnV02VDb7RZz7b29Hx5cg8wdhtoDbPNL1f3KTiIRLEDBdZ5LJqVkROQwJ4LKF70x2j4K0zOsCJ8hlJ0BSJy2HJ5F2Pi0DIe14pGCvpxtvXaUubd4oYMsw8V3Mq+oDZ4oI=~~0~~1~~9~~99999999999~~2015-08-09 23:21:09~~2~~N~~2015-08-10 02:00:10~~`+kpyf+`~~2015-08-10 03:00:51~~111000000~~111019200~~111019200~~~~~~~~01~~0`) + "\n")

			// 10-02
			dstFile2.WriteString(enc.ConvertString(tslsh+`~~111001571071~~`+fphm+`~~2~~百年孤独~~无~~~~1.00000000~~26.19469000~~26.19~~.13~~3.41~~2015-08-10 02:00:13~~`+kpyf+`~~2015-08-10 03:00:51~~111000000~~~~~~~~~~~~01`) + "\n")

			if i%10000 == 0 {
				logger.Println("current", i)
			}

		}

	}

	logger.Println("current", p2)

	logger.Println("write data to file end")

}
