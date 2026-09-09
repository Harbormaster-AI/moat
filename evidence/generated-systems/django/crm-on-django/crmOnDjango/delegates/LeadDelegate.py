from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from crmOnDjango.models.Lead import Lead
from crmOnDjango.models.Organization import Organization
from crmOnDjango.models.User import User
from crmOnDjango.models.Activity import Activity
from crmOnDjango.models.Campaign import Campaign
from crmOnDjango.models.Account import Account
from crmOnDjango.models.Contact import Contact
from crmOnDjango.models.Opportunity import Opportunity
from crmOnDjango.models.Note import Note
from crmOnDjango.models.EmailMessage import EmailMessage
from crmOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Lead
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LeadDelegate Declaration
#======================================================================
class LeadDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, leadId ):
		try:	
			lead = Lead.objects.filter(id=leadId)
			return lead.first();
		except Lead.DoesNotExist:
			raise ProcessingError("Lead with id " + str(leadId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, lead):
		for model in serializers.deserialize("json", lead):
			model.save()
			return model;

	def create(self, lead):
		lead.save()
		return lead;

	def saveFromJson(self, lead):
		for model in serializers.deserialize("json", lead):
			model.save()
			return lead;
	
	def save(self, lead):
		lead.save()
		return lead;
	
	def delete(self, leadId ):
		errMsg = "Failed to delete Lead from db using id " + str(leadId)
		
		try:
			lead = Lead.objects.get(id=leadId)
			lead.delete()
			return True
		except Lead.DoesNotExist:
			raise ProcessingError("Lead with id " + str(leadId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Lead.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Lead from db")
		except Exception:
			return None;
		
	def assignOrganization( self, leadId, organizationId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OrganizationDelegate import OrganizationDelegate

		errMsg = "Failed to assign element " + str(organizationId) + " for Organization on Lead"

		try:
			# get the Lead from db
			lead = self.get( leadId ).first()	
			
			# get the Organization from db
			organization = OrganizationDelegate().get(organizationId).first();
			
			# assign the Organization		
			lead.organization = organization
			
			#save it
			lead.save()

			# reload and return the appropriate version					
			return self.get( leadId );
		except Lead.DoesNotExist:
			raise ProcessingError(errMsg + " : Lead with id " + str(leadId) + " does not exist.")
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrganization( self, leadId ):
		errMsg = "Failed to unassign element " + str(organizationId) + " for Organization on Lead"

		try:
			# get the Lead from db
			lead = self.get( leadId ).first()	
			
			# assign to None for unassignment
			lead.organization = None			

			#save it
			lead.save()

			# reload and return the appropriate version					
			return self.get( leadId );
		except Lead.DoesNotExist:
			raise ProcessingError(errMsg + " : Lead with id " + str(leadId) + " does not exist.")
		except Exception:
			return None;
		
	def assignOwner( self, leadId, ownerId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.UserDelegate import UserDelegate

		errMsg = "Failed to assign element " + str(ownerId) + " for Owner on Lead"

		try:
			# get the Lead from db
			lead = self.get( leadId ).first()	
			
			# get the User from db
			user = UserDelegate().get(ownerId).first();
			
			# assign the Owner		
			lead.owner = user
			
			#save it
			lead.save()

			# reload and return the appropriate version					
			return self.get( leadId );
		except Lead.DoesNotExist:
			raise ProcessingError(errMsg + " : Lead with id " + str(leadId) + " does not exist.")
		except User.DoesNotExist:
			raise ProcessingError(errMsg + " : User with id " + str(ownerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOwner( self, leadId ):
		errMsg = "Failed to unassign element " + str(ownerId) + " for Owner on Lead"

		try:
			# get the Lead from db
			lead = self.get( leadId ).first()	
			
			# assign to None for unassignment
			lead.user = None			

			#save it
			lead.save()

			# reload and return the appropriate version					
			return self.get( leadId );
		except Lead.DoesNotExist:
			raise ProcessingError(errMsg + " : Lead with id " + str(leadId) + " does not exist.")
		except Exception:
			return None;
		
	def assignConvertedAccount( self, leadId, convertedAccountId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.AccountDelegate import AccountDelegate

		errMsg = "Failed to assign element " + str(convertedAccountId) + " for ConvertedAccount on Lead"

		try:
			# get the Lead from db
			lead = self.get( leadId ).first()	
			
			# get the Account from db
			account = AccountDelegate().get(convertedAccountId).first();
			
			# assign the ConvertedAccount		
			lead.convertedAccount = account
			
			#save it
			lead.save()

			# reload and return the appropriate version					
			return self.get( leadId );
		except Lead.DoesNotExist:
			raise ProcessingError(errMsg + " : Lead with id " + str(leadId) + " does not exist.")
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(convertedAccountId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignConvertedAccount( self, leadId ):
		errMsg = "Failed to unassign element " + str(convertedAccountId) + " for ConvertedAccount on Lead"

		try:
			# get the Lead from db
			lead = self.get( leadId ).first()	
			
			# assign to None for unassignment
			lead.account = None			

			#save it
			lead.save()

			# reload and return the appropriate version					
			return self.get( leadId );
		except Lead.DoesNotExist:
			raise ProcessingError(errMsg + " : Lead with id " + str(leadId) + " does not exist.")
		except Exception:
			return None;
		
	def assignConvertedContact( self, leadId, convertedContactId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.ContactDelegate import ContactDelegate

		errMsg = "Failed to assign element " + str(convertedContactId) + " for ConvertedContact on Lead"

		try:
			# get the Lead from db
			lead = self.get( leadId ).first()	
			
			# get the Contact from db
			contact = ContactDelegate().get(convertedContactId).first();
			
			# assign the ConvertedContact		
			lead.convertedContact = contact
			
			#save it
			lead.save()

			# reload and return the appropriate version					
			return self.get( leadId );
		except Lead.DoesNotExist:
			raise ProcessingError(errMsg + " : Lead with id " + str(leadId) + " does not exist.")
		except Contact.DoesNotExist:
			raise ProcessingError(errMsg + " : Contact with id " + str(convertedContactId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignConvertedContact( self, leadId ):
		errMsg = "Failed to unassign element " + str(convertedContactId) + " for ConvertedContact on Lead"

		try:
			# get the Lead from db
			lead = self.get( leadId ).first()	
			
			# assign to None for unassignment
			lead.contact = None			

			#save it
			lead.save()

			# reload and return the appropriate version					
			return self.get( leadId );
		except Lead.DoesNotExist:
			raise ProcessingError(errMsg + " : Lead with id " + str(leadId) + " does not exist.")
		except Exception:
			return None;
		
	def assignConvertedOpportunity( self, leadId, convertedOpportunityId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OpportunityDelegate import OpportunityDelegate

		errMsg = "Failed to assign element " + str(convertedOpportunityId) + " for ConvertedOpportunity on Lead"

		try:
			# get the Lead from db
			lead = self.get( leadId ).first()	
			
			# get the Opportunity from db
			opportunity = OpportunityDelegate().get(convertedOpportunityId).first();
			
			# assign the ConvertedOpportunity		
			lead.convertedOpportunity = opportunity
			
			#save it
			lead.save()

			# reload and return the appropriate version					
			return self.get( leadId );
		except Lead.DoesNotExist:
			raise ProcessingError(errMsg + " : Lead with id " + str(leadId) + " does not exist.")
		except Opportunity.DoesNotExist:
			raise ProcessingError(errMsg + " : Opportunity with id " + str(convertedOpportunityId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignConvertedOpportunity( self, leadId ):
		errMsg = "Failed to unassign element " + str(convertedOpportunityId) + " for ConvertedOpportunity on Lead"

		try:
			# get the Lead from db
			lead = self.get( leadId ).first()	
			
			# assign to None for unassignment
			lead.opportunity = None			

			#save it
			lead.save()

			# reload and return the appropriate version					
			return self.get( leadId );
		except Lead.DoesNotExist:
			raise ProcessingError(errMsg + " : Lead with id " + str(leadId) + " does not exist.")
		except Exception:
			return None;
		
	def addActivities( self, leadId, activitiesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.ActivityDelegate import ActivityDelegate

		errMsg = "Failed to add elements " + str(activitiesIds) + " for Activities on Lead"

		try:
			# get the Lead
			lead = self.get( leadId ).first()
				
			# split on a comma with no spaces
			idList = activitiesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Activity		
				activity = ActivityDelegate().get(id).first();	
				# add the Activity
				lead.activities.add(activity)
				
			# save it		
			lead.save()
			
			# reload and return the appropriate version
			return self.get( leadId );
		except Lead.DoesNotExist:
			raise ProcessingError(errMsg + " : Lead with id " + str(leadId) + " does not exist.")
		except Activity.DoesNotExist:
			raise ProcessingError(errMsg + " : Activity does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeActivities( self, leadId, activitiesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.ActivityDelegate import ActivityDelegate

		errMsg = "Failed to remove elements " + str(activitiesIds) + " for Activities on Lead"

		try:
			# get the Lead
			lead = self.get( leadId ).first()
				
			# split on a comma with no spaces
			idList = activitiesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Activity		
				activity = ActivityDelegate().get(id).first();	
				# add the Activity
				lead.activities.remove(activity)
				
			# save it		
			lead.save()
			
			# reload and return the appropriate version
			return self.get( leadId );
		except Lead.DoesNotExist:
			raise ProcessingError(errMsg + " : Lead with id " + str(leadId) + " does not exist.")
		except Activity.DoesNotExist:
			raise ProcessingError(errMsg + " : Activity does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addCampaigns( self, leadId, campaignsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.CampaignDelegate import CampaignDelegate

		errMsg = "Failed to add elements " + str(campaignsIds) + " for Campaigns on Lead"

		try:
			# get the Lead
			lead = self.get( leadId ).first()
				
			# split on a comma with no spaces
			idList = campaignsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Campaign		
				campaign = CampaignDelegate().get(id).first();	
				# add the Campaign
				lead.campaigns.add(campaign)
				
			# save it		
			lead.save()
			
			# reload and return the appropriate version
			return self.get( leadId );
		except Lead.DoesNotExist:
			raise ProcessingError(errMsg + " : Lead with id " + str(leadId) + " does not exist.")
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCampaigns( self, leadId, campaignsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.CampaignDelegate import CampaignDelegate

		errMsg = "Failed to remove elements " + str(campaignsIds) + " for Campaigns on Lead"

		try:
			# get the Lead
			lead = self.get( leadId ).first()
				
			# split on a comma with no spaces
			idList = campaignsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Campaign		
				campaign = CampaignDelegate().get(id).first();	
				# add the Campaign
				lead.campaigns.remove(campaign)
				
			# save it		
			lead.save()
			
			# reload and return the appropriate version
			return self.get( leadId );
		except Lead.DoesNotExist:
			raise ProcessingError(errMsg + " : Lead with id " + str(leadId) + " does not exist.")
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addNotes( self, leadId, notesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.NoteDelegate import NoteDelegate

		errMsg = "Failed to add elements " + str(notesIds) + " for Notes on Lead"

		try:
			# get the Lead
			lead = self.get( leadId ).first()
				
			# split on a comma with no spaces
			idList = notesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Note		
				note = NoteDelegate().get(id).first();	
				# add the Note
				lead.notes.add(note)
				
			# save it		
			lead.save()
			
			# reload and return the appropriate version
			return self.get( leadId );
		except Lead.DoesNotExist:
			raise ProcessingError(errMsg + " : Lead with id " + str(leadId) + " does not exist.")
		except Note.DoesNotExist:
			raise ProcessingError(errMsg + " : Note does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeNotes( self, leadId, notesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.NoteDelegate import NoteDelegate

		errMsg = "Failed to remove elements " + str(notesIds) + " for Notes on Lead"

		try:
			# get the Lead
			lead = self.get( leadId ).first()
				
			# split on a comma with no spaces
			idList = notesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Note		
				note = NoteDelegate().get(id).first();	
				# add the Note
				lead.notes.remove(note)
				
			# save it		
			lead.save()
			
			# reload and return the appropriate version
			return self.get( leadId );
		except Lead.DoesNotExist:
			raise ProcessingError(errMsg + " : Lead with id " + str(leadId) + " does not exist.")
		except Note.DoesNotExist:
			raise ProcessingError(errMsg + " : Note does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addEmailMessages( self, leadId, emailMessagesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.EmailMessageDelegate import EmailMessageDelegate

		errMsg = "Failed to add elements " + str(emailMessagesIds) + " for EmailMessages on Lead"

		try:
			# get the Lead
			lead = self.get( leadId ).first()
				
			# split on a comma with no spaces
			idList = emailMessagesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the EmailMessage		
				emailMessage = EmailMessageDelegate().get(id).first();	
				# add the EmailMessage
				lead.emailMessages.add(emailMessage)
				
			# save it		
			lead.save()
			
			# reload and return the appropriate version
			return self.get( leadId );
		except Lead.DoesNotExist:
			raise ProcessingError(errMsg + " : Lead with id " + str(leadId) + " does not exist.")
		except EmailMessage.DoesNotExist:
			raise ProcessingError(errMsg + " : EmailMessage does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeEmailMessages( self, leadId, emailMessagesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.EmailMessageDelegate import EmailMessageDelegate

		errMsg = "Failed to remove elements " + str(emailMessagesIds) + " for EmailMessages on Lead"

		try:
			# get the Lead
			lead = self.get( leadId ).first()
				
			# split on a comma with no spaces
			idList = emailMessagesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the EmailMessage		
				emailMessage = EmailMessageDelegate().get(id).first();	
				# add the EmailMessage
				lead.emailMessages.remove(emailMessage)
				
			# save it		
			lead.save()
			
			# reload and return the appropriate version
			return self.get( leadId );
		except Lead.DoesNotExist:
			raise ProcessingError(errMsg + " : Lead with id " + str(leadId) + " does not exist.")
		except EmailMessage.DoesNotExist:
			raise ProcessingError(errMsg + " : EmailMessage does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
