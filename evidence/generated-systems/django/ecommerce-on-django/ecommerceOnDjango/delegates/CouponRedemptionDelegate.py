from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from ecommerceOnDjango.models.CouponRedemption import CouponRedemption
from ecommerceOnDjango.models.Coupon import Coupon
from ecommerceOnDjango.models.Order import Order
from ecommerceOnDjango.models.Customer import Customer
from ecommerceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model CouponRedemption
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CouponRedemptionDelegate Declaration
#======================================================================
class CouponRedemptionDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, couponRedemptionId ):
		try:	
			couponRedemption = CouponRedemption.objects.filter(id=couponRedemptionId)
			return couponRedemption.first();
		except CouponRedemption.DoesNotExist:
			raise ProcessingError("CouponRedemption with id " + str(couponRedemptionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, couponRedemption):
		for model in serializers.deserialize("json", couponRedemption):
			model.save()
			return model;

	def create(self, couponRedemption):
		couponRedemption.save()
		return couponRedemption;

	def saveFromJson(self, couponRedemption):
		for model in serializers.deserialize("json", couponRedemption):
			model.save()
			return couponRedemption;
	
	def save(self, couponRedemption):
		couponRedemption.save()
		return couponRedemption;
	
	def delete(self, couponRedemptionId ):
		errMsg = "Failed to delete CouponRedemption from db using id " + str(couponRedemptionId)
		
		try:
			couponRedemption = CouponRedemption.objects.get(id=couponRedemptionId)
			couponRedemption.delete()
			return True
		except CouponRedemption.DoesNotExist:
			raise ProcessingError("CouponRedemption with id " + str(couponRedemptionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = CouponRedemption.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all CouponRedemption from db")
		except Exception:
			return None;
		
	def assignCoupon( self, couponRedemptionId, couponId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.CouponDelegate import CouponDelegate

		errMsg = "Failed to assign element " + str(couponId) + " for Coupon on CouponRedemption"

		try:
			# get the CouponRedemption from db
			couponRedemption = self.get( couponRedemptionId ).first()	
			
			# get the Coupon from db
			coupon = CouponDelegate().get(couponId).first();
			
			# assign the Coupon		
			couponRedemption.coupon = coupon
			
			#save it
			couponRedemption.save()

			# reload and return the appropriate version					
			return self.get( couponRedemptionId );
		except CouponRedemption.DoesNotExist:
			raise ProcessingError(errMsg + " : CouponRedemption with id " + str(couponRedemptionId) + " does not exist.")
		except Coupon.DoesNotExist:
			raise ProcessingError(errMsg + " : Coupon with id " + str(couponId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCoupon( self, couponRedemptionId ):
		errMsg = "Failed to unassign element " + str(couponId) + " for Coupon on CouponRedemption"

		try:
			# get the CouponRedemption from db
			couponRedemption = self.get( couponRedemptionId ).first()	
			
			# assign to None for unassignment
			couponRedemption.coupon = None			

			#save it
			couponRedemption.save()

			# reload and return the appropriate version					
			return self.get( couponRedemptionId );
		except CouponRedemption.DoesNotExist:
			raise ProcessingError(errMsg + " : CouponRedemption with id " + str(couponRedemptionId) + " does not exist.")
		except Exception:
			return None;
		
	def assignOrder( self, couponRedemptionId, orderId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.OrderDelegate import OrderDelegate

		errMsg = "Failed to assign element " + str(orderId) + " for Order on CouponRedemption"

		try:
			# get the CouponRedemption from db
			couponRedemption = self.get( couponRedemptionId ).first()	
			
			# get the Order from db
			order = OrderDelegate().get(orderId).first();
			
			# assign the Order		
			couponRedemption.order = order
			
			#save it
			couponRedemption.save()

			# reload and return the appropriate version					
			return self.get( couponRedemptionId );
		except CouponRedemption.DoesNotExist:
			raise ProcessingError(errMsg + " : CouponRedemption with id " + str(couponRedemptionId) + " does not exist.")
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order with id " + str(orderId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrder( self, couponRedemptionId ):
		errMsg = "Failed to unassign element " + str(orderId) + " for Order on CouponRedemption"

		try:
			# get the CouponRedemption from db
			couponRedemption = self.get( couponRedemptionId ).first()	
			
			# assign to None for unassignment
			couponRedemption.order = None			

			#save it
			couponRedemption.save()

			# reload and return the appropriate version					
			return self.get( couponRedemptionId );
		except CouponRedemption.DoesNotExist:
			raise ProcessingError(errMsg + " : CouponRedemption with id " + str(couponRedemptionId) + " does not exist.")
		except Exception:
			return None;
		
	def assignCustomer( self, couponRedemptionId, customerId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.CustomerDelegate import CustomerDelegate

		errMsg = "Failed to assign element " + str(customerId) + " for Customer on CouponRedemption"

		try:
			# get the CouponRedemption from db
			couponRedemption = self.get( couponRedemptionId ).first()	
			
			# get the Customer from db
			customer = CustomerDelegate().get(customerId).first();
			
			# assign the Customer		
			couponRedemption.customer = customer
			
			#save it
			couponRedemption.save()

			# reload and return the appropriate version					
			return self.get( couponRedemptionId );
		except CouponRedemption.DoesNotExist:
			raise ProcessingError(errMsg + " : CouponRedemption with id " + str(couponRedemptionId) + " does not exist.")
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCustomer( self, couponRedemptionId ):
		errMsg = "Failed to unassign element " + str(customerId) + " for Customer on CouponRedemption"

		try:
			# get the CouponRedemption from db
			couponRedemption = self.get( couponRedemptionId ).first()	
			
			# assign to None for unassignment
			couponRedemption.customer = None			

			#save it
			couponRedemption.save()

			# reload and return the appropriate version					
			return self.get( couponRedemptionId );
		except CouponRedemption.DoesNotExist:
			raise ProcessingError(errMsg + " : CouponRedemption with id " + str(couponRedemptionId) + " does not exist.")
		except Exception:
			return None;
		
