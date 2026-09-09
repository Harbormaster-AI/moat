from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from healthcareOnDjango.models.InventoryItem import InventoryItem
from healthcareOnDjango.models.Facility import Facility
from healthcareOnDjango.models.MedicalSupplier import MedicalSupplier
from healthcareOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model InventoryItem
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InventoryItemDelegate Declaration
#======================================================================
class InventoryItemDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, inventoryItemId ):
		try:	
			inventoryItem = InventoryItem.objects.filter(id=inventoryItemId)
			return inventoryItem.first();
		except InventoryItem.DoesNotExist:
			raise ProcessingError("InventoryItem with id " + str(inventoryItemId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, inventoryItem):
		for model in serializers.deserialize("json", inventoryItem):
			model.save()
			return model;

	def create(self, inventoryItem):
		inventoryItem.save()
		return inventoryItem;

	def saveFromJson(self, inventoryItem):
		for model in serializers.deserialize("json", inventoryItem):
			model.save()
			return inventoryItem;
	
	def save(self, inventoryItem):
		inventoryItem.save()
		return inventoryItem;
	
	def delete(self, inventoryItemId ):
		errMsg = "Failed to delete InventoryItem from db using id " + str(inventoryItemId)
		
		try:
			inventoryItem = InventoryItem.objects.get(id=inventoryItemId)
			inventoryItem.delete()
			return True
		except InventoryItem.DoesNotExist:
			raise ProcessingError("InventoryItem with id " + str(inventoryItemId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = InventoryItem.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all InventoryItem from db")
		except Exception:
			return None;
		
	def assignFacility( self, inventoryItemId, facilityId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.FacilityDelegate import FacilityDelegate

		errMsg = "Failed to assign element " + str(facilityId) + " for Facility on InventoryItem"

		try:
			# get the InventoryItem from db
			inventoryItem = self.get( inventoryItemId ).first()	
			
			# get the Facility from db
			facility = FacilityDelegate().get(facilityId).first();
			
			# assign the Facility		
			inventoryItem.facility = facility
			
			#save it
			inventoryItem.save()

			# reload and return the appropriate version					
			return self.get( inventoryItemId );
		except InventoryItem.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryItem with id " + str(inventoryItemId) + " does not exist.")
		except Facility.DoesNotExist:
			raise ProcessingError(errMsg + " : Facility with id " + str(facilityId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignFacility( self, inventoryItemId ):
		errMsg = "Failed to unassign element " + str(facilityId) + " for Facility on InventoryItem"

		try:
			# get the InventoryItem from db
			inventoryItem = self.get( inventoryItemId ).first()	
			
			# assign to None for unassignment
			inventoryItem.facility = None			

			#save it
			inventoryItem.save()

			# reload and return the appropriate version					
			return self.get( inventoryItemId );
		except InventoryItem.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryItem with id " + str(inventoryItemId) + " does not exist.")
		except Exception:
			return None;
		
	def assignSupplier( self, inventoryItemId, supplierId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.MedicalSupplierDelegate import MedicalSupplierDelegate

		errMsg = "Failed to assign element " + str(supplierId) + " for Supplier on InventoryItem"

		try:
			# get the InventoryItem from db
			inventoryItem = self.get( inventoryItemId ).first()	
			
			# get the MedicalSupplier from db
			medicalSupplier = MedicalSupplierDelegate().get(supplierId).first();
			
			# assign the Supplier		
			inventoryItem.supplier = medicalSupplier
			
			#save it
			inventoryItem.save()

			# reload and return the appropriate version					
			return self.get( inventoryItemId );
		except InventoryItem.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryItem with id " + str(inventoryItemId) + " does not exist.")
		except MedicalSupplier.DoesNotExist:
			raise ProcessingError(errMsg + " : MedicalSupplier with id " + str(supplierId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignSupplier( self, inventoryItemId ):
		errMsg = "Failed to unassign element " + str(supplierId) + " for Supplier on InventoryItem"

		try:
			# get the InventoryItem from db
			inventoryItem = self.get( inventoryItemId ).first()	
			
			# assign to None for unassignment
			inventoryItem.medicalSupplier = None			

			#save it
			inventoryItem.save()

			# reload and return the appropriate version					
			return self.get( inventoryItemId );
		except InventoryItem.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryItem with id " + str(inventoryItemId) + " does not exist.")
		except Exception:
			return None;
		
