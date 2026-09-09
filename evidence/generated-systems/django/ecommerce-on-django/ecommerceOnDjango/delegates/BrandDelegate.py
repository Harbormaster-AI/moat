from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from ecommerceOnDjango.models.Brand import Brand
from ecommerceOnDjango.models.Merchant import Merchant
from ecommerceOnDjango.models.Product import Product
from ecommerceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Brand
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BrandDelegate Declaration
#======================================================================
class BrandDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, brandId ):
		try:	
			brand = Brand.objects.filter(id=brandId)
			return brand.first();
		except Brand.DoesNotExist:
			raise ProcessingError("Brand with id " + str(brandId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, brand):
		for model in serializers.deserialize("json", brand):
			model.save()
			return model;

	def create(self, brand):
		brand.save()
		return brand;

	def saveFromJson(self, brand):
		for model in serializers.deserialize("json", brand):
			model.save()
			return brand;
	
	def save(self, brand):
		brand.save()
		return brand;
	
	def delete(self, brandId ):
		errMsg = "Failed to delete Brand from db using id " + str(brandId)
		
		try:
			brand = Brand.objects.get(id=brandId)
			brand.delete()
			return True
		except Brand.DoesNotExist:
			raise ProcessingError("Brand with id " + str(brandId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Brand.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Brand from db")
		except Exception:
			return None;
		
	def assignMerchant( self, brandId, merchantId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.MerchantDelegate import MerchantDelegate

		errMsg = "Failed to assign element " + str(merchantId) + " for Merchant on Brand"

		try:
			# get the Brand from db
			brand = self.get( brandId ).first()	
			
			# get the Merchant from db
			merchant = MerchantDelegate().get(merchantId).first();
			
			# assign the Merchant		
			brand.merchant = merchant
			
			#save it
			brand.save()

			# reload and return the appropriate version					
			return self.get( brandId );
		except Brand.DoesNotExist:
			raise ProcessingError(errMsg + " : Brand with id " + str(brandId) + " does not exist.")
		except Merchant.DoesNotExist:
			raise ProcessingError(errMsg + " : Merchant with id " + str(merchantId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignMerchant( self, brandId ):
		errMsg = "Failed to unassign element " + str(merchantId) + " for Merchant on Brand"

		try:
			# get the Brand from db
			brand = self.get( brandId ).first()	
			
			# assign to None for unassignment
			brand.merchant = None			

			#save it
			brand.save()

			# reload and return the appropriate version					
			return self.get( brandId );
		except Brand.DoesNotExist:
			raise ProcessingError(errMsg + " : Brand with id " + str(brandId) + " does not exist.")
		except Exception:
			return None;
		
	def addProducts( self, brandId, productsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ProductDelegate import ProductDelegate

		errMsg = "Failed to add elements " + str(productsIds) + " for Products on Brand"

		try:
			# get the Brand
			brand = self.get( brandId ).first()
				
			# split on a comma with no spaces
			idList = productsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Product		
				product = ProductDelegate().get(id).first();	
				# add the Product
				brand.products.add(product)
				
			# save it		
			brand.save()
			
			# reload and return the appropriate version
			return self.get( brandId );
		except Brand.DoesNotExist:
			raise ProcessingError(errMsg + " : Brand with id " + str(brandId) + " does not exist.")
		except Product.DoesNotExist:
			raise ProcessingError(errMsg + " : Product does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeProducts( self, brandId, productsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ProductDelegate import ProductDelegate

		errMsg = "Failed to remove elements " + str(productsIds) + " for Products on Brand"

		try:
			# get the Brand
			brand = self.get( brandId ).first()
				
			# split on a comma with no spaces
			idList = productsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Product		
				product = ProductDelegate().get(id).first();	
				# add the Product
				brand.products.remove(product)
				
			# save it		
			brand.save()
			
			# reload and return the appropriate version
			return self.get( brandId );
		except Brand.DoesNotExist:
			raise ProcessingError(errMsg + " : Brand with id " + str(brandId) + " does not exist.")
		except Product.DoesNotExist:
			raise ProcessingError(errMsg + " : Product does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
