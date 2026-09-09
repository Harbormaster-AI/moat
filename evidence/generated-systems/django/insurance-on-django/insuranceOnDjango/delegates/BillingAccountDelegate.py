from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from insuranceOnDjango.models.BillingAccount import BillingAccount
from insuranceOnDjango.models.Customer import Customer
from insuranceOnDjango.models.Policy import Policy
from insuranceOnDjango.models.Invoice import Invoice
from insuranceOnDjango.models.Payment import Payment
from insuranceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model BillingAccount
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BillingAccountDelegate Declaration
#======================================================================
class BillingAccountDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, billingAccountId ):
		try:	
			billingAccount = BillingAccount.objects.filter(id=billingAccountId)
			return billingAccount.first();
		except BillingAccount.DoesNotExist:
			raise ProcessingError("BillingAccount with id " + str(billingAccountId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, billingAccount):
		for model in serializers.deserialize("json", billingAccount):
			model.save()
			return model;

	def create(self, billingAccount):
		billingAccount.save()
		return billingAccount;

	def saveFromJson(self, billingAccount):
		for model in serializers.deserialize("json", billingAccount):
			model.save()
			return billingAccount;
	
	def save(self, billingAccount):
		billingAccount.save()
		return billingAccount;
	
	def delete(self, billingAccountId ):
		errMsg = "Failed to delete BillingAccount from db using id " + str(billingAccountId)
		
		try:
			billingAccount = BillingAccount.objects.get(id=billingAccountId)
			billingAccount.delete()
			return True
		except BillingAccount.DoesNotExist:
			raise ProcessingError("BillingAccount with id " + str(billingAccountId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = BillingAccount.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all BillingAccount from db")
		except Exception:
			return None;
		
	def assignCustomer( self, billingAccountId, customerId ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.CustomerDelegate import CustomerDelegate

		errMsg = "Failed to assign element " + str(customerId) + " for Customer on BillingAccount"

		try:
			# get the BillingAccount from db
			billingAccount = self.get( billingAccountId ).first()	
			
			# get the Customer from db
			customer = CustomerDelegate().get(customerId).first();
			
			# assign the Customer		
			billingAccount.customer = customer
			
			#save it
			billingAccount.save()

			# reload and return the appropriate version					
			return self.get( billingAccountId );
		except BillingAccount.DoesNotExist:
			raise ProcessingError(errMsg + " : BillingAccount with id " + str(billingAccountId) + " does not exist.")
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCustomer( self, billingAccountId ):
		errMsg = "Failed to unassign element " + str(customerId) + " for Customer on BillingAccount"

		try:
			# get the BillingAccount from db
			billingAccount = self.get( billingAccountId ).first()	
			
			# assign to None for unassignment
			billingAccount.customer = None			

			#save it
			billingAccount.save()

			# reload and return the appropriate version					
			return self.get( billingAccountId );
		except BillingAccount.DoesNotExist:
			raise ProcessingError(errMsg + " : BillingAccount with id " + str(billingAccountId) + " does not exist.")
		except Exception:
			return None;
		
	def addPolicies( self, billingAccountId, policiesIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.PolicyDelegate import PolicyDelegate

		errMsg = "Failed to add elements " + str(policiesIds) + " for Policies on BillingAccount"

		try:
			# get the BillingAccount
			billingAccount = self.get( billingAccountId ).first()
				
			# split on a comma with no spaces
			idList = policiesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Policy		
				policy = PolicyDelegate().get(id).first();	
				# add the Policy
				billingAccount.policies.add(policy)
				
			# save it		
			billingAccount.save()
			
			# reload and return the appropriate version
			return self.get( billingAccountId );
		except BillingAccount.DoesNotExist:
			raise ProcessingError(errMsg + " : BillingAccount with id " + str(billingAccountId) + " does not exist.")
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePolicies( self, billingAccountId, policiesIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.PolicyDelegate import PolicyDelegate

		errMsg = "Failed to remove elements " + str(policiesIds) + " for Policies on BillingAccount"

		try:
			# get the BillingAccount
			billingAccount = self.get( billingAccountId ).first()
				
			# split on a comma with no spaces
			idList = policiesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Policy		
				policy = PolicyDelegate().get(id).first();	
				# add the Policy
				billingAccount.policies.remove(policy)
				
			# save it		
			billingAccount.save()
			
			# reload and return the appropriate version
			return self.get( billingAccountId );
		except BillingAccount.DoesNotExist:
			raise ProcessingError(errMsg + " : BillingAccount with id " + str(billingAccountId) + " does not exist.")
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addInvoices( self, billingAccountId, invoicesIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.InvoiceDelegate import InvoiceDelegate

		errMsg = "Failed to add elements " + str(invoicesIds) + " for Invoices on BillingAccount"

		try:
			# get the BillingAccount
			billingAccount = self.get( billingAccountId ).first()
				
			# split on a comma with no spaces
			idList = invoicesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Invoice		
				invoice = InvoiceDelegate().get(id).first();	
				# add the Invoice
				billingAccount.invoices.add(invoice)
				
			# save it		
			billingAccount.save()
			
			# reload and return the appropriate version
			return self.get( billingAccountId );
		except BillingAccount.DoesNotExist:
			raise ProcessingError(errMsg + " : BillingAccount with id " + str(billingAccountId) + " does not exist.")
		except Invoice.DoesNotExist:
			raise ProcessingError(errMsg + " : Invoice does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeInvoices( self, billingAccountId, invoicesIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.InvoiceDelegate import InvoiceDelegate

		errMsg = "Failed to remove elements " + str(invoicesIds) + " for Invoices on BillingAccount"

		try:
			# get the BillingAccount
			billingAccount = self.get( billingAccountId ).first()
				
			# split on a comma with no spaces
			idList = invoicesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Invoice		
				invoice = InvoiceDelegate().get(id).first();	
				# add the Invoice
				billingAccount.invoices.remove(invoice)
				
			# save it		
			billingAccount.save()
			
			# reload and return the appropriate version
			return self.get( billingAccountId );
		except BillingAccount.DoesNotExist:
			raise ProcessingError(errMsg + " : BillingAccount with id " + str(billingAccountId) + " does not exist.")
		except Invoice.DoesNotExist:
			raise ProcessingError(errMsg + " : Invoice does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addPayments( self, billingAccountId, paymentsIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.PaymentDelegate import PaymentDelegate

		errMsg = "Failed to add elements " + str(paymentsIds) + " for Payments on BillingAccount"

		try:
			# get the BillingAccount
			billingAccount = self.get( billingAccountId ).first()
				
			# split on a comma with no spaces
			idList = paymentsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Payment		
				payment = PaymentDelegate().get(id).first();	
				# add the Payment
				billingAccount.payments.add(payment)
				
			# save it		
			billingAccount.save()
			
			# reload and return the appropriate version
			return self.get( billingAccountId );
		except BillingAccount.DoesNotExist:
			raise ProcessingError(errMsg + " : BillingAccount with id " + str(billingAccountId) + " does not exist.")
		except Payment.DoesNotExist:
			raise ProcessingError(errMsg + " : Payment does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePayments( self, billingAccountId, paymentsIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.PaymentDelegate import PaymentDelegate

		errMsg = "Failed to remove elements " + str(paymentsIds) + " for Payments on BillingAccount"

		try:
			# get the BillingAccount
			billingAccount = self.get( billingAccountId ).first()
				
			# split on a comma with no spaces
			idList = paymentsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Payment		
				payment = PaymentDelegate().get(id).first();	
				# add the Payment
				billingAccount.payments.remove(payment)
				
			# save it		
			billingAccount.save()
			
			# reload and return the appropriate version
			return self.get( billingAccountId );
		except BillingAccount.DoesNotExist:
			raise ProcessingError(errMsg + " : BillingAccount with id " + str(billingAccountId) + " does not exist.")
		except Payment.DoesNotExist:
			raise ProcessingError(errMsg + " : Payment does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
