import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from manufacturingOnDjango.delegates.ShiftDelegate import ShiftDelegate

 #======================================================================
# 
# Encapsulates data for View Shift
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ShiftView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Shift index.")

def get(request, shiftId ):
	delegate = ShiftDelegate()
	responseData = delegate.get( shiftId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	shift = json.loads(request.body)
	delegate = ShiftDelegate()
	responseData = delegate.createFromJson( shift )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	shift = json.loads(request.body)
	delegate = ShiftDelegate()
	responseData = delegate.save( shift )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, shiftId ):
	delegate = ShiftDelegate()
	responseData = delegate.delete( shiftId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ShiftDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPlant( request, shiftId, PlantId ):
	delegate = ShiftDelegate()
	responseData = delegate.savePlant( shiftId, PlantId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPlant( request, shiftId ):
	delegate = ShiftDelegate()
	responseData = delegate.deletePlant( shiftId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addAssignments( request, shiftId, AssignmentsIds ):
	delegate = ShiftDelegate()
	responseData = delegate.addAssignments( shiftId, AssignmentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAssignments( request, shiftId, AssignmentsIds ):
	delegate = ShiftDelegate()
	responseData = delegate.removeAssignments( shiftId, AssignmentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

