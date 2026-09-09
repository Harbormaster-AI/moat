import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from manufacturingOnDjango.delegates.ShiftAssignmentDelegate import ShiftAssignmentDelegate

 #======================================================================
# 
# Encapsulates data for View ShiftAssignment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ShiftAssignmentView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the ShiftAssignment index.")

def get(request, shiftAssignmentId ):
	delegate = ShiftAssignmentDelegate()
	responseData = delegate.get( shiftAssignmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	shiftAssignment = json.loads(request.body)
	delegate = ShiftAssignmentDelegate()
	responseData = delegate.createFromJson( shiftAssignment )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	shiftAssignment = json.loads(request.body)
	delegate = ShiftAssignmentDelegate()
	responseData = delegate.save( shiftAssignment )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, shiftAssignmentId ):
	delegate = ShiftAssignmentDelegate()
	responseData = delegate.delete( shiftAssignmentId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ShiftAssignmentDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignShift( request, shiftAssignmentId, ShiftId ):
	delegate = ShiftAssignmentDelegate()
	responseData = delegate.saveShift( shiftAssignmentId, ShiftId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignShift( request, shiftAssignmentId ):
	delegate = ShiftAssignmentDelegate()
	responseData = delegate.deleteShift( shiftAssignmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignEmployee( request, shiftAssignmentId, EmployeeId ):
	delegate = ShiftAssignmentDelegate()
	responseData = delegate.saveEmployee( shiftAssignmentId, EmployeeId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignEmployee( request, shiftAssignmentId ):
	delegate = ShiftAssignmentDelegate()
	responseData = delegate.deleteEmployee( shiftAssignmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignWorkCenter( request, shiftAssignmentId, WorkCenterId ):
	delegate = ShiftAssignmentDelegate()
	responseData = delegate.saveWorkCenter( shiftAssignmentId, WorkCenterId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignWorkCenter( request, shiftAssignmentId ):
	delegate = ShiftAssignmentDelegate()
	responseData = delegate.deleteWorkCenter( shiftAssignmentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

