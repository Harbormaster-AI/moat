from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from crmOnDjango.models.User import User
from crmOnDjango.models.Organization import Organization
from crmOnDjango.models.Team import Team
from crmOnDjango.models.Activity import Activity
from crmOnDjango.models.Account import Account
from crmOnDjango.models.Lead import Lead
from crmOnDjango.models.Opportunity import Opportunity
from crmOnDjango.models.Case_ import Case_
from crmOnDjango.models.Quote import Quote
from crmOnDjango.models.Order import Order
from crmOnDjango.models.Contract import Contract
from crmOnDjango.models.EmailMessage import EmailMessage
from crmOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model User
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class UserDelegate Declaration
#======================================================================
class UserDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, userId ):
		try:	
			user = User.objects.filter(id=userId)
			return user.first();
		except User.DoesNotExist:
			raise ProcessingError("User with id " + str(userId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, user):
		for model in serializers.deserialize("json", user):
			model.save()
			return model;

	def create(self, user):
		user.save()
		return user;

	def saveFromJson(self, user):
		for model in serializers.deserialize("json", user):
			model.save()
			return user;
	
	def save(self, user):
		user.save()
		return user;
	
	def delete(self, userId ):
		errMsg = "Failed to delete User from db using id " + str(userId)
		
		try:
			user = User.objects.get(id=userId)
			user.delete()
			return True
		except User.DoesNotExist:
			raise ProcessingError("User with id " + str(userId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = User.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all User from db")
		except Exception:
			return None;
		
	def assignOrganization( self, userId, organizationId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OrganizationDelegate import OrganizationDelegate

		errMsg = "Failed to assign element " + str(organizationId) + " for Organization on User"

		try:
			# get the User from db
			user = self.get( userId ).first()	
			
			# get the Organization from db
			organization = OrganizationDelegate().get(organizationId).first();
			
			# assign the Organization		
			user.organization = organization
			
			#save it
			user.save()

			# reload and return the appropriate version					
			return self.get( userId );
		except User.DoesNotExist:
			raise ProcessingError(errMsg + " : User with id " + str(userId) + " does not exist.")
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrganization( self, userId ):
		errMsg = "Failed to unassign element " + str(organizationId) + " for Organization on User"

		try:
			# get the User from db
			user = self.get( userId ).first()	
			
			# assign to None for unassignment
			user.organization = None			

			#save it
			user.save()

			# reload and return the appropriate version					
			return self.get( userId );
		except User.DoesNotExist:
			raise ProcessingError(errMsg + " : User with id " + str(userId) + " does not exist.")
		except Exception:
			return None;
		
	def addTeams( self, userId, teamsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.TeamDelegate import TeamDelegate

		errMsg = "Failed to add elements " + str(teamsIds) + " for Teams on User"

		try:
			# get the User
			user = self.get( userId ).first()
				
			# split on a comma with no spaces
			idList = teamsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Team		
				team = TeamDelegate().get(id).first();	
				# add the Team
				user.teams.add(team)
				
			# save it		
			user.save()
			
			# reload and return the appropriate version
			return self.get( userId );
		except User.DoesNotExist:
			raise ProcessingError(errMsg + " : User with id " + str(userId) + " does not exist.")
		except Team.DoesNotExist:
			raise ProcessingError(errMsg + " : Team does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeTeams( self, userId, teamsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.TeamDelegate import TeamDelegate

		errMsg = "Failed to remove elements " + str(teamsIds) + " for Teams on User"

		try:
			# get the User
			user = self.get( userId ).first()
				
			# split on a comma with no spaces
			idList = teamsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Team		
				team = TeamDelegate().get(id).first();	
				# add the Team
				user.teams.remove(team)
				
			# save it		
			user.save()
			
			# reload and return the appropriate version
			return self.get( userId );
		except User.DoesNotExist:
			raise ProcessingError(errMsg + " : User with id " + str(userId) + " does not exist.")
		except Team.DoesNotExist:
			raise ProcessingError(errMsg + " : Team does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addActivities( self, userId, activitiesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.ActivityDelegate import ActivityDelegate

		errMsg = "Failed to add elements " + str(activitiesIds) + " for Activities on User"

		try:
			# get the User
			user = self.get( userId ).first()
				
			# split on a comma with no spaces
			idList = activitiesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Activity		
				activity = ActivityDelegate().get(id).first();	
				# add the Activity
				user.activities.add(activity)
				
			# save it		
			user.save()
			
			# reload and return the appropriate version
			return self.get( userId );
		except User.DoesNotExist:
			raise ProcessingError(errMsg + " : User with id " + str(userId) + " does not exist.")
		except Activity.DoesNotExist:
			raise ProcessingError(errMsg + " : Activity does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeActivities( self, userId, activitiesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.ActivityDelegate import ActivityDelegate

		errMsg = "Failed to remove elements " + str(activitiesIds) + " for Activities on User"

		try:
			# get the User
			user = self.get( userId ).first()
				
			# split on a comma with no spaces
			idList = activitiesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Activity		
				activity = ActivityDelegate().get(id).first();	
				# add the Activity
				user.activities.remove(activity)
				
			# save it		
			user.save()
			
			# reload and return the appropriate version
			return self.get( userId );
		except User.DoesNotExist:
			raise ProcessingError(errMsg + " : User with id " + str(userId) + " does not exist.")
		except Activity.DoesNotExist:
			raise ProcessingError(errMsg + " : Activity does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addOwnedAccounts( self, userId, ownedAccountsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.AccountDelegate import AccountDelegate

		errMsg = "Failed to add elements " + str(ownedAccountsIds) + " for OwnedAccounts on User"

		try:
			# get the User
			user = self.get( userId ).first()
				
			# split on a comma with no spaces
			idList = ownedAccountsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Account		
				account = AccountDelegate().get(id).first();	
				# add the Account
				user.ownedAccounts.add(account)
				
			# save it		
			user.save()
			
			# reload and return the appropriate version
			return self.get( userId );
		except User.DoesNotExist:
			raise ProcessingError(errMsg + " : User with id " + str(userId) + " does not exist.")
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeOwnedAccounts( self, userId, ownedAccountsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.AccountDelegate import AccountDelegate

		errMsg = "Failed to remove elements " + str(ownedAccountsIds) + " for OwnedAccounts on User"

		try:
			# get the User
			user = self.get( userId ).first()
				
			# split on a comma with no spaces
			idList = ownedAccountsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Account		
				account = AccountDelegate().get(id).first();	
				# add the Account
				user.ownedAccounts.remove(account)
				
			# save it		
			user.save()
			
			# reload and return the appropriate version
			return self.get( userId );
		except User.DoesNotExist:
			raise ProcessingError(errMsg + " : User with id " + str(userId) + " does not exist.")
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addOwnedLeads( self, userId, ownedLeadsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.LeadDelegate import LeadDelegate

		errMsg = "Failed to add elements " + str(ownedLeadsIds) + " for OwnedLeads on User"

		try:
			# get the User
			user = self.get( userId ).first()
				
			# split on a comma with no spaces
			idList = ownedLeadsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Lead		
				lead = LeadDelegate().get(id).first();	
				# add the Lead
				user.ownedLeads.add(lead)
				
			# save it		
			user.save()
			
			# reload and return the appropriate version
			return self.get( userId );
		except User.DoesNotExist:
			raise ProcessingError(errMsg + " : User with id " + str(userId) + " does not exist.")
		except Lead.DoesNotExist:
			raise ProcessingError(errMsg + " : Lead does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeOwnedLeads( self, userId, ownedLeadsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.LeadDelegate import LeadDelegate

		errMsg = "Failed to remove elements " + str(ownedLeadsIds) + " for OwnedLeads on User"

		try:
			# get the User
			user = self.get( userId ).first()
				
			# split on a comma with no spaces
			idList = ownedLeadsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Lead		
				lead = LeadDelegate().get(id).first();	
				# add the Lead
				user.ownedLeads.remove(lead)
				
			# save it		
			user.save()
			
			# reload and return the appropriate version
			return self.get( userId );
		except User.DoesNotExist:
			raise ProcessingError(errMsg + " : User with id " + str(userId) + " does not exist.")
		except Lead.DoesNotExist:
			raise ProcessingError(errMsg + " : Lead does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addOwnedOpportunities( self, userId, ownedOpportunitiesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OpportunityDelegate import OpportunityDelegate

		errMsg = "Failed to add elements " + str(ownedOpportunitiesIds) + " for OwnedOpportunities on User"

		try:
			# get the User
			user = self.get( userId ).first()
				
			# split on a comma with no spaces
			idList = ownedOpportunitiesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Opportunity		
				opportunity = OpportunityDelegate().get(id).first();	
				# add the Opportunity
				user.ownedOpportunities.add(opportunity)
				
			# save it		
			user.save()
			
			# reload and return the appropriate version
			return self.get( userId );
		except User.DoesNotExist:
			raise ProcessingError(errMsg + " : User with id " + str(userId) + " does not exist.")
		except Opportunity.DoesNotExist:
			raise ProcessingError(errMsg + " : Opportunity does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeOwnedOpportunities( self, userId, ownedOpportunitiesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OpportunityDelegate import OpportunityDelegate

		errMsg = "Failed to remove elements " + str(ownedOpportunitiesIds) + " for OwnedOpportunities on User"

		try:
			# get the User
			user = self.get( userId ).first()
				
			# split on a comma with no spaces
			idList = ownedOpportunitiesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Opportunity		
				opportunity = OpportunityDelegate().get(id).first();	
				# add the Opportunity
				user.ownedOpportunities.remove(opportunity)
				
			# save it		
			user.save()
			
			# reload and return the appropriate version
			return self.get( userId );
		except User.DoesNotExist:
			raise ProcessingError(errMsg + " : User with id " + str(userId) + " does not exist.")
		except Opportunity.DoesNotExist:
			raise ProcessingError(errMsg + " : Opportunity does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addOwnedCases( self, userId, ownedCasesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.Case_Delegate import Case_Delegate

		errMsg = "Failed to add elements " + str(ownedCasesIds) + " for OwnedCases on User"

		try:
			# get the User
			user = self.get( userId ).first()
				
			# split on a comma with no spaces
			idList = ownedCasesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Case_		
				case_ = Case_Delegate().get(id).first();	
				# add the Case_
				user.ownedCases.add(case_)
				
			# save it		
			user.save()
			
			# reload and return the appropriate version
			return self.get( userId );
		except User.DoesNotExist:
			raise ProcessingError(errMsg + " : User with id " + str(userId) + " does not exist.")
		except Case_.DoesNotExist:
			raise ProcessingError(errMsg + " : Case_ does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeOwnedCases( self, userId, ownedCasesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.Case_Delegate import Case_Delegate

		errMsg = "Failed to remove elements " + str(ownedCasesIds) + " for OwnedCases on User"

		try:
			# get the User
			user = self.get( userId ).first()
				
			# split on a comma with no spaces
			idList = ownedCasesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Case_		
				case_ = Case_Delegate().get(id).first();	
				# add the Case_
				user.ownedCases.remove(case_)
				
			# save it		
			user.save()
			
			# reload and return the appropriate version
			return self.get( userId );
		except User.DoesNotExist:
			raise ProcessingError(errMsg + " : User with id " + str(userId) + " does not exist.")
		except Case_.DoesNotExist:
			raise ProcessingError(errMsg + " : Case_ does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addQuotes( self, userId, quotesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.QuoteDelegate import QuoteDelegate

		errMsg = "Failed to add elements " + str(quotesIds) + " for Quotes on User"

		try:
			# get the User
			user = self.get( userId ).first()
				
			# split on a comma with no spaces
			idList = quotesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Quote		
				quote = QuoteDelegate().get(id).first();	
				# add the Quote
				user.quotes.add(quote)
				
			# save it		
			user.save()
			
			# reload and return the appropriate version
			return self.get( userId );
		except User.DoesNotExist:
			raise ProcessingError(errMsg + " : User with id " + str(userId) + " does not exist.")
		except Quote.DoesNotExist:
			raise ProcessingError(errMsg + " : Quote does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeQuotes( self, userId, quotesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.QuoteDelegate import QuoteDelegate

		errMsg = "Failed to remove elements " + str(quotesIds) + " for Quotes on User"

		try:
			# get the User
			user = self.get( userId ).first()
				
			# split on a comma with no spaces
			idList = quotesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Quote		
				quote = QuoteDelegate().get(id).first();	
				# add the Quote
				user.quotes.remove(quote)
				
			# save it		
			user.save()
			
			# reload and return the appropriate version
			return self.get( userId );
		except User.DoesNotExist:
			raise ProcessingError(errMsg + " : User with id " + str(userId) + " does not exist.")
		except Quote.DoesNotExist:
			raise ProcessingError(errMsg + " : Quote does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addOrders( self, userId, ordersIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OrderDelegate import OrderDelegate

		errMsg = "Failed to add elements " + str(ordersIds) + " for Orders on User"

		try:
			# get the User
			user = self.get( userId ).first()
				
			# split on a comma with no spaces
			idList = ordersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Order		
				order = OrderDelegate().get(id).first();	
				# add the Order
				user.orders.add(order)
				
			# save it		
			user.save()
			
			# reload and return the appropriate version
			return self.get( userId );
		except User.DoesNotExist:
			raise ProcessingError(errMsg + " : User with id " + str(userId) + " does not exist.")
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeOrders( self, userId, ordersIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OrderDelegate import OrderDelegate

		errMsg = "Failed to remove elements " + str(ordersIds) + " for Orders on User"

		try:
			# get the User
			user = self.get( userId ).first()
				
			# split on a comma with no spaces
			idList = ordersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Order		
				order = OrderDelegate().get(id).first();	
				# add the Order
				user.orders.remove(order)
				
			# save it		
			user.save()
			
			# reload and return the appropriate version
			return self.get( userId );
		except User.DoesNotExist:
			raise ProcessingError(errMsg + " : User with id " + str(userId) + " does not exist.")
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addContracts( self, userId, contractsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.ContractDelegate import ContractDelegate

		errMsg = "Failed to add elements " + str(contractsIds) + " for Contracts on User"

		try:
			# get the User
			user = self.get( userId ).first()
				
			# split on a comma with no spaces
			idList = contractsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Contract		
				contract = ContractDelegate().get(id).first();	
				# add the Contract
				user.contracts.add(contract)
				
			# save it		
			user.save()
			
			# reload and return the appropriate version
			return self.get( userId );
		except User.DoesNotExist:
			raise ProcessingError(errMsg + " : User with id " + str(userId) + " does not exist.")
		except Contract.DoesNotExist:
			raise ProcessingError(errMsg + " : Contract does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeContracts( self, userId, contractsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.ContractDelegate import ContractDelegate

		errMsg = "Failed to remove elements " + str(contractsIds) + " for Contracts on User"

		try:
			# get the User
			user = self.get( userId ).first()
				
			# split on a comma with no spaces
			idList = contractsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Contract		
				contract = ContractDelegate().get(id).first();	
				# add the Contract
				user.contracts.remove(contract)
				
			# save it		
			user.save()
			
			# reload and return the appropriate version
			return self.get( userId );
		except User.DoesNotExist:
			raise ProcessingError(errMsg + " : User with id " + str(userId) + " does not exist.")
		except Contract.DoesNotExist:
			raise ProcessingError(errMsg + " : Contract does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addEmailMessages( self, userId, emailMessagesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.EmailMessageDelegate import EmailMessageDelegate

		errMsg = "Failed to add elements " + str(emailMessagesIds) + " for EmailMessages on User"

		try:
			# get the User
			user = self.get( userId ).first()
				
			# split on a comma with no spaces
			idList = emailMessagesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the EmailMessage		
				emailMessage = EmailMessageDelegate().get(id).first();	
				# add the EmailMessage
				user.emailMessages.add(emailMessage)
				
			# save it		
			user.save()
			
			# reload and return the appropriate version
			return self.get( userId );
		except User.DoesNotExist:
			raise ProcessingError(errMsg + " : User with id " + str(userId) + " does not exist.")
		except EmailMessage.DoesNotExist:
			raise ProcessingError(errMsg + " : EmailMessage does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeEmailMessages( self, userId, emailMessagesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.EmailMessageDelegate import EmailMessageDelegate

		errMsg = "Failed to remove elements " + str(emailMessagesIds) + " for EmailMessages on User"

		try:
			# get the User
			user = self.get( userId ).first()
				
			# split on a comma with no spaces
			idList = emailMessagesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the EmailMessage		
				emailMessage = EmailMessageDelegate().get(id).first();	
				# add the EmailMessage
				user.emailMessages.remove(emailMessage)
				
			# save it		
			user.save()
			
			# reload and return the appropriate version
			return self.get( userId );
		except User.DoesNotExist:
			raise ProcessingError(errMsg + " : User with id " + str(userId) + " does not exist.")
		except EmailMessage.DoesNotExist:
			raise ProcessingError(errMsg + " : EmailMessage does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
