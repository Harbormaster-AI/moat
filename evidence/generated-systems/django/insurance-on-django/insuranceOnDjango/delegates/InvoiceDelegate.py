from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from insuranceOnDjango.models.Invoice import Invoice
from insuranceOnDjango.models.BillingAccount import BillingAccount
from insuranceOnDjango.models.Policy import Policy
from insuranceOnDjango.models.Payment import Payment
from insuranceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Invoice
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InvoiceDelegate Declaration
#======================================================================
class InvoiceDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, invoiceId ):
		try:	
			invoice = Invoice.objects.filter(id=invoiceId)
			return invoice.first();
		except Invoice.DoesNotExist:
			raise ProcessingError("Invoice with id " + str(invoiceId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, invoice):
		for model in serializers.deserialize("json", invoice):
			model.save()
			return model;

	def create(self, invoice):
		invoice.save()
		return invoice;

	def saveFromJson(self, invoice):
		for model in serializers.deserialize("json", invoice):
			model.save()
			return invoice;
	
	def save(self, invoice):
		invoice.save()
		return invoice;
	
	def delete(self, invoiceId ):
		errMsg = "Failed to delete Invoice from db using id " + str(invoiceId)
		
		try:
			invoice = Invoice.objects.get(id=invoiceId)
			invoice.delete()
			return True
		except Invoice.DoesNotExist:
			raise ProcessingError("Invoice with id " + str(invoiceId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Invoice.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Invoice from db")
		except Exception:
			return None;
		
	def assignBillingAccount( self, invoiceId, billingAccountId ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.BillingAccountDelegate import BillingAccountDelegate

		errMsg = "Failed to assign element " + str(billingAccountId) + " for BillingAccount on Invoice"

		try:
			# get the Invoice from db
			invoice = self.get( invoiceId ).first()	
			
			# get the BillingAccount from db
			billingAccount = BillingAccountDelegate().get(billingAccountId).first();
			
			# assign the BillingAccount		
			invoice.billingAccount = billingAccount
			
			#save it
			invoice.save()

			# reload and return the appropriate version					
			return self.get( invoiceId );
		except Invoice.DoesNotExist:
			raise ProcessingError(errMsg + " : Invoice with id " + str(invoiceId) + " does not exist.")
		except BillingAccount.DoesNotExist:
			raise ProcessingError(errMsg + " : BillingAccount with id " + str(billingAccountId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignBillingAccount( self, invoiceId ):
		errMsg = "Failed to unassign element " + str(billingAccountId) + " for BillingAccount on Invoice"

		try:
			# get the Invoice from db
			invoice = self.get( invoiceId ).first()	
			
			# assign to None for unassignment
			invoice.billingAccount = None			

			#save it
			invoice.save()

			# reload and return the appropriate version					
			return self.get( invoiceId );
		except Invoice.DoesNotExist:
			raise ProcessingError(errMsg + " : Invoice with id " + str(invoiceId) + " does not exist.")
		except Exception:
			return None;
		
	def assignPolicy( self, invoiceId, policyId ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.PolicyDelegate import PolicyDelegate

		errMsg = "Failed to assign element " + str(policyId) + " for Policy on Invoice"

		try:
			# get the Invoice from db
			invoice = self.get( invoiceId ).first()	
			
			# get the Policy from db
			policy = PolicyDelegate().get(policyId).first();
			
			# assign the Policy		
			invoice.policy = policy
			
			#save it
			invoice.save()

			# reload and return the appropriate version					
			return self.get( invoiceId );
		except Invoice.DoesNotExist:
			raise ProcessingError(errMsg + " : Invoice with id " + str(invoiceId) + " does not exist.")
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPolicy( self, invoiceId ):
		errMsg = "Failed to unassign element " + str(policyId) + " for Policy on Invoice"

		try:
			# get the Invoice from db
			invoice = self.get( invoiceId ).first()	
			
			# assign to None for unassignment
			invoice.policy = None			

			#save it
			invoice.save()

			# reload and return the appropriate version					
			return self.get( invoiceId );
		except Invoice.DoesNotExist:
			raise ProcessingError(errMsg + " : Invoice with id " + str(invoiceId) + " does not exist.")
		except Exception:
			return None;
		
	def addPayments( self, invoiceId, paymentsIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.PaymentDelegate import PaymentDelegate

		errMsg = "Failed to add elements " + str(paymentsIds) + " for Payments on Invoice"

		try:
			# get the Invoice
			invoice = self.get( invoiceId ).first()
				
			# split on a comma with no spaces
			idList = paymentsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Payment		
				payment = PaymentDelegate().get(id).first();	
				# add the Payment
				invoice.payments.add(payment)
				
			# save it		
			invoice.save()
			
			# reload and return the appropriate version
			return self.get( invoiceId );
		except Invoice.DoesNotExist:
			raise ProcessingError(errMsg + " : Invoice with id " + str(invoiceId) + " does not exist.")
		except Payment.DoesNotExist:
			raise ProcessingError(errMsg + " : Payment does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePayments( self, invoiceId, paymentsIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.PaymentDelegate import PaymentDelegate

		errMsg = "Failed to remove elements " + str(paymentsIds) + " for Payments on Invoice"

		try:
			# get the Invoice
			invoice = self.get( invoiceId ).first()
				
			# split on a comma with no spaces
			idList = paymentsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Payment		
				payment = PaymentDelegate().get(id).first();	
				# add the Payment
				invoice.payments.remove(payment)
				
			# save it		
			invoice.save()
			
			# reload and return the appropriate version
			return self.get( invoiceId );
		except Invoice.DoesNotExist:
			raise ProcessingError(errMsg + " : Invoice with id " + str(invoiceId) + " does not exist.")
		except Payment.DoesNotExist:
			raise ProcessingError(errMsg + " : Payment does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
