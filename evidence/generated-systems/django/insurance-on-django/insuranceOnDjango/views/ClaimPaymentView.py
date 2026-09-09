import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from insuranceOnDjango.delegates.ClaimPaymentDelegate import ClaimPaymentDelegate

 #======================================================================
# 
# Encapsulates data for View ClaimPayment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ClaimPaymentView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the ClaimPayment index.")

def get(request, claimPaymentId ):
	delegate = ClaimPaymentDelegate()
	responseData = delegate.get( claimPaymentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	claimPayment = json.loads(request.body)
	delegate = ClaimPaymentDelegate()
	responseData = delegate.createFromJson( claimPayment )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	claimPayment = json.loads(request.body)
	delegate = ClaimPaymentDelegate()
	responseData = delegate.save( claimPayment )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, claimPaymentId ):
	delegate = ClaimPaymentDelegate()
	responseData = delegate.delete( claimPaymentId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ClaimPaymentDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignClaim( request, claimPaymentId, ClaimId ):
	delegate = ClaimPaymentDelegate()
	responseData = delegate.saveClaim( claimPaymentId, ClaimId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignClaim( request, claimPaymentId ):
	delegate = ClaimPaymentDelegate()
	responseData = delegate.deleteClaim( claimPaymentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignExposure( request, claimPaymentId, ExposureId ):
	delegate = ClaimPaymentDelegate()
	responseData = delegate.saveExposure( claimPaymentId, ExposureId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignExposure( request, claimPaymentId ):
	delegate = ClaimPaymentDelegate()
	responseData = delegate.deleteExposure( claimPaymentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignBeneficiary( request, claimPaymentId, BeneficiaryId ):
	delegate = ClaimPaymentDelegate()
	responseData = delegate.saveBeneficiary( claimPaymentId, BeneficiaryId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignBeneficiary( request, claimPaymentId ):
	delegate = ClaimPaymentDelegate()
	responseData = delegate.deleteBeneficiary( claimPaymentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignServiceProvider( request, claimPaymentId, ServiceProviderId ):
	delegate = ClaimPaymentDelegate()
	responseData = delegate.saveServiceProvider( claimPaymentId, ServiceProviderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignServiceProvider( request, claimPaymentId ):
	delegate = ClaimPaymentDelegate()
	responseData = delegate.deleteServiceProvider( claimPaymentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCustomer( request, claimPaymentId, CustomerId ):
	delegate = ClaimPaymentDelegate()
	responseData = delegate.saveCustomer( claimPaymentId, CustomerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCustomer( request, claimPaymentId ):
	delegate = ClaimPaymentDelegate()
	responseData = delegate.deleteCustomer( claimPaymentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

