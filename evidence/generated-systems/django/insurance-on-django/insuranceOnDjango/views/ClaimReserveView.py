import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from insuranceOnDjango.delegates.ClaimReserveDelegate import ClaimReserveDelegate

 #======================================================================
# 
# Encapsulates data for View ClaimReserve
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ClaimReserveView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the ClaimReserve index.")

def get(request, claimReserveId ):
	delegate = ClaimReserveDelegate()
	responseData = delegate.get( claimReserveId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	claimReserve = json.loads(request.body)
	delegate = ClaimReserveDelegate()
	responseData = delegate.createFromJson( claimReserve )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	claimReserve = json.loads(request.body)
	delegate = ClaimReserveDelegate()
	responseData = delegate.save( claimReserve )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, claimReserveId ):
	delegate = ClaimReserveDelegate()
	responseData = delegate.delete( claimReserveId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ClaimReserveDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignClaim( request, claimReserveId, ClaimId ):
	delegate = ClaimReserveDelegate()
	responseData = delegate.saveClaim( claimReserveId, ClaimId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignClaim( request, claimReserveId ):
	delegate = ClaimReserveDelegate()
	responseData = delegate.deleteClaim( claimReserveId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignExposure( request, claimReserveId, ExposureId ):
	delegate = ClaimReserveDelegate()
	responseData = delegate.saveExposure( claimReserveId, ExposureId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignExposure( request, claimReserveId ):
	delegate = ClaimReserveDelegate()
	responseData = delegate.deleteExposure( claimReserveId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

