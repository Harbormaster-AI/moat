from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from crmOnDjango.models.Organization import Organization
from crmOnDjango.models.User import User
from crmOnDjango.models.Account import Account
from crmOnDjango.models.Team import Team
from crmOnDjango.models.Territory import Territory
from crmOnDjango.models.Product import Product
from crmOnDjango.models.PriceBook import PriceBook
from crmOnDjango.models.Campaign import Campaign
from crmOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Organization
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OrganizationDelegate Declaration
#======================================================================
class OrganizationDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, organizationId ):
		try:	
			organization = Organization.objects.filter(id=organizationId)
			return organization.first();
		except Organization.DoesNotExist:
			raise ProcessingError("Organization with id " + str(organizationId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, organization):
		for model in serializers.deserialize("json", organization):
			model.save()
			return model;

	def create(self, organization):
		organization.save()
		return organization;

	def saveFromJson(self, organization):
		for model in serializers.deserialize("json", organization):
			model.save()
			return organization;
	
	def save(self, organization):
		organization.save()
		return organization;
	
	def delete(self, organizationId ):
		errMsg = "Failed to delete Organization from db using id " + str(organizationId)
		
		try:
			organization = Organization.objects.get(id=organizationId)
			organization.delete()
			return True
		except Organization.DoesNotExist:
			raise ProcessingError("Organization with id " + str(organizationId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Organization.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Organization from db")
		except Exception:
			return None;
		
	def addUsers( self, organizationId, usersIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.UserDelegate import UserDelegate

		errMsg = "Failed to add elements " + str(usersIds) + " for Users on Organization"

		try:
			# get the Organization
			organization = self.get( organizationId ).first()
				
			# split on a comma with no spaces
			idList = usersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the User		
				user = UserDelegate().get(id).first();	
				# add the User
				organization.users.add(user)
				
			# save it		
			organization.save()
			
			# reload and return the appropriate version
			return self.get( organizationId );
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except User.DoesNotExist:
			raise ProcessingError(errMsg + " : User does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeUsers( self, organizationId, usersIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.UserDelegate import UserDelegate

		errMsg = "Failed to remove elements " + str(usersIds) + " for Users on Organization"

		try:
			# get the Organization
			organization = self.get( organizationId ).first()
				
			# split on a comma with no spaces
			idList = usersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the User		
				user = UserDelegate().get(id).first();	
				# add the User
				organization.users.remove(user)
				
			# save it		
			organization.save()
			
			# reload and return the appropriate version
			return self.get( organizationId );
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except User.DoesNotExist:
			raise ProcessingError(errMsg + " : User does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addAccounts( self, organizationId, accountsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.AccountDelegate import AccountDelegate

		errMsg = "Failed to add elements " + str(accountsIds) + " for Accounts on Organization"

		try:
			# get the Organization
			organization = self.get( organizationId ).first()
				
			# split on a comma with no spaces
			idList = accountsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Account		
				account = AccountDelegate().get(id).first();	
				# add the Account
				organization.accounts.add(account)
				
			# save it		
			organization.save()
			
			# reload and return the appropriate version
			return self.get( organizationId );
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAccounts( self, organizationId, accountsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.AccountDelegate import AccountDelegate

		errMsg = "Failed to remove elements " + str(accountsIds) + " for Accounts on Organization"

		try:
			# get the Organization
			organization = self.get( organizationId ).first()
				
			# split on a comma with no spaces
			idList = accountsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Account		
				account = AccountDelegate().get(id).first();	
				# add the Account
				organization.accounts.remove(account)
				
			# save it		
			organization.save()
			
			# reload and return the appropriate version
			return self.get( organizationId );
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addTeams( self, organizationId, teamsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.TeamDelegate import TeamDelegate

		errMsg = "Failed to add elements " + str(teamsIds) + " for Teams on Organization"

		try:
			# get the Organization
			organization = self.get( organizationId ).first()
				
			# split on a comma with no spaces
			idList = teamsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Team		
				team = TeamDelegate().get(id).first();	
				# add the Team
				organization.teams.add(team)
				
			# save it		
			organization.save()
			
			# reload and return the appropriate version
			return self.get( organizationId );
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Team.DoesNotExist:
			raise ProcessingError(errMsg + " : Team does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeTeams( self, organizationId, teamsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.TeamDelegate import TeamDelegate

		errMsg = "Failed to remove elements " + str(teamsIds) + " for Teams on Organization"

		try:
			# get the Organization
			organization = self.get( organizationId ).first()
				
			# split on a comma with no spaces
			idList = teamsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Team		
				team = TeamDelegate().get(id).first();	
				# add the Team
				organization.teams.remove(team)
				
			# save it		
			organization.save()
			
			# reload and return the appropriate version
			return self.get( organizationId );
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Team.DoesNotExist:
			raise ProcessingError(errMsg + " : Team does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addTerritories( self, organizationId, territoriesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.TerritoryDelegate import TerritoryDelegate

		errMsg = "Failed to add elements " + str(territoriesIds) + " for Territories on Organization"

		try:
			# get the Organization
			organization = self.get( organizationId ).first()
				
			# split on a comma with no spaces
			idList = territoriesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Territory		
				territory = TerritoryDelegate().get(id).first();	
				# add the Territory
				organization.territories.add(territory)
				
			# save it		
			organization.save()
			
			# reload and return the appropriate version
			return self.get( organizationId );
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Territory.DoesNotExist:
			raise ProcessingError(errMsg + " : Territory does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeTerritories( self, organizationId, territoriesIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.TerritoryDelegate import TerritoryDelegate

		errMsg = "Failed to remove elements " + str(territoriesIds) + " for Territories on Organization"

		try:
			# get the Organization
			organization = self.get( organizationId ).first()
				
			# split on a comma with no spaces
			idList = territoriesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Territory		
				territory = TerritoryDelegate().get(id).first();	
				# add the Territory
				organization.territories.remove(territory)
				
			# save it		
			organization.save()
			
			# reload and return the appropriate version
			return self.get( organizationId );
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Territory.DoesNotExist:
			raise ProcessingError(errMsg + " : Territory does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addProducts( self, organizationId, productsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.ProductDelegate import ProductDelegate

		errMsg = "Failed to add elements " + str(productsIds) + " for Products on Organization"

		try:
			# get the Organization
			organization = self.get( organizationId ).first()
				
			# split on a comma with no spaces
			idList = productsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Product		
				product = ProductDelegate().get(id).first();	
				# add the Product
				organization.products.add(product)
				
			# save it		
			organization.save()
			
			# reload and return the appropriate version
			return self.get( organizationId );
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Product.DoesNotExist:
			raise ProcessingError(errMsg + " : Product does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeProducts( self, organizationId, productsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.ProductDelegate import ProductDelegate

		errMsg = "Failed to remove elements " + str(productsIds) + " for Products on Organization"

		try:
			# get the Organization
			organization = self.get( organizationId ).first()
				
			# split on a comma with no spaces
			idList = productsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Product		
				product = ProductDelegate().get(id).first();	
				# add the Product
				organization.products.remove(product)
				
			# save it		
			organization.save()
			
			# reload and return the appropriate version
			return self.get( organizationId );
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Product.DoesNotExist:
			raise ProcessingError(errMsg + " : Product does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addPriceBooks( self, organizationId, priceBooksIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.PriceBookDelegate import PriceBookDelegate

		errMsg = "Failed to add elements " + str(priceBooksIds) + " for PriceBooks on Organization"

		try:
			# get the Organization
			organization = self.get( organizationId ).first()
				
			# split on a comma with no spaces
			idList = priceBooksIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the PriceBook		
				priceBook = PriceBookDelegate().get(id).first();	
				# add the PriceBook
				organization.priceBooks.add(priceBook)
				
			# save it		
			organization.save()
			
			# reload and return the appropriate version
			return self.get( organizationId );
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except PriceBook.DoesNotExist:
			raise ProcessingError(errMsg + " : PriceBook does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePriceBooks( self, organizationId, priceBooksIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.PriceBookDelegate import PriceBookDelegate

		errMsg = "Failed to remove elements " + str(priceBooksIds) + " for PriceBooks on Organization"

		try:
			# get the Organization
			organization = self.get( organizationId ).first()
				
			# split on a comma with no spaces
			idList = priceBooksIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the PriceBook		
				priceBook = PriceBookDelegate().get(id).first();	
				# add the PriceBook
				organization.priceBooks.remove(priceBook)
				
			# save it		
			organization.save()
			
			# reload and return the appropriate version
			return self.get( organizationId );
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except PriceBook.DoesNotExist:
			raise ProcessingError(errMsg + " : PriceBook does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addCampaigns( self, organizationId, campaignsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.CampaignDelegate import CampaignDelegate

		errMsg = "Failed to add elements " + str(campaignsIds) + " for Campaigns on Organization"

		try:
			# get the Organization
			organization = self.get( organizationId ).first()
				
			# split on a comma with no spaces
			idList = campaignsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Campaign		
				campaign = CampaignDelegate().get(id).first();	
				# add the Campaign
				organization.campaigns.add(campaign)
				
			# save it		
			organization.save()
			
			# reload and return the appropriate version
			return self.get( organizationId );
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCampaigns( self, organizationId, campaignsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.CampaignDelegate import CampaignDelegate

		errMsg = "Failed to remove elements " + str(campaignsIds) + " for Campaigns on Organization"

		try:
			# get the Organization
			organization = self.get( organizationId ).first()
				
			# split on a comma with no spaces
			idList = campaignsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Campaign		
				campaign = CampaignDelegate().get(id).first();	
				# add the Campaign
				organization.campaigns.remove(campaign)
				
			# save it		
			organization.save()
			
			# reload and return the appropriate version
			return self.get( organizationId );
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
