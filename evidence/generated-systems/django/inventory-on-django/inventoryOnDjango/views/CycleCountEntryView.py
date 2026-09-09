import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from inventoryOnDjango.delegates.CycleCountEntryDelegate import CycleCountEntryDelegate

 #======================================================================
# 
# Encapsulates data for View CycleCountEntry
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CycleCountEntryView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the CycleCountEntry index.")

def get(request, cycleCountEntryId ):
	delegate = CycleCountEntryDelegate()
	responseData = delegate.get( cycleCountEntryId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	cycleCountEntry = json.loads(request.body)
	delegate = CycleCountEntryDelegate()
	responseData = delegate.createFromJson( cycleCountEntry )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	cycleCountEntry = json.loads(request.body)
	delegate = CycleCountEntryDelegate()
	responseData = delegate.save( cycleCountEntry )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, cycleCountEntryId ):
	delegate = CycleCountEntryDelegate()
	responseData = delegate.delete( cycleCountEntryId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = CycleCountEntryDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCycleCount( request, cycleCountEntryId, CycleCountId ):
	delegate = CycleCountEntryDelegate()
	responseData = delegate.saveCycleCount( cycleCountEntryId, CycleCountId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCycleCount( request, cycleCountEntryId ):
	delegate = CycleCountEntryDelegate()
	responseData = delegate.deleteCycleCount( cycleCountEntryId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignSku( request, cycleCountEntryId, SkuId ):
	delegate = CycleCountEntryDelegate()
	responseData = delegate.saveSku( cycleCountEntryId, SkuId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignSku( request, cycleCountEntryId ):
	delegate = CycleCountEntryDelegate()
	responseData = delegate.deleteSku( cycleCountEntryId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignLot( request, cycleCountEntryId, LotId ):
	delegate = CycleCountEntryDelegate()
	responseData = delegate.saveLot( cycleCountEntryId, LotId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignLot( request, cycleCountEntryId ):
	delegate = CycleCountEntryDelegate()
	responseData = delegate.deleteLot( cycleCountEntryId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignLocation( request, cycleCountEntryId, LocationId ):
	delegate = CycleCountEntryDelegate()
	responseData = delegate.saveLocation( cycleCountEntryId, LocationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignLocation( request, cycleCountEntryId ):
	delegate = CycleCountEntryDelegate()
	responseData = delegate.deleteLocation( cycleCountEntryId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addSerialNumbers( request, cycleCountEntryId, SerialNumbersIds ):
	delegate = CycleCountEntryDelegate()
	responseData = delegate.addSerialNumbers( cycleCountEntryId, SerialNumbersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeSerialNumbers( request, cycleCountEntryId, SerialNumbersIds ):
	delegate = CycleCountEntryDelegate()
	responseData = delegate.removeSerialNumbers( cycleCountEntryId, SerialNumbersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

