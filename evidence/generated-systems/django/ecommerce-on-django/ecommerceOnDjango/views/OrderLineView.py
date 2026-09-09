import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from ecommerceOnDjango.delegates.OrderLineDelegate import OrderLineDelegate

 #======================================================================
# 
# Encapsulates data for View OrderLine
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OrderLineView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the OrderLine index.")

def get(request, orderLineId ):
	delegate = OrderLineDelegate()
	responseData = delegate.get( orderLineId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	orderLine = json.loads(request.body)
	delegate = OrderLineDelegate()
	responseData = delegate.createFromJson( orderLine )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	orderLine = json.loads(request.body)
	delegate = OrderLineDelegate()
	responseData = delegate.save( orderLine )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, orderLineId ):
	delegate = OrderLineDelegate()
	responseData = delegate.delete( orderLineId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = OrderLineDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrder( request, orderLineId, OrderId ):
	delegate = OrderLineDelegate()
	responseData = delegate.saveOrder( orderLineId, OrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrder( request, orderLineId ):
	delegate = OrderLineDelegate()
	responseData = delegate.deleteOrder( orderLineId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignVariant( request, orderLineId, VariantId ):
	delegate = OrderLineDelegate()
	responseData = delegate.saveVariant( orderLineId, VariantId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignVariant( request, orderLineId ):
	delegate = OrderLineDelegate()
	responseData = delegate.deleteVariant( orderLineId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addAppliedPromotions( request, orderLineId, AppliedPromotionsIds ):
	delegate = OrderLineDelegate()
	responseData = delegate.addAppliedPromotions( orderLineId, AppliedPromotionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAppliedPromotions( request, orderLineId, AppliedPromotionsIds ):
	delegate = OrderLineDelegate()
	responseData = delegate.removeAppliedPromotions( orderLineId, AppliedPromotionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

