import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from ecommerceOnDjango.delegates.ProductPricingDelegate import ProductPricingDelegate

 #======================================================================
# 
# Encapsulates data for View ProductPricing
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ProductPricingView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the ProductPricing index.")

def get(request, productPricingId ):
	delegate = ProductPricingDelegate()
	responseData = delegate.get( productPricingId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	productPricing = json.loads(request.body)
	delegate = ProductPricingDelegate()
	responseData = delegate.createFromJson( productPricing )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	productPricing = json.loads(request.body)
	delegate = ProductPricingDelegate()
	responseData = delegate.save( productPricing )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, productPricingId ):
	delegate = ProductPricingDelegate()
	responseData = delegate.delete( productPricingId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ProductPricingDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignVariant( request, productPricingId, VariantId ):
	delegate = ProductPricingDelegate()
	responseData = delegate.saveVariant( productPricingId, VariantId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignVariant( request, productPricingId ):
	delegate = ProductPricingDelegate()
	responseData = delegate.deleteVariant( productPricingId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignChannel( request, productPricingId, ChannelId ):
	delegate = ProductPricingDelegate()
	responseData = delegate.saveChannel( productPricingId, ChannelId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignChannel( request, productPricingId ):
	delegate = ProductPricingDelegate()
	responseData = delegate.deleteChannel( productPricingId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

