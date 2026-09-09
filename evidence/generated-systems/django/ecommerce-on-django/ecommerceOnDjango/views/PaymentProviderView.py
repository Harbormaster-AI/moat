import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from ecommerceOnDjango.delegates.PaymentProviderDelegate import PaymentProviderDelegate

 #======================================================================
# 
# Encapsulates data for View PaymentProvider
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PaymentProviderView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the PaymentProvider index.")

def get(request, paymentProviderId ):
	delegate = PaymentProviderDelegate()
	responseData = delegate.get( paymentProviderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	paymentProvider = json.loads(request.body)
	delegate = PaymentProviderDelegate()
	responseData = delegate.createFromJson( paymentProvider )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	paymentProvider = json.loads(request.body)
	delegate = PaymentProviderDelegate()
	responseData = delegate.save( paymentProvider )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, paymentProviderId ):
	delegate = PaymentProviderDelegate()
	responseData = delegate.delete( paymentProviderId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = PaymentProviderDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignMerchant( request, paymentProviderId, MerchantId ):
	delegate = PaymentProviderDelegate()
	responseData = delegate.saveMerchant( paymentProviderId, MerchantId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignMerchant( request, paymentProviderId ):
	delegate = PaymentProviderDelegate()
	responseData = delegate.deleteMerchant( paymentProviderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addChannels( request, paymentProviderId, ChannelsIds ):
	delegate = PaymentProviderDelegate()
	responseData = delegate.addChannels( paymentProviderId, ChannelsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeChannels( request, paymentProviderId, ChannelsIds ):
	delegate = PaymentProviderDelegate()
	responseData = delegate.removeChannels( paymentProviderId, ChannelsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPayments( request, paymentProviderId, PaymentsIds ):
	delegate = PaymentProviderDelegate()
	responseData = delegate.addPayments( paymentProviderId, PaymentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePayments( request, paymentProviderId, PaymentsIds ):
	delegate = PaymentProviderDelegate()
	responseData = delegate.removePayments( paymentProviderId, PaymentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addSubscriptions( request, paymentProviderId, SubscriptionsIds ):
	delegate = PaymentProviderDelegate()
	responseData = delegate.addSubscriptions( paymentProviderId, SubscriptionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeSubscriptions( request, paymentProviderId, SubscriptionsIds ):
	delegate = PaymentProviderDelegate()
	responseData = delegate.removeSubscriptions( paymentProviderId, SubscriptionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

