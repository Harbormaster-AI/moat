from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from advertisingOnDjango.models.User import User
from advertisingOnDjango.models.Agency import Agency
from advertisingOnDjango.models.Team import Team
from advertisingOnDjango.models.AdAccount import AdAccount
from advertisingOnDjango.exceptions import Exceptions

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
		
	def assignAgency( self, userId, agencyId ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.AgencyDelegate import AgencyDelegate

		errMsg = "Failed to assign element " + str(agencyId) + " for Agency on User"

		try:
			# get the User from db
			user = self.get( userId ).first()	
			
			# get the Agency from db
			agency = AgencyDelegate().get(agencyId).first();
			
			# assign the Agency		
			user.agency = agency
			
			#save it
			user.save()

			# reload and return the appropriate version					
			return self.get( userId );
		except User.DoesNotExist:
			raise ProcessingError(errMsg + " : User with id " + str(userId) + " does not exist.")
		except Agency.DoesNotExist:
			raise ProcessingError(errMsg + " : Agency with id " + str(agencyId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignAgency( self, userId ):
		errMsg = "Failed to unassign element " + str(agencyId) + " for Agency on User"

		try:
			# get the User from db
			user = self.get( userId ).first()	
			
			# assign to None for unassignment
			user.agency = None			

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
		from advertisingOnDjango.delegates.TeamDelegate import TeamDelegate

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
		from advertisingOnDjango.delegates.TeamDelegate import TeamDelegate

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
		
	def addAdAccounts( self, userId, adAccountsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.AdAccountDelegate import AdAccountDelegate

		errMsg = "Failed to add elements " + str(adAccountsIds) + " for AdAccounts on User"

		try:
			# get the User
			user = self.get( userId ).first()
				
			# split on a comma with no spaces
			idList = adAccountsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the AdAccount		
				adAccount = AdAccountDelegate().get(id).first();	
				# add the AdAccount
				user.adAccounts.add(adAccount)
				
			# save it		
			user.save()
			
			# reload and return the appropriate version
			return self.get( userId );
		except User.DoesNotExist:
			raise ProcessingError(errMsg + " : User with id " + str(userId) + " does not exist.")
		except AdAccount.DoesNotExist:
			raise ProcessingError(errMsg + " : AdAccount does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAdAccounts( self, userId, adAccountsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.AdAccountDelegate import AdAccountDelegate

		errMsg = "Failed to remove elements " + str(adAccountsIds) + " for AdAccounts on User"

		try:
			# get the User
			user = self.get( userId ).first()
				
			# split on a comma with no spaces
			idList = adAccountsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the AdAccount		
				adAccount = AdAccountDelegate().get(id).first();	
				# add the AdAccount
				user.adAccounts.remove(adAccount)
				
			# save it		
			user.save()
			
			# reload and return the appropriate version
			return self.get( userId );
		except User.DoesNotExist:
			raise ProcessingError(errMsg + " : User with id " + str(userId) + " does not exist.")
		except AdAccount.DoesNotExist:
			raise ProcessingError(errMsg + " : AdAccount does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
