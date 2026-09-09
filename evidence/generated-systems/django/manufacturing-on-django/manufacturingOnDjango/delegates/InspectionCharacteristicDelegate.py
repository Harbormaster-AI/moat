from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from manufacturingOnDjango.models.InspectionCharacteristic import InspectionCharacteristic
from manufacturingOnDjango.models.InspectionPlan import InspectionPlan
from manufacturingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model InspectionCharacteristic
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InspectionCharacteristicDelegate Declaration
#======================================================================
class InspectionCharacteristicDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, inspectionCharacteristicId ):
		try:	
			inspectionCharacteristic = InspectionCharacteristic.objects.filter(id=inspectionCharacteristicId)
			return inspectionCharacteristic.first();
		except InspectionCharacteristic.DoesNotExist:
			raise ProcessingError("InspectionCharacteristic with id " + str(inspectionCharacteristicId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, inspectionCharacteristic):
		for model in serializers.deserialize("json", inspectionCharacteristic):
			model.save()
			return model;

	def create(self, inspectionCharacteristic):
		inspectionCharacteristic.save()
		return inspectionCharacteristic;

	def saveFromJson(self, inspectionCharacteristic):
		for model in serializers.deserialize("json", inspectionCharacteristic):
			model.save()
			return inspectionCharacteristic;
	
	def save(self, inspectionCharacteristic):
		inspectionCharacteristic.save()
		return inspectionCharacteristic;
	
	def delete(self, inspectionCharacteristicId ):
		errMsg = "Failed to delete InspectionCharacteristic from db using id " + str(inspectionCharacteristicId)
		
		try:
			inspectionCharacteristic = InspectionCharacteristic.objects.get(id=inspectionCharacteristicId)
			inspectionCharacteristic.delete()
			return True
		except InspectionCharacteristic.DoesNotExist:
			raise ProcessingError("InspectionCharacteristic with id " + str(inspectionCharacteristicId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = InspectionCharacteristic.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all InspectionCharacteristic from db")
		except Exception:
			return None;
		
	def assignInspectionPlan( self, inspectionCharacteristicId, inspectionPlanId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.InspectionPlanDelegate import InspectionPlanDelegate

		errMsg = "Failed to assign element " + str(inspectionPlanId) + " for InspectionPlan on InspectionCharacteristic"

		try:
			# get the InspectionCharacteristic from db
			inspectionCharacteristic = self.get( inspectionCharacteristicId ).first()	
			
			# get the InspectionPlan from db
			inspectionPlan = InspectionPlanDelegate().get(inspectionPlanId).first();
			
			# assign the InspectionPlan		
			inspectionCharacteristic.inspectionPlan = inspectionPlan
			
			#save it
			inspectionCharacteristic.save()

			# reload and return the appropriate version					
			return self.get( inspectionCharacteristicId );
		except InspectionCharacteristic.DoesNotExist:
			raise ProcessingError(errMsg + " : InspectionCharacteristic with id " + str(inspectionCharacteristicId) + " does not exist.")
		except InspectionPlan.DoesNotExist:
			raise ProcessingError(errMsg + " : InspectionPlan with id " + str(inspectionPlanId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignInspectionPlan( self, inspectionCharacteristicId ):
		errMsg = "Failed to unassign element " + str(inspectionPlanId) + " for InspectionPlan on InspectionCharacteristic"

		try:
			# get the InspectionCharacteristic from db
			inspectionCharacteristic = self.get( inspectionCharacteristicId ).first()	
			
			# assign to None for unassignment
			inspectionCharacteristic.inspectionPlan = None			

			#save it
			inspectionCharacteristic.save()

			# reload and return the appropriate version					
			return self.get( inspectionCharacteristicId );
		except InspectionCharacteristic.DoesNotExist:
			raise ProcessingError(errMsg + " : InspectionCharacteristic with id " + str(inspectionCharacteristicId) + " does not exist.")
		except Exception:
			return None;
		
