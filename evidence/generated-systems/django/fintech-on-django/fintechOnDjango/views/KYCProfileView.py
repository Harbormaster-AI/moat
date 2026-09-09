import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from fintechOnDjango.delegates.KYCProfileDelegate import KYCProfileDelegate

 #======================================================================
# 
# Encapsulates data for View KYCProfile
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class KYCProfileView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the KYCProfile index.")

def get(request, kYCProfileId ):
	delegate = KYCProfileDelegate()
	responseData = delegate.get( kYCProfileId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	kYCProfile = json.loads(request.body)
	delegate = KYCProfileDelegate()
	responseData = delegate.createFromJson( kYCProfile )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	kYCProfile = json.loads(request.body)
	delegate = KYCProfileDelegate()
	responseData = delegate.save( kYCProfile )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, kYCProfileId ):
	delegate = KYCProfileDelegate()
	responseData = delegate.delete( kYCProfileId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = KYCProfileDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCustomer( request, kYCProfileId, CustomerId ):
	delegate = KYCProfileDelegate()
	responseData = delegate.saveCustomer( kYCProfileId, CustomerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCustomer( request, kYCProfileId ):
	delegate = KYCProfileDelegate()
	responseData = delegate.deleteCustomer( kYCProfileId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDocuments( request, kYCProfileId, DocumentsIds ):
	delegate = KYCProfileDelegate()
	responseData = delegate.addDocuments( kYCProfileId, DocumentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDocuments( request, kYCProfileId, DocumentsIds ):
	delegate = KYCProfileDelegate()
	responseData = delegate.removeDocuments( kYCProfileId, DocumentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addScreenings( request, kYCProfileId, ScreeningsIds ):
	delegate = KYCProfileDelegate()
	responseData = delegate.addScreenings( kYCProfileId, ScreeningsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeScreenings( request, kYCProfileId, ScreeningsIds ):
	delegate = KYCProfileDelegate()
	responseData = delegate.removeScreenings( kYCProfileId, ScreeningsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addAddresses( request, kYCProfileId, AddressesIds ):
	delegate = KYCProfileDelegate()
	responseData = delegate.addAddresses( kYCProfileId, AddressesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAddresses( request, kYCProfileId, AddressesIds ):
	delegate = KYCProfileDelegate()
	responseData = delegate.removeAddresses( kYCProfileId, AddressesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

