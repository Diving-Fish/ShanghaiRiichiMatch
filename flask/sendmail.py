import smtplib
from email.mime.text import MIMEText
from email.header import Header

mail_host = "smtp.163.com"
mail_user = "dfyshisb@163.com"
mail_pass = "Dfyshisb123"

f = open('source.txt', 'r', encoding='utf-8')
s = f.readlines()
for line in s:
    arr = line.replace('\n', '').split('\t')

    sender = 'dfyshisb@163.com'
    receivers = [arr[4]]

    message = MIMEText('您的账号信息如下：\n账号: %s\n密码: %s\n\n请迅速前往http://47.100.50.175:8080/ 修改您负责人账号的密码。\n祝好！' % (arr[1] + '000', arr[5]), 'plain', 'utf-8')
    message['From'] = "2019上海高校立直麻将个人赛组委会<dfyshisb@163.com>"
    message['To'] = arr[4]

    subject = '2019上海高校立直麻将个人赛负责人邮件'
    message['Subject'] = Header(subject, 'utf-8')

    smtpObj = smtplib.SMTP()
    smtpObj.connect(mail_host, 25)    # 25 为 SMTP 端口号
    smtpObj.login(mail_user, mail_pass)
    smtpObj.sendmail(sender, receivers, message.as_string())
    print("邮件发送成功")