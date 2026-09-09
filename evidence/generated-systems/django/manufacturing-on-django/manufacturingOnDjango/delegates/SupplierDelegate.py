from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from manufacturingOnDjango.models.Supplier import Supplier
from manufacturingOnDjango.models.Enterprise import Enterprise
from manufacturingOnDjango.models.Item import Item
from manufacturingOnDjango.models.PurchaseOrder import PurchaseOrder
from manufacturingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Supplier
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SupplierDelegate Declaration
#======================================================================
class SupplierDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, supplierId ):
		try:	
			supplier = Supplier.objects.filter(id=supplierId)
			return supplier.first();
		except Supplier.DoesNotExist:
			raise ProcessingError("Supplier with id " + str(supplierId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, supplier):
		for model in serializers.deserialize("json", supplier):
			model.save()
			return model;

	def create(self, supplier):
		supplier.save()
		return supplier;

	def saveFromJson(self, supplier):
		for model in serializers.deserialize("json", supplier):
			model.save()
			return supplier;
	
	def save(self, supplier):
		supplier.save()
		return supplier;
	
	def delete(self, supplierId ):
		errMsg = "Failed to delete Supplier from db using id " + str(supplierId)
		
		try:
			supplier = Supplier.objects.get(id=supplierId)
			supplier.delete()
			return True
		except Supplier.DoesNotExist:
			raise ProcessingError("Supplier with id " + str(supplierId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Supplier.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Supplier from db")
		except Exception:
			return None;
		
	def addEnterprises( self, supplierId, enterprisesIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.EnterpriseDelegate import EnterpriseDelegate

		errMsg = "Failed to add elements " + str(enterprisesIds) + " for Enterprises on Supplier"

		try:
			# get the Supplier
			supplier = self.get( supplierId ).first()
				
			# split on a comma with no spaces
			idList = enterprisesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Enterprise		
				enterprise = EnterpriseDelegate().get(id).first();	
				# add the Enterprise
				supplier.enterprises.add(enterprise)
				
			# save it		
			supplier.save()
			
			# reload and return the appropriate version
			return self.get( supplierId );
		except Supplier.DoesNotExist:
			raise ProcessingError(errMsg + " : Supplier with id " + str(supplierId) + " does not exist.")
		except Enterprise.DoesNotExist:
			raise ProcessingError(errMsg + " : Enterprise does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeEnterprises( self, supplierId, enterprisesIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.EnterpriseDelegate import EnterpriseDelegate

		errMsg = "Failed to remove elements " + str(enterprisesIds) + " for Enterprises on Supplier"

		try:
			# get the Supplier
			supplier = self.get( supplierId ).first()
				
			# split on a comma with no spaces
			idList = enterprisesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Enterprise		
				enterprise = EnterpriseDelegate().get(id).first();	
				# add the Enterprise
				supplier.enterprises.remove(enterprise)
				
			# save it		
			supplier.save()
			
			# reload and return the appropriate version
			return self.get( supplierId );
		except Supplier.DoesNotExist:
			raise ProcessingError(errMsg + " : Supplier with id " + str(supplierId) + " does not exist.")
		except Enterprise.DoesNotExist:
			raise ProcessingError(errMsg + " : Enterprise does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addItems( self, supplierId, itemsIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.ItemDelegate import ItemDelegate

		errMsg = "Failed to add elements " + str(itemsIds) + " for Items on Supplier"

		try:
			# get the Supplier
			supplier = self.get( supplierId ).first()
				
			# split on a comma with no spaces
			idList = itemsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Item		
				item = ItemDelegate().get(id).first();	
				# add the Item
				supplier.items.add(item)
				
			# save it		
			supplier.save()
			
			# reload and return the appropriate version
			return self.get( supplierId );
		except Supplier.DoesNotExist:
			raise ProcessingError(errMsg + " : Supplier with id " + str(supplierId) + " does not exist.")
		except Item.DoesNotExist:
			raise ProcessingError(errMsg + " : Item does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeItems( self, supplierId, itemsIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.ItemDelegate import ItemDelegate

		errMsg = "Failed to remove elements " + str(itemsIds) + " for Items on Supplier"

		try:
			# get the Supplier
			supplier = self.get( supplierId ).first()
				
			# split on a comma with no spaces
			idList = itemsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Item		
				item = ItemDelegate().get(id).first();	
				# add the Item
				supplier.items.remove(item)
				
			# save it		
			supplier.save()
			
			# reload and return the appropriate version
			return self.get( supplierId );
		except Supplier.DoesNotExist:
			raise ProcessingError(errMsg + " : Supplier with id " + str(supplierId) + " does not exist.")
		except Item.DoesNotExist:
			raise ProcessingError(errMsg + " : Item does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addPurchaseOrders( self, supplierId, purchaseOrdersIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.PurchaseOrderDelegate import PurchaseOrderDelegate

		errMsg = "Failed to add elements " + str(purchaseOrdersIds) + " for PurchaseOrders on Supplier"

		try:
			# get the Supplier
			supplier = self.get( supplierId ).first()
				
			# split on a comma with no spaces
			idList = purchaseOrdersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the PurchaseOrder		
				purchaseOrder = PurchaseOrderDelegate().get(id).first();	
				# add the PurchaseOrder
				supplier.purchaseOrders.add(purchaseOrder)
				
			# save it		
			supplier.save()
			
			# reload and return the appropriate version
			return self.get( supplierId );
		except Supplier.DoesNotExist:
			raise ProcessingError(errMsg + " : Supplier with id " + str(supplierId) + " does not exist.")
		except PurchaseOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : PurchaseOrder does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePurchaseOrders( self, supplierId, purchaseOrdersIds ):
		# lazy importing avoids circular dependencies
		from manufacturingOnDjango.delegates.PurchaseOrderDelegate import PurchaseOrderDelegate

		errMsg = "Failed to remove elements " + str(purchaseOrdersIds) + " for PurchaseOrders on Supplier"

		try:
			# get the Supplier
			supplier = self.get( supplierId ).first()
				
			# split on a comma with no spaces
			idList = purchaseOrdersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the PurchaseOrder		
				purchaseOrder = PurchaseOrderDelegate().get(id).first();	
				# add the PurchaseOrder
				supplier.purchaseOrders.remove(purchaseOrder)
				
			# save it		
			supplier.save()
			
			# reload and return the appropriate version
			return self.get( supplierId );
		except Supplier.DoesNotExist:
			raise ProcessingError(errMsg + " : Supplier with id " + str(supplierId) + " does not exist.")
		except PurchaseOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : PurchaseOrder does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
