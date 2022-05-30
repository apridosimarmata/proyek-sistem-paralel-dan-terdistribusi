import requests, http
from flask import jsonify, request
from .constants import JSON_CONTENT

def post_data(url, data):
    try:
        with requests.post(url, json = data) as result:
            return result.content, result.status_code, JSON_CONTENT

    except Exception as e:
        return make_response(http.HTTPStatus.INTERNAL_SERVER_ERROR, extract_exception_message(e), None)

def get_data(url):
    try:
        with requests.get(url) as result:
            return result.content, result.status_code, JSON_CONTENT
    except Exception as e:
        return make_response(http.HTTPStatus.INTERNAL_SERVER_ERROR, extract_exception_message(e), None)

def patch_data(url, data):
    try:
        with requests.patch(url, json = data) as result:
            print(result.request)
            return result.content, result.status_code, JSON_CONTENT
    except Exception as e:
        return make_response(http.HTTPStatus.INTERNAL_SERVER_ERROR, extract_exception_message(e), None)

def remove_prefix(url, prefix):
    return url[len(prefix) + 1:]

def extract_exception_message(e):
    return f"Exception caused by {e.__class__.__name__}"

def make_response(status_code, message, result):
    meta = {
        'code' : status_code,
        'message' : message
    }

    return jsonify({
        'meta' : meta,
        'result' : result
    }), 200, {'Content-Type': 'application/json'}