import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from insuranceOnDjango.delegates.InvoiceDelegate import InvoiceDelegate

 #======================================================================
# 
# Encapsulates data for View Invoice
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InvoiceView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Invoice index.")

def get(request, invoiceId ):
	delegate = InvoiceDelegate()
	responseData = delegate.get( invoiceId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	invoice = json.loads(request.body)
	delegate = InvoiceDelegate()
	responseData = delegate.createFromJson( invoice )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	invoice = json.loads(request.body)
	delegate = InvoiceDelegate()
	responseData = delegate.save( invoice )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, invoiceId ):
	delegate = InvoiceDelegate()
	responseData = delegate.delete( invoiceId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = InvoiceDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignBillingAccount( request, invoiceId, BillingAccountId ):
	delegate = InvoiceDelegate()
	responseData = delegate.saveBillingAccount( invoiceId, BillingAccountId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignBillingAccount( request, invoiceId ):
	delegate = InvoiceDelegate()
	responseData = delegate.deleteBillingAccount( invoiceId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignPolicy( request, invoiceId, PolicyId ):
	delegate = InvoiceDelegate()
	responseData = delegate.savePolicy( invoiceId, PolicyId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPolicy( request, invoiceId ):
	delegate = InvoiceDelegate()
	responseData = delegate.deletePolicy( invoiceId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPayments( request, invoiceId, PaymentsIds ):
	delegate = InvoiceDelegate()
	responseData = delegate.addPayments( invoiceId, PaymentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePayments( request, invoiceId, PaymentsIds ):
	delegate = InvoiceDelegate()
	responseData = delegate.removePayments( invoiceId, PaymentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

