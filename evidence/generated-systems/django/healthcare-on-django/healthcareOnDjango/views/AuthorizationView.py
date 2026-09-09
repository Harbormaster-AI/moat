import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from healthcareOnDjango.delegates.AuthorizationDelegate import AuthorizationDelegate

 #======================================================================
# 
# Encapsulates data for View Authorization
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AuthorizationView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Authorization index.")

def get(request, authorizationId ):
	delegate = AuthorizationDelegate()
	responseData = delegate.get( authorizationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	authorization = json.loads(request.body)
	delegate = AuthorizationDelegate()
	responseData = delegate.createFromJson( authorization )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	authorization = json.loads(request.body)
	delegate = AuthorizationDelegate()
	responseData = delegate.save( authorization )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, authorizationId ):
	delegate = AuthorizationDelegate()
	responseData = delegate.delete( authorizationId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = AuthorizationDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCoverage( request, authorizationId, CoverageId ):
	delegate = AuthorizationDelegate()
	responseData = delegate.saveCoverage( authorizationId, CoverageId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCoverage( request, authorizationId ):
	delegate = AuthorizationDelegate()
	responseData = delegate.deleteCoverage( authorizationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrder( request, authorizationId, OrderId ):
	delegate = AuthorizationDelegate()
	responseData = delegate.saveOrder( authorizationId, OrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrder( request, authorizationId ):
	delegate = AuthorizationDelegate()
	responseData = delegate.deleteOrder( authorizationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

