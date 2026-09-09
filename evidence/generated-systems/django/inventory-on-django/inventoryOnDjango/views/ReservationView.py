import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from inventoryOnDjango.delegates.ReservationDelegate import ReservationDelegate

 #======================================================================
# 
# Encapsulates data for View Reservation
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ReservationView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Reservation index.")

def get(request, reservationId ):
	delegate = ReservationDelegate()
	responseData = delegate.get( reservationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	reservation = json.loads(request.body)
	delegate = ReservationDelegate()
	responseData = delegate.createFromJson( reservation )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	reservation = json.loads(request.body)
	delegate = ReservationDelegate()
	responseData = delegate.save( reservation )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, reservationId ):
	delegate = ReservationDelegate()
	responseData = delegate.delete( reservationId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ReservationDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignSku( request, reservationId, SkuId ):
	delegate = ReservationDelegate()
	responseData = delegate.saveSku( reservationId, SkuId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignSku( request, reservationId ):
	delegate = ReservationDelegate()
	responseData = delegate.deleteSku( reservationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignWarehouse( request, reservationId, WarehouseId ):
	delegate = ReservationDelegate()
	responseData = delegate.saveWarehouse( reservationId, WarehouseId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignWarehouse( request, reservationId ):
	delegate = ReservationDelegate()
	responseData = delegate.deleteWarehouse( reservationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignLocation( request, reservationId, LocationId ):
	delegate = ReservationDelegate()
	responseData = delegate.saveLocation( reservationId, LocationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignLocation( request, reservationId ):
	delegate = ReservationDelegate()
	responseData = delegate.deleteLocation( reservationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignInventoryItem( request, reservationId, InventoryItemId ):
	delegate = ReservationDelegate()
	responseData = delegate.saveInventoryItem( reservationId, InventoryItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignInventoryItem( request, reservationId ):
	delegate = ReservationDelegate()
	responseData = delegate.deleteInventoryItem( reservationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignLot( request, reservationId, LotId ):
	delegate = ReservationDelegate()
	responseData = delegate.saveLot( reservationId, LotId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignLot( request, reservationId ):
	delegate = ReservationDelegate()
	responseData = delegate.deleteLot( reservationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignDemandSignal( request, reservationId, DemandSignalId ):
	delegate = ReservationDelegate()
	responseData = delegate.saveDemandSignal( reservationId, DemandSignalId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignDemandSignal( request, reservationId ):
	delegate = ReservationDelegate()
	responseData = delegate.deleteDemandSignal( reservationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addSerialNumbers( request, reservationId, SerialNumbersIds ):
	delegate = ReservationDelegate()
	responseData = delegate.addSerialNumbers( reservationId, SerialNumbersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeSerialNumbers( request, reservationId, SerialNumbersIds ):
	delegate = ReservationDelegate()
	responseData = delegate.removeSerialNumbers( reservationId, SerialNumbersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

