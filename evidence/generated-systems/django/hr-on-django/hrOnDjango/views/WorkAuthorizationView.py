import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from hrOnDjango.delegates.WorkAuthorizationDelegate import WorkAuthorizationDelegate

 #======================================================================
# 
# Encapsulates data for View WorkAuthorization
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class WorkAuthorizationView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the WorkAuthorization index.")

def get(request, workAuthorizationId ):
	delegate = WorkAuthorizationDelegate()
	responseData = delegate.get( workAuthorizationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	workAuthorization = json.loads(request.body)
	delegate = WorkAuthorizationDelegate()
	responseData = delegate.createFromJson( workAuthorization )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	workAuthorization = json.loads(request.body)
	delegate = WorkAuthorizationDelegate()
	responseData = delegate.save( workAuthorization )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, workAuthorizationId ):
	delegate = WorkAuthorizationDelegate()
	responseData = delegate.delete( workAuthorizationId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = WorkAuthorizationDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignEmployee( request, workAuthorizationId, EmployeeId ):
	delegate = WorkAuthorizationDelegate()
	responseData = delegate.saveEmployee( workAuthorizationId, EmployeeId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignEmployee( request, workAuthorizationId ):
	delegate = WorkAuthorizationDelegate()
	responseData = delegate.deleteEmployee( workAuthorizationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addDocuments( request, workAuthorizationId, DocumentsIds ):
	delegate = WorkAuthorizationDelegate()
	responseData = delegate.addDocuments( workAuthorizationId, DocumentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeDocuments( request, workAuthorizationId, DocumentsIds ):
	delegate = WorkAuthorizationDelegate()
	responseData = delegate.removeDocuments( workAuthorizationId, DocumentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

