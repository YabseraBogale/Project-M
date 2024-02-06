from flask import Flask,render_template,request
from flask_socketio import SocketIO, emit


app=Flask(__name__)
app.config['SECRET_KEY'] = 'secret!'
socketio = SocketIO(app)


@app.route('/')
def main():
    return render_template('base.html')

@socketio.on('message')
def handle_message(data):
    print('received message: ' + data)

@socketio.on('json')
def handle_json(json):
    print('received json: ' + str(json))

if __name__=="__main__":
    socketio.run(app)