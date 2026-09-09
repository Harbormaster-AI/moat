from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from manufacturingOnDjango.models.Item import Item
from manufacturingOnDjango.models.BusinessUnit import BusinessUnit
from manufacturingOnDjango.models.BOM import BOM
from manufacturingOnDjango.models.Routing import Routing
from manufacturingOnDjango.models.Supplier import Supplier
from manufacturingOnDjango.models.QualitySpecification import QualitySpecification
from manufacturingOnDjango.models.InventoryItem import InventoryItem
from manufacturingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Item
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ItemDelegate Declaration
#======================================================================
class ItemDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, itemId ):
		try:	
			item = Item.objects.filter(id=itemId)
			return item.first();
		except Item.DoesNotExist:
			raise ProcessingError("Item with id " + str(itemId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, item):
		for model in serializers.deserialize("json", item):
			model.save()
			return model;

	def create(self, item):
		item.save()
		return item;

	def saveFromJson(self, item):
		for model in serializers.deserialize("json", item):
			model.save()
			return item;
	
	def save(self, item):
		item.save()
		return item;
	
	def delete(self, itemId ):
		errMsg = "Failed to delete Item from db using id " + str(itemId)
		
		try:
			item = Item.objects.get(id=itemId)
			item.delete()
			return True
		except Item.DoesNotExist:
			raise ProcessingError("Item with id " + str(itemId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Item.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Item from db")
		except Exception:
			return None;
		
	def assignBusinessUnit( self, itemId, businessUnitId ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.BusinessUnitDelegate import BusinessUnitDelegate

		errMsg = "Failed to assign element " + str(businessUnitId) + " for BusinessUnit on Item"

		try:
			# get the Item from db
			item = self.get( itemId ).first()	
			
			# get the BusinessUnit from db
			businessUnit = BusinessUnitDelegate().get(businessUnitId).first();
			
			# assign the BusinessUnit		
			item.businessUnit = businessUnit
			
			#save it
			item.save()

			# reload and return the appropriate version					
			return self.get( itemId );
		except Item.DoesNotExist:
			raise ProcessingError(errMsg + " : Item with id " + str(itemId) + " does not exist.")
		except BusinessUnit.DoesNotExist:
			raise ProcessingError(errMsg + " : BusinessUnit with id " + str(businessUnitId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignBusinessUnit( self, itemId ):
		errMsg = "Failed to unassign element " + str(businessUnitId) + " for BusinessUnit on Item"

		try:
			# get the Item from db
			item = self.get( itemId ).first()	
			
			# assign to None for unassignment
			item.businessUnit = None			

			#save it
			item.save()

			# reload and return the appropriate version					
			return self.get( itemId );
		except Item.DoesNotExist:
			raise ProcessingError(errMsg + " : Item with id " + str(itemId) + " does not exist.")
		except Exception:
			return None;
		
	def addBoms( self, itemId, bomsIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.BOMDelegate import BOMDelegate

		errMsg = "Failed to add elements " + str(bomsIds) + " for Boms on Item"

		try:
			# get the Item
			item = self.get( itemId ).first()
				
			# split on a comma with no spaces
			idList = bomsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the BOM		
				bOM = BOMDelegate().get(id).first();	
				# add the BOM
				item.boms.add(bOM)
				
			# save it		
			item.save()
			
			# reload and return the appropriate version
			return self.get( itemId );
		except Item.DoesNotExist:
			raise ProcessingError(errMsg + " : Item with id " + str(itemId) + " does not exist.")
		except BOM.DoesNotExist:
			raise ProcessingError(errMsg + " : BOM does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeBoms( self, itemId, bomsIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.BOMDelegate import BOMDelegate

		errMsg = "Failed to remove elements " + str(bomsIds) + " for Boms on Item"

		try:
			# get the Item
			item = self.get( itemId ).first()
				
			# split on a comma with no spaces
			idList = bomsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the BOM		
				bOM = BOMDelegate().get(id).first();	
				# add the BOM
				item.boms.remove(bOM)
				
			# save it		
			item.save()
			
			# reload and return the appropriate version
			return self.get( itemId );
		except Item.DoesNotExist:
			raise ProcessingError(errMsg + " : Item with id " + str(itemId) + " does not exist.")
		except BOM.DoesNotExist:
			raise ProcessingError(errMsg + " : BOM does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addRoutings( self, itemId, routingsIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.RoutingDelegate import RoutingDelegate

		errMsg = "Failed to add elements " + str(routingsIds) + " for Routings on Item"

		try:
			# get the Item
			item = self.get( itemId ).first()
				
			# split on a comma with no spaces
			idList = routingsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Routing		
				routing = RoutingDelegate().get(id).first();	
				# add the Routing
				item.routings.add(routing)
				
			# save it		
			item.save()
			
			# reload and return the appropriate version
			return self.get( itemId );
		except Item.DoesNotExist:
			raise ProcessingError(errMsg + " : Item with id " + str(itemId) + " does not exist.")
		except Routing.DoesNotExist:
			raise ProcessingError(errMsg + " : Routing does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeRoutings( self, itemId, routingsIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.RoutingDelegate import RoutingDelegate

		errMsg = "Failed to remove elements " + str(routingsIds) + " for Routings on Item"

		try:
			# get the Item
			item = self.get( itemId ).first()
				
			# split on a comma with no spaces
			idList = routingsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Routing		
				routing = RoutingDelegate().get(id).first();	
				# add the Routing
				item.routings.remove(routing)
				
			# save it		
			item.save()
			
			# reload and return the appropriate version
			return self.get( itemId );
		except Item.DoesNotExist:
			raise ProcessingError(errMsg + " : Item with id " + str(itemId) + " does not exist.")
		except Routing.DoesNotExist:
			raise ProcessingError(errMsg + " : Routing does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addSuppliers( self, itemId, suppliersIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.SupplierDelegate import SupplierDelegate

		errMsg = "Failed to add elements " + str(suppliersIds) + " for Suppliers on Item"

		try:
			# get the Item
			item = self.get( itemId ).first()
				
			# split on a comma with no spaces
			idList = suppliersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Supplier		
				supplier = SupplierDelegate().get(id).first();	
				# add the Supplier
				item.suppliers.add(supplier)
				
			# save it		
			item.save()
			
			# reload and return the appropriate version
			return self.get( itemId );
		except Item.DoesNotExist:
			raise ProcessingError(errMsg + " : Item with id " + str(itemId) + " does not exist.")
		except Supplier.DoesNotExist:
			raise ProcessingError(errMsg + " : Supplier does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeSuppliers( self, itemId, suppliersIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.SupplierDelegate import SupplierDelegate

		errMsg = "Failed to remove elements " + str(suppliersIds) + " for Suppliers on Item"

		try:
			# get the Item
			item = self.get( itemId ).first()
				
			# split on a comma with no spaces
			idList = suppliersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Supplier		
				supplier = SupplierDelegate().get(id).first();	
				# add the Supplier
				item.suppliers.remove(supplier)
				
			# save it		
			item.save()
			
			# reload and return the appropriate version
			return self.get( itemId );
		except Item.DoesNotExist:
			raise ProcessingError(errMsg + " : Item with id " + str(itemId) + " does not exist.")
		except Supplier.DoesNotExist:
			raise ProcessingError(errMsg + " : Supplier does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addQualitySpecifications( self, itemId, qualitySpecificationsIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.QualitySpecificationDelegate import QualitySpecificationDelegate

		errMsg = "Failed to add elements " + str(qualitySpecificationsIds) + " for QualitySpecifications on Item"

		try:
			# get the Item
			item = self.get( itemId ).first()
				
			# split on a comma with no spaces
			idList = qualitySpecificationsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the QualitySpecification		
				qualitySpecification = QualitySpecificationDelegate().get(id).first();	
				# add the QualitySpecification
				item.qualitySpecifications.add(qualitySpecification)
				
			# save it		
			item.save()
			
			# reload and return the appropriate version
			return self.get( itemId );
		except Item.DoesNotExist:
			raise ProcessingError(errMsg + " : Item with id " + str(itemId) + " does not exist.")
		except QualitySpecification.DoesNotExist:
			raise ProcessingError(errMsg + " : QualitySpecification does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeQualitySpecifications( self, itemId, qualitySpecificationsIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.QualitySpecificationDelegate import QualitySpecificationDelegate

		errMsg = "Failed to remove elements " + str(qualitySpecificationsIds) + " for QualitySpecifications on Item"

		try:
			# get the Item
			item = self.get( itemId ).first()
				
			# split on a comma with no spaces
			idList = qualitySpecificationsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the QualitySpecification		
				qualitySpecification = QualitySpecificationDelegate().get(id).first();	
				# add the QualitySpecification
				item.qualitySpecifications.remove(qualitySpecification)
				
			# save it		
			item.save()
			
			# reload and return the appropriate version
			return self.get( itemId );
		except Item.DoesNotExist:
			raise ProcessingError(errMsg + " : Item with id " + str(itemId) + " does not exist.")
		except QualitySpecification.DoesNotExist:
			raise ProcessingError(errMsg + " : QualitySpecification does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addInventoryItems( self, itemId, inventoryItemsIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.InventoryItemDelegate import InventoryItemDelegate

		errMsg = "Failed to add elements " + str(inventoryItemsIds) + " for InventoryItems on Item"

		try:
			# get the Item
			item = self.get( itemId ).first()
				
			# split on a comma with no spaces
			idList = inventoryItemsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the InventoryItem		
				inventoryItem = InventoryItemDelegate().get(id).first();	
				# add the InventoryItem
				item.inventoryItems.add(inventoryItem)
				
			# save it		
			item.save()
			
			# reload and return the appropriate version
			return self.get( itemId );
		except Item.DoesNotExist:
			raise ProcessingError(errMsg + " : Item with id " + str(itemId) + " does not exist.")
		except InventoryItem.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryItem does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeInventoryItems( self, itemId, inventoryItemsIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.InventoryItemDelegate import InventoryItemDelegate

		errMsg = "Failed to remove elements " + str(inventoryItemsIds) + " for InventoryItems on Item"

		try:
			# get the Item
			item = self.get( itemId ).first()
				
			# split on a comma with no spaces
			idList = inventoryItemsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the InventoryItem		
				inventoryItem = InventoryItemDelegate().get(id).first();	
				# add the InventoryItem
				item.inventoryItems.remove(inventoryItem)
				
			# save it		
			item.save()
			
			# reload and return the appropriate version
			return self.get( itemId );
		except Item.DoesNotExist:
			raise ProcessingError(errMsg + " : Item with id " + str(itemId) + " does not exist.")
		except InventoryItem.DoesNotExist:
			raise ProcessingError(errMsg + " : InventoryItem does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
