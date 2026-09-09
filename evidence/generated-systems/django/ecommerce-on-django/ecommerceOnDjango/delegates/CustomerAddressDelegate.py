from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from ecommerceOnDjango.models.CustomerAddress import CustomerAddress
from ecommerceOnDjango.models.Customer import Customer
from ecommerceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model CustomerAddress
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CustomerAddressDelegate Declaration
#======================================================================
class CustomerAddressDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, customerAddressId ):
		try:	
			customerAddress = CustomerAddress.objects.filter(id=customerAddressId)
			return customerAddress.first();
		except CustomerAddress.DoesNotExist:
			raise ProcessingError("CustomerAddress with id " + str(customerAddressId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, customerAddress):
		for model in serializers.deserialize("json", customerAddress):
			model.save()
			return model;

	def create(self, customerAddress):
		customerAddress.save()
		return customerAddress;

	def saveFromJson(self, customerAddress):
		for model in serializers.deserialize("json", customerAddress):
			model.save()
			return customerAddress;
	
	def save(self, customerAddress):
		customerAddress.save()
		return customerAddress;
	
	def delete(self, customerAddressId ):
		errMsg = "Failed to delete CustomerAddress from db using id " + str(customerAddressId)
		
		try:
			customerAddress = CustomerAddress.objects.get(id=customerAddressId)
			customerAddress.delete()
			return True
		except CustomerAddress.DoesNotExist:
			raise ProcessingError("CustomerAddress with id " + str(customerAddressId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = CustomerAddress.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all CustomerAddress from db")
		except Exception:
			return None;
		
	def assignCustomer( self, customerAddressId, customerId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.CustomerDelegate import CustomerDelegate

		errMsg = "Failed to assign element " + str(customerId) + " for Customer on CustomerAddress"

		try:
			# get the CustomerAddress from db
			customerAddress = self.get( customerAddressId ).first()	
			
			# get the Customer from db
			customer = CustomerDelegate().get(customerId).first();
			
			# assign the Customer		
			customerAddress.customer = customer
			
			#save it
			customerAddress.save()

			# reload and return the appropriate version					
			return self.get( customerAddressId );
		except CustomerAddress.DoesNotExist:
			raise ProcessingError(errMsg + " : CustomerAddress with id " + str(customerAddressId) + " does not exist.")
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCustomer( self, customerAddressId ):
		errMsg = "Failed to unassign element " + str(customerId) + " for Customer on CustomerAddress"

		try:
			# get the CustomerAddress from db
			customerAddress = self.get( customerAddressId ).first()	
			
			# assign to None for unassignment
			customerAddress.customer = None			

			#save it
			customerAddress.save()

			# reload and return the appropriate version					
			return self.get( customerAddressId );
		except CustomerAddress.DoesNotExist:
			raise ProcessingError(errMsg + " : CustomerAddress with id " + str(customerAddressId) + " does not exist.")
		except Exception:
			return None;
		
