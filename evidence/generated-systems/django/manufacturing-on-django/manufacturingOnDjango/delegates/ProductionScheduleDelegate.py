from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from manufacturingOnDjango.models.ProductionSchedule import ProductionSchedule
from manufacturingOnDjango.models.Plant import Plant
from manufacturingOnDjango.models.WorkOrder import WorkOrder
from manufacturingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model ProductionSchedule
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ProductionScheduleDelegate Declaration
#======================================================================
class ProductionScheduleDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, productionScheduleId ):
		try:	
			productionSchedule = ProductionSchedule.objects.filter(id=productionScheduleId)
			return productionSchedule.first();
		except ProductionSchedule.DoesNotExist:
			raise ProcessingError("ProductionSchedule with id " + str(productionScheduleId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, productionSchedule):
		for model in serializers.deserialize("json", productionSchedule):
			model.save()
			return model;

	def create(self, productionSchedule):
		productionSchedule.save()
		return productionSchedule;

	def saveFromJson(self, productionSchedule):
		for model in serializers.deserialize("json", productionSchedule):
			model.save()
			return productionSchedule;
	
	def save(self, productionSchedule):
		productionSchedule.save()
		return productionSchedule;
	
	def delete(self, productionScheduleId ):
		errMsg = "Failed to delete ProductionSchedule from db using id " + str(productionScheduleId)
		
		try:
			productionSchedule = ProductionSchedule.objects.get(id=productionScheduleId)
			productionSchedule.delete()
			return True
		except ProductionSchedule.DoesNotExist:
			raise ProcessingError("ProductionSchedule with id " + str(productionScheduleId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = ProductionSchedule.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all ProductionSchedule from db")
		except Exception:
			return None;
		
	def assignPlant( self, productionScheduleId, plantId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.PlantDelegate import PlantDelegate

		errMsg = "Failed to assign element " + str(plantId) + " for Plant on ProductionSchedule"

		try:
			# get the ProductionSchedule from db
			productionSchedule = self.get( productionScheduleId ).first()	
			
			# get the Plant from db
			plant = PlantDelegate().get(plantId).first();
			
			# assign the Plant		
			productionSchedule.plant = plant
			
			#save it
			productionSchedule.save()

			# reload and return the appropriate version					
			return self.get( productionScheduleId );
		except ProductionSchedule.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductionSchedule with id " + str(productionScheduleId) + " does not exist.")
		except Plant.DoesNotExist:
			raise ProcessingError(errMsg + " : Plant with id " + str(plantId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPlant( self, productionScheduleId ):
		errMsg = "Failed to unassign element " + str(plantId) + " for Plant on ProductionSchedule"

		try:
			# get the ProductionSchedule from db
			productionSchedule = self.get( productionScheduleId ).first()	
			
			# assign to None for unassignment
			productionSchedule.plant = None			

			#save it
			productionSchedule.save()

			# reload and return the appropriate version					
			return self.get( productionScheduleId );
		except ProductionSchedule.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductionSchedule with id " + str(productionScheduleId) + " does not exist.")
		except Exception:
			return None;
		
	def addWorkOrders( self, productionScheduleId, workOrdersIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.WorkOrderDelegate import WorkOrderDelegate

		errMsg = "Failed to add elements " + str(workOrdersIds) + " for WorkOrders on ProductionSchedule"

		try:
			# get the ProductionSchedule
			productionSchedule = self.get( productionScheduleId ).first()
				
			# split on a comma with no spaces
			idList = workOrdersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the WorkOrder		
				workOrder = WorkOrderDelegate().get(id).first();	
				# add the WorkOrder
				productionSchedule.workOrders.add(workOrder)
				
			# save it		
			productionSchedule.save()
			
			# reload and return the appropriate version
			return self.get( productionScheduleId );
		except ProductionSchedule.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductionSchedule with id " + str(productionScheduleId) + " does not exist.")
		except WorkOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : WorkOrder does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeWorkOrders( self, productionScheduleId, workOrdersIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.WorkOrderDelegate import WorkOrderDelegate

		errMsg = "Failed to remove elements " + str(workOrdersIds) + " for WorkOrders on ProductionSchedule"

		try:
			# get the ProductionSchedule
			productionSchedule = self.get( productionScheduleId ).first()
				
			# split on a comma with no spaces
			idList = workOrdersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the WorkOrder		
				workOrder = WorkOrderDelegate().get(id).first();	
				# add the WorkOrder
				productionSchedule.workOrders.remove(workOrder)
				
			# save it		
			productionSchedule.save()
			
			# reload and return the appropriate version
			return self.get( productionScheduleId );
		except ProductionSchedule.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductionSchedule with id " + str(productionScheduleId) + " does not exist.")
		except WorkOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : WorkOrder does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
