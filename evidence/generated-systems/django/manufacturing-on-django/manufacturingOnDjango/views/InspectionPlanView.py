import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from manufacturingOnDjango.delegates.InspectionPlanDelegate import InspectionPlanDelegate

 #======================================================================
# 
# Encapsulates data for View InspectionPlan
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InspectionPlanView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the InspectionPlan index.")

def get(request, inspectionPlanId ):
	delegate = InspectionPlanDelegate()
	responseData = delegate.get( inspectionPlanId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	inspectionPlan = json.loads(request.body)
	delegate = InspectionPlanDelegate()
	responseData = delegate.createFromJson( inspectionPlan )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	inspectionPlan = json.loads(request.body)
	delegate = InspectionPlanDelegate()
	responseData = delegate.save( inspectionPlan )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, inspectionPlanId ):
	delegate = InspectionPlanDelegate()
	responseData = delegate.delete( inspectionPlanId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = InspectionPlanDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignItem( request, inspectionPlanId, ItemId ):
	delegate = InspectionPlanDelegate()
	responseData = delegate.saveItem( inspectionPlanId, ItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignItem( request, inspectionPlanId ):
	delegate = InspectionPlanDelegate()
	responseData = delegate.deleteItem( inspectionPlanId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCharacteristics( request, inspectionPlanId, CharacteristicsIds ):
	delegate = InspectionPlanDelegate()
	responseData = delegate.addCharacteristics( inspectionPlanId, CharacteristicsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCharacteristics( request, inspectionPlanId, CharacteristicsIds ):
	delegate = InspectionPlanDelegate()
	responseData = delegate.removeCharacteristics( inspectionPlanId, CharacteristicsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

