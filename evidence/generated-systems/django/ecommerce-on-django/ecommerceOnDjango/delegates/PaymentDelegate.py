from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from ecommerceOnDjango.models.Payment import Payment
from ecommerceOnDjango.models.Order import Order
from ecommerceOnDjango.models.Customer import Customer
from ecommerceOnDjango.models.PaymentProvider import PaymentProvider
from ecommerceOnDjango.models.Refund import Refund
from ecommerceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Payment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PaymentDelegate Declaration
#======================================================================
class PaymentDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, paymentId ):
		try:	
			payment = Payment.objects.filter(id=paymentId)
			return payment.first();
		except Payment.DoesNotExist:
			raise ProcessingError("Payment with id " + str(paymentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, payment):
		for model in serializers.deserialize("json", payment):
			model.save()
			return model;

	def create(self, payment):
		payment.save()
		return payment;

	def saveFromJson(self, payment):
		for model in serializers.deserialize("json", payment):
			model.save()
			return payment;
	
	def save(self, payment):
		payment.save()
		return payment;
	
	def delete(self, paymentId ):
		errMsg = "Failed to delete Payment from db using id " + str(paymentId)
		
		try:
			payment = Payment.objects.get(id=paymentId)
			payment.delete()
			return True
		except Payment.DoesNotExist:
			raise ProcessingError("Payment with id " + str(paymentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Payment.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Payment from db")
		except Exception:
			return None;
		
	def assignOrder( self, paymentId, orderId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.OrderDelegate import OrderDelegate

		errMsg = "Failed to assign element " + str(orderId) + " for Order on Payment"

		try:
			# get the Payment from db
			payment = self.get( paymentId ).first()	
			
			# get the Order from db
			order = OrderDelegate().get(orderId).first();
			
			# assign the Order		
			payment.order = order
			
			#save it
			payment.save()

			# reload and return the appropriate version					
			return self.get( paymentId );
		except Payment.DoesNotExist:
			raise ProcessingError(errMsg + " : Payment with id " + str(paymentId) + " does not exist.")
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order with id " + str(orderId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrder( self, paymentId ):
		errMsg = "Failed to unassign element " + str(orderId) + " for Order on Payment"

		try:
			# get the Payment from db
			payment = self.get( paymentId ).first()	
			
			# assign to None for unassignment
			payment.order = None			

			#save it
			payment.save()

			# reload and return the appropriate version					
			return self.get( paymentId );
		except Payment.DoesNotExist:
			raise ProcessingError(errMsg + " : Payment with id " + str(paymentId) + " does not exist.")
		except Exception:
			return None;
		
	def assignCustomer( self, paymentId, customerId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.CustomerDelegate import CustomerDelegate

		errMsg = "Failed to assign element " + str(customerId) + " for Customer on Payment"

		try:
			# get the Payment from db
			payment = self.get( paymentId ).first()	
			
			# get the Customer from db
			customer = CustomerDelegate().get(customerId).first();
			
			# assign the Customer		
			payment.customer = customer
			
			#save it
			payment.save()

			# reload and return the appropriate version					
			return self.get( paymentId );
		except Payment.DoesNotExist:
			raise ProcessingError(errMsg + " : Payment with id " + str(paymentId) + " does not exist.")
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCustomer( self, paymentId ):
		errMsg = "Failed to unassign element " + str(customerId) + " for Customer on Payment"

		try:
			# get the Payment from db
			payment = self.get( paymentId ).first()	
			
			# assign to None for unassignment
			payment.customer = None			

			#save it
			payment.save()

			# reload and return the appropriate version					
			return self.get( paymentId );
		except Payment.DoesNotExist:
			raise ProcessingError(errMsg + " : Payment with id " + str(paymentId) + " does not exist.")
		except Exception:
			return None;
		
	def assignPaymentProvider( self, paymentId, paymentProviderId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.PaymentProviderDelegate import PaymentProviderDelegate

		errMsg = "Failed to assign element " + str(paymentProviderId) + " for PaymentProvider on Payment"

		try:
			# get the Payment from db
			payment = self.get( paymentId ).first()	
			
			# get the PaymentProvider from db
			paymentProvider = PaymentProviderDelegate().get(paymentProviderId).first();
			
			# assign the PaymentProvider		
			payment.paymentProvider = paymentProvider
			
			#save it
			payment.save()

			# reload and return the appropriate version					
			return self.get( paymentId );
		except Payment.DoesNotExist:
			raise ProcessingError(errMsg + " : Payment with id " + str(paymentId) + " does not exist.")
		except PaymentProvider.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentProvider with id " + str(paymentProviderId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPaymentProvider( self, paymentId ):
		errMsg = "Failed to unassign element " + str(paymentProviderId) + " for PaymentProvider on Payment"

		try:
			# get the Payment from db
			payment = self.get( paymentId ).first()	
			
			# assign to None for unassignment
			payment.paymentProvider = None			

			#save it
			payment.save()

			# reload and return the appropriate version					
			return self.get( paymentId );
		except Payment.DoesNotExist:
			raise ProcessingError(errMsg + " : Payment with id " + str(paymentId) + " does not exist.")
		except Exception:
			return None;
		
	def addRefunds( self, paymentId, refundsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.RefundDelegate import RefundDelegate

		errMsg = "Failed to add elements " + str(refundsIds) + " for Refunds on Payment"

		try:
			# get the Payment
			payment = self.get( paymentId ).first()
				
			# split on a comma with no spaces
			idList = refundsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Refund		
				refund = RefundDelegate().get(id).first();	
				# add the Refund
				payment.refunds.add(refund)
				
			# save it		
			payment.save()
			
			# reload and return the appropriate version
			return self.get( paymentId );
		except Payment.DoesNotExist:
			raise ProcessingError(errMsg + " : Payment with id " + str(paymentId) + " does not exist.")
		except Refund.DoesNotExist:
			raise ProcessingError(errMsg + " : Refund does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeRefunds( self, paymentId, refundsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.RefundDelegate import RefundDelegate

		errMsg = "Failed to remove elements " + str(refundsIds) + " for Refunds on Payment"

		try:
			# get the Payment
			payment = self.get( paymentId ).first()
				
			# split on a comma with no spaces
			idList = refundsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Refund		
				refund = RefundDelegate().get(id).first();	
				# add the Refund
				payment.refunds.remove(refund)
				
			# save it		
			payment.save()
			
			# reload and return the appropriate version
			return self.get( paymentId );
		except Payment.DoesNotExist:
			raise ProcessingError(errMsg + " : Payment with id " + str(paymentId) + " does not exist.")
		except Refund.DoesNotExist:
			raise ProcessingError(errMsg + " : Refund does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
