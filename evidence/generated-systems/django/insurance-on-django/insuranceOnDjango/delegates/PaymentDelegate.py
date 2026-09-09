from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from insuranceOnDjango.models.Payment import Payment
from insuranceOnDjango.models.Invoice import Invoice
from insuranceOnDjango.models.BillingAccount import BillingAccount
from insuranceOnDjango.models.Policy import Policy
from insuranceOnDjango.exceptions import Exceptions

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
		
	def assignInvoice( self, paymentId, invoiceId ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.InvoiceDelegate import InvoiceDelegate

		errMsg = "Failed to assign element " + str(invoiceId) + " for Invoice on Payment"

		try:
			# get the Payment from db
			payment = self.get( paymentId ).first()	
			
			# get the Invoice from db
			invoice = InvoiceDelegate().get(invoiceId).first();
			
			# assign the Invoice		
			payment.invoice = invoice
			
			#save it
			payment.save()

			# reload and return the appropriate version					
			return self.get( paymentId );
		except Payment.DoesNotExist:
			raise ProcessingError(errMsg + " : Payment with id " + str(paymentId) + " does not exist.")
		except Invoice.DoesNotExist:
			raise ProcessingError(errMsg + " : Invoice with id " + str(invoiceId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignInvoice( self, paymentId ):
		errMsg = "Failed to unassign element " + str(invoiceId) + " for Invoice on Payment"

		try:
			# get the Payment from db
			payment = self.get( paymentId ).first()	
			
			# assign to None for unassignment
			payment.invoice = None			

			#save it
			payment.save()

			# reload and return the appropriate version					
			return self.get( paymentId );
		except Payment.DoesNotExist:
			raise ProcessingError(errMsg + " : Payment with id " + str(paymentId) + " does not exist.")
		except Exception:
			return None;
		
	def assignBillingAccount( self, paymentId, billingAccountId ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.BillingAccountDelegate import BillingAccountDelegate

		errMsg = "Failed to assign element " + str(billingAccountId) + " for BillingAccount on Payment"

		try:
			# get the Payment from db
			payment = self.get( paymentId ).first()	
			
			# get the BillingAccount from db
			billingAccount = BillingAccountDelegate().get(billingAccountId).first();
			
			# assign the BillingAccount		
			payment.billingAccount = billingAccount
			
			#save it
			payment.save()

			# reload and return the appropriate version					
			return self.get( paymentId );
		except Payment.DoesNotExist:
			raise ProcessingError(errMsg + " : Payment with id " + str(paymentId) + " does not exist.")
		except BillingAccount.DoesNotExist:
			raise ProcessingError(errMsg + " : BillingAccount with id " + str(billingAccountId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignBillingAccount( self, paymentId ):
		errMsg = "Failed to unassign element " + str(billingAccountId) + " for BillingAccount on Payment"

		try:
			# get the Payment from db
			payment = self.get( paymentId ).first()	
			
			# assign to None for unassignment
			payment.billingAccount = None			

			#save it
			payment.save()

			# reload and return the appropriate version					
			return self.get( paymentId );
		except Payment.DoesNotExist:
			raise ProcessingError(errMsg + " : Payment with id " + str(paymentId) + " does not exist.")
		except Exception:
			return None;
		
	def assignPolicy( self, paymentId, policyId ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.PolicyDelegate import PolicyDelegate

		errMsg = "Failed to assign element " + str(policyId) + " for Policy on Payment"

		try:
			# get the Payment from db
			payment = self.get( paymentId ).first()	
			
			# get the Policy from db
			policy = PolicyDelegate().get(policyId).first();
			
			# assign the Policy		
			payment.policy = policy
			
			#save it
			payment.save()

			# reload and return the appropriate version					
			return self.get( paymentId );
		except Payment.DoesNotExist:
			raise ProcessingError(errMsg + " : Payment with id " + str(paymentId) + " does not exist.")
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPolicy( self, paymentId ):
		errMsg = "Failed to unassign element " + str(policyId) + " for Policy on Payment"

		try:
			# get the Payment from db
			payment = self.get( paymentId ).first()	
			
			# assign to None for unassignment
			payment.policy = None			

			#save it
			payment.save()

			# reload and return the appropriate version					
			return self.get( paymentId );
		except Payment.DoesNotExist:
			raise ProcessingError(errMsg + " : Payment with id " + str(paymentId) + " does not exist.")
		except Exception:
			return None;
		
