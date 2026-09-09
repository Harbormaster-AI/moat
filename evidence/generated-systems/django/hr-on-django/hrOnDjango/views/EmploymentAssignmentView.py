import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from hrOnDjango.delegates.EmploymentAssignmentDelegate import EmploymentAssignmentDelegate

 #======================================================================
# 
# Encapsulates data for View EmploymentAssignment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class EmploymentAssignmentView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the EmploymentAssignment index.")

def get(request, employmentAssignmentId ):
	delegate = EmploymentAssignmentDelegate()
	responseData = delegate.get( employmentAssignmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	employmentAssignment = json.loads(request.body)
	delegate = EmploymentAssignmentDelegate()
	responseData = delegate.createFromJson( employmentAssignment )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	employmentAssignment = json.loads(request.body)
	delegate = EmploymentAssignmentDelegate()
	responseData = delegate.save( employmentAssignment )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, employmentAssignmentId ):
	delegate = EmploymentAssignmentDelegate()
	responseData = delegate.delete( employmentAssignmentId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = EmploymentAssignmentDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignEmployee( request, employmentAssignmentId, EmployeeId ):
	delegate = EmploymentAssignmentDelegate()
	responseData = delegate.saveEmployee( employmentAssignmentId, EmployeeId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignEmployee( request, employmentAssignmentId ):
	delegate = EmploymentAssignmentDelegate()
	responseData = delegate.deleteEmployee( employmentAssignmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPosition( request, employmentAssignmentId, PositionId ):
	delegate = EmploymentAssignmentDelegate()
	responseData = delegate.savePosition( employmentAssignmentId, PositionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPosition( request, employmentAssignmentId ):
	delegate = EmploymentAssignmentDelegate()
	responseData = delegate.deletePosition( employmentAssignmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignSupervisor( request, employmentAssignmentId, SupervisorId ):
	delegate = EmploymentAssignmentDelegate()
	responseData = delegate.saveSupervisor( employmentAssignmentId, SupervisorId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignSupervisor( request, employmentAssignmentId ):
	delegate = EmploymentAssignmentDelegate()
	responseData = delegate.deleteSupervisor( employmentAssignmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

