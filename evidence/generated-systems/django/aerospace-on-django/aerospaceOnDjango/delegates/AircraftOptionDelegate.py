from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from aerospaceOnDjango.models.AircraftOption import AircraftOption
from aerospaceOnDjango.models.AircraftVariant import AircraftVariant
from aerospaceOnDjango.models.AircraftPackage import AircraftPackage
from aerospaceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model AircraftOption
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AircraftOptionDelegate Declaration
#======================================================================
class AircraftOptionDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, aircraftOptionId ):
		try:	
			aircraftOption = AircraftOption.objects.filter(id=aircraftOptionId)
			return aircraftOption.first();
		except AircraftOption.DoesNotExist:
			raise ProcessingError("AircraftOption with id " + str(aircraftOptionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, aircraftOption):
		for model in serializers.deserialize("json", aircraftOption):
			model.save()
			return model;

	def create(self, aircraftOption):
		aircraftOption.save()
		return aircraftOption;

	def saveFromJson(self, aircraftOption):
		for model in serializers.deserialize("json", aircraftOption):
			model.save()
			return aircraftOption;
	
	def save(self, aircraftOption):
		aircraftOption.save()
		return aircraftOption;
	
	def delete(self, aircraftOptionId ):
		errMsg = "Failed to delete AircraftOption from db using id " + str(aircraftOptionId)
		
		try:
			aircraftOption = AircraftOption.objects.get(id=aircraftOptionId)
			aircraftOption.delete()
			return True
		except AircraftOption.DoesNotExist:
			raise ProcessingError("AircraftOption with id " + str(aircraftOptionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = AircraftOption.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all AircraftOption from db")
		except Exception:
			return None;
		
	def addVariants( self, aircraftOptionId, variantsIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftVariantDelegate import AircraftVariantDelegate

		errMsg = "Failed to add elements " + str(variantsIds) + " for Variants on AircraftOption"

		try:
			# get the AircraftOption
			aircraftOption = self.get( aircraftOptionId ).first()
				
			# split on a comma with no spaces
			idList = variantsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the AircraftVariant		
				aircraftVariant = AircraftVariantDelegate().get(id).first();	
				# add the AircraftVariant
				aircraftOption.variants.add(aircraftVariant)
				
			# save it		
			aircraftOption.save()
			
			# reload and return the appropriate version
			return self.get( aircraftOptionId );
		except AircraftOption.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftOption with id " + str(aircraftOptionId) + " does not exist.")
		except AircraftVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftVariant does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeVariants( self, aircraftOptionId, variantsIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftVariantDelegate import AircraftVariantDelegate

		errMsg = "Failed to remove elements " + str(variantsIds) + " for Variants on AircraftOption"

		try:
			# get the AircraftOption
			aircraftOption = self.get( aircraftOptionId ).first()
				
			# split on a comma with no spaces
			idList = variantsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the AircraftVariant		
				aircraftVariant = AircraftVariantDelegate().get(id).first();	
				# add the AircraftVariant
				aircraftOption.variants.remove(aircraftVariant)
				
			# save it		
			aircraftOption.save()
			
			# reload and return the appropriate version
			return self.get( aircraftOptionId );
		except AircraftOption.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftOption with id " + str(aircraftOptionId) + " does not exist.")
		except AircraftVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftVariant does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addPackages( self, aircraftOptionId, packagesIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftPackageDelegate import AircraftPackageDelegate

		errMsg = "Failed to add elements " + str(packagesIds) + " for Packages on AircraftOption"

		try:
			# get the AircraftOption
			aircraftOption = self.get( aircraftOptionId ).first()
				
			# split on a comma with no spaces
			idList = packagesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the AircraftPackage		
				aircraftPackage = AircraftPackageDelegate().get(id).first();	
				# add the AircraftPackage
				aircraftOption.packages.add(aircraftPackage)
				
			# save it		
			aircraftOption.save()
			
			# reload and return the appropriate version
			return self.get( aircraftOptionId );
		except AircraftOption.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftOption with id " + str(aircraftOptionId) + " does not exist.")
		except AircraftPackage.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftPackage does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePackages( self, aircraftOptionId, packagesIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftPackageDelegate import AircraftPackageDelegate

		errMsg = "Failed to remove elements " + str(packagesIds) + " for Packages on AircraftOption"

		try:
			# get the AircraftOption
			aircraftOption = self.get( aircraftOptionId ).first()
				
			# split on a comma with no spaces
			idList = packagesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the AircraftPackage		
				aircraftPackage = AircraftPackageDelegate().get(id).first();	
				# add the AircraftPackage
				aircraftOption.packages.remove(aircraftPackage)
				
			# save it		
			aircraftOption.save()
			
			# reload and return the appropriate version
			return self.get( aircraftOptionId );
		except AircraftOption.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftOption with id " + str(aircraftOptionId) + " does not exist.")
		except AircraftPackage.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftPackage does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
