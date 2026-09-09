import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from manufacturingOnDjango.delegates.InspectionResultDelegate import InspectionResultDelegate

 #======================================================================
# 
# Encapsulates data for View InspectionResult
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InspectionResultView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the InspectionResult index.")

def get(request, inspectionResultId ):
	delegate = InspectionResultDelegate()
	responseData = delegate.get( inspectionResultId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	inspectionResult = json.loads(request.body)
	delegate = InspectionResultDelegate()
	responseData = delegate.createFromJson( inspectionResult )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	inspectionResult = json.loads(request.body)
	delegate = InspectionResultDelegate()
	responseData = delegate.save( inspectionResult )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, inspectionResultId ):
	delegate = InspectionResultDelegate()
	responseData = delegate.delete( inspectionResultId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = InspectionResultDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignInspectionLot( request, inspectionResultId, InspectionLotId ):
	delegate = InspectionResultDelegate()
	responseData = delegate.saveInspectionLot( inspectionResultId, InspectionLotId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignInspectionLot( request, inspectionResultId ):
	delegate = InspectionResultDelegate()
	responseData = delegate.deleteInspectionLot( inspectionResultId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCharacteristic( request, inspectionResultId, CharacteristicId ):
	delegate = InspectionResultDelegate()
	responseData = delegate.saveCharacteristic( inspectionResultId, CharacteristicId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCharacteristic( request, inspectionResultId ):
	delegate = InspectionResultDelegate()
	responseData = delegate.deleteCharacteristic( inspectionResultId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

