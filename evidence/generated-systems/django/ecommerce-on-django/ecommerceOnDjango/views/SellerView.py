import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from ecommerceOnDjango.delegates.SellerDelegate import SellerDelegate

 #======================================================================
# 
# Encapsulates data for View Seller
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SellerView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Seller index.")

def get(request, sellerId ):
	delegate = SellerDelegate()
	responseData = delegate.get( sellerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	seller = json.loads(request.body)
	delegate = SellerDelegate()
	responseData = delegate.createFromJson( seller )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	seller = json.loads(request.body)
	delegate = SellerDelegate()
	responseData = delegate.save( seller )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, sellerId ):
	delegate = SellerDelegate()
	responseData = delegate.delete( sellerId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = SellerDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignMerchant( request, sellerId, MerchantId ):
	delegate = SellerDelegate()
	responseData = delegate.saveMerchant( sellerId, MerchantId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignMerchant( request, sellerId ):
	delegate = SellerDelegate()
	responseData = delegate.deleteMerchant( sellerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addProducts( request, sellerId, ProductsIds ):
	delegate = SellerDelegate()
	responseData = delegate.addProducts( sellerId, ProductsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeProducts( request, sellerId, ProductsIds ):
	delegate = SellerDelegate()
	responseData = delegate.removeProducts( sellerId, ProductsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPayouts( request, sellerId, PayoutsIds ):
	delegate = SellerDelegate()
	responseData = delegate.addPayouts( sellerId, PayoutsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePayouts( request, sellerId, PayoutsIds ):
	delegate = SellerDelegate()
	responseData = delegate.removePayouts( sellerId, PayoutsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addOrders( request, sellerId, OrdersIds ):
	delegate = SellerDelegate()
	responseData = delegate.addOrders( sellerId, OrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeOrders( request, sellerId, OrdersIds ):
	delegate = SellerDelegate()
	responseData = delegate.removeOrders( sellerId, OrdersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

