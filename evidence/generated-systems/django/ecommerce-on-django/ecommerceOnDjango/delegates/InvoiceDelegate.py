from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from ecommerceOnDjango.models.Invoice import Invoice
from ecommerceOnDjango.models.Order import Order
from ecommerceOnDjango.exceptions import Exceptions

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
		
	def assignOrder( self, invoiceId, orderId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.OrderDelegate import OrderDelegate

		errMsg = "Failed to assign element " + str(orderId) + " for Order on Invoice"

		try:
			# get the Invoice from db
			invoice = self.get( invoiceId ).first()	
			
			# get the Order from db
			order = OrderDelegate().get(orderId).first();
			
			# assign the Order		
			invoice.order = order
			
			#save it
			invoice.save()

			# reload and return the appropriate version					
			return self.get( invoiceId );
		except Invoice.DoesNotExist:
			raise ProcessingError(errMsg + " : Invoice with id " + str(invoiceId) + " does not exist.")
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order with id " + str(orderId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrder( self, invoiceId ):
		errMsg = "Failed to unassign element " + str(orderId) + " for Order on Invoice"

		try:
			# get the Invoice from db
			invoice = self.get( invoiceId ).first()	
			
			# assign to None for unassignment
			invoice.order = None			

			#save it
			invoice.save()

			# reload and return the appropriate version					
			return self.get( invoiceId );
		except Invoice.DoesNotExist:
			raise ProcessingError(errMsg + " : Invoice with id " + str(invoiceId) + " does not exist.")
		except Exception:
			return None;
		
