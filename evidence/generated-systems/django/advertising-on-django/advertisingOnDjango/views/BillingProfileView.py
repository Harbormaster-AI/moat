import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from advertisingOnDjango.delegates.BillingProfileDelegate import BillingProfileDelegate

 #======================================================================
# 
# Encapsulates data for View BillingProfile
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BillingProfileView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the BillingProfile index.")

def get(request, billingProfileId ):
	delegate = BillingProfileDelegate()
	responseData = delegate.get( billingProfileId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	billingProfile = json.loads(request.body)
	delegate = BillingProfileDelegate()
	responseData = delegate.createFromJson( billingProfile )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	billingProfile = json.loads(request.body)
	delegate = BillingProfileDelegate()
	responseData = delegate.save( billingProfile )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, billingProfileId ):
	delegate = BillingProfileDelegate()
	responseData = delegate.delete( billingProfileId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = BillingProfileDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignAdvertiser( request, billingProfileId, AdvertiserId ):
	delegate = BillingProfileDelegate()
	responseData = delegate.saveAdvertiser( billingProfileId, AdvertiserId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignAdvertiser( request, billingProfileId ):
	delegate = BillingProfileDelegate()
	responseData = delegate.deleteAdvertiser( billingProfileId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPaymentMethods( request, billingProfileId, PaymentMethodsIds ):
	delegate = BillingProfileDelegate()
	responseData = delegate.addPaymentMethods( billingProfileId, PaymentMethodsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePaymentMethods( request, billingProfileId, PaymentMethodsIds ):
	delegate = BillingProfileDelegate()
	responseData = delegate.removePaymentMethods( billingProfileId, PaymentMethodsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addAdAccounts( request, billingProfileId, AdAccountsIds ):
	delegate = BillingProfileDelegate()
	responseData = delegate.addAdAccounts( billingProfileId, AdAccountsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeAdAccounts( request, billingProfileId, AdAccountsIds ):
	delegate = BillingProfileDelegate()
	responseData = delegate.removeAdAccounts( billingProfileId, AdAccountsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

