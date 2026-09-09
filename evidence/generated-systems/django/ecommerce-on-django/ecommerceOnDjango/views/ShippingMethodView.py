import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from ecommerceOnDjango.delegates.ShippingMethodDelegate import ShippingMethodDelegate

 #======================================================================
# 
# Encapsulates data for View ShippingMethod
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ShippingMethodView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the ShippingMethod index.")

def get(request, shippingMethodId ):
	delegate = ShippingMethodDelegate()
	responseData = delegate.get( shippingMethodId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	shippingMethod = json.loads(request.body)
	delegate = ShippingMethodDelegate()
	responseData = delegate.createFromJson( shippingMethod )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	shippingMethod = json.loads(request.body)
	delegate = ShippingMethodDelegate()
	responseData = delegate.save( shippingMethod )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, shippingMethodId ):
	delegate = ShippingMethodDelegate()
	responseData = delegate.delete( shippingMethodId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ShippingMethodDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCarrierService( request, shippingMethodId, CarrierServiceId ):
	delegate = ShippingMethodDelegate()
	responseData = delegate.saveCarrierService( shippingMethodId, CarrierServiceId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCarrierService( request, shippingMethodId ):
	delegate = ShippingMethodDelegate()
	responseData = delegate.deleteCarrierService( shippingMethodId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addChannels( request, shippingMethodId, ChannelsIds ):
	delegate = ShippingMethodDelegate()
	responseData = delegate.addChannels( shippingMethodId, ChannelsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeChannels( request, shippingMethodId, ChannelsIds ):
	delegate = ShippingMethodDelegate()
	responseData = delegate.removeChannels( shippingMethodId, ChannelsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

