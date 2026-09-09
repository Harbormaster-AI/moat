import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from manufacturingOnDjango.delegates.InspectionCharacteristicDelegate import InspectionCharacteristicDelegate

 #======================================================================
# 
# Encapsulates data for View InspectionCharacteristic
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InspectionCharacteristicView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the InspectionCharacteristic index.")

def get(request, inspectionCharacteristicId ):
	delegate = InspectionCharacteristicDelegate()
	responseData = delegate.get( inspectionCharacteristicId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	inspectionCharacteristic = json.loads(request.body)
	delegate = InspectionCharacteristicDelegate()
	responseData = delegate.createFromJson( inspectionCharacteristic )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	inspectionCharacteristic = json.loads(request.body)
	delegate = InspectionCharacteristicDelegate()
	responseData = delegate.save( inspectionCharacteristic )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, inspectionCharacteristicId ):
	delegate = InspectionCharacteristicDelegate()
	responseData = delegate.delete( inspectionCharacteristicId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = InspectionCharacteristicDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignInspectionPlan( request, inspectionCharacteristicId, InspectionPlanId ):
	delegate = InspectionCharacteristicDelegate()
	responseData = delegate.saveInspectionPlan( inspectionCharacteristicId, InspectionPlanId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignInspectionPlan( request, inspectionCharacteristicId ):
	delegate = InspectionCharacteristicDelegate()
	responseData = delegate.deleteInspectionPlan( inspectionCharacteristicId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

