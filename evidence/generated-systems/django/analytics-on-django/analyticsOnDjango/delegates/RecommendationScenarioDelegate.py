from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from analyticsOnDjango.models.RecommendationScenario import RecommendationScenario
from analyticsOnDjango.models.Model import Model
from analyticsOnDjango.models.DataSet import DataSet
from analyticsOnDjango.models.Experiment import Experiment
from analyticsOnDjango.models.Alert import Alert
from analyticsOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model RecommendationScenario
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RecommendationScenarioDelegate Declaration
#======================================================================
class RecommendationScenarioDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, recommendationScenarioId ):
		try:	
			recommendationScenario = RecommendationScenario.objects.filter(id=recommendationScenarioId)
			return recommendationScenario.first();
		except RecommendationScenario.DoesNotExist:
			raise ProcessingError("RecommendationScenario with id " + str(recommendationScenarioId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, recommendationScenario):
		for model in serializers.deserialize("json", recommendationScenario):
			model.save()
			return model;

	def create(self, recommendationScenario):
		recommendationScenario.save()
		return recommendationScenario;

	def saveFromJson(self, recommendationScenario):
		for model in serializers.deserialize("json", recommendationScenario):
			model.save()
			return recommendationScenario;
	
	def save(self, recommendationScenario):
		recommendationScenario.save()
		return recommendationScenario;
	
	def delete(self, recommendationScenarioId ):
		errMsg = "Failed to delete RecommendationScenario from db using id " + str(recommendationScenarioId)
		
		try:
			recommendationScenario = RecommendationScenario.objects.get(id=recommendationScenarioId)
			recommendationScenario.delete()
			return True
		except RecommendationScenario.DoesNotExist:
			raise ProcessingError("RecommendationScenario with id " + str(recommendationScenarioId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = RecommendationScenario.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all RecommendationScenario from db")
		except Exception:
			return None;
		
	def addModels( self, recommendationScenarioId, modelsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.ModelDelegate import ModelDelegate

		errMsg = "Failed to add elements " + str(modelsIds) + " for Models on RecommendationScenario"

		try:
			# get the RecommendationScenario
			recommendationScenario = self.get( recommendationScenarioId ).first()
				
			# split on a comma with no spaces
			idList = modelsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Model		
				model = ModelDelegate().get(id).first();	
				# add the Model
				recommendationScenario.models.add(model)
				
			# save it		
			recommendationScenario.save()
			
			# reload and return the appropriate version
			return self.get( recommendationScenarioId );
		except RecommendationScenario.DoesNotExist:
			raise ProcessingError(errMsg + " : RecommendationScenario with id " + str(recommendationScenarioId) + " does not exist.")
		except Model.DoesNotExist:
			raise ProcessingError(errMsg + " : Model does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeModels( self, recommendationScenarioId, modelsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.ModelDelegate import ModelDelegate

		errMsg = "Failed to remove elements " + str(modelsIds) + " for Models on RecommendationScenario"

		try:
			# get the RecommendationScenario
			recommendationScenario = self.get( recommendationScenarioId ).first()
				
			# split on a comma with no spaces
			idList = modelsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Model		
				model = ModelDelegate().get(id).first();	
				# add the Model
				recommendationScenario.models.remove(model)
				
			# save it		
			recommendationScenario.save()
			
			# reload and return the appropriate version
			return self.get( recommendationScenarioId );
		except RecommendationScenario.DoesNotExist:
			raise ProcessingError(errMsg + " : RecommendationScenario with id " + str(recommendationScenarioId) + " does not exist.")
		except Model.DoesNotExist:
			raise ProcessingError(errMsg + " : Model does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addDatasets( self, recommendationScenarioId, datasetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to add elements " + str(datasetsIds) + " for Datasets on RecommendationScenario"

		try:
			# get the RecommendationScenario
			recommendationScenario = self.get( recommendationScenarioId ).first()
				
			# split on a comma with no spaces
			idList = datasetsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataSet		
				dataSet = DataSetDelegate().get(id).first();	
				# add the DataSet
				recommendationScenario.datasets.add(dataSet)
				
			# save it		
			recommendationScenario.save()
			
			# reload and return the appropriate version
			return self.get( recommendationScenarioId );
		except RecommendationScenario.DoesNotExist:
			raise ProcessingError(errMsg + " : RecommendationScenario with id " + str(recommendationScenarioId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDatasets( self, recommendationScenarioId, datasetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to remove elements " + str(datasetsIds) + " for Datasets on RecommendationScenario"

		try:
			# get the RecommendationScenario
			recommendationScenario = self.get( recommendationScenarioId ).first()
				
			# split on a comma with no spaces
			idList = datasetsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataSet		
				dataSet = DataSetDelegate().get(id).first();	
				# add the DataSet
				recommendationScenario.datasets.remove(dataSet)
				
			# save it		
			recommendationScenario.save()
			
			# reload and return the appropriate version
			return self.get( recommendationScenarioId );
		except RecommendationScenario.DoesNotExist:
			raise ProcessingError(errMsg + " : RecommendationScenario with id " + str(recommendationScenarioId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addExperiments( self, recommendationScenarioId, experimentsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.ExperimentDelegate import ExperimentDelegate

		errMsg = "Failed to add elements " + str(experimentsIds) + " for Experiments on RecommendationScenario"

		try:
			# get the RecommendationScenario
			recommendationScenario = self.get( recommendationScenarioId ).first()
				
			# split on a comma with no spaces
			idList = experimentsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Experiment		
				experiment = ExperimentDelegate().get(id).first();	
				# add the Experiment
				recommendationScenario.experiments.add(experiment)
				
			# save it		
			recommendationScenario.save()
			
			# reload and return the appropriate version
			return self.get( recommendationScenarioId );
		except RecommendationScenario.DoesNotExist:
			raise ProcessingError(errMsg + " : RecommendationScenario with id " + str(recommendationScenarioId) + " does not exist.")
		except Experiment.DoesNotExist:
			raise ProcessingError(errMsg + " : Experiment does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeExperiments( self, recommendationScenarioId, experimentsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.ExperimentDelegate import ExperimentDelegate

		errMsg = "Failed to remove elements " + str(experimentsIds) + " for Experiments on RecommendationScenario"

		try:
			# get the RecommendationScenario
			recommendationScenario = self.get( recommendationScenarioId ).first()
				
			# split on a comma with no spaces
			idList = experimentsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Experiment		
				experiment = ExperimentDelegate().get(id).first();	
				# add the Experiment
				recommendationScenario.experiments.remove(experiment)
				
			# save it		
			recommendationScenario.save()
			
			# reload and return the appropriate version
			return self.get( recommendationScenarioId );
		except RecommendationScenario.DoesNotExist:
			raise ProcessingError(errMsg + " : RecommendationScenario with id " + str(recommendationScenarioId) + " does not exist.")
		except Experiment.DoesNotExist:
			raise ProcessingError(errMsg + " : Experiment does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addAlerts( self, recommendationScenarioId, alertsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.AlertDelegate import AlertDelegate

		errMsg = "Failed to add elements " + str(alertsIds) + " for Alerts on RecommendationScenario"

		try:
			# get the RecommendationScenario
			recommendationScenario = self.get( recommendationScenarioId ).first()
				
			# split on a comma with no spaces
			idList = alertsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Alert		
				alert = AlertDelegate().get(id).first();	
				# add the Alert
				recommendationScenario.alerts.add(alert)
				
			# save it		
			recommendationScenario.save()
			
			# reload and return the appropriate version
			return self.get( recommendationScenarioId );
		except RecommendationScenario.DoesNotExist:
			raise ProcessingError(errMsg + " : RecommendationScenario with id " + str(recommendationScenarioId) + " does not exist.")
		except Alert.DoesNotExist:
			raise ProcessingError(errMsg + " : Alert does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAlerts( self, recommendationScenarioId, alertsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.AlertDelegate import AlertDelegate

		errMsg = "Failed to remove elements " + str(alertsIds) + " for Alerts on RecommendationScenario"

		try:
			# get the RecommendationScenario
			recommendationScenario = self.get( recommendationScenarioId ).first()
				
			# split on a comma with no spaces
			idList = alertsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Alert		
				alert = AlertDelegate().get(id).first();	
				# add the Alert
				recommendationScenario.alerts.remove(alert)
				
			# save it		
			recommendationScenario.save()
			
			# reload and return the appropriate version
			return self.get( recommendationScenarioId );
		except RecommendationScenario.DoesNotExist:
			raise ProcessingError(errMsg + " : RecommendationScenario with id " + str(recommendationScenarioId) + " does not exist.")
		except Alert.DoesNotExist:
			raise ProcessingError(errMsg + " : Alert does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
