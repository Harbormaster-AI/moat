from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from crmOnDjango.models.EmailMessage import EmailMessage
from crmOnDjango.models.Organization import Organization
from crmOnDjango.models.User import User
from crmOnDjango.models.Account import Account
from crmOnDjango.models.Contact import Contact
from crmOnDjango.models.Lead import Lead
from crmOnDjango.models.Case_ import Case_
from crmOnDjango.models.Opportunity import Opportunity
from crmOnDjango.models.Campaign import Campaign
from crmOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model EmailMessage
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class EmailMessageDelegate Declaration
#======================================================================
class EmailMessageDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, emailMessageId ):
		try:	
			emailMessage = EmailMessage.objects.filter(id=emailMessageId)
			return emailMessage.first();
		except EmailMessage.DoesNotExist:
			raise ProcessingError("EmailMessage with id " + str(emailMessageId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, emailMessage):
		for model in serializers.deserialize("json", emailMessage):
			model.save()
			return model;

	def create(self, emailMessage):
		emailMessage.save()
		return emailMessage;

	def saveFromJson(self, emailMessage):
		for model in serializers.deserialize("json", emailMessage):
			model.save()
			return emailMessage;
	
	def save(self, emailMessage):
		emailMessage.save()
		return emailMessage;
	
	def delete(self, emailMessageId ):
		errMsg = "Failed to delete EmailMessage from db using id " + str(emailMessageId)
		
		try:
			emailMessage = EmailMessage.objects.get(id=emailMessageId)
			emailMessage.delete()
			return True
		except EmailMessage.DoesNotExist:
			raise ProcessingError("EmailMessage with id " + str(emailMessageId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = EmailMessage.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all EmailMessage from db")
		except Exception:
			return None;
		
	def assignOrganization( self, emailMessageId, organizationId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OrganizationDelegate import OrganizationDelegate

		errMsg = "Failed to assign element " + str(organizationId) + " for Organization on EmailMessage"

		try:
			# get the EmailMessage from db
			emailMessage = self.get( emailMessageId ).first()	
			
			# get the Organization from db
			organization = OrganizationDelegate().get(organizationId).first();
			
			# assign the Organization		
			emailMessage.organization = organization
			
			#save it
			emailMessage.save()

			# reload and return the appropriate version					
			return self.get( emailMessageId );
		except EmailMessage.DoesNotExist:
			raise ProcessingError(errMsg + " : EmailMessage with id " + str(emailMessageId) + " does not exist.")
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrganization( self, emailMessageId ):
		errMsg = "Failed to unassign element " + str(organizationId) + " for Organization on EmailMessage"

		try:
			# get the EmailMessage from db
			emailMessage = self.get( emailMessageId ).first()	
			
			# assign to None for unassignment
			emailMessage.organization = None			

			#save it
			emailMessage.save()

			# reload and return the appropriate version					
			return self.get( emailMessageId );
		except EmailMessage.DoesNotExist:
			raise ProcessingError(errMsg + " : EmailMessage with id " + str(emailMessageId) + " does not exist.")
		except Exception:
			return None;
		
	def assignOwner( self, emailMessageId, ownerId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.UserDelegate import UserDelegate

		errMsg = "Failed to assign element " + str(ownerId) + " for Owner on EmailMessage"

		try:
			# get the EmailMessage from db
			emailMessage = self.get( emailMessageId ).first()	
			
			# get the User from db
			user = UserDelegate().get(ownerId).first();
			
			# assign the Owner		
			emailMessage.owner = user
			
			#save it
			emailMessage.save()

			# reload and return the appropriate version					
			return self.get( emailMessageId );
		except EmailMessage.DoesNotExist:
			raise ProcessingError(errMsg + " : EmailMessage with id " + str(emailMessageId) + " does not exist.")
		except User.DoesNotExist:
			raise ProcessingError(errMsg + " : User with id " + str(ownerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOwner( self, emailMessageId ):
		errMsg = "Failed to unassign element " + str(ownerId) + " for Owner on EmailMessage"

		try:
			# get the EmailMessage from db
			emailMessage = self.get( emailMessageId ).first()	
			
			# assign to None for unassignment
			emailMessage.user = None			

			#save it
			emailMessage.save()

			# reload and return the appropriate version					
			return self.get( emailMessageId );
		except EmailMessage.DoesNotExist:
			raise ProcessingError(errMsg + " : EmailMessage with id " + str(emailMessageId) + " does not exist.")
		except Exception:
			return None;
		
	def assignAccount( self, emailMessageId, accountId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.AccountDelegate import AccountDelegate

		errMsg = "Failed to assign element " + str(accountId) + " for Account on EmailMessage"

		try:
			# get the EmailMessage from db
			emailMessage = self.get( emailMessageId ).first()	
			
			# get the Account from db
			account = AccountDelegate().get(accountId).first();
			
			# assign the Account		
			emailMessage.account = account
			
			#save it
			emailMessage.save()

			# reload and return the appropriate version					
			return self.get( emailMessageId );
		except EmailMessage.DoesNotExist:
			raise ProcessingError(errMsg + " : EmailMessage with id " + str(emailMessageId) + " does not exist.")
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignAccount( self, emailMessageId ):
		errMsg = "Failed to unassign element " + str(accountId) + " for Account on EmailMessage"

		try:
			# get the EmailMessage from db
			emailMessage = self.get( emailMessageId ).first()	
			
			# assign to None for unassignment
			emailMessage.account = None			

			#save it
			emailMessage.save()

			# reload and return the appropriate version					
			return self.get( emailMessageId );
		except EmailMessage.DoesNotExist:
			raise ProcessingError(errMsg + " : EmailMessage with id " + str(emailMessageId) + " does not exist.")
		except Exception:
			return None;
		
	def assignContact( self, emailMessageId, contactId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.ContactDelegate import ContactDelegate

		errMsg = "Failed to assign element " + str(contactId) + " for Contact on EmailMessage"

		try:
			# get the EmailMessage from db
			emailMessage = self.get( emailMessageId ).first()	
			
			# get the Contact from db
			contact = ContactDelegate().get(contactId).first();
			
			# assign the Contact		
			emailMessage.contact = contact
			
			#save it
			emailMessage.save()

			# reload and return the appropriate version					
			return self.get( emailMessageId );
		except EmailMessage.DoesNotExist:
			raise ProcessingError(errMsg + " : EmailMessage with id " + str(emailMessageId) + " does not exist.")
		except Contact.DoesNotExist:
			raise ProcessingError(errMsg + " : Contact with id " + str(contactId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignContact( self, emailMessageId ):
		errMsg = "Failed to unassign element " + str(contactId) + " for Contact on EmailMessage"

		try:
			# get the EmailMessage from db
			emailMessage = self.get( emailMessageId ).first()	
			
			# assign to None for unassignment
			emailMessage.contact = None			

			#save it
			emailMessage.save()

			# reload and return the appropriate version					
			return self.get( emailMessageId );
		except EmailMessage.DoesNotExist:
			raise ProcessingError(errMsg + " : EmailMessage with id " + str(emailMessageId) + " does not exist.")
		except Exception:
			return None;
		
	def assignLead( self, emailMessageId, leadId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.LeadDelegate import LeadDelegate

		errMsg = "Failed to assign element " + str(leadId) + " for Lead on EmailMessage"

		try:
			# get the EmailMessage from db
			emailMessage = self.get( emailMessageId ).first()	
			
			# get the Lead from db
			lead = LeadDelegate().get(leadId).first();
			
			# assign the Lead		
			emailMessage.lead = lead
			
			#save it
			emailMessage.save()

			# reload and return the appropriate version					
			return self.get( emailMessageId );
		except EmailMessage.DoesNotExist:
			raise ProcessingError(errMsg + " : EmailMessage with id " + str(emailMessageId) + " does not exist.")
		except Lead.DoesNotExist:
			raise ProcessingError(errMsg + " : Lead with id " + str(leadId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignLead( self, emailMessageId ):
		errMsg = "Failed to unassign element " + str(leadId) + " for Lead on EmailMessage"

		try:
			# get the EmailMessage from db
			emailMessage = self.get( emailMessageId ).first()	
			
			# assign to None for unassignment
			emailMessage.lead = None			

			#save it
			emailMessage.save()

			# reload and return the appropriate version					
			return self.get( emailMessageId );
		except EmailMessage.DoesNotExist:
			raise ProcessingError(errMsg + " : EmailMessage with id " + str(emailMessageId) + " does not exist.")
		except Exception:
			return None;
		
	def assignCase( self, emailMessageId, caseId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.Case_Delegate import Case_Delegate

		errMsg = "Failed to assign element " + str(caseId) + " for Case on EmailMessage"

		try:
			# get the EmailMessage from db
			emailMessage = self.get( emailMessageId ).first()	
			
			# get the Case_ from db
			case_ = Case_Delegate().get(caseId).first();
			
			# assign the Case		
			emailMessage.case = case_
			
			#save it
			emailMessage.save()

			# reload and return the appropriate version					
			return self.get( emailMessageId );
		except EmailMessage.DoesNotExist:
			raise ProcessingError(errMsg + " : EmailMessage with id " + str(emailMessageId) + " does not exist.")
		except Case_.DoesNotExist:
			raise ProcessingError(errMsg + " : Case_ with id " + str(caseId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCase( self, emailMessageId ):
		errMsg = "Failed to unassign element " + str(caseId) + " for Case on EmailMessage"

		try:
			# get the EmailMessage from db
			emailMessage = self.get( emailMessageId ).first()	
			
			# assign to None for unassignment
			emailMessage.case_ = None			

			#save it
			emailMessage.save()

			# reload and return the appropriate version					
			return self.get( emailMessageId );
		except EmailMessage.DoesNotExist:
			raise ProcessingError(errMsg + " : EmailMessage with id " + str(emailMessageId) + " does not exist.")
		except Exception:
			return None;
		
	def assignOpportunity( self, emailMessageId, opportunityId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OpportunityDelegate import OpportunityDelegate

		errMsg = "Failed to assign element " + str(opportunityId) + " for Opportunity on EmailMessage"

		try:
			# get the EmailMessage from db
			emailMessage = self.get( emailMessageId ).first()	
			
			# get the Opportunity from db
			opportunity = OpportunityDelegate().get(opportunityId).first();
			
			# assign the Opportunity		
			emailMessage.opportunity = opportunity
			
			#save it
			emailMessage.save()

			# reload and return the appropriate version					
			return self.get( emailMessageId );
		except EmailMessage.DoesNotExist:
			raise ProcessingError(errMsg + " : EmailMessage with id " + str(emailMessageId) + " does not exist.")
		except Opportunity.DoesNotExist:
			raise ProcessingError(errMsg + " : Opportunity with id " + str(opportunityId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOpportunity( self, emailMessageId ):
		errMsg = "Failed to unassign element " + str(opportunityId) + " for Opportunity on EmailMessage"

		try:
			# get the EmailMessage from db
			emailMessage = self.get( emailMessageId ).first()	
			
			# assign to None for unassignment
			emailMessage.opportunity = None			

			#save it
			emailMessage.save()

			# reload and return the appropriate version					
			return self.get( emailMessageId );
		except EmailMessage.DoesNotExist:
			raise ProcessingError(errMsg + " : EmailMessage with id " + str(emailMessageId) + " does not exist.")
		except Exception:
			return None;
		
	def assignCampaign( self, emailMessageId, campaignId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.CampaignDelegate import CampaignDelegate

		errMsg = "Failed to assign element " + str(campaignId) + " for Campaign on EmailMessage"

		try:
			# get the EmailMessage from db
			emailMessage = self.get( emailMessageId ).first()	
			
			# get the Campaign from db
			campaign = CampaignDelegate().get(campaignId).first();
			
			# assign the Campaign		
			emailMessage.campaign = campaign
			
			#save it
			emailMessage.save()

			# reload and return the appropriate version					
			return self.get( emailMessageId );
		except EmailMessage.DoesNotExist:
			raise ProcessingError(errMsg + " : EmailMessage with id " + str(emailMessageId) + " does not exist.")
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign with id " + str(campaignId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCampaign( self, emailMessageId ):
		errMsg = "Failed to unassign element " + str(campaignId) + " for Campaign on EmailMessage"

		try:
			# get the EmailMessage from db
			emailMessage = self.get( emailMessageId ).first()	
			
			# assign to None for unassignment
			emailMessage.campaign = None			

			#save it
			emailMessage.save()

			# reload and return the appropriate version					
			return self.get( emailMessageId );
		except EmailMessage.DoesNotExist:
			raise ProcessingError(errMsg + " : EmailMessage with id " + str(emailMessageId) + " does not exist.")
		except Exception:
			return None;
		
