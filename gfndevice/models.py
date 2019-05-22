import re

from gfndevice.db import db
from datetime import datetime
from sqlalchemy.orm import validates

class Room(db.Model):
    __tablename__ = 'rooms'
    id = db.Column(db.Integer, primary_key=True)
    name = db.Column(db.String(30), nullable=False, unique=True)

    devices = db.relationship('Device', back_populates="room")
    

    def to_dict(self):
        return {
            "id": self.id,
            "name": self.name,
            "devices": [d.id for d in self.devices]
        }


    @staticmethod
    def create(name=None):
        if not name:
            name = "Room{}".format(Room.query.count() + 1)
        
        return Room(name=name)


class Device(db.Model):
    __tablename__ = 'devices'
    id = db.Column(db.Integer, primary_key=True)
    name = db.Column(db.String(30), nullable=False)
    mac = db.Column(db.String(17), nullable=False, unique=True) # MACs are 12 digits + 5 ':'

    room_id = db.Column(db.Integer, db.ForeignKey('rooms.id'), nullable=False)
    room = db.relationship('Room', back_populates="devices")

    readings = db.relationship('Reading', back_populates="device")


    @validates('mac')
    def validate_mac(self, key, mac):
        if not mac:
            raise AssertionError("No MAC provided")

        if not re.match("[0-9a-f]{2}([-:]?)[0-9a-f]{2}(\\1[0-9a-f]{2}){4}$", mac.lower()):
            raise AssertionError("Invalid MAC address")

        if Device.query.filter_by(mac=mac).first():
            raise AssertionError("MAC already in use")

        return mac


    @validates('name')
    def validate_name(self, key, name):
        if not name:
            raise AssertionError("No name provided")

        if Device.query.filter_by(name=name).first():
            raise AssertionError("Name already in use")

        return name


    def to_dict(self):
        return {
            "id": self.id,
            "name": self.name,
            "mac": self.mac,
            "room": self.room_id
        }


    def __repr__(self): # pragma: no cover
        return "<Device {}, name={}, mac={}, room={}>".format(self.id, self.name, self.mac, self.room_id)


class Reading(db.Model):
    __tablename__ = 'readings'
    id = db.Column(db.Integer, primary_key=True)
    timestamp = db.Column(db.DateTime, nullable=False, default=datetime.utcnow)
    name = db.Column(db.String, nullable=False)
    value = db.Column(db.Float, nullable=False)
    unit = db.Column(db.String, default=None)

    device_id = db.Column(db.Integer, db.ForeignKey('devices.id'), nullable=False)
    device = db.relationship('Device', back_populates="readings")


    @validates('name', 'unit')
    def validate_name(self, key, name):
        if not name:
            raise AssertionError("No name provided")

        if not isinstance(name, str):
            raise AssertionError("Invalid {}".format(key))

        return name


    @validates('value')
    def validate_value(self, key, value):
        if not value:
            raise AssertionError("No value provided")

        if not isinstance(value, int):
            raise AssertionError("Invalid value")

        return value


    def to_dict(self):
        return {
            "id": self.id,
            "timestamp": self.timestamp,
            "name": self.name,
            "value": self.value,
            "unit": self.unit,
            "device": self.device.id
        }


    def __repr__(self): # pragma: no cover
        return "<Reading {}, name={}, value={}, unit={}, device={}>".format(self.id, self.name, self.value, self.unit, self.device.id)
