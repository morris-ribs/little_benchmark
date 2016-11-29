#!flask/bin/python
from flask import Flask, jsonify, request

app = Flask(__name__)

@app.route('/')
def index():
    """Default endpoint"""
    return "Hello, World!"


@app.route('/bombhere', methods=['POST'])
def bombhere():
    """Handles the POST requests"""
    if not request.json:
        return 400
    dummyinfo = {
        'DummyString': request.json['DummyString'],
        'DummyInt': request.json['DummyInt'],
        'DummyFloat': request.json['DummyFloat'],
        'DummyDate': request.json['DummyDate']
    }
    return jsonify({'dummyinfo': dummyinfo}), 200


if __name__ == '__main__':
    app.run()
