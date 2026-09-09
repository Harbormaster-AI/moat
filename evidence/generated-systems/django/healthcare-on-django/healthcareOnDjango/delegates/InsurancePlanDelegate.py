from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from healthcareOnDjango.models.InsurancePlan import InsurancePlan
from healthcareOnDjango.models.InsurancePayer import InsurancePayer
from healthcareOnDjango.models.Coverage import Coverage
from healthcareOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model InsurancePlan
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InsurancePlanDelegate Declaration
#======================================================================
class InsurancePlanDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, insurancePlanId ):
		try:	
			insurancePlan = InsurancePlan.objects.filter(id=insurancePlanId)
			return insurancePlan.first();
		except InsurancePlan.DoesNotExist:
			raise ProcessingError("InsurancePlan with id " + str(insurancePlanId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, insurancePlan):
		for model in serializers.deserialize("json", insurancePlan):
			model.save()
			return model;

	def create(self, insurancePlan):
		insurancePlan.save()
		return insurancePlan;

	def saveFromJson(self, insurancePlan):
		for model in serializers.deserialize("json", insurancePlan):
			model.save()
			return insurancePlan;
	
	def save(self, insurancePlan):
		insurancePlan.save()
		return insurancePlan;
	
	def delete(self, insurancePlanId ):
		errMsg = "Failed to delete InsurancePlan from db using id " + str(insurancePlanId)
		
		try:
			insurancePlan = InsurancePlan.objects.get(id=insurancePlanId)
			insurancePlan.delete()
			return True
		except InsurancePlan.DoesNotExist:
			raise ProcessingError("InsurancePlan with id " + str(insurancePlanId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = InsurancePlan.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all InsurancePlan from db")
		except Exception:
			return None;
		
	def assignPayer( self, insurancePlanId, payerId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.InsurancePayerDelegate import InsurancePayerDelegate

		errMsg = "Failed to assign element " + str(payerId) + " for Payer on InsurancePlan"

		try:
			# get the InsurancePlan from db
			insurancePlan = self.get( insurancePlanId ).first()	
			
			# get the InsurancePayer from db
			insurancePayer = InsurancePayerDelegate().get(payerId).first();
			
			# assign the Payer		
			insurancePlan.payer = insurancePayer
			
			#save it
			insurancePlan.save()

			# reload and return the appropriate version					
			return self.get( insurancePlanId );
		except InsurancePlan.DoesNotExist:
			raise ProcessingError(errMsg + " : InsurancePlan with id " + str(insurancePlanId) + " does not exist.")
		except InsurancePayer.DoesNotExist:
			raise ProcessingError(errMsg + " : InsurancePayer with id " + str(payerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPayer( self, insurancePlanId ):
		errMsg = "Failed to unassign element " + str(payerId) + " for Payer on InsurancePlan"

		try:
			# get the InsurancePlan from db
			insurancePlan = self.get( insurancePlanId ).first()	
			
			# assign to None for unassignment
			insurancePlan.insurancePayer = None			

			#save it
			insurancePlan.save()

			# reload and return the appropriate version					
			return self.get( insurancePlanId );
		except InsurancePlan.DoesNotExist:
			raise ProcessingError(errMsg + " : InsurancePlan with id " + str(insurancePlanId) + " does not exist.")
		except Exception:
			return None;
		
	def addCoverages( self, insurancePlanId, coveragesIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.CoverageDelegate import CoverageDelegate

		errMsg = "Failed to add elements " + str(coveragesIds) + " for Coverages on InsurancePlan"

		try:
			# get the InsurancePlan
			insurancePlan = self.get( insurancePlanId ).first()
				
			# split on a comma with no spaces
			idList = coveragesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Coverage		
				coverage = CoverageDelegate().get(id).first();	
				# add the Coverage
				insurancePlan.coverages.add(coverage)
				
			# save it		
			insurancePlan.save()
			
			# reload and return the appropriate version
			return self.get( insurancePlanId );
		except InsurancePlan.DoesNotExist:
			raise ProcessingError(errMsg + " : InsurancePlan with id " + str(insurancePlanId) + " does not exist.")
		except Coverage.DoesNotExist:
			raise ProcessingError(errMsg + " : Coverage does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCoverages( self, insurancePlanId, coveragesIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.CoverageDelegate import CoverageDelegate

		errMsg = "Failed to remove elements " + str(coveragesIds) + " for Coverages on InsurancePlan"

		try:
			# get the InsurancePlan
			insurancePlan = self.get( insurancePlanId ).first()
				
			# split on a comma with no spaces
			idList = coveragesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Coverage		
				coverage = CoverageDelegate().get(id).first();	
				# add the Coverage
				insurancePlan.coverages.remove(coverage)
				
			# save it		
			insurancePlan.save()
			
			# reload and return the appropriate version
			return self.get( insurancePlanId );
		except InsurancePlan.DoesNotExist:
			raise ProcessingError(errMsg + " : InsurancePlan with id " + str(insurancePlanId) + " does not exist.")
		except Coverage.DoesNotExist:
			raise ProcessingError(errMsg + " : Coverage does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
