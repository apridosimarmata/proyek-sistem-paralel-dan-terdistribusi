from flask import Blueprint, jsonify, request, make_response
import requests, os, http
from utils import auth as authutils
from utils.request import *
from auth import authentication, authorization

prefix = 'auth'

auth = Blueprint(prefix, __name__)


@auth.route('/authentication/email', methods = ['POST'])
def authenticate_by_email():
    return authentication.by_email(request)

@auth.route('/authentication/phone', methods = ['POST'])
def authenticate_by_phone():
    return authentication.by_phone(request)

@auth.route('/authentication/phone/<whatsapp_no>', methods = ['GET'])
def request_phone_authentication(whatsapp_no):
    return authentication.request_code(request)

from utils.auth import access_token_required

@auth.route('/authorization', methods = ['GET'])
@access_token_required
def is_token_valid():
    return authorization.authorize(request)

@auth.route('/authorization/refresh')
def refresh_token():
    if request.headers.get('Access-Refresh-Token') is None:
        return make_response(http.HTTPStatus.UNAUTHORIZED, 'Access-Refresh-Token missing', None)
    return authorization.refresh_token(request.headers.get('Access-Refresh-Token'))
