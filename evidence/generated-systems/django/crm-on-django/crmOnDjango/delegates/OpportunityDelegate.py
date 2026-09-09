from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from crmOnDjango.models.Opportunity import Opportunity
from crmOnDjango.models.Organization import Organization
from crmOnDjango.models.Account import Account
from crmOnDjango.models.User import User
from crmOnDjango.models.Contact import Contact
from crmOnDjango.models.OpportunityLineItem import OpportunityLineItem
from crmOnDjango.models.OpportunityStageHistory import OpportunityStageHistory
from crmOnDjango.models.Quote import Quote
from crmOnDjango.models.Order import Order
from crmOnDjango.models.Campaign import Campaign
from crmOnDjango.models.Activity import Activity
from crmOnDjango.models.Team import Team
from crmOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Opportunity
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OpportunityDelegate Declaration
#======================================================================
class OpportunityDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, opportunityId ):
		try:	
			opportunity = Opportunity.objects.filter(id=opportunityId)
			return opportunity.first();
		except Opportunity.DoesNotExist:
			raise ProcessingError("Opportunity with id " + str(opportunityId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, opportunity):
		for model in serializers.deserialize("json", opportunity):
			model.save()
			return model;

	def create(self, opportunity):
		opportunity.save()
		return opportunity;

	def saveFromJson(self, opportunity):
		for model in serializers.deserialize("json", opportunity):
			model.save()
			return opportunity;
	
	def save(self, opportunity):
		opportunity.save()
		return opportunity;
	
	def delete(self, opportunityId ):
		errMsg = "Failed to delete Opportunity from db using id " + str(opportunityId)
		
		try:
			opportunity = Opportunity.objects.get(id=opportunityId)
			opportunity.delete()
			return True
		except Opportunity.DoesNotExist:
			raise ProcessingError("Opportunity with id " + str(opportunityId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Opportunity.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Opportunity from db")
		except Exception:
			return None;
		
	def assignOrganization( self, opportunityId, organizationId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OrganizationDelegate import OrganizationDelegate

		errMsg = "Failed to assign element " + str(organizationId) + " for Organization on Opportunity"

		try:
			# get the Opportunity from db
			opportunity = self.get( opportunityId ).first()	
			
			# get the Organization from db
			organization = OrganizationDelegate().get(organizationId).first();
			
			# assign the Organization		
			opportunity.organization = organization
			
			#save it
			opportunity.save()

			# reload and return the appropriate version					
			return self.get( opportunityId );
		except Opportunity.DoesNotExist:
			raise ProcessingError(errMsg + " : Opportunity with id " + str(opportunityId) + " does not exist.")
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrganization( self, opportunityId ):
		errMsg = "Failed to unassign element " + str(organizationId) + " for Organization on Opportunity"

		try:
			# get the Opportunity from db
			opportunity = self.get( opportunityId ).first()	
			
			# assign to None for unassignment
			opportunity.organization = None			

			#save it
			opportunity.save()

			# reload and return the appropriate version					
			return self.get( opportunityId );
		except Opportunity.DoesNotExist:
			raise ProcessingError(errMsg + " : Opportunity with id " + str(opportunityId) + " does not exist.")
		except Exception:
			return None;
		
	def assignAccount( self, opportunityId, accountId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.AccountDelegate import AccountDelegate

		errMsg = "Failed to assign element " + str(accountId) + " for Account on Opportunity"

		try:
			# get the Opportunity from db
			opportunity = self.get( opportunityId ).first()	
			
			# get the Account from db
			account = AccountDelegate().get(accountId).first();
			
			# assign the Account		
			opportunity.account = account
			
			#save it
			opportunity.save()

			# reload and return the appropriate version					
			return self.get( opportunityId );
		except Opportunity.DoesNotExist:
			raise ProcessingError(errMsg + " : Opportunity with id " + str(opportunityId) + " does not exist.")
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account with id " + str(accountId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignAccount( self, opportunityId ):
		errMsg = "Failed to unassign element " + str(accountId) + " for Account on Opportunity"

		try:
			# get the Opportunity from db
			opportunity = self.get( opportunityId ).first()	
			
			# assign to None for unassignment
			opportunity.account = None			

			#save it
			opportunity.save()

			# reload and return the appropriate version					
			return self.get( opportunityId );
		except Opportunity.DoesNotExist:
			raise ProcessingError(errMsg + " : Opportunity with id " + str(opportunityId) + " does not exist.")
		except Exception:
			return None;
		
	def assignOwner( self, opportunityId, ownerId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.UserDelegate import UserDelegate

		errMsg = "Failed to assign element " + str(ownerId) + " for Owner on Opportunity"

		try:
			# get the Opportunity from db
			opportunity = self.get( opportunityId ).first()	
			
			# get the User from db
			user = UserDelegate().get(ownerId).first();
			
			# assign the Owner		
			opportunity.owner = user
			
			#save it
			opportunity.save()

			# reload and return the appropriate version					
			return self.get( opportunityId );
		except Opportunity.DoesNotExist:
			raise ProcessingError(errMsg + " : Opportunity with id " + str(opportunityId) + " does not exist.")
		except User.DoesNotExist:
			raise ProcessingError(errMsg + " : User with id " + str(ownerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOwner( self, opportunityId ):
		errMsg = "Failed to unassign element " + str(ownerId) + " for Owner on Opportunity"

		try:
			# get the Opportunity from db
			opportunity = self.get( opportunityId ).first()	
			
			# assign to None for unassignment
			opportunity.user = None			

			#save it
			opportunity.save()

			# reload and return the appropriate version					
			return self.get( opportunityId );
		except Opportunity.DoesNotExist:
			raise ProcessingError(errMsg + " : Opportunity with id " + str(opportunityId) + " does not exist.")
		except Exception:
			return None;
		
	def addContacts( self, opportunityId, contactsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.ContactDelegate import ContactDelegate

		errMsg = "Failed to add elements " + str(contactsIds) + " for Contacts on Opportunity"

		try:
			# get the Opportunity
			opportunity = self.get( opportunityId ).first()
				
			# split on a comma with no spaces
			idList = contactsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Contact		
				contact = ContactDelegate().get(id).first();	
				# add the Contact
				opportunity.contacts.add(contact)
				
			# save it		
			opportunity.save()
			
			# reload and return the appropriate version
			return self.get( opportunityId );
		except Opportunity.DoesNotExist:
			raise ProcessingError(errMsg + " : Opportunity with id " + str(opportunityId) + " does not exist.")
		except Contact.DoesNotExist:
			raise ProcessingError(errMsg + " : Contact does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeContacts( self, opportunityId, contactsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.ContactDelegate import ContactDelegate

		errMsg = "Failed to remove elements " + str(contactsIds) + " for Contacts on Opportunity"

		try:
			# get the Opportunity
			opportunity = self.get( opportunityId ).first()
				
			# split on a comma with no spaces
			idList = contactsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Contact		
				contact = ContactDelegate().get(id).first();	
				# add the Contact
				opportunity.contacts.remove(contact)
				
			# save it		
			opportunity.save()
			
			# reload and return the appropriate version
			return self.get( opportunityId );
		except Opportunity.DoesNotExist:
			raise ProcessingError(errMsg + " : Opportunity with id " + str(opportunityId) + " does not exist.")
		except Contact.DoesNotExist:
			raise ProcessingError(errMsg + " : Contact does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addLineItems( self, opportunityId, lineItemsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OpportunityLineItemDelegate import OpportunityLineItemDelegate

		errMsg = "Failed to add elements " + str(lineItemsIds) + " for LineItems on Opportunity"

		try:
			# get the Opportunity
			opportunity = self.get( opportunityId ).first()
				
			# split on a comma with no spaces
			idList = lineItemsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the OpportunityLineItem		
				opportunityLineItem = OpportunityLineItemDelegate().get(id).first();	
				# add the OpportunityLineItem
				opportunity.lineItems.add(opportunityLineItem)
				
			# save it		
			opportunity.save()
			
			# reload and return the appropriate version
			return self.get( opportunityId );
		except Opportunity.DoesNotExist:
			raise ProcessingError(errMsg + " : Opportunity with id " + str(opportunityId) + " does not exist.")
		except OpportunityLineItem.DoesNotExist:
			raise ProcessingError(errMsg + " : OpportunityLineItem does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeLineItems( self, opportunityId, lineItemsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OpportunityLineItemDelegate import OpportunityLineItemDelegate

		errMsg = "Failed to remove elements " + str(lineItemsIds) + " for LineItems on Opportunity"

		try:
			# get the Opportunity
			opportunity = self.get( opportunityId ).first()
				
			# split on a comma with no spaces
			idList = lineItemsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the OpportunityLineItem		
				opportunityLineItem = OpportunityLineItemDelegate().get(id).first();	
				# add the OpportunityLineItem
				opportunity.lineItems.remove(opportunityLineItem)
				
			# save it		
			opportunity.save()
			
			# reload and return the appropriate version
			return self.get( opportunityId );
		except Opportunity.DoesNotExist:
			raise ProcessingError(errMsg + " : Opportunity with id " + str(opportunityId) + " does not exist.")
		except OpportunityLineItem.DoesNotExist:
			raise ProcessingError(errMsg + " : OpportunityLineItem does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addStageHistory( self, opportunityId, stageHistoryIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OpportunityStageHistoryDelegate import OpportunityStageHistoryDelegate

		errMsg = "Failed to add elements " + str(stageHistoryIds) + " for StageHistory on Opportunity"

		try:
			# get the Opportunity
			opportunity = self.get( opportunityId ).first()
				
			# split on a comma with no spaces
			idList = stageHistoryIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the OpportunityStageHistory		
				opportunityStageHistory = OpportunityStageHistoryDelegate().get(id).first();	
				# add the OpportunityStageHistory
				opportunity.stageHistory.add(opportunityStageHistory)
				
			# save it		
			opportunity.save()
			
			# reload and return the appropriate version
			return self.get( opportunityId );
		except Opportunity.DoesNotExist:
			raise ProcessingError(errMsg + " : Opportunity with id " + str(opportunityId) + " does not exist.")
		except OpportunityStageHistory.DoesNotExist:
			raise ProcessingError(errMsg + " : OpportunityStageHistory does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeStageHistory( self, opportunityId, stageHistoryIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OpportunityStageHistoryDelegate import OpportunityStageHistoryDelegate

		errMsg = "Failed to remove elements " + str(stageHistoryIds) + " for StageHistory on Opportunity"

		try:
			# get the Opportunity
			opportunity = self.get( opportunityId ).first()
				
			# split on a comma with no spaces
			idList = stageHistoryIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the OpportunityStageHistory		
				opportunityStageHistory = OpportunityStageHistoryDelegate().get(id).first();	
				# add the OpportunityStageHistory
				opportunity.stageHistory.remove(opportunityStageHistory)
				
			# save it		
			opportunity.save()
			
			# reload and return the appropriate version
			return self.get( opportunityId );
		except Opportunity.DoesNotExist:
			raise ProcessingError(errMsg + " : Opportunity with id " + str(opportunityId) + " does not exist.")
		except OpportunityStageHistory.DoesNotExist:
			raise ProcessingError(errMsg + " : OpportunityStageHistory does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addQuotes( self, opportunityId, quotesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.QuoteDelegate import QuoteDelegate

		errMsg = "Failed to add elements " + str(quotesIds) + " for Quotes on Opportunity"

		try:
			# get the Opportunity
			opportunity = self.get( opportunityId ).first()
				
			# split on a comma with no spaces
			idList = quotesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Quote		
				quote = QuoteDelegate().get(id).first();	
				# add the Quote
				opportunity.quotes.add(quote)
				
			# save it		
			opportunity.save()
			
			# reload and return the appropriate version
			return self.get( opportunityId );
		except Opportunity.DoesNotExist:
			raise ProcessingError(errMsg + " : Opportunity with id " + str(opportunityId) + " does not exist.")
		except Quote.DoesNotExist:
			raise ProcessingError(errMsg + " : Quote does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeQuotes( self, opportunityId, quotesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.QuoteDelegate import QuoteDelegate

		errMsg = "Failed to remove elements " + str(quotesIds) + " for Quotes on Opportunity"

		try:
			# get the Opportunity
			opportunity = self.get( opportunityId ).first()
				
			# split on a comma with no spaces
			idList = quotesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Quote		
				quote = QuoteDelegate().get(id).first();	
				# add the Quote
				opportunity.quotes.remove(quote)
				
			# save it		
			opportunity.save()
			
			# reload and return the appropriate version
			return self.get( opportunityId );
		except Opportunity.DoesNotExist:
			raise ProcessingError(errMsg + " : Opportunity with id " + str(opportunityId) + " does not exist.")
		except Quote.DoesNotExist:
			raise ProcessingError(errMsg + " : Quote does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addOrders( self, opportunityId, ordersIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OrderDelegate import OrderDelegate

		errMsg = "Failed to add elements " + str(ordersIds) + " for Orders on Opportunity"

		try:
			# get the Opportunity
			opportunity = self.get( opportunityId ).first()
				
			# split on a comma with no spaces
			idList = ordersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Order		
				order = OrderDelegate().get(id).first();	
				# add the Order
				opportunity.orders.add(order)
				
			# save it		
			opportunity.save()
			
			# reload and return the appropriate version
			return self.get( opportunityId );
		except Opportunity.DoesNotExist:
			raise ProcessingError(errMsg + " : Opportunity with id " + str(opportunityId) + " does not exist.")
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeOrders( self, opportunityId, ordersIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OrderDelegate import OrderDelegate

		errMsg = "Failed to remove elements " + str(ordersIds) + " for Orders on Opportunity"

		try:
			# get the Opportunity
			opportunity = self.get( opportunityId ).first()
				
			# split on a comma with no spaces
			idList = ordersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Order		
				order = OrderDelegate().get(id).first();	
				# add the Order
				opportunity.orders.remove(order)
				
			# save it		
			opportunity.save()
			
			# reload and return the appropriate version
			return self.get( opportunityId );
		except Opportunity.DoesNotExist:
			raise ProcessingError(errMsg + " : Opportunity with id " + str(opportunityId) + " does not exist.")
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addCampaigns( self, opportunityId, campaignsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.CampaignDelegate import CampaignDelegate

		errMsg = "Failed to add elements " + str(campaignsIds) + " for Campaigns on Opportunity"

		try:
			# get the Opportunity
			opportunity = self.get( opportunityId ).first()
				
			# split on a comma with no spaces
			idList = campaignsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Campaign		
				campaign = CampaignDelegate().get(id).first();	
				# add the Campaign
				opportunity.campaigns.add(campaign)
				
			# save it		
			opportunity.save()
			
			# reload and return the appropriate version
			return self.get( opportunityId );
		except Opportunity.DoesNotExist:
			raise ProcessingError(errMsg + " : Opportunity with id " + str(opportunityId) + " does not exist.")
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCampaigns( self, opportunityId, campaignsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.CampaignDelegate import CampaignDelegate

		errMsg = "Failed to remove elements " + str(campaignsIds) + " for Campaigns on Opportunity"

		try:
			# get the Opportunity
			opportunity = self.get( opportunityId ).first()
				
			# split on a comma with no spaces
			idList = campaignsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Campaign		
				campaign = CampaignDelegate().get(id).first();	
				# add the Campaign
				opportunity.campaigns.remove(campaign)
				
			# save it		
			opportunity.save()
			
			# reload and return the appropriate version
			return self.get( opportunityId );
		except Opportunity.DoesNotExist:
			raise ProcessingError(errMsg + " : Opportunity with id " + str(opportunityId) + " does not exist.")
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addActivities( self, opportunityId, activitiesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.ActivityDelegate import ActivityDelegate

		errMsg = "Failed to add elements " + str(activitiesIds) + " for Activities on Opportunity"

		try:
			# get the Opportunity
			opportunity = self.get( opportunityId ).first()
				
			# split on a comma with no spaces
			idList = activitiesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Activity		
				activity = ActivityDelegate().get(id).first();	
				# add the Activity
				opportunity.activities.add(activity)
				
			# save it		
			opportunity.save()
			
			# reload and return the appropriate version
			return self.get( opportunityId );
		except Opportunity.DoesNotExist:
			raise ProcessingError(errMsg + " : Opportunity with id " + str(opportunityId) + " does not exist.")
		except Activity.DoesNotExist:
			raise ProcessingError(errMsg + " : Activity does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeActivities( self, opportunityId, activitiesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.ActivityDelegate import ActivityDelegate

		errMsg = "Failed to remove elements " + str(activitiesIds) + " for Activities on Opportunity"

		try:
			# get the Opportunity
			opportunity = self.get( opportunityId ).first()
				
			# split on a comma with no spaces
			idList = activitiesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Activity		
				activity = ActivityDelegate().get(id).first();	
				# add the Activity
				opportunity.activities.remove(activity)
				
			# save it		
			opportunity.save()
			
			# reload and return the appropriate version
			return self.get( opportunityId );
		except Opportunity.DoesNotExist:
			raise ProcessingError(errMsg + " : Opportunity with id " + str(opportunityId) + " does not exist.")
		except Activity.DoesNotExist:
			raise ProcessingError(errMsg + " : Activity does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addTeams( self, opportunityId, teamsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.TeamDelegate import TeamDelegate

		errMsg = "Failed to add elements " + str(teamsIds) + " for Teams on Opportunity"

		try:
			# get the Opportunity
			opportunity = self.get( opportunityId ).first()
				
			# split on a comma with no spaces
			idList = teamsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Team		
				team = TeamDelegate().get(id).first();	
				# add the Team
				opportunity.teams.add(team)
				
			# save it		
			opportunity.save()
			
			# reload and return the appropriate version
			return self.get( opportunityId );
		except Opportunity.DoesNotExist:
			raise ProcessingError(errMsg + " : Opportunity with id " + str(opportunityId) + " does not exist.")
		except Team.DoesNotExist:
			raise ProcessingError(errMsg + " : Team does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeTeams( self, opportunityId, teamsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.TeamDelegate import TeamDelegate

		errMsg = "Failed to remove elements " + str(teamsIds) + " for Teams on Opportunity"

		try:
			# get the Opportunity
			opportunity = self.get( opportunityId ).first()
				
			# split on a comma with no spaces
			idList = teamsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Team		
				team = TeamDelegate().get(id).first();	
				# add the Team
				opportunity.teams.remove(team)
				
			# save it		
			opportunity.save()
			
			# reload and return the appropriate version
			return self.get( opportunityId );
		except Opportunity.DoesNotExist:
			raise ProcessingError(errMsg + " : Opportunity with id " + str(opportunityId) + " does not exist.")
		except Team.DoesNotExist:
			raise ProcessingError(errMsg + " : Team does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
