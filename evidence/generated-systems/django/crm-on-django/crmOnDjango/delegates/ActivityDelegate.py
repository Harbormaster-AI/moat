from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from crmOnDjango.models.Activity import Activity
from crmOnDjango.models.Organization import Organization
from crmOnDjango.models.User import User
from crmOnDjango.models.Account import Account
from crmOnDjango.models.Contact import Contact
from crmOnDjango.models.Lead import Lead
from crmOnDjango.models.Opportunity import Opportunity
from crmOnDjango.models.Case_ import Case_
from crmOnDjango.models.Campaign import Campaign
from crmOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Activity
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ActivityDelegate Declaration
#======================================================================
class ActivityDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, activityId ):
		try:	
			activity = Activity.objects.filter(id=activityId)
			return activity.first();
		except Activity.DoesNotExist:
			raise ProcessingError("Activity with id " + str(activityId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, activity):
		for model in serializers.deserialize("json", activity):
			model.save()
			return model;

	def create(self, activity):
		activity.save()
		return activity;

	def saveFromJson(self, activity):
		for model in serializers.deserialize("json", activity):
			model.save()
			return activity;
	
	def save(self, activity):
		activity.save()
		return activity;
	
	def delete(self, activityId ):
		errMsg = "Failed to delete Activity from db using id " + str(activityId)
		
		try:
			activity = Activity.objects.get(id=activityId)
			activity.delete()
			return True
		except Activity.DoesNotExist:
			raise ProcessingError("Activity with id " + str(activityId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Activity.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Activity from db")
		except Exception:
			return None;
		
	def assignOrganization( self, activityId, organizationId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OrganizationDelegate import OrganizationDelegate

		errMsg = "Failed to assign element " + str(organizationId) + " for Organization on Activity"

		try:
			# get the Activity from db
			activity = self.get( activityId ).first()	
			
			# get the Organization from db
			organization = OrganizationDelegate().get(organizationId).first();
			
			# assign the Organization		
			activity.organization = organization
			
			#save it
			activity.save()

			# reload and return the appropriate version					
			return self.get( activityId );
		except Activity.DoesNotExist:
			raise ProcessingError(errMsg + " : Activity with id " + str(activityId) + " does not exist.")
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrganization( self, activityId ):
		errMsg = "Failed to unassign element " + str(organizationId) + " for Organization on Activity"

		try:
			# get the Activity from db
			activity = self.get( activityId ).first()	
			
			# assign to None for unassignment
			activity.organization = None			

			#save it
			activity.save()

			# reload and return the appropriate version					
			return self.get( activityId );
		except Activity.DoesNotExist:
			raise ProcessingError(errMsg + " : Activity with id " + str(activityId) + " does not exist.")
		except Exception:
			return None;
		
	def assignOwner( self, activityId, ownerId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.UserDelegate import UserDelegate

		errMsg = "Failed to assign element " + str(ownerId) + " for Owner on Activity"

		try:
			# get the Activity from db
			activity = self.get( activityId ).first()	
			
			# get the User from db
			user = UserDelegate().get(ownerId).first();
			
			# assign the Owner		
			activity.owner = user
			
			#save it
			activity.save()

			# reload and return the appropriate version					
			return self.get( activityId );
		except Activity.DoesNotExist:
			raise ProcessingError(errMsg + " : Activity with id " + str(activityId) + " does not exist.")
		except User.DoesNotExist:
			raise ProcessingError(errMsg + " : User with id " + str(ownerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOwner( self, activityId ):
		errMsg = "Failed to unassign element " + str(ownerId) + " for Owner on Activity"

		try:
			# get the Activity from db
			activity = self.get( activityId ).first()	
			
			# assign to None for unassignment
			activity.user = None			

			#save it
			activity.save()

			# reload and return the appropriate version					
			return self.get( activityId );
		except Activity.DoesNotExist:
			raise ProcessingError(errMsg + " : Activity with id " + str(activityId) + " does not exist.")
		except Exception:
			return None;
		
	def assignAccount( self, activityId, accountId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.AccountDelegate import AccountDelegate

		errMsg = "Failed to assign element " + str(accountId) + " for Account on Activity"

		try:
			# get the Activity from db
			activity = self.get( activityId ).first()	
			
			# get the Account from db
			account = AccountDelegate().get(accountId).first();
			
			# assign the Account		
			activity.account = account
			
			#save it
			activity.save()

			# reload and return the appropriate version					
			return self.get( activityId );
		except Activity.DoesNotExist:
			raise ProcessingError(errMsg + " : Activity with id " + str(activityId) + " does not exist.")
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignAccount( self, activityId ):
		errMsg = "Failed to unassign element " + str(accountId) + " for Account on Activity"

		try:
			# get the Activity from db
			activity = self.get( activityId ).first()	
			
			# assign to None for unassignment
			activity.account = None			

			#save it
			activity.save()

			# reload and return the appropriate version					
			return self.get( activityId );
		except Activity.DoesNotExist:
			raise ProcessingError(errMsg + " : Activity with id " + str(activityId) + " does not exist.")
		except Exception:
			return None;
		
	def assignContact( self, activityId, contactId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.ContactDelegate import ContactDelegate

		errMsg = "Failed to assign element " + str(contactId) + " for Contact on Activity"

		try:
			# get the Activity from db
			activity = self.get( activityId ).first()	
			
			# get the Contact from db
			contact = ContactDelegate().get(contactId).first();
			
			# assign the Contact		
			activity.contact = contact
			
			#save it
			activity.save()

			# reload and return the appropriate version					
			return self.get( activityId );
		except Activity.DoesNotExist:
			raise ProcessingError(errMsg + " : Activity with id " + str(activityId) + " does not exist.")
		except Contact.DoesNotExist:
			raise ProcessingError(errMsg + " : Contact with id " + str(contactId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignContact( self, activityId ):
		errMsg = "Failed to unassign element " + str(contactId) + " for Contact on Activity"

		try:
			# get the Activity from db
			activity = self.get( activityId ).first()	
			
			# assign to None for unassignment
			activity.contact = None			

			#save it
			activity.save()

			# reload and return the appropriate version					
			return self.get( activityId );
		except Activity.DoesNotExist:
			raise ProcessingError(errMsg + " : Activity with id " + str(activityId) + " does not exist.")
		except Exception:
			return None;
		
	def assignLead( self, activityId, leadId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.LeadDelegate import LeadDelegate

		errMsg = "Failed to assign element " + str(leadId) + " for Lead on Activity"

		try:
			# get the Activity from db
			activity = self.get( activityId ).first()	
			
			# get the Lead from db
			lead = LeadDelegate().get(leadId).first();
			
			# assign the Lead		
			activity.lead = lead
			
			#save it
			activity.save()

			# reload and return the appropriate version					
			return self.get( activityId );
		except Activity.DoesNotExist:
			raise ProcessingError(errMsg + " : Activity with id " + str(activityId) + " does not exist.")
		except Lead.DoesNotExist:
			raise ProcessingError(errMsg + " : Lead with id " + str(leadId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignLead( self, activityId ):
		errMsg = "Failed to unassign element " + str(leadId) + " for Lead on Activity"

		try:
			# get the Activity from db
			activity = self.get( activityId ).first()	
			
			# assign to None for unassignment
			activity.lead = None			

			#save it
			activity.save()

			# reload and return the appropriate version					
			return self.get( activityId );
		except Activity.DoesNotExist:
			raise ProcessingError(errMsg + " : Activity with id " + str(activityId) + " does not exist.")
		except Exception:
			return None;
		
	def assignOpportunity( self, activityId, opportunityId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OpportunityDelegate import OpportunityDelegate

		errMsg = "Failed to assign element " + str(opportunityId) + " for Opportunity on Activity"

		try:
			# get the Activity from db
			activity = self.get( activityId ).first()	
			
			# get the Opportunity from db
			opportunity = OpportunityDelegate().get(opportunityId).first();
			
			# assign the Opportunity		
			activity.opportunity = opportunity
			
			#save it
			activity.save()

			# reload and return the appropriate version					
			return self.get( activityId );
		except Activity.DoesNotExist:
			raise ProcessingError(errMsg + " : Activity with id " + str(activityId) + " does not exist.")
		except Opportunity.DoesNotExist:
			raise ProcessingError(errMsg + " : Opportunity with id " + str(opportunityId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOpportunity( self, activityId ):
		errMsg = "Failed to unassign element " + str(opportunityId) + " for Opportunity on Activity"

		try:
			# get the Activity from db
			activity = self.get( activityId ).first()	
			
			# assign to None for unassignment
			activity.opportunity = None			

			#save it
			activity.save()

			# reload and return the appropriate version					
			return self.get( activityId );
		except Activity.DoesNotExist:
			raise ProcessingError(errMsg + " : Activity with id " + str(activityId) + " does not exist.")
		except Exception:
			return None;
		
	def assignCase( self, activityId, caseId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.Case_Delegate import Case_Delegate

		errMsg = "Failed to assign element " + str(caseId) + " for Case on Activity"

		try:
			# get the Activity from db
			activity = self.get( activityId ).first()	
			
			# get the Case_ from db
			case_ = Case_Delegate().get(caseId).first();
			
			# assign the Case		
			activity.case = case_
			
			#save it
			activity.save()

			# reload and return the appropriate version					
			return self.get( activityId );
		except Activity.DoesNotExist:
			raise ProcessingError(errMsg + " : Activity with id " + str(activityId) + " does not exist.")
		except Case_.DoesNotExist:
			raise ProcessingError(errMsg + " : Case_ with id " + str(caseId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCase( self, activityId ):
		errMsg = "Failed to unassign element " + str(caseId) + " for Case on Activity"

		try:
			# get the Activity from db
			activity = self.get( activityId ).first()	
			
			# assign to None for unassignment
			activity.case_ = None			

			#save it
			activity.save()

			# reload and return the appropriate version					
			return self.get( activityId );
		except Activity.DoesNotExist:
			raise ProcessingError(errMsg + " : Activity with id " + str(activityId) + " does not exist.")
		except Exception:
			return None;
		
	def assignCampaign( self, activityId, campaignId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.CampaignDelegate import CampaignDelegate

		errMsg = "Failed to assign element " + str(campaignId) + " for Campaign on Activity"

		try:
			# get the Activity from db
			activity = self.get( activityId ).first()	
			
			# get the Campaign from db
			campaign = CampaignDelegate().get(campaignId).first();
			
			# assign the Campaign		
			activity.campaign = campaign
			
			#save it
			activity.save()

			# reload and return the appropriate version					
			return self.get( activityId );
		except Activity.DoesNotExist:
			raise ProcessingError(errMsg + " : Activity with id " + str(activityId) + " does not exist.")
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign with id " + str(campaignId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCampaign( self, activityId ):
		errMsg = "Failed to unassign element " + str(campaignId) + " for Campaign on Activity"

		try:
			# get the Activity from db
			activity = self.get( activityId ).first()	
			
			# assign to None for unassignment
			activity.campaign = None			

			#save it
			activity.save()

			# reload and return the appropriate version					
			return self.get( activityId );
		except Activity.DoesNotExist:
			raise ProcessingError(errMsg + " : Activity with id " + str(activityId) + " does not exist.")
		except Exception:
			return None;
		
