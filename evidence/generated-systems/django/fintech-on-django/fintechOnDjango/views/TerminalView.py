import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from fintechOnDjango.delegates.TerminalDelegate import TerminalDelegate

 #======================================================================
# 
# Encapsulates data for View Terminal
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TerminalView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the Terminal index.")

def get(request, terminalId ):
	delegate = TerminalDelegate()
	responseData = delegate.get( terminalId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	terminal = json.loads(request.body)
	delegate = TerminalDelegate()
	responseData = delegate.createFromJson( terminal )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	terminal = json.loads(request.body)
	delegate = TerminalDelegate()
	responseData = delegate.save( terminal )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, terminalId ):
	delegate = TerminalDelegate()
	responseData = delegate.delete( terminalId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = TerminalDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignMerchant( request, terminalId, MerchantId ):
	delegate = TerminalDelegate()
	responseData = delegate.saveMerchant( terminalId, MerchantId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignMerchant( request, terminalId ):
	delegate = TerminalDelegate()
	responseData = delegate.deleteMerchant( terminalId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

