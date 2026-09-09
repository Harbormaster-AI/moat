import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from fintechOnDjango.delegates.VerifiedAddressDelegate import VerifiedAddressDelegate

 #======================================================================
# 
# Encapsulates data for View VerifiedAddress
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class VerifiedAddressView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the VerifiedAddress index.")

def get(request, verifiedAddressId ):
	delegate = VerifiedAddressDelegate()
	responseData = delegate.get( verifiedAddressId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	verifiedAddress = json.loads(request.body)
	delegate = VerifiedAddressDelegate()
	responseData = delegate.createFromJson( verifiedAddress )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	verifiedAddress = json.loads(request.body)
	delegate = VerifiedAddressDelegate()
	responseData = delegate.save( verifiedAddress )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, verifiedAddressId ):
	delegate = VerifiedAddressDelegate()
	responseData = delegate.delete( verifiedAddressId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = VerifiedAddressDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignKycProfile( request, verifiedAddressId, KycProfileId ):
	delegate = VerifiedAddressDelegate()
	responseData = delegate.saveKycProfile( verifiedAddressId, KycProfileId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignKycProfile( request, verifiedAddressId ):
	delegate = VerifiedAddressDelegate()
	responseData = delegate.deleteKycProfile( verifiedAddressId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

