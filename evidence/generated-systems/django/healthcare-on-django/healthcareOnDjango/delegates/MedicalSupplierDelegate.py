from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from healthcareOnDjango.models.MedicalSupplier import MedicalSupplier
from healthcareOnDjango.models.Facility import Facility
from healthcareOnDjango.models.InventoryItem import InventoryItem
from healthcareOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model MedicalSupplier
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MedicalSupplierDelegate Declaration
#======================================================================
class MedicalSupplierDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, medicalSupplierId ):
		try:	
			medicalSupplier = MedicalSupplier.objects.filter(id=medicalSupplierId)
			return medicalSupplier.first();
		except MedicalSupplier.DoesNotExist:
			raise ProcessingError("MedicalSupplier with id " + str(medicalSupplierId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, medicalSupplier):
		for model in serializers.deserialize("json", medicalSupplier):
			model.save()
			return model;

	def create(self, medicalSupplier):
		medicalSupplier.save()
		return medicalSupplier;

	def saveFromJson(self, medicalSupplier):
		for model in serializers.deserialize("json", medicalSupplier):
			model.save()
			return medicalSupplier;
	
	def save(self, medicalSupplier):
		medicalSupplier.save()
		return medicalSupplier;
	
	def delete(self, medicalSupplierId ):
		errMsg = "Failed to delete MedicalSupplier from db using id " + str(medicalSupplierId)
		
		try:
			medicalSupplier = MedicalSupplier.objects.get(id=medicalSupplierId)
			medicalSupplier.delete()
			return True
		except MedicalSupplier.DoesNotExist:
			raise ProcessingError("MedicalSupplier with id " + str(medicalSupplierId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = MedicalSupplier.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all MedicalSupplier from db")
		except Exception:
			return None;
		
	def addFacilities( self, medicalSupplierId, facilitiesIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.FacilityDelegate import FacilityDelegate

		errMsg = "Failed to add elements " + str(facilitiesIds) + " for Facilities on MedicalSupplier"

		try:
			# get the MedicalSupplier
			medicalSupplier = self.get( medicalSupplierId ).first()
				
			# split on a comma with no spaces
			idList = facilitiesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Facility		
				facility = FacilityDelegate().get(id).first();	
				# add the Facility
				medicalSupplier.facilities.add(facility)
				
			# save it		
			medicalSupplier.save()
			
			# reload and return the appropriate version
			return self.get( medicalSupplierId );
		except MedicalSupplier.DoesNotExist:
			raise ProcessingError(errMsg + " : MedicalSupplier with id " + str(medicalSupplierId) + " does not exist.")
		except Facility.DoesNotExist:
			raise ProcessingError(errMsg + " : Facility does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeFacilities( self, medicalSupplierId, facilitiesIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.FacilityDelegate import FacilityDelegate

		errMsg = "Failed to remove elements " + str(facilitiesIds) + " for Facilities on MedicalSupplier"

		try:
			# get the MedicalSupplier
			medicalSupplier = self.get( medicalSupplierId ).first()
				
			# split on a comma with no spaces
			idList = facilitiesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Facility		
				facility = FacilityDelegate().get(id).first();	
				# add the Facility
				medicalSupplier.facilities.remove(facility)
				
			# save it		
			medicalSupplier.save()
			
			# reload and return the appropriate version
			return self.get( medicalSupplierId );
		except MedicalSupplier.DoesNotExist:
			raise ProcessingError(errMsg + " : MedicalSupplier with id " + str(medicalSupplierId) + " does not exist.")
		except Facility.DoesNotExist:
			raise ProcessingError(errMsg + " : Facility does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addInventoryItems( self, medicalSupplierId, inventoryItemsIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.InventoryItemDelegate import InventoryItemDelegate

		errMsg = "Failed to add elements " + str(inventoryItemsIds) + " for InventoryItems on MedicalSupplier"

		try:
			# get the MedicalSupplier
			medicalSupplier = self.get( medicalSupplierId ).first()
				
			# split on a comma with no spaces
			idList = inventoryItemsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the InventoryItem		
				inventoryItem = InventoryItemDelegate().get(id).first();	
				# add the InventoryItem
				medicalSupplier.inventoryItems.add(inventoryItem)
				
			# save it		
			medicalSupplier.save()
			
			# reload and return the appropriate version
			return self.get( medicalSupplierId );
		except MedicalSupplier.DoesNotExist:
			raise ProcessingError(errMsg + " : MedicalSupplier with id " + str(medicalSupplierId) + " does not exist.")
		except InventoryItem.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryItem does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeInventoryItems( self, medicalSupplierId, inventoryItemsIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.InventoryItemDelegate import InventoryItemDelegate

		errMsg = "Failed to remove elements " + str(inventoryItemsIds) + " for InventoryItems on MedicalSupplier"

		try:
			# get the MedicalSupplier
			medicalSupplier = self.get( medicalSupplierId ).first()
				
			# split on a comma with no spaces
			idList = inventoryItemsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the InventoryItem		
				inventoryItem = InventoryItemDelegate().get(id).first();	
				# add the InventoryItem
				medicalSupplier.inventoryItems.remove(inventoryItem)
				
			# save it		
			medicalSupplier.save()
			
			# reload and return the appropriate version
			return self.get( medicalSupplierId );
		except MedicalSupplier.DoesNotExist:
			raise ProcessingError(errMsg + " : MedicalSupplier with id " + str(medicalSupplierId) + " does not exist.")
		except InventoryItem.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryItem does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
