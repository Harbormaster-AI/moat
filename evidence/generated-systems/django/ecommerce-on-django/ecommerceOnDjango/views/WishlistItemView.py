import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from ecommerceOnDjango.delegates.WishlistItemDelegate import WishlistItemDelegate

 #======================================================================
# 
# Encapsulates data for View WishlistItem
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class WishlistItemView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the WishlistItem index.")

def get(request, wishlistItemId ):
	delegate = WishlistItemDelegate()
	responseData = delegate.get( wishlistItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	wishlistItem = json.loads(request.body)
	delegate = WishlistItemDelegate()
	responseData = delegate.createFromJson( wishlistItem )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	wishlistItem = json.loads(request.body)
	delegate = WishlistItemDelegate()
	responseData = delegate.save( wishlistItem )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, wishlistItemId ):
	delegate = WishlistItemDelegate()
	responseData = delegate.delete( wishlistItemId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = WishlistItemDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignWishlist( request, wishlistItemId, WishlistId ):
	delegate = WishlistItemDelegate()
	responseData = delegate.saveWishlist( wishlistItemId, WishlistId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignWishlist( request, wishlistItemId ):
	delegate = WishlistItemDelegate()
	responseData = delegate.deleteWishlist( wishlistItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignVariant( request, wishlistItemId, VariantId ):
	delegate = WishlistItemDelegate()
	responseData = delegate.saveVariant( wishlistItemId, VariantId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignVariant( request, wishlistItemId ):
	delegate = WishlistItemDelegate()
	responseData = delegate.deleteVariant( wishlistItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

