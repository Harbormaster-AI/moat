from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from advertisingOnDjango.models.Team import Team
from advertisingOnDjango.models.Agency import Agency
from advertisingOnDjango.models.User import User
from advertisingOnDjango.models.AdAccount import AdAccount
from advertisingOnDjango.exceptions import Exceptions

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
		
	def assignAgency( self, teamId, agencyId ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.AgencyDelegate import AgencyDelegate

		errMsg = "Failed to assign element " + str(agencyId) + " for Agency on Team"

		try:
			# get the Team from db
			team = self.get( teamId ).first()	
			
			# get the Agency from db
			agency = AgencyDelegate().get(agencyId).first();
			
			# assign the Agency		
			team.agency = agency
			
			#save it
			team.save()

			# reload and return the appropriate version					
			return self.get( teamId );
		except Team.DoesNotExist:
			raise ProcessingError(errMsg + " : Team with id " + str(teamId) + " does not exist.")
		except Agency.DoesNotExist:
			raise ProcessingError(errMsg + " : Agency with id " + str(agencyId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignAgency( self, teamId ):
		errMsg = "Failed to unassign element " + str(agencyId) + " for Agency on Team"

		try:
			# get the Team from db
			team = self.get( teamId ).first()	
			
			# assign to None for unassignment
			team.agency = None			

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
		from advertisingOnDjango.delegates.UserDelegate import UserDelegate

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
		from advertisingOnDjango.delegates.UserDelegate import UserDelegate

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
		
	def addAdAccounts( self, teamId, adAccountsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.AdAccountDelegate import AdAccountDelegate

		errMsg = "Failed to add elements " + str(adAccountsIds) + " for AdAccounts on Team"

		try:
			# get the Team
			team = self.get( teamId ).first()
				
			# split on a comma with no spaces
			idList = adAccountsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the AdAccount		
				adAccount = AdAccountDelegate().get(id).first();	
				# add the AdAccount
				team.adAccounts.add(adAccount)
				
			# save it		
			team.save()
			
			# reload and return the appropriate version
			return self.get( teamId );
		except Team.DoesNotExist:
			raise ProcessingError(errMsg + " : Team with id " + str(teamId) + " does not exist.")
		except AdAccount.DoesNotExist:
			raise ProcessingError(errMsg + " : AdAccount does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAdAccounts( self, teamId, adAccountsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.AdAccountDelegate import AdAccountDelegate

		errMsg = "Failed to remove elements " + str(adAccountsIds) + " for AdAccounts on Team"

		try:
			# get the Team
			team = self.get( teamId ).first()
				
			# split on a comma with no spaces
			idList = adAccountsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the AdAccount		
				adAccount = AdAccountDelegate().get(id).first();	
				# add the AdAccount
				team.adAccounts.remove(adAccount)
				
			# save it		
			team.save()
			
			# reload and return the appropriate version
			return self.get( teamId );
		except Team.DoesNotExist:
			raise ProcessingError(errMsg + " : Team with id " + str(teamId) + " does not exist.")
		except AdAccount.DoesNotExist:
			raise ProcessingError(errMsg + " : AdAccount does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
