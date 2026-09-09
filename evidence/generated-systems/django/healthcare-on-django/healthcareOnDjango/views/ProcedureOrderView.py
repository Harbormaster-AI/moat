import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from healthcareOnDjango.delegates.ProcedureOrderDelegate import ProcedureOrderDelegate

 #======================================================================
# 
# Encapsulates data for View ProcedureOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ProcedureOrderView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the ProcedureOrder index.")

def get(request, procedureOrderId ):
	delegate = ProcedureOrderDelegate()
	responseData = delegate.get( procedureOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	procedureOrder = json.loads(request.body)
	delegate = ProcedureOrderDelegate()
	responseData = delegate.createFromJson( procedureOrder )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	procedureOrder = json.loads(request.body)
	delegate = ProcedureOrderDelegate()
	responseData = delegate.save( procedureOrder )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, procedureOrderId ):
	delegate = ProcedureOrderDelegate()
	responseData = delegate.delete( procedureOrderId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ProcedureOrderDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrder( request, procedureOrderId, OrderId ):
	delegate = ProcedureOrderDelegate()
	responseData = delegate.saveOrder( procedureOrderId, OrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrder( request, procedureOrderId ):
	delegate = ProcedureOrderDelegate()
	responseData = delegate.deleteOrder( procedureOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignFacility( request, procedureOrderId, FacilityId ):
	delegate = ProcedureOrderDelegate()
	responseData = delegate.saveFacility( procedureOrderId, FacilityId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignFacility( request, procedureOrderId ):
	delegate = ProcedureOrderDelegate()
	responseData = delegate.deleteFacility( procedureOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignProcedure( request, procedureOrderId, ProcedureId ):
	delegate = ProcedureOrderDelegate()
	responseData = delegate.saveProcedure( procedureOrderId, ProcedureId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignProcedure( request, procedureOrderId ):
	delegate = ProcedureOrderDelegate()
	responseData = delegate.deleteProcedure( procedureOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

