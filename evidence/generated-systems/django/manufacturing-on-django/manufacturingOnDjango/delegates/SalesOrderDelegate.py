from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from manufacturingOnDjango.models.SalesOrder import SalesOrder
from manufacturingOnDjango.models.Customer import Customer
from manufacturingOnDjango.models.Plant import Plant
from manufacturingOnDjango.models.SalesOrderLine import SalesOrderLine
from manufacturingOnDjango.models.WorkOrder import WorkOrder
from manufacturingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model SalesOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SalesOrderDelegate Declaration
#======================================================================
class SalesOrderDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, salesOrderId ):
		try:	
			salesOrder = SalesOrder.objects.filter(id=salesOrderId)
			return salesOrder.first();
		except SalesOrder.DoesNotExist:
			raise ProcessingError("SalesOrder with id " + str(salesOrderId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, salesOrder):
		for model in serializers.deserialize("json", salesOrder):
			model.save()
			return model;

	def create(self, salesOrder):
		salesOrder.save()
		return salesOrder;

	def saveFromJson(self, salesOrder):
		for model in serializers.deserialize("json", salesOrder):
			model.save()
			return salesOrder;
	
	def save(self, salesOrder):
		salesOrder.save()
		return salesOrder;
	
	def delete(self, salesOrderId ):
		errMsg = "Failed to delete SalesOrder from db using id " + str(salesOrderId)
		
		try:
			salesOrder = SalesOrder.objects.get(id=salesOrderId)
			salesOrder.delete()
			return True
		except SalesOrder.DoesNotExist:
			raise ProcessingError("SalesOrder with id " + str(salesOrderId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = SalesOrder.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all SalesOrder from db")
		except Exception:
			return None;
		
	def assignCustomer( self, salesOrderId, customerId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.CustomerDelegate import CustomerDelegate

		errMsg = "Failed to assign element " + str(customerId) + " for Customer on SalesOrder"

		try:
			# get the SalesOrder from db
			salesOrder = self.get( salesOrderId ).first()	
			
			# get the Customer from db
			customer = CustomerDelegate().get(customerId).first();
			
			# assign the Customer		
			salesOrder.customer = customer
			
			#save it
			salesOrder.save()

			# reload and return the appropriate version					
			return self.get( salesOrderId );
		except SalesOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : SalesOrder with id " + str(salesOrderId) + " does not exist.")
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCustomer( self, salesOrderId ):
		errMsg = "Failed to unassign element " + str(customerId) + " for Customer on SalesOrder"

		try:
			# get the SalesOrder from db
			salesOrder = self.get( salesOrderId ).first()	
			
			# assign to None for unassignment
			salesOrder.customer = None			

			#save it
			salesOrder.save()

			# reload and return the appropriate version					
			return self.get( salesOrderId );
		except SalesOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : SalesOrder with id " + str(salesOrderId) + " does not exist.")
		except Exception:
			return None;
		
	def assignPlant( self, salesOrderId, plantId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.PlantDelegate import PlantDelegate

		errMsg = "Failed to assign element " + str(plantId) + " for Plant on SalesOrder"

		try:
			# get the SalesOrder from db
			salesOrder = self.get( salesOrderId ).first()	
			
			# get the Plant from db
			plant = PlantDelegate().get(plantId).first();
			
			# assign the Plant		
			salesOrder.plant = plant
			
			#save it
			salesOrder.save()

			# reload and return the appropriate version					
			return self.get( salesOrderId );
		except SalesOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : SalesOrder with id " + str(salesOrderId) + " does not exist.")
		except Plant.DoesNotExist:
			raise ProcessingError(errMsg + " : Plant with id " + str(plantId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPlant( self, salesOrderId ):
		errMsg = "Failed to unassign element " + str(plantId) + " for Plant on SalesOrder"

		try:
			# get the SalesOrder from db
			salesOrder = self.get( salesOrderId ).first()	
			
			# assign to None for unassignment
			salesOrder.plant = None			

			#save it
			salesOrder.save()

			# reload and return the appropriate version					
			return self.get( salesOrderId );
		except SalesOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : SalesOrder with id " + str(salesOrderId) + " does not exist.")
		except Exception:
			return None;
		
	def addLines( self, salesOrderId, linesIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.SalesOrderLineDelegate import SalesOrderLineDelegate

		errMsg = "Failed to add elements " + str(linesIds) + " for Lines on SalesOrder"

		try:
			# get the SalesOrder
			salesOrder = self.get( salesOrderId ).first()
				
			# split on a comma with no spaces
			idList = linesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the SalesOrderLine		
				salesOrderLine = SalesOrderLineDelegate().get(id).first();	
				# add the SalesOrderLine
				salesOrder.lines.add(salesOrderLine)
				
			# save it		
			salesOrder.save()
			
			# reload and return the appropriate version
			return self.get( salesOrderId );
		except SalesOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : SalesOrder with id " + str(salesOrderId) + " does not exist.")
		except SalesOrderLine.DoesNotExist:
			raise ProcessingError(errMsg + " : SalesOrderLine does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeLines( self, salesOrderId, linesIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.SalesOrderLineDelegate import SalesOrderLineDelegate

		errMsg = "Failed to remove elements " + str(linesIds) + " for Lines on SalesOrder"

		try:
			# get the SalesOrder
			salesOrder = self.get( salesOrderId ).first()
				
			# split on a comma with no spaces
			idList = linesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the SalesOrderLine		
				salesOrderLine = SalesOrderLineDelegate().get(id).first();	
				# add the SalesOrderLine
				salesOrder.lines.remove(salesOrderLine)
				
			# save it		
			salesOrder.save()
			
			# reload and return the appropriate version
			return self.get( salesOrderId );
		except SalesOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : SalesOrder with id " + str(salesOrderId) + " does not exist.")
		except SalesOrderLine.DoesNotExist:
			raise ProcessingError(errMsg + " : SalesOrderLine does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addWorkOrders( self, salesOrderId, workOrdersIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.WorkOrderDelegate import WorkOrderDelegate

		errMsg = "Failed to add elements " + str(workOrdersIds) + " for WorkOrders on SalesOrder"

		try:
			# get the SalesOrder
			salesOrder = self.get( salesOrderId ).first()
				
			# split on a comma with no spaces
			idList = workOrdersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the WorkOrder		
				workOrder = WorkOrderDelegate().get(id).first();	
				# add the WorkOrder
				salesOrder.workOrders.add(workOrder)
				
			# save it		
			salesOrder.save()
			
			# reload and return the appropriate version
			return self.get( salesOrderId );
		except SalesOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : SalesOrder with id " + str(salesOrderId) + " does not exist.")
		except WorkOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : WorkOrder does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeWorkOrders( self, salesOrderId, workOrdersIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.WorkOrderDelegate import WorkOrderDelegate

		errMsg = "Failed to remove elements " + str(workOrdersIds) + " for WorkOrders on SalesOrder"

		try:
			# get the SalesOrder
			salesOrder = self.get( salesOrderId ).first()
				
			# split on a comma with no spaces
			idList = workOrdersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the WorkOrder		
				workOrder = WorkOrderDelegate().get(id).first();	
				# add the WorkOrder
				salesOrder.workOrders.remove(workOrder)
				
			# save it		
			salesOrder.save()
			
			# reload and return the appropriate version
			return self.get( salesOrderId );
		except SalesOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : SalesOrder with id " + str(salesOrderId) + " does not exist.")
		except WorkOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : WorkOrder does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
