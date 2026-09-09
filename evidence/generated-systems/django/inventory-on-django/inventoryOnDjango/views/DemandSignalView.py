import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from inventoryOnDjango.delegates.DemandSignalDelegate import DemandSignalDelegate

 #======================================================================
# 
# Encapsulates data for View DemandSignal
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DemandSignalView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the DemandSignal index.")

def get(request, demandSignalId ):
	delegate = DemandSignalDelegate()
	responseData = delegate.get( demandSignalId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	demandSignal = json.loads(request.body)
	delegate = DemandSignalDelegate()
	responseData = delegate.createFromJson( demandSignal )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	demandSignal = json.loads(request.body)
	delegate = DemandSignalDelegate()
	responseData = delegate.save( demandSignal )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, demandSignalId ):
	delegate = DemandSignalDelegate()
	responseData = delegate.delete( demandSignalId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = DemandSignalDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignSku( request, demandSignalId, SkuId ):
	delegate = DemandSignalDelegate()
	responseData = delegate.saveSku( demandSignalId, SkuId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignSku( request, demandSignalId ):
	delegate = DemandSignalDelegate()
	responseData = delegate.deleteSku( demandSignalId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addReservations( request, demandSignalId, ReservationsIds ):
	delegate = DemandSignalDelegate()
	responseData = delegate.addReservations( demandSignalId, ReservationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeReservations( request, demandSignalId, ReservationsIds ):
	delegate = DemandSignalDelegate()
	responseData = delegate.removeReservations( demandSignalId, ReservationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

