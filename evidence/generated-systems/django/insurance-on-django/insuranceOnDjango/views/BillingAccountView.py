import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from insuranceOnDjango.delegates.BillingAccountDelegate import BillingAccountDelegate

 #======================================================================
# 
# Encapsulates data for View BillingAccount
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BillingAccountView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the BillingAccount index.")

def get(request, billingAccountId ):
	delegate = BillingAccountDelegate()
	responseData = delegate.get( billingAccountId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	billingAccount = json.loads(request.body)
	delegate = BillingAccountDelegate()
	responseData = delegate.createFromJson( billingAccount )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	billingAccount = json.loads(request.body)
	delegate = BillingAccountDelegate()
	responseData = delegate.save( billingAccount )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, billingAccountId ):
	delegate = BillingAccountDelegate()
	responseData = delegate.delete( billingAccountId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = BillingAccountDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCustomer( request, billingAccountId, CustomerId ):
	delegate = BillingAccountDelegate()
	responseData = delegate.saveCustomer( billingAccountId, CustomerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCustomer( request, billingAccountId ):
	delegate = BillingAccountDelegate()
	responseData = delegate.deleteCustomer( billingAccountId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPolicies( request, billingAccountId, PoliciesIds ):
	delegate = BillingAccountDelegate()
	responseData = delegate.addPolicies( billingAccountId, PoliciesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePolicies( request, billingAccountId, PoliciesIds ):
	delegate = BillingAccountDelegate()
	responseData = delegate.removePolicies( billingAccountId, PoliciesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addInvoices( request, billingAccountId, InvoicesIds ):
	delegate = BillingAccountDelegate()
	responseData = delegate.addInvoices( billingAccountId, InvoicesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeInvoices( request, billingAccountId, InvoicesIds ):
	delegate = BillingAccountDelegate()
	responseData = delegate.removeInvoices( billingAccountId, InvoicesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addPayments( request, billingAccountId, PaymentsIds ):
	delegate = BillingAccountDelegate()
	responseData = delegate.addPayments( billingAccountId, PaymentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removePayments( request, billingAccountId, PaymentsIds ):
	delegate = BillingAccountDelegate()
	responseData = delegate.removePayments( billingAccountId, PaymentsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

