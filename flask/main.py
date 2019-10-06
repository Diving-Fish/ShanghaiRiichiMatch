from flask import Flask
from flask import request
from selenium import webdriver
from time import sleep
from selenium.webdriver.chrome.options import Options
from selenium.common.exceptions import NoSuchElementException


class G:
    lock = False


app = Flask(__name__)
chrome_options = Options()
# chrome_options.add_argument('--headless')
chrome_options.add_argument('--no-sandbox')
driver = webdriver.Chrome(chrome_options=chrome_options)
driver.get("https://majsoul.union-game.com/dhs/")
driver.find_element_by_id('username').send_keys('Diver_Fish')
driver.find_element_by_id('password').send_keys('Dfyshisb123')
sleep(0.5)
driver.find_element_by_xpath('//*[@id="root"]/div/div[2]/div/div/div[2]/div/form/div[3]/button').click()
sleep(2)
driver.find_element_by_xpath('//*[@id="root"]/div/div[1]/main/div[2]/ul/li[2]/div/div[5]/a/button').click()


def _get_name(user_id):
    driver.find_element_by_xpath('//*[@id="root"]/div/header/div/div[3]/div/div/div/div/button[1]').click()
    sleep(0.2)
    driver.find_element_by_xpath('//*[@id="root"]/div/header/div/div[3]/div/div/div/div/button[2]').click()
    driver.find_element_by_xpath('//*[@id="root"]/div/div[1]/main/div[2]/div/div/button[2]').click()
    driver.find_element_by_xpath('/html/body/div[2]/div[2]/div/div[2]/div/div[1]/div/div/div/textarea[3]').send_keys(
        user_id)
    driver.find_element_by_xpath('/html/body/div[2]/div[2]/div/div[3]/button[1]').click()
    sleep(0.2)
    st = driver.find_element_by_xpath('/html/body/div[2]/div[2]/div/div[2]/div/div[2]/div/div/div/textarea[3]').text
    sleep(0.2)
    print(st)
    if st == '':
        driver.find_element_by_xpath('/html/body/div[2]/div[2]/div/div[3]/button[2]').click()
        return '获取角色信息出错'
    else:
        driver.find_element_by_xpath('/html/body/div[2]/div[2]/div/div[3]/button[3]').click()
        return st


def _select(data):
    index = 1
    ready = 0
    for ob in data:
        name = ob['name']
        point = ob['point']
        while True:
            try:
                node = driver.find_element_by_xpath(
                    '//*[@id="root"]/div/div[1]/main/div[2]/div/div[1]/div[1]/div[2]/ul/li[%d]' % index)
                if node.text == name:
                    print("%s %d" % (name, point))
                    driver.find_element_by_xpath('//*[@id="root"]/div/div[1]/main/div[2]/div/div[1]/div[1]/div[2]/ul/li[%d]/div[2]/button' % index).click()
                    ready += 1
                    sleep(0.3)
                    if driver.find_element_by_xpath('//*[@id="root"]/div/div[1]/main/div[2]/div/div[1]/div[2]/div[1]/div[%d]/div[2]/div[1]/p' % ready).text != name:
                        return -1
                    driver.find_element_by_xpath('//*[@id="root"]/div/div[1]/main/div[2]/div/div[1]/div[2]/div[1]/div[%d]/div[2]/div[2]/div/input' % ready).send_keys(str(point))
                    index = 1
                    break
                index += 1
            except NoSuchElementException:
                break
    return ready


def _start():
    driver.find_element_by_xpath('//*[@id="root"]/div/div[1]/main/div[2]/div/div[1]/div[2]/div[2]/label[1]/span[1]/span[1]/input').click()
    driver.find_element_by_xpath('//*[@id="root"]/div/div[1]/main/div[2]/div/div[1]/div[2]/div[3]/button[3]').click()


@app.route('/get_username/<user_id>')
def get_username(user_id):
    if G.lock:
        return "busy"
    G.lock = True
    name = _get_name(user_id)
    G.lock = False
    return name


@app.route('/start_match', methods=['POST'])
def start_match():
    if G.lock:
        return "busy"
    G.lock = True
    driver.find_element_by_xpath('//*[@id="root"]/div/header/div/div[3]/div/div/div/div/button[1]').click()
    sleep(0.2)
    driver.find_element_by_xpath('//*[@id="root"]/div/header/div/div[3]/div/div/div/div/button[3]').click()
    sleep(0.2)
    j = request.get_json()
    ready = _select(j['data'])
    print(ready)
    if ready != 4:
        G.lock = False
        return "bad"
    _start()
    G.lock = False
    return "ok"


@app.route('/get_now_info')
def get_now_info():
    if G.lock:
        return "busy"
    G.lock = True
    driver.find_element_by_xpath('//*[@id="root"]/div/header/div/div[3]/div/div/div/div/button[3]').click()
    sleep(0.2)
    index = 1
    ready = []
    playing = []
    while True:
        try:
            ready.append(driver.find_element_by_xpath('//*[@id="root"]/div/div[1]/main/div[2]/div/div[1]/div[1]/div[2]/ul/li[%d]' % index).text)
            index += 1
        except NoSuchElementException:
            break
    index = 1
    while True:
        try:
            l = []
            for i in range(1, 4):
                l.append(driver.find_element_by_xpath('//*[@id="root"]/div/div[1]/main/div[2]/div/div[2]/div[%d]/div/ul/li[%d]/div/span' % (index, i)).text)
            playing.append(l)
            index += 1
        except NoSuchElementException:
            break
    G.lock = False
    return {"ready": ready, "playing": playing}
