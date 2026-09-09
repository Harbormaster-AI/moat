import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from fintechOnDjango.delegates.AppliedFeeDelegate import AppliedFeeDelegate

 #======================================================================
# 
# Encapsulates data for View AppliedFee
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AppliedFeeView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the AppliedFee index.")

def get(request, appliedFeeId ):
	delegate = AppliedFeeDelegate()
	responseData = delegate.get( appliedFeeId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	appliedFee = json.loads(request.body)
	delegate = AppliedFeeDelegate()
	responseData = delegate.createFromJson( appliedFee )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	appliedFee = json.loads(request.body)
	delegate = AppliedFeeDelegate()
	responseData = delegate.save( appliedFee )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, appliedFeeId ):
	delegate = AppliedFeeDelegate()
	responseData = delegate.delete( appliedFeeId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = AppliedFeeDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPaymentOrder( request, appliedFeeId, PaymentOrderId ):
	delegate = AppliedFeeDelegate()
	responseData = delegate.savePaymentOrder( appliedFeeId, PaymentOrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPaymentOrder( request, appliedFeeId ):
	delegate = AppliedFeeDelegate()
	responseData = delegate.deletePaymentOrder( appliedFeeId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignTransaction( request, appliedFeeId, TransactionId ):
	delegate = AppliedFeeDelegate()
	responseData = delegate.saveTransaction( appliedFeeId, TransactionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignTransaction( request, appliedFeeId ):
	delegate = AppliedFeeDelegate()
	responseData = delegate.deleteTransaction( appliedFeeId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

