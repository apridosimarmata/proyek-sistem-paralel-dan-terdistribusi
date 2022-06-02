from flask import Flask, render_template, request, jsonify, redirect, url_for
from functools import wraps
import requests, json

app = Flask(__name__, template_folder='templates', static_folder='static', static_url_path='/static')


@app.route('/login')
def login():
    return render_template('login.html')

def access_token_required(f):
    @wraps(f)
    def check_if_access_token_exist(*args, **kwargs):
        if request.cookies.get('token') is None:
            return redirect(url_for('login'))
        else:
            authorize = requests.get('http://staff-service:8000/authorize', headers = request.headers)
            if authorize.status_code != 200:
                print(authorize.status_code)
                return redirect(url_for('login'))
        return f(*args, **kwargs)
    return check_if_access_token_exist

@app.route('/')
@access_token_required
def index():
    reservations = requests.get('http://staff-service:5004/list')
    return render_template('index.html', reservations =  json.loads(reservations.content))

if __name__ == '__main__':
    app.run(debug=True, host="0.0.0.0")
