from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from insuranceOnDjango.models.Beneficiary import Beneficiary
from insuranceOnDjango.models.Policy import Policy
from insuranceOnDjango.models.Customer import Customer
from insuranceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Beneficiary
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class BeneficiaryDelegate Declaration
#======================================================================
class BeneficiaryDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, beneficiaryId ):
		try:	
			beneficiary = Beneficiary.objects.filter(id=beneficiaryId)
			return beneficiary.first();
		except Beneficiary.DoesNotExist:
			raise ProcessingError("Beneficiary with id " + str(beneficiaryId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, beneficiary):
		for model in serializers.deserialize("json", beneficiary):
			model.save()
			return model;

	def create(self, beneficiary):
		beneficiary.save()
		return beneficiary;

	def saveFromJson(self, beneficiary):
		for model in serializers.deserialize("json", beneficiary):
			model.save()
			return beneficiary;
	
	def save(self, beneficiary):
		beneficiary.save()
		return beneficiary;
	
	def delete(self, beneficiaryId ):
		errMsg = "Failed to delete Beneficiary from db using id " + str(beneficiaryId)
		
		try:
			beneficiary = Beneficiary.objects.get(id=beneficiaryId)
			beneficiary.delete()
			return True
		except Beneficiary.DoesNotExist:
			raise ProcessingError("Beneficiary with id " + str(beneficiaryId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Beneficiary.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Beneficiary from db")
		except Exception:
			return None;
		
	def assignPolicy( self, beneficiaryId, policyId ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.PolicyDelegate import PolicyDelegate

		errMsg = "Failed to assign element " + str(policyId) + " for Policy on Beneficiary"

		try:
			# get the Beneficiary from db
			beneficiary = self.get( beneficiaryId ).first()	
			
			# get the Policy from db
			policy = PolicyDelegate().get(policyId).first();
			
			# assign the Policy		
			beneficiary.policy = policy
			
			#save it
			beneficiary.save()

			# reload and return the appropriate version					
			return self.get( beneficiaryId );
		except Beneficiary.DoesNotExist:
			raise ProcessingError(errMsg + " : Beneficiary with id " + str(beneficiaryId) + " does not exist.")
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPolicy( self, beneficiaryId ):
		errMsg = "Failed to unassign element " + str(policyId) + " for Policy on Beneficiary"

		try:
			# get the Beneficiary from db
			beneficiary = self.get( beneficiaryId ).first()	
			
			# assign to None for unassignment
			beneficiary.policy = None			

			#save it
			beneficiary.save()

			# reload and return the appropriate version					
			return self.get( beneficiaryId );
		except Beneficiary.DoesNotExist:
			raise ProcessingError(errMsg + " : Beneficiary with id " + str(beneficiaryId) + " does not exist.")
		except Exception:
			return None;
		
	def assignCustomer( self, beneficiaryId, customerId ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.CustomerDelegate import CustomerDelegate

		errMsg = "Failed to assign element " + str(customerId) + " for Customer on Beneficiary"

		try:
			# get the Beneficiary from db
			beneficiary = self.get( beneficiaryId ).first()	
			
			# get the Customer from db
			customer = CustomerDelegate().get(customerId).first();
			
			# assign the Customer		
			beneficiary.customer = customer
			
			#save it
			beneficiary.save()

			# reload and return the appropriate version					
			return self.get( beneficiaryId );
		except Beneficiary.DoesNotExist:
			raise ProcessingError(errMsg + " : Beneficiary with id " + str(beneficiaryId) + " does not exist.")
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCustomer( self, beneficiaryId ):
		errMsg = "Failed to unassign element " + str(customerId) + " for Customer on Beneficiary"

		try:
			# get the Beneficiary from db
			beneficiary = self.get( beneficiaryId ).first()	
			
			# assign to None for unassignment
			beneficiary.customer = None			

			#save it
			beneficiary.save()

			# reload and return the appropriate version					
			return self.get( beneficiaryId );
		except Beneficiary.DoesNotExist:
			raise ProcessingError(errMsg + " : Beneficiary with id " + str(beneficiaryId) + " does not exist.")
		except Exception:
			return None;
		
