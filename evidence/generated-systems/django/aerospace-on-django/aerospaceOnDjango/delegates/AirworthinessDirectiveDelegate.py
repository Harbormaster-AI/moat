from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from aerospaceOnDjango.models.AirworthinessDirective import AirworthinessDirective
from aerospaceOnDjango.models.MaintenanceWorkOrder import MaintenanceWorkOrder
from aerospaceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model AirworthinessDirective
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AirworthinessDirectiveDelegate Declaration
#======================================================================
class AirworthinessDirectiveDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, airworthinessDirectiveId ):
		try:	
			airworthinessDirective = AirworthinessDirective.objects.filter(id=airworthinessDirectiveId)
			return airworthinessDirective.first();
		except AirworthinessDirective.DoesNotExist:
			raise ProcessingError("AirworthinessDirective with id " + str(airworthinessDirectiveId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, airworthinessDirective):
		for model in serializers.deserialize("json", airworthinessDirective):
			model.save()
			return model;

	def create(self, airworthinessDirective):
		airworthinessDirective.save()
		return airworthinessDirective;

	def saveFromJson(self, airworthinessDirective):
		for model in serializers.deserialize("json", airworthinessDirective):
			model.save()
			return airworthinessDirective;
	
	def save(self, airworthinessDirective):
		airworthinessDirective.save()
		return airworthinessDirective;
	
	def delete(self, airworthinessDirectiveId ):
		errMsg = "Failed to delete AirworthinessDirective from db using id " + str(airworthinessDirectiveId)
		
		try:
			airworthinessDirective = AirworthinessDirective.objects.get(id=airworthinessDirectiveId)
			airworthinessDirective.delete()
			return True
		except AirworthinessDirective.DoesNotExist:
			raise ProcessingError("AirworthinessDirective with id " + str(airworthinessDirectiveId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = AirworthinessDirective.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all AirworthinessDirective from db")
		except Exception:
			return None;
		
	def addWorkOrders( self, airworthinessDirectiveId, workOrdersIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.MaintenanceWorkOrderDelegate import MaintenanceWorkOrderDelegate

		errMsg = "Failed to add elements " + str(workOrdersIds) + " for WorkOrders on AirworthinessDirective"

		try:
			# get the AirworthinessDirective
			airworthinessDirective = self.get( airworthinessDirectiveId ).first()
				
			# split on a comma with no spaces
			idList = workOrdersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the MaintenanceWorkOrder		
				maintenanceWorkOrder = MaintenanceWorkOrderDelegate().get(id).first();	
				# add the MaintenanceWorkOrder
				airworthinessDirective.workOrders.add(maintenanceWorkOrder)
				
			# save it		
			airworthinessDirective.save()
			
			# reload and return the appropriate version
			return self.get( airworthinessDirectiveId );
		except AirworthinessDirective.DoesNotExist:
			raise ProcessingError(errMsg + " : AirworthinessDirective with id " + str(airworthinessDirectiveId) + " does not exist.")
		except MaintenanceWorkOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : MaintenanceWorkOrder does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeWorkOrders( self, airworthinessDirectiveId, workOrdersIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.MaintenanceWorkOrderDelegate import MaintenanceWorkOrderDelegate

		errMsg = "Failed to remove elements " + str(workOrdersIds) + " for WorkOrders on AirworthinessDirective"

		try:
			# get the AirworthinessDirective
			airworthinessDirective = self.get( airworthinessDirectiveId ).first()
				
			# split on a comma with no spaces
			idList = workOrdersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the MaintenanceWorkOrder		
				maintenanceWorkOrder = MaintenanceWorkOrderDelegate().get(id).first();	
				# add the MaintenanceWorkOrder
				airworthinessDirective.workOrders.remove(maintenanceWorkOrder)
				
			# save it		
			airworthinessDirective.save()
			
			# reload and return the appropriate version
			return self.get( airworthinessDirectiveId );
		except AirworthinessDirective.DoesNotExist:
			raise ProcessingError(errMsg + " : AirworthinessDirective with id " + str(airworthinessDirectiveId) + " does not exist.")
		except MaintenanceWorkOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : MaintenanceWorkOrder does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
