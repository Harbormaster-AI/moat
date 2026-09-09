import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from hrOnDjango.delegates.TerminationDelegate import TerminationDelegate

 #======================================================================
# 
# Encapsulates data for View Termination
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TerminationView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Termination index.")

def get(request, terminationId ):
	delegate = TerminationDelegate()
	responseData = delegate.get( terminationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	termination = json.loads(request.body)
	delegate = TerminationDelegate()
	responseData = delegate.createFromJson( termination )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	termination = json.loads(request.body)
	delegate = TerminationDelegate()
	responseData = delegate.save( termination )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, terminationId ):
	delegate = TerminationDelegate()
	responseData = delegate.delete( terminationId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = TerminationDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignEmployee( request, terminationId, EmployeeId ):
	delegate = TerminationDelegate()
	responseData = delegate.saveEmployee( terminationId, EmployeeId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignEmployee( request, terminationId ):
	delegate = TerminationDelegate()
	responseData = delegate.deleteEmployee( terminationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignAssignment( request, terminationId, AssignmentId ):
	delegate = TerminationDelegate()
	responseData = delegate.saveAssignment( terminationId, AssignmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignAssignment( request, terminationId ):
	delegate = TerminationDelegate()
	responseData = delegate.deleteAssignment( terminationId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

