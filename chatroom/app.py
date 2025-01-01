from flask import Flask, render_template, request
from flask_socketio import SocketIO, join_room, leave_room, send

app = Flask(__name__)
app.config['SECRET_KEY'] = 'your-secret-key'
socketio = SocketIO(app)

# Keep track of active rooms
active_rooms = {}

@app.route('/')
def index():
    return render_template('index.html')

@app.route('/chat', methods=['POST'])
def chat():
    username = request.form.get('username')
    room = request.form.get('room')
    return render_template('chat.html', username=username, room=room)

@socketio.on('join')
def on_join(data):
    username = data['username']
    room = data['room']

    # Allow only two users in the room
    if room not in active_rooms:
        active_rooms[room] = []

    if len(active_rooms[room]) < 2:
        join_room(room)
        active_rooms[room].append(username)
        send(f"{username} has joined the room.", to=room)
    else:
        send("Room is full.", to=request.sid)

@socketio.on('message')
def handle_message(data):
    room = data['room']
    send(f"{data['username']}: {data['message']}", to=room)

@socketio.on('leave')
def on_leave(data):
    username = data['username']
    room = data['room']
    leave_room(room)
    if room in active_rooms:
        active_rooms[room].remove(username)
        send(f"{username} has left the room.", to=room)

if __name__ == '__main__':
    socketio.run(app, debug=True)
