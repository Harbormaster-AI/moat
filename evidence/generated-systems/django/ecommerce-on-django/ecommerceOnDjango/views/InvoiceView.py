import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from ecommerceOnDjango.delegates.InvoiceDelegate import InvoiceDelegate

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

def assignOrder( request, invoiceId, OrderId ):
	delegate = InvoiceDelegate()
	responseData = delegate.saveOrder( invoiceId, OrderId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignOrder( request, invoiceId ):
	delegate = InvoiceDelegate()
	responseData = delegate.deleteOrder( invoiceId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

