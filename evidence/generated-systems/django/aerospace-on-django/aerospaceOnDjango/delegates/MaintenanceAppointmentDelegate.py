from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from aerospaceOnDjango.models.MaintenanceAppointment import MaintenanceAppointment
from aerospaceOnDjango.models.Aircraft import Aircraft
from aerospaceOnDjango.models.MROFacility import MROFacility
from aerospaceOnDjango.models.MaintenanceWorkOrder import MaintenanceWorkOrder
from aerospaceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model MaintenanceAppointment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MaintenanceAppointmentDelegate Declaration
#======================================================================
class MaintenanceAppointmentDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, maintenanceAppointmentId ):
		try:	
			maintenanceAppointment = MaintenanceAppointment.objects.filter(id=maintenanceAppointmentId)
			return maintenanceAppointment.first();
		except MaintenanceAppointment.DoesNotExist:
			raise ProcessingError("MaintenanceAppointment with id " + str(maintenanceAppointmentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, maintenanceAppointment):
		for model in serializers.deserialize("json", maintenanceAppointment):
			model.save()
			return model;

	def create(self, maintenanceAppointment):
		maintenanceAppointment.save()
		return maintenanceAppointment;

	def saveFromJson(self, maintenanceAppointment):
		for model in serializers.deserialize("json", maintenanceAppointment):
			model.save()
			return maintenanceAppointment;
	
	def save(self, maintenanceAppointment):
		maintenanceAppointment.save()
		return maintenanceAppointment;
	
	def delete(self, maintenanceAppointmentId ):
		errMsg = "Failed to delete MaintenanceAppointment from db using id " + str(maintenanceAppointmentId)
		
		try:
			maintenanceAppointment = MaintenanceAppointment.objects.get(id=maintenanceAppointmentId)
			maintenanceAppointment.delete()
			return True
		except MaintenanceAppointment.DoesNotExist:
			raise ProcessingError("MaintenanceAppointment with id " + str(maintenanceAppointmentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = MaintenanceAppointment.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all MaintenanceAppointment from db")
		except Exception:
			return None;
		
	def assignAircraft( self, maintenanceAppointmentId, aircraftId ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftDelegate import AircraftDelegate

		errMsg = "Failed to assign element " + str(aircraftId) + " for Aircraft on MaintenanceAppointment"

		try:
			# get the MaintenanceAppointment from db
			maintenanceAppointment = self.get( maintenanceAppointmentId ).first()	
			
			# get the Aircraft from db
			aircraft = AircraftDelegate().get(aircraftId).first();
			
			# assign the Aircraft		
			maintenanceAppointment.aircraft = aircraft
			
			#save it
			maintenanceAppointment.save()

			# reload and return the appropriate version					
			return self.get( maintenanceAppointmentId );
		except MaintenanceAppointment.DoesNotExist:
			raise ProcessingError(errMsg + " : MaintenanceAppointment with id " + str(maintenanceAppointmentId) + " does not exist.")
		except Aircraft.DoesNotExist:
			raise ProcessingError(errMsg + " : Aircraft with id " + str(aircraftId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignAircraft( self, maintenanceAppointmentId ):
		errMsg = "Failed to unassign element " + str(aircraftId) + " for Aircraft on MaintenanceAppointment"

		try:
			# get the MaintenanceAppointment from db
			maintenanceAppointment = self.get( maintenanceAppointmentId ).first()	
			
			# assign to None for unassignment
			maintenanceAppointment.aircraft = None			

			#save it
			maintenanceAppointment.save()

			# reload and return the appropriate version					
			return self.get( maintenanceAppointmentId );
		except MaintenanceAppointment.DoesNotExist:
			raise ProcessingError(errMsg + " : MaintenanceAppointment with id " + str(maintenanceAppointmentId) + " does not exist.")
		except Exception:
			return None;
		
	def assignMroFacility( self, maintenanceAppointmentId, mroFacilityId ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.MROFacilityDelegate import MROFacilityDelegate

		errMsg = "Failed to assign element " + str(mroFacilityId) + " for MroFacility on MaintenanceAppointment"

		try:
			# get the MaintenanceAppointment from db
			maintenanceAppointment = self.get( maintenanceAppointmentId ).first()	
			
			# get the MROFacility from db
			mROFacility = MROFacilityDelegate().get(mroFacilityId).first();
			
			# assign the MroFacility		
			maintenanceAppointment.mroFacility = mROFacility
			
			#save it
			maintenanceAppointment.save()

			# reload and return the appropriate version					
			return self.get( maintenanceAppointmentId );
		except MaintenanceAppointment.DoesNotExist:
			raise ProcessingError(errMsg + " : MaintenanceAppointment with id " + str(maintenanceAppointmentId) + " does not exist.")
		except MROFacility.DoesNotExist:
			raise ProcessingError(errMsg + " : MROFacility with id " + str(mroFacilityId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignMroFacility( self, maintenanceAppointmentId ):
		errMsg = "Failed to unassign element " + str(mroFacilityId) + " for MroFacility on MaintenanceAppointment"

		try:
			# get the MaintenanceAppointment from db
			maintenanceAppointment = self.get( maintenanceAppointmentId ).first()	
			
			# assign to None for unassignment
			maintenanceAppointment.mROFacility = None			

			#save it
			maintenanceAppointment.save()

			# reload and return the appropriate version					
			return self.get( maintenanceAppointmentId );
		except MaintenanceAppointment.DoesNotExist:
			raise ProcessingError(errMsg + " : MaintenanceAppointment with id " + str(maintenanceAppointmentId) + " does not exist.")
		except Exception:
			return None;
		
	def assignWorkOrder( self, maintenanceAppointmentId, workOrderId ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.MaintenanceWorkOrderDelegate import MaintenanceWorkOrderDelegate

		errMsg = "Failed to assign element " + str(workOrderId) + " for WorkOrder on MaintenanceAppointment"

		try:
			# get the MaintenanceAppointment from db
			maintenanceAppointment = self.get( maintenanceAppointmentId ).first()	
			
			# get the MaintenanceWorkOrder from db
			maintenanceWorkOrder = MaintenanceWorkOrderDelegate().get(workOrderId).first();
			
			# assign the WorkOrder		
			maintenanceAppointment.workOrder = maintenanceWorkOrder
			
			#save it
			maintenanceAppointment.save()

			# reload and return the appropriate version					
			return self.get( maintenanceAppointmentId );
		except MaintenanceAppointment.DoesNotExist:
			raise ProcessingError(errMsg + " : MaintenanceAppointment with id " + str(maintenanceAppointmentId) + " does not exist.")
		except MaintenanceWorkOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : MaintenanceWorkOrder with id " + str(workOrderId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignWorkOrder( self, maintenanceAppointmentId ):
		errMsg = "Failed to unassign element " + str(workOrderId) + " for WorkOrder on MaintenanceAppointment"

		try:
			# get the MaintenanceAppointment from db
			maintenanceAppointment = self.get( maintenanceAppointmentId ).first()	
			
			# assign to None for unassignment
			maintenanceAppointment.maintenanceWorkOrder = None			

			#save it
			maintenanceAppointment.save()

			# reload and return the appropriate version					
			return self.get( maintenanceAppointmentId );
		except MaintenanceAppointment.DoesNotExist:
			raise ProcessingError(errMsg + " : MaintenanceAppointment with id " + str(maintenanceAppointmentId) + " does not exist.")
		except Exception:
			return None;
		
