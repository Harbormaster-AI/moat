import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from healthcareOnDjango.delegates.CarePlanDelegate import CarePlanDelegate

 #======================================================================
# 
# Encapsulates data for View CarePlan
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CarePlanView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the CarePlan index.")

def get(request, carePlanId ):
	delegate = CarePlanDelegate()
	responseData = delegate.get( carePlanId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	carePlan = json.loads(request.body)
	delegate = CarePlanDelegate()
	responseData = delegate.createFromJson( carePlan )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	carePlan = json.loads(request.body)
	delegate = CarePlanDelegate()
	responseData = delegate.save( carePlan )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, carePlanId ):
	delegate = CarePlanDelegate()
	responseData = delegate.delete( carePlanId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = CarePlanDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPatient( request, carePlanId, PatientId ):
	delegate = CarePlanDelegate()
	responseData = delegate.savePatient( carePlanId, PatientId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPatient( request, carePlanId ):
	delegate = CarePlanDelegate()
	responseData = delegate.deletePatient( carePlanId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCareTeam( request, carePlanId, CareTeamId ):
	delegate = CarePlanDelegate()
	responseData = delegate.saveCareTeam( carePlanId, CareTeamId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCareTeam( request, carePlanId ):
	delegate = CarePlanDelegate()
	responseData = delegate.deleteCareTeam( carePlanId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addEncounters( request, carePlanId, EncountersIds ):
	delegate = CarePlanDelegate()
	responseData = delegate.addEncounters( carePlanId, EncountersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeEncounters( request, carePlanId, EncountersIds ):
	delegate = CarePlanDelegate()
	responseData = delegate.removeEncounters( carePlanId, EncountersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addTasks( request, carePlanId, TasksIds ):
	delegate = CarePlanDelegate()
	responseData = delegate.addTasks( carePlanId, TasksIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeTasks( request, carePlanId, TasksIds ):
	delegate = CarePlanDelegate()
	responseData = delegate.removeTasks( carePlanId, TasksIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

