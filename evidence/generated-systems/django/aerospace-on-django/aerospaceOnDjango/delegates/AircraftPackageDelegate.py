from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from aerospaceOnDjango.models.AircraftPackage import AircraftPackage
from aerospaceOnDjango.models.AircraftOption import AircraftOption
from aerospaceOnDjango.models.AircraftVariant import AircraftVariant
from aerospaceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model AircraftPackage
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AircraftPackageDelegate Declaration
#======================================================================
class AircraftPackageDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, aircraftPackageId ):
		try:	
			aircraftPackage = AircraftPackage.objects.filter(id=aircraftPackageId)
			return aircraftPackage.first();
		except AircraftPackage.DoesNotExist:
			raise ProcessingError("AircraftPackage with id " + str(aircraftPackageId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, aircraftPackage):
		for model in serializers.deserialize("json", aircraftPackage):
			model.save()
			return model;

	def create(self, aircraftPackage):
		aircraftPackage.save()
		return aircraftPackage;

	def saveFromJson(self, aircraftPackage):
		for model in serializers.deserialize("json", aircraftPackage):
			model.save()
			return aircraftPackage;
	
	def save(self, aircraftPackage):
		aircraftPackage.save()
		return aircraftPackage;
	
	def delete(self, aircraftPackageId ):
		errMsg = "Failed to delete AircraftPackage from db using id " + str(aircraftPackageId)
		
		try:
			aircraftPackage = AircraftPackage.objects.get(id=aircraftPackageId)
			aircraftPackage.delete()
			return True
		except AircraftPackage.DoesNotExist:
			raise ProcessingError("AircraftPackage with id " + str(aircraftPackageId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = AircraftPackage.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all AircraftPackage from db")
		except Exception:
			return None;
		
	def addOptions( self, aircraftPackageId, optionsIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftOptionDelegate import AircraftOptionDelegate

		errMsg = "Failed to add elements " + str(optionsIds) + " for Options on AircraftPackage"

		try:
			# get the AircraftPackage
			aircraftPackage = self.get( aircraftPackageId ).first()
				
			# split on a comma with no spaces
			idList = optionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the AircraftOption		
				aircraftOption = AircraftOptionDelegate().get(id).first();	
				# add the AircraftOption
				aircraftPackage.options.add(aircraftOption)
				
			# save it		
			aircraftPackage.save()
			
			# reload and return the appropriate version
			return self.get( aircraftPackageId );
		except AircraftPackage.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftPackage with id " + str(aircraftPackageId) + " does not exist.")
		except AircraftOption.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftOption does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeOptions( self, aircraftPackageId, optionsIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftOptionDelegate import AircraftOptionDelegate

		errMsg = "Failed to remove elements " + str(optionsIds) + " for Options on AircraftPackage"

		try:
			# get the AircraftPackage
			aircraftPackage = self.get( aircraftPackageId ).first()
				
			# split on a comma with no spaces
			idList = optionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the AircraftOption		
				aircraftOption = AircraftOptionDelegate().get(id).first();	
				# add the AircraftOption
				aircraftPackage.options.remove(aircraftOption)
				
			# save it		
			aircraftPackage.save()
			
			# reload and return the appropriate version
			return self.get( aircraftPackageId );
		except AircraftPackage.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftPackage with id " + str(aircraftPackageId) + " does not exist.")
		except AircraftOption.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftOption does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addVariants( self, aircraftPackageId, variantsIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftVariantDelegate import AircraftVariantDelegate

		errMsg = "Failed to add elements " + str(variantsIds) + " for Variants on AircraftPackage"

		try:
			# get the AircraftPackage
			aircraftPackage = self.get( aircraftPackageId ).first()
				
			# split on a comma with no spaces
			idList = variantsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the AircraftVariant		
				aircraftVariant = AircraftVariantDelegate().get(id).first();	
				# add the AircraftVariant
				aircraftPackage.variants.add(aircraftVariant)
				
			# save it		
			aircraftPackage.save()
			
			# reload and return the appropriate version
			return self.get( aircraftPackageId );
		except AircraftPackage.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftPackage with id " + str(aircraftPackageId) + " does not exist.")
		except AircraftVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftVariant does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeVariants( self, aircraftPackageId, variantsIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftVariantDelegate import AircraftVariantDelegate

		errMsg = "Failed to remove elements " + str(variantsIds) + " for Variants on AircraftPackage"

		try:
			# get the AircraftPackage
			aircraftPackage = self.get( aircraftPackageId ).first()
				
			# split on a comma with no spaces
			idList = variantsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the AircraftVariant		
				aircraftVariant = AircraftVariantDelegate().get(id).first();	
				# add the AircraftVariant
				aircraftPackage.variants.remove(aircraftVariant)
				
			# save it		
			aircraftPackage.save()
			
			# reload and return the appropriate version
			return self.get( aircraftPackageId );
		except AircraftPackage.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftPackage with id " + str(aircraftPackageId) + " does not exist.")
		except AircraftVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftVariant does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
