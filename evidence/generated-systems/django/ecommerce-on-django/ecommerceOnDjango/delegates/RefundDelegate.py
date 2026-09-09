from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from ecommerceOnDjango.models.Refund import Refund
from ecommerceOnDjango.models.Payment import Payment
from ecommerceOnDjango.models.Order import Order
from ecommerceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Refund
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RefundDelegate Declaration
#======================================================================
class RefundDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, refundId ):
		try:	
			refund = Refund.objects.filter(id=refundId)
			return refund.first();
		except Refund.DoesNotExist:
			raise ProcessingError("Refund with id " + str(refundId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, refund):
		for model in serializers.deserialize("json", refund):
			model.save()
			return model;

	def create(self, refund):
		refund.save()
		return refund;

	def saveFromJson(self, refund):
		for model in serializers.deserialize("json", refund):
			model.save()
			return refund;
	
	def save(self, refund):
		refund.save()
		return refund;
	
	def delete(self, refundId ):
		errMsg = "Failed to delete Refund from db using id " + str(refundId)
		
		try:
			refund = Refund.objects.get(id=refundId)
			refund.delete()
			return True
		except Refund.DoesNotExist:
			raise ProcessingError("Refund with id " + str(refundId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Refund.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Refund from db")
		except Exception:
			return None;
		
	def assignPayment( self, refundId, paymentId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.PaymentDelegate import PaymentDelegate

		errMsg = "Failed to assign element " + str(paymentId) + " for Payment on Refund"

		try:
			# get the Refund from db
			refund = self.get( refundId ).first()	
			
			# get the Payment from db
			payment = PaymentDelegate().get(paymentId).first();
			
			# assign the Payment		
			refund.payment = payment
			
			#save it
			refund.save()

			# reload and return the appropriate version					
			return self.get( refundId );
		except Refund.DoesNotExist:
			raise ProcessingError(errMsg + " : Refund with id " + str(refundId) + " does not exist.")
		except Payment.DoesNotExist:
			raise ProcessingError(errMsg + " : Payment with id " + str(paymentId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPayment( self, refundId ):
		errMsg = "Failed to unassign element " + str(paymentId) + " for Payment on Refund"

		try:
			# get the Refund from db
			refund = self.get( refundId ).first()	
			
			# assign to None for unassignment
			refund.payment = None			

			#save it
			refund.save()

			# reload and return the appropriate version					
			return self.get( refundId );
		except Refund.DoesNotExist:
			raise ProcessingError(errMsg + " : Refund with id " + str(refundId) + " does not exist.")
		except Exception:
			return None;
		
	def assignOrder( self, refundId, orderId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.OrderDelegate import OrderDelegate

		errMsg = "Failed to assign element " + str(orderId) + " for Order on Refund"

		try:
			# get the Refund from db
			refund = self.get( refundId ).first()	
			
			# get the Order from db
			order = OrderDelegate().get(orderId).first();
			
			# assign the Order		
			refund.order = order
			
			#save it
			refund.save()

			# reload and return the appropriate version					
			return self.get( refundId );
		except Refund.DoesNotExist:
			raise ProcessingError(errMsg + " : Refund with id " + str(refundId) + " does not exist.")
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order with id " + str(orderId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrder( self, refundId ):
		errMsg = "Failed to unassign element " + str(orderId) + " for Order on Refund"

		try:
			# get the Refund from db
			refund = self.get( refundId ).first()	
			
			# assign to None for unassignment
			refund.order = None			

			#save it
			refund.save()

			# reload and return the appropriate version					
			return self.get( refundId );
		except Refund.DoesNotExist:
			raise ProcessingError(errMsg + " : Refund with id " + str(refundId) + " does not exist.")
		except Exception:
			return None;
		
