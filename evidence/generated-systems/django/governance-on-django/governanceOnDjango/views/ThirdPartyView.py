import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from governanceOnDjango.delegates.ThirdPartyDelegate import ThirdPartyDelegate

 #======================================================================
# 
# Encapsulates data for View ThirdParty
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ThirdPartyView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the ThirdParty index.")

def get(request, thirdPartyId ):
	delegate = ThirdPartyDelegate()
	responseData = delegate.get( thirdPartyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	thirdParty = json.loads(request.body)
	delegate = ThirdPartyDelegate()
	responseData = delegate.createFromJson( thirdParty )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	thirdParty = json.loads(request.body)
	delegate = ThirdPartyDelegate()
	responseData = delegate.save( thirdParty )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, thirdPartyId ):
	delegate = ThirdPartyDelegate()
	responseData = delegate.delete( thirdPartyId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ThirdPartyDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrganization( request, thirdPartyId, OrganizationId ):
	delegate = ThirdPartyDelegate()
	responseData = delegate.saveOrganization( thirdPartyId, OrganizationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrganization( request, thirdPartyId ):
	delegate = ThirdPartyDelegate()
	responseData = delegate.deleteOrganization( thirdPartyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addProcessingActivities( request, thirdPartyId, ProcessingActivitiesIds ):
	delegate = ThirdPartyDelegate()
	responseData = delegate.addProcessingActivities( thirdPartyId, ProcessingActivitiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeProcessingActivities( request, thirdPartyId, ProcessingActivitiesIds ):
	delegate = ThirdPartyDelegate()
	responseData = delegate.removeProcessingActivities( thirdPartyId, ProcessingActivitiesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addAssessments( request, thirdPartyId, AssessmentsIds ):
	delegate = ThirdPartyDelegate()
	responseData = delegate.addAssessments( thirdPartyId, AssessmentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAssessments( request, thirdPartyId, AssessmentsIds ):
	delegate = ThirdPartyDelegate()
	responseData = delegate.removeAssessments( thirdPartyId, AssessmentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addContracts( request, thirdPartyId, ContractsIds ):
	delegate = ThirdPartyDelegate()
	responseData = delegate.addContracts( thirdPartyId, ContractsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeContracts( request, thirdPartyId, ContractsIds ):
	delegate = ThirdPartyDelegate()
	responseData = delegate.removeContracts( thirdPartyId, ContractsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addObligations( request, thirdPartyId, ObligationsIds ):
	delegate = ThirdPartyDelegate()
	responseData = delegate.addObligations( thirdPartyId, ObligationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeObligations( request, thirdPartyId, ObligationsIds ):
	delegate = ThirdPartyDelegate()
	responseData = delegate.removeObligations( thirdPartyId, ObligationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDataBreaches( request, thirdPartyId, DataBreachesIds ):
	delegate = ThirdPartyDelegate()
	responseData = delegate.addDataBreaches( thirdPartyId, DataBreachesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDataBreaches( request, thirdPartyId, DataBreachesIds ):
	delegate = ThirdPartyDelegate()
	responseData = delegate.removeDataBreaches( thirdPartyId, DataBreachesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

