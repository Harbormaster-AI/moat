from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from aerospaceOnDjango.models.ServiceBulletin import ServiceBulletin
from aerospaceOnDjango.models.MaintenanceWorkOrder import MaintenanceWorkOrder
from aerospaceOnDjango.models.AircraftVariant import AircraftVariant
from aerospaceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model ServiceBulletin
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ServiceBulletinDelegate Declaration
#======================================================================
class ServiceBulletinDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, serviceBulletinId ):
		try:	
			serviceBulletin = ServiceBulletin.objects.filter(id=serviceBulletinId)
			return serviceBulletin.first();
		except ServiceBulletin.DoesNotExist:
			raise ProcessingError("ServiceBulletin with id " + str(serviceBulletinId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, serviceBulletin):
		for model in serializers.deserialize("json", serviceBulletin):
			model.save()
			return model;

	def create(self, serviceBulletin):
		serviceBulletin.save()
		return serviceBulletin;

	def saveFromJson(self, serviceBulletin):
		for model in serializers.deserialize("json", serviceBulletin):
			model.save()
			return serviceBulletin;
	
	def save(self, serviceBulletin):
		serviceBulletin.save()
		return serviceBulletin;
	
	def delete(self, serviceBulletinId ):
		errMsg = "Failed to delete ServiceBulletin from db using id " + str(serviceBulletinId)
		
		try:
			serviceBulletin = ServiceBulletin.objects.get(id=serviceBulletinId)
			serviceBulletin.delete()
			return True
		except ServiceBulletin.DoesNotExist:
			raise ProcessingError("ServiceBulletin with id " + str(serviceBulletinId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = ServiceBulletin.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all ServiceBulletin from db")
		except Exception:
			return None;
		
	def addWorkOrders( self, serviceBulletinId, workOrdersIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.MaintenanceWorkOrderDelegate import MaintenanceWorkOrderDelegate

		errMsg = "Failed to add elements " + str(workOrdersIds) + " for WorkOrders on ServiceBulletin"

		try:
			# get the ServiceBulletin
			serviceBulletin = self.get( serviceBulletinId ).first()
				
			# split on a comma with no spaces
			idList = workOrdersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the MaintenanceWorkOrder		
				maintenanceWorkOrder = MaintenanceWorkOrderDelegate().get(id).first();	
				# add the MaintenanceWorkOrder
				serviceBulletin.workOrders.add(maintenanceWorkOrder)
				
			# save it		
			serviceBulletin.save()
			
			# reload and return the appropriate version
			return self.get( serviceBulletinId );
		except ServiceBulletin.DoesNotExist:
			raise ProcessingError(errMsg + " : ServiceBulletin with id " + str(serviceBulletinId) + " does not exist.")
		except MaintenanceWorkOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : MaintenanceWorkOrder does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeWorkOrders( self, serviceBulletinId, workOrdersIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.MaintenanceWorkOrderDelegate import MaintenanceWorkOrderDelegate

		errMsg = "Failed to remove elements " + str(workOrdersIds) + " for WorkOrders on ServiceBulletin"

		try:
			# get the ServiceBulletin
			serviceBulletin = self.get( serviceBulletinId ).first()
				
			# split on a comma with no spaces
			idList = workOrdersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the MaintenanceWorkOrder		
				maintenanceWorkOrder = MaintenanceWorkOrderDelegate().get(id).first();	
				# add the MaintenanceWorkOrder
				serviceBulletin.workOrders.remove(maintenanceWorkOrder)
				
			# save it		
			serviceBulletin.save()
			
			# reload and return the appropriate version
			return self.get( serviceBulletinId );
		except ServiceBulletin.DoesNotExist:
			raise ProcessingError(errMsg + " : ServiceBulletin with id " + str(serviceBulletinId) + " does not exist.")
		except MaintenanceWorkOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : MaintenanceWorkOrder does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addVariants( self, serviceBulletinId, variantsIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftVariantDelegate import AircraftVariantDelegate

		errMsg = "Failed to add elements " + str(variantsIds) + " for Variants on ServiceBulletin"

		try:
			# get the ServiceBulletin
			serviceBulletin = self.get( serviceBulletinId ).first()
				
			# split on a comma with no spaces
			idList = variantsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the AircraftVariant		
				aircraftVariant = AircraftVariantDelegate().get(id).first();	
				# add the AircraftVariant
				serviceBulletin.variants.add(aircraftVariant)
				
			# save it		
			serviceBulletin.save()
			
			# reload and return the appropriate version
			return self.get( serviceBulletinId );
		except ServiceBulletin.DoesNotExist:
			raise ProcessingError(errMsg + " : ServiceBulletin with id " + str(serviceBulletinId) + " does not exist.")
		except AircraftVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftVariant does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeVariants( self, serviceBulletinId, variantsIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.AircraftVariantDelegate import AircraftVariantDelegate

		errMsg = "Failed to remove elements " + str(variantsIds) + " for Variants on ServiceBulletin"

		try:
			# get the ServiceBulletin
			serviceBulletin = self.get( serviceBulletinId ).first()
				
			# split on a comma with no spaces
			idList = variantsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the AircraftVariant		
				aircraftVariant = AircraftVariantDelegate().get(id).first();	
				# add the AircraftVariant
				serviceBulletin.variants.remove(aircraftVariant)
				
			# save it		
			serviceBulletin.save()
			
			# reload and return the appropriate version
			return self.get( serviceBulletinId );
		except ServiceBulletin.DoesNotExist:
			raise ProcessingError(errMsg + " : ServiceBulletin with id " + str(serviceBulletinId) + " does not exist.")
		except AircraftVariant.DoesNotExist:
			raise ProcessingError(errMsg + " : AircraftVariant does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
