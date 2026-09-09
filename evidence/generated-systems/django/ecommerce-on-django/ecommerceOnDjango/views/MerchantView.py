import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from ecommerceOnDjango.delegates.MerchantDelegate import MerchantDelegate

 #======================================================================
# 
# Encapsulates data for View Merchant
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MerchantView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Merchant index.")

def get(request, merchantId ):
	delegate = MerchantDelegate()
	responseData = delegate.get( merchantId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	merchant = json.loads(request.body)
	delegate = MerchantDelegate()
	responseData = delegate.createFromJson( merchant )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	merchant = json.loads(request.body)
	delegate = MerchantDelegate()
	responseData = delegate.save( merchant )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, merchantId ):
	delegate = MerchantDelegate()
	responseData = delegate.delete( merchantId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = MerchantDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addChannels( request, merchantId, ChannelsIds ):
	delegate = MerchantDelegate()
	responseData = delegate.addChannels( merchantId, ChannelsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeChannels( request, merchantId, ChannelsIds ):
	delegate = MerchantDelegate()
	responseData = delegate.removeChannels( merchantId, ChannelsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addBrands( request, merchantId, BrandsIds ):
	delegate = MerchantDelegate()
	responseData = delegate.addBrands( merchantId, BrandsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeBrands( request, merchantId, BrandsIds ):
	delegate = MerchantDelegate()
	responseData = delegate.removeBrands( merchantId, BrandsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addFulfillmentCenters( request, merchantId, FulfillmentCentersIds ):
	delegate = MerchantDelegate()
	responseData = delegate.addFulfillmentCenters( merchantId, FulfillmentCentersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeFulfillmentCenters( request, merchantId, FulfillmentCentersIds ):
	delegate = MerchantDelegate()
	responseData = delegate.removeFulfillmentCenters( merchantId, FulfillmentCentersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addTaxRules( request, merchantId, TaxRulesIds ):
	delegate = MerchantDelegate()
	responseData = delegate.addTaxRules( merchantId, TaxRulesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeTaxRules( request, merchantId, TaxRulesIds ):
	delegate = MerchantDelegate()
	responseData = delegate.removeTaxRules( merchantId, TaxRulesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPaymentProviders( request, merchantId, PaymentProvidersIds ):
	delegate = MerchantDelegate()
	responseData = delegate.addPaymentProviders( merchantId, PaymentProvidersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePaymentProviders( request, merchantId, PaymentProvidersIds ):
	delegate = MerchantDelegate()
	responseData = delegate.removePaymentProviders( merchantId, PaymentProvidersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addSellers( request, merchantId, SellersIds ):
	delegate = MerchantDelegate()
	responseData = delegate.addSellers( merchantId, SellersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeSellers( request, merchantId, SellersIds ):
	delegate = MerchantDelegate()
	responseData = delegate.removeSellers( merchantId, SellersIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPromotions( request, merchantId, PromotionsIds ):
	delegate = MerchantDelegate()
	responseData = delegate.addPromotions( merchantId, PromotionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePromotions( request, merchantId, PromotionsIds ):
	delegate = MerchantDelegate()
	responseData = delegate.removePromotions( merchantId, PromotionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

