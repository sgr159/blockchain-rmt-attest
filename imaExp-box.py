import subprocess
import sys
import json
import urllib
import urllib2
import time
hostname = sys.argv[1]
export_dest_ip = sys.argv[2]
export_dest_port = sys.argv[3]
skip_line = 0


while True:
    f = open("ascii_runtime_measurements_diff","r")
    line_num = 0
    for line in f.readlines():
        line_num+=1
        if skip_line != 0 and line_num <=skip_line:
            continue
        skip_line+=1
        print line
        linearr = str(line).split(" ")
        print(linearr)
        binName = str(linearr[-1]).split("/")[-1].strip('\n')
        binHash = linearr[3].split(":")[-1]
        verifyBinReq = {}
        verifyBinReq["Name"] = binName
        verifyBinReq["Measurement"] = binHash
        verifyBinReq["Base"] = 16
        hostReq = {}
        hostReq["Name"] = hostname
        hostReq["Data"] = [verifyBinReq]

        headers = {
                    'content-type': "application/json",
                }
        url = "http://"+export_dest_ip+":"+export_dest_port+"/verify"
        print (hostReq)
        
        req = urllib2.Request(url, json.dumps(hostReq))
        try:
            response = urllib2.urlopen(req)
            result = response.read()
            print(result)
        except:
            pass
    time.sleep(2)

