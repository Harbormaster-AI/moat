import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from fintechOnDjango.delegates.KYCDocumentDelegate import KYCDocumentDelegate

 #======================================================================
# 
# Encapsulates data for View KYCDocument
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class KYCDocumentView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the KYCDocument index.")

def get(request, kYCDocumentId ):
	delegate = KYCDocumentDelegate()
	responseData = delegate.get( kYCDocumentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	kYCDocument = json.loads(request.body)
	delegate = KYCDocumentDelegate()
	responseData = delegate.createFromJson( kYCDocument )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	kYCDocument = json.loads(request.body)
	delegate = KYCDocumentDelegate()
	responseData = delegate.save( kYCDocument )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, kYCDocumentId ):
	delegate = KYCDocumentDelegate()
	responseData = delegate.delete( kYCDocumentId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = KYCDocumentDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignKycProfile( request, kYCDocumentId, KycProfileId ):
	delegate = KYCDocumentDelegate()
	responseData = delegate.saveKycProfile( kYCDocumentId, KycProfileId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignKycProfile( request, kYCDocumentId ):
	delegate = KYCDocumentDelegate()
	responseData = delegate.deleteKycProfile( kYCDocumentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

