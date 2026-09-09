import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from insuranceOnDjango.delegates.ClaimDelegate import ClaimDelegate

 #======================================================================
# 
# Encapsulates data for View Claim
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ClaimView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Claim index.")

def get(request, claimId ):
	delegate = ClaimDelegate()
	responseData = delegate.get( claimId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	claim = json.loads(request.body)
	delegate = ClaimDelegate()
	responseData = delegate.createFromJson( claim )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	claim = json.loads(request.body)
	delegate = ClaimDelegate()
	responseData = delegate.save( claim )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, claimId ):
	delegate = ClaimDelegate()
	responseData = delegate.delete( claimId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ClaimDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPolicy( request, claimId, PolicyId ):
	delegate = ClaimDelegate()
	responseData = delegate.savePolicy( claimId, PolicyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPolicy( request, claimId ):
	delegate = ClaimDelegate()
	responseData = delegate.deletePolicy( claimId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCustomer( request, claimId, CustomerId ):
	delegate = ClaimDelegate()
	responseData = delegate.saveCustomer( claimId, CustomerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCustomer( request, claimId ):
	delegate = ClaimDelegate()
	responseData = delegate.deleteCustomer( claimId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignAdjuster( request, claimId, AdjusterId ):
	delegate = ClaimDelegate()
	responseData = delegate.saveAdjuster( claimId, AdjusterId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignAdjuster( request, claimId ):
	delegate = ClaimDelegate()
	responseData = delegate.deleteAdjuster( claimId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignIncident( request, claimId, IncidentId ):
	delegate = ClaimDelegate()
	responseData = delegate.saveIncident( claimId, IncidentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignIncident( request, claimId ):
	delegate = ClaimDelegate()
	responseData = delegate.deleteIncident( claimId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addExposures( request, claimId, ExposuresIds ):
	delegate = ClaimDelegate()
	responseData = delegate.addExposures( claimId, ExposuresIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeExposures( request, claimId, ExposuresIds ):
	delegate = ClaimDelegate()
	responseData = delegate.removeExposures( claimId, ExposuresIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addReserves( request, claimId, ReservesIds ):
	delegate = ClaimDelegate()
	responseData = delegate.addReserves( claimId, ReservesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeReserves( request, claimId, ReservesIds ):
	delegate = ClaimDelegate()
	responseData = delegate.removeReserves( claimId, ReservesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addClaimPayments( request, claimId, ClaimPaymentsIds ):
	delegate = ClaimDelegate()
	responseData = delegate.addClaimPayments( claimId, ClaimPaymentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeClaimPayments( request, claimId, ClaimPaymentsIds ):
	delegate = ClaimDelegate()
	responseData = delegate.removeClaimPayments( claimId, ClaimPaymentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addServiceProviders( request, claimId, ServiceProvidersIds ):
	delegate = ClaimDelegate()
	responseData = delegate.addServiceProviders( claimId, ServiceProvidersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeServiceProviders( request, claimId, ServiceProvidersIds ):
	delegate = ClaimDelegate()
	responseData = delegate.removeServiceProviders( claimId, ServiceProvidersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addSubrogations( request, claimId, SubrogationsIds ):
	delegate = ClaimDelegate()
	responseData = delegate.addSubrogations( claimId, SubrogationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeSubrogations( request, claimId, SubrogationsIds ):
	delegate = ClaimDelegate()
	responseData = delegate.removeSubrogations( claimId, SubrogationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

