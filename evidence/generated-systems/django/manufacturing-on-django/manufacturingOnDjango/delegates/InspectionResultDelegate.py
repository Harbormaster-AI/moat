from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from manufacturingOnDjango.models.InspectionResult import InspectionResult
from manufacturingOnDjango.models.InspectionLot import InspectionLot
from manufacturingOnDjango.models.InspectionCharacteristic import InspectionCharacteristic
from manufacturingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model InspectionResult
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InspectionResultDelegate Declaration
#======================================================================
class InspectionResultDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, inspectionResultId ):
		try:	
			inspectionResult = InspectionResult.objects.filter(id=inspectionResultId)
			return inspectionResult.first();
		except InspectionResult.DoesNotExist:
			raise ProcessingError("InspectionResult with id " + str(inspectionResultId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, inspectionResult):
		for model in serializers.deserialize("json", inspectionResult):
			model.save()
			return model;

	def create(self, inspectionResult):
		inspectionResult.save()
		return inspectionResult;

	def saveFromJson(self, inspectionResult):
		for model in serializers.deserialize("json", inspectionResult):
			model.save()
			return inspectionResult;
	
	def save(self, inspectionResult):
		inspectionResult.save()
		return inspectionResult;
	
	def delete(self, inspectionResultId ):
		errMsg = "Failed to delete InspectionResult from db using id " + str(inspectionResultId)
		
		try:
			inspectionResult = InspectionResult.objects.get(id=inspectionResultId)
			inspectionResult.delete()
			return True
		except InspectionResult.DoesNotExist:
			raise ProcessingError("InspectionResult with id " + str(inspectionResultId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = InspectionResult.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all InspectionResult from db")
		except Exception:
			return None;
		
	def assignInspectionLot( self, inspectionResultId, inspectionLotId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.InspectionLotDelegate import InspectionLotDelegate

		errMsg = "Failed to assign element " + str(inspectionLotId) + " for InspectionLot on InspectionResult"

		try:
			# get the InspectionResult from db
			inspectionResult = self.get( inspectionResultId ).first()	
			
			# get the InspectionLot from db
			inspectionLot = InspectionLotDelegate().get(inspectionLotId).first();
			
			# assign the InspectionLot		
			inspectionResult.inspectionLot = inspectionLot
			
			#save it
			inspectionResult.save()

			# reload and return the appropriate version					
			return self.get( inspectionResultId );
		except InspectionResult.DoesNotExist:
			raise ProcessingError(errMsg + " : InspectionResult with id " + str(inspectionResultId) + " does not exist.")
		except InspectionLot.DoesNotExist:
			raise ProcessingError(errMsg + " : InspectionLot with id " + str(inspectionLotId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignInspectionLot( self, inspectionResultId ):
		errMsg = "Failed to unassign element " + str(inspectionLotId) + " for InspectionLot on InspectionResult"

		try:
			# get the InspectionResult from db
			inspectionResult = self.get( inspectionResultId ).first()	
			
			# assign to None for unassignment
			inspectionResult.inspectionLot = None			

			#save it
			inspectionResult.save()

			# reload and return the appropriate version					
			return self.get( inspectionResultId );
		except InspectionResult.DoesNotExist:
			raise ProcessingError(errMsg + " : InspectionResult with id " + str(inspectionResultId) + " does not exist.")
		except Exception:
			return None;
		
	def assignCharacteristic( self, inspectionResultId, characteristicId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.InspectionCharacteristicDelegate import InspectionCharacteristicDelegate

		errMsg = "Failed to assign element " + str(characteristicId) + " for Characteristic on InspectionResult"

		try:
			# get the InspectionResult from db
			inspectionResult = self.get( inspectionResultId ).first()	
			
			# get the InspectionCharacteristic from db
			inspectionCharacteristic = InspectionCharacteristicDelegate().get(characteristicId).first();
			
			# assign the Characteristic		
			inspectionResult.characteristic = inspectionCharacteristic
			
			#save it
			inspectionResult.save()

			# reload and return the appropriate version					
			return self.get( inspectionResultId );
		except InspectionResult.DoesNotExist:
			raise ProcessingError(errMsg + " : InspectionResult with id " + str(inspectionResultId) + " does not exist.")
		except InspectionCharacteristic.DoesNotExist:
			raise ProcessingError(errMsg + " : InspectionCharacteristic with id " + str(characteristicId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCharacteristic( self, inspectionResultId ):
		errMsg = "Failed to unassign element " + str(characteristicId) + " for Characteristic on InspectionResult"

		try:
			# get the InspectionResult from db
			inspectionResult = self.get( inspectionResultId ).first()	
			
			# assign to None for unassignment
			inspectionResult.inspectionCharacteristic = None			

			#save it
			inspectionResult.save()

			# reload and return the appropriate version					
			return self.get( inspectionResultId );
		except InspectionResult.DoesNotExist:
			raise ProcessingError(errMsg + " : InspectionResult with id " + str(inspectionResultId) + " does not exist.")
		except Exception:
			return None;
		
