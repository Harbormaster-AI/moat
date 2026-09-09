from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from insuranceOnDjango.models.Claim import Claim
from insuranceOnDjango.models.Policy import Policy
from insuranceOnDjango.models.Customer import Customer
from insuranceOnDjango.models.Adjuster import Adjuster
from insuranceOnDjango.models.Incident import Incident
from insuranceOnDjango.models.Exposure import Exposure
from insuranceOnDjango.models.ClaimReserve import ClaimReserve
from insuranceOnDjango.models.ClaimPayment import ClaimPayment
from insuranceOnDjango.models.ServiceProvider import ServiceProvider
from insuranceOnDjango.models.SubrogationRecovery import SubrogationRecovery
from insuranceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Claim
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ClaimDelegate Declaration
#======================================================================
class ClaimDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, claimId ):
		try:	
			claim = Claim.objects.filter(id=claimId)
			return claim.first();
		except Claim.DoesNotExist:
			raise ProcessingError("Claim with id " + str(claimId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, claim):
		for model in serializers.deserialize("json", claim):
			model.save()
			return model;

	def create(self, claim):
		claim.save()
		return claim;

	def saveFromJson(self, claim):
		for model in serializers.deserialize("json", claim):
			model.save()
			return claim;
	
	def save(self, claim):
		claim.save()
		return claim;
	
	def delete(self, claimId ):
		errMsg = "Failed to delete Claim from db using id " + str(claimId)
		
		try:
			claim = Claim.objects.get(id=claimId)
			claim.delete()
			return True
		except Claim.DoesNotExist:
			raise ProcessingError("Claim with id " + str(claimId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Claim.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Claim from db")
		except Exception:
			return None;
		
	def assignPolicy( self, claimId, policyId ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.PolicyDelegate import PolicyDelegate

		errMsg = "Failed to assign element " + str(policyId) + " for Policy on Claim"

		try:
			# get the Claim from db
			claim = self.get( claimId ).first()	
			
			# get the Policy from db
			policy = PolicyDelegate().get(policyId).first();
			
			# assign the Policy		
			claim.policy = policy
			
			#save it
			claim.save()

			# reload and return the appropriate version					
			return self.get( claimId );
		except Claim.DoesNotExist:
			raise ProcessingError(errMsg + " : Claim with id " + str(claimId) + " does not exist.")
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPolicy( self, claimId ):
		errMsg = "Failed to unassign element " + str(policyId) + " for Policy on Claim"

		try:
			# get the Claim from db
			claim = self.get( claimId ).first()	
			
			# assign to None for unassignment
			claim.policy = None			

			#save it
			claim.save()

			# reload and return the appropriate version					
			return self.get( claimId );
		except Claim.DoesNotExist:
			raise ProcessingError(errMsg + " : Claim with id " + str(claimId) + " does not exist.")
		except Exception:
			return None;
		
	def assignCustomer( self, claimId, customerId ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.CustomerDelegate import CustomerDelegate

		errMsg = "Failed to assign element " + str(customerId) + " for Customer on Claim"

		try:
			# get the Claim from db
			claim = self.get( claimId ).first()	
			
			# get the Customer from db
			customer = CustomerDelegate().get(customerId).first();
			
			# assign the Customer		
			claim.customer = customer
			
			#save it
			claim.save()

			# reload and return the appropriate version					
			return self.get( claimId );
		except Claim.DoesNotExist:
			raise ProcessingError(errMsg + " : Claim with id " + str(claimId) + " does not exist.")
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCustomer( self, claimId ):
		errMsg = "Failed to unassign element " + str(customerId) + " for Customer on Claim"

		try:
			# get the Claim from db
			claim = self.get( claimId ).first()	
			
			# assign to None for unassignment
			claim.customer = None			

			#save it
			claim.save()

			# reload and return the appropriate version					
			return self.get( claimId );
		except Claim.DoesNotExist:
			raise ProcessingError(errMsg + " : Claim with id " + str(claimId) + " does not exist.")
		except Exception:
			return None;
		
	def assignAdjuster( self, claimId, adjusterId ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.AdjusterDelegate import AdjusterDelegate

		errMsg = "Failed to assign element " + str(adjusterId) + " for Adjuster on Claim"

		try:
			# get the Claim from db
			claim = self.get( claimId ).first()	
			
			# get the Adjuster from db
			adjuster = AdjusterDelegate().get(adjusterId).first();
			
			# assign the Adjuster		
			claim.adjuster = adjuster
			
			#save it
			claim.save()

			# reload and return the appropriate version					
			return self.get( claimId );
		except Claim.DoesNotExist:
			raise ProcessingError(errMsg + " : Claim with id " + str(claimId) + " does not exist.")
		except Adjuster.DoesNotExist:
			raise ProcessingError(errMsg + " : Adjuster with id " + str(adjusterId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignAdjuster( self, claimId ):
		errMsg = "Failed to unassign element " + str(adjusterId) + " for Adjuster on Claim"

		try:
			# get the Claim from db
			claim = self.get( claimId ).first()	
			
			# assign to None for unassignment
			claim.adjuster = None			

			#save it
			claim.save()

			# reload and return the appropriate version					
			return self.get( claimId );
		except Claim.DoesNotExist:
			raise ProcessingError(errMsg + " : Claim with id " + str(claimId) + " does not exist.")
		except Exception:
			return None;
		
	def assignIncident( self, claimId, incidentId ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.IncidentDelegate import IncidentDelegate

		errMsg = "Failed to assign element " + str(incidentId) + " for Incident on Claim"

		try:
			# get the Claim from db
			claim = self.get( claimId ).first()	
			
			# get the Incident from db
			incident = IncidentDelegate().get(incidentId).first();
			
			# assign the Incident		
			claim.incident = incident
			
			#save it
			claim.save()

			# reload and return the appropriate version					
			return self.get( claimId );
		except Claim.DoesNotExist:
			raise ProcessingError(errMsg + " : Claim with id " + str(claimId) + " does not exist.")
		except Incident.DoesNotExist:
			raise ProcessingError(errMsg + " : Incident with id " + str(incidentId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignIncident( self, claimId ):
		errMsg = "Failed to unassign element " + str(incidentId) + " for Incident on Claim"

		try:
			# get the Claim from db
			claim = self.get( claimId ).first()	
			
			# assign to None for unassignment
			claim.incident = None			

			#save it
			claim.save()

			# reload and return the appropriate version					
			return self.get( claimId );
		except Claim.DoesNotExist:
			raise ProcessingError(errMsg + " : Claim with id " + str(claimId) + " does not exist.")
		except Exception:
			return None;
		
	def addExposures( self, claimId, exposuresIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.ExposureDelegate import ExposureDelegate

		errMsg = "Failed to add elements " + str(exposuresIds) + " for Exposures on Claim"

		try:
			# get the Claim
			claim = self.get( claimId ).first()
				
			# split on a comma with no spaces
			idList = exposuresIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Exposure		
				exposure = ExposureDelegate().get(id).first();	
				# add the Exposure
				claim.exposures.add(exposure)
				
			# save it		
			claim.save()
			
			# reload and return the appropriate version
			return self.get( claimId );
		except Claim.DoesNotExist:
			raise ProcessingError(errMsg + " : Claim with id " + str(claimId) + " does not exist.")
		except Exposure.DoesNotExist:
			raise ProcessingError(errMsg + " : Exposure does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeExposures( self, claimId, exposuresIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.ExposureDelegate import ExposureDelegate

		errMsg = "Failed to remove elements " + str(exposuresIds) + " for Exposures on Claim"

		try:
			# get the Claim
			claim = self.get( claimId ).first()
				
			# split on a comma with no spaces
			idList = exposuresIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Exposure		
				exposure = ExposureDelegate().get(id).first();	
				# add the Exposure
				claim.exposures.remove(exposure)
				
			# save it		
			claim.save()
			
			# reload and return the appropriate version
			return self.get( claimId );
		except Claim.DoesNotExist:
			raise ProcessingError(errMsg + " : Claim with id " + str(claimId) + " does not exist.")
		except Exposure.DoesNotExist:
			raise ProcessingError(errMsg + " : Exposure does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addReserves( self, claimId, reservesIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.ClaimReserveDelegate import ClaimReserveDelegate

		errMsg = "Failed to add elements " + str(reservesIds) + " for Reserves on Claim"

		try:
			# get the Claim
			claim = self.get( claimId ).first()
				
			# split on a comma with no spaces
			idList = reservesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the ClaimReserve		
				claimReserve = ClaimReserveDelegate().get(id).first();	
				# add the ClaimReserve
				claim.reserves.add(claimReserve)
				
			# save it		
			claim.save()
			
			# reload and return the appropriate version
			return self.get( claimId );
		except Claim.DoesNotExist:
			raise ProcessingError(errMsg + " : Claim with id " + str(claimId) + " does not exist.")
		except ClaimReserve.DoesNotExist:
			raise ProcessingError(errMsg + " : ClaimReserve does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeReserves( self, claimId, reservesIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.ClaimReserveDelegate import ClaimReserveDelegate

		errMsg = "Failed to remove elements " + str(reservesIds) + " for Reserves on Claim"

		try:
			# get the Claim
			claim = self.get( claimId ).first()
				
			# split on a comma with no spaces
			idList = reservesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the ClaimReserve		
				claimReserve = ClaimReserveDelegate().get(id).first();	
				# add the ClaimReserve
				claim.reserves.remove(claimReserve)
				
			# save it		
			claim.save()
			
			# reload and return the appropriate version
			return self.get( claimId );
		except Claim.DoesNotExist:
			raise ProcessingError(errMsg + " : Claim with id " + str(claimId) + " does not exist.")
		except ClaimReserve.DoesNotExist:
			raise ProcessingError(errMsg + " : ClaimReserve does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addClaimPayments( self, claimId, claimPaymentsIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.ClaimPaymentDelegate import ClaimPaymentDelegate

		errMsg = "Failed to add elements " + str(claimPaymentsIds) + " for ClaimPayments on Claim"

		try:
			# get the Claim
			claim = self.get( claimId ).first()
				
			# split on a comma with no spaces
			idList = claimPaymentsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the ClaimPayment		
				claimPayment = ClaimPaymentDelegate().get(id).first();	
				# add the ClaimPayment
				claim.claimPayments.add(claimPayment)
				
			# save it		
			claim.save()
			
			# reload and return the appropriate version
			return self.get( claimId );
		except Claim.DoesNotExist:
			raise ProcessingError(errMsg + " : Claim with id " + str(claimId) + " does not exist.")
		except ClaimPayment.DoesNotExist:
			raise ProcessingError(errMsg + " : ClaimPayment does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeClaimPayments( self, claimId, claimPaymentsIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.ClaimPaymentDelegate import ClaimPaymentDelegate

		errMsg = "Failed to remove elements " + str(claimPaymentsIds) + " for ClaimPayments on Claim"

		try:
			# get the Claim
			claim = self.get( claimId ).first()
				
			# split on a comma with no spaces
			idList = claimPaymentsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the ClaimPayment		
				claimPayment = ClaimPaymentDelegate().get(id).first();	
				# add the ClaimPayment
				claim.claimPayments.remove(claimPayment)
				
			# save it		
			claim.save()
			
			# reload and return the appropriate version
			return self.get( claimId );
		except Claim.DoesNotExist:
			raise ProcessingError(errMsg + " : Claim with id " + str(claimId) + " does not exist.")
		except ClaimPayment.DoesNotExist:
			raise ProcessingError(errMsg + " : ClaimPayment does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addServiceProviders( self, claimId, serviceProvidersIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.ServiceProviderDelegate import ServiceProviderDelegate

		errMsg = "Failed to add elements " + str(serviceProvidersIds) + " for ServiceProviders on Claim"

		try:
			# get the Claim
			claim = self.get( claimId ).first()
				
			# split on a comma with no spaces
			idList = serviceProvidersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the ServiceProvider		
				serviceProvider = ServiceProviderDelegate().get(id).first();	
				# add the ServiceProvider
				claim.serviceProviders.add(serviceProvider)
				
			# save it		
			claim.save()
			
			# reload and return the appropriate version
			return self.get( claimId );
		except Claim.DoesNotExist:
			raise ProcessingError(errMsg + " : Claim with id " + str(claimId) + " does not exist.")
		except ServiceProvider.DoesNotExist:
			raise ProcessingError(errMsg + " : ServiceProvider does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeServiceProviders( self, claimId, serviceProvidersIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.ServiceProviderDelegate import ServiceProviderDelegate

		errMsg = "Failed to remove elements " + str(serviceProvidersIds) + " for ServiceProviders on Claim"

		try:
			# get the Claim
			claim = self.get( claimId ).first()
				
			# split on a comma with no spaces
			idList = serviceProvidersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the ServiceProvider		
				serviceProvider = ServiceProviderDelegate().get(id).first();	
				# add the ServiceProvider
				claim.serviceProviders.remove(serviceProvider)
				
			# save it		
			claim.save()
			
			# reload and return the appropriate version
			return self.get( claimId );
		except Claim.DoesNotExist:
			raise ProcessingError(errMsg + " : Claim with id " + str(claimId) + " does not exist.")
		except ServiceProvider.DoesNotExist:
			raise ProcessingError(errMsg + " : ServiceProvider does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addSubrogations( self, claimId, subrogationsIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.SubrogationRecoveryDelegate import SubrogationRecoveryDelegate

		errMsg = "Failed to add elements " + str(subrogationsIds) + " for Subrogations on Claim"

		try:
			# get the Claim
			claim = self.get( claimId ).first()
				
			# split on a comma with no spaces
			idList = subrogationsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the SubrogationRecovery		
				subrogationRecovery = SubrogationRecoveryDelegate().get(id).first();	
				# add the SubrogationRecovery
				claim.subrogations.add(subrogationRecovery)
				
			# save it		
			claim.save()
			
			# reload and return the appropriate version
			return self.get( claimId );
		except Claim.DoesNotExist:
			raise ProcessingError(errMsg + " : Claim with id " + str(claimId) + " does not exist.")
		except SubrogationRecovery.DoesNotExist:
			raise ProcessingError(errMsg + " : SubrogationRecovery does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeSubrogations( self, claimId, subrogationsIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.SubrogationRecoveryDelegate import SubrogationRecoveryDelegate

		errMsg = "Failed to remove elements " + str(subrogationsIds) + " for Subrogations on Claim"

		try:
			# get the Claim
			claim = self.get( claimId ).first()
				
			# split on a comma with no spaces
			idList = subrogationsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the SubrogationRecovery		
				subrogationRecovery = SubrogationRecoveryDelegate().get(id).first();	
				# add the SubrogationRecovery
				claim.subrogations.remove(subrogationRecovery)
				
			# save it		
			claim.save()
			
			# reload and return the appropriate version
			return self.get( claimId );
		except Claim.DoesNotExist:
			raise ProcessingError(errMsg + " : Claim with id " + str(claimId) + " does not exist.")
		except SubrogationRecovery.DoesNotExist:
			raise ProcessingError(errMsg + " : SubrogationRecovery does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
