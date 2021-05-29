import random
import os
txt = "LQ BF SI DZ LB PK FL QT QC DM QU IN HR PB BE AC QU FD YY QS DZ PT FC XX DO UT OP QD RR OY XH QB EH AE ZU YT CA FQ MQ KH IS SI CU ZB DU XN PL TS DQ KC WM UU EB OC YP PH NT OT PK XN PL QH QB PH OV EB MH BU PT QC OC ZD TU KC XN TG KP HO HR FC EH QC MV DV RX MT SI FA HT OY ER US AB RR XI BQ TW EL BQ QI TM OR FC YD PT UB OP RO HT TH MH CL DY IK MH"


ar = []
let = ['a','b','c','d','e','f','g','h','i','j','k','l','m','n','o','p','q','r','s','t','s','v','w','x','y','z']
tsplit = txt.split(" ")

done = []

def ar_to_txt(tmp):
    lst = ""
    for t in tmp:
        lst += t
    return lst

def process(letters):
    tmp = []
    tmp_txt = txt
    c = 0
    for t in ar:
        tmp.append([t,letters[c]])
        c+=1
        if c >= 25:
            break
    for q in tmp:
        tmp_txt = tmp_txt.replace(q[0],q[1])
    #print(tmp_txt)
    tmp_txt = tmp_txt.replace(" ","")
    return tmp_txt



for t in tsplit:
    if t not in ar:
        ar.append(t)

print("count: " + str(len(ar)))

process(let)

checking = []
import subprocess as subp


with open('harry.txt') as f:
    for line in f:
       # For Python3, use print(line)
        line = line.lower()
       
        if " " in line:
           s = line.split(" ")
           for i in s:
               i = i.strip()
               if i not in checking:
                   checking.append(i)
        else:
                line = line.strip()
                if line not in checking:
                       checking.append(line)
#for c in checking:
#    print(c)

def autoHash():
    countlines = 0
    total = 32
    ques = "?????????"
    for x in range(total):
        print(x)
        subp.check_call(str("C:\\hashcat-6.2.1\\hashcat.exe -m 1800 ass2.txt harry.txt --force " + ques), shell=True)
        ques += "?"
        #os.system('cmd /k "ping google.com"')
        
        #with open('C:\hashcat-6.2.1\hashcat.potfile') as f:
        #    for line in f:
            # For Python3, use print(line)
        #        countlines += 1


autoHash()

#print(ar_to_txt(let))
total = pow(26,26)
def start():
    for x in range(total):
        r = let
        random.shuffle(r)
        comp = ar_to_txt(r)
        if comp not in done:
            results = process(r)
            done.append(comp)
            if "flag" in results:
                print(results)
                
            #for res in checking:
            #    if res in results:
            #        print(results)
            #        break
            #print(str(x) + "/" + str(total) + " or " + str((x/total)*100))

#start()
        
