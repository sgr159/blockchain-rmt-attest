import subprocess
import sys
import requests
import json

f = subprocess.Popen(['tail -F -n +1 ascii_runtime_measurements'],\
                stdout=subprocess.PIPE,stderr=subprocess.PIPE,shell=True)

hostname = sys.argv[1]
export_dest_ip = sys.argv[2]
export_dest_port = sys.argv[3]



while True:
        line = f.stdout.readline()
        linearr = str(line).split(" ")
        binName = str(linearr[-1]).split("/")[-1].strip('\n')[:-3]
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
        r = requests.post("http://"+export_dest_ip+":"+export_dest_port+"/verify",data=json.dumps(hostReq))
        print (hostReq)
        print(r.status_code, r.reason)


