import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from inventoryOnDjango.delegates.UoMConversionDelegate import UoMConversionDelegate

 #======================================================================
# 
# Encapsulates data for View UoMConversion
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class UoMConversionView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the UoMConversion index.")

def get(request, uoMConversionId ):
	delegate = UoMConversionDelegate()
	responseData = delegate.get( uoMConversionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	uoMConversion = json.loads(request.body)
	delegate = UoMConversionDelegate()
	responseData = delegate.createFromJson( uoMConversion )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	uoMConversion = json.loads(request.body)
	delegate = UoMConversionDelegate()
	responseData = delegate.save( uoMConversion )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, uoMConversionId ):
	delegate = UoMConversionDelegate()
	responseData = delegate.delete( uoMConversionId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = UoMConversionDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignSku( request, uoMConversionId, SkuId ):
	delegate = UoMConversionDelegate()
	responseData = delegate.saveSku( uoMConversionId, SkuId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignSku( request, uoMConversionId ):
	delegate = UoMConversionDelegate()
	responseData = delegate.deleteSku( uoMConversionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

