import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from hrOnDjango.delegates.WorkShiftDelegate import WorkShiftDelegate

 #======================================================================
# 
# Encapsulates data for View WorkShift
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class WorkShiftView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the WorkShift index.")

def get(request, workShiftId ):
	delegate = WorkShiftDelegate()
	responseData = delegate.get( workShiftId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	workShift = json.loads(request.body)
	delegate = WorkShiftDelegate()
	responseData = delegate.createFromJson( workShift )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	workShift = json.loads(request.body)
	delegate = WorkShiftDelegate()
	responseData = delegate.save( workShift )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, workShiftId ):
	delegate = WorkShiftDelegate()
	responseData = delegate.delete( workShiftId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = WorkShiftDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignWorkSchedule( request, workShiftId, WorkScheduleId ):
	delegate = WorkShiftDelegate()
	responseData = delegate.saveWorkSchedule( workShiftId, WorkScheduleId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignWorkSchedule( request, workShiftId ):
	delegate = WorkShiftDelegate()
	responseData = delegate.deleteWorkSchedule( workShiftId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

