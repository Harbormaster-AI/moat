import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from aerospaceOnDjango.delegates.APUDelegate import APUDelegate

 #======================================================================
# 
# Encapsulates data for View APU
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class APUView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the APU index.")

def get(request, aPUId ):
	delegate = APUDelegate()
	responseData = delegate.get( aPUId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	aPU = json.loads(request.body)
	delegate = APUDelegate()
	responseData = delegate.createFromJson( aPU )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	aPU = json.loads(request.body)
	delegate = APUDelegate()
	responseData = delegate.save( aPU )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, aPUId ):
	delegate = APUDelegate()
	responseData = delegate.delete( aPUId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = APUDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignSupplier( request, aPUId, SupplierId ):
	delegate = APUDelegate()
	responseData = delegate.saveSupplier( aPUId, SupplierId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignSupplier( request, aPUId ):
	delegate = APUDelegate()
	responseData = delegate.deleteSupplier( aPUId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addVariants( request, aPUId, VariantsIds ):
	delegate = APUDelegate()
	responseData = delegate.addVariants( aPUId, VariantsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeVariants( request, aPUId, VariantsIds ):
	delegate = APUDelegate()
	responseData = delegate.removeVariants( aPUId, VariantsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

