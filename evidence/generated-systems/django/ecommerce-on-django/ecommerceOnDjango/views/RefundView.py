import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from ecommerceOnDjango.delegates.RefundDelegate import RefundDelegate

 #======================================================================
# 
# Encapsulates data for View Refund
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RefundView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Refund index.")

def get(request, refundId ):
	delegate = RefundDelegate()
	responseData = delegate.get( refundId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	refund = json.loads(request.body)
	delegate = RefundDelegate()
	responseData = delegate.createFromJson( refund )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	refund = json.loads(request.body)
	delegate = RefundDelegate()
	responseData = delegate.save( refund )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, refundId ):
	delegate = RefundDelegate()
	responseData = delegate.delete( refundId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = RefundDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPayment( request, refundId, PaymentId ):
	delegate = RefundDelegate()
	responseData = delegate.savePayment( refundId, PaymentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPayment( request, refundId ):
	delegate = RefundDelegate()
	responseData = delegate.deletePayment( refundId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignOrder( request, refundId, OrderId ):
	delegate = RefundDelegate()
	responseData = delegate.saveOrder( refundId, OrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrder( request, refundId ):
	delegate = RefundDelegate()
	responseData = delegate.deleteOrder( refundId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

