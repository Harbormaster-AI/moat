from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from crmOnDjango.models.Case_ import Case_
from crmOnDjango.models.Organization import Organization
from crmOnDjango.models.Account import Account
from crmOnDjango.models.Contact import Contact
from crmOnDjango.models.User import User
from crmOnDjango.models.Team import Team
from crmOnDjango.models.Activity import Activity
from crmOnDjango.models.Note import Note
from crmOnDjango.models.EmailMessage import EmailMessage
from crmOnDjango.models.Opportunity import Opportunity
from crmOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Case_
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Case_Delegate Declaration
#======================================================================
class Case_Delegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, case_Id ):
		try:	
			case_ = Case_.objects.filter(id=case_Id)
			return case_.first();
		except Case_.DoesNotExist:
			raise ProcessingError("Case_ with id " + str(case_Id) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, case_):
		for model in serializers.deserialize("json", case_):
			model.save()
			return model;

	def create(self, case_):
		case_.save()
		return case_;

	def saveFromJson(self, case_):
		for model in serializers.deserialize("json", case_):
			model.save()
			return case_;
	
	def save(self, case_):
		case_.save()
		return case_;
	
	def delete(self, case_Id ):
		errMsg = "Failed to delete Case_ from db using id " + str(case_Id)
		
		try:
			case_ = Case_.objects.get(id=case_Id)
			case_.delete()
			return True
		except Case_.DoesNotExist:
			raise ProcessingError("Case_ with id " + str(case_Id) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Case_.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Case_ from db")
		except Exception:
			return None;
		
	def assignOrganization( self, case_Id, organizationId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OrganizationDelegate import OrganizationDelegate

		errMsg = "Failed to assign element " + str(organizationId) + " for Organization on Case_"

		try:
			# get the Case_ from db
			case_ = self.get( case_Id ).first()	
			
			# get the Organization from db
			organization = OrganizationDelegate().get(organizationId).first();
			
			# assign the Organization		
			case_.organization = organization
			
			#save it
			case_.save()

			# reload and return the appropriate version					
			return self.get( case_Id );
		except Case_.DoesNotExist:
			raise ProcessingError(errMsg + " : Case_ with id " + str(case_Id) + " does not exist.")
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrganization( self, case_Id ):
		errMsg = "Failed to unassign element " + str(organizationId) + " for Organization on Case_"

		try:
			# get the Case_ from db
			case_ = self.get( case_Id ).first()	
			
			# assign to None for unassignment
			case_.organization = None			

			#save it
			case_.save()

			# reload and return the appropriate version					
			return self.get( case_Id );
		except Case_.DoesNotExist:
			raise ProcessingError(errMsg + " : Case_ with id " + str(case_Id) + " does not exist.")
		except Exception:
			return None;
		
	def assignAccount( self, case_Id, accountId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.AccountDelegate import AccountDelegate

		errMsg = "Failed to assign element " + str(accountId) + " for Account on Case_"

		try:
			# get the Case_ from db
			case_ = self.get( case_Id ).first()	
			
			# get the Account from db
			account = AccountDelegate().get(accountId).first();
			
			# assign the Account		
			case_.account = account
			
			#save it
			case_.save()

			# reload and return the appropriate version					
			return self.get( case_Id );
		except Case_.DoesNotExist:
			raise ProcessingError(errMsg + " : Case_ with id " + str(case_Id) + " does not exist.")
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignAccount( self, case_Id ):
		errMsg = "Failed to unassign element " + str(accountId) + " for Account on Case_"

		try:
			# get the Case_ from db
			case_ = self.get( case_Id ).first()	
			
			# assign to None for unassignment
			case_.account = None			

			#save it
			case_.save()

			# reload and return the appropriate version					
			return self.get( case_Id );
		except Case_.DoesNotExist:
			raise ProcessingError(errMsg + " : Case_ with id " + str(case_Id) + " does not exist.")
		except Exception:
			return None;
		
	def assignContact( self, case_Id, contactId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.ContactDelegate import ContactDelegate

		errMsg = "Failed to assign element " + str(contactId) + " for Contact on Case_"

		try:
			# get the Case_ from db
			case_ = self.get( case_Id ).first()	
			
			# get the Contact from db
			contact = ContactDelegate().get(contactId).first();
			
			# assign the Contact		
			case_.contact = contact
			
			#save it
			case_.save()

			# reload and return the appropriate version					
			return self.get( case_Id );
		except Case_.DoesNotExist:
			raise ProcessingError(errMsg + " : Case_ with id " + str(case_Id) + " does not exist.")
		except Contact.DoesNotExist:
			raise ProcessingError(errMsg + " : Contact with id " + str(contactId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignContact( self, case_Id ):
		errMsg = "Failed to unassign element " + str(contactId) + " for Contact on Case_"

		try:
			# get the Case_ from db
			case_ = self.get( case_Id ).first()	
			
			# assign to None for unassignment
			case_.contact = None			

			#save it
			case_.save()

			# reload and return the appropriate version					
			return self.get( case_Id );
		except Case_.DoesNotExist:
			raise ProcessingError(errMsg + " : Case_ with id " + str(case_Id) + " does not exist.")
		except Exception:
			return None;
		
	def assignOwner( self, case_Id, ownerId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.UserDelegate import UserDelegate

		errMsg = "Failed to assign element " + str(ownerId) + " for Owner on Case_"

		try:
			# get the Case_ from db
			case_ = self.get( case_Id ).first()	
			
			# get the User from db
			user = UserDelegate().get(ownerId).first();
			
			# assign the Owner		
			case_.owner = user
			
			#save it
			case_.save()

			# reload and return the appropriate version					
			return self.get( case_Id );
		except Case_.DoesNotExist:
			raise ProcessingError(errMsg + " : Case_ with id " + str(case_Id) + " does not exist.")
		except User.DoesNotExist:
			raise ProcessingError(errMsg + " : User with id " + str(ownerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOwner( self, case_Id ):
		errMsg = "Failed to unassign element " + str(ownerId) + " for Owner on Case_"

		try:
			# get the Case_ from db
			case_ = self.get( case_Id ).first()	
			
			# assign to None for unassignment
			case_.user = None			

			#save it
			case_.save()

			# reload and return the appropriate version					
			return self.get( case_Id );
		except Case_.DoesNotExist:
			raise ProcessingError(errMsg + " : Case_ with id " + str(case_Id) + " does not exist.")
		except Exception:
			return None;
		
	def assignTeam( self, case_Id, teamId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.TeamDelegate import TeamDelegate

		errMsg = "Failed to assign element " + str(teamId) + " for Team on Case_"

		try:
			# get the Case_ from db
			case_ = self.get( case_Id ).first()	
			
			# get the Team from db
			team = TeamDelegate().get(teamId).first();
			
			# assign the Team		
			case_.team = team
			
			#save it
			case_.save()

			# reload and return the appropriate version					
			return self.get( case_Id );
		except Case_.DoesNotExist:
			raise ProcessingError(errMsg + " : Case_ with id " + str(case_Id) + " does not exist.")
		except Team.DoesNotExist:
			raise ProcessingError(errMsg + " : Team with id " + str(teamId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignTeam( self, case_Id ):
		errMsg = "Failed to unassign element " + str(teamId) + " for Team on Case_"

		try:
			# get the Case_ from db
			case_ = self.get( case_Id ).first()	
			
			# assign to None for unassignment
			case_.team = None			

			#save it
			case_.save()

			# reload and return the appropriate version					
			return self.get( case_Id );
		except Case_.DoesNotExist:
			raise ProcessingError(errMsg + " : Case_ with id " + str(case_Id) + " does not exist.")
		except Exception:
			return None;
		
	def addActivities( self, case_Id, activitiesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.ActivityDelegate import ActivityDelegate

		errMsg = "Failed to add elements " + str(activitiesIds) + " for Activities on Case_"

		try:
			# get the Case_
			case_ = self.get( case_Id ).first()
				
			# split on a comma with no spaces
			idList = activitiesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Activity		
				activity = ActivityDelegate().get(id).first();	
				# add the Activity
				case_.activities.add(activity)
				
			# save it		
			case_.save()
			
			# reload and return the appropriate version
			return self.get( case_Id );
		except Case_.DoesNotExist:
			raise ProcessingError(errMsg + " : Case_ with id " + str(case_Id) + " does not exist.")
		except Activity.DoesNotExist:
			raise ProcessingError(errMsg + " : Activity does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeActivities( self, case_Id, activitiesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.ActivityDelegate import ActivityDelegate

		errMsg = "Failed to remove elements " + str(activitiesIds) + " for Activities on Case_"

		try:
			# get the Case_
			case_ = self.get( case_Id ).first()
				
			# split on a comma with no spaces
			idList = activitiesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Activity		
				activity = ActivityDelegate().get(id).first();	
				# add the Activity
				case_.activities.remove(activity)
				
			# save it		
			case_.save()
			
			# reload and return the appropriate version
			return self.get( case_Id );
		except Case_.DoesNotExist:
			raise ProcessingError(errMsg + " : Case_ with id " + str(case_Id) + " does not exist.")
		except Activity.DoesNotExist:
			raise ProcessingError(errMsg + " : Activity does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addCaseComments( self, case_Id, caseCommentsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.NoteDelegate import NoteDelegate

		errMsg = "Failed to add elements " + str(caseCommentsIds) + " for CaseComments on Case_"

		try:
			# get the Case_
			case_ = self.get( case_Id ).first()
				
			# split on a comma with no spaces
			idList = caseCommentsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Note		
				note = NoteDelegate().get(id).first();	
				# add the Note
				case_.caseComments.add(note)
				
			# save it		
			case_.save()
			
			# reload and return the appropriate version
			return self.get( case_Id );
		except Case_.DoesNotExist:
			raise ProcessingError(errMsg + " : Case_ with id " + str(case_Id) + " does not exist.")
		except Note.DoesNotExist:
			raise ProcessingError(errMsg + " : Note does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCaseComments( self, case_Id, caseCommentsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.NoteDelegate import NoteDelegate

		errMsg = "Failed to remove elements " + str(caseCommentsIds) + " for CaseComments on Case_"

		try:
			# get the Case_
			case_ = self.get( case_Id ).first()
				
			# split on a comma with no spaces
			idList = caseCommentsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Note		
				note = NoteDelegate().get(id).first();	
				# add the Note
				case_.caseComments.remove(note)
				
			# save it		
			case_.save()
			
			# reload and return the appropriate version
			return self.get( case_Id );
		except Case_.DoesNotExist:
			raise ProcessingError(errMsg + " : Case_ with id " + str(case_Id) + " does not exist.")
		except Note.DoesNotExist:
			raise ProcessingError(errMsg + " : Note does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addEmails( self, case_Id, emailsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.EmailMessageDelegate import EmailMessageDelegate

		errMsg = "Failed to add elements " + str(emailsIds) + " for Emails on Case_"

		try:
			# get the Case_
			case_ = self.get( case_Id ).first()
				
			# split on a comma with no spaces
			idList = emailsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the EmailMessage		
				emailMessage = EmailMessageDelegate().get(id).first();	
				# add the EmailMessage
				case_.emails.add(emailMessage)
				
			# save it		
			case_.save()
			
			# reload and return the appropriate version
			return self.get( case_Id );
		except Case_.DoesNotExist:
			raise ProcessingError(errMsg + " : Case_ with id " + str(case_Id) + " does not exist.")
		except EmailMessage.DoesNotExist:
			raise ProcessingError(errMsg + " : EmailMessage does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeEmails( self, case_Id, emailsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.EmailMessageDelegate import EmailMessageDelegate

		errMsg = "Failed to remove elements " + str(emailsIds) + " for Emails on Case_"

		try:
			# get the Case_
			case_ = self.get( case_Id ).first()
				
			# split on a comma with no spaces
			idList = emailsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the EmailMessage		
				emailMessage = EmailMessageDelegate().get(id).first();	
				# add the EmailMessage
				case_.emails.remove(emailMessage)
				
			# save it		
			case_.save()
			
			# reload and return the appropriate version
			return self.get( case_Id );
		except Case_.DoesNotExist:
			raise ProcessingError(errMsg + " : Case_ with id " + str(case_Id) + " does not exist.")
		except EmailMessage.DoesNotExist:
			raise ProcessingError(errMsg + " : EmailMessage does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addRelatedOpportunities( self, case_Id, relatedOpportunitiesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OpportunityDelegate import OpportunityDelegate

		errMsg = "Failed to add elements " + str(relatedOpportunitiesIds) + " for RelatedOpportunities on Case_"

		try:
			# get the Case_
			case_ = self.get( case_Id ).first()
				
			# split on a comma with no spaces
			idList = relatedOpportunitiesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Opportunity		
				opportunity = OpportunityDelegate().get(id).first();	
				# add the Opportunity
				case_.relatedOpportunities.add(opportunity)
				
			# save it		
			case_.save()
			
			# reload and return the appropriate version
			return self.get( case_Id );
		except Case_.DoesNotExist:
			raise ProcessingError(errMsg + " : Case_ with id " + str(case_Id) + " does not exist.")
		except Opportunity.DoesNotExist:
			raise ProcessingError(errMsg + " : Opportunity does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeRelatedOpportunities( self, case_Id, relatedOpportunitiesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OpportunityDelegate import OpportunityDelegate

		errMsg = "Failed to remove elements " + str(relatedOpportunitiesIds) + " for RelatedOpportunities on Case_"

		try:
			# get the Case_
			case_ = self.get( case_Id ).first()
				
			# split on a comma with no spaces
			idList = relatedOpportunitiesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Opportunity		
				opportunity = OpportunityDelegate().get(id).first();	
				# add the Opportunity
				case_.relatedOpportunities.remove(opportunity)
				
			# save it		
			case_.save()
			
			# reload and return the appropriate version
			return self.get( case_Id );
		except Case_.DoesNotExist:
			raise ProcessingError(errMsg + " : Case_ with id " + str(case_Id) + " does not exist.")
		except Opportunity.DoesNotExist:
			raise ProcessingError(errMsg + " : Opportunity does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
