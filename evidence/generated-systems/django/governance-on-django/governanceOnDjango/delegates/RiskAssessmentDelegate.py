from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from governanceOnDjango.models.RiskAssessment import RiskAssessment
from governanceOnDjango.models.Risk import Risk
from governanceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model RiskAssessment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RiskAssessmentDelegate Declaration
#======================================================================
class RiskAssessmentDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, riskAssessmentId ):
		try:	
			riskAssessment = RiskAssessment.objects.filter(id=riskAssessmentId)
			return riskAssessment.first();
		except RiskAssessment.DoesNotExist:
			raise ProcessingError("RiskAssessment with id " + str(riskAssessmentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, riskAssessment):
		for model in serializers.deserialize("json", riskAssessment):
			model.save()
			return model;

	def create(self, riskAssessment):
		riskAssessment.save()
		return riskAssessment;

	def saveFromJson(self, riskAssessment):
		for model in serializers.deserialize("json", riskAssessment):
			model.save()
			return riskAssessment;
	
	def save(self, riskAssessment):
		riskAssessment.save()
		return riskAssessment;
	
	def delete(self, riskAssessmentId ):
		errMsg = "Failed to delete RiskAssessment from db using id " + str(riskAssessmentId)
		
		try:
			riskAssessment = RiskAssessment.objects.get(id=riskAssessmentId)
			riskAssessment.delete()
			return True
		except RiskAssessment.DoesNotExist:
			raise ProcessingError("RiskAssessment with id " + str(riskAssessmentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = RiskAssessment.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all RiskAssessment from db")
		except Exception:
			return None;
		
	def assignRisk( self, riskAssessmentId, riskId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.RiskDelegate import RiskDelegate

		errMsg = "Failed to assign element " + str(riskId) + " for Risk on RiskAssessment"

		try:
			# get the RiskAssessment from db
			riskAssessment = self.get( riskAssessmentId ).first()	
			
			# get the Risk from db
			risk = RiskDelegate().get(riskId).first();
			
			# assign the Risk		
			riskAssessment.risk = risk
			
			#save it
			riskAssessment.save()

			# reload and return the appropriate version					
			return self.get( riskAssessmentId );
		except RiskAssessment.DoesNotExist:
			raise ProcessingError(errMsg + " : RiskAssessment with id " + str(riskAssessmentId) + " does not exist.")
		except Risk.DoesNotExist:
			raise ProcessingError(errMsg + " : Risk with id " + str(riskId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignRisk( self, riskAssessmentId ):
		errMsg = "Failed to unassign element " + str(riskId) + " for Risk on RiskAssessment"

		try:
			# get the RiskAssessment from db
			riskAssessment = self.get( riskAssessmentId ).first()	
			
			# assign to None for unassignment
			riskAssessment.risk = None			

			#save it
			riskAssessment.save()

			# reload and return the appropriate version					
			return self.get( riskAssessmentId );
		except RiskAssessment.DoesNotExist:
			raise ProcessingError(errMsg + " : RiskAssessment with id " + str(riskAssessmentId) + " does not exist.")
		except Exception:
			return None;
		
