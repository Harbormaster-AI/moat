import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from ecommerceOnDjango.delegates.WishlistDelegate import WishlistDelegate

 #======================================================================
# 
# Encapsulates data for View Wishlist
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class WishlistView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Wishlist index.")

def get(request, wishlistId ):
	delegate = WishlistDelegate()
	responseData = delegate.get( wishlistId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	wishlist = json.loads(request.body)
	delegate = WishlistDelegate()
	responseData = delegate.createFromJson( wishlist )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	wishlist = json.loads(request.body)
	delegate = WishlistDelegate()
	responseData = delegate.save( wishlist )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, wishlistId ):
	delegate = WishlistDelegate()
	responseData = delegate.delete( wishlistId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = WishlistDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCustomer( request, wishlistId, CustomerId ):
	delegate = WishlistDelegate()
	responseData = delegate.saveCustomer( wishlistId, CustomerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCustomer( request, wishlistId ):
	delegate = WishlistDelegate()
	responseData = delegate.deleteCustomer( wishlistId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addItems( request, wishlistId, ItemsIds ):
	delegate = WishlistDelegate()
	responseData = delegate.addItems( wishlistId, ItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeItems( request, wishlistId, ItemsIds ):
	delegate = WishlistDelegate()
	responseData = delegate.removeItems( wishlistId, ItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

