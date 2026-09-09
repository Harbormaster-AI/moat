import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from ecommerceOnDjango.delegates.GiftCardDelegate import GiftCardDelegate

 #======================================================================
# 
# Encapsulates data for View GiftCard
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class GiftCardView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the GiftCard index.")

def get(request, giftCardId ):
	delegate = GiftCardDelegate()
	responseData = delegate.get( giftCardId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	giftCard = json.loads(request.body)
	delegate = GiftCardDelegate()
	responseData = delegate.createFromJson( giftCard )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	giftCard = json.loads(request.body)
	delegate = GiftCardDelegate()
	responseData = delegate.save( giftCard )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, giftCardId ):
	delegate = GiftCardDelegate()
	responseData = delegate.delete( giftCardId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = GiftCardDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCustomer( request, giftCardId, CustomerId ):
	delegate = GiftCardDelegate()
	responseData = delegate.saveCustomer( giftCardId, CustomerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCustomer( request, giftCardId ):
	delegate = GiftCardDelegate()
	responseData = delegate.deleteCustomer( giftCardId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignIssuedOrder( request, giftCardId, IssuedOrderId ):
	delegate = GiftCardDelegate()
	responseData = delegate.saveIssuedOrder( giftCardId, IssuedOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignIssuedOrder( request, giftCardId ):
	delegate = GiftCardDelegate()
	responseData = delegate.deleteIssuedOrder( giftCardId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addRedemptions( request, giftCardId, RedemptionsIds ):
	delegate = GiftCardDelegate()
	responseData = delegate.addRedemptions( giftCardId, RedemptionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeRedemptions( request, giftCardId, RedemptionsIds ):
	delegate = GiftCardDelegate()
	responseData = delegate.removeRedemptions( giftCardId, RedemptionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

