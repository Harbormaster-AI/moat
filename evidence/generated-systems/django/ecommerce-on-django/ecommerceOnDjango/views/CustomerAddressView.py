import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from ecommerceOnDjango.delegates.CustomerAddressDelegate import CustomerAddressDelegate

 #======================================================================
# 
# Encapsulates data for View CustomerAddress
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CustomerAddressView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the CustomerAddress index.")

def get(request, customerAddressId ):
	delegate = CustomerAddressDelegate()
	responseData = delegate.get( customerAddressId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	customerAddress = json.loads(request.body)
	delegate = CustomerAddressDelegate()
	responseData = delegate.createFromJson( customerAddress )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	customerAddress = json.loads(request.body)
	delegate = CustomerAddressDelegate()
	responseData = delegate.save( customerAddress )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, customerAddressId ):
	delegate = CustomerAddressDelegate()
	responseData = delegate.delete( customerAddressId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = CustomerAddressDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCustomer( request, customerAddressId, CustomerId ):
	delegate = CustomerAddressDelegate()
	responseData = delegate.saveCustomer( customerAddressId, CustomerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCustomer( request, customerAddressId ):
	delegate = CustomerAddressDelegate()
	responseData = delegate.deleteCustomer( customerAddressId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

