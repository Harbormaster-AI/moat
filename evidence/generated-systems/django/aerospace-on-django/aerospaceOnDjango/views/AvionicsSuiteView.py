import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from aerospaceOnDjango.delegates.AvionicsSuiteDelegate import AvionicsSuiteDelegate

 #======================================================================
# 
# Encapsulates data for View AvionicsSuite
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AvionicsSuiteView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the AvionicsSuite index.")

def get(request, avionicsSuiteId ):
	delegate = AvionicsSuiteDelegate()
	responseData = delegate.get( avionicsSuiteId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	avionicsSuite = json.loads(request.body)
	delegate = AvionicsSuiteDelegate()
	responseData = delegate.createFromJson( avionicsSuite )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	avionicsSuite = json.loads(request.body)
	delegate = AvionicsSuiteDelegate()
	responseData = delegate.save( avionicsSuite )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, avionicsSuiteId ):
	delegate = AvionicsSuiteDelegate()
	responseData = delegate.delete( avionicsSuiteId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = AvionicsSuiteDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignSupplier( request, avionicsSuiteId, SupplierId ):
	delegate = AvionicsSuiteDelegate()
	responseData = delegate.saveSupplier( avionicsSuiteId, SupplierId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignSupplier( request, avionicsSuiteId ):
	delegate = AvionicsSuiteDelegate()
	responseData = delegate.deleteSupplier( avionicsSuiteId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addVariants( request, avionicsSuiteId, VariantsIds ):
	delegate = AvionicsSuiteDelegate()
	responseData = delegate.addVariants( avionicsSuiteId, VariantsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeVariants( request, avionicsSuiteId, VariantsIds ):
	delegate = AvionicsSuiteDelegate()
	responseData = delegate.removeVariants( avionicsSuiteId, VariantsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addSoftwareLoads( request, avionicsSuiteId, SoftwareLoadsIds ):
	delegate = AvionicsSuiteDelegate()
	responseData = delegate.addSoftwareLoads( avionicsSuiteId, SoftwareLoadsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeSoftwareLoads( request, avionicsSuiteId, SoftwareLoadsIds ):
	delegate = AvionicsSuiteDelegate()
	responseData = delegate.removeSoftwareLoads( avionicsSuiteId, SoftwareLoadsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

