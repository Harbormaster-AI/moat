import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from manufacturingOnDjango.delegates.EmployeeDelegate import EmployeeDelegate

 #======================================================================
# 
# Encapsulates data for View Employee
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class EmployeeView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Employee index.")

def get(request, employeeId ):
	delegate = EmployeeDelegate()
	responseData = delegate.get( employeeId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	employee = json.loads(request.body)
	delegate = EmployeeDelegate()
	responseData = delegate.createFromJson( employee )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	employee = json.loads(request.body)
	delegate = EmployeeDelegate()
	responseData = delegate.save( employee )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, employeeId ):
	delegate = EmployeeDelegate()
	responseData = delegate.delete( employeeId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = EmployeeDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignWorkCenter( request, employeeId, WorkCenterId ):
	delegate = EmployeeDelegate()
	responseData = delegate.saveWorkCenter( employeeId, WorkCenterId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignWorkCenter( request, employeeId ):
	delegate = EmployeeDelegate()
	responseData = delegate.deleteWorkCenter( employeeId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addShiftAssignments( request, employeeId, ShiftAssignmentsIds ):
	delegate = EmployeeDelegate()
	responseData = delegate.addShiftAssignments( employeeId, ShiftAssignmentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeShiftAssignments( request, employeeId, ShiftAssignmentsIds ):
	delegate = EmployeeDelegate()
	responseData = delegate.removeShiftAssignments( employeeId, ShiftAssignmentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCorrectiveActions( request, employeeId, CorrectiveActionsIds ):
	delegate = EmployeeDelegate()
	responseData = delegate.addCorrectiveActions( employeeId, CorrectiveActionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCorrectiveActions( request, employeeId, CorrectiveActionsIds ):
	delegate = EmployeeDelegate()
	responseData = delegate.removeCorrectiveActions( employeeId, CorrectiveActionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

