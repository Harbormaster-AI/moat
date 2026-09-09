import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from ecommerceOnDjango.delegates.CartDelegate import CartDelegate

 #======================================================================
# 
# Encapsulates data for View Cart
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CartView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Cart index.")

def get(request, cartId ):
	delegate = CartDelegate()
	responseData = delegate.get( cartId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	cart = json.loads(request.body)
	delegate = CartDelegate()
	responseData = delegate.createFromJson( cart )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	cart = json.loads(request.body)
	delegate = CartDelegate()
	responseData = delegate.save( cart )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, cartId ):
	delegate = CartDelegate()
	responseData = delegate.delete( cartId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = CartDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCustomer( request, cartId, CustomerId ):
	delegate = CartDelegate()
	responseData = delegate.saveCustomer( cartId, CustomerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCustomer( request, cartId ):
	delegate = CartDelegate()
	responseData = delegate.deleteCustomer( cartId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignChannel( request, cartId, ChannelId ):
	delegate = CartDelegate()
	responseData = delegate.saveChannel( cartId, ChannelId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignChannel( request, cartId ):
	delegate = CartDelegate()
	responseData = delegate.deleteChannel( cartId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addItems( request, cartId, ItemsIds ):
	delegate = CartDelegate()
	responseData = delegate.addItems( cartId, ItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeItems( request, cartId, ItemsIds ):
	delegate = CartDelegate()
	responseData = delegate.removeItems( cartId, ItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addAppliedPromotions( request, cartId, AppliedPromotionsIds ):
	delegate = CartDelegate()
	responseData = delegate.addAppliedPromotions( cartId, AppliedPromotionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAppliedPromotions( request, cartId, AppliedPromotionsIds ):
	delegate = CartDelegate()
	responseData = delegate.removeAppliedPromotions( cartId, AppliedPromotionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

