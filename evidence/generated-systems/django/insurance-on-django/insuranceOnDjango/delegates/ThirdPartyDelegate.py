from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from insuranceOnDjango.models.ThirdParty import ThirdParty
from insuranceOnDjango.models.SubrogationRecovery import SubrogationRecovery
from insuranceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model ThirdParty
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ThirdPartyDelegate Declaration
#======================================================================
class ThirdPartyDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, thirdPartyId ):
		try:	
			thirdParty = ThirdParty.objects.filter(id=thirdPartyId)
			return thirdParty.first();
		except ThirdParty.DoesNotExist:
			raise ProcessingError("ThirdParty with id " + str(thirdPartyId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, thirdParty):
		for model in serializers.deserialize("json", thirdParty):
			model.save()
			return model;

	def create(self, thirdParty):
		thirdParty.save()
		return thirdParty;

	def saveFromJson(self, thirdParty):
		for model in serializers.deserialize("json", thirdParty):
			model.save()
			return thirdParty;
	
	def save(self, thirdParty):
		thirdParty.save()
		return thirdParty;
	
	def delete(self, thirdPartyId ):
		errMsg = "Failed to delete ThirdParty from db using id " + str(thirdPartyId)
		
		try:
			thirdParty = ThirdParty.objects.get(id=thirdPartyId)
			thirdParty.delete()
			return True
		except ThirdParty.DoesNotExist:
			raise ProcessingError("ThirdParty with id " + str(thirdPartyId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = ThirdParty.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all ThirdParty from db")
		except Exception:
			return None;
		
	def addSubrogations( self, thirdPartyId, subrogationsIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.SubrogationRecoveryDelegate import SubrogationRecoveryDelegate

		errMsg = "Failed to add elements " + str(subrogationsIds) + " for Subrogations on ThirdParty"

		try:
			# get the ThirdParty
			thirdParty = self.get( thirdPartyId ).first()
				
			# split on a comma with no spaces
			idList = subrogationsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the SubrogationRecovery		
				subrogationRecovery = SubrogationRecoveryDelegate().get(id).first();	
				# add the SubrogationRecovery
				thirdParty.subrogations.add(subrogationRecovery)
				
			# save it		
			thirdParty.save()
			
			# reload and return the appropriate version
			return self.get( thirdPartyId );
		except ThirdParty.DoesNotExist:
			raise ProcessingError(errMsg + " : ThirdParty with id " + str(thirdPartyId) + " does not exist.")
		except SubrogationRecovery.DoesNotExist:
			raise ProcessingError(errMsg + " : SubrogationRecovery does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeSubrogations( self, thirdPartyId, subrogationsIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.SubrogationRecoveryDelegate import SubrogationRecoveryDelegate

		errMsg = "Failed to remove elements " + str(subrogationsIds) + " for Subrogations on ThirdParty"

		try:
			# get the ThirdParty
			thirdParty = self.get( thirdPartyId ).first()
				
			# split on a comma with no spaces
			idList = subrogationsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the SubrogationRecovery		
				subrogationRecovery = SubrogationRecoveryDelegate().get(id).first();	
				# add the SubrogationRecovery
				thirdParty.subrogations.remove(subrogationRecovery)
				
			# save it		
			thirdParty.save()
			
			# reload and return the appropriate version
			return self.get( thirdPartyId );
		except ThirdParty.DoesNotExist:
			raise ProcessingError(errMsg + " : ThirdParty with id " + str(thirdPartyId) + " does not exist.")
		except SubrogationRecovery.DoesNotExist:
			raise ProcessingError(errMsg + " : SubrogationRecovery does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
