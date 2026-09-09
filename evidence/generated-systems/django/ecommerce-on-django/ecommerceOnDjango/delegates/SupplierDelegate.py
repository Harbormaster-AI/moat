from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from ecommerceOnDjango.models.Supplier import Supplier
from ecommerceOnDjango.models.Merchant import Merchant
from ecommerceOnDjango.models.Product import Product
from ecommerceOnDjango.models.FulfillmentCenter import FulfillmentCenter
from ecommerceOnDjango.exceptions import Exceptions

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
		
	def assignMerchant( self, supplierId, merchantId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.MerchantDelegate import MerchantDelegate

		errMsg = "Failed to assign element " + str(merchantId) + " for Merchant on Supplier"

		try:
			# get the Supplier from db
			supplier = self.get( supplierId ).first()	
			
			# get the Merchant from db
			merchant = MerchantDelegate().get(merchantId).first();
			
			# assign the Merchant		
			supplier.merchant = merchant
			
			#save it
			supplier.save()

			# reload and return the appropriate version					
			return self.get( supplierId );
		except Supplier.DoesNotExist:
			raise ProcessingError(errMsg + " : Supplier with id " + str(supplierId) + " does not exist.")
		except Merchant.DoesNotExist:
			raise ProcessingError(errMsg + " : Merchant with id " + str(merchantId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignMerchant( self, supplierId ):
		errMsg = "Failed to unassign element " + str(merchantId) + " for Merchant on Supplier"

		try:
			# get the Supplier from db
			supplier = self.get( supplierId ).first()	
			
			# assign to None for unassignment
			supplier.merchant = None			

			#save it
			supplier.save()

			# reload and return the appropriate version					
			return self.get( supplierId );
		except Supplier.DoesNotExist:
			raise ProcessingError(errMsg + " : Supplier with id " + str(supplierId) + " does not exist.")
		except Exception:
			return None;
		
	def addProducts( self, supplierId, productsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ProductDelegate import ProductDelegate

		errMsg = "Failed to add elements " + str(productsIds) + " for Products on Supplier"

		try:
			# get the Supplier
			supplier = self.get( supplierId ).first()
				
			# split on a comma with no spaces
			idList = productsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Product		
				product = ProductDelegate().get(id).first();	
				# add the Product
				supplier.products.add(product)
				
			# save it		
			supplier.save()
			
			# reload and return the appropriate version
			return self.get( supplierId );
		except Supplier.DoesNotExist:
			raise ProcessingError(errMsg + " : Supplier with id " + str(supplierId) + " does not exist.")
		except Product.DoesNotExist:
			raise ProcessingError(errMsg + " : Product does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeProducts( self, supplierId, productsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ProductDelegate import ProductDelegate

		errMsg = "Failed to remove elements " + str(productsIds) + " for Products on Supplier"

		try:
			# get the Supplier
			supplier = self.get( supplierId ).first()
				
			# split on a comma with no spaces
			idList = productsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Product		
				product = ProductDelegate().get(id).first();	
				# add the Product
				supplier.products.remove(product)
				
			# save it		
			supplier.save()
			
			# reload and return the appropriate version
			return self.get( supplierId );
		except Supplier.DoesNotExist:
			raise ProcessingError(errMsg + " : Supplier with id " + str(supplierId) + " does not exist.")
		except Product.DoesNotExist:
			raise ProcessingError(errMsg + " : Product does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addFulfillmentCenters( self, supplierId, fulfillmentCentersIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.FulfillmentCenterDelegate import FulfillmentCenterDelegate

		errMsg = "Failed to add elements " + str(fulfillmentCentersIds) + " for FulfillmentCenters on Supplier"

		try:
			# get the Supplier
			supplier = self.get( supplierId ).first()
				
			# split on a comma with no spaces
			idList = fulfillmentCentersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the FulfillmentCenter		
				fulfillmentCenter = FulfillmentCenterDelegate().get(id).first();	
				# add the FulfillmentCenter
				supplier.fulfillmentCenters.add(fulfillmentCenter)
				
			# save it		
			supplier.save()
			
			# reload and return the appropriate version
			return self.get( supplierId );
		except Supplier.DoesNotExist:
			raise ProcessingError(errMsg + " : Supplier with id " + str(supplierId) + " does not exist.")
		except FulfillmentCenter.DoesNotExist:
			raise ProcessingError(errMsg + " : FulfillmentCenter does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeFulfillmentCenters( self, supplierId, fulfillmentCentersIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.FulfillmentCenterDelegate import FulfillmentCenterDelegate

		errMsg = "Failed to remove elements " + str(fulfillmentCentersIds) + " for FulfillmentCenters on Supplier"

		try:
			# get the Supplier
			supplier = self.get( supplierId ).first()
				
			# split on a comma with no spaces
			idList = fulfillmentCentersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the FulfillmentCenter		
				fulfillmentCenter = FulfillmentCenterDelegate().get(id).first();	
				# add the FulfillmentCenter
				supplier.fulfillmentCenters.remove(fulfillmentCenter)
				
			# save it		
			supplier.save()
			
			# reload and return the appropriate version
			return self.get( supplierId );
		except Supplier.DoesNotExist:
			raise ProcessingError(errMsg + " : Supplier with id " + str(supplierId) + " does not exist.")
		except FulfillmentCenter.DoesNotExist:
			raise ProcessingError(errMsg + " : FulfillmentCenter does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
