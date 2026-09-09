from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from manufacturingOnDjango.models.PlannedOrder import PlannedOrder
from manufacturingOnDjango.models.MRPRun import MRPRun
from manufacturingOnDjango.models.Item import Item
from manufacturingOnDjango.models.Plant import Plant
from manufacturingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model PlannedOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PlannedOrderDelegate Declaration
#======================================================================
class PlannedOrderDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, plannedOrderId ):
		try:	
			plannedOrder = PlannedOrder.objects.filter(id=plannedOrderId)
			return plannedOrder.first();
		except PlannedOrder.DoesNotExist:
			raise ProcessingError("PlannedOrder with id " + str(plannedOrderId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, plannedOrder):
		for model in serializers.deserialize("json", plannedOrder):
			model.save()
			return model;

	def create(self, plannedOrder):
		plannedOrder.save()
		return plannedOrder;

	def saveFromJson(self, plannedOrder):
		for model in serializers.deserialize("json", plannedOrder):
			model.save()
			return plannedOrder;
	
	def save(self, plannedOrder):
		plannedOrder.save()
		return plannedOrder;
	
	def delete(self, plannedOrderId ):
		errMsg = "Failed to delete PlannedOrder from db using id " + str(plannedOrderId)
		
		try:
			plannedOrder = PlannedOrder.objects.get(id=plannedOrderId)
			plannedOrder.delete()
			return True
		except PlannedOrder.DoesNotExist:
			raise ProcessingError("PlannedOrder with id " + str(plannedOrderId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = PlannedOrder.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all PlannedOrder from db")
		except Exception:
			return None;
		
	def assignMrpRun( self, plannedOrderId, mrpRunId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.MRPRunDelegate import MRPRunDelegate

		errMsg = "Failed to assign element " + str(mrpRunId) + " for MrpRun on PlannedOrder"

		try:
			# get the PlannedOrder from db
			plannedOrder = self.get( plannedOrderId ).first()	
			
			# get the MRPRun from db
			mRPRun = MRPRunDelegate().get(mrpRunId).first();
			
			# assign the MrpRun		
			plannedOrder.mrpRun = mRPRun
			
			#save it
			plannedOrder.save()

			# reload and return the appropriate version					
			return self.get( plannedOrderId );
		except PlannedOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : PlannedOrder with id " + str(plannedOrderId) + " does not exist.")
		except MRPRun.DoesNotExist:
			raise ProcessingError(errMsg + " : MRPRun with id " + str(mrpRunId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignMrpRun( self, plannedOrderId ):
		errMsg = "Failed to unassign element " + str(mrpRunId) + " for MrpRun on PlannedOrder"

		try:
			# get the PlannedOrder from db
			plannedOrder = self.get( plannedOrderId ).first()	
			
			# assign to None for unassignment
			plannedOrder.mRPRun = None			

			#save it
			plannedOrder.save()

			# reload and return the appropriate version					
			return self.get( plannedOrderId );
		except PlannedOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : PlannedOrder with id " + str(plannedOrderId) + " does not exist.")
		except Exception:
			return None;
		
	def assignItem( self, plannedOrderId, itemId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.ItemDelegate import ItemDelegate

		errMsg = "Failed to assign element " + str(itemId) + " for Item on PlannedOrder"

		try:
			# get the PlannedOrder from db
			plannedOrder = self.get( plannedOrderId ).first()	
			
			# get the Item from db
			item = ItemDelegate().get(itemId).first();
			
			# assign the Item		
			plannedOrder.item = item
			
			#save it
			plannedOrder.save()

			# reload and return the appropriate version					
			return self.get( plannedOrderId );
		except PlannedOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : PlannedOrder with id " + str(plannedOrderId) + " does not exist.")
		except Item.DoesNotExist:
			raise ProcessingError(errMsg + " : Item with id " + str(itemId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignItem( self, plannedOrderId ):
		errMsg = "Failed to unassign element " + str(itemId) + " for Item on PlannedOrder"

		try:
			# get the PlannedOrder from db
			plannedOrder = self.get( plannedOrderId ).first()	
			
			# assign to None for unassignment
			plannedOrder.item = None			

			#save it
			plannedOrder.save()

			# reload and return the appropriate version					
			return self.get( plannedOrderId );
		except PlannedOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : PlannedOrder with id " + str(plannedOrderId) + " does not exist.")
		except Exception:
			return None;
		
	def assignPlant( self, plannedOrderId, plantId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.PlantDelegate import PlantDelegate

		errMsg = "Failed to assign element " + str(plantId) + " for Plant on PlannedOrder"

		try:
			# get the PlannedOrder from db
			plannedOrder = self.get( plannedOrderId ).first()	
			
			# get the Plant from db
			plant = PlantDelegate().get(plantId).first();
			
			# assign the Plant		
			plannedOrder.plant = plant
			
			#save it
			plannedOrder.save()

			# reload and return the appropriate version					
			return self.get( plannedOrderId );
		except PlannedOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : PlannedOrder with id " + str(plannedOrderId) + " does not exist.")
		except Plant.DoesNotExist:
			raise ProcessingError(errMsg + " : Plant with id " + str(plantId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPlant( self, plannedOrderId ):
		errMsg = "Failed to unassign element " + str(plantId) + " for Plant on PlannedOrder"

		try:
			# get the PlannedOrder from db
			plannedOrder = self.get( plannedOrderId ).first()	
			
			# assign to None for unassignment
			plannedOrder.plant = None			

			#save it
			plannedOrder.save()

			# reload and return the appropriate version					
			return self.get( plannedOrderId );
		except PlannedOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : PlannedOrder with id " + str(plannedOrderId) + " does not exist.")
		except Exception:
			return None;
		
