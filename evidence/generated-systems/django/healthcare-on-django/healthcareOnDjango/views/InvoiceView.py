import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from healthcareOnDjango.delegates.InvoiceDelegate import InvoiceDelegate

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

def assignPatient( request, invoiceId, PatientId ):
	delegate = InvoiceDelegate()
	responseData = delegate.savePatient( invoiceId, PatientId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignPatient( request, invoiceId ):
	delegate = InvoiceDelegate()
	responseData = delegate.deletePatient( invoiceId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignClaim( request, invoiceId, ClaimId ):
	delegate = InvoiceDelegate()
	responseData = delegate.saveClaim( invoiceId, ClaimId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignClaim( request, invoiceId ):
	delegate = InvoiceDelegate()
	responseData = delegate.deleteClaim( invoiceId )
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

