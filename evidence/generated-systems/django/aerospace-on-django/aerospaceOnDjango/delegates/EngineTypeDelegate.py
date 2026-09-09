from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from aerospaceOnDjango.models.EngineType import EngineType
from aerospaceOnDjango.models.Supplier import Supplier
from aerospaceOnDjango.models.AircraftModel import AircraftModel
from aerospaceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model EngineType
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class EngineTypeDelegate Declaration
#======================================================================
class EngineTypeDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, engineTypeId ):
		try:	
			engineType = EngineType.objects.filter(id=engineTypeId)
			return engineType.first();
		except EngineType.DoesNotExist:
			raise ProcessingError("EngineType with id " + str(engineTypeId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, engineType):
		for model in serializers.deserialize("json", engineType):
			model.save()
			return model;

	def create(self, engineType):
		engineType.save()
		return engineType;

	def saveFromJson(self, engineType):
		for model in serializers.deserialize("json", engineType):
			model.save()
			return engineType;
	
	def save(self, engineType):
		engineType.save()
		return engineType;
	
	def delete(self, engineTypeId ):
		errMsg = "Failed to delete EngineType from db using id " + str(engineTypeId)
		
		try:
			engineType = EngineType.objects.get(id=engineTypeId)
			engineType.delete()
			return True
		except EngineType.DoesNotExist:
			raise ProcessingError("EngineType with id " + str(engineTypeId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = EngineType.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all EngineType from db")
		except Exception:
			return None;
		
	def assignSupplier( self, engineTypeId, supplierId ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.SupplierDelegate import SupplierDelegate

		errMsg = "Failed to assign element " + str(supplierId) + " for Supplier on EngineType"

		try:
			# get the EngineType from db
			engineType = self.get( engineTypeId ).first()	
			
			# get the Supplier from db
			supplier = SupplierDelegate().get(supplierId).first();
			
			# assign the Supplier		
			engineType.supplier = supplier
			
			#save it
			engineType.save()

			# reload and return the appropriate version					
			return self.get( engineTypeId );
		except EngineType.DoesNotExist:
			raise ProcessingError(errMsg + " : EngineType with id " + str(engineTypeId) + " does not exist.")
		except Supplier.DoesNotExist:
			raise ProcessingError(errMsg + " : Supplier with id " + str(supplierId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignSupplier( self, engineTypeId ):
		errMsg = "Failed to unassign element " + str(supplierId) + " for Supplier on EngineType"

		try:
			# get the EngineType from db
			engineType = self.get( engineTypeId ).first()	
			
			# assign to None for unassignment
			engineType.supplier = None			

			#save it
			engineType.save()

			# reload and return the appropriate version					
			return self.get( engineTypeId );
		except EngineType.DoesNotExist:
			raise ProcessingError(errMsg + " : EngineType with id " + str(engineTypeId) + " does not exist.")
		except Exception:
			return None;
		
	def addCompatibleModels( self, engineTypeId, compatibleModelsIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftModelDelegate import AircraftModelDelegate

		errMsg = "Failed to add elements " + str(compatibleModelsIds) + " for CompatibleModels on EngineType"

		try:
			# get the EngineType
			engineType = self.get( engineTypeId ).first()
				
			# split on a comma with no spaces
			idList = compatibleModelsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the AircraftModel		
				aircraftModel = AircraftModelDelegate().get(id).first();	
				# add the AircraftModel
				engineType.compatibleModels.add(aircraftModel)
				
			# save it		
			engineType.save()
			
			# reload and return the appropriate version
			return self.get( engineTypeId );
		except EngineType.DoesNotExist:
			raise ProcessingError(errMsg + " : EngineType with id " + str(engineTypeId) + " does not exist.")
		except AircraftModel.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftModel does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCompatibleModels( self, engineTypeId, compatibleModelsIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftModelDelegate import AircraftModelDelegate

		errMsg = "Failed to remove elements " + str(compatibleModelsIds) + " for CompatibleModels on EngineType"

		try:
			# get the EngineType
			engineType = self.get( engineTypeId ).first()
				
			# split on a comma with no spaces
			idList = compatibleModelsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the AircraftModel		
				aircraftModel = AircraftModelDelegate().get(id).first();	
				# add the AircraftModel
				engineType.compatibleModels.remove(aircraftModel)
				
			# save it		
			engineType.save()
			
			# reload and return the appropriate version
			return self.get( engineTypeId );
		except EngineType.DoesNotExist:
			raise ProcessingError(errMsg + " : EngineType with id " + str(engineTypeId) + " does not exist.")
		except AircraftModel.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftModel does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
