from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from crmOnDjango.models.Account import Account
from crmOnDjango.models.Organization import Organization
from crmOnDjango.models.Contact import Contact
from crmOnDjango.models.Opportunity import Opportunity
from crmOnDjango.models.Case_ import Case_
from crmOnDjango.models.User import User
from crmOnDjango.models.Territory import Territory
from crmOnDjango.models.Activity import Activity
from crmOnDjango.models.Campaign import Campaign
from crmOnDjango.models.Quote import Quote
from crmOnDjango.models.Order import Order
from crmOnDjango.models.Contract import Contract
from crmOnDjango.models.Note import Note
from crmOnDjango.models.EmailMessage import EmailMessage
from crmOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Account
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AccountDelegate Declaration
#======================================================================
class AccountDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, accountId ):
		try:	
			account = Account.objects.filter(id=accountId)
			return account.first();
		except Account.DoesNotExist:
			raise ProcessingError("Account with id " + str(accountId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, account):
		for model in serializers.deserialize("json", account):
			model.save()
			return model;

	def create(self, account):
		account.save()
		return account;

	def saveFromJson(self, account):
		for model in serializers.deserialize("json", account):
			model.save()
			return account;
	
	def save(self, account):
		account.save()
		return account;
	
	def delete(self, accountId ):
		errMsg = "Failed to delete Account from db using id " + str(accountId)
		
		try:
			account = Account.objects.get(id=accountId)
			account.delete()
			return True
		except Account.DoesNotExist:
			raise ProcessingError("Account with id " + str(accountId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Account.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Account from db")
		except Exception:
			return None;
		
	def assignOrganization( self, accountId, organizationId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OrganizationDelegate import OrganizationDelegate

		errMsg = "Failed to assign element " + str(organizationId) + " for Organization on Account"

		try:
			# get the Account from db
			account = self.get( accountId ).first()	
			
			# get the Organization from db
			organization = OrganizationDelegate().get(organizationId).first();
			
			# assign the Organization		
			account.organization = organization
			
			#save it
			account.save()

			# reload and return the appropriate version					
			return self.get( accountId );
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrganization( self, accountId ):
		errMsg = "Failed to unassign element " + str(organizationId) + " for Organization on Account"

		try:
			# get the Account from db
			account = self.get( accountId ).first()	
			
			# assign to None for unassignment
			account.organization = None			

			#save it
			account.save()

			# reload and return the appropriate version					
			return self.get( accountId );
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except Exception:
			return None;
		
	def assignParentAccount( self, accountId, parentAccountId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.AccountDelegate import AccountDelegate

		errMsg = "Failed to assign element " + str(parentAccountId) + " for ParentAccount on Account"

		try:
			# get the Account from db
			account = self.get( accountId ).first()	
			
			# get the Account from db
			account = AccountDelegate().get(parentAccountId).first();
			
			# assign the ParentAccount		
			account.parentAccount = account
			
			#save it
			account.save()

			# reload and return the appropriate version					
			return self.get( accountId );
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(parentAccountId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignParentAccount( self, accountId ):
		errMsg = "Failed to unassign element " + str(parentAccountId) + " for ParentAccount on Account"

		try:
			# get the Account from db
			account = self.get( accountId ).first()	
			
			# assign to None for unassignment
			account.account = None			

			#save it
			account.save()

			# reload and return the appropriate version					
			return self.get( accountId );
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except Exception:
			return None;
		
	def assignOwner( self, accountId, ownerId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.UserDelegate import UserDelegate

		errMsg = "Failed to assign element " + str(ownerId) + " for Owner on Account"

		try:
			# get the Account from db
			account = self.get( accountId ).first()	
			
			# get the User from db
			user = UserDelegate().get(ownerId).first();
			
			# assign the Owner		
			account.owner = user
			
			#save it
			account.save()

			# reload and return the appropriate version					
			return self.get( accountId );
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except User.DoesNotExist:
			raise ProcessingError(errMsg + " : User with id " + str(ownerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOwner( self, accountId ):
		errMsg = "Failed to unassign element " + str(ownerId) + " for Owner on Account"

		try:
			# get the Account from db
			account = self.get( accountId ).first()	
			
			# assign to None for unassignment
			account.user = None			

			#save it
			account.save()

			# reload and return the appropriate version					
			return self.get( accountId );
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except Exception:
			return None;
		
	def assignTerritory( self, accountId, territoryId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.TerritoryDelegate import TerritoryDelegate

		errMsg = "Failed to assign element " + str(territoryId) + " for Territory on Account"

		try:
			# get the Account from db
			account = self.get( accountId ).first()	
			
			# get the Territory from db
			territory = TerritoryDelegate().get(territoryId).first();
			
			# assign the Territory		
			account.territory = territory
			
			#save it
			account.save()

			# reload and return the appropriate version					
			return self.get( accountId );
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except Territory.DoesNotExist:
			raise ProcessingError(errMsg + " : Territory with id " + str(territoryId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignTerritory( self, accountId ):
		errMsg = "Failed to unassign element " + str(territoryId) + " for Territory on Account"

		try:
			# get the Account from db
			account = self.get( accountId ).first()	
			
			# assign to None for unassignment
			account.territory = None			

			#save it
			account.save()

			# reload and return the appropriate version					
			return self.get( accountId );
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except Exception:
			return None;
		
	def addChildAccounts( self, accountId, childAccountsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.AccountDelegate import AccountDelegate

		errMsg = "Failed to add elements " + str(childAccountsIds) + " for ChildAccounts on Account"

		try:
			# get the Account
			account = self.get( accountId ).first()
				
			# split on a comma with no spaces
			idList = childAccountsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Account		
				account = AccountDelegate().get(id).first();	
				# add the Account
				account.childAccounts.add(account)
				
			# save it		
			account.save()
			
			# reload and return the appropriate version
			return self.get( accountId );
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeChildAccounts( self, accountId, childAccountsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.AccountDelegate import AccountDelegate

		errMsg = "Failed to remove elements " + str(childAccountsIds) + " for ChildAccounts on Account"

		try:
			# get the Account
			account = self.get( accountId ).first()
				
			# split on a comma with no spaces
			idList = childAccountsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Account		
				account = AccountDelegate().get(id).first();	
				# add the Account
				account.childAccounts.remove(account)
				
			# save it		
			account.save()
			
			# reload and return the appropriate version
			return self.get( accountId );
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addContacts( self, accountId, contactsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.ContactDelegate import ContactDelegate

		errMsg = "Failed to add elements " + str(contactsIds) + " for Contacts on Account"

		try:
			# get the Account
			account = self.get( accountId ).first()
				
			# split on a comma with no spaces
			idList = contactsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Contact		
				contact = ContactDelegate().get(id).first();	
				# add the Contact
				account.contacts.add(contact)
				
			# save it		
			account.save()
			
			# reload and return the appropriate version
			return self.get( accountId );
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except Contact.DoesNotExist:
			raise ProcessingError(errMsg + " : Contact does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeContacts( self, accountId, contactsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.ContactDelegate import ContactDelegate

		errMsg = "Failed to remove elements " + str(contactsIds) + " for Contacts on Account"

		try:
			# get the Account
			account = self.get( accountId ).first()
				
			# split on a comma with no spaces
			idList = contactsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Contact		
				contact = ContactDelegate().get(id).first();	
				# add the Contact
				account.contacts.remove(contact)
				
			# save it		
			account.save()
			
			# reload and return the appropriate version
			return self.get( accountId );
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except Contact.DoesNotExist:
			raise ProcessingError(errMsg + " : Contact does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addOpportunities( self, accountId, opportunitiesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OpportunityDelegate import OpportunityDelegate

		errMsg = "Failed to add elements " + str(opportunitiesIds) + " for Opportunities on Account"

		try:
			# get the Account
			account = self.get( accountId ).first()
				
			# split on a comma with no spaces
			idList = opportunitiesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Opportunity		
				opportunity = OpportunityDelegate().get(id).first();	
				# add the Opportunity
				account.opportunities.add(opportunity)
				
			# save it		
			account.save()
			
			# reload and return the appropriate version
			return self.get( accountId );
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except Opportunity.DoesNotExist:
			raise ProcessingError(errMsg + " : Opportunity does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeOpportunities( self, accountId, opportunitiesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OpportunityDelegate import OpportunityDelegate

		errMsg = "Failed to remove elements " + str(opportunitiesIds) + " for Opportunities on Account"

		try:
			# get the Account
			account = self.get( accountId ).first()
				
			# split on a comma with no spaces
			idList = opportunitiesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Opportunity		
				opportunity = OpportunityDelegate().get(id).first();	
				# add the Opportunity
				account.opportunities.remove(opportunity)
				
			# save it		
			account.save()
			
			# reload and return the appropriate version
			return self.get( accountId );
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except Opportunity.DoesNotExist:
			raise ProcessingError(errMsg + " : Opportunity does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addCases( self, accountId, casesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.Case_Delegate import Case_Delegate

		errMsg = "Failed to add elements " + str(casesIds) + " for Cases on Account"

		try:
			# get the Account
			account = self.get( accountId ).first()
				
			# split on a comma with no spaces
			idList = casesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Case_		
				case_ = Case_Delegate().get(id).first();	
				# add the Case_
				account.cases.add(case_)
				
			# save it		
			account.save()
			
			# reload and return the appropriate version
			return self.get( accountId );
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except Case_.DoesNotExist:
			raise ProcessingError(errMsg + " : Case_ does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCases( self, accountId, casesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.Case_Delegate import Case_Delegate

		errMsg = "Failed to remove elements " + str(casesIds) + " for Cases on Account"

		try:
			# get the Account
			account = self.get( accountId ).first()
				
			# split on a comma with no spaces
			idList = casesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Case_		
				case_ = Case_Delegate().get(id).first();	
				# add the Case_
				account.cases.remove(case_)
				
			# save it		
			account.save()
			
			# reload and return the appropriate version
			return self.get( accountId );
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except Case_.DoesNotExist:
			raise ProcessingError(errMsg + " : Case_ does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addActivities( self, accountId, activitiesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.ActivityDelegate import ActivityDelegate

		errMsg = "Failed to add elements " + str(activitiesIds) + " for Activities on Account"

		try:
			# get the Account
			account = self.get( accountId ).first()
				
			# split on a comma with no spaces
			idList = activitiesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Activity		
				activity = ActivityDelegate().get(id).first();	
				# add the Activity
				account.activities.add(activity)
				
			# save it		
			account.save()
			
			# reload and return the appropriate version
			return self.get( accountId );
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except Activity.DoesNotExist:
			raise ProcessingError(errMsg + " : Activity does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeActivities( self, accountId, activitiesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.ActivityDelegate import ActivityDelegate

		errMsg = "Failed to remove elements " + str(activitiesIds) + " for Activities on Account"

		try:
			# get the Account
			account = self.get( accountId ).first()
				
			# split on a comma with no spaces
			idList = activitiesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Activity		
				activity = ActivityDelegate().get(id).first();	
				# add the Activity
				account.activities.remove(activity)
				
			# save it		
			account.save()
			
			# reload and return the appropriate version
			return self.get( accountId );
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except Activity.DoesNotExist:
			raise ProcessingError(errMsg + " : Activity does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addCampaigns( self, accountId, campaignsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.CampaignDelegate import CampaignDelegate

		errMsg = "Failed to add elements " + str(campaignsIds) + " for Campaigns on Account"

		try:
			# get the Account
			account = self.get( accountId ).first()
				
			# split on a comma with no spaces
			idList = campaignsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Campaign		
				campaign = CampaignDelegate().get(id).first();	
				# add the Campaign
				account.campaigns.add(campaign)
				
			# save it		
			account.save()
			
			# reload and return the appropriate version
			return self.get( accountId );
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCampaigns( self, accountId, campaignsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.CampaignDelegate import CampaignDelegate

		errMsg = "Failed to remove elements " + str(campaignsIds) + " for Campaigns on Account"

		try:
			# get the Account
			account = self.get( accountId ).first()
				
			# split on a comma with no spaces
			idList = campaignsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Campaign		
				campaign = CampaignDelegate().get(id).first();	
				# add the Campaign
				account.campaigns.remove(campaign)
				
			# save it		
			account.save()
			
			# reload and return the appropriate version
			return self.get( accountId );
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addQuotes( self, accountId, quotesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.QuoteDelegate import QuoteDelegate

		errMsg = "Failed to add elements " + str(quotesIds) + " for Quotes on Account"

		try:
			# get the Account
			account = self.get( accountId ).first()
				
			# split on a comma with no spaces
			idList = quotesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Quote		
				quote = QuoteDelegate().get(id).first();	
				# add the Quote
				account.quotes.add(quote)
				
			# save it		
			account.save()
			
			# reload and return the appropriate version
			return self.get( accountId );
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except Quote.DoesNotExist:
			raise ProcessingError(errMsg + " : Quote does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeQuotes( self, accountId, quotesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.QuoteDelegate import QuoteDelegate

		errMsg = "Failed to remove elements " + str(quotesIds) + " for Quotes on Account"

		try:
			# get the Account
			account = self.get( accountId ).first()
				
			# split on a comma with no spaces
			idList = quotesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Quote		
				quote = QuoteDelegate().get(id).first();	
				# add the Quote
				account.quotes.remove(quote)
				
			# save it		
			account.save()
			
			# reload and return the appropriate version
			return self.get( accountId );
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except Quote.DoesNotExist:
			raise ProcessingError(errMsg + " : Quote does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addOrders( self, accountId, ordersIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OrderDelegate import OrderDelegate

		errMsg = "Failed to add elements " + str(ordersIds) + " for Orders on Account"

		try:
			# get the Account
			account = self.get( accountId ).first()
				
			# split on a comma with no spaces
			idList = ordersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Order		
				order = OrderDelegate().get(id).first();	
				# add the Order
				account.orders.add(order)
				
			# save it		
			account.save()
			
			# reload and return the appropriate version
			return self.get( accountId );
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeOrders( self, accountId, ordersIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OrderDelegate import OrderDelegate

		errMsg = "Failed to remove elements " + str(ordersIds) + " for Orders on Account"

		try:
			# get the Account
			account = self.get( accountId ).first()
				
			# split on a comma with no spaces
			idList = ordersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Order		
				order = OrderDelegate().get(id).first();	
				# add the Order
				account.orders.remove(order)
				
			# save it		
			account.save()
			
			# reload and return the appropriate version
			return self.get( accountId );
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addContracts( self, accountId, contractsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.ContractDelegate import ContractDelegate

		errMsg = "Failed to add elements " + str(contractsIds) + " for Contracts on Account"

		try:
			# get the Account
			account = self.get( accountId ).first()
				
			# split on a comma with no spaces
			idList = contractsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Contract		
				contract = ContractDelegate().get(id).first();	
				# add the Contract
				account.contracts.add(contract)
				
			# save it		
			account.save()
			
			# reload and return the appropriate version
			return self.get( accountId );
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except Contract.DoesNotExist:
			raise ProcessingError(errMsg + " : Contract does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeContracts( self, accountId, contractsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.ContractDelegate import ContractDelegate

		errMsg = "Failed to remove elements " + str(contractsIds) + " for Contracts on Account"

		try:
			# get the Account
			account = self.get( accountId ).first()
				
			# split on a comma with no spaces
			idList = contractsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Contract		
				contract = ContractDelegate().get(id).first();	
				# add the Contract
				account.contracts.remove(contract)
				
			# save it		
			account.save()
			
			# reload and return the appropriate version
			return self.get( accountId );
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except Contract.DoesNotExist:
			raise ProcessingError(errMsg + " : Contract does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addNotes( self, accountId, notesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.NoteDelegate import NoteDelegate

		errMsg = "Failed to add elements " + str(notesIds) + " for Notes on Account"

		try:
			# get the Account
			account = self.get( accountId ).first()
				
			# split on a comma with no spaces
			idList = notesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Note		
				note = NoteDelegate().get(id).first();	
				# add the Note
				account.notes.add(note)
				
			# save it		
			account.save()
			
			# reload and return the appropriate version
			return self.get( accountId );
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except Note.DoesNotExist:
			raise ProcessingError(errMsg + " : Note does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeNotes( self, accountId, notesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.NoteDelegate import NoteDelegate

		errMsg = "Failed to remove elements " + str(notesIds) + " for Notes on Account"

		try:
			# get the Account
			account = self.get( accountId ).first()
				
			# split on a comma with no spaces
			idList = notesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Note		
				note = NoteDelegate().get(id).first();	
				# add the Note
				account.notes.remove(note)
				
			# save it		
			account.save()
			
			# reload and return the appropriate version
			return self.get( accountId );
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except Note.DoesNotExist:
			raise ProcessingError(errMsg + " : Note does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addEmailMessages( self, accountId, emailMessagesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.EmailMessageDelegate import EmailMessageDelegate

		errMsg = "Failed to add elements " + str(emailMessagesIds) + " for EmailMessages on Account"

		try:
			# get the Account
			account = self.get( accountId ).first()
				
			# split on a comma with no spaces
			idList = emailMessagesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the EmailMessage		
				emailMessage = EmailMessageDelegate().get(id).first();	
				# add the EmailMessage
				account.emailMessages.add(emailMessage)
				
			# save it		
			account.save()
			
			# reload and return the appropriate version
			return self.get( accountId );
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except EmailMessage.DoesNotExist:
			raise ProcessingError(errMsg + " : EmailMessage does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeEmailMessages( self, accountId, emailMessagesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.EmailMessageDelegate import EmailMessageDelegate

		errMsg = "Failed to remove elements " + str(emailMessagesIds) + " for EmailMessages on Account"

		try:
			# get the Account
			account = self.get( accountId ).first()
				
			# split on a comma with no spaces
			idList = emailMessagesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the EmailMessage		
				emailMessage = EmailMessageDelegate().get(id).first();	
				# add the EmailMessage
				account.emailMessages.remove(emailMessage)
				
			# save it		
			account.save()
			
			# reload and return the appropriate version
			return self.get( accountId );
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except EmailMessage.DoesNotExist:
			raise ProcessingError(errMsg + " : EmailMessage does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
