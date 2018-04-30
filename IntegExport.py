# -*- coding: utf-8 -*-
"""
Spyder Editor

This is a temporary script file.
"""


import subprocess
import os

filename = 'C:\\Users\\sagsriva\\ima\\src\\main\\ima\\ascii_runtime_measurements'

print(os.getcwd())

f = subprocess.Popen(['tail','-F',filename],\
        stdout=subprocess.PIPE,stderr=subprocess.PIPE)
while True:
    line = f.stdout.readline()
    print(line)