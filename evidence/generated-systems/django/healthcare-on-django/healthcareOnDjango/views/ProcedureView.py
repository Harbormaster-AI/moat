import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from healthcareOnDjango.delegates.ProcedureDelegate import ProcedureDelegate

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

def assignEncounter( request, procedureId, EncounterId ):
	delegate = ProcedureDelegate()
	responseData = delegate.saveEncounter( procedureId, EncounterId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignEncounter( request, procedureId ):
	delegate = ProcedureDelegate()
	responseData = delegate.deleteEncounter( procedureId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPerformer( request, procedureId, PerformerId ):
	delegate = ProcedureDelegate()
	responseData = delegate.savePerformer( procedureId, PerformerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPerformer( request, procedureId ):
	delegate = ProcedureDelegate()
	responseData = delegate.deletePerformer( procedureId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignProcedureOrder( request, procedureId, ProcedureOrderId ):
	delegate = ProcedureDelegate()
	responseData = delegate.saveProcedureOrder( procedureId, ProcedureOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignProcedureOrder( request, procedureId ):
	delegate = ProcedureDelegate()
	responseData = delegate.deleteProcedureOrder( procedureId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

