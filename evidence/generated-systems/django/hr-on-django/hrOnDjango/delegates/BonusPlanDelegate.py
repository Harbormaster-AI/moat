from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from hrOnDjango.models.BonusPlan import BonusPlan
from hrOnDjango.models.CompensationPackage import CompensationPackage
from hrOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model BonusPlan
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BonusPlanDelegate Declaration
#======================================================================
class BonusPlanDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, bonusPlanId ):
		try:	
			bonusPlan = BonusPlan.objects.filter(id=bonusPlanId)
			return bonusPlan.first();
		except BonusPlan.DoesNotExist:
			raise ProcessingError("BonusPlan with id " + str(bonusPlanId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, bonusPlan):
		for model in serializers.deserialize("json", bonusPlan):
			model.save()
			return model;

	def create(self, bonusPlan):
		bonusPlan.save()
		return bonusPlan;

	def saveFromJson(self, bonusPlan):
		for model in serializers.deserialize("json", bonusPlan):
			model.save()
			return bonusPlan;
	
	def save(self, bonusPlan):
		bonusPlan.save()
		return bonusPlan;
	
	def delete(self, bonusPlanId ):
		errMsg = "Failed to delete BonusPlan from db using id " + str(bonusPlanId)
		
		try:
			bonusPlan = BonusPlan.objects.get(id=bonusPlanId)
			bonusPlan.delete()
			return True
		except BonusPlan.DoesNotExist:
			raise ProcessingError("BonusPlan with id " + str(bonusPlanId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = BonusPlan.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all BonusPlan from db")
		except Exception:
			return None;
		
	def addCompensationPackages( self, bonusPlanId, compensationPackagesIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.CompensationPackageDelegate import CompensationPackageDelegate

		errMsg = "Failed to add elements " + str(compensationPackagesIds) + " for CompensationPackages on BonusPlan"

		try:
			# get the BonusPlan
			bonusPlan = self.get( bonusPlanId ).first()
				
			# split on a comma with no spaces
			idList = compensationPackagesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the CompensationPackage		
				compensationPackage = CompensationPackageDelegate().get(id).first();	
				# add the CompensationPackage
				bonusPlan.compensationPackages.add(compensationPackage)
				
			# save it		
			bonusPlan.save()
			
			# reload and return the appropriate version
			return self.get( bonusPlanId );
		except BonusPlan.DoesNotExist:
			raise ProcessingError(errMsg + " : BonusPlan with id " + str(bonusPlanId) + " does not exist.")
		except CompensationPackage.DoesNotExist:
			raise ProcessingError(errMsg + " : CompensationPackage does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCompensationPackages( self, bonusPlanId, compensationPackagesIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.CompensationPackageDelegate import CompensationPackageDelegate

		errMsg = "Failed to remove elements " + str(compensationPackagesIds) + " for CompensationPackages on BonusPlan"

		try:
			# get the BonusPlan
			bonusPlan = self.get( bonusPlanId ).first()
				
			# split on a comma with no spaces
			idList = compensationPackagesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the CompensationPackage		
				compensationPackage = CompensationPackageDelegate().get(id).first();	
				# add the CompensationPackage
				bonusPlan.compensationPackages.remove(compensationPackage)
				
			# save it		
			bonusPlan.save()
			
			# reload and return the appropriate version
			return self.get( bonusPlanId );
		except BonusPlan.DoesNotExist:
			raise ProcessingError(errMsg + " : BonusPlan with id " + str(bonusPlanId) + " does not exist.")
		except CompensationPackage.DoesNotExist:
			raise ProcessingError(errMsg + " : CompensationPackage does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
