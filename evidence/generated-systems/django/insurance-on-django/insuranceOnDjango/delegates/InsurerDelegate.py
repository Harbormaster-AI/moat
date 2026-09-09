from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from insuranceOnDjango.models.Insurer import Insurer
from insuranceOnDjango.models.InsuranceProduct import InsuranceProduct
from insuranceOnDjango.models.Distributor import Distributor
from insuranceOnDjango.models.Policy import Policy
from insuranceOnDjango.models.Claim import Claim
from insuranceOnDjango.models.ReinsuranceAgreement import ReinsuranceAgreement
from insuranceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Insurer
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InsurerDelegate Declaration
#======================================================================
class InsurerDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, insurerId ):
		try:	
			insurer = Insurer.objects.filter(id=insurerId)
			return insurer.first();
		except Insurer.DoesNotExist:
			raise ProcessingError("Insurer with id " + str(insurerId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, insurer):
		for model in serializers.deserialize("json", insurer):
			model.save()
			return model;

	def create(self, insurer):
		insurer.save()
		return insurer;

	def saveFromJson(self, insurer):
		for model in serializers.deserialize("json", insurer):
			model.save()
			return insurer;
	
	def save(self, insurer):
		insurer.save()
		return insurer;
	
	def delete(self, insurerId ):
		errMsg = "Failed to delete Insurer from db using id " + str(insurerId)
		
		try:
			insurer = Insurer.objects.get(id=insurerId)
			insurer.delete()
			return True
		except Insurer.DoesNotExist:
			raise ProcessingError("Insurer with id " + str(insurerId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Insurer.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Insurer from db")
		except Exception:
			return None;
		
	def addProducts( self, insurerId, productsIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.InsuranceProductDelegate import InsuranceProductDelegate

		errMsg = "Failed to add elements " + str(productsIds) + " for Products on Insurer"

		try:
			# get the Insurer
			insurer = self.get( insurerId ).first()
				
			# split on a comma with no spaces
			idList = productsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the InsuranceProduct		
				insuranceProduct = InsuranceProductDelegate().get(id).first();	
				# add the InsuranceProduct
				insurer.products.add(insuranceProduct)
				
			# save it		
			insurer.save()
			
			# reload and return the appropriate version
			return self.get( insurerId );
		except Insurer.DoesNotExist:
			raise ProcessingError(errMsg + " : Insurer with id " + str(insurerId) + " does not exist.")
		except InsuranceProduct.DoesNotExist:
			raise ProcessingError(errMsg + " : InsuranceProduct does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeProducts( self, insurerId, productsIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.InsuranceProductDelegate import InsuranceProductDelegate

		errMsg = "Failed to remove elements " + str(productsIds) + " for Products on Insurer"

		try:
			# get the Insurer
			insurer = self.get( insurerId ).first()
				
			# split on a comma with no spaces
			idList = productsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the InsuranceProduct		
				insuranceProduct = InsuranceProductDelegate().get(id).first();	
				# add the InsuranceProduct
				insurer.products.remove(insuranceProduct)
				
			# save it		
			insurer.save()
			
			# reload and return the appropriate version
			return self.get( insurerId );
		except Insurer.DoesNotExist:
			raise ProcessingError(errMsg + " : Insurer with id " + str(insurerId) + " does not exist.")
		except InsuranceProduct.DoesNotExist:
			raise ProcessingError(errMsg + " : InsuranceProduct does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addDistributionPartners( self, insurerId, distributionPartnersIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.DistributorDelegate import DistributorDelegate

		errMsg = "Failed to add elements " + str(distributionPartnersIds) + " for DistributionPartners on Insurer"

		try:
			# get the Insurer
			insurer = self.get( insurerId ).first()
				
			# split on a comma with no spaces
			idList = distributionPartnersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Distributor		
				distributor = DistributorDelegate().get(id).first();	
				# add the Distributor
				insurer.distributionPartners.add(distributor)
				
			# save it		
			insurer.save()
			
			# reload and return the appropriate version
			return self.get( insurerId );
		except Insurer.DoesNotExist:
			raise ProcessingError(errMsg + " : Insurer with id " + str(insurerId) + " does not exist.")
		except Distributor.DoesNotExist:
			raise ProcessingError(errMsg + " : Distributor does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDistributionPartners( self, insurerId, distributionPartnersIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.DistributorDelegate import DistributorDelegate

		errMsg = "Failed to remove elements " + str(distributionPartnersIds) + " for DistributionPartners on Insurer"

		try:
			# get the Insurer
			insurer = self.get( insurerId ).first()
				
			# split on a comma with no spaces
			idList = distributionPartnersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Distributor		
				distributor = DistributorDelegate().get(id).first();	
				# add the Distributor
				insurer.distributionPartners.remove(distributor)
				
			# save it		
			insurer.save()
			
			# reload and return the appropriate version
			return self.get( insurerId );
		except Insurer.DoesNotExist:
			raise ProcessingError(errMsg + " : Insurer with id " + str(insurerId) + " does not exist.")
		except Distributor.DoesNotExist:
			raise ProcessingError(errMsg + " : Distributor does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addPolicies( self, insurerId, policiesIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.PolicyDelegate import PolicyDelegate

		errMsg = "Failed to add elements " + str(policiesIds) + " for Policies on Insurer"

		try:
			# get the Insurer
			insurer = self.get( insurerId ).first()
				
			# split on a comma with no spaces
			idList = policiesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Policy		
				policy = PolicyDelegate().get(id).first();	
				# add the Policy
				insurer.policies.add(policy)
				
			# save it		
			insurer.save()
			
			# reload and return the appropriate version
			return self.get( insurerId );
		except Insurer.DoesNotExist:
			raise ProcessingError(errMsg + " : Insurer with id " + str(insurerId) + " does not exist.")
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePolicies( self, insurerId, policiesIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.PolicyDelegate import PolicyDelegate

		errMsg = "Failed to remove elements " + str(policiesIds) + " for Policies on Insurer"

		try:
			# get the Insurer
			insurer = self.get( insurerId ).first()
				
			# split on a comma with no spaces
			idList = policiesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Policy		
				policy = PolicyDelegate().get(id).first();	
				# add the Policy
				insurer.policies.remove(policy)
				
			# save it		
			insurer.save()
			
			# reload and return the appropriate version
			return self.get( insurerId );
		except Insurer.DoesNotExist:
			raise ProcessingError(errMsg + " : Insurer with id " + str(insurerId) + " does not exist.")
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addClaims( self, insurerId, claimsIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.ClaimDelegate import ClaimDelegate

		errMsg = "Failed to add elements " + str(claimsIds) + " for Claims on Insurer"

		try:
			# get the Insurer
			insurer = self.get( insurerId ).first()
				
			# split on a comma with no spaces
			idList = claimsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Claim		
				claim = ClaimDelegate().get(id).first();	
				# add the Claim
				insurer.claims.add(claim)
				
			# save it		
			insurer.save()
			
			# reload and return the appropriate version
			return self.get( insurerId );
		except Insurer.DoesNotExist:
			raise ProcessingError(errMsg + " : Insurer with id " + str(insurerId) + " does not exist.")
		except Claim.DoesNotExist:
			raise ProcessingError(errMsg + " : Claim does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeClaims( self, insurerId, claimsIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.ClaimDelegate import ClaimDelegate

		errMsg = "Failed to remove elements " + str(claimsIds) + " for Claims on Insurer"

		try:
			# get the Insurer
			insurer = self.get( insurerId ).first()
				
			# split on a comma with no spaces
			idList = claimsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Claim		
				claim = ClaimDelegate().get(id).first();	
				# add the Claim
				insurer.claims.remove(claim)
				
			# save it		
			insurer.save()
			
			# reload and return the appropriate version
			return self.get( insurerId );
		except Insurer.DoesNotExist:
			raise ProcessingError(errMsg + " : Insurer with id " + str(insurerId) + " does not exist.")
		except Claim.DoesNotExist:
			raise ProcessingError(errMsg + " : Claim does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addReinsuranceAgreements( self, insurerId, reinsuranceAgreementsIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.ReinsuranceAgreementDelegate import ReinsuranceAgreementDelegate

		errMsg = "Failed to add elements " + str(reinsuranceAgreementsIds) + " for ReinsuranceAgreements on Insurer"

		try:
			# get the Insurer
			insurer = self.get( insurerId ).first()
				
			# split on a comma with no spaces
			idList = reinsuranceAgreementsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the ReinsuranceAgreement		
				reinsuranceAgreement = ReinsuranceAgreementDelegate().get(id).first();	
				# add the ReinsuranceAgreement
				insurer.reinsuranceAgreements.add(reinsuranceAgreement)
				
			# save it		
			insurer.save()
			
			# reload and return the appropriate version
			return self.get( insurerId );
		except Insurer.DoesNotExist:
			raise ProcessingError(errMsg + " : Insurer with id " + str(insurerId) + " does not exist.")
		except ReinsuranceAgreement.DoesNotExist:
			raise ProcessingError(errMsg + " : ReinsuranceAgreement does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeReinsuranceAgreements( self, insurerId, reinsuranceAgreementsIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.ReinsuranceAgreementDelegate import ReinsuranceAgreementDelegate

		errMsg = "Failed to remove elements " + str(reinsuranceAgreementsIds) + " for ReinsuranceAgreements on Insurer"

		try:
			# get the Insurer
			insurer = self.get( insurerId ).first()
				
			# split on a comma with no spaces
			idList = reinsuranceAgreementsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the ReinsuranceAgreement		
				reinsuranceAgreement = ReinsuranceAgreementDelegate().get(id).first();	
				# add the ReinsuranceAgreement
				insurer.reinsuranceAgreements.remove(reinsuranceAgreement)
				
			# save it		
			insurer.save()
			
			# reload and return the appropriate version
			return self.get( insurerId );
		except Insurer.DoesNotExist:
			raise ProcessingError(errMsg + " : Insurer with id " + str(insurerId) + " does not exist.")
		except ReinsuranceAgreement.DoesNotExist:
			raise ProcessingError(errMsg + " : ReinsuranceAgreement does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
