import requests, jwt, json
from flask import jsonify
from routes.user import url as user_url
from utils import request as req
from .models import AuthenticationResponse


def create_authentication_response(user):
    resp = AuthenticationResponse(user)
    return resp.as_dict()

def by_email(request):
    result = req.post_data(user_url + '/verification/password', request.json)
    if json.loads(result[0]).get('meta').get('code') == 200:
        return jsonify({
            'meta' : {
                'code' : 200,
                'message' : 'Berhasil masuk',
            },
            'result' : create_authentication_response(json.loads(result[0]).get('result'))
        })
    return result

def by_phone(request):
    result = req.post_data(user_url + '/verification/code', request.json)
    if json.loads(result[0]).get('meta').get('code') == 200:
        return jsonify({
            'meta' : {
                'code' : 200,
                'message' : 'Berhasil masuk',
            },
            'result' : create_authentication_response(json.loads(result[0]).get('result'))
        })
    return result

def request_code(request):
    return req.get_data(user_url + '/verification/code/' + request.view_args['whatsapp_no'])
