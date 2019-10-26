from flask import Flask
import requests
import demjson
from flask import request
from selenium import webdriver
from time import sleep
from selenium.webdriver.chrome.options import Options
from selenium.common.exceptions import NoSuchElementException
from selenium.common.exceptions import WebDriverException


def submit_score():
    driver.find_element_by_xpath('//*[@id="root"]/div/header/div/div[3]/div/div/div/div/button[4]').click()
    sleep(2)
    driver.find_element_by_xpath('//*[@id="root"]/div/div[1]/main/div[2]/div/div[2]/div/div/div/div[2]/div/div').click()
    sleep(1)
    driver.find_element_by_xpath('//*[@id="menu-"]/div[2]/ul/li[3]').click()
    sleep(1)
    f = open('snapshot.key', 'r', encoding='utf-8')
    snapshot = f.read()
    f.close()
    index = 1
    while True:
        try:
            path = '//*[@id="root"]/div/div[1]/main/div[2]/div/div[2]/div/table/tbody/tr[%d]' % index
            txt = driver.find_element_by_xpath(path).text
        except NoSuchElementException:
            break
        if index == 1:
            fw = open('snapshot.key', 'w', encoding='utf-8')
            fw.write(txt)
            fw.close()
        if snapshot == txt:
            break
        print("[Starter] Submitted: %s", txt)
        for i in [2, 3, 4, 5]:
            s2 = driver.find_element_by_xpath(path + "/td[%d]" % i).text
            arr = s2.split(' ')
            arr[1] = float(arr[1])
            resp = requests.post("http://localhost:8088/api/public/push_score", json={
                "round": 1,
                "name": arr[0],
                "point": arr[1]
            })
        index = index + 1
    driver.find_element_by_xpath('//*[@id="root"]/div/header/div/div[3]/div/div/div/div/button[1]').click()


def __start_match():
    # print("try to start...")
    resp = requests.get("http://localhost:5000/get_now_info")
    j = demjson.decode(resp.text, 'utf-8')
    ready = []
    for r in j["ready"]:
        resp2 = requests.get("http://localhost:8088/api/public/score", {"round": 1, "name": r})
        j2 = demjson.decode(resp2.text, 'utf-8')
        if j2["scores"] is None:
            j2["scores"] = []
        if j2["id"] == 0 or len(j2["scores"]) >= 6:
            continue
        ready.append(r)
    if len(ready) >= 4:
        requests.post("http://localhost:5000/start_match", None, {
            "data": [
                {"name": ready[0], "point": 25000},
                {"name": ready[1], "point": 25000},
                {"name": ready[2], "point": 25000},
                {"name": ready[3], "point": 25000}
            ]
        })
        print("[Starter] Started: %s %s %s %s" % (ready[0], ready[1], ready[2], ready[3]))
        return 1
    return 0


def start_match():
    while __start_match() == 1:
        pass


chrome_options = Options()
chrome_options.add_argument('--headless')
chrome_options.add_argument('--no-sandbox')
driver = webdriver.Chrome(chrome_options=chrome_options)
driver.get("https://majsoul.com/dhs/")
driver.set_window_size(1920, 1080)
driver.find_element_by_id('username').send_keys('Dennis_Frank')
driver.find_element_by_id('password').send_keys('Dfyshisb123')
sleep(0.5)
driver.find_element_by_xpath('//*[@id="root"]/div/div[2]/div/div/div[2]/div/form/div[3]/button').click()
sleep(2)
driver.find_element_by_xpath('//*[@id="root"]/div/div[1]/main/div[2]/ul/li/div/div[5]/a/button').click()
sleep(3)
driver.find_element_by_xpath('//*[@id="root"]/div/header/div/div[3]/div/div/div/div/button[4]').click()
print("started successfully!")
sleep(10)
while True:
    try:
        sleep(5)
        submit_score()
        sleep(2)
        __start_match()
    except WebDriverException:
        driver.close()
        exit(0)
