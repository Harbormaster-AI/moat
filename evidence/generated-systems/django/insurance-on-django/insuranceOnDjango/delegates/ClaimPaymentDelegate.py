from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from insuranceOnDjango.models.ClaimPayment import ClaimPayment
from insuranceOnDjango.models.Claim import Claim
from insuranceOnDjango.models.Exposure import Exposure
from insuranceOnDjango.models.Beneficiary import Beneficiary
from insuranceOnDjango.models.ServiceProvider import ServiceProvider
from insuranceOnDjango.models.Customer import Customer
from insuranceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model ClaimPayment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ClaimPaymentDelegate Declaration
#======================================================================
class ClaimPaymentDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, claimPaymentId ):
		try:	
			claimPayment = ClaimPayment.objects.filter(id=claimPaymentId)
			return claimPayment.first();
		except ClaimPayment.DoesNotExist:
			raise ProcessingError("ClaimPayment with id " + str(claimPaymentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, claimPayment):
		for model in serializers.deserialize("json", claimPayment):
			model.save()
			return model;

	def create(self, claimPayment):
		claimPayment.save()
		return claimPayment;

	def saveFromJson(self, claimPayment):
		for model in serializers.deserialize("json", claimPayment):
			model.save()
			return claimPayment;
	
	def save(self, claimPayment):
		claimPayment.save()
		return claimPayment;
	
	def delete(self, claimPaymentId ):
		errMsg = "Failed to delete ClaimPayment from db using id " + str(claimPaymentId)
		
		try:
			claimPayment = ClaimPayment.objects.get(id=claimPaymentId)
			claimPayment.delete()
			return True
		except ClaimPayment.DoesNotExist:
			raise ProcessingError("ClaimPayment with id " + str(claimPaymentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = ClaimPayment.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all ClaimPayment from db")
		except Exception:
			return None;
		
	def assignClaim( self, claimPaymentId, claimId ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.ClaimDelegate import ClaimDelegate

		errMsg = "Failed to assign element " + str(claimId) + " for Claim on ClaimPayment"

		try:
			# get the ClaimPayment from db
			claimPayment = self.get( claimPaymentId ).first()	
			
			# get the Claim from db
			claim = ClaimDelegate().get(claimId).first();
			
			# assign the Claim		
			claimPayment.claim = claim
			
			#save it
			claimPayment.save()

			# reload and return the appropriate version					
			return self.get( claimPaymentId );
		except ClaimPayment.DoesNotExist:
			raise ProcessingError(errMsg + " : ClaimPayment with id " + str(claimPaymentId) + " does not exist.")
		except Claim.DoesNotExist:
			raise ProcessingError(errMsg + " : Claim with id " + str(claimId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignClaim( self, claimPaymentId ):
		errMsg = "Failed to unassign element " + str(claimId) + " for Claim on ClaimPayment"

		try:
			# get the ClaimPayment from db
			claimPayment = self.get( claimPaymentId ).first()	
			
			# assign to None for unassignment
			claimPayment.claim = None			

			#save it
			claimPayment.save()

			# reload and return the appropriate version					
			return self.get( claimPaymentId );
		except ClaimPayment.DoesNotExist:
			raise ProcessingError(errMsg + " : ClaimPayment with id " + str(claimPaymentId) + " does not exist.")
		except Exception:
			return None;
		
	def assignExposure( self, claimPaymentId, exposureId ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.ExposureDelegate import ExposureDelegate

		errMsg = "Failed to assign element " + str(exposureId) + " for Exposure on ClaimPayment"

		try:
			# get the ClaimPayment from db
			claimPayment = self.get( claimPaymentId ).first()	
			
			# get the Exposure from db
			exposure = ExposureDelegate().get(exposureId).first();
			
			# assign the Exposure		
			claimPayment.exposure = exposure
			
			#save it
			claimPayment.save()

			# reload and return the appropriate version					
			return self.get( claimPaymentId );
		except ClaimPayment.DoesNotExist:
			raise ProcessingError(errMsg + " : ClaimPayment with id " + str(claimPaymentId) + " does not exist.")
		except Exposure.DoesNotExist:
			raise ProcessingError(errMsg + " : Exposure with id " + str(exposureId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignExposure( self, claimPaymentId ):
		errMsg = "Failed to unassign element " + str(exposureId) + " for Exposure on ClaimPayment"

		try:
			# get the ClaimPayment from db
			claimPayment = self.get( claimPaymentId ).first()	
			
			# assign to None for unassignment
			claimPayment.exposure = None			

			#save it
			claimPayment.save()

			# reload and return the appropriate version					
			return self.get( claimPaymentId );
		except ClaimPayment.DoesNotExist:
			raise ProcessingError(errMsg + " : ClaimPayment with id " + str(claimPaymentId) + " does not exist.")
		except Exception:
			return None;
		
	def assignBeneficiary( self, claimPaymentId, beneficiaryId ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.BeneficiaryDelegate import BeneficiaryDelegate

		errMsg = "Failed to assign element " + str(beneficiaryId) + " for Beneficiary on ClaimPayment"

		try:
			# get the ClaimPayment from db
			claimPayment = self.get( claimPaymentId ).first()	
			
			# get the Beneficiary from db
			beneficiary = BeneficiaryDelegate().get(beneficiaryId).first();
			
			# assign the Beneficiary		
			claimPayment.beneficiary = beneficiary
			
			#save it
			claimPayment.save()

			# reload and return the appropriate version					
			return self.get( claimPaymentId );
		except ClaimPayment.DoesNotExist:
			raise ProcessingError(errMsg + " : ClaimPayment with id " + str(claimPaymentId) + " does not exist.")
		except Beneficiary.DoesNotExist:
			raise ProcessingError(errMsg + " : Beneficiary with id " + str(beneficiaryId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignBeneficiary( self, claimPaymentId ):
		errMsg = "Failed to unassign element " + str(beneficiaryId) + " for Beneficiary on ClaimPayment"

		try:
			# get the ClaimPayment from db
			claimPayment = self.get( claimPaymentId ).first()	
			
			# assign to None for unassignment
			claimPayment.beneficiary = None			

			#save it
			claimPayment.save()

			# reload and return the appropriate version					
			return self.get( claimPaymentId );
		except ClaimPayment.DoesNotExist:
			raise ProcessingError(errMsg + " : ClaimPayment with id " + str(claimPaymentId) + " does not exist.")
		except Exception:
			return None;
		
	def assignServiceProvider( self, claimPaymentId, serviceProviderId ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.ServiceProviderDelegate import ServiceProviderDelegate

		errMsg = "Failed to assign element " + str(serviceProviderId) + " for ServiceProvider on ClaimPayment"

		try:
			# get the ClaimPayment from db
			claimPayment = self.get( claimPaymentId ).first()	
			
			# get the ServiceProvider from db
			serviceProvider = ServiceProviderDelegate().get(serviceProviderId).first();
			
			# assign the ServiceProvider		
			claimPayment.serviceProvider = serviceProvider
			
			#save it
			claimPayment.save()

			# reload and return the appropriate version					
			return self.get( claimPaymentId );
		except ClaimPayment.DoesNotExist:
			raise ProcessingError(errMsg + " : ClaimPayment with id " + str(claimPaymentId) + " does not exist.")
		except ServiceProvider.DoesNotExist:
			raise ProcessingError(errMsg + " : ServiceProvider with id " + str(serviceProviderId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignServiceProvider( self, claimPaymentId ):
		errMsg = "Failed to unassign element " + str(serviceProviderId) + " for ServiceProvider on ClaimPayment"

		try:
			# get the ClaimPayment from db
			claimPayment = self.get( claimPaymentId ).first()	
			
			# assign to None for unassignment
			claimPayment.serviceProvider = None			

			#save it
			claimPayment.save()

			# reload and return the appropriate version					
			return self.get( claimPaymentId );
		except ClaimPayment.DoesNotExist:
			raise ProcessingError(errMsg + " : ClaimPayment with id " + str(claimPaymentId) + " does not exist.")
		except Exception:
			return None;
		
	def assignCustomer( self, claimPaymentId, customerId ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.CustomerDelegate import CustomerDelegate

		errMsg = "Failed to assign element " + str(customerId) + " for Customer on ClaimPayment"

		try:
			# get the ClaimPayment from db
			claimPayment = self.get( claimPaymentId ).first()	
			
			# get the Customer from db
			customer = CustomerDelegate().get(customerId).first();
			
			# assign the Customer		
			claimPayment.customer = customer
			
			#save it
			claimPayment.save()

			# reload and return the appropriate version					
			return self.get( claimPaymentId );
		except ClaimPayment.DoesNotExist:
			raise ProcessingError(errMsg + " : ClaimPayment with id " + str(claimPaymentId) + " does not exist.")
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCustomer( self, claimPaymentId ):
		errMsg = "Failed to unassign element " + str(customerId) + " for Customer on ClaimPayment"

		try:
			# get the ClaimPayment from db
			claimPayment = self.get( claimPaymentId ).first()	
			
			# assign to None for unassignment
			claimPayment.customer = None			

			#save it
			claimPayment.save()

			# reload and return the appropriate version					
			return self.get( claimPaymentId );
		except ClaimPayment.DoesNotExist:
			raise ProcessingError(errMsg + " : ClaimPayment with id " + str(claimPaymentId) + " does not exist.")
		except Exception:
			return None;
		
