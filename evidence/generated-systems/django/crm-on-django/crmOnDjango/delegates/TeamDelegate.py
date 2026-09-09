from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from crmOnDjango.models.Team import Team
from crmOnDjango.models.Organization import Organization
from crmOnDjango.models.User import User
from crmOnDjango.models.Account import Account
from crmOnDjango.models.Opportunity import Opportunity
from crmOnDjango.models.Case_ import Case_
from crmOnDjango.models.Campaign import Campaign
from crmOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Team
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TeamDelegate Declaration
#======================================================================
class TeamDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, teamId ):
		try:	
			team = Team.objects.filter(id=teamId)
			return team.first();
		except Team.DoesNotExist:
			raise ProcessingError("Team with id " + str(teamId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, team):
		for model in serializers.deserialize("json", team):
			model.save()
			return model;

	def create(self, team):
		team.save()
		return team;

	def saveFromJson(self, team):
		for model in serializers.deserialize("json", team):
			model.save()
			return team;
	
	def save(self, team):
		team.save()
		return team;
	
	def delete(self, teamId ):
		errMsg = "Failed to delete Team from db using id " + str(teamId)
		
		try:
			team = Team.objects.get(id=teamId)
			team.delete()
			return True
		except Team.DoesNotExist:
			raise ProcessingError("Team with id " + str(teamId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Team.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Team from db")
		except Exception:
			return None;
		
	def assignOrganization( self, teamId, organizationId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OrganizationDelegate import OrganizationDelegate

		errMsg = "Failed to assign element " + str(organizationId) + " for Organization on Team"

		try:
			# get the Team from db
			team = self.get( teamId ).first()	
			
			# get the Organization from db
			organization = OrganizationDelegate().get(organizationId).first();
			
			# assign the Organization		
			team.organization = organization
			
			#save it
			team.save()

			# reload and return the appropriate version					
			return self.get( teamId );
		except Team.DoesNotExist:
			raise ProcessingError(errMsg + " : Team with id " + str(teamId) + " does not exist.")
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrganization( self, teamId ):
		errMsg = "Failed to unassign element " + str(organizationId) + " for Organization on Team"

		try:
			# get the Team from db
			team = self.get( teamId ).first()	
			
			# assign to None for unassignment
			team.organization = None			

			#save it
			team.save()

			# reload and return the appropriate version					
			return self.get( teamId );
		except Team.DoesNotExist:
			raise ProcessingError(errMsg + " : Team with id " + str(teamId) + " does not exist.")
		except Exception:
			return None;
		
	def addUsers( self, teamId, usersIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.UserDelegate import UserDelegate

		errMsg = "Failed to add elements " + str(usersIds) + " for Users on Team"

		try:
			# get the Team
			team = self.get( teamId ).first()
				
			# split on a comma with no spaces
			idList = usersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the User		
				user = UserDelegate().get(id).first();	
				# add the User
				team.users.add(user)
				
			# save it		
			team.save()
			
			# reload and return the appropriate version
			return self.get( teamId );
		except Team.DoesNotExist:
			raise ProcessingError(errMsg + " : Team with id " + str(teamId) + " does not exist.")
		except User.DoesNotExist:
			raise ProcessingError(errMsg + " : User does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeUsers( self, teamId, usersIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.UserDelegate import UserDelegate

		errMsg = "Failed to remove elements " + str(usersIds) + " for Users on Team"

		try:
			# get the Team
			team = self.get( teamId ).first()
				
			# split on a comma with no spaces
			idList = usersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the User		
				user = UserDelegate().get(id).first();	
				# add the User
				team.users.remove(user)
				
			# save it		
			team.save()
			
			# reload and return the appropriate version
			return self.get( teamId );
		except Team.DoesNotExist:
			raise ProcessingError(errMsg + " : Team with id " + str(teamId) + " does not exist.")
		except User.DoesNotExist:
			raise ProcessingError(errMsg + " : User does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addAccounts( self, teamId, accountsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.AccountDelegate import AccountDelegate

		errMsg = "Failed to add elements " + str(accountsIds) + " for Accounts on Team"

		try:
			# get the Team
			team = self.get( teamId ).first()
				
			# split on a comma with no spaces
			idList = accountsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Account		
				account = AccountDelegate().get(id).first();	
				# add the Account
				team.accounts.add(account)
				
			# save it		
			team.save()
			
			# reload and return the appropriate version
			return self.get( teamId );
		except Team.DoesNotExist:
			raise ProcessingError(errMsg + " : Team with id " + str(teamId) + " does not exist.")
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAccounts( self, teamId, accountsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.AccountDelegate import AccountDelegate

		errMsg = "Failed to remove elements " + str(accountsIds) + " for Accounts on Team"

		try:
			# get the Team
			team = self.get( teamId ).first()
				
			# split on a comma with no spaces
			idList = accountsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Account		
				account = AccountDelegate().get(id).first();	
				# add the Account
				team.accounts.remove(account)
				
			# save it		
			team.save()
			
			# reload and return the appropriate version
			return self.get( teamId );
		except Team.DoesNotExist:
			raise ProcessingError(errMsg + " : Team with id " + str(teamId) + " does not exist.")
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addOpportunities( self, teamId, opportunitiesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OpportunityDelegate import OpportunityDelegate

		errMsg = "Failed to add elements " + str(opportunitiesIds) + " for Opportunities on Team"

		try:
			# get the Team
			team = self.get( teamId ).first()
				
			# split on a comma with no spaces
			idList = opportunitiesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Opportunity		
				opportunity = OpportunityDelegate().get(id).first();	
				# add the Opportunity
				team.opportunities.add(opportunity)
				
			# save it		
			team.save()
			
			# reload and return the appropriate version
			return self.get( teamId );
		except Team.DoesNotExist:
			raise ProcessingError(errMsg + " : Team with id " + str(teamId) + " does not exist.")
		except Opportunity.DoesNotExist:
			raise ProcessingError(errMsg + " : Opportunity does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeOpportunities( self, teamId, opportunitiesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OpportunityDelegate import OpportunityDelegate

		errMsg = "Failed to remove elements " + str(opportunitiesIds) + " for Opportunities on Team"

		try:
			# get the Team
			team = self.get( teamId ).first()
				
			# split on a comma with no spaces
			idList = opportunitiesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Opportunity		
				opportunity = OpportunityDelegate().get(id).first();	
				# add the Opportunity
				team.opportunities.remove(opportunity)
				
			# save it		
			team.save()
			
			# reload and return the appropriate version
			return self.get( teamId );
		except Team.DoesNotExist:
			raise ProcessingError(errMsg + " : Team with id " + str(teamId) + " does not exist.")
		except Opportunity.DoesNotExist:
			raise ProcessingError(errMsg + " : Opportunity does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addCases( self, teamId, casesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.Case_Delegate import Case_Delegate

		errMsg = "Failed to add elements " + str(casesIds) + " for Cases on Team"

		try:
			# get the Team
			team = self.get( teamId ).first()
				
			# split on a comma with no spaces
			idList = casesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Case_		
				case_ = Case_Delegate().get(id).first();	
				# add the Case_
				team.cases.add(case_)
				
			# save it		
			team.save()
			
			# reload and return the appropriate version
			return self.get( teamId );
		except Team.DoesNotExist:
			raise ProcessingError(errMsg + " : Team with id " + str(teamId) + " does not exist.")
		except Case_.DoesNotExist:
			raise ProcessingError(errMsg + " : Case_ does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCases( self, teamId, casesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.Case_Delegate import Case_Delegate

		errMsg = "Failed to remove elements " + str(casesIds) + " for Cases on Team"

		try:
			# get the Team
			team = self.get( teamId ).first()
				
			# split on a comma with no spaces
			idList = casesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Case_		
				case_ = Case_Delegate().get(id).first();	
				# add the Case_
				team.cases.remove(case_)
				
			# save it		
			team.save()
			
			# reload and return the appropriate version
			return self.get( teamId );
		except Team.DoesNotExist:
			raise ProcessingError(errMsg + " : Team with id " + str(teamId) + " does not exist.")
		except Case_.DoesNotExist:
			raise ProcessingError(errMsg + " : Case_ does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addCampaigns( self, teamId, campaignsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.CampaignDelegate import CampaignDelegate

		errMsg = "Failed to add elements " + str(campaignsIds) + " for Campaigns on Team"

		try:
			# get the Team
			team = self.get( teamId ).first()
				
			# split on a comma with no spaces
			idList = campaignsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Campaign		
				campaign = CampaignDelegate().get(id).first();	
				# add the Campaign
				team.campaigns.add(campaign)
				
			# save it		
			team.save()
			
			# reload and return the appropriate version
			return self.get( teamId );
		except Team.DoesNotExist:
			raise ProcessingError(errMsg + " : Team with id " + str(teamId) + " does not exist.")
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCampaigns( self, teamId, campaignsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.CampaignDelegate import CampaignDelegate

		errMsg = "Failed to remove elements " + str(campaignsIds) + " for Campaigns on Team"

		try:
			# get the Team
			team = self.get( teamId ).first()
				
			# split on a comma with no spaces
			idList = campaignsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Campaign		
				campaign = CampaignDelegate().get(id).first();	
				# add the Campaign
				team.campaigns.remove(campaign)
				
			# save it		
			team.save()
			
			# reload and return the appropriate version
			return self.get( teamId );
		except Team.DoesNotExist:
			raise ProcessingError(errMsg + " : Team with id " + str(teamId) + " does not exist.")
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
