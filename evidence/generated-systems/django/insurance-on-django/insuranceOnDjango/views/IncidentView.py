import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from insuranceOnDjango.delegates.IncidentDelegate import IncidentDelegate

 #======================================================================
# 
# Encapsulates data for View Incident
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class IncidentView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Incident index.")

def get(request, incidentId ):
	delegate = IncidentDelegate()
	responseData = delegate.get( incidentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	incident = json.loads(request.body)
	delegate = IncidentDelegate()
	responseData = delegate.createFromJson( incident )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	incident = json.loads(request.body)
	delegate = IncidentDelegate()
	responseData = delegate.save( incident )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, incidentId ):
	delegate = IncidentDelegate()
	responseData = delegate.delete( incidentId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = IncidentDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignClaim( request, incidentId, ClaimId ):
	delegate = IncidentDelegate()
	responseData = delegate.saveClaim( incidentId, ClaimId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignClaim( request, incidentId ):
	delegate = IncidentDelegate()
	responseData = delegate.deleteClaim( incidentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addInsuredObjects( request, incidentId, InsuredObjectsIds ):
	delegate = IncidentDelegate()
	responseData = delegate.addInsuredObjects( incidentId, InsuredObjectsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeInsuredObjects( request, incidentId, InsuredObjectsIds ):
	delegate = IncidentDelegate()
	responseData = delegate.removeInsuredObjects( incidentId, InsuredObjectsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

