import sqlite3
from pprint import pprint
connection=sqlite3.Connection("freelance-data-v2.db")

pointer=connection.cursor()

pointer.execute("select * from Software")
result=pointer.fetchmany(10)
pprint(result)
