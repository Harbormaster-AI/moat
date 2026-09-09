from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from crmOnDjango.models.Note import Note
from crmOnDjango.models.Organization import Organization
from crmOnDjango.models.User import User
from crmOnDjango.models.Account import Account
from crmOnDjango.models.Contact import Contact
from crmOnDjango.models.Opportunity import Opportunity
from crmOnDjango.models.Case_ import Case_
from crmOnDjango.models.Lead import Lead
from crmOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Note
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class NoteDelegate Declaration
#======================================================================
class NoteDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, noteId ):
		try:	
			note = Note.objects.filter(id=noteId)
			return note.first();
		except Note.DoesNotExist:
			raise ProcessingError("Note with id " + str(noteId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, note):
		for model in serializers.deserialize("json", note):
			model.save()
			return model;

	def create(self, note):
		note.save()
		return note;

	def saveFromJson(self, note):
		for model in serializers.deserialize("json", note):
			model.save()
			return note;
	
	def save(self, note):
		note.save()
		return note;
	
	def delete(self, noteId ):
		errMsg = "Failed to delete Note from db using id " + str(noteId)
		
		try:
			note = Note.objects.get(id=noteId)
			note.delete()
			return True
		except Note.DoesNotExist:
			raise ProcessingError("Note with id " + str(noteId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Note.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Note from db")
		except Exception:
			return None;
		
	def assignOrganization( self, noteId, organizationId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OrganizationDelegate import OrganizationDelegate

		errMsg = "Failed to assign element " + str(organizationId) + " for Organization on Note"

		try:
			# get the Note from db
			note = self.get( noteId ).first()	
			
			# get the Organization from db
			organization = OrganizationDelegate().get(organizationId).first();
			
			# assign the Organization		
			note.organization = organization
			
			#save it
			note.save()

			# reload and return the appropriate version					
			return self.get( noteId );
		except Note.DoesNotExist:
			raise ProcessingError(errMsg + " : Note with id " + str(noteId) + " does not exist.")
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrganization( self, noteId ):
		errMsg = "Failed to unassign element " + str(organizationId) + " for Organization on Note"

		try:
			# get the Note from db
			note = self.get( noteId ).first()	
			
			# assign to None for unassignment
			note.organization = None			

			#save it
			note.save()

			# reload and return the appropriate version					
			return self.get( noteId );
		except Note.DoesNotExist:
			raise ProcessingError(errMsg + " : Note with id " + str(noteId) + " does not exist.")
		except Exception:
			return None;
		
	def assignOwner( self, noteId, ownerId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.UserDelegate import UserDelegate

		errMsg = "Failed to assign element " + str(ownerId) + " for Owner on Note"

		try:
			# get the Note from db
			note = self.get( noteId ).first()	
			
			# get the User from db
			user = UserDelegate().get(ownerId).first();
			
			# assign the Owner		
			note.owner = user
			
			#save it
			note.save()

			# reload and return the appropriate version					
			return self.get( noteId );
		except Note.DoesNotExist:
			raise ProcessingError(errMsg + " : Note with id " + str(noteId) + " does not exist.")
		except User.DoesNotExist:
			raise ProcessingError(errMsg + " : User with id " + str(ownerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOwner( self, noteId ):
		errMsg = "Failed to unassign element " + str(ownerId) + " for Owner on Note"

		try:
			# get the Note from db
			note = self.get( noteId ).first()	
			
			# assign to None for unassignment
			note.user = None			

			#save it
			note.save()

			# reload and return the appropriate version					
			return self.get( noteId );
		except Note.DoesNotExist:
			raise ProcessingError(errMsg + " : Note with id " + str(noteId) + " does not exist.")
		except Exception:
			return None;
		
	def assignAccount( self, noteId, accountId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.AccountDelegate import AccountDelegate

		errMsg = "Failed to assign element " + str(accountId) + " for Account on Note"

		try:
			# get the Note from db
			note = self.get( noteId ).first()	
			
			# get the Account from db
			account = AccountDelegate().get(accountId).first();
			
			# assign the Account		
			note.account = account
			
			#save it
			note.save()

			# reload and return the appropriate version					
			return self.get( noteId );
		except Note.DoesNotExist:
			raise ProcessingError(errMsg + " : Note with id " + str(noteId) + " does not exist.")
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignAccount( self, noteId ):
		errMsg = "Failed to unassign element " + str(accountId) + " for Account on Note"

		try:
			# get the Note from db
			note = self.get( noteId ).first()	
			
			# assign to None for unassignment
			note.account = None			

			#save it
			note.save()

			# reload and return the appropriate version					
			return self.get( noteId );
		except Note.DoesNotExist:
			raise ProcessingError(errMsg + " : Note with id " + str(noteId) + " does not exist.")
		except Exception:
			return None;
		
	def assignContact( self, noteId, contactId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.ContactDelegate import ContactDelegate

		errMsg = "Failed to assign element " + str(contactId) + " for Contact on Note"

		try:
			# get the Note from db
			note = self.get( noteId ).first()	
			
			# get the Contact from db
			contact = ContactDelegate().get(contactId).first();
			
			# assign the Contact		
			note.contact = contact
			
			#save it
			note.save()

			# reload and return the appropriate version					
			return self.get( noteId );
		except Note.DoesNotExist:
			raise ProcessingError(errMsg + " : Note with id " + str(noteId) + " does not exist.")
		except Contact.DoesNotExist:
			raise ProcessingError(errMsg + " : Contact with id " + str(contactId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignContact( self, noteId ):
		errMsg = "Failed to unassign element " + str(contactId) + " for Contact on Note"

		try:
			# get the Note from db
			note = self.get( noteId ).first()	
			
			# assign to None for unassignment
			note.contact = None			

			#save it
			note.save()

			# reload and return the appropriate version					
			return self.get( noteId );
		except Note.DoesNotExist:
			raise ProcessingError(errMsg + " : Note with id " + str(noteId) + " does not exist.")
		except Exception:
			return None;
		
	def assignOpportunity( self, noteId, opportunityId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OpportunityDelegate import OpportunityDelegate

		errMsg = "Failed to assign element " + str(opportunityId) + " for Opportunity on Note"

		try:
			# get the Note from db
			note = self.get( noteId ).first()	
			
			# get the Opportunity from db
			opportunity = OpportunityDelegate().get(opportunityId).first();
			
			# assign the Opportunity		
			note.opportunity = opportunity
			
			#save it
			note.save()

			# reload and return the appropriate version					
			return self.get( noteId );
		except Note.DoesNotExist:
			raise ProcessingError(errMsg + " : Note with id " + str(noteId) + " does not exist.")
		except Opportunity.DoesNotExist:
			raise ProcessingError(errMsg + " : Opportunity with id " + str(opportunityId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOpportunity( self, noteId ):
		errMsg = "Failed to unassign element " + str(opportunityId) + " for Opportunity on Note"

		try:
			# get the Note from db
			note = self.get( noteId ).first()	
			
			# assign to None for unassignment
			note.opportunity = None			

			#save it
			note.save()

			# reload and return the appropriate version					
			return self.get( noteId );
		except Note.DoesNotExist:
			raise ProcessingError(errMsg + " : Note with id " + str(noteId) + " does not exist.")
		except Exception:
			return None;
		
	def assignCase( self, noteId, caseId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.Case_Delegate import Case_Delegate

		errMsg = "Failed to assign element " + str(caseId) + " for Case on Note"

		try:
			# get the Note from db
			note = self.get( noteId ).first()	
			
			# get the Case_ from db
			case_ = Case_Delegate().get(caseId).first();
			
			# assign the Case		
			note.case = case_
			
			#save it
			note.save()

			# reload and return the appropriate version					
			return self.get( noteId );
		except Note.DoesNotExist:
			raise ProcessingError(errMsg + " : Note with id " + str(noteId) + " does not exist.")
		except Case_.DoesNotExist:
			raise ProcessingError(errMsg + " : Case_ with id " + str(caseId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCase( self, noteId ):
		errMsg = "Failed to unassign element " + str(caseId) + " for Case on Note"

		try:
			# get the Note from db
			note = self.get( noteId ).first()	
			
			# assign to None for unassignment
			note.case_ = None			

			#save it
			note.save()

			# reload and return the appropriate version					
			return self.get( noteId );
		except Note.DoesNotExist:
			raise ProcessingError(errMsg + " : Note with id " + str(noteId) + " does not exist.")
		except Exception:
			return None;
		
	def assignLead( self, noteId, leadId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.LeadDelegate import LeadDelegate

		errMsg = "Failed to assign element " + str(leadId) + " for Lead on Note"

		try:
			# get the Note from db
			note = self.get( noteId ).first()	
			
			# get the Lead from db
			lead = LeadDelegate().get(leadId).first();
			
			# assign the Lead		
			note.lead = lead
			
			#save it
			note.save()

			# reload and return the appropriate version					
			return self.get( noteId );
		except Note.DoesNotExist:
			raise ProcessingError(errMsg + " : Note with id " + str(noteId) + " does not exist.")
		except Lead.DoesNotExist:
			raise ProcessingError(errMsg + " : Lead with id " + str(leadId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignLead( self, noteId ):
		errMsg = "Failed to unassign element " + str(leadId) + " for Lead on Note"

		try:
			# get the Note from db
			note = self.get( noteId ).first()	
			
			# assign to None for unassignment
			note.lead = None			

			#save it
			note.save()

			# reload and return the appropriate version					
			return self.get( noteId );
		except Note.DoesNotExist:
			raise ProcessingError(errMsg + " : Note with id " + str(noteId) + " does not exist.")
		except Exception:
			return None;
		
