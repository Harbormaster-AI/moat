import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from ecommerceOnDjango.delegates.CartItemDelegate import CartItemDelegate

 #======================================================================
# 
# Encapsulates data for View CartItem
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CartItemView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the CartItem index.")

def get(request, cartItemId ):
	delegate = CartItemDelegate()
	responseData = delegate.get( cartItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	cartItem = json.loads(request.body)
	delegate = CartItemDelegate()
	responseData = delegate.createFromJson( cartItem )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	cartItem = json.loads(request.body)
	delegate = CartItemDelegate()
	responseData = delegate.save( cartItem )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, cartItemId ):
	delegate = CartItemDelegate()
	responseData = delegate.delete( cartItemId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = CartItemDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCart( request, cartItemId, CartId ):
	delegate = CartItemDelegate()
	responseData = delegate.saveCart( cartItemId, CartId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCart( request, cartItemId ):
	delegate = CartItemDelegate()
	responseData = delegate.deleteCart( cartItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignVariant( request, cartItemId, VariantId ):
	delegate = CartItemDelegate()
	responseData = delegate.saveVariant( cartItemId, VariantId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignVariant( request, cartItemId ):
	delegate = CartItemDelegate()
	responseData = delegate.deleteVariant( cartItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addAppliedPromotions( request, cartItemId, AppliedPromotionsIds ):
	delegate = CartItemDelegate()
	responseData = delegate.addAppliedPromotions( cartItemId, AppliedPromotionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAppliedPromotions( request, cartItemId, AppliedPromotionsIds ):
	delegate = CartItemDelegate()
	responseData = delegate.removeAppliedPromotions( cartItemId, AppliedPromotionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

