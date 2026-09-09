import json

from django.core import serializers
from django.shortcuts import render
from django.http import HttpResponse

from crmOnDjango.delegates.CampaignMemberDelegate import CampaignMemberDelegate

 #======================================================================
# 
# Encapsulates data for View CampaignMember
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CampaignMemberView function declarations
#======================================================================
def index(request):
	return HttpResponse("Hello, world. You're at the CampaignMember index.")

def get(request, campaignMemberId ):
	delegate = CampaignMemberDelegate()
	responseData = delegate.get( campaignMemberId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def create(request):
	campaignMember = json.loads(request.body)
	delegate = CampaignMemberDelegate()
	responseData = delegate.createFromJson( campaignMember )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def save(request):
	campaignMember = json.loads(request.body)
	delegate = CampaignMemberDelegate()
	responseData = delegate.save( campaignMember )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def delete(request, campaignMemberId ):
	delegate = CampaignMemberDelegate()
	responseData = delegate.delete( campaignMemberId )
	return HttpResponse(responseData, content_type="application/json");

def getAll(request):
	delegate = CampaignMemberDelegate()
	responseData = delegate.getAll()
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignCampaign( request, campaignMemberId, CampaignId ):
	delegate = CampaignMemberDelegate()
	responseData = delegate.saveCampaign( campaignMemberId, CampaignId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignCampaign( request, campaignMemberId ):
	delegate = CampaignMemberDelegate()
	responseData = delegate.deleteCampaign( campaignMemberId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignLead( request, campaignMemberId, LeadId ):
	delegate = CampaignMemberDelegate()
	responseData = delegate.saveLead( campaignMemberId, LeadId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignLead( request, campaignMemberId ):
	delegate = CampaignMemberDelegate()
	responseData = delegate.deleteLead( campaignMemberId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

def assignContact( request, campaignMemberId, ContactId ):
	delegate = CampaignMemberDelegate()
	responseData = delegate.saveContact( campaignMemberId, ContactId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");
	
def unassignContact( request, campaignMemberId ):
	delegate = CampaignMemberDelegate()
	responseData = delegate.deleteContact( campaignMemberId )
	asJson = serializers.serialize("json", responseData)
	return HttpResponse(asJson, content_type="application/json");

