from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from analyticsOnDjango.models.AnalyticsWorkspace import AnalyticsWorkspace
from analyticsOnDjango.models.DataSet import DataSet
from analyticsOnDjango.models.DataSource import DataSource
from analyticsOnDjango.models.DataPipeline import DataPipeline
from analyticsOnDjango.models.Dashboard import Dashboard
from analyticsOnDjango.models.Report import Report
from analyticsOnDjango.models.Notebook import Notebook
from analyticsOnDjango.models.Model import Model
from analyticsOnDjango.models.FeatureSet import FeatureSet
from analyticsOnDjango.models.AccessPolicy import AccessPolicy
from analyticsOnDjango.models.LineageNode import LineageNode
from analyticsOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model AnalyticsWorkspace
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AnalyticsWorkspaceDelegate Declaration
#======================================================================
class AnalyticsWorkspaceDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, analyticsWorkspaceId ):
		try:	
			analyticsWorkspace = AnalyticsWorkspace.objects.filter(id=analyticsWorkspaceId)
			return analyticsWorkspace.first();
		except AnalyticsWorkspace.DoesNotExist:
			raise ProcessingError("AnalyticsWorkspace with id " + str(analyticsWorkspaceId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, analyticsWorkspace):
		for model in serializers.deserialize("json", analyticsWorkspace):
			model.save()
			return model;

	def create(self, analyticsWorkspace):
		analyticsWorkspace.save()
		return analyticsWorkspace;

	def saveFromJson(self, analyticsWorkspace):
		for model in serializers.deserialize("json", analyticsWorkspace):
			model.save()
			return analyticsWorkspace;
	
	def save(self, analyticsWorkspace):
		analyticsWorkspace.save()
		return analyticsWorkspace;
	
	def delete(self, analyticsWorkspaceId ):
		errMsg = "Failed to delete AnalyticsWorkspace from db using id " + str(analyticsWorkspaceId)
		
		try:
			analyticsWorkspace = AnalyticsWorkspace.objects.get(id=analyticsWorkspaceId)
			analyticsWorkspace.delete()
			return True
		except AnalyticsWorkspace.DoesNotExist:
			raise ProcessingError("AnalyticsWorkspace with id " + str(analyticsWorkspaceId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = AnalyticsWorkspace.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all AnalyticsWorkspace from db")
		except Exception:
			return None;
		
	def addDatasets( self, analyticsWorkspaceId, datasetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to add elements " + str(datasetsIds) + " for Datasets on AnalyticsWorkspace"

		try:
			# get the AnalyticsWorkspace
			analyticsWorkspace = self.get( analyticsWorkspaceId ).first()
				
			# split on a comma with no spaces
			idList = datasetsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataSet		
				dataSet = DataSetDelegate().get(id).first();	
				# add the DataSet
				analyticsWorkspace.datasets.add(dataSet)
				
			# save it		
			analyticsWorkspace.save()
			
			# reload and return the appropriate version
			return self.get( analyticsWorkspaceId );
		except AnalyticsWorkspace.DoesNotExist:
			raise ProcessingError(errMsg + " : AnalyticsWorkspace with id " + str(analyticsWorkspaceId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDatasets( self, analyticsWorkspaceId, datasetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to remove elements " + str(datasetsIds) + " for Datasets on AnalyticsWorkspace"

		try:
			# get the AnalyticsWorkspace
			analyticsWorkspace = self.get( analyticsWorkspaceId ).first()
				
			# split on a comma with no spaces
			idList = datasetsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataSet		
				dataSet = DataSetDelegate().get(id).first();	
				# add the DataSet
				analyticsWorkspace.datasets.remove(dataSet)
				
			# save it		
			analyticsWorkspace.save()
			
			# reload and return the appropriate version
			return self.get( analyticsWorkspaceId );
		except AnalyticsWorkspace.DoesNotExist:
			raise ProcessingError(errMsg + " : AnalyticsWorkspace with id " + str(analyticsWorkspaceId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addDataSources( self, analyticsWorkspaceId, dataSourcesIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSourceDelegate import DataSourceDelegate

		errMsg = "Failed to add elements " + str(dataSourcesIds) + " for DataSources on AnalyticsWorkspace"

		try:
			# get the AnalyticsWorkspace
			analyticsWorkspace = self.get( analyticsWorkspaceId ).first()
				
			# split on a comma with no spaces
			idList = dataSourcesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataSource		
				dataSource = DataSourceDelegate().get(id).first();	
				# add the DataSource
				analyticsWorkspace.dataSources.add(dataSource)
				
			# save it		
			analyticsWorkspace.save()
			
			# reload and return the appropriate version
			return self.get( analyticsWorkspaceId );
		except AnalyticsWorkspace.DoesNotExist:
			raise ProcessingError(errMsg + " : AnalyticsWorkspace with id " + str(analyticsWorkspaceId) + " does not exist.")
		except DataSource.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSource does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDataSources( self, analyticsWorkspaceId, dataSourcesIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSourceDelegate import DataSourceDelegate

		errMsg = "Failed to remove elements " + str(dataSourcesIds) + " for DataSources on AnalyticsWorkspace"

		try:
			# get the AnalyticsWorkspace
			analyticsWorkspace = self.get( analyticsWorkspaceId ).first()
				
			# split on a comma with no spaces
			idList = dataSourcesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataSource		
				dataSource = DataSourceDelegate().get(id).first();	
				# add the DataSource
				analyticsWorkspace.dataSources.remove(dataSource)
				
			# save it		
			analyticsWorkspace.save()
			
			# reload and return the appropriate version
			return self.get( analyticsWorkspaceId );
		except AnalyticsWorkspace.DoesNotExist:
			raise ProcessingError(errMsg + " : AnalyticsWorkspace with id " + str(analyticsWorkspaceId) + " does not exist.")
		except DataSource.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSource does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addPipelines( self, analyticsWorkspaceId, pipelinesIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataPipelineDelegate import DataPipelineDelegate

		errMsg = "Failed to add elements " + str(pipelinesIds) + " for Pipelines on AnalyticsWorkspace"

		try:
			# get the AnalyticsWorkspace
			analyticsWorkspace = self.get( analyticsWorkspaceId ).first()
				
			# split on a comma with no spaces
			idList = pipelinesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataPipeline		
				dataPipeline = DataPipelineDelegate().get(id).first();	
				# add the DataPipeline
				analyticsWorkspace.pipelines.add(dataPipeline)
				
			# save it		
			analyticsWorkspace.save()
			
			# reload and return the appropriate version
			return self.get( analyticsWorkspaceId );
		except AnalyticsWorkspace.DoesNotExist:
			raise ProcessingError(errMsg + " : AnalyticsWorkspace with id " + str(analyticsWorkspaceId) + " does not exist.")
		except DataPipeline.DoesNotExist:
			raise ProcessingError(errMsg + " : DataPipeline does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePipelines( self, analyticsWorkspaceId, pipelinesIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataPipelineDelegate import DataPipelineDelegate

		errMsg = "Failed to remove elements " + str(pipelinesIds) + " for Pipelines on AnalyticsWorkspace"

		try:
			# get the AnalyticsWorkspace
			analyticsWorkspace = self.get( analyticsWorkspaceId ).first()
				
			# split on a comma with no spaces
			idList = pipelinesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataPipeline		
				dataPipeline = DataPipelineDelegate().get(id).first();	
				# add the DataPipeline
				analyticsWorkspace.pipelines.remove(dataPipeline)
				
			# save it		
			analyticsWorkspace.save()
			
			# reload and return the appropriate version
			return self.get( analyticsWorkspaceId );
		except AnalyticsWorkspace.DoesNotExist:
			raise ProcessingError(errMsg + " : AnalyticsWorkspace with id " + str(analyticsWorkspaceId) + " does not exist.")
		except DataPipeline.DoesNotExist:
			raise ProcessingError(errMsg + " : DataPipeline does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addDashboards( self, analyticsWorkspaceId, dashboardsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DashboardDelegate import DashboardDelegate

		errMsg = "Failed to add elements " + str(dashboardsIds) + " for Dashboards on AnalyticsWorkspace"

		try:
			# get the AnalyticsWorkspace
			analyticsWorkspace = self.get( analyticsWorkspaceId ).first()
				
			# split on a comma with no spaces
			idList = dashboardsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Dashboard		
				dashboard = DashboardDelegate().get(id).first();	
				# add the Dashboard
				analyticsWorkspace.dashboards.add(dashboard)
				
			# save it		
			analyticsWorkspace.save()
			
			# reload and return the appropriate version
			return self.get( analyticsWorkspaceId );
		except AnalyticsWorkspace.DoesNotExist:
			raise ProcessingError(errMsg + " : AnalyticsWorkspace with id " + str(analyticsWorkspaceId) + " does not exist.")
		except Dashboard.DoesNotExist:
			raise ProcessingError(errMsg + " : Dashboard does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDashboards( self, analyticsWorkspaceId, dashboardsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DashboardDelegate import DashboardDelegate

		errMsg = "Failed to remove elements " + str(dashboardsIds) + " for Dashboards on AnalyticsWorkspace"

		try:
			# get the AnalyticsWorkspace
			analyticsWorkspace = self.get( analyticsWorkspaceId ).first()
				
			# split on a comma with no spaces
			idList = dashboardsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Dashboard		
				dashboard = DashboardDelegate().get(id).first();	
				# add the Dashboard
				analyticsWorkspace.dashboards.remove(dashboard)
				
			# save it		
			analyticsWorkspace.save()
			
			# reload and return the appropriate version
			return self.get( analyticsWorkspaceId );
		except AnalyticsWorkspace.DoesNotExist:
			raise ProcessingError(errMsg + " : AnalyticsWorkspace with id " + str(analyticsWorkspaceId) + " does not exist.")
		except Dashboard.DoesNotExist:
			raise ProcessingError(errMsg + " : Dashboard does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addReports( self, analyticsWorkspaceId, reportsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.ReportDelegate import ReportDelegate

		errMsg = "Failed to add elements " + str(reportsIds) + " for Reports on AnalyticsWorkspace"

		try:
			# get the AnalyticsWorkspace
			analyticsWorkspace = self.get( analyticsWorkspaceId ).first()
				
			# split on a comma with no spaces
			idList = reportsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Report		
				report = ReportDelegate().get(id).first();	
				# add the Report
				analyticsWorkspace.reports.add(report)
				
			# save it		
			analyticsWorkspace.save()
			
			# reload and return the appropriate version
			return self.get( analyticsWorkspaceId );
		except AnalyticsWorkspace.DoesNotExist:
			raise ProcessingError(errMsg + " : AnalyticsWorkspace with id " + str(analyticsWorkspaceId) + " does not exist.")
		except Report.DoesNotExist:
			raise ProcessingError(errMsg + " : Report does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeReports( self, analyticsWorkspaceId, reportsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.ReportDelegate import ReportDelegate

		errMsg = "Failed to remove elements " + str(reportsIds) + " for Reports on AnalyticsWorkspace"

		try:
			# get the AnalyticsWorkspace
			analyticsWorkspace = self.get( analyticsWorkspaceId ).first()
				
			# split on a comma with no spaces
			idList = reportsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Report		
				report = ReportDelegate().get(id).first();	
				# add the Report
				analyticsWorkspace.reports.remove(report)
				
			# save it		
			analyticsWorkspace.save()
			
			# reload and return the appropriate version
			return self.get( analyticsWorkspaceId );
		except AnalyticsWorkspace.DoesNotExist:
			raise ProcessingError(errMsg + " : AnalyticsWorkspace with id " + str(analyticsWorkspaceId) + " does not exist.")
		except Report.DoesNotExist:
			raise ProcessingError(errMsg + " : Report does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addNotebooks( self, analyticsWorkspaceId, notebooksIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.NotebookDelegate import NotebookDelegate

		errMsg = "Failed to add elements " + str(notebooksIds) + " for Notebooks on AnalyticsWorkspace"

		try:
			# get the AnalyticsWorkspace
			analyticsWorkspace = self.get( analyticsWorkspaceId ).first()
				
			# split on a comma with no spaces
			idList = notebooksIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Notebook		
				notebook = NotebookDelegate().get(id).first();	
				# add the Notebook
				analyticsWorkspace.notebooks.add(notebook)
				
			# save it		
			analyticsWorkspace.save()
			
			# reload and return the appropriate version
			return self.get( analyticsWorkspaceId );
		except AnalyticsWorkspace.DoesNotExist:
			raise ProcessingError(errMsg + " : AnalyticsWorkspace with id " + str(analyticsWorkspaceId) + " does not exist.")
		except Notebook.DoesNotExist:
			raise ProcessingError(errMsg + " : Notebook does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeNotebooks( self, analyticsWorkspaceId, notebooksIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.NotebookDelegate import NotebookDelegate

		errMsg = "Failed to remove elements " + str(notebooksIds) + " for Notebooks on AnalyticsWorkspace"

		try:
			# get the AnalyticsWorkspace
			analyticsWorkspace = self.get( analyticsWorkspaceId ).first()
				
			# split on a comma with no spaces
			idList = notebooksIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Notebook		
				notebook = NotebookDelegate().get(id).first();	
				# add the Notebook
				analyticsWorkspace.notebooks.remove(notebook)
				
			# save it		
			analyticsWorkspace.save()
			
			# reload and return the appropriate version
			return self.get( analyticsWorkspaceId );
		except AnalyticsWorkspace.DoesNotExist:
			raise ProcessingError(errMsg + " : AnalyticsWorkspace with id " + str(analyticsWorkspaceId) + " does not exist.")
		except Notebook.DoesNotExist:
			raise ProcessingError(errMsg + " : Notebook does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addModels( self, analyticsWorkspaceId, modelsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.ModelDelegate import ModelDelegate

		errMsg = "Failed to add elements " + str(modelsIds) + " for Models on AnalyticsWorkspace"

		try:
			# get the AnalyticsWorkspace
			analyticsWorkspace = self.get( analyticsWorkspaceId ).first()
				
			# split on a comma with no spaces
			idList = modelsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Model		
				model = ModelDelegate().get(id).first();	
				# add the Model
				analyticsWorkspace.models.add(model)
				
			# save it		
			analyticsWorkspace.save()
			
			# reload and return the appropriate version
			return self.get( analyticsWorkspaceId );
		except AnalyticsWorkspace.DoesNotExist:
			raise ProcessingError(errMsg + " : AnalyticsWorkspace with id " + str(analyticsWorkspaceId) + " does not exist.")
		except Model.DoesNotExist:
			raise ProcessingError(errMsg + " : Model does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeModels( self, analyticsWorkspaceId, modelsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.ModelDelegate import ModelDelegate

		errMsg = "Failed to remove elements " + str(modelsIds) + " for Models on AnalyticsWorkspace"

		try:
			# get the AnalyticsWorkspace
			analyticsWorkspace = self.get( analyticsWorkspaceId ).first()
				
			# split on a comma with no spaces
			idList = modelsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Model		
				model = ModelDelegate().get(id).first();	
				# add the Model
				analyticsWorkspace.models.remove(model)
				
			# save it		
			analyticsWorkspace.save()
			
			# reload and return the appropriate version
			return self.get( analyticsWorkspaceId );
		except AnalyticsWorkspace.DoesNotExist:
			raise ProcessingError(errMsg + " : AnalyticsWorkspace with id " + str(analyticsWorkspaceId) + " does not exist.")
		except Model.DoesNotExist:
			raise ProcessingError(errMsg + " : Model does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addFeatureSets( self, analyticsWorkspaceId, featureSetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.FeatureSetDelegate import FeatureSetDelegate

		errMsg = "Failed to add elements " + str(featureSetsIds) + " for FeatureSets on AnalyticsWorkspace"

		try:
			# get the AnalyticsWorkspace
			analyticsWorkspace = self.get( analyticsWorkspaceId ).first()
				
			# split on a comma with no spaces
			idList = featureSetsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the FeatureSet		
				featureSet = FeatureSetDelegate().get(id).first();	
				# add the FeatureSet
				analyticsWorkspace.featureSets.add(featureSet)
				
			# save it		
			analyticsWorkspace.save()
			
			# reload and return the appropriate version
			return self.get( analyticsWorkspaceId );
		except AnalyticsWorkspace.DoesNotExist:
			raise ProcessingError(errMsg + " : AnalyticsWorkspace with id " + str(analyticsWorkspaceId) + " does not exist.")
		except FeatureSet.DoesNotExist:
			raise ProcessingError(errMsg + " : FeatureSet does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeFeatureSets( self, analyticsWorkspaceId, featureSetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.FeatureSetDelegate import FeatureSetDelegate

		errMsg = "Failed to remove elements " + str(featureSetsIds) + " for FeatureSets on AnalyticsWorkspace"

		try:
			# get the AnalyticsWorkspace
			analyticsWorkspace = self.get( analyticsWorkspaceId ).first()
				
			# split on a comma with no spaces
			idList = featureSetsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the FeatureSet		
				featureSet = FeatureSetDelegate().get(id).first();	
				# add the FeatureSet
				analyticsWorkspace.featureSets.remove(featureSet)
				
			# save it		
			analyticsWorkspace.save()
			
			# reload and return the appropriate version
			return self.get( analyticsWorkspaceId );
		except AnalyticsWorkspace.DoesNotExist:
			raise ProcessingError(errMsg + " : AnalyticsWorkspace with id " + str(analyticsWorkspaceId) + " does not exist.")
		except FeatureSet.DoesNotExist:
			raise ProcessingError(errMsg + " : FeatureSet does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addPolicies( self, analyticsWorkspaceId, policiesIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.AccessPolicyDelegate import AccessPolicyDelegate

		errMsg = "Failed to add elements " + str(policiesIds) + " for Policies on AnalyticsWorkspace"

		try:
			# get the AnalyticsWorkspace
			analyticsWorkspace = self.get( analyticsWorkspaceId ).first()
				
			# split on a comma with no spaces
			idList = policiesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the AccessPolicy		
				accessPolicy = AccessPolicyDelegate().get(id).first();	
				# add the AccessPolicy
				analyticsWorkspace.policies.add(accessPolicy)
				
			# save it		
			analyticsWorkspace.save()
			
			# reload and return the appropriate version
			return self.get( analyticsWorkspaceId );
		except AnalyticsWorkspace.DoesNotExist:
			raise ProcessingError(errMsg + " : AnalyticsWorkspace with id " + str(analyticsWorkspaceId) + " does not exist.")
		except AccessPolicy.DoesNotExist:
			raise ProcessingError(errMsg + " : AccessPolicy does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePolicies( self, analyticsWorkspaceId, policiesIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.AccessPolicyDelegate import AccessPolicyDelegate

		errMsg = "Failed to remove elements " + str(policiesIds) + " for Policies on AnalyticsWorkspace"

		try:
			# get the AnalyticsWorkspace
			analyticsWorkspace = self.get( analyticsWorkspaceId ).first()
				
			# split on a comma with no spaces
			idList = policiesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the AccessPolicy		
				accessPolicy = AccessPolicyDelegate().get(id).first();	
				# add the AccessPolicy
				analyticsWorkspace.policies.remove(accessPolicy)
				
			# save it		
			analyticsWorkspace.save()
			
			# reload and return the appropriate version
			return self.get( analyticsWorkspaceId );
		except AnalyticsWorkspace.DoesNotExist:
			raise ProcessingError(errMsg + " : AnalyticsWorkspace with id " + str(analyticsWorkspaceId) + " does not exist.")
		except AccessPolicy.DoesNotExist:
			raise ProcessingError(errMsg + " : AccessPolicy does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addLineageNodes( self, analyticsWorkspaceId, lineageNodesIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.LineageNodeDelegate import LineageNodeDelegate

		errMsg = "Failed to add elements " + str(lineageNodesIds) + " for LineageNodes on AnalyticsWorkspace"

		try:
			# get the AnalyticsWorkspace
			analyticsWorkspace = self.get( analyticsWorkspaceId ).first()
				
			# split on a comma with no spaces
			idList = lineageNodesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the LineageNode		
				lineageNode = LineageNodeDelegate().get(id).first();	
				# add the LineageNode
				analyticsWorkspace.lineageNodes.add(lineageNode)
				
			# save it		
			analyticsWorkspace.save()
			
			# reload and return the appropriate version
			return self.get( analyticsWorkspaceId );
		except AnalyticsWorkspace.DoesNotExist:
			raise ProcessingError(errMsg + " : AnalyticsWorkspace with id " + str(analyticsWorkspaceId) + " does not exist.")
		except LineageNode.DoesNotExist:
			raise ProcessingError(errMsg + " : LineageNode does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeLineageNodes( self, analyticsWorkspaceId, lineageNodesIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.LineageNodeDelegate import LineageNodeDelegate

		errMsg = "Failed to remove elements " + str(lineageNodesIds) + " for LineageNodes on AnalyticsWorkspace"

		try:
			# get the AnalyticsWorkspace
			analyticsWorkspace = self.get( analyticsWorkspaceId ).first()
				
			# split on a comma with no spaces
			idList = lineageNodesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the LineageNode		
				lineageNode = LineageNodeDelegate().get(id).first();	
				# add the LineageNode
				analyticsWorkspace.lineageNodes.remove(lineageNode)
				
			# save it		
			analyticsWorkspace.save()
			
			# reload and return the appropriate version
			return self.get( analyticsWorkspaceId );
		except AnalyticsWorkspace.DoesNotExist:
			raise ProcessingError(errMsg + " : AnalyticsWorkspace with id " + str(analyticsWorkspaceId) + " does not exist.")
		except LineageNode.DoesNotExist:
			raise ProcessingError(errMsg + " : LineageNode does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
