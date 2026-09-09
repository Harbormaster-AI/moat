from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from healthcareOnDjango.models.InsurancePayer import InsurancePayer
from healthcareOnDjango.models.InsurancePlan import InsurancePlan
from healthcareOnDjango.models.Claim import Claim
from healthcareOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model InsurancePayer
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InsurancePayerDelegate Declaration
#======================================================================
class InsurancePayerDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, insurancePayerId ):
		try:	
			insurancePayer = InsurancePayer.objects.filter(id=insurancePayerId)
			return insurancePayer.first();
		except InsurancePayer.DoesNotExist:
			raise ProcessingError("InsurancePayer with id " + str(insurancePayerId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, insurancePayer):
		for model in serializers.deserialize("json", insurancePayer):
			model.save()
			return model;

	def create(self, insurancePayer):
		insurancePayer.save()
		return insurancePayer;

	def saveFromJson(self, insurancePayer):
		for model in serializers.deserialize("json", insurancePayer):
			model.save()
			return insurancePayer;
	
	def save(self, insurancePayer):
		insurancePayer.save()
		return insurancePayer;
	
	def delete(self, insurancePayerId ):
		errMsg = "Failed to delete InsurancePayer from db using id " + str(insurancePayerId)
		
		try:
			insurancePayer = InsurancePayer.objects.get(id=insurancePayerId)
			insurancePayer.delete()
			return True
		except InsurancePayer.DoesNotExist:
			raise ProcessingError("InsurancePayer with id " + str(insurancePayerId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = InsurancePayer.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all InsurancePayer from db")
		except Exception:
			return None;
		
	def addPlans( self, insurancePayerId, plansIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.InsurancePlanDelegate import InsurancePlanDelegate

		errMsg = "Failed to add elements " + str(plansIds) + " for Plans on InsurancePayer"

		try:
			# get the InsurancePayer
			insurancePayer = self.get( insurancePayerId ).first()
				
			# split on a comma with no spaces
			idList = plansIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the InsurancePlan		
				insurancePlan = InsurancePlanDelegate().get(id).first();	
				# add the InsurancePlan
				insurancePayer.plans.add(insurancePlan)
				
			# save it		
			insurancePayer.save()
			
			# reload and return the appropriate version
			return self.get( insurancePayerId );
		except InsurancePayer.DoesNotExist:
			raise ProcessingError(errMsg + " : InsurancePayer with id " + str(insurancePayerId) + " does not exist.")
		except InsurancePlan.DoesNotExist:
			raise ProcessingError(errMsg + " : InsurancePlan does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePlans( self, insurancePayerId, plansIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.InsurancePlanDelegate import InsurancePlanDelegate

		errMsg = "Failed to remove elements " + str(plansIds) + " for Plans on InsurancePayer"

		try:
			# get the InsurancePayer
			insurancePayer = self.get( insurancePayerId ).first()
				
			# split on a comma with no spaces
			idList = plansIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the InsurancePlan		
				insurancePlan = InsurancePlanDelegate().get(id).first();	
				# add the InsurancePlan
				insurancePayer.plans.remove(insurancePlan)
				
			# save it		
			insurancePayer.save()
			
			# reload and return the appropriate version
			return self.get( insurancePayerId );
		except InsurancePayer.DoesNotExist:
			raise ProcessingError(errMsg + " : InsurancePayer with id " + str(insurancePayerId) + " does not exist.")
		except InsurancePlan.DoesNotExist:
			raise ProcessingError(errMsg + " : InsurancePlan does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addClaims( self, insurancePayerId, claimsIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ClaimDelegate import ClaimDelegate

		errMsg = "Failed to add elements " + str(claimsIds) + " for Claims on InsurancePayer"

		try:
			# get the InsurancePayer
			insurancePayer = self.get( insurancePayerId ).first()
				
			# split on a comma with no spaces
			idList = claimsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Claim		
				claim = ClaimDelegate().get(id).first();	
				# add the Claim
				insurancePayer.claims.add(claim)
				
			# save it		
			insurancePayer.save()
			
			# reload and return the appropriate version
			return self.get( insurancePayerId );
		except InsurancePayer.DoesNotExist:
			raise ProcessingError(errMsg + " : InsurancePayer with id " + str(insurancePayerId) + " does not exist.")
		except Claim.DoesNotExist:
			raise ProcessingError(errMsg + " : Claim does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeClaims( self, insurancePayerId, claimsIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ClaimDelegate import ClaimDelegate

		errMsg = "Failed to remove elements " + str(claimsIds) + " for Claims on InsurancePayer"

		try:
			# get the InsurancePayer
			insurancePayer = self.get( insurancePayerId ).first()
				
			# split on a comma with no spaces
			idList = claimsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Claim		
				claim = ClaimDelegate().get(id).first();	
				# add the Claim
				insurancePayer.claims.remove(claim)
				
			# save it		
			insurancePayer.save()
			
			# reload and return the appropriate version
			return self.get( insurancePayerId );
		except InsurancePayer.DoesNotExist:
			raise ProcessingError(errMsg + " : InsurancePayer with id " + str(insurancePayerId) + " does not exist.")
		except Claim.DoesNotExist:
			raise ProcessingError(errMsg + " : Claim does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
