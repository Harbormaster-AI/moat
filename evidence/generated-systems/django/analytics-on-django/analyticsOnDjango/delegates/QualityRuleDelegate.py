from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from analyticsOnDjango.models.QualityRule import QualityRule
from analyticsOnDjango.models.DataSet import DataSet
from analyticsOnDjango.models.QualityCheck import QualityCheck
from analyticsOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model QualityRule
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class QualityRuleDelegate Declaration
#======================================================================
class QualityRuleDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, qualityRuleId ):
		try:	
			qualityRule = QualityRule.objects.filter(id=qualityRuleId)
			return qualityRule.first();
		except QualityRule.DoesNotExist:
			raise ProcessingError("QualityRule with id " + str(qualityRuleId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, qualityRule):
		for model in serializers.deserialize("json", qualityRule):
			model.save()
			return model;

	def create(self, qualityRule):
		qualityRule.save()
		return qualityRule;

	def saveFromJson(self, qualityRule):
		for model in serializers.deserialize("json", qualityRule):
			model.save()
			return qualityRule;
	
	def save(self, qualityRule):
		qualityRule.save()
		return qualityRule;
	
	def delete(self, qualityRuleId ):
		errMsg = "Failed to delete QualityRule from db using id " + str(qualityRuleId)
		
		try:
			qualityRule = QualityRule.objects.get(id=qualityRuleId)
			qualityRule.delete()
			return True
		except QualityRule.DoesNotExist:
			raise ProcessingError("QualityRule with id " + str(qualityRuleId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = QualityRule.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all QualityRule from db")
		except Exception:
			return None;
		
	def assignDataset( self, qualityRuleId, datasetId ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to assign element " + str(datasetId) + " for Dataset on QualityRule"

		try:
			# get the QualityRule from db
			qualityRule = self.get( qualityRuleId ).first()	
			
			# get the DataSet from db
			dataSet = DataSetDelegate().get(datasetId).first();
			
			# assign the Dataset		
			qualityRule.dataset = dataSet
			
			#save it
			qualityRule.save()

			# reload and return the appropriate version					
			return self.get( qualityRuleId );
		except QualityRule.DoesNotExist:
			raise ProcessingError(errMsg + " : QualityRule with id " + str(qualityRuleId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet with id " + str(datasetId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignDataset( self, qualityRuleId ):
		errMsg = "Failed to unassign element " + str(datasetId) + " for Dataset on QualityRule"

		try:
			# get the QualityRule from db
			qualityRule = self.get( qualityRuleId ).first()	
			
			# assign to None for unassignment
			qualityRule.dataSet = None			

			#save it
			qualityRule.save()

			# reload and return the appropriate version					
			return self.get( qualityRuleId );
		except QualityRule.DoesNotExist:
			raise ProcessingError(errMsg + " : QualityRule with id " + str(qualityRuleId) + " does not exist.")
		except Exception:
			return None;
		
	def addChecks( self, qualityRuleId, checksIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.QualityCheckDelegate import QualityCheckDelegate

		errMsg = "Failed to add elements " + str(checksIds) + " for Checks on QualityRule"

		try:
			# get the QualityRule
			qualityRule = self.get( qualityRuleId ).first()
				
			# split on a comma with no spaces
			idList = checksIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the QualityCheck		
				qualityCheck = QualityCheckDelegate().get(id).first();	
				# add the QualityCheck
				qualityRule.checks.add(qualityCheck)
				
			# save it		
			qualityRule.save()
			
			# reload and return the appropriate version
			return self.get( qualityRuleId );
		except QualityRule.DoesNotExist:
			raise ProcessingError(errMsg + " : QualityRule with id " + str(qualityRuleId) + " does not exist.")
		except QualityCheck.DoesNotExist:
			raise ProcessingError(errMsg + " : QualityCheck does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeChecks( self, qualityRuleId, checksIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.QualityCheckDelegate import QualityCheckDelegate

		errMsg = "Failed to remove elements " + str(checksIds) + " for Checks on QualityRule"

		try:
			# get the QualityRule
			qualityRule = self.get( qualityRuleId ).first()
				
			# split on a comma with no spaces
			idList = checksIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the QualityCheck		
				qualityCheck = QualityCheckDelegate().get(id).first();	
				# add the QualityCheck
				qualityRule.checks.remove(qualityCheck)
				
			# save it		
			qualityRule.save()
			
			# reload and return the appropriate version
			return self.get( qualityRuleId );
		except QualityRule.DoesNotExist:
			raise ProcessingError(errMsg + " : QualityRule with id " + str(qualityRuleId) + " does not exist.")
		except QualityCheck.DoesNotExist:
			raise ProcessingError(errMsg + " : QualityCheck does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
