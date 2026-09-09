import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from insuranceOnDjango.delegates.DocumentDelegate import DocumentDelegate

 #======================================================================
# 
# Encapsulates data for View Document
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DocumentView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Document index.")

def get(request, documentId ):
	delegate = DocumentDelegate()
	responseData = delegate.get( documentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	document = json.loads(request.body)
	delegate = DocumentDelegate()
	responseData = delegate.createFromJson( document )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	document = json.loads(request.body)
	delegate = DocumentDelegate()
	responseData = delegate.save( document )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, documentId ):
	delegate = DocumentDelegate()
	responseData = delegate.delete( documentId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = DocumentDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPolicy( request, documentId, PolicyId ):
	delegate = DocumentDelegate()
	responseData = delegate.savePolicy( documentId, PolicyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPolicy( request, documentId ):
	delegate = DocumentDelegate()
	responseData = delegate.deletePolicy( documentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignClaim( request, documentId, ClaimId ):
	delegate = DocumentDelegate()
	responseData = delegate.saveClaim( documentId, ClaimId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignClaim( request, documentId ):
	delegate = DocumentDelegate()
	responseData = delegate.deleteClaim( documentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignApplication( request, documentId, ApplicationId ):
	delegate = DocumentDelegate()
	responseData = delegate.saveApplication( documentId, ApplicationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignApplication( request, documentId ):
	delegate = DocumentDelegate()
	responseData = delegate.deleteApplication( documentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCustomer( request, documentId, CustomerId ):
	delegate = DocumentDelegate()
	responseData = delegate.saveCustomer( documentId, CustomerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCustomer( request, documentId ):
	delegate = DocumentDelegate()
	responseData = delegate.deleteCustomer( documentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

