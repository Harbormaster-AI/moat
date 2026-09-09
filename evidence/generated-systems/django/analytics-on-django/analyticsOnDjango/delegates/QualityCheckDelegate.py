from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from analyticsOnDjango.models.QualityCheck import QualityCheck
from analyticsOnDjango.models.QualityRule import QualityRule
from analyticsOnDjango.models.DataSet import DataSet
from analyticsOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model QualityCheck
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class QualityCheckDelegate Declaration
#======================================================================
class QualityCheckDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, qualityCheckId ):
		try:	
			qualityCheck = QualityCheck.objects.filter(id=qualityCheckId)
			return qualityCheck.first();
		except QualityCheck.DoesNotExist:
			raise ProcessingError("QualityCheck with id " + str(qualityCheckId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, qualityCheck):
		for model in serializers.deserialize("json", qualityCheck):
			model.save()
			return model;

	def create(self, qualityCheck):
		qualityCheck.save()
		return qualityCheck;

	def saveFromJson(self, qualityCheck):
		for model in serializers.deserialize("json", qualityCheck):
			model.save()
			return qualityCheck;
	
	def save(self, qualityCheck):
		qualityCheck.save()
		return qualityCheck;
	
	def delete(self, qualityCheckId ):
		errMsg = "Failed to delete QualityCheck from db using id " + str(qualityCheckId)
		
		try:
			qualityCheck = QualityCheck.objects.get(id=qualityCheckId)
			qualityCheck.delete()
			return True
		except QualityCheck.DoesNotExist:
			raise ProcessingError("QualityCheck with id " + str(qualityCheckId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = QualityCheck.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all QualityCheck from db")
		except Exception:
			return None;
		
	def assignRule( self, qualityCheckId, ruleId ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.QualityRuleDelegate import QualityRuleDelegate

		errMsg = "Failed to assign element " + str(ruleId) + " for Rule on QualityCheck"

		try:
			# get the QualityCheck from db
			qualityCheck = self.get( qualityCheckId ).first()	
			
			# get the QualityRule from db
			qualityRule = QualityRuleDelegate().get(ruleId).first();
			
			# assign the Rule		
			qualityCheck.rule = qualityRule
			
			#save it
			qualityCheck.save()

			# reload and return the appropriate version					
			return self.get( qualityCheckId );
		except QualityCheck.DoesNotExist:
			raise ProcessingError(errMsg + " : QualityCheck with id " + str(qualityCheckId) + " does not exist.")
		except QualityRule.DoesNotExist:
			raise ProcessingError(errMsg + " : QualityRule with id " + str(ruleId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignRule( self, qualityCheckId ):
		errMsg = "Failed to unassign element " + str(ruleId) + " for Rule on QualityCheck"

		try:
			# get the QualityCheck from db
			qualityCheck = self.get( qualityCheckId ).first()	
			
			# assign to None for unassignment
			qualityCheck.qualityRule = None			

			#save it
			qualityCheck.save()

			# reload and return the appropriate version					
			return self.get( qualityCheckId );
		except QualityCheck.DoesNotExist:
			raise ProcessingError(errMsg + " : QualityCheck with id " + str(qualityCheckId) + " does not exist.")
		except Exception:
			return None;
		
	def assignDataset( self, qualityCheckId, datasetId ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to assign element " + str(datasetId) + " for Dataset on QualityCheck"

		try:
			# get the QualityCheck from db
			qualityCheck = self.get( qualityCheckId ).first()	
			
			# get the DataSet from db
			dataSet = DataSetDelegate().get(datasetId).first();
			
			# assign the Dataset		
			qualityCheck.dataset = dataSet
			
			#save it
			qualityCheck.save()

			# reload and return the appropriate version					
			return self.get( qualityCheckId );
		except QualityCheck.DoesNotExist:
			raise ProcessingError(errMsg + " : QualityCheck with id " + str(qualityCheckId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet with id " + str(datasetId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignDataset( self, qualityCheckId ):
		errMsg = "Failed to unassign element " + str(datasetId) + " for Dataset on QualityCheck"

		try:
			# get the QualityCheck from db
			qualityCheck = self.get( qualityCheckId ).first()	
			
			# assign to None for unassignment
			qualityCheck.dataSet = None			

			#save it
			qualityCheck.save()

			# reload and return the appropriate version					
			return self.get( qualityCheckId );
		except QualityCheck.DoesNotExist:
			raise ProcessingError(errMsg + " : QualityCheck with id " + str(qualityCheckId) + " does not exist.")
		except Exception:
			return None;
		
