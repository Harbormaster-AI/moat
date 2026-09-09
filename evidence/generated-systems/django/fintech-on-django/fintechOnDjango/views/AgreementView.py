import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from fintechOnDjango.delegates.AgreementDelegate import AgreementDelegate

 #======================================================================
# 
# Encapsulates data for View Agreement
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AgreementView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Agreement index.")

def get(request, agreementId ):
	delegate = AgreementDelegate()
	responseData = delegate.get( agreementId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	agreement = json.loads(request.body)
	delegate = AgreementDelegate()
	responseData = delegate.createFromJson( agreement )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	agreement = json.loads(request.body)
	delegate = AgreementDelegate()
	responseData = delegate.save( agreement )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, agreementId ):
	delegate = AgreementDelegate()
	responseData = delegate.delete( agreementId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = AgreementDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCustomer( request, agreementId, CustomerId ):
	delegate = AgreementDelegate()
	responseData = delegate.saveCustomer( agreementId, CustomerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCustomer( request, agreementId ):
	delegate = AgreementDelegate()
	responseData = delegate.deleteCustomer( agreementId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignProductOffering( request, agreementId, ProductOfferingId ):
	delegate = AgreementDelegate()
	responseData = delegate.saveProductOffering( agreementId, ProductOfferingId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignProductOffering( request, agreementId ):
	delegate = AgreementDelegate()
	responseData = delegate.deleteProductOffering( agreementId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

