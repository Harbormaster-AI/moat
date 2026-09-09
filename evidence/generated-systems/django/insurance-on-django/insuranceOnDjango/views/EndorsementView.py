import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from insuranceOnDjango.delegates.EndorsementDelegate import EndorsementDelegate

 #======================================================================
# 
# Encapsulates data for View Endorsement
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class EndorsementView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Endorsement index.")

def get(request, endorsementId ):
	delegate = EndorsementDelegate()
	responseData = delegate.get( endorsementId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	endorsement = json.loads(request.body)
	delegate = EndorsementDelegate()
	responseData = delegate.createFromJson( endorsement )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	endorsement = json.loads(request.body)
	delegate = EndorsementDelegate()
	responseData = delegate.save( endorsement )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, endorsementId ):
	delegate = EndorsementDelegate()
	responseData = delegate.delete( endorsementId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = EndorsementDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPolicy( request, endorsementId, PolicyId ):
	delegate = EndorsementDelegate()
	responseData = delegate.savePolicy( endorsementId, PolicyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPolicy( request, endorsementId ):
	delegate = EndorsementDelegate()
	responseData = delegate.deletePolicy( endorsementId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

