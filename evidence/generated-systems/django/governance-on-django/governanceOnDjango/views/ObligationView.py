import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from governanceOnDjango.delegates.ObligationDelegate import ObligationDelegate

 #======================================================================
# 
# Encapsulates data for View Obligation
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ObligationView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Obligation index.")

def get(request, obligationId ):
	delegate = ObligationDelegate()
	responseData = delegate.get( obligationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	obligation = json.loads(request.body)
	delegate = ObligationDelegate()
	responseData = delegate.createFromJson( obligation )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	obligation = json.loads(request.body)
	delegate = ObligationDelegate()
	responseData = delegate.save( obligation )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, obligationId ):
	delegate = ObligationDelegate()
	responseData = delegate.delete( obligationId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ObligationDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignRegulation( request, obligationId, RegulationId ):
	delegate = ObligationDelegate()
	responseData = delegate.saveRegulation( obligationId, RegulationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignRegulation( request, obligationId ):
	delegate = ObligationDelegate()
	responseData = delegate.deleteRegulation( obligationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addControls( request, obligationId, ControlsIds ):
	delegate = ObligationDelegate()
	responseData = delegate.addControls( obligationId, ControlsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeControls( request, obligationId, ControlsIds ):
	delegate = ObligationDelegate()
	responseData = delegate.removeControls( obligationId, ControlsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPolicies( request, obligationId, PoliciesIds ):
	delegate = ObligationDelegate()
	responseData = delegate.addPolicies( obligationId, PoliciesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePolicies( request, obligationId, PoliciesIds ):
	delegate = ObligationDelegate()
	responseData = delegate.removePolicies( obligationId, PoliciesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addContracts( request, obligationId, ContractsIds ):
	delegate = ObligationDelegate()
	responseData = delegate.addContracts( obligationId, ContractsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeContracts( request, obligationId, ContractsIds ):
	delegate = ObligationDelegate()
	responseData = delegate.removeContracts( obligationId, ContractsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

