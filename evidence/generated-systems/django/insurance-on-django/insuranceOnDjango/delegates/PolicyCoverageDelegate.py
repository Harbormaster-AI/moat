from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from insuranceOnDjango.models.PolicyCoverage import PolicyCoverage
from insuranceOnDjango.models.Policy import Policy
from insuranceOnDjango.models.InsuredObject import InsuredObject
from insuranceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model PolicyCoverage
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PolicyCoverageDelegate Declaration
#======================================================================
class PolicyCoverageDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, policyCoverageId ):
		try:	
			policyCoverage = PolicyCoverage.objects.filter(id=policyCoverageId)
			return policyCoverage.first();
		except PolicyCoverage.DoesNotExist:
			raise ProcessingError("PolicyCoverage with id " + str(policyCoverageId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, policyCoverage):
		for model in serializers.deserialize("json", policyCoverage):
			model.save()
			return model;

	def create(self, policyCoverage):
		policyCoverage.save()
		return policyCoverage;

	def saveFromJson(self, policyCoverage):
		for model in serializers.deserialize("json", policyCoverage):
			model.save()
			return policyCoverage;
	
	def save(self, policyCoverage):
		policyCoverage.save()
		return policyCoverage;
	
	def delete(self, policyCoverageId ):
		errMsg = "Failed to delete PolicyCoverage from db using id " + str(policyCoverageId)
		
		try:
			policyCoverage = PolicyCoverage.objects.get(id=policyCoverageId)
			policyCoverage.delete()
			return True
		except PolicyCoverage.DoesNotExist:
			raise ProcessingError("PolicyCoverage with id " + str(policyCoverageId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = PolicyCoverage.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all PolicyCoverage from db")
		except Exception:
			return None;
		
	def assignPolicy( self, policyCoverageId, policyId ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.PolicyDelegate import PolicyDelegate

		errMsg = "Failed to assign element " + str(policyId) + " for Policy on PolicyCoverage"

		try:
			# get the PolicyCoverage from db
			policyCoverage = self.get( policyCoverageId ).first()	
			
			# get the Policy from db
			policy = PolicyDelegate().get(policyId).first();
			
			# assign the Policy		
			policyCoverage.policy = policy
			
			#save it
			policyCoverage.save()

			# reload and return the appropriate version					
			return self.get( policyCoverageId );
		except PolicyCoverage.DoesNotExist:
			raise ProcessingError(errMsg + " : PolicyCoverage with id " + str(policyCoverageId) + " does not exist.")
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPolicy( self, policyCoverageId ):
		errMsg = "Failed to unassign element " + str(policyId) + " for Policy on PolicyCoverage"

		try:
			# get the PolicyCoverage from db
			policyCoverage = self.get( policyCoverageId ).first()	
			
			# assign to None for unassignment
			policyCoverage.policy = None			

			#save it
			policyCoverage.save()

			# reload and return the appropriate version					
			return self.get( policyCoverageId );
		except PolicyCoverage.DoesNotExist:
			raise ProcessingError(errMsg + " : PolicyCoverage with id " + str(policyCoverageId) + " does not exist.")
		except Exception:
			return None;
		
	def addInsuredObjects( self, policyCoverageId, insuredObjectsIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.InsuredObjectDelegate import InsuredObjectDelegate

		errMsg = "Failed to add elements " + str(insuredObjectsIds) + " for InsuredObjects on PolicyCoverage"

		try:
			# get the PolicyCoverage
			policyCoverage = self.get( policyCoverageId ).first()
				
			# split on a comma with no spaces
			idList = insuredObjectsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the InsuredObject		
				insuredObject = InsuredObjectDelegate().get(id).first();	
				# add the InsuredObject
				policyCoverage.insuredObjects.add(insuredObject)
				
			# save it		
			policyCoverage.save()
			
			# reload and return the appropriate version
			return self.get( policyCoverageId );
		except PolicyCoverage.DoesNotExist:
			raise ProcessingError(errMsg + " : PolicyCoverage with id " + str(policyCoverageId) + " does not exist.")
		except InsuredObject.DoesNotExist:
			raise ProcessingError(errMsg + " : InsuredObject does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeInsuredObjects( self, policyCoverageId, insuredObjectsIds ):
		# lazy importing avoids circular dependencies
		from insuranceOnDjango.delegates.InsuredObjectDelegate import InsuredObjectDelegate

		errMsg = "Failed to remove elements " + str(insuredObjectsIds) + " for InsuredObjects on PolicyCoverage"

		try:
			# get the PolicyCoverage
			policyCoverage = self.get( policyCoverageId ).first()
				
			# split on a comma with no spaces
			idList = insuredObjectsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the InsuredObject		
				insuredObject = InsuredObjectDelegate().get(id).first();	
				# add the InsuredObject
				policyCoverage.insuredObjects.remove(insuredObject)
				
			# save it		
			policyCoverage.save()
			
			# reload and return the appropriate version
			return self.get( policyCoverageId );
		except PolicyCoverage.DoesNotExist:
			raise ProcessingError(errMsg + " : PolicyCoverage with id " + str(policyCoverageId) + " does not exist.")
		except InsuredObject.DoesNotExist:
			raise ProcessingError(errMsg + " : InsuredObject does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
