import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from crmOnDjango.delegates.OrderItemDelegate import OrderItemDelegate

 #======================================================================
# 
# Encapsulates data for View OrderItem
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OrderItemView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the OrderItem index.")

def get(request, orderItemId ):
	delegate = OrderItemDelegate()
	responseData = delegate.get( orderItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	orderItem = json.loads(request.body)
	delegate = OrderItemDelegate()
	responseData = delegate.createFromJson( orderItem )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	orderItem = json.loads(request.body)
	delegate = OrderItemDelegate()
	responseData = delegate.save( orderItem )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, orderItemId ):
	delegate = OrderItemDelegate()
	responseData = delegate.delete( orderItemId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = OrderItemDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrder( request, orderItemId, OrderId ):
	delegate = OrderItemDelegate()
	responseData = delegate.saveOrder( orderItemId, OrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrder( request, orderItemId ):
	delegate = OrderItemDelegate()
	responseData = delegate.deleteOrder( orderItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignProduct( request, orderItemId, ProductId ):
	delegate = OrderItemDelegate()
	responseData = delegate.saveProduct( orderItemId, ProductId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignProduct( request, orderItemId ):
	delegate = OrderItemDelegate()
	responseData = delegate.deleteProduct( orderItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPriceBookEntry( request, orderItemId, PriceBookEntryId ):
	delegate = OrderItemDelegate()
	responseData = delegate.savePriceBookEntry( orderItemId, PriceBookEntryId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPriceBookEntry( request, orderItemId ):
	delegate = OrderItemDelegate()
	responseData = delegate.deletePriceBookEntry( orderItemId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

