from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from crmOnDjango.models.Contact import Contact
from crmOnDjango.models.Organization import Organization
from crmOnDjango.models.Account import Account
from crmOnDjango.models.User import User
from crmOnDjango.models.Activity import Activity
from crmOnDjango.models.Opportunity import Opportunity
from crmOnDjango.models.Case_ import Case_
from crmOnDjango.models.Campaign import Campaign
from crmOnDjango.models.Note import Note
from crmOnDjango.models.EmailMessage import EmailMessage
from crmOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Contact
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ContactDelegate Declaration
#======================================================================
class ContactDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, contactId ):
		try:	
			contact = Contact.objects.filter(id=contactId)
			return contact.first();
		except Contact.DoesNotExist:
			raise ProcessingError("Contact with id " + str(contactId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, contact):
		for model in serializers.deserialize("json", contact):
			model.save()
			return model;

	def create(self, contact):
		contact.save()
		return contact;

	def saveFromJson(self, contact):
		for model in serializers.deserialize("json", contact):
			model.save()
			return contact;
	
	def save(self, contact):
		contact.save()
		return contact;
	
	def delete(self, contactId ):
		errMsg = "Failed to delete Contact from db using id " + str(contactId)
		
		try:
			contact = Contact.objects.get(id=contactId)
			contact.delete()
			return True
		except Contact.DoesNotExist:
			raise ProcessingError("Contact with id " + str(contactId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Contact.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Contact from db")
		except Exception:
			return None;
		
	def assignOrganization( self, contactId, organizationId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OrganizationDelegate import OrganizationDelegate

		errMsg = "Failed to assign element " + str(organizationId) + " for Organization on Contact"

		try:
			# get the Contact from db
			contact = self.get( contactId ).first()	
			
			# get the Organization from db
			organization = OrganizationDelegate().get(organizationId).first();
			
			# assign the Organization		
			contact.organization = organization
			
			#save it
			contact.save()

			# reload and return the appropriate version					
			return self.get( contactId );
		except Contact.DoesNotExist:
			raise ProcessingError(errMsg + " : Contact with id " + str(contactId) + " does not exist.")
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrganization( self, contactId ):
		errMsg = "Failed to unassign element " + str(organizationId) + " for Organization on Contact"

		try:
			# get the Contact from db
			contact = self.get( contactId ).first()	
			
			# assign to None for unassignment
			contact.organization = None			

			#save it
			contact.save()

			# reload and return the appropriate version					
			return self.get( contactId );
		except Contact.DoesNotExist:
			raise ProcessingError(errMsg + " : Contact with id " + str(contactId) + " does not exist.")
		except Exception:
			return None;
		
	def assignAccount( self, contactId, accountId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.AccountDelegate import AccountDelegate

		errMsg = "Failed to assign element " + str(accountId) + " for Account on Contact"

		try:
			# get the Contact from db
			contact = self.get( contactId ).first()	
			
			# get the Account from db
			account = AccountDelegate().get(accountId).first();
			
			# assign the Account		
			contact.account = account
			
			#save it
			contact.save()

			# reload and return the appropriate version					
			return self.get( contactId );
		except Contact.DoesNotExist:
			raise ProcessingError(errMsg + " : Contact with id " + str(contactId) + " does not exist.")
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignAccount( self, contactId ):
		errMsg = "Failed to unassign element " + str(accountId) + " for Account on Contact"

		try:
			# get the Contact from db
			contact = self.get( contactId ).first()	
			
			# assign to None for unassignment
			contact.account = None			

			#save it
			contact.save()

			# reload and return the appropriate version					
			return self.get( contactId );
		except Contact.DoesNotExist:
			raise ProcessingError(errMsg + " : Contact with id " + str(contactId) + " does not exist.")
		except Exception:
			return None;
		
	def assignOwner( self, contactId, ownerId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.UserDelegate import UserDelegate

		errMsg = "Failed to assign element " + str(ownerId) + " for Owner on Contact"

		try:
			# get the Contact from db
			contact = self.get( contactId ).first()	
			
			# get the User from db
			user = UserDelegate().get(ownerId).first();
			
			# assign the Owner		
			contact.owner = user
			
			#save it
			contact.save()

			# reload and return the appropriate version					
			return self.get( contactId );
		except Contact.DoesNotExist:
			raise ProcessingError(errMsg + " : Contact with id " + str(contactId) + " does not exist.")
		except User.DoesNotExist:
			raise ProcessingError(errMsg + " : User with id " + str(ownerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOwner( self, contactId ):
		errMsg = "Failed to unassign element " + str(ownerId) + " for Owner on Contact"

		try:
			# get the Contact from db
			contact = self.get( contactId ).first()	
			
			# assign to None for unassignment
			contact.user = None			

			#save it
			contact.save()

			# reload and return the appropriate version					
			return self.get( contactId );
		except Contact.DoesNotExist:
			raise ProcessingError(errMsg + " : Contact with id " + str(contactId) + " does not exist.")
		except Exception:
			return None;
		
	def addActivities( self, contactId, activitiesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.ActivityDelegate import ActivityDelegate

		errMsg = "Failed to add elements " + str(activitiesIds) + " for Activities on Contact"

		try:
			# get the Contact
			contact = self.get( contactId ).first()
				
			# split on a comma with no spaces
			idList = activitiesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Activity		
				activity = ActivityDelegate().get(id).first();	
				# add the Activity
				contact.activities.add(activity)
				
			# save it		
			contact.save()
			
			# reload and return the appropriate version
			return self.get( contactId );
		except Contact.DoesNotExist:
			raise ProcessingError(errMsg + " : Contact with id " + str(contactId) + " does not exist.")
		except Activity.DoesNotExist:
			raise ProcessingError(errMsg + " : Activity does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeActivities( self, contactId, activitiesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.ActivityDelegate import ActivityDelegate

		errMsg = "Failed to remove elements " + str(activitiesIds) + " for Activities on Contact"

		try:
			# get the Contact
			contact = self.get( contactId ).first()
				
			# split on a comma with no spaces
			idList = activitiesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Activity		
				activity = ActivityDelegate().get(id).first();	
				# add the Activity
				contact.activities.remove(activity)
				
			# save it		
			contact.save()
			
			# reload and return the appropriate version
			return self.get( contactId );
		except Contact.DoesNotExist:
			raise ProcessingError(errMsg + " : Contact with id " + str(contactId) + " does not exist.")
		except Activity.DoesNotExist:
			raise ProcessingError(errMsg + " : Activity does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addOpportunities( self, contactId, opportunitiesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OpportunityDelegate import OpportunityDelegate

		errMsg = "Failed to add elements " + str(opportunitiesIds) + " for Opportunities on Contact"

		try:
			# get the Contact
			contact = self.get( contactId ).first()
				
			# split on a comma with no spaces
			idList = opportunitiesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Opportunity		
				opportunity = OpportunityDelegate().get(id).first();	
				# add the Opportunity
				contact.opportunities.add(opportunity)
				
			# save it		
			contact.save()
			
			# reload and return the appropriate version
			return self.get( contactId );
		except Contact.DoesNotExist:
			raise ProcessingError(errMsg + " : Contact with id " + str(contactId) + " does not exist.")
		except Opportunity.DoesNotExist:
			raise ProcessingError(errMsg + " : Opportunity does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeOpportunities( self, contactId, opportunitiesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OpportunityDelegate import OpportunityDelegate

		errMsg = "Failed to remove elements " + str(opportunitiesIds) + " for Opportunities on Contact"

		try:
			# get the Contact
			contact = self.get( contactId ).first()
				
			# split on a comma with no spaces
			idList = opportunitiesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Opportunity		
				opportunity = OpportunityDelegate().get(id).first();	
				# add the Opportunity
				contact.opportunities.remove(opportunity)
				
			# save it		
			contact.save()
			
			# reload and return the appropriate version
			return self.get( contactId );
		except Contact.DoesNotExist:
			raise ProcessingError(errMsg + " : Contact with id " + str(contactId) + " does not exist.")
		except Opportunity.DoesNotExist:
			raise ProcessingError(errMsg + " : Opportunity does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addCases( self, contactId, casesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.Case_Delegate import Case_Delegate

		errMsg = "Failed to add elements " + str(casesIds) + " for Cases on Contact"

		try:
			# get the Contact
			contact = self.get( contactId ).first()
				
			# split on a comma with no spaces
			idList = casesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Case_		
				case_ = Case_Delegate().get(id).first();	
				# add the Case_
				contact.cases.add(case_)
				
			# save it		
			contact.save()
			
			# reload and return the appropriate version
			return self.get( contactId );
		except Contact.DoesNotExist:
			raise ProcessingError(errMsg + " : Contact with id " + str(contactId) + " does not exist.")
		except Case_.DoesNotExist:
			raise ProcessingError(errMsg + " : Case_ does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCases( self, contactId, casesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.Case_Delegate import Case_Delegate

		errMsg = "Failed to remove elements " + str(casesIds) + " for Cases on Contact"

		try:
			# get the Contact
			contact = self.get( contactId ).first()
				
			# split on a comma with no spaces
			idList = casesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Case_		
				case_ = Case_Delegate().get(id).first();	
				# add the Case_
				contact.cases.remove(case_)
				
			# save it		
			contact.save()
			
			# reload and return the appropriate version
			return self.get( contactId );
		except Contact.DoesNotExist:
			raise ProcessingError(errMsg + " : Contact with id " + str(contactId) + " does not exist.")
		except Case_.DoesNotExist:
			raise ProcessingError(errMsg + " : Case_ does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addCampaigns( self, contactId, campaignsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.CampaignDelegate import CampaignDelegate

		errMsg = "Failed to add elements " + str(campaignsIds) + " for Campaigns on Contact"

		try:
			# get the Contact
			contact = self.get( contactId ).first()
				
			# split on a comma with no spaces
			idList = campaignsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Campaign		
				campaign = CampaignDelegate().get(id).first();	
				# add the Campaign
				contact.campaigns.add(campaign)
				
			# save it		
			contact.save()
			
			# reload and return the appropriate version
			return self.get( contactId );
		except Contact.DoesNotExist:
			raise ProcessingError(errMsg + " : Contact with id " + str(contactId) + " does not exist.")
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCampaigns( self, contactId, campaignsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.CampaignDelegate import CampaignDelegate

		errMsg = "Failed to remove elements " + str(campaignsIds) + " for Campaigns on Contact"

		try:
			# get the Contact
			contact = self.get( contactId ).first()
				
			# split on a comma with no spaces
			idList = campaignsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Campaign		
				campaign = CampaignDelegate().get(id).first();	
				# add the Campaign
				contact.campaigns.remove(campaign)
				
			# save it		
			contact.save()
			
			# reload and return the appropriate version
			return self.get( contactId );
		except Contact.DoesNotExist:
			raise ProcessingError(errMsg + " : Contact with id " + str(contactId) + " does not exist.")
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addNotes( self, contactId, notesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.NoteDelegate import NoteDelegate

		errMsg = "Failed to add elements " + str(notesIds) + " for Notes on Contact"

		try:
			# get the Contact
			contact = self.get( contactId ).first()
				
			# split on a comma with no spaces
			idList = notesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Note		
				note = NoteDelegate().get(id).first();	
				# add the Note
				contact.notes.add(note)
				
			# save it		
			contact.save()
			
			# reload and return the appropriate version
			return self.get( contactId );
		except Contact.DoesNotExist:
			raise ProcessingError(errMsg + " : Contact with id " + str(contactId) + " does not exist.")
		except Note.DoesNotExist:
			raise ProcessingError(errMsg + " : Note does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeNotes( self, contactId, notesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.NoteDelegate import NoteDelegate

		errMsg = "Failed to remove elements " + str(notesIds) + " for Notes on Contact"

		try:
			# get the Contact
			contact = self.get( contactId ).first()
				
			# split on a comma with no spaces
			idList = notesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Note		
				note = NoteDelegate().get(id).first();	
				# add the Note
				contact.notes.remove(note)
				
			# save it		
			contact.save()
			
			# reload and return the appropriate version
			return self.get( contactId );
		except Contact.DoesNotExist:
			raise ProcessingError(errMsg + " : Contact with id " + str(contactId) + " does not exist.")
		except Note.DoesNotExist:
			raise ProcessingError(errMsg + " : Note does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addEmailMessages( self, contactId, emailMessagesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.EmailMessageDelegate import EmailMessageDelegate

		errMsg = "Failed to add elements " + str(emailMessagesIds) + " for EmailMessages on Contact"

		try:
			# get the Contact
			contact = self.get( contactId ).first()
				
			# split on a comma with no spaces
			idList = emailMessagesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the EmailMessage		
				emailMessage = EmailMessageDelegate().get(id).first();	
				# add the EmailMessage
				contact.emailMessages.add(emailMessage)
				
			# save it		
			contact.save()
			
			# reload and return the appropriate version
			return self.get( contactId );
		except Contact.DoesNotExist:
			raise ProcessingError(errMsg + " : Contact with id " + str(contactId) + " does not exist.")
		except EmailMessage.DoesNotExist:
			raise ProcessingError(errMsg + " : EmailMessage does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeEmailMessages( self, contactId, emailMessagesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.EmailMessageDelegate import EmailMessageDelegate

		errMsg = "Failed to remove elements " + str(emailMessagesIds) + " for EmailMessages on Contact"

		try:
			# get the Contact
			contact = self.get( contactId ).first()
				
			# split on a comma with no spaces
			idList = emailMessagesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the EmailMessage		
				emailMessage = EmailMessageDelegate().get(id).first();	
				# add the EmailMessage
				contact.emailMessages.remove(emailMessage)
				
			# save it		
			contact.save()
			
			# reload and return the appropriate version
			return self.get( contactId );
		except Contact.DoesNotExist:
			raise ProcessingError(errMsg + " : Contact with id " + str(contactId) + " does not exist.")
		except EmailMessage.DoesNotExist:
			raise ProcessingError(errMsg + " : EmailMessage does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
