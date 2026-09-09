import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from hrOnDjango.delegates.PolicyAcknowledgementDelegate import PolicyAcknowledgementDelegate

 #======================================================================
# 
# Encapsulates data for View PolicyAcknowledgement
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PolicyAcknowledgementView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the PolicyAcknowledgement index.")

def get(request, policyAcknowledgementId ):
	delegate = PolicyAcknowledgementDelegate()
	responseData = delegate.get( policyAcknowledgementId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	policyAcknowledgement = json.loads(request.body)
	delegate = PolicyAcknowledgementDelegate()
	responseData = delegate.createFromJson( policyAcknowledgement )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	policyAcknowledgement = json.loads(request.body)
	delegate = PolicyAcknowledgementDelegate()
	responseData = delegate.save( policyAcknowledgement )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, policyAcknowledgementId ):
	delegate = PolicyAcknowledgementDelegate()
	responseData = delegate.delete( policyAcknowledgementId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = PolicyAcknowledgementDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPolicy( request, policyAcknowledgementId, PolicyId ):
	delegate = PolicyAcknowledgementDelegate()
	responseData = delegate.savePolicy( policyAcknowledgementId, PolicyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPolicy( request, policyAcknowledgementId ):
	delegate = PolicyAcknowledgementDelegate()
	responseData = delegate.deletePolicy( policyAcknowledgementId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignEmployee( request, policyAcknowledgementId, EmployeeId ):
	delegate = PolicyAcknowledgementDelegate()
	responseData = delegate.saveEmployee( policyAcknowledgementId, EmployeeId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignEmployee( request, policyAcknowledgementId ):
	delegate = PolicyAcknowledgementDelegate()
	responseData = delegate.deleteEmployee( policyAcknowledgementId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

