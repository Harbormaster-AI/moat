from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from insuranceOnDjango.models.InsuredObject import InsuredObject
from insuranceOnDjango.models.Policy import Policy
from insuranceOnDjango.models.PolicyCoverage import PolicyCoverage
from insuranceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model InsuredObject
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InsuredObjectDelegate Declaration
#======================================================================
class InsuredObjectDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, insuredObjectId ):
		try:	
			insuredObject = InsuredObject.objects.filter(id=insuredObjectId)
			return insuredObject.first();
		except InsuredObject.DoesNotExist:
			raise ProcessingError("InsuredObject with id " + str(insuredObjectId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, insuredObject):
		for model in serializers.deserialize("json", insuredObject):
			model.save()
			return model;

	def create(self, insuredObject):
		insuredObject.save()
		return insuredObject;

	def saveFromJson(self, insuredObject):
		for model in serializers.deserialize("json", insuredObject):
			model.save()
			return insuredObject;
	
	def save(self, insuredObject):
		insuredObject.save()
		return insuredObject;
	
	def delete(self, insuredObjectId ):
		errMsg = "Failed to delete InsuredObject from db using id " + str(insuredObjectId)
		
		try:
			insuredObject = InsuredObject.objects.get(id=insuredObjectId)
			insuredObject.delete()
			return True
		except InsuredObject.DoesNotExist:
			raise ProcessingError("InsuredObject with id " + str(insuredObjectId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = InsuredObject.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all InsuredObject from db")
		except Exception:
			return None;
		
	def assignPolicy( self, insuredObjectId, policyId ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.PolicyDelegate import PolicyDelegate

		errMsg = "Failed to assign element " + str(policyId) + " for Policy on InsuredObject"

		try:
			# get the InsuredObject from db
			insuredObject = self.get( insuredObjectId ).first()	
			
			# get the Policy from db
			policy = PolicyDelegate().get(policyId).first();
			
			# assign the Policy		
			insuredObject.policy = policy
			
			#save it
			insuredObject.save()

			# reload and return the appropriate version					
			return self.get( insuredObjectId );
		except InsuredObject.DoesNotExist:
			raise ProcessingError(errMsg + " : InsuredObject with id " + str(insuredObjectId) + " does not exist.")
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPolicy( self, insuredObjectId ):
		errMsg = "Failed to unassign element " + str(policyId) + " for Policy on InsuredObject"

		try:
			# get the InsuredObject from db
			insuredObject = self.get( insuredObjectId ).first()	
			
			# assign to None for unassignment
			insuredObject.policy = None			

			#save it
			insuredObject.save()

			# reload and return the appropriate version					
			return self.get( insuredObjectId );
		except InsuredObject.DoesNotExist:
			raise ProcessingError(errMsg + " : InsuredObject with id " + str(insuredObjectId) + " does not exist.")
		except Exception:
			return None;
		
	def addCoverages( self, insuredObjectId, coveragesIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.PolicyCoverageDelegate import PolicyCoverageDelegate

		errMsg = "Failed to add elements " + str(coveragesIds) + " for Coverages on InsuredObject"

		try:
			# get the InsuredObject
			insuredObject = self.get( insuredObjectId ).first()
				
			# split on a comma with no spaces
			idList = coveragesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the PolicyCoverage		
				policyCoverage = PolicyCoverageDelegate().get(id).first();	
				# add the PolicyCoverage
				insuredObject.coverages.add(policyCoverage)
				
			# save it		
			insuredObject.save()
			
			# reload and return the appropriate version
			return self.get( insuredObjectId );
		except InsuredObject.DoesNotExist:
			raise ProcessingError(errMsg + " : InsuredObject with id " + str(insuredObjectId) + " does not exist.")
		except PolicyCoverage.DoesNotExist:
			raise ProcessingError(errMsg + " : PolicyCoverage does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCoverages( self, insuredObjectId, coveragesIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.PolicyCoverageDelegate import PolicyCoverageDelegate

		errMsg = "Failed to remove elements " + str(coveragesIds) + " for Coverages on InsuredObject"

		try:
			# get the InsuredObject
			insuredObject = self.get( insuredObjectId ).first()
				
			# split on a comma with no spaces
			idList = coveragesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the PolicyCoverage		
				policyCoverage = PolicyCoverageDelegate().get(id).first();	
				# add the PolicyCoverage
				insuredObject.coverages.remove(policyCoverage)
				
			# save it		
			insuredObject.save()
			
			# reload and return the appropriate version
			return self.get( insuredObjectId );
		except InsuredObject.DoesNotExist:
			raise ProcessingError(errMsg + " : InsuredObject with id " + str(insuredObjectId) + " does not exist.")
		except PolicyCoverage.DoesNotExist:
			raise ProcessingError(errMsg + " : PolicyCoverage does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
