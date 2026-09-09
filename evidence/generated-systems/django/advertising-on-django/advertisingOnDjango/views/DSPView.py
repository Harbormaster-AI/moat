import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from advertisingOnDjango.delegates.DSPDelegate import DSPDelegate

 #======================================================================
# 
# Encapsulates data for View DSP
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DSPView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the DSP index.")

def get(request, dSPId ):
	delegate = DSPDelegate()
	responseData = delegate.get( dSPId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	dSP = json.loads(request.body)
	delegate = DSPDelegate()
	responseData = delegate.createFromJson( dSP )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	dSP = json.loads(request.body)
	delegate = DSPDelegate()
	responseData = delegate.save( dSP )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, dSPId ):
	delegate = DSPDelegate()
	responseData = delegate.delete( dSPId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = DSPDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addAdAccounts( request, dSPId, AdAccountsIds ):
	delegate = DSPDelegate()
	responseData = delegate.addAdAccounts( dSPId, AdAccountsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAdAccounts( request, dSPId, AdAccountsIds ):
	delegate = DSPDelegate()
	responseData = delegate.removeAdAccounts( dSPId, AdAccountsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

