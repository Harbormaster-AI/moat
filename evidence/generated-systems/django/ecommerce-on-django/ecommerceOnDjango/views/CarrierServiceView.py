import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from ecommerceOnDjango.delegates.CarrierServiceDelegate import CarrierServiceDelegate

 #======================================================================
# 
# Encapsulates data for View CarrierService
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CarrierServiceView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the CarrierService index.")

def get(request, carrierServiceId ):
	delegate = CarrierServiceDelegate()
	responseData = delegate.get( carrierServiceId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	carrierService = json.loads(request.body)
	delegate = CarrierServiceDelegate()
	responseData = delegate.createFromJson( carrierService )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	carrierService = json.loads(request.body)
	delegate = CarrierServiceDelegate()
	responseData = delegate.save( carrierService )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, carrierServiceId ):
	delegate = CarrierServiceDelegate()
	responseData = delegate.delete( carrierServiceId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = CarrierServiceDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addShippingMethods( request, carrierServiceId, ShippingMethodsIds ):
	delegate = CarrierServiceDelegate()
	responseData = delegate.addShippingMethods( carrierServiceId, ShippingMethodsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeShippingMethods( request, carrierServiceId, ShippingMethodsIds ):
	delegate = CarrierServiceDelegate()
	responseData = delegate.removeShippingMethods( carrierServiceId, ShippingMethodsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

