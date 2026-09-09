import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from ecommerceOnDjango.delegates.ChannelDelegate import ChannelDelegate

 #======================================================================
# 
# Encapsulates data for View Channel
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ChannelView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Channel index.")

def get(request, channelId ):
	delegate = ChannelDelegate()
	responseData = delegate.get( channelId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	channel = json.loads(request.body)
	delegate = ChannelDelegate()
	responseData = delegate.createFromJson( channel )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	channel = json.loads(request.body)
	delegate = ChannelDelegate()
	responseData = delegate.save( channel )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, channelId ):
	delegate = ChannelDelegate()
	responseData = delegate.delete( channelId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = ChannelDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignMerchant( request, channelId, MerchantId ):
	delegate = ChannelDelegate()
	responseData = delegate.saveMerchant( channelId, MerchantId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignMerchant( request, channelId ):
	delegate = ChannelDelegate()
	responseData = delegate.deleteMerchant( channelId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCatalogs( request, channelId, CatalogsIds ):
	delegate = ChannelDelegate()
	responseData = delegate.addCatalogs( channelId, CatalogsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCatalogs( request, channelId, CatalogsIds ):
	delegate = ChannelDelegate()
	responseData = delegate.removeCatalogs( channelId, CatalogsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPromotions( request, channelId, PromotionsIds ):
	delegate = ChannelDelegate()
	responseData = delegate.addPromotions( channelId, PromotionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePromotions( request, channelId, PromotionsIds ):
	delegate = ChannelDelegate()
	responseData = delegate.removePromotions( channelId, PromotionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addShippingMethods( request, channelId, ShippingMethodsIds ):
	delegate = ChannelDelegate()
	responseData = delegate.addShippingMethods( channelId, ShippingMethodsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeShippingMethods( request, channelId, ShippingMethodsIds ):
	delegate = ChannelDelegate()
	responseData = delegate.removeShippingMethods( channelId, ShippingMethodsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPaymentProviders( request, channelId, PaymentProvidersIds ):
	delegate = ChannelDelegate()
	responseData = delegate.addPaymentProviders( channelId, PaymentProvidersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePaymentProviders( request, channelId, PaymentProvidersIds ):
	delegate = ChannelDelegate()
	responseData = delegate.removePaymentProviders( channelId, PaymentProvidersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

