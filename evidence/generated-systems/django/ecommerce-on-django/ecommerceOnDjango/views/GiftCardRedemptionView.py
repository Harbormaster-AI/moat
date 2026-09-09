import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from ecommerceOnDjango.delegates.GiftCardRedemptionDelegate import GiftCardRedemptionDelegate

 #======================================================================
# 
# Encapsulates data for View GiftCardRedemption
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class GiftCardRedemptionView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the GiftCardRedemption index.")

def get(request, giftCardRedemptionId ):
	delegate = GiftCardRedemptionDelegate()
	responseData = delegate.get( giftCardRedemptionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	giftCardRedemption = json.loads(request.body)
	delegate = GiftCardRedemptionDelegate()
	responseData = delegate.createFromJson( giftCardRedemption )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	giftCardRedemption = json.loads(request.body)
	delegate = GiftCardRedemptionDelegate()
	responseData = delegate.save( giftCardRedemption )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, giftCardRedemptionId ):
	delegate = GiftCardRedemptionDelegate()
	responseData = delegate.delete( giftCardRedemptionId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = GiftCardRedemptionDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignGiftCard( request, giftCardRedemptionId, GiftCardId ):
	delegate = GiftCardRedemptionDelegate()
	responseData = delegate.saveGiftCard( giftCardRedemptionId, GiftCardId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignGiftCard( request, giftCardRedemptionId ):
	delegate = GiftCardRedemptionDelegate()
	responseData = delegate.deleteGiftCard( giftCardRedemptionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrder( request, giftCardRedemptionId, OrderId ):
	delegate = GiftCardRedemptionDelegate()
	responseData = delegate.saveOrder( giftCardRedemptionId, OrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrder( request, giftCardRedemptionId ):
	delegate = GiftCardRedemptionDelegate()
	responseData = delegate.deleteOrder( giftCardRedemptionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

