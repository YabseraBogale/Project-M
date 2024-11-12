import sqlite3
import time

connection=sqlite3.connect("gold.db")
pointer=connection.cursor()
# SC TN SD
file = open("UT.csv")

j=0

for i in file.readlines()[1:]:
    try:
        statment="Insert Into UserData(Email,Firstname,Lastname,Country,HQCompanyName,Sent) values(?,'','','USA',?,'not_sent')"
        email=i.split(",")[3]
        companyname=i.split(",")[2]
        pointer.execute(statment,(email,companyname))
        connection.commit()
        j+=1
        if j%1000==0:
            print("At line ...",j)
            time.sleep(3)
            
    except Exception as e:
        print(str(e))

