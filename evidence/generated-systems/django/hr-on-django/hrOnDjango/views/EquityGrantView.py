import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from hrOnDjango.delegates.EquityGrantDelegate import EquityGrantDelegate

 #======================================================================
# 
# Encapsulates data for View EquityGrant
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class EquityGrantView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the EquityGrant index.")

def get(request, equityGrantId ):
	delegate = EquityGrantDelegate()
	responseData = delegate.get( equityGrantId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	equityGrant = json.loads(request.body)
	delegate = EquityGrantDelegate()
	responseData = delegate.createFromJson( equityGrant )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	equityGrant = json.loads(request.body)
	delegate = EquityGrantDelegate()
	responseData = delegate.save( equityGrant )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, equityGrantId ):
	delegate = EquityGrantDelegate()
	responseData = delegate.delete( equityGrantId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = EquityGrantDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCompensationPackage( request, equityGrantId, CompensationPackageId ):
	delegate = EquityGrantDelegate()
	responseData = delegate.saveCompensationPackage( equityGrantId, CompensationPackageId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCompensationPackage( request, equityGrantId ):
	delegate = EquityGrantDelegate()
	responseData = delegate.deleteCompensationPackage( equityGrantId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

