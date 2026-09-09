import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from ecommerceOnDjango.delegates.SubscriptionDelegate import SubscriptionDelegate

 #======================================================================
# 
# Encapsulates data for View Subscription
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SubscriptionView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Subscription index.")

def get(request, subscriptionId ):
	delegate = SubscriptionDelegate()
	responseData = delegate.get( subscriptionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	subscription = json.loads(request.body)
	delegate = SubscriptionDelegate()
	responseData = delegate.createFromJson( subscription )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	subscription = json.loads(request.body)
	delegate = SubscriptionDelegate()
	responseData = delegate.save( subscription )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, subscriptionId ):
	delegate = SubscriptionDelegate()
	responseData = delegate.delete( subscriptionId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = SubscriptionDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCustomer( request, subscriptionId, CustomerId ):
	delegate = SubscriptionDelegate()
	responseData = delegate.saveCustomer( subscriptionId, CustomerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCustomer( request, subscriptionId ):
	delegate = SubscriptionDelegate()
	responseData = delegate.deleteCustomer( subscriptionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignVariant( request, subscriptionId, VariantId ):
	delegate = SubscriptionDelegate()
	responseData = delegate.saveVariant( subscriptionId, VariantId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignVariant( request, subscriptionId ):
	delegate = SubscriptionDelegate()
	responseData = delegate.deleteVariant( subscriptionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPaymentProvider( request, subscriptionId, PaymentProviderId ):
	delegate = SubscriptionDelegate()
	responseData = delegate.savePaymentProvider( subscriptionId, PaymentProviderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPaymentProvider( request, subscriptionId ):
	delegate = SubscriptionDelegate()
	responseData = delegate.deletePaymentProvider( subscriptionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignChannel( request, subscriptionId, ChannelId ):
	delegate = SubscriptionDelegate()
	responseData = delegate.saveChannel( subscriptionId, ChannelId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignChannel( request, subscriptionId ):
	delegate = SubscriptionDelegate()
	responseData = delegate.deleteChannel( subscriptionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

