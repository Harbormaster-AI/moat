import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from insuranceOnDjango.delegates.ExposureDelegate import ExposureDelegate

 #======================================================================
# 
# Encapsulates data for View Exposure
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ExposureView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Exposure index.")

def get(request, exposureId ):
	delegate = ExposureDelegate()
	responseData = delegate.get( exposureId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	exposure = json.loads(request.body)
	delegate = ExposureDelegate()
	responseData = delegate.createFromJson( exposure )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	exposure = json.loads(request.body)
	delegate = ExposureDelegate()
	responseData = delegate.save( exposure )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, exposureId ):
	delegate = ExposureDelegate()
	responseData = delegate.delete( exposureId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ExposureDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignClaim( request, exposureId, ClaimId ):
	delegate = ExposureDelegate()
	responseData = delegate.saveClaim( exposureId, ClaimId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignClaim( request, exposureId ):
	delegate = ExposureDelegate()
	responseData = delegate.deleteClaim( exposureId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPolicyCoverage( request, exposureId, PolicyCoverageId ):
	delegate = ExposureDelegate()
	responseData = delegate.savePolicyCoverage( exposureId, PolicyCoverageId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPolicyCoverage( request, exposureId ):
	delegate = ExposureDelegate()
	responseData = delegate.deletePolicyCoverage( exposureId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignInsuredObject( request, exposureId, InsuredObjectId ):
	delegate = ExposureDelegate()
	responseData = delegate.saveInsuredObject( exposureId, InsuredObjectId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignInsuredObject( request, exposureId ):
	delegate = ExposureDelegate()
	responseData = delegate.deleteInsuredObject( exposureId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addReserves( request, exposureId, ReservesIds ):
	delegate = ExposureDelegate()
	responseData = delegate.addReserves( exposureId, ReservesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeReserves( request, exposureId, ReservesIds ):
	delegate = ExposureDelegate()
	responseData = delegate.removeReserves( exposureId, ReservesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPayments( request, exposureId, PaymentsIds ):
	delegate = ExposureDelegate()
	responseData = delegate.addPayments( exposureId, PaymentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePayments( request, exposureId, PaymentsIds ):
	delegate = ExposureDelegate()
	responseData = delegate.removePayments( exposureId, PaymentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

