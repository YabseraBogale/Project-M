import psycopg2 as pg


conn=pg.connect(dbname="FriElVehicleHandover",user="yourusername",password="yourpassword")
pointer=conn.cursor()
statment="""
    INSERT INTO public.users(
	employeeid, firstname, lastname, email, phonenumber, department)
	VALUES (1, 'Yabsera', 'Bogale', 'yabserabogalercd30932013@gmail.com', '920201161', 'IT');
"""
pointer.execute(statment)
conn.commit()
