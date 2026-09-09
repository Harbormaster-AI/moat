import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from hrOnDjango.delegates.DocumentDelegate import DocumentDelegate

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

def assignCandidate( request, documentId, CandidateId ):
	delegate = DocumentDelegate()
	responseData = delegate.saveCandidate( documentId, CandidateId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCandidate( request, documentId ):
	delegate = DocumentDelegate()
	responseData = delegate.deleteCandidate( documentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignEmployee( request, documentId, EmployeeId ):
	delegate = DocumentDelegate()
	responseData = delegate.saveEmployee( documentId, EmployeeId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignEmployee( request, documentId ):
	delegate = DocumentDelegate()
	responseData = delegate.deleteEmployee( documentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

