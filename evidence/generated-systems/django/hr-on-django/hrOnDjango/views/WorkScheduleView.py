import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from hrOnDjango.delegates.WorkScheduleDelegate import WorkScheduleDelegate

 #======================================================================
# 
# Encapsulates data for View WorkSchedule
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class WorkScheduleView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the WorkSchedule index.")

def get(request, workScheduleId ):
	delegate = WorkScheduleDelegate()
	responseData = delegate.get( workScheduleId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	workSchedule = json.loads(request.body)
	delegate = WorkScheduleDelegate()
	responseData = delegate.createFromJson( workSchedule )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	workSchedule = json.loads(request.body)
	delegate = WorkScheduleDelegate()
	responseData = delegate.save( workSchedule )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, workScheduleId ):
	delegate = WorkScheduleDelegate()
	responseData = delegate.delete( workScheduleId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = WorkScheduleDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addContracts( request, workScheduleId, ContractsIds ):
	delegate = WorkScheduleDelegate()
	responseData = delegate.addContracts( workScheduleId, ContractsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeContracts( request, workScheduleId, ContractsIds ):
	delegate = WorkScheduleDelegate()
	responseData = delegate.removeContracts( workScheduleId, ContractsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addShifts( request, workScheduleId, ShiftsIds ):
	delegate = WorkScheduleDelegate()
	responseData = delegate.addShifts( workScheduleId, ShiftsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeShifts( request, workScheduleId, ShiftsIds ):
	delegate = WorkScheduleDelegate()
	responseData = delegate.removeShifts( workScheduleId, ShiftsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addExceptions( request, workScheduleId, ExceptionsIds ):
	delegate = WorkScheduleDelegate()
	responseData = delegate.addExceptions( workScheduleId, ExceptionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeExceptions( request, workScheduleId, ExceptionsIds ):
	delegate = WorkScheduleDelegate()
	responseData = delegate.removeExceptions( workScheduleId, ExceptionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

