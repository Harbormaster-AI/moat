import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from hrOnDjango.delegates.PolicyDelegate import PolicyDelegate

 #======================================================================
# 
# Encapsulates data for View Policy
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PolicyView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Policy index.")

def get(request, policyId ):
	delegate = PolicyDelegate()
	responseData = delegate.get( policyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	policy = json.loads(request.body)
	delegate = PolicyDelegate()
	responseData = delegate.createFromJson( policy )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	policy = json.loads(request.body)
	delegate = PolicyDelegate()
	responseData = delegate.save( policy )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, policyId ):
	delegate = PolicyDelegate()
	responseData = delegate.delete( policyId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = PolicyDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrganization( request, policyId, OrganizationId ):
	delegate = PolicyDelegate()
	responseData = delegate.saveOrganization( policyId, OrganizationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrganization( request, policyId ):
	delegate = PolicyDelegate()
	responseData = delegate.deleteOrganization( policyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addAcknowledgements( request, policyId, AcknowledgementsIds ):
	delegate = PolicyDelegate()
	responseData = delegate.addAcknowledgements( policyId, AcknowledgementsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAcknowledgements( request, policyId, AcknowledgementsIds ):
	delegate = PolicyDelegate()
	responseData = delegate.removeAcknowledgements( policyId, AcknowledgementsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

