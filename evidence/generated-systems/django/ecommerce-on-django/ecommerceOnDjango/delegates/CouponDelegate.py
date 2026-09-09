from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from ecommerceOnDjango.models.Coupon import Coupon
from ecommerceOnDjango.models.Promotion import Promotion
from ecommerceOnDjango.models.CouponRedemption import CouponRedemption
from ecommerceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Coupon
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CouponDelegate Declaration
#======================================================================
class CouponDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, couponId ):
		try:	
			coupon = Coupon.objects.filter(id=couponId)
			return coupon.first();
		except Coupon.DoesNotExist:
			raise ProcessingError("Coupon with id " + str(couponId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, coupon):
		for model in serializers.deserialize("json", coupon):
			model.save()
			return model;

	def create(self, coupon):
		coupon.save()
		return coupon;

	def saveFromJson(self, coupon):
		for model in serializers.deserialize("json", coupon):
			model.save()
			return coupon;
	
	def save(self, coupon):
		coupon.save()
		return coupon;
	
	def delete(self, couponId ):
		errMsg = "Failed to delete Coupon from db using id " + str(couponId)
		
		try:
			coupon = Coupon.objects.get(id=couponId)
			coupon.delete()
			return True
		except Coupon.DoesNotExist:
			raise ProcessingError("Coupon with id " + str(couponId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Coupon.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Coupon from db")
		except Exception:
			return None;
		
	def assignPromotion( self, couponId, promotionId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.PromotionDelegate import PromotionDelegate

		errMsg = "Failed to assign element " + str(promotionId) + " for Promotion on Coupon"

		try:
			# get the Coupon from db
			coupon = self.get( couponId ).first()	
			
			# get the Promotion from db
			promotion = PromotionDelegate().get(promotionId).first();
			
			# assign the Promotion		
			coupon.promotion = promotion
			
			#save it
			coupon.save()

			# reload and return the appropriate version					
			return self.get( couponId );
		except Coupon.DoesNotExist:
			raise ProcessingError(errMsg + " : Coupon with id " + str(couponId) + " does not exist.")
		except Promotion.DoesNotExist:
			raise ProcessingError(errMsg + " : Promotion with id " + str(promotionId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPromotion( self, couponId ):
		errMsg = "Failed to unassign element " + str(promotionId) + " for Promotion on Coupon"

		try:
			# get the Coupon from db
			coupon = self.get( couponId ).first()	
			
			# assign to None for unassignment
			coupon.promotion = None			

			#save it
			coupon.save()

			# reload and return the appropriate version					
			return self.get( couponId );
		except Coupon.DoesNotExist:
			raise ProcessingError(errMsg + " : Coupon with id " + str(couponId) + " does not exist.")
		except Exception:
			return None;
		
	def addRedemptions( self, couponId, redemptionsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.CouponRedemptionDelegate import CouponRedemptionDelegate

		errMsg = "Failed to add elements " + str(redemptionsIds) + " for Redemptions on Coupon"

		try:
			# get the Coupon
			coupon = self.get( couponId ).first()
				
			# split on a comma with no spaces
			idList = redemptionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the CouponRedemption		
				couponRedemption = CouponRedemptionDelegate().get(id).first();	
				# add the CouponRedemption
				coupon.redemptions.add(couponRedemption)
				
			# save it		
			coupon.save()
			
			# reload and return the appropriate version
			return self.get( couponId );
		except Coupon.DoesNotExist:
			raise ProcessingError(errMsg + " : Coupon with id " + str(couponId) + " does not exist.")
		except CouponRedemption.DoesNotExist:
			raise ProcessingError(errMsg + " : CouponRedemption does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeRedemptions( self, couponId, redemptionsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.CouponRedemptionDelegate import CouponRedemptionDelegate

		errMsg = "Failed to remove elements " + str(redemptionsIds) + " for Redemptions on Coupon"

		try:
			# get the Coupon
			coupon = self.get( couponId ).first()
				
			# split on a comma with no spaces
			idList = redemptionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the CouponRedemption		
				couponRedemption = CouponRedemptionDelegate().get(id).first();	
				# add the CouponRedemption
				coupon.redemptions.remove(couponRedemption)
				
			# save it		
			coupon.save()
			
			# reload and return the appropriate version
			return self.get( couponId );
		except Coupon.DoesNotExist:
			raise ProcessingError(errMsg + " : Coupon with id " + str(couponId) + " does not exist.")
		except CouponRedemption.DoesNotExist:
			raise ProcessingError(errMsg + " : CouponRedemption does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
