from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from advertisingOnDjango.models.Agency import Agency
from advertisingOnDjango.models.Advertiser import Advertiser
from advertisingOnDjango.models.Team import Team
from advertisingOnDjango.models.User import User
from advertisingOnDjango.models.InsertionOrder import InsertionOrder
from advertisingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Agency
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AgencyDelegate Declaration
#======================================================================
class AgencyDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, agencyId ):
		try:	
			agency = Agency.objects.filter(id=agencyId)
			return agency.first();
		except Agency.DoesNotExist:
			raise ProcessingError("Agency with id " + str(agencyId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, agency):
		for model in serializers.deserialize("json", agency):
			model.save()
			return model;

	def create(self, agency):
		agency.save()
		return agency;

	def saveFromJson(self, agency):
		for model in serializers.deserialize("json", agency):
			model.save()
			return agency;
	
	def save(self, agency):
		agency.save()
		return agency;
	
	def delete(self, agencyId ):
		errMsg = "Failed to delete Agency from db using id " + str(agencyId)
		
		try:
			agency = Agency.objects.get(id=agencyId)
			agency.delete()
			return True
		except Agency.DoesNotExist:
			raise ProcessingError("Agency with id " + str(agencyId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Agency.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Agency from db")
		except Exception:
			return None;
		
	def addAdvertisers( self, agencyId, advertisersIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.AdvertiserDelegate import AdvertiserDelegate

		errMsg = "Failed to add elements " + str(advertisersIds) + " for Advertisers on Agency"

		try:
			# get the Agency
			agency = self.get( agencyId ).first()
				
			# split on a comma with no spaces
			idList = advertisersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Advertiser		
				advertiser = AdvertiserDelegate().get(id).first();	
				# add the Advertiser
				agency.advertisers.add(advertiser)
				
			# save it		
			agency.save()
			
			# reload and return the appropriate version
			return self.get( agencyId );
		except Agency.DoesNotExist:
			raise ProcessingError(errMsg + " : Agency with id " + str(agencyId) + " does not exist.")
		except Advertiser.DoesNotExist:
			raise ProcessingError(errMsg + " : Advertiser does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAdvertisers( self, agencyId, advertisersIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.AdvertiserDelegate import AdvertiserDelegate

		errMsg = "Failed to remove elements " + str(advertisersIds) + " for Advertisers on Agency"

		try:
			# get the Agency
			agency = self.get( agencyId ).first()
				
			# split on a comma with no spaces
			idList = advertisersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Advertiser		
				advertiser = AdvertiserDelegate().get(id).first();	
				# add the Advertiser
				agency.advertisers.remove(advertiser)
				
			# save it		
			agency.save()
			
			# reload and return the appropriate version
			return self.get( agencyId );
		except Agency.DoesNotExist:
			raise ProcessingError(errMsg + " : Agency with id " + str(agencyId) + " does not exist.")
		except Advertiser.DoesNotExist:
			raise ProcessingError(errMsg + " : Advertiser does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addTeams( self, agencyId, teamsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.TeamDelegate import TeamDelegate

		errMsg = "Failed to add elements " + str(teamsIds) + " for Teams on Agency"

		try:
			# get the Agency
			agency = self.get( agencyId ).first()
				
			# split on a comma with no spaces
			idList = teamsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Team		
				team = TeamDelegate().get(id).first();	
				# add the Team
				agency.teams.add(team)
				
			# save it		
			agency.save()
			
			# reload and return the appropriate version
			return self.get( agencyId );
		except Agency.DoesNotExist:
			raise ProcessingError(errMsg + " : Agency with id " + str(agencyId) + " does not exist.")
		except Team.DoesNotExist:
			raise ProcessingError(errMsg + " : Team does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeTeams( self, agencyId, teamsIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.TeamDelegate import TeamDelegate

		errMsg = "Failed to remove elements " + str(teamsIds) + " for Teams on Agency"

		try:
			# get the Agency
			agency = self.get( agencyId ).first()
				
			# split on a comma with no spaces
			idList = teamsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Team		
				team = TeamDelegate().get(id).first();	
				# add the Team
				agency.teams.remove(team)
				
			# save it		
			agency.save()
			
			# reload and return the appropriate version
			return self.get( agencyId );
		except Agency.DoesNotExist:
			raise ProcessingError(errMsg + " : Agency with id " + str(agencyId) + " does not exist.")
		except Team.DoesNotExist:
			raise ProcessingError(errMsg + " : Team does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addUsers( self, agencyId, usersIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.UserDelegate import UserDelegate

		errMsg = "Failed to add elements " + str(usersIds) + " for Users on Agency"

		try:
			# get the Agency
			agency = self.get( agencyId ).first()
				
			# split on a comma with no spaces
			idList = usersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the User		
				user = UserDelegate().get(id).first();	
				# add the User
				agency.users.add(user)
				
			# save it		
			agency.save()
			
			# reload and return the appropriate version
			return self.get( agencyId );
		except Agency.DoesNotExist:
			raise ProcessingError(errMsg + " : Agency with id " + str(agencyId) + " does not exist.")
		except User.DoesNotExist:
			raise ProcessingError(errMsg + " : User does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeUsers( self, agencyId, usersIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.UserDelegate import UserDelegate

		errMsg = "Failed to remove elements " + str(usersIds) + " for Users on Agency"

		try:
			# get the Agency
			agency = self.get( agencyId ).first()
				
			# split on a comma with no spaces
			idList = usersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the User		
				user = UserDelegate().get(id).first();	
				# add the User
				agency.users.remove(user)
				
			# save it		
			agency.save()
			
			# reload and return the appropriate version
			return self.get( agencyId );
		except Agency.DoesNotExist:
			raise ProcessingError(errMsg + " : Agency with id " + str(agencyId) + " does not exist.")
		except User.DoesNotExist:
			raise ProcessingError(errMsg + " : User does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addInsertionOrders( self, agencyId, insertionOrdersIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.InsertionOrderDelegate import InsertionOrderDelegate

		errMsg = "Failed to add elements " + str(insertionOrdersIds) + " for InsertionOrders on Agency"

		try:
			# get the Agency
			agency = self.get( agencyId ).first()
				
			# split on a comma with no spaces
			idList = insertionOrdersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the InsertionOrder		
				insertionOrder = InsertionOrderDelegate().get(id).first();	
				# add the InsertionOrder
				agency.insertionOrders.add(insertionOrder)
				
			# save it		
			agency.save()
			
			# reload and return the appropriate version
			return self.get( agencyId );
		except Agency.DoesNotExist:
			raise ProcessingError(errMsg + " : Agency with id " + str(agencyId) + " does not exist.")
		except InsertionOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : InsertionOrder does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeInsertionOrders( self, agencyId, insertionOrdersIds ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.InsertionOrderDelegate import InsertionOrderDelegate

		errMsg = "Failed to remove elements " + str(insertionOrdersIds) + " for InsertionOrders on Agency"

		try:
			# get the Agency
			agency = self.get( agencyId ).first()
				
			# split on a comma with no spaces
			idList = insertionOrdersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the InsertionOrder		
				insertionOrder = InsertionOrderDelegate().get(id).first();	
				# add the InsertionOrder
				agency.insertionOrders.remove(insertionOrder)
				
			# save it		
			agency.save()
			
			# reload and return the appropriate version
			return self.get( agencyId );
		except Agency.DoesNotExist:
			raise ProcessingError(errMsg + " : Agency with id " + str(agencyId) + " does not exist.")
		except InsertionOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : InsertionOrder does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
