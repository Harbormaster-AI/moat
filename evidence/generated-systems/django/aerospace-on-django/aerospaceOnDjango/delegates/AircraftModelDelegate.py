from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from aerospaceOnDjango.models.AircraftModel import AircraftModel
from aerospaceOnDjango.models.AircraftFamily import AircraftFamily
from aerospaceOnDjango.models.AircraftVariant import AircraftVariant
from aerospaceOnDjango.models.EngineType import EngineType
from aerospaceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model AircraftModel
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AircraftModelDelegate Declaration
#======================================================================
class AircraftModelDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, aircraftModelId ):
		try:	
			aircraftModel = AircraftModel.objects.filter(id=aircraftModelId)
			return aircraftModel.first();
		except AircraftModel.DoesNotExist:
			raise ProcessingError("AircraftModel with id " + str(aircraftModelId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, aircraftModel):
		for model in serializers.deserialize("json", aircraftModel):
			model.save()
			return model;

	def create(self, aircraftModel):
		aircraftModel.save()
		return aircraftModel;

	def saveFromJson(self, aircraftModel):
		for model in serializers.deserialize("json", aircraftModel):
			model.save()
			return aircraftModel;
	
	def save(self, aircraftModel):
		aircraftModel.save()
		return aircraftModel;
	
	def delete(self, aircraftModelId ):
		errMsg = "Failed to delete AircraftModel from db using id " + str(aircraftModelId)
		
		try:
			aircraftModel = AircraftModel.objects.get(id=aircraftModelId)
			aircraftModel.delete()
			return True
		except AircraftModel.DoesNotExist:
			raise ProcessingError("AircraftModel with id " + str(aircraftModelId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = AircraftModel.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all AircraftModel from db")
		except Exception:
			return None;
		
	def assignFamily( self, aircraftModelId, familyId ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftFamilyDelegate import AircraftFamilyDelegate

		errMsg = "Failed to assign element " + str(familyId) + " for Family on AircraftModel"

		try:
			# get the AircraftModel from db
			aircraftModel = self.get( aircraftModelId ).first()	
			
			# get the AircraftFamily from db
			aircraftFamily = AircraftFamilyDelegate().get(familyId).first();
			
			# assign the Family		
			aircraftModel.family = aircraftFamily
			
			#save it
			aircraftModel.save()

			# reload and return the appropriate version					
			return self.get( aircraftModelId );
		except AircraftModel.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftModel with id " + str(aircraftModelId) + " does not exist.")
		except AircraftFamily.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftFamily with id " + str(familyId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignFamily( self, aircraftModelId ):
		errMsg = "Failed to unassign element " + str(familyId) + " for Family on AircraftModel"

		try:
			# get the AircraftModel from db
			aircraftModel = self.get( aircraftModelId ).first()	
			
			# assign to None for unassignment
			aircraftModel.aircraftFamily = None			

			#save it
			aircraftModel.save()

			# reload and return the appropriate version					
			return self.get( aircraftModelId );
		except AircraftModel.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftModel with id " + str(aircraftModelId) + " does not exist.")
		except Exception:
			return None;
		
	def addVariants( self, aircraftModelId, variantsIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftVariantDelegate import AircraftVariantDelegate

		errMsg = "Failed to add elements " + str(variantsIds) + " for Variants on AircraftModel"

		try:
			# get the AircraftModel
			aircraftModel = self.get( aircraftModelId ).first()
				
			# split on a comma with no spaces
			idList = variantsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the AircraftVariant		
				aircraftVariant = AircraftVariantDelegate().get(id).first();	
				# add the AircraftVariant
				aircraftModel.variants.add(aircraftVariant)
				
			# save it		
			aircraftModel.save()
			
			# reload and return the appropriate version
			return self.get( aircraftModelId );
		except AircraftModel.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftModel with id " + str(aircraftModelId) + " does not exist.")
		except AircraftVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftVariant does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeVariants( self, aircraftModelId, variantsIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftVariantDelegate import AircraftVariantDelegate

		errMsg = "Failed to remove elements " + str(variantsIds) + " for Variants on AircraftModel"

		try:
			# get the AircraftModel
			aircraftModel = self.get( aircraftModelId ).first()
				
			# split on a comma with no spaces
			idList = variantsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the AircraftVariant		
				aircraftVariant = AircraftVariantDelegate().get(id).first();	
				# add the AircraftVariant
				aircraftModel.variants.remove(aircraftVariant)
				
			# save it		
			aircraftModel.save()
			
			# reload and return the appropriate version
			return self.get( aircraftModelId );
		except AircraftModel.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftModel with id " + str(aircraftModelId) + " does not exist.")
		except AircraftVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftVariant does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addEngineTypes( self, aircraftModelId, engineTypesIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.EngineTypeDelegate import EngineTypeDelegate

		errMsg = "Failed to add elements " + str(engineTypesIds) + " for EngineTypes on AircraftModel"

		try:
			# get the AircraftModel
			aircraftModel = self.get( aircraftModelId ).first()
				
			# split on a comma with no spaces
			idList = engineTypesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the EngineType		
				engineType = EngineTypeDelegate().get(id).first();	
				# add the EngineType
				aircraftModel.engineTypes.add(engineType)
				
			# save it		
			aircraftModel.save()
			
			# reload and return the appropriate version
			return self.get( aircraftModelId );
		except AircraftModel.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftModel with id " + str(aircraftModelId) + " does not exist.")
		except EngineType.DoesNotExist:
			raise ProcessingError(errMsg + " : EngineType does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeEngineTypes( self, aircraftModelId, engineTypesIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.EngineTypeDelegate import EngineTypeDelegate

		errMsg = "Failed to remove elements " + str(engineTypesIds) + " for EngineTypes on AircraftModel"

		try:
			# get the AircraftModel
			aircraftModel = self.get( aircraftModelId ).first()
				
			# split on a comma with no spaces
			idList = engineTypesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the EngineType		
				engineType = EngineTypeDelegate().get(id).first();	
				# add the EngineType
				aircraftModel.engineTypes.remove(engineType)
				
			# save it		
			aircraftModel.save()
			
			# reload and return the appropriate version
			return self.get( aircraftModelId );
		except AircraftModel.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftModel with id " + str(aircraftModelId) + " does not exist.")
		except EngineType.DoesNotExist:
			raise ProcessingError(errMsg + " : EngineType does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
