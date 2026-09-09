from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from fintechOnDjango.models.Agreement import Agreement
from fintechOnDjango.models.Customer import Customer
from fintechOnDjango.models.ProductOffering import ProductOffering
from fintechOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Agreement
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AgreementDelegate Declaration
#======================================================================
class AgreementDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, agreementId ):
		try:	
			agreement = Agreement.objects.filter(id=agreementId)
			return agreement.first();
		except Agreement.DoesNotExist:
			raise ProcessingError("Agreement with id " + str(agreementId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, agreement):
		for model in serializers.deserialize("json", agreement):
			model.save()
			return model;

	def create(self, agreement):
		agreement.save()
		return agreement;

	def saveFromJson(self, agreement):
		for model in serializers.deserialize("json", agreement):
			model.save()
			return agreement;
	
	def save(self, agreement):
		agreement.save()
		return agreement;
	
	def delete(self, agreementId ):
		errMsg = "Failed to delete Agreement from db using id " + str(agreementId)
		
		try:
			agreement = Agreement.objects.get(id=agreementId)
			agreement.delete()
			return True
		except Agreement.DoesNotExist:
			raise ProcessingError("Agreement with id " + str(agreementId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Agreement.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Agreement from db")
		except Exception:
			return None;
		
	def assignCustomer( self, agreementId, customerId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.CustomerDelegate import CustomerDelegate

		errMsg = "Failed to assign element " + str(customerId) + " for Customer on Agreement"

		try:
			# get the Agreement from db
			agreement = self.get( agreementId ).first()	
			
			# get the Customer from db
			customer = CustomerDelegate().get(customerId).first();
			
			# assign the Customer		
			agreement.customer = customer
			
			#save it
			agreement.save()

			# reload and return the appropriate version					
			return self.get( agreementId );
		except Agreement.DoesNotExist:
			raise ProcessingError(errMsg + " : Agreement with id " + str(agreementId) + " does not exist.")
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCustomer( self, agreementId ):
		errMsg = "Failed to unassign element " + str(customerId) + " for Customer on Agreement"

		try:
			# get the Agreement from db
			agreement = self.get( agreementId ).first()	
			
			# assign to None for unassignment
			agreement.customer = None			

			#save it
			agreement.save()

			# reload and return the appropriate version					
			return self.get( agreementId );
		except Agreement.DoesNotExist:
			raise ProcessingError(errMsg + " : Agreement with id " + str(agreementId) + " does not exist.")
		except Exception:
			return None;
		
	def assignProductOffering( self, agreementId, productOfferingId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.ProductOfferingDelegate import ProductOfferingDelegate

		errMsg = "Failed to assign element " + str(productOfferingId) + " for ProductOffering on Agreement"

		try:
			# get the Agreement from db
			agreement = self.get( agreementId ).first()	
			
			# get the ProductOffering from db
			productOffering = ProductOfferingDelegate().get(productOfferingId).first();
			
			# assign the ProductOffering		
			agreement.productOffering = productOffering
			
			#save it
			agreement.save()

			# reload and return the appropriate version					
			return self.get( agreementId );
		except Agreement.DoesNotExist:
			raise ProcessingError(errMsg + " : Agreement with id " + str(agreementId) + " does not exist.")
		except ProductOffering.DoesNotExist:
			raise ProcessingError(errMsg + " : ProductOffering with id " + str(productOfferingId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignProductOffering( self, agreementId ):
		errMsg = "Failed to unassign element " + str(productOfferingId) + " for ProductOffering on Agreement"

		try:
			# get the Agreement from db
			agreement = self.get( agreementId ).first()	
			
			# assign to None for unassignment
			agreement.productOffering = None			

			#save it
			agreement.save()

			# reload and return the appropriate version					
			return self.get( agreementId );
		except Agreement.DoesNotExist:
			raise ProcessingError(errMsg + " : Agreement with id " + str(agreementId) + " does not exist.")
		except Exception:
			return None;
		
