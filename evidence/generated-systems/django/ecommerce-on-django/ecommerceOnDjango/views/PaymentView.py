import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from ecommerceOnDjango.delegates.PaymentDelegate import PaymentDelegate

 #======================================================================
# 
# Encapsulates data for View Payment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PaymentView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Payment index.")

def get(request, paymentId ):
	delegate = PaymentDelegate()
	responseData = delegate.get( paymentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	payment = json.loads(request.body)
	delegate = PaymentDelegate()
	responseData = delegate.createFromJson( payment )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	payment = json.loads(request.body)
	delegate = PaymentDelegate()
	responseData = delegate.save( payment )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, paymentId ):
	delegate = PaymentDelegate()
	responseData = delegate.delete( paymentId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = PaymentDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrder( request, paymentId, OrderId ):
	delegate = PaymentDelegate()
	responseData = delegate.saveOrder( paymentId, OrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrder( request, paymentId ):
	delegate = PaymentDelegate()
	responseData = delegate.deleteOrder( paymentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCustomer( request, paymentId, CustomerId ):
	delegate = PaymentDelegate()
	responseData = delegate.saveCustomer( paymentId, CustomerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCustomer( request, paymentId ):
	delegate = PaymentDelegate()
	responseData = delegate.deleteCustomer( paymentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPaymentProvider( request, paymentId, PaymentProviderId ):
	delegate = PaymentDelegate()
	responseData = delegate.savePaymentProvider( paymentId, PaymentProviderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPaymentProvider( request, paymentId ):
	delegate = PaymentDelegate()
	responseData = delegate.deletePaymentProvider( paymentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addRefunds( request, paymentId, RefundsIds ):
	delegate = PaymentDelegate()
	responseData = delegate.addRefunds( paymentId, RefundsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeRefunds( request, paymentId, RefundsIds ):
	delegate = PaymentDelegate()
	responseData = delegate.removeRefunds( paymentId, RefundsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

