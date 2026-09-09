import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from insuranceOnDjango.delegates.SubrogationRecoveryDelegate import SubrogationRecoveryDelegate

 #======================================================================
# 
# Encapsulates data for View SubrogationRecovery
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SubrogationRecoveryView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the SubrogationRecovery index.")

def get(request, subrogationRecoveryId ):
	delegate = SubrogationRecoveryDelegate()
	responseData = delegate.get( subrogationRecoveryId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	subrogationRecovery = json.loads(request.body)
	delegate = SubrogationRecoveryDelegate()
	responseData = delegate.createFromJson( subrogationRecovery )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	subrogationRecovery = json.loads(request.body)
	delegate = SubrogationRecoveryDelegate()
	responseData = delegate.save( subrogationRecovery )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, subrogationRecoveryId ):
	delegate = SubrogationRecoveryDelegate()
	responseData = delegate.delete( subrogationRecoveryId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = SubrogationRecoveryDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignClaim( request, subrogationRecoveryId, ClaimId ):
	delegate = SubrogationRecoveryDelegate()
	responseData = delegate.saveClaim( subrogationRecoveryId, ClaimId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignClaim( request, subrogationRecoveryId ):
	delegate = SubrogationRecoveryDelegate()
	responseData = delegate.deleteClaim( subrogationRecoveryId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignExposure( request, subrogationRecoveryId, ExposureId ):
	delegate = SubrogationRecoveryDelegate()
	responseData = delegate.saveExposure( subrogationRecoveryId, ExposureId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignExposure( request, subrogationRecoveryId ):
	delegate = SubrogationRecoveryDelegate()
	responseData = delegate.deleteExposure( subrogationRecoveryId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCounterparty( request, subrogationRecoveryId, CounterpartyId ):
	delegate = SubrogationRecoveryDelegate()
	responseData = delegate.saveCounterparty( subrogationRecoveryId, CounterpartyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCounterparty( request, subrogationRecoveryId ):
	delegate = SubrogationRecoveryDelegate()
	responseData = delegate.deleteCounterparty( subrogationRecoveryId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

