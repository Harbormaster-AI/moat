from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from fintechOnDjango.models.Invoice import Invoice
from fintechOnDjango.models.Merchant import Merchant
from fintechOnDjango.models.PaymentOrder import PaymentOrder
from fintechOnDjango.exceptions import Exceptions

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
		
	def assignMerchant( self, invoiceId, merchantId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.MerchantDelegate import MerchantDelegate

		errMsg = "Failed to assign element " + str(merchantId) + " for Merchant on Invoice"

		try:
			# get the Invoice from db
			invoice = self.get( invoiceId ).first()	
			
			# get the Merchant from db
			merchant = MerchantDelegate().get(merchantId).first();
			
			# assign the Merchant		
			invoice.merchant = merchant
			
			#save it
			invoice.save()

			# reload and return the appropriate version					
			return self.get( invoiceId );
		except Invoice.DoesNotExist:
			raise ProcessingError(errMsg + " : Invoice with id " + str(invoiceId) + " does not exist.")
		except Merchant.DoesNotExist:
			raise ProcessingError(errMsg + " : Merchant with id " + str(merchantId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignMerchant( self, invoiceId ):
		errMsg = "Failed to unassign element " + str(merchantId) + " for Merchant on Invoice"

		try:
			# get the Invoice from db
			invoice = self.get( invoiceId ).first()	
			
			# assign to None for unassignment
			invoice.merchant = None			

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
		from fintechOnDjango.delegates.PaymentOrderDelegate import PaymentOrderDelegate

		errMsg = "Failed to add elements " + str(paymentsIds) + " for Payments on Invoice"

		try:
			# get the Invoice
			invoice = self.get( invoiceId ).first()
				
			# split on a comma with no spaces
			idList = paymentsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the PaymentOrder		
				paymentOrder = PaymentOrderDelegate().get(id).first();	
				# add the PaymentOrder
				invoice.payments.add(paymentOrder)
				
			# save it		
			invoice.save()
			
			# reload and return the appropriate version
			return self.get( invoiceId );
		except Invoice.DoesNotExist:
			raise ProcessingError(errMsg + " : Invoice with id " + str(invoiceId) + " does not exist.")
		except PaymentOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentOrder does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePayments( self, invoiceId, paymentsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.PaymentOrderDelegate import PaymentOrderDelegate

		errMsg = "Failed to remove elements " + str(paymentsIds) + " for Payments on Invoice"

		try:
			# get the Invoice
			invoice = self.get( invoiceId ).first()
				
			# split on a comma with no spaces
			idList = paymentsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the PaymentOrder		
				paymentOrder = PaymentOrderDelegate().get(id).first();	
				# add the PaymentOrder
				invoice.payments.remove(paymentOrder)
				
			# save it		
			invoice.save()
			
			# reload and return the appropriate version
			return self.get( invoiceId );
		except Invoice.DoesNotExist:
			raise ProcessingError(errMsg + " : Invoice with id " + str(invoiceId) + " does not exist.")
		except PaymentOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : PaymentOrder does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
