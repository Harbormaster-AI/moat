import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from hrOnDjango.delegates.BankAccountDelegate import BankAccountDelegate

 #======================================================================
# 
# Encapsulates data for View BankAccount
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BankAccountView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the BankAccount index.")

def get(request, bankAccountId ):
	delegate = BankAccountDelegate()
	responseData = delegate.get( bankAccountId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	bankAccount = json.loads(request.body)
	delegate = BankAccountDelegate()
	responseData = delegate.createFromJson( bankAccount )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	bankAccount = json.loads(request.body)
	delegate = BankAccountDelegate()
	responseData = delegate.save( bankAccount )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, bankAccountId ):
	delegate = BankAccountDelegate()
	responseData = delegate.delete( bankAccountId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = BankAccountDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

