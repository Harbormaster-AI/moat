import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from governanceOnDjango.delegates.ProcedureDelegate import ProcedureDelegate

 #======================================================================
# 
# Encapsulates data for View Procedure
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ProcedureView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Procedure index.")

def get(request, procedureId ):
	delegate = ProcedureDelegate()
	responseData = delegate.get( procedureId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	procedure = json.loads(request.body)
	delegate = ProcedureDelegate()
	responseData = delegate.createFromJson( procedure )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	procedure = json.loads(request.body)
	delegate = ProcedureDelegate()
	responseData = delegate.save( procedure )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, procedureId ):
	delegate = ProcedureDelegate()
	responseData = delegate.delete( procedureId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ProcedureDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPolicy( request, procedureId, PolicyId ):
	delegate = ProcedureDelegate()
	responseData = delegate.savePolicy( procedureId, PolicyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPolicy( request, procedureId ):
	delegate = ProcedureDelegate()
	responseData = delegate.deletePolicy( procedureId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addControls( request, procedureId, ControlsIds ):
	delegate = ProcedureDelegate()
	responseData = delegate.addControls( procedureId, ControlsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeControls( request, procedureId, ControlsIds ):
	delegate = ProcedureDelegate()
	responseData = delegate.removeControls( procedureId, ControlsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

