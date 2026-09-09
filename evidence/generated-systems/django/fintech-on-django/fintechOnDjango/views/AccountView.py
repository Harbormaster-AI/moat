import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from fintechOnDjango.delegates.AccountDelegate import AccountDelegate

 #======================================================================
# 
# Encapsulates data for View Account
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AccountView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Account index.")

def get(request, accountId ):
	delegate = AccountDelegate()
	responseData = delegate.get( accountId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	account = json.loads(request.body)
	delegate = AccountDelegate()
	responseData = delegate.createFromJson( account )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	account = json.loads(request.body)
	delegate = AccountDelegate()
	responseData = delegate.save( account )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, accountId ):
	delegate = AccountDelegate()
	responseData = delegate.delete( accountId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = AccountDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCustomer( request, accountId, CustomerId ):
	delegate = AccountDelegate()
	responseData = delegate.saveCustomer( accountId, CustomerId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCustomer( request, accountId ):
	delegate = AccountDelegate()
	responseData = delegate.deleteCustomer( accountId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignInstitution( request, accountId, InstitutionId ):
	delegate = AccountDelegate()
	responseData = delegate.saveInstitution( accountId, InstitutionId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignInstitution( request, accountId ):
	delegate = AccountDelegate()
	responseData = delegate.deleteInstitution( accountId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addTransactions( request, accountId, TransactionsIds ):
	delegate = AccountDelegate()
	responseData = delegate.addTransactions( accountId, TransactionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeTransactions( request, accountId, TransactionsIds ):
	delegate = AccountDelegate()
	responseData = delegate.removeTransactions( accountId, TransactionsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addCards( request, accountId, CardsIds ):
	delegate = AccountDelegate()
	responseData = delegate.addCards( accountId, CardsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeCards( request, accountId, CardsIds ):
	delegate = AccountDelegate()
	responseData = delegate.removeCards( accountId, CardsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addStatements( request, accountId, StatementsIds ):
	delegate = AccountDelegate()
	responseData = delegate.addStatements( accountId, StatementsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeStatements( request, accountId, StatementsIds ):
	delegate = AccountDelegate()
	responseData = delegate.removeStatements( accountId, StatementsIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def addMandates( request, accountId, MandatesIds ):
	delegate = AccountDelegate()
	responseData = delegate.addMandates( accountId, MandatesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def removeMandates( request, accountId, MandatesIds ):
	delegate = AccountDelegate()
	responseData = delegate.removeMandates( accountId, MandatesIds )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

