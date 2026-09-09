from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from analyticsOnDjango.models.FraudScenario import FraudScenario
from analyticsOnDjango.models.Model import Model
from analyticsOnDjango.models.DataSet import DataSet
from analyticsOnDjango.models.Alert import Alert
from analyticsOnDjango.models.FraudSignal import FraudSignal
from analyticsOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model FraudScenario
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class FraudScenarioDelegate Declaration
#======================================================================
class FraudScenarioDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, fraudScenarioId ):
		try:	
			fraudScenario = FraudScenario.objects.filter(id=fraudScenarioId)
			return fraudScenario.first();
		except FraudScenario.DoesNotExist:
			raise ProcessingError("FraudScenario with id " + str(fraudScenarioId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, fraudScenario):
		for model in serializers.deserialize("json", fraudScenario):
			model.save()
			return model;

	def create(self, fraudScenario):
		fraudScenario.save()
		return fraudScenario;

	def saveFromJson(self, fraudScenario):
		for model in serializers.deserialize("json", fraudScenario):
			model.save()
			return fraudScenario;
	
	def save(self, fraudScenario):
		fraudScenario.save()
		return fraudScenario;
	
	def delete(self, fraudScenarioId ):
		errMsg = "Failed to delete FraudScenario from db using id " + str(fraudScenarioId)
		
		try:
			fraudScenario = FraudScenario.objects.get(id=fraudScenarioId)
			fraudScenario.delete()
			return True
		except FraudScenario.DoesNotExist:
			raise ProcessingError("FraudScenario with id " + str(fraudScenarioId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = FraudScenario.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all FraudScenario from db")
		except Exception:
			return None;
		
	def addModels( self, fraudScenarioId, modelsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.ModelDelegate import ModelDelegate

		errMsg = "Failed to add elements " + str(modelsIds) + " for Models on FraudScenario"

		try:
			# get the FraudScenario
			fraudScenario = self.get( fraudScenarioId ).first()
				
			# split on a comma with no spaces
			idList = modelsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Model		
				model = ModelDelegate().get(id).first();	
				# add the Model
				fraudScenario.models.add(model)
				
			# save it		
			fraudScenario.save()
			
			# reload and return the appropriate version
			return self.get( fraudScenarioId );
		except FraudScenario.DoesNotExist:
			raise ProcessingError(errMsg + " : FraudScenario with id " + str(fraudScenarioId) + " does not exist.")
		except Model.DoesNotExist:
			raise ProcessingError(errMsg + " : Model does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeModels( self, fraudScenarioId, modelsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.ModelDelegate import ModelDelegate

		errMsg = "Failed to remove elements " + str(modelsIds) + " for Models on FraudScenario"

		try:
			# get the FraudScenario
			fraudScenario = self.get( fraudScenarioId ).first()
				
			# split on a comma with no spaces
			idList = modelsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Model		
				model = ModelDelegate().get(id).first();	
				# add the Model
				fraudScenario.models.remove(model)
				
			# save it		
			fraudScenario.save()
			
			# reload and return the appropriate version
			return self.get( fraudScenarioId );
		except FraudScenario.DoesNotExist:
			raise ProcessingError(errMsg + " : FraudScenario with id " + str(fraudScenarioId) + " does not exist.")
		except Model.DoesNotExist:
			raise ProcessingError(errMsg + " : Model does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addDatasets( self, fraudScenarioId, datasetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to add elements " + str(datasetsIds) + " for Datasets on FraudScenario"

		try:
			# get the FraudScenario
			fraudScenario = self.get( fraudScenarioId ).first()
				
			# split on a comma with no spaces
			idList = datasetsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataSet		
				dataSet = DataSetDelegate().get(id).first();	
				# add the DataSet
				fraudScenario.datasets.add(dataSet)
				
			# save it		
			fraudScenario.save()
			
			# reload and return the appropriate version
			return self.get( fraudScenarioId );
		except FraudScenario.DoesNotExist:
			raise ProcessingError(errMsg + " : FraudScenario with id " + str(fraudScenarioId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDatasets( self, fraudScenarioId, datasetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to remove elements " + str(datasetsIds) + " for Datasets on FraudScenario"

		try:
			# get the FraudScenario
			fraudScenario = self.get( fraudScenarioId ).first()
				
			# split on a comma with no spaces
			idList = datasetsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataSet		
				dataSet = DataSetDelegate().get(id).first();	
				# add the DataSet
				fraudScenario.datasets.remove(dataSet)
				
			# save it		
			fraudScenario.save()
			
			# reload and return the appropriate version
			return self.get( fraudScenarioId );
		except FraudScenario.DoesNotExist:
			raise ProcessingError(errMsg + " : FraudScenario with id " + str(fraudScenarioId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addAlerts( self, fraudScenarioId, alertsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.AlertDelegate import AlertDelegate

		errMsg = "Failed to add elements " + str(alertsIds) + " for Alerts on FraudScenario"

		try:
			# get the FraudScenario
			fraudScenario = self.get( fraudScenarioId ).first()
				
			# split on a comma with no spaces
			idList = alertsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Alert		
				alert = AlertDelegate().get(id).first();	
				# add the Alert
				fraudScenario.alerts.add(alert)
				
			# save it		
			fraudScenario.save()
			
			# reload and return the appropriate version
			return self.get( fraudScenarioId );
		except FraudScenario.DoesNotExist:
			raise ProcessingError(errMsg + " : FraudScenario with id " + str(fraudScenarioId) + " does not exist.")
		except Alert.DoesNotExist:
			raise ProcessingError(errMsg + " : Alert does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAlerts( self, fraudScenarioId, alertsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.AlertDelegate import AlertDelegate

		errMsg = "Failed to remove elements " + str(alertsIds) + " for Alerts on FraudScenario"

		try:
			# get the FraudScenario
			fraudScenario = self.get( fraudScenarioId ).first()
				
			# split on a comma with no spaces
			idList = alertsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Alert		
				alert = AlertDelegate().get(id).first();	
				# add the Alert
				fraudScenario.alerts.remove(alert)
				
			# save it		
			fraudScenario.save()
			
			# reload and return the appropriate version
			return self.get( fraudScenarioId );
		except FraudScenario.DoesNotExist:
			raise ProcessingError(errMsg + " : FraudScenario with id " + str(fraudScenarioId) + " does not exist.")
		except Alert.DoesNotExist:
			raise ProcessingError(errMsg + " : Alert does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addSignals( self, fraudScenarioId, signalsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.FraudSignalDelegate import FraudSignalDelegate

		errMsg = "Failed to add elements " + str(signalsIds) + " for Signals on FraudScenario"

		try:
			# get the FraudScenario
			fraudScenario = self.get( fraudScenarioId ).first()
				
			# split on a comma with no spaces
			idList = signalsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the FraudSignal		
				fraudSignal = FraudSignalDelegate().get(id).first();	
				# add the FraudSignal
				fraudScenario.signals.add(fraudSignal)
				
			# save it		
			fraudScenario.save()
			
			# reload and return the appropriate version
			return self.get( fraudScenarioId );
		except FraudScenario.DoesNotExist:
			raise ProcessingError(errMsg + " : FraudScenario with id " + str(fraudScenarioId) + " does not exist.")
		except FraudSignal.DoesNotExist:
			raise ProcessingError(errMsg + " : FraudSignal does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeSignals( self, fraudScenarioId, signalsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.FraudSignalDelegate import FraudSignalDelegate

		errMsg = "Failed to remove elements " + str(signalsIds) + " for Signals on FraudScenario"

		try:
			# get the FraudScenario
			fraudScenario = self.get( fraudScenarioId ).first()
				
			# split on a comma with no spaces
			idList = signalsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the FraudSignal		
				fraudSignal = FraudSignalDelegate().get(id).first();	
				# add the FraudSignal
				fraudScenario.signals.remove(fraudSignal)
				
			# save it		
			fraudScenario.save()
			
			# reload and return the appropriate version
			return self.get( fraudScenarioId );
		except FraudScenario.DoesNotExist:
			raise ProcessingError(errMsg + " : FraudScenario with id " + str(fraudScenarioId) + " does not exist.")
		except FraudSignal.DoesNotExist:
			raise ProcessingError(errMsg + " : FraudSignal does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
