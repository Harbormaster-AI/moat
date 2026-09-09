import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from insuranceOnDjango.delegates.ThirdPartyDelegate import ThirdPartyDelegate

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

def addSubrogations( request, thirdPartyId, SubrogationsIds ):
	delegate = ThirdPartyDelegate()
	responseData = delegate.addSubrogations( thirdPartyId, SubrogationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeSubrogations( request, thirdPartyId, SubrogationsIds ):
	delegate = ThirdPartyDelegate()
	responseData = delegate.removeSubrogations( thirdPartyId, SubrogationsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

