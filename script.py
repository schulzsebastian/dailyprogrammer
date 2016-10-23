import re

with open('table.txt', 'r') as f:
    text = f.read()
    cells = re.findall(r'<t[dh]>(.*?)</t[dh]>', text)
    mj_cells = re.findall(r'<t[dh]>([MJ].*?)</t[dh]>', text)
    email_cells = re.findall(r'<td>([\w\.]+@[\w\.]+?)</td>', text)
    domain_cells = re.findall(r'<td>([\w.@]+example\.com)</td>', text)
