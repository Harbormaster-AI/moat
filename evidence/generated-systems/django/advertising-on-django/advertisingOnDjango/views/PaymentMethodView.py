import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from advertisingOnDjango.delegates.PaymentMethodDelegate import PaymentMethodDelegate

 #======================================================================
# 
# Encapsulates data for View PaymentMethod
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PaymentMethodView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the PaymentMethod index.")

def get(request, paymentMethodId ):
	delegate = PaymentMethodDelegate()
	responseData = delegate.get( paymentMethodId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	paymentMethod = json.loads(request.body)
	delegate = PaymentMethodDelegate()
	responseData = delegate.createFromJson( paymentMethod )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	paymentMethod = json.loads(request.body)
	delegate = PaymentMethodDelegate()
	responseData = delegate.save( paymentMethod )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, paymentMethodId ):
	delegate = PaymentMethodDelegate()
	responseData = delegate.delete( paymentMethodId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = PaymentMethodDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignBillingProfile( request, paymentMethodId, BillingProfileId ):
	delegate = PaymentMethodDelegate()
	responseData = delegate.saveBillingProfile( paymentMethodId, BillingProfileId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignBillingProfile( request, paymentMethodId ):
	delegate = PaymentMethodDelegate()
	responseData = delegate.deleteBillingProfile( paymentMethodId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

