from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from manufacturingOnDjango.models.InspectionPlan import InspectionPlan
from manufacturingOnDjango.models.Item import Item
from manufacturingOnDjango.models.InspectionCharacteristic import InspectionCharacteristic
from manufacturingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model InspectionPlan
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InspectionPlanDelegate Declaration
#======================================================================
class InspectionPlanDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, inspectionPlanId ):
		try:	
			inspectionPlan = InspectionPlan.objects.filter(id=inspectionPlanId)
			return inspectionPlan.first();
		except InspectionPlan.DoesNotExist:
			raise ProcessingError("InspectionPlan with id " + str(inspectionPlanId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, inspectionPlan):
		for model in serializers.deserialize("json", inspectionPlan):
			model.save()
			return model;

	def create(self, inspectionPlan):
		inspectionPlan.save()
		return inspectionPlan;

	def saveFromJson(self, inspectionPlan):
		for model in serializers.deserialize("json", inspectionPlan):
			model.save()
			return inspectionPlan;
	
	def save(self, inspectionPlan):
		inspectionPlan.save()
		return inspectionPlan;
	
	def delete(self, inspectionPlanId ):
		errMsg = "Failed to delete InspectionPlan from db using id " + str(inspectionPlanId)
		
		try:
			inspectionPlan = InspectionPlan.objects.get(id=inspectionPlanId)
			inspectionPlan.delete()
			return True
		except InspectionPlan.DoesNotExist:
			raise ProcessingError("InspectionPlan with id " + str(inspectionPlanId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = InspectionPlan.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all InspectionPlan from db")
		except Exception:
			return None;
		
	def assignItem( self, inspectionPlanId, itemId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.ItemDelegate import ItemDelegate

		errMsg = "Failed to assign element " + str(itemId) + " for Item on InspectionPlan"

		try:
			# get the InspectionPlan from db
			inspectionPlan = self.get( inspectionPlanId ).first()	
			
			# get the Item from db
			item = ItemDelegate().get(itemId).first();
			
			# assign the Item		
			inspectionPlan.item = item
			
			#save it
			inspectionPlan.save()

			# reload and return the appropriate version					
			return self.get( inspectionPlanId );
		except InspectionPlan.DoesNotExist:
			raise ProcessingError(errMsg + " : InspectionPlan with id " + str(inspectionPlanId) + " does not exist.")
		except Item.DoesNotExist:
			raise ProcessingError(errMsg + " : Item with id " + str(itemId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignItem( self, inspectionPlanId ):
		errMsg = "Failed to unassign element " + str(itemId) + " for Item on InspectionPlan"

		try:
			# get the InspectionPlan from db
			inspectionPlan = self.get( inspectionPlanId ).first()	
			
			# assign to None for unassignment
			inspectionPlan.item = None			

			#save it
			inspectionPlan.save()

			# reload and return the appropriate version					
			return self.get( inspectionPlanId );
		except InspectionPlan.DoesNotExist:
			raise ProcessingError(errMsg + " : InspectionPlan with id " + str(inspectionPlanId) + " does not exist.")
		except Exception:
			return None;
		
	def addCharacteristics( self, inspectionPlanId, characteristicsIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.InspectionCharacteristicDelegate import InspectionCharacteristicDelegate

		errMsg = "Failed to add elements " + str(characteristicsIds) + " for Characteristics on InspectionPlan"

		try:
			# get the InspectionPlan
			inspectionPlan = self.get( inspectionPlanId ).first()
				
			# split on a comma with no spaces
			idList = characteristicsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the InspectionCharacteristic		
				inspectionCharacteristic = InspectionCharacteristicDelegate().get(id).first();	
				# add the InspectionCharacteristic
				inspectionPlan.characteristics.add(inspectionCharacteristic)
				
			# save it		
			inspectionPlan.save()
			
			# reload and return the appropriate version
			return self.get( inspectionPlanId );
		except InspectionPlan.DoesNotExist:
			raise ProcessingError(errMsg + " : InspectionPlan with id " + str(inspectionPlanId) + " does not exist.")
		except InspectionCharacteristic.DoesNotExist:
			raise ProcessingError(errMsg + " : InspectionCharacteristic does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCharacteristics( self, inspectionPlanId, characteristicsIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.InspectionCharacteristicDelegate import InspectionCharacteristicDelegate

		errMsg = "Failed to remove elements " + str(characteristicsIds) + " for Characteristics on InspectionPlan"

		try:
			# get the InspectionPlan
			inspectionPlan = self.get( inspectionPlanId ).first()
				
			# split on a comma with no spaces
			idList = characteristicsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the InspectionCharacteristic		
				inspectionCharacteristic = InspectionCharacteristicDelegate().get(id).first();	
				# add the InspectionCharacteristic
				inspectionPlan.characteristics.remove(inspectionCharacteristic)
				
			# save it		
			inspectionPlan.save()
			
			# reload and return the appropriate version
			return self.get( inspectionPlanId );
		except InspectionPlan.DoesNotExist:
			raise ProcessingError(errMsg + " : InspectionPlan with id " + str(inspectionPlanId) + " does not exist.")
		except InspectionCharacteristic.DoesNotExist:
			raise ProcessingError(errMsg + " : InspectionCharacteristic does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
