import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from aerospaceOnDjango.delegates.Component_Delegate import Component_Delegate

 #======================================================================
# 
# Encapsulates data for View Component_
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Component_View function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Component_ index.")

def get(request, component_Id ):
	delegate = Component_Delegate()
	responseData = delegate.get( component_Id )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	component_ = json.loads(request.body)
	delegate = Component_Delegate()
	responseData = delegate.createFromJson( component_ )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	component_ = json.loads(request.body)
	delegate = Component_Delegate()
	responseData = delegate.save( component_ )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, component_Id ):
	delegate = Component_Delegate()
	responseData = delegate.delete( component_Id )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = Component_Delegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignSupplier( request, component_Id, SupplierId ):
	delegate = Component_Delegate()
	responseData = delegate.saveSupplier( component_Id, SupplierId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignSupplier( request, component_Id ):
	delegate = Component_Delegate()
	responseData = delegate.deleteSupplier( component_Id )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

