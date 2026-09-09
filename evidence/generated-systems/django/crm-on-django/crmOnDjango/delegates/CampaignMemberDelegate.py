from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from crmOnDjango.models.CampaignMember import CampaignMember
from crmOnDjango.models.Campaign import Campaign
from crmOnDjango.models.Lead import Lead
from crmOnDjango.models.Contact import Contact
from crmOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model CampaignMember
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CampaignMemberDelegate Declaration
#======================================================================
class CampaignMemberDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, campaignMemberId ):
		try:	
			campaignMember = CampaignMember.objects.filter(id=campaignMemberId)
			return campaignMember.first();
		except CampaignMember.DoesNotExist:
			raise ProcessingError("CampaignMember with id " + str(campaignMemberId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, campaignMember):
		for model in serializers.deserialize("json", campaignMember):
			model.save()
			return model;

	def create(self, campaignMember):
		campaignMember.save()
		return campaignMember;

	def saveFromJson(self, campaignMember):
		for model in serializers.deserialize("json", campaignMember):
			model.save()
			return campaignMember;
	
	def save(self, campaignMember):
		campaignMember.save()
		return campaignMember;
	
	def delete(self, campaignMemberId ):
		errMsg = "Failed to delete CampaignMember from db using id " + str(campaignMemberId)
		
		try:
			campaignMember = CampaignMember.objects.get(id=campaignMemberId)
			campaignMember.delete()
			return True
		except CampaignMember.DoesNotExist:
			raise ProcessingError("CampaignMember with id " + str(campaignMemberId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = CampaignMember.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all CampaignMember from db")
		except Exception:
			return None;
		
	def assignCampaign( self, campaignMemberId, campaignId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.CampaignDelegate import CampaignDelegate

		errMsg = "Failed to assign element " + str(campaignId) + " for Campaign on CampaignMember"

		try:
			# get the CampaignMember from db
			campaignMember = self.get( campaignMemberId ).first()	
			
			# get the Campaign from db
			campaign = CampaignDelegate().get(campaignId).first();
			
			# assign the Campaign		
			campaignMember.campaign = campaign
			
			#save it
			campaignMember.save()

			# reload and return the appropriate version					
			return self.get( campaignMemberId );
		except CampaignMember.DoesNotExist:
			raise ProcessingError(errMsg + " : CampaignMember with id " + str(campaignMemberId) + " does not exist.")
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign with id " + str(campaignId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCampaign( self, campaignMemberId ):
		errMsg = "Failed to unassign element " + str(campaignId) + " for Campaign on CampaignMember"

		try:
			# get the CampaignMember from db
			campaignMember = self.get( campaignMemberId ).first()	
			
			# assign to None for unassignment
			campaignMember.campaign = None			

			#save it
			campaignMember.save()

			# reload and return the appropriate version					
			return self.get( campaignMemberId );
		except CampaignMember.DoesNotExist:
			raise ProcessingError(errMsg + " : CampaignMember with id " + str(campaignMemberId) + " does not exist.")
		except Exception:
			return None;
		
	def assignLead( self, campaignMemberId, leadId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.LeadDelegate import LeadDelegate

		errMsg = "Failed to assign element " + str(leadId) + " for Lead on CampaignMember"

		try:
			# get the CampaignMember from db
			campaignMember = self.get( campaignMemberId ).first()	
			
			# get the Lead from db
			lead = LeadDelegate().get(leadId).first();
			
			# assign the Lead		
			campaignMember.lead = lead
			
			#save it
			campaignMember.save()

			# reload and return the appropriate version					
			return self.get( campaignMemberId );
		except CampaignMember.DoesNotExist:
			raise ProcessingError(errMsg + " : CampaignMember with id " + str(campaignMemberId) + " does not exist.")
		except Lead.DoesNotExist:
			raise ProcessingError(errMsg + " : Lead with id " + str(leadId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignLead( self, campaignMemberId ):
		errMsg = "Failed to unassign element " + str(leadId) + " for Lead on CampaignMember"

		try:
			# get the CampaignMember from db
			campaignMember = self.get( campaignMemberId ).first()	
			
			# assign to None for unassignment
			campaignMember.lead = None			

			#save it
			campaignMember.save()

			# reload and return the appropriate version					
			return self.get( campaignMemberId );
		except CampaignMember.DoesNotExist:
			raise ProcessingError(errMsg + " : CampaignMember with id " + str(campaignMemberId) + " does not exist.")
		except Exception:
			return None;
		
	def assignContact( self, campaignMemberId, contactId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.ContactDelegate import ContactDelegate

		errMsg = "Failed to assign element " + str(contactId) + " for Contact on CampaignMember"

		try:
			# get the CampaignMember from db
			campaignMember = self.get( campaignMemberId ).first()	
			
			# get the Contact from db
			contact = ContactDelegate().get(contactId).first();
			
			# assign the Contact		
			campaignMember.contact = contact
			
			#save it
			campaignMember.save()

			# reload and return the appropriate version					
			return self.get( campaignMemberId );
		except CampaignMember.DoesNotExist:
			raise ProcessingError(errMsg + " : CampaignMember with id " + str(campaignMemberId) + " does not exist.")
		except Contact.DoesNotExist:
			raise ProcessingError(errMsg + " : Contact with id " + str(contactId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignContact( self, campaignMemberId ):
		errMsg = "Failed to unassign element " + str(contactId) + " for Contact on CampaignMember"

		try:
			# get the CampaignMember from db
			campaignMember = self.get( campaignMemberId ).first()	
			
			# assign to None for unassignment
			campaignMember.contact = None			

			#save it
			campaignMember.save()

			# reload and return the appropriate version					
			return self.get( campaignMemberId );
		except CampaignMember.DoesNotExist:
			raise ProcessingError(errMsg + " : CampaignMember with id " + str(campaignMemberId) + " does not exist.")
		except Exception:
			return None;
		
