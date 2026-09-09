import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from governanceOnDjango.delegates.MatterDelegate import MatterDelegate

 #======================================================================
# 
# Encapsulates data for View Matter
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MatterView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Matter index.")

def get(request, matterId ):
	delegate = MatterDelegate()
	responseData = delegate.get( matterId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	matter = json.loads(request.body)
	delegate = MatterDelegate()
	responseData = delegate.createFromJson( matter )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	matter = json.loads(request.body)
	delegate = MatterDelegate()
	responseData = delegate.save( matter )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, matterId ):
	delegate = MatterDelegate()
	responseData = delegate.delete( matterId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = MatterDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrganization( request, matterId, OrganizationId ):
	delegate = MatterDelegate()
	responseData = delegate.saveOrganization( matterId, OrganizationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrganization( request, matterId ):
	delegate = MatterDelegate()
	responseData = delegate.deleteOrganization( matterId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addLegalHolds( request, matterId, LegalHoldsIds ):
	delegate = MatterDelegate()
	responseData = delegate.addLegalHolds( matterId, LegalHoldsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeLegalHolds( request, matterId, LegalHoldsIds ):
	delegate = MatterDelegate()
	responseData = delegate.removeLegalHolds( matterId, LegalHoldsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDataBreaches( request, matterId, DataBreachesIds ):
	delegate = MatterDelegate()
	responseData = delegate.addDataBreaches( matterId, DataBreachesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDataBreaches( request, matterId, DataBreachesIds ):
	delegate = MatterDelegate()
	responseData = delegate.removeDataBreaches( matterId, DataBreachesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addContracts( request, matterId, ContractsIds ):
	delegate = MatterDelegate()
	responseData = delegate.addContracts( matterId, ContractsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeContracts( request, matterId, ContractsIds ):
	delegate = MatterDelegate()
	responseData = delegate.removeContracts( matterId, ContractsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

