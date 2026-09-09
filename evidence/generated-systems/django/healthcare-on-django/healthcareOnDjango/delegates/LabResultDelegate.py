from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from healthcareOnDjango.models.LabResult import LabResult
from healthcareOnDjango.models.LaboratoryOrder import LaboratoryOrder
from healthcareOnDjango.models.Observation import Observation
from healthcareOnDjango.models.Laboratory import Laboratory
from healthcareOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model LabResult
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LabResultDelegate Declaration
#======================================================================
class LabResultDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, labResultId ):
		try:	
			labResult = LabResult.objects.filter(id=labResultId)
			return labResult.first();
		except LabResult.DoesNotExist:
			raise ProcessingError("LabResult with id " + str(labResultId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, labResult):
		for model in serializers.deserialize("json", labResult):
			model.save()
			return model;

	def create(self, labResult):
		labResult.save()
		return labResult;

	def saveFromJson(self, labResult):
		for model in serializers.deserialize("json", labResult):
			model.save()
			return labResult;
	
	def save(self, labResult):
		labResult.save()
		return labResult;
	
	def delete(self, labResultId ):
		errMsg = "Failed to delete LabResult from db using id " + str(labResultId)
		
		try:
			labResult = LabResult.objects.get(id=labResultId)
			labResult.delete()
			return True
		except LabResult.DoesNotExist:
			raise ProcessingError("LabResult with id " + str(labResultId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = LabResult.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all LabResult from db")
		except Exception:
			return None;
		
	def assignLaboratoryOrder( self, labResultId, laboratoryOrderId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.LaboratoryOrderDelegate import LaboratoryOrderDelegate

		errMsg = "Failed to assign element " + str(laboratoryOrderId) + " for LaboratoryOrder on LabResult"

		try:
			# get the LabResult from db
			labResult = self.get( labResultId ).first()	
			
			# get the LaboratoryOrder from db
			laboratoryOrder = LaboratoryOrderDelegate().get(laboratoryOrderId).first();
			
			# assign the LaboratoryOrder		
			labResult.laboratoryOrder = laboratoryOrder
			
			#save it
			labResult.save()

			# reload and return the appropriate version					
			return self.get( labResultId );
		except LabResult.DoesNotExist:
			raise ProcessingError(errMsg + " : LabResult with id " + str(labResultId) + " does not exist.")
		except LaboratoryOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : LaboratoryOrder with id " + str(laboratoryOrderId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignLaboratoryOrder( self, labResultId ):
		errMsg = "Failed to unassign element " + str(laboratoryOrderId) + " for LaboratoryOrder on LabResult"

		try:
			# get the LabResult from db
			labResult = self.get( labResultId ).first()	
			
			# assign to None for unassignment
			labResult.laboratoryOrder = None			

			#save it
			labResult.save()

			# reload and return the appropriate version					
			return self.get( labResultId );
		except LabResult.DoesNotExist:
			raise ProcessingError(errMsg + " : LabResult with id " + str(labResultId) + " does not exist.")
		except Exception:
			return None;
		
	def assignLaboratory( self, labResultId, laboratoryId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.LaboratoryDelegate import LaboratoryDelegate

		errMsg = "Failed to assign element " + str(laboratoryId) + " for Laboratory on LabResult"

		try:
			# get the LabResult from db
			labResult = self.get( labResultId ).first()	
			
			# get the Laboratory from db
			laboratory = LaboratoryDelegate().get(laboratoryId).first();
			
			# assign the Laboratory		
			labResult.laboratory = laboratory
			
			#save it
			labResult.save()

			# reload and return the appropriate version					
			return self.get( labResultId );
		except LabResult.DoesNotExist:
			raise ProcessingError(errMsg + " : LabResult with id " + str(labResultId) + " does not exist.")
		except Laboratory.DoesNotExist:
			raise ProcessingError(errMsg + " : Laboratory with id " + str(laboratoryId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignLaboratory( self, labResultId ):
		errMsg = "Failed to unassign element " + str(laboratoryId) + " for Laboratory on LabResult"

		try:
			# get the LabResult from db
			labResult = self.get( labResultId ).first()	
			
			# assign to None for unassignment
			labResult.laboratory = None			

			#save it
			labResult.save()

			# reload and return the appropriate version					
			return self.get( labResultId );
		except LabResult.DoesNotExist:
			raise ProcessingError(errMsg + " : LabResult with id " + str(labResultId) + " does not exist.")
		except Exception:
			return None;
		
	def addObservations( self, labResultId, observationsIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ObservationDelegate import ObservationDelegate

		errMsg = "Failed to add elements " + str(observationsIds) + " for Observations on LabResult"

		try:
			# get the LabResult
			labResult = self.get( labResultId ).first()
				
			# split on a comma with no spaces
			idList = observationsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Observation		
				observation = ObservationDelegate().get(id).first();	
				# add the Observation
				labResult.observations.add(observation)
				
			# save it		
			labResult.save()
			
			# reload and return the appropriate version
			return self.get( labResultId );
		except LabResult.DoesNotExist:
			raise ProcessingError(errMsg + " : LabResult with id " + str(labResultId) + " does not exist.")
		except Observation.DoesNotExist:
			raise ProcessingError(errMsg + " : Observation does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeObservations( self, labResultId, observationsIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ObservationDelegate import ObservationDelegate

		errMsg = "Failed to remove elements " + str(observationsIds) + " for Observations on LabResult"

		try:
			# get the LabResult
			labResult = self.get( labResultId ).first()
				
			# split on a comma with no spaces
			idList = observationsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Observation		
				observation = ObservationDelegate().get(id).first();	
				# add the Observation
				labResult.observations.remove(observation)
				
			# save it		
			labResult.save()
			
			# reload and return the appropriate version
			return self.get( labResultId );
		except LabResult.DoesNotExist:
			raise ProcessingError(errMsg + " : LabResult with id " + str(labResultId) + " does not exist.")
		except Observation.DoesNotExist:
			raise ProcessingError(errMsg + " : Observation does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
