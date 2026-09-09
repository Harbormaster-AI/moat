import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from ecommerceOnDjango.delegates.ProductVariantDelegate import ProductVariantDelegate

 #======================================================================
# 
# Encapsulates data for View ProductVariant
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ProductVariantView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the ProductVariant index.")

def get(request, productVariantId ):
	delegate = ProductVariantDelegate()
	responseData = delegate.get( productVariantId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	productVariant = json.loads(request.body)
	delegate = ProductVariantDelegate()
	responseData = delegate.createFromJson( productVariant )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	productVariant = json.loads(request.body)
	delegate = ProductVariantDelegate()
	responseData = delegate.save( productVariant )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, productVariantId ):
	delegate = ProductVariantDelegate()
	responseData = delegate.delete( productVariantId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ProductVariantDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignProduct( request, productVariantId, ProductId ):
	delegate = ProductVariantDelegate()
	responseData = delegate.saveProduct( productVariantId, ProductId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignProduct( request, productVariantId ):
	delegate = ProductVariantDelegate()
	responseData = delegate.deleteProduct( productVariantId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPricing( request, productVariantId, PricingIds ):
	delegate = ProductVariantDelegate()
	responseData = delegate.addPricing( productVariantId, PricingIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePricing( request, productVariantId, PricingIds ):
	delegate = ProductVariantDelegate()
	responseData = delegate.removePricing( productVariantId, PricingIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addInventoryItems( request, productVariantId, InventoryItemsIds ):
	delegate = ProductVariantDelegate()
	responseData = delegate.addInventoryItems( productVariantId, InventoryItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeInventoryItems( request, productVariantId, InventoryItemsIds ):
	delegate = ProductVariantDelegate()
	responseData = delegate.removeInventoryItems( productVariantId, InventoryItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addMediaAssets( request, productVariantId, MediaAssetsIds ):
	delegate = ProductVariantDelegate()
	responseData = delegate.addMediaAssets( productVariantId, MediaAssetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeMediaAssets( request, productVariantId, MediaAssetsIds ):
	delegate = ProductVariantDelegate()
	responseData = delegate.removeMediaAssets( productVariantId, MediaAssetsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addSubscriptions( request, productVariantId, SubscriptionsIds ):
	delegate = ProductVariantDelegate()
	responseData = delegate.addSubscriptions( productVariantId, SubscriptionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeSubscriptions( request, productVariantId, SubscriptionsIds ):
	delegate = ProductVariantDelegate()
	responseData = delegate.removeSubscriptions( productVariantId, SubscriptionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCartItems( request, productVariantId, CartItemsIds ):
	delegate = ProductVariantDelegate()
	responseData = delegate.addCartItems( productVariantId, CartItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCartItems( request, productVariantId, CartItemsIds ):
	delegate = ProductVariantDelegate()
	responseData = delegate.removeCartItems( productVariantId, CartItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addOrderLines( request, productVariantId, OrderLinesIds ):
	delegate = ProductVariantDelegate()
	responseData = delegate.addOrderLines( productVariantId, OrderLinesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeOrderLines( request, productVariantId, OrderLinesIds ):
	delegate = ProductVariantDelegate()
	responseData = delegate.removeOrderLines( productVariantId, OrderLinesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addWishlistItems( request, productVariantId, WishlistItemsIds ):
	delegate = ProductVariantDelegate()
	responseData = delegate.addWishlistItems( productVariantId, WishlistItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeWishlistItems( request, productVariantId, WishlistItemsIds ):
	delegate = ProductVariantDelegate()
	responseData = delegate.removeWishlistItems( productVariantId, WishlistItemsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

