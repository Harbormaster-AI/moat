import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from governanceOnDjango.delegates.ContractDelegate import ContractDelegate

 #======================================================================
# 
# Encapsulates data for View Contract
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ContractView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Contract index.")

def get(request, contractId ):
	delegate = ContractDelegate()
	responseData = delegate.get( contractId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	contract = json.loads(request.body)
	delegate = ContractDelegate()
	responseData = delegate.createFromJson( contract )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	contract = json.loads(request.body)
	delegate = ContractDelegate()
	responseData = delegate.save( contract )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, contractId ):
	delegate = ContractDelegate()
	responseData = delegate.delete( contractId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ContractDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignThirdParty( request, contractId, ThirdPartyId ):
	delegate = ContractDelegate()
	responseData = delegate.saveThirdParty( contractId, ThirdPartyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignThirdParty( request, contractId ):
	delegate = ContractDelegate()
	responseData = delegate.deleteThirdParty( contractId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignMatter( request, contractId, MatterId ):
	delegate = ContractDelegate()
	responseData = delegate.saveMatter( contractId, MatterId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignMatter( request, contractId ):
	delegate = ContractDelegate()
	responseData = delegate.deleteMatter( contractId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addObligations( request, contractId, ObligationsIds ):
	delegate = ContractDelegate()
	responseData = delegate.addObligations( contractId, ObligationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeObligations( request, contractId, ObligationsIds ):
	delegate = ContractDelegate()
	responseData = delegate.removeObligations( contractId, ObligationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDataProcessingActivities( request, contractId, DataProcessingActivitiesIds ):
	delegate = ContractDelegate()
	responseData = delegate.addDataProcessingActivities( contractId, DataProcessingActivitiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDataProcessingActivities( request, contractId, DataProcessingActivitiesIds ):
	delegate = ContractDelegate()
	responseData = delegate.removeDataProcessingActivities( contractId, DataProcessingActivitiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

