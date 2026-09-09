from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from insuranceOnDjango.models.Exposure import Exposure
from insuranceOnDjango.models.Claim import Claim
from insuranceOnDjango.models.PolicyCoverage import PolicyCoverage
from insuranceOnDjango.models.InsuredObject import InsuredObject
from insuranceOnDjango.models.ClaimReserve import ClaimReserve
from insuranceOnDjango.models.ClaimPayment import ClaimPayment
from insuranceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Exposure
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ExposureDelegate Declaration
#======================================================================
class ExposureDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, exposureId ):
		try:	
			exposure = Exposure.objects.filter(id=exposureId)
			return exposure.first();
		except Exposure.DoesNotExist:
			raise ProcessingError("Exposure with id " + str(exposureId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, exposure):
		for model in serializers.deserialize("json", exposure):
			model.save()
			return model;

	def create(self, exposure):
		exposure.save()
		return exposure;

	def saveFromJson(self, exposure):
		for model in serializers.deserialize("json", exposure):
			model.save()
			return exposure;
	
	def save(self, exposure):
		exposure.save()
		return exposure;
	
	def delete(self, exposureId ):
		errMsg = "Failed to delete Exposure from db using id " + str(exposureId)
		
		try:
			exposure = Exposure.objects.get(id=exposureId)
			exposure.delete()
			return True
		except Exposure.DoesNotExist:
			raise ProcessingError("Exposure with id " + str(exposureId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Exposure.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Exposure from db")
		except Exception:
			return None;
		
	def assignClaim( self, exposureId, claimId ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.ClaimDelegate import ClaimDelegate

		errMsg = "Failed to assign element " + str(claimId) + " for Claim on Exposure"

		try:
			# get the Exposure from db
			exposure = self.get( exposureId ).first()	
			
			# get the Claim from db
			claim = ClaimDelegate().get(claimId).first();
			
			# assign the Claim		
			exposure.claim = claim
			
			#save it
			exposure.save()

			# reload and return the appropriate version					
			return self.get( exposureId );
		except Exposure.DoesNotExist:
			raise ProcessingError(errMsg + " : Exposure with id " + str(exposureId) + " does not exist.")
		except Claim.DoesNotExist:
			raise ProcessingError(errMsg + " : Claim with id " + str(claimId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignClaim( self, exposureId ):
		errMsg = "Failed to unassign element " + str(claimId) + " for Claim on Exposure"

		try:
			# get the Exposure from db
			exposure = self.get( exposureId ).first()	
			
			# assign to None for unassignment
			exposure.claim = None			

			#save it
			exposure.save()

			# reload and return the appropriate version					
			return self.get( exposureId );
		except Exposure.DoesNotExist:
			raise ProcessingError(errMsg + " : Exposure with id " + str(exposureId) + " does not exist.")
		except Exception:
			return None;
		
	def assignPolicyCoverage( self, exposureId, policyCoverageId ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.PolicyCoverageDelegate import PolicyCoverageDelegate

		errMsg = "Failed to assign element " + str(policyCoverageId) + " for PolicyCoverage on Exposure"

		try:
			# get the Exposure from db
			exposure = self.get( exposureId ).first()	
			
			# get the PolicyCoverage from db
			policyCoverage = PolicyCoverageDelegate().get(policyCoverageId).first();
			
			# assign the PolicyCoverage		
			exposure.policyCoverage = policyCoverage
			
			#save it
			exposure.save()

			# reload and return the appropriate version					
			return self.get( exposureId );
		except Exposure.DoesNotExist:
			raise ProcessingError(errMsg + " : Exposure with id " + str(exposureId) + " does not exist.")
		except PolicyCoverage.DoesNotExist:
			raise ProcessingError(errMsg + " : PolicyCoverage with id " + str(policyCoverageId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPolicyCoverage( self, exposureId ):
		errMsg = "Failed to unassign element " + str(policyCoverageId) + " for PolicyCoverage on Exposure"

		try:
			# get the Exposure from db
			exposure = self.get( exposureId ).first()	
			
			# assign to None for unassignment
			exposure.policyCoverage = None			

			#save it
			exposure.save()

			# reload and return the appropriate version					
			return self.get( exposureId );
		except Exposure.DoesNotExist:
			raise ProcessingError(errMsg + " : Exposure with id " + str(exposureId) + " does not exist.")
		except Exception:
			return None;
		
	def assignInsuredObject( self, exposureId, insuredObjectId ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.InsuredObjectDelegate import InsuredObjectDelegate

		errMsg = "Failed to assign element " + str(insuredObjectId) + " for InsuredObject on Exposure"

		try:
			# get the Exposure from db
			exposure = self.get( exposureId ).first()	
			
			# get the InsuredObject from db
			insuredObject = InsuredObjectDelegate().get(insuredObjectId).first();
			
			# assign the InsuredObject		
			exposure.insuredObject = insuredObject
			
			#save it
			exposure.save()

			# reload and return the appropriate version					
			return self.get( exposureId );
		except Exposure.DoesNotExist:
			raise ProcessingError(errMsg + " : Exposure with id " + str(exposureId) + " does not exist.")
		except InsuredObject.DoesNotExist:
			raise ProcessingError(errMsg + " : InsuredObject with id " + str(insuredObjectId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignInsuredObject( self, exposureId ):
		errMsg = "Failed to unassign element " + str(insuredObjectId) + " for InsuredObject on Exposure"

		try:
			# get the Exposure from db
			exposure = self.get( exposureId ).first()	
			
			# assign to None for unassignment
			exposure.insuredObject = None			

			#save it
			exposure.save()

			# reload and return the appropriate version					
			return self.get( exposureId );
		except Exposure.DoesNotExist:
			raise ProcessingError(errMsg + " : Exposure with id " + str(exposureId) + " does not exist.")
		except Exception:
			return None;
		
	def addReserves( self, exposureId, reservesIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.ClaimReserveDelegate import ClaimReserveDelegate

		errMsg = "Failed to add elements " + str(reservesIds) + " for Reserves on Exposure"

		try:
			# get the Exposure
			exposure = self.get( exposureId ).first()
				
			# split on a comma with no spaces
			idList = reservesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the ClaimReserve		
				claimReserve = ClaimReserveDelegate().get(id).first();	
				# add the ClaimReserve
				exposure.reserves.add(claimReserve)
				
			# save it		
			exposure.save()
			
			# reload and return the appropriate version
			return self.get( exposureId );
		except Exposure.DoesNotExist:
			raise ProcessingError(errMsg + " : Exposure with id " + str(exposureId) + " does not exist.")
		except ClaimReserve.DoesNotExist:
			raise ProcessingError(errMsg + " : ClaimReserve does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeReserves( self, exposureId, reservesIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.ClaimReserveDelegate import ClaimReserveDelegate

		errMsg = "Failed to remove elements " + str(reservesIds) + " for Reserves on Exposure"

		try:
			# get the Exposure
			exposure = self.get( exposureId ).first()
				
			# split on a comma with no spaces
			idList = reservesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the ClaimReserve		
				claimReserve = ClaimReserveDelegate().get(id).first();	
				# add the ClaimReserve
				exposure.reserves.remove(claimReserve)
				
			# save it		
			exposure.save()
			
			# reload and return the appropriate version
			return self.get( exposureId );
		except Exposure.DoesNotExist:
			raise ProcessingError(errMsg + " : Exposure with id " + str(exposureId) + " does not exist.")
		except ClaimReserve.DoesNotExist:
			raise ProcessingError(errMsg + " : ClaimReserve does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addPayments( self, exposureId, paymentsIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.ClaimPaymentDelegate import ClaimPaymentDelegate

		errMsg = "Failed to add elements " + str(paymentsIds) + " for Payments on Exposure"

		try:
			# get the Exposure
			exposure = self.get( exposureId ).first()
				
			# split on a comma with no spaces
			idList = paymentsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the ClaimPayment		
				claimPayment = ClaimPaymentDelegate().get(id).first();	
				# add the ClaimPayment
				exposure.payments.add(claimPayment)
				
			# save it		
			exposure.save()
			
			# reload and return the appropriate version
			return self.get( exposureId );
		except Exposure.DoesNotExist:
			raise ProcessingError(errMsg + " : Exposure with id " + str(exposureId) + " does not exist.")
		except ClaimPayment.DoesNotExist:
			raise ProcessingError(errMsg + " : ClaimPayment does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePayments( self, exposureId, paymentsIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.ClaimPaymentDelegate import ClaimPaymentDelegate

		errMsg = "Failed to remove elements " + str(paymentsIds) + " for Payments on Exposure"

		try:
			# get the Exposure
			exposure = self.get( exposureId ).first()
				
			# split on a comma with no spaces
			idList = paymentsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the ClaimPayment		
				claimPayment = ClaimPaymentDelegate().get(id).first();	
				# add the ClaimPayment
				exposure.payments.remove(claimPayment)
				
			# save it		
			exposure.save()
			
			# reload and return the appropriate version
			return self.get( exposureId );
		except Exposure.DoesNotExist:
			raise ProcessingError(errMsg + " : Exposure with id " + str(exposureId) + " does not exist.")
		except ClaimPayment.DoesNotExist:
			raise ProcessingError(errMsg + " : ClaimPayment does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
