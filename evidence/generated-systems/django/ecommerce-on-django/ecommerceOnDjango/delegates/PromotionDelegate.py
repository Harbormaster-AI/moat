from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from ecommerceOnDjango.models.Promotion import Promotion
from ecommerceOnDjango.models.Merchant import Merchant
from ecommerceOnDjango.models.Channel import Channel
from ecommerceOnDjango.models.Product import Product
from ecommerceOnDjango.models.Category import Category
from ecommerceOnDjango.models.Coupon import Coupon
from ecommerceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Promotion
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PromotionDelegate Declaration
#======================================================================
class PromotionDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, promotionId ):
		try:	
			promotion = Promotion.objects.filter(id=promotionId)
			return promotion.first();
		except Promotion.DoesNotExist:
			raise ProcessingError("Promotion with id " + str(promotionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, promotion):
		for model in serializers.deserialize("json", promotion):
			model.save()
			return model;

	def create(self, promotion):
		promotion.save()
		return promotion;

	def saveFromJson(self, promotion):
		for model in serializers.deserialize("json", promotion):
			model.save()
			return promotion;
	
	def save(self, promotion):
		promotion.save()
		return promotion;
	
	def delete(self, promotionId ):
		errMsg = "Failed to delete Promotion from db using id " + str(promotionId)
		
		try:
			promotion = Promotion.objects.get(id=promotionId)
			promotion.delete()
			return True
		except Promotion.DoesNotExist:
			raise ProcessingError("Promotion with id " + str(promotionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Promotion.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Promotion from db")
		except Exception:
			return None;
		
	def assignMerchant( self, promotionId, merchantId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.MerchantDelegate import MerchantDelegate

		errMsg = "Failed to assign element " + str(merchantId) + " for Merchant on Promotion"

		try:
			# get the Promotion from db
			promotion = self.get( promotionId ).first()	
			
			# get the Merchant from db
			merchant = MerchantDelegate().get(merchantId).first();
			
			# assign the Merchant		
			promotion.merchant = merchant
			
			#save it
			promotion.save()

			# reload and return the appropriate version					
			return self.get( promotionId );
		except Promotion.DoesNotExist:
			raise ProcessingError(errMsg + " : Promotion with id " + str(promotionId) + " does not exist.")
		except Merchant.DoesNotExist:
			raise ProcessingError(errMsg + " : Merchant with id " + str(merchantId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignMerchant( self, promotionId ):
		errMsg = "Failed to unassign element " + str(merchantId) + " for Merchant on Promotion"

		try:
			# get the Promotion from db
			promotion = self.get( promotionId ).first()	
			
			# assign to None for unassignment
			promotion.merchant = None			

			#save it
			promotion.save()

			# reload and return the appropriate version					
			return self.get( promotionId );
		except Promotion.DoesNotExist:
			raise ProcessingError(errMsg + " : Promotion with id " + str(promotionId) + " does not exist.")
		except Exception:
			return None;
		
	def addChannels( self, promotionId, channelsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ChannelDelegate import ChannelDelegate

		errMsg = "Failed to add elements " + str(channelsIds) + " for Channels on Promotion"

		try:
			# get the Promotion
			promotion = self.get( promotionId ).first()
				
			# split on a comma with no spaces
			idList = channelsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Channel		
				channel = ChannelDelegate().get(id).first();	
				# add the Channel
				promotion.channels.add(channel)
				
			# save it		
			promotion.save()
			
			# reload and return the appropriate version
			return self.get( promotionId );
		except Promotion.DoesNotExist:
			raise ProcessingError(errMsg + " : Promotion with id " + str(promotionId) + " does not exist.")
		except Channel.DoesNotExist:
			raise ProcessingError(errMsg + " : Channel does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeChannels( self, promotionId, channelsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ChannelDelegate import ChannelDelegate

		errMsg = "Failed to remove elements " + str(channelsIds) + " for Channels on Promotion"

		try:
			# get the Promotion
			promotion = self.get( promotionId ).first()
				
			# split on a comma with no spaces
			idList = channelsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Channel		
				channel = ChannelDelegate().get(id).first();	
				# add the Channel
				promotion.channels.remove(channel)
				
			# save it		
			promotion.save()
			
			# reload and return the appropriate version
			return self.get( promotionId );
		except Promotion.DoesNotExist:
			raise ProcessingError(errMsg + " : Promotion with id " + str(promotionId) + " does not exist.")
		except Channel.DoesNotExist:
			raise ProcessingError(errMsg + " : Channel does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addApplicableProducts( self, promotionId, applicableProductsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ProductDelegate import ProductDelegate

		errMsg = "Failed to add elements " + str(applicableProductsIds) + " for ApplicableProducts on Promotion"

		try:
			# get the Promotion
			promotion = self.get( promotionId ).first()
				
			# split on a comma with no spaces
			idList = applicableProductsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Product		
				product = ProductDelegate().get(id).first();	
				# add the Product
				promotion.applicableProducts.add(product)
				
			# save it		
			promotion.save()
			
			# reload and return the appropriate version
			return self.get( promotionId );
		except Promotion.DoesNotExist:
			raise ProcessingError(errMsg + " : Promotion with id " + str(promotionId) + " does not exist.")
		except Product.DoesNotExist:
			raise ProcessingError(errMsg + " : Product does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeApplicableProducts( self, promotionId, applicableProductsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ProductDelegate import ProductDelegate

		errMsg = "Failed to remove elements " + str(applicableProductsIds) + " for ApplicableProducts on Promotion"

		try:
			# get the Promotion
			promotion = self.get( promotionId ).first()
				
			# split on a comma with no spaces
			idList = applicableProductsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Product		
				product = ProductDelegate().get(id).first();	
				# add the Product
				promotion.applicableProducts.remove(product)
				
			# save it		
			promotion.save()
			
			# reload and return the appropriate version
			return self.get( promotionId );
		except Promotion.DoesNotExist:
			raise ProcessingError(errMsg + " : Promotion with id " + str(promotionId) + " does not exist.")
		except Product.DoesNotExist:
			raise ProcessingError(errMsg + " : Product does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addApplicableCategories( self, promotionId, applicableCategoriesIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.CategoryDelegate import CategoryDelegate

		errMsg = "Failed to add elements " + str(applicableCategoriesIds) + " for ApplicableCategories on Promotion"

		try:
			# get the Promotion
			promotion = self.get( promotionId ).first()
				
			# split on a comma with no spaces
			idList = applicableCategoriesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Category		
				category = CategoryDelegate().get(id).first();	
				# add the Category
				promotion.applicableCategories.add(category)
				
			# save it		
			promotion.save()
			
			# reload and return the appropriate version
			return self.get( promotionId );
		except Promotion.DoesNotExist:
			raise ProcessingError(errMsg + " : Promotion with id " + str(promotionId) + " does not exist.")
		except Category.DoesNotExist:
			raise ProcessingError(errMsg + " : Category does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeApplicableCategories( self, promotionId, applicableCategoriesIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.CategoryDelegate import CategoryDelegate

		errMsg = "Failed to remove elements " + str(applicableCategoriesIds) + " for ApplicableCategories on Promotion"

		try:
			# get the Promotion
			promotion = self.get( promotionId ).first()
				
			# split on a comma with no spaces
			idList = applicableCategoriesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Category		
				category = CategoryDelegate().get(id).first();	
				# add the Category
				promotion.applicableCategories.remove(category)
				
			# save it		
			promotion.save()
			
			# reload and return the appropriate version
			return self.get( promotionId );
		except Promotion.DoesNotExist:
			raise ProcessingError(errMsg + " : Promotion with id " + str(promotionId) + " does not exist.")
		except Category.DoesNotExist:
			raise ProcessingError(errMsg + " : Category does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addCoupons( self, promotionId, couponsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.CouponDelegate import CouponDelegate

		errMsg = "Failed to add elements " + str(couponsIds) + " for Coupons on Promotion"

		try:
			# get the Promotion
			promotion = self.get( promotionId ).first()
				
			# split on a comma with no spaces
			idList = couponsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Coupon		
				coupon = CouponDelegate().get(id).first();	
				# add the Coupon
				promotion.coupons.add(coupon)
				
			# save it		
			promotion.save()
			
			# reload and return the appropriate version
			return self.get( promotionId );
		except Promotion.DoesNotExist:
			raise ProcessingError(errMsg + " : Promotion with id " + str(promotionId) + " does not exist.")
		except Coupon.DoesNotExist:
			raise ProcessingError(errMsg + " : Coupon does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCoupons( self, promotionId, couponsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.CouponDelegate import CouponDelegate

		errMsg = "Failed to remove elements " + str(couponsIds) + " for Coupons on Promotion"

		try:
			# get the Promotion
			promotion = self.get( promotionId ).first()
				
			# split on a comma with no spaces
			idList = couponsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Coupon		
				coupon = CouponDelegate().get(id).first();	
				# add the Coupon
				promotion.coupons.remove(coupon)
				
			# save it		
			promotion.save()
			
			# reload and return the appropriate version
			return self.get( promotionId );
		except Promotion.DoesNotExist:
			raise ProcessingError(errMsg + " : Promotion with id " + str(promotionId) + " does not exist.")
		except Coupon.DoesNotExist:
			raise ProcessingError(errMsg + " : Coupon does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
