from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from aerospaceOnDjango.models.MaintenanceWorkOrder import MaintenanceWorkOrder
from aerospaceOnDjango.models.Aircraft import Aircraft
from aerospaceOnDjango.models.AirworthinessDirective import AirworthinessDirective
from aerospaceOnDjango.models.ServiceBulletin import ServiceBulletin
from aerospaceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model MaintenanceWorkOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MaintenanceWorkOrderDelegate Declaration
#======================================================================
class MaintenanceWorkOrderDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, maintenanceWorkOrderId ):
		try:	
			maintenanceWorkOrder = MaintenanceWorkOrder.objects.filter(id=maintenanceWorkOrderId)
			return maintenanceWorkOrder.first();
		except MaintenanceWorkOrder.DoesNotExist:
			raise ProcessingError("MaintenanceWorkOrder with id " + str(maintenanceWorkOrderId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, maintenanceWorkOrder):
		for model in serializers.deserialize("json", maintenanceWorkOrder):
			model.save()
			return model;

	def create(self, maintenanceWorkOrder):
		maintenanceWorkOrder.save()
		return maintenanceWorkOrder;

	def saveFromJson(self, maintenanceWorkOrder):
		for model in serializers.deserialize("json", maintenanceWorkOrder):
			model.save()
			return maintenanceWorkOrder;
	
	def save(self, maintenanceWorkOrder):
		maintenanceWorkOrder.save()
		return maintenanceWorkOrder;
	
	def delete(self, maintenanceWorkOrderId ):
		errMsg = "Failed to delete MaintenanceWorkOrder from db using id " + str(maintenanceWorkOrderId)
		
		try:
			maintenanceWorkOrder = MaintenanceWorkOrder.objects.get(id=maintenanceWorkOrderId)
			maintenanceWorkOrder.delete()
			return True
		except MaintenanceWorkOrder.DoesNotExist:
			raise ProcessingError("MaintenanceWorkOrder with id " + str(maintenanceWorkOrderId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = MaintenanceWorkOrder.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all MaintenanceWorkOrder from db")
		except Exception:
			return None;
		
	def assignAircraft( self, maintenanceWorkOrderId, aircraftId ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftDelegate import AircraftDelegate

		errMsg = "Failed to assign element " + str(aircraftId) + " for Aircraft on MaintenanceWorkOrder"

		try:
			# get the MaintenanceWorkOrder from db
			maintenanceWorkOrder = self.get( maintenanceWorkOrderId ).first()	
			
			# get the Aircraft from db
			aircraft = AircraftDelegate().get(aircraftId).first();
			
			# assign the Aircraft		
			maintenanceWorkOrder.aircraft = aircraft
			
			#save it
			maintenanceWorkOrder.save()

			# reload and return the appropriate version					
			return self.get( maintenanceWorkOrderId );
		except MaintenanceWorkOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : MaintenanceWorkOrder with id " + str(maintenanceWorkOrderId) + " does not exist.")
		except Aircraft.DoesNotExist:
			raise ProcessingError(errMsg + " : Aircraft with id " + str(aircraftId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignAircraft( self, maintenanceWorkOrderId ):
		errMsg = "Failed to unassign element " + str(aircraftId) + " for Aircraft on MaintenanceWorkOrder"

		try:
			# get the MaintenanceWorkOrder from db
			maintenanceWorkOrder = self.get( maintenanceWorkOrderId ).first()	
			
			# assign to None for unassignment
			maintenanceWorkOrder.aircraft = None			

			#save it
			maintenanceWorkOrder.save()

			# reload and return the appropriate version					
			return self.get( maintenanceWorkOrderId );
		except MaintenanceWorkOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : MaintenanceWorkOrder with id " + str(maintenanceWorkOrderId) + " does not exist.")
		except Exception:
			return None;
		
	def assignAirworthinessDirective( self, maintenanceWorkOrderId, airworthinessDirectiveId ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AirworthinessDirectiveDelegate import AirworthinessDirectiveDelegate

		errMsg = "Failed to assign element " + str(airworthinessDirectiveId) + " for AirworthinessDirective on MaintenanceWorkOrder"

		try:
			# get the MaintenanceWorkOrder from db
			maintenanceWorkOrder = self.get( maintenanceWorkOrderId ).first()	
			
			# get the AirworthinessDirective from db
			airworthinessDirective = AirworthinessDirectiveDelegate().get(airworthinessDirectiveId).first();
			
			# assign the AirworthinessDirective		
			maintenanceWorkOrder.airworthinessDirective = airworthinessDirective
			
			#save it
			maintenanceWorkOrder.save()

			# reload and return the appropriate version					
			return self.get( maintenanceWorkOrderId );
		except MaintenanceWorkOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : MaintenanceWorkOrder with id " + str(maintenanceWorkOrderId) + " does not exist.")
		except AirworthinessDirective.DoesNotExist:
			raise ProcessingError(errMsg + " : AirworthinessDirective with id " + str(airworthinessDirectiveId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignAirworthinessDirective( self, maintenanceWorkOrderId ):
		errMsg = "Failed to unassign element " + str(airworthinessDirectiveId) + " for AirworthinessDirective on MaintenanceWorkOrder"

		try:
			# get the MaintenanceWorkOrder from db
			maintenanceWorkOrder = self.get( maintenanceWorkOrderId ).first()	
			
			# assign to None for unassignment
			maintenanceWorkOrder.airworthinessDirective = None			

			#save it
			maintenanceWorkOrder.save()

			# reload and return the appropriate version					
			return self.get( maintenanceWorkOrderId );
		except MaintenanceWorkOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : MaintenanceWorkOrder with id " + str(maintenanceWorkOrderId) + " does not exist.")
		except Exception:
			return None;
		
	def assignServiceBulletin( self, maintenanceWorkOrderId, serviceBulletinId ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.ServiceBulletinDelegate import ServiceBulletinDelegate

		errMsg = "Failed to assign element " + str(serviceBulletinId) + " for ServiceBulletin on MaintenanceWorkOrder"

		try:
			# get the MaintenanceWorkOrder from db
			maintenanceWorkOrder = self.get( maintenanceWorkOrderId ).first()	
			
			# get the ServiceBulletin from db
			serviceBulletin = ServiceBulletinDelegate().get(serviceBulletinId).first();
			
			# assign the ServiceBulletin		
			maintenanceWorkOrder.serviceBulletin = serviceBulletin
			
			#save it
			maintenanceWorkOrder.save()

			# reload and return the appropriate version					
			return self.get( maintenanceWorkOrderId );
		except MaintenanceWorkOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : MaintenanceWorkOrder with id " + str(maintenanceWorkOrderId) + " does not exist.")
		except ServiceBulletin.DoesNotExist:
			raise ProcessingError(errMsg + " : ServiceBulletin with id " + str(serviceBulletinId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignServiceBulletin( self, maintenanceWorkOrderId ):
		errMsg = "Failed to unassign element " + str(serviceBulletinId) + " for ServiceBulletin on MaintenanceWorkOrder"

		try:
			# get the MaintenanceWorkOrder from db
			maintenanceWorkOrder = self.get( maintenanceWorkOrderId ).first()	
			
			# assign to None for unassignment
			maintenanceWorkOrder.serviceBulletin = None			

			#save it
			maintenanceWorkOrder.save()

			# reload and return the appropriate version					
			return self.get( maintenanceWorkOrderId );
		except MaintenanceWorkOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : MaintenanceWorkOrder with id " + str(maintenanceWorkOrderId) + " does not exist.")
		except Exception:
			return None;
		
