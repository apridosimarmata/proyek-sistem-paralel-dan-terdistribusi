import datetime
import json

import requests
from flask import Flask, jsonify, request
from flask_restful import abort
from flask_sqlalchemy import SQLAlchemy
from sqlalchemy_utils import UUIDType
import uuid

app = Flask(__name__)
app.config["SQLALCHEMY_DATABASE_URI"] = "sqlite:///test.db"
#app.config["SQLALCHEMY_DATABASE_URI"] = "mysql://root:proyekspt@localhost:33060/proyekspt"
app.config["SQLALCHEMY_TRACK_MODIFICATIONS"] = False
db = SQLAlchemy(app)


class Reservation(db.Model):
    __tablename__ = 'reservations'
    id = db.Column(UUIDType, primary_key=True)
    name = db.Column(db.String(100), nullable=False)
    phone_number = db.Column(db.String(100), nullable=False)
    email = db.Column(db.String(100), nullable=False)
    reservation_date = db.Column(db.DateTime, nullable=False)
    status = db.Column(db.String(100), nullable=False)
    created_at = db.Column(db.DateTime, nullable=False, default=datetime.datetime.now())
    created_by = db.Column(db.String(100), nullable=True)
    modified_at = db.Column(db.DateTime, nullable=False, default=datetime.datetime.now())
    modified_by = db.Column(db.String(100), nullable=True)

    def to_json(self):
        return {
            'id': self.id,
            'name': self.name,
            'phone_number': self.phone_number,
            'email': self.email,
            'reservation_date': self.reservation_date,
            'status': self.status
        }


@app.route("/list", methods=["GET"])
def get_reservations():
    reservations = Reservation.query.all()
    return jsonify([reservation.to_json() for reservation in reservations])


@app.route("/<reservation_id>", methods=["GET"])
def get_reservation(reservation_id):
    reservation = Reservation.query.get(reservation_id)
    if reservation is None:
        abort(404)
    return jsonify(reservation.to_json()), 200, {'Content-Type': 'application/json'}


@app.route("/create", methods=["POST"])
def create_reservation():
    req = requests.get("http://staff-service:8000/authorize", headers=request.headers)
    if req.status_code != 200:
        abort(req.status_code)
    actor = json.loads(req.content).get('result')
    if not request.json:
        abort(400)
    reservation = Reservation(
        id=uuid.uuid4(),
        name=request.json.get('name'),
        phone_number=request.json.get('phone_number'),
        status="reserved",
        email=request.json.get('email'),
        created_by=actor,
        reservation_date= datetime.datetime(int(request.json.get('year')), int(request.json.get('month')), int(request.json.get('date')),int(request.json.get('hour')))
    )

    db.session.add(reservation)
    db.session.commit()
    return jsonify(reservation.to_json()), 201

@app.route("/<reservation_id>", methods=["DELETE"])
def delete_reservation(reservation_id):
    reservation = Reservation.query.get(reservation_id)
    if reservation is None:
        abort(404)
    db.session.delete(reservation)
    db.session.commit()
    return jsonify({'result': True})


@app.route("/<reservation_id>", methods=["POST"])
def update_reservation(reservation_id):
    req = requests.get("http://staff-service:8000/authorize", headers=request.headers)
    if req.status_code != 200:
        abort(req.status_code)
    actor = json.loads(req.content).get('result')
    if not request.json:
        abort(400)
    reservation = Reservation.query.get(reservation_id)
    if reservation is None:
        abort(404)

    reservation.status= request.json.get('status')

    reservation.modified_by = actor

    reservation.modified_at = datetime.datetime.now()

    db.session.commit()
    return jsonify(reservation.to_json())


# db.create_all()
# db.session.add(Reservation(name="Flaskaa",phone_number="000000", status="status", email="exasmple@example.com"))
# db.session.commit()

if __name__ == '__main__':
    try:
        db.create_all()
    except Exception as e:
        pass
    app.run(debug=True, host="0.0.0.0")
