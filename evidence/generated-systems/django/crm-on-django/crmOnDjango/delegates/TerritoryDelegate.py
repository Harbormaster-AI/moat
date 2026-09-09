from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from crmOnDjango.models.Territory import Territory
from crmOnDjango.models.Organization import Organization
from crmOnDjango.models.Account import Account
from crmOnDjango.models.User import User
from crmOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Territory
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TerritoryDelegate Declaration
#======================================================================
class TerritoryDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, territoryId ):
		try:	
			territory = Territory.objects.filter(id=territoryId)
			return territory.first();
		except Territory.DoesNotExist:
			raise ProcessingError("Territory with id " + str(territoryId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, territory):
		for model in serializers.deserialize("json", territory):
			model.save()
			return model;

	def create(self, territory):
		territory.save()
		return territory;

	def saveFromJson(self, territory):
		for model in serializers.deserialize("json", territory):
			model.save()
			return territory;
	
	def save(self, territory):
		territory.save()
		return territory;
	
	def delete(self, territoryId ):
		errMsg = "Failed to delete Territory from db using id " + str(territoryId)
		
		try:
			territory = Territory.objects.get(id=territoryId)
			territory.delete()
			return True
		except Territory.DoesNotExist:
			raise ProcessingError("Territory with id " + str(territoryId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Territory.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Territory from db")
		except Exception:
			return None;
		
	def assignOrganization( self, territoryId, organizationId ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.OrganizationDelegate import OrganizationDelegate

		errMsg = "Failed to assign element " + str(organizationId) + " for Organization on Territory"

		try:
			# get the Territory from db
			territory = self.get( territoryId ).first()	
			
			# get the Organization from db
			organization = OrganizationDelegate().get(organizationId).first();
			
			# assign the Organization		
			territory.organization = organization
			
			#save it
			territory.save()

			# reload and return the appropriate version					
			return self.get( territoryId );
		except Territory.DoesNotExist:
			raise ProcessingError(errMsg + " : Territory with id " + str(territoryId) + " does not exist.")
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrganization( self, territoryId ):
		errMsg = "Failed to unassign element " + str(organizationId) + " for Organization on Territory"

		try:
			# get the Territory from db
			territory = self.get( territoryId ).first()	
			
			# assign to None for unassignment
			territory.organization = None			

			#save it
			territory.save()

			# reload and return the appropriate version					
			return self.get( territoryId );
		except Territory.DoesNotExist:
			raise ProcessingError(errMsg + " : Territory with id " + str(territoryId) + " does not exist.")
		except Exception:
			return None;
		
	def addAccounts( self, territoryId, accountsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.AccountDelegate import AccountDelegate

		errMsg = "Failed to add elements " + str(accountsIds) + " for Accounts on Territory"

		try:
			# get the Territory
			territory = self.get( territoryId ).first()
				
			# split on a comma with no spaces
			idList = accountsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Account		
				account = AccountDelegate().get(id).first();	
				# add the Account
				territory.accounts.add(account)
				
			# save it		
			territory.save()
			
			# reload and return the appropriate version
			return self.get( territoryId );
		except Territory.DoesNotExist:
			raise ProcessingError(errMsg + " : Territory with id " + str(territoryId) + " does not exist.")
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAccounts( self, territoryId, accountsIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.AccountDelegate import AccountDelegate

		errMsg = "Failed to remove elements " + str(accountsIds) + " for Accounts on Territory"

		try:
			# get the Territory
			territory = self.get( territoryId ).first()
				
			# split on a comma with no spaces
			idList = accountsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Account		
				account = AccountDelegate().get(id).first();	
				# add the Account
				territory.accounts.remove(account)
				
			# save it		
			territory.save()
			
			# reload and return the appropriate version
			return self.get( territoryId );
		except Territory.DoesNotExist:
			raise ProcessingError(errMsg + " : Territory with id " + str(territoryId) + " does not exist.")
		except Account.DoesNotExist:
			raise ProcessingError(errMsg + " : Account does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addUsers( self, territoryId, usersIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.UserDelegate import UserDelegate

		errMsg = "Failed to add elements " + str(usersIds) + " for Users on Territory"

		try:
			# get the Territory
			territory = self.get( territoryId ).first()
				
			# split on a comma with no spaces
			idList = usersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the User		
				user = UserDelegate().get(id).first();	
				# add the User
				territory.users.add(user)
				
			# save it		
			territory.save()
			
			# reload and return the appropriate version
			return self.get( territoryId );
		except Territory.DoesNotExist:
			raise ProcessingError(errMsg + " : Territory with id " + str(territoryId) + " does not exist.")
		except User.DoesNotExist:
			raise ProcessingError(errMsg + " : User does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeUsers( self, territoryId, usersIds ):
		# lazy importing avoids circular dependencies
		from crmOnDjango.delegates.UserDelegate import UserDelegate

		errMsg = "Failed to remove elements " + str(usersIds) + " for Users on Territory"

		try:
			# get the Territory
			territory = self.get( territoryId ).first()
				
			# split on a comma with no spaces
			idList = usersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the User		
				user = UserDelegate().get(id).first();	
				# add the User
				territory.users.remove(user)
				
			# save it		
			territory.save()
			
			# reload and return the appropriate version
			return self.get( territoryId );
		except Territory.DoesNotExist:
			raise ProcessingError(errMsg + " : Territory with id " + str(territoryId) + " does not exist.")
		except User.DoesNotExist:
			raise ProcessingError(errMsg + " : User does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
