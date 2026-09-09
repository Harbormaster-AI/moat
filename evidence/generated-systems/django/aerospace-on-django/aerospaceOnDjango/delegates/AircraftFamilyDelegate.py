from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from aerospaceOnDjango.models.AircraftFamily import AircraftFamily
from aerospaceOnDjango.models.AircraftProgram import AircraftProgram
from aerospaceOnDjango.models.AircraftModel import AircraftModel
from aerospaceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model AircraftFamily
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AircraftFamilyDelegate Declaration
#======================================================================
class AircraftFamilyDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, aircraftFamilyId ):
		try:	
			aircraftFamily = AircraftFamily.objects.filter(id=aircraftFamilyId)
			return aircraftFamily.first();
		except AircraftFamily.DoesNotExist:
			raise ProcessingError("AircraftFamily with id " + str(aircraftFamilyId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, aircraftFamily):
		for model in serializers.deserialize("json", aircraftFamily):
			model.save()
			return model;

	def create(self, aircraftFamily):
		aircraftFamily.save()
		return aircraftFamily;

	def saveFromJson(self, aircraftFamily):
		for model in serializers.deserialize("json", aircraftFamily):
			model.save()
			return aircraftFamily;
	
	def save(self, aircraftFamily):
		aircraftFamily.save()
		return aircraftFamily;
	
	def delete(self, aircraftFamilyId ):
		errMsg = "Failed to delete AircraftFamily from db using id " + str(aircraftFamilyId)
		
		try:
			aircraftFamily = AircraftFamily.objects.get(id=aircraftFamilyId)
			aircraftFamily.delete()
			return True
		except AircraftFamily.DoesNotExist:
			raise ProcessingError("AircraftFamily with id " + str(aircraftFamilyId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = AircraftFamily.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all AircraftFamily from db")
		except Exception:
			return None;
		
	def assignProgram( self, aircraftFamilyId, programId ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftProgramDelegate import AircraftProgramDelegate

		errMsg = "Failed to assign element " + str(programId) + " for Program on AircraftFamily"

		try:
			# get the AircraftFamily from db
			aircraftFamily = self.get( aircraftFamilyId ).first()	
			
			# get the AircraftProgram from db
			aircraftProgram = AircraftProgramDelegate().get(programId).first();
			
			# assign the Program		
			aircraftFamily.program = aircraftProgram
			
			#save it
			aircraftFamily.save()

			# reload and return the appropriate version					
			return self.get( aircraftFamilyId );
		except AircraftFamily.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftFamily with id " + str(aircraftFamilyId) + " does not exist.")
		except AircraftProgram.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftProgram with id " + str(programId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignProgram( self, aircraftFamilyId ):
		errMsg = "Failed to unassign element " + str(programId) + " for Program on AircraftFamily"

		try:
			# get the AircraftFamily from db
			aircraftFamily = self.get( aircraftFamilyId ).first()	
			
			# assign to None for unassignment
			aircraftFamily.aircraftProgram = None			

			#save it
			aircraftFamily.save()

			# reload and return the appropriate version					
			return self.get( aircraftFamilyId );
		except AircraftFamily.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftFamily with id " + str(aircraftFamilyId) + " does not exist.")
		except Exception:
			return None;
		
	def addAircraftModels( self, aircraftFamilyId, aircraftModelsIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftModelDelegate import AircraftModelDelegate

		errMsg = "Failed to add elements " + str(aircraftModelsIds) + " for AircraftModels on AircraftFamily"

		try:
			# get the AircraftFamily
			aircraftFamily = self.get( aircraftFamilyId ).first()
				
			# split on a comma with no spaces
			idList = aircraftModelsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the AircraftModel		
				aircraftModel = AircraftModelDelegate().get(id).first();	
				# add the AircraftModel
				aircraftFamily.aircraftModels.add(aircraftModel)
				
			# save it		
			aircraftFamily.save()
			
			# reload and return the appropriate version
			return self.get( aircraftFamilyId );
		except AircraftFamily.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftFamily with id " + str(aircraftFamilyId) + " does not exist.")
		except AircraftModel.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftModel does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAircraftModels( self, aircraftFamilyId, aircraftModelsIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftModelDelegate import AircraftModelDelegate

		errMsg = "Failed to remove elements " + str(aircraftModelsIds) + " for AircraftModels on AircraftFamily"

		try:
			# get the AircraftFamily
			aircraftFamily = self.get( aircraftFamilyId ).first()
				
			# split on a comma with no spaces
			idList = aircraftModelsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the AircraftModel		
				aircraftModel = AircraftModelDelegate().get(id).first();	
				# add the AircraftModel
				aircraftFamily.aircraftModels.remove(aircraftModel)
				
			# save it		
			aircraftFamily.save()
			
			# reload and return the appropriate version
			return self.get( aircraftFamilyId );
		except AircraftFamily.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftFamily with id " + str(aircraftFamilyId) + " does not exist.")
		except AircraftModel.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftModel does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
