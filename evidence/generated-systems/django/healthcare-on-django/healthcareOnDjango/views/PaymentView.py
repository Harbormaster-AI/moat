import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from healthcareOnDjango.delegates.PaymentDelegate import PaymentDelegate

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

def assignInvoice( request, paymentId, InvoiceId ):
	delegate = PaymentDelegate()
	responseData = delegate.saveInvoice( paymentId, InvoiceId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignInvoice( request, paymentId ):
	delegate = PaymentDelegate()
	responseData = delegate.deleteInvoice( paymentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPayer( request, paymentId, PayerId ):
	delegate = PaymentDelegate()
	responseData = delegate.savePayer( paymentId, PayerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPayer( request, paymentId ):
	delegate = PaymentDelegate()
	responseData = delegate.deletePayer( paymentId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

