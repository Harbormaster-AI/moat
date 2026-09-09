from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from analyticsOnDjango.models.Visualization import Visualization
from analyticsOnDjango.models.Dashboard import Dashboard
from analyticsOnDjango.models.Report import Report
from analyticsOnDjango.models.Metric import Metric
from analyticsOnDjango.models.Dimension import Dimension
from analyticsOnDjango.models.DataSet import DataSet
from analyticsOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Visualization
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class VisualizationDelegate Declaration
#======================================================================
class VisualizationDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, visualizationId ):
		try:	
			visualization = Visualization.objects.filter(id=visualizationId)
			return visualization.first();
		except Visualization.DoesNotExist:
			raise ProcessingError("Visualization with id " + str(visualizationId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, visualization):
		for model in serializers.deserialize("json", visualization):
			model.save()
			return model;

	def create(self, visualization):
		visualization.save()
		return visualization;

	def saveFromJson(self, visualization):
		for model in serializers.deserialize("json", visualization):
			model.save()
			return visualization;
	
	def save(self, visualization):
		visualization.save()
		return visualization;
	
	def delete(self, visualizationId ):
		errMsg = "Failed to delete Visualization from db using id " + str(visualizationId)
		
		try:
			visualization = Visualization.objects.get(id=visualizationId)
			visualization.delete()
			return True
		except Visualization.DoesNotExist:
			raise ProcessingError("Visualization with id " + str(visualizationId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Visualization.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Visualization from db")
		except Exception:
			return None;
		
	def assignDashboard( self, visualizationId, dashboardId ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DashboardDelegate import DashboardDelegate

		errMsg = "Failed to assign element " + str(dashboardId) + " for Dashboard on Visualization"

		try:
			# get the Visualization from db
			visualization = self.get( visualizationId ).first()	
			
			# get the Dashboard from db
			dashboard = DashboardDelegate().get(dashboardId).first();
			
			# assign the Dashboard		
			visualization.dashboard = dashboard
			
			#save it
			visualization.save()

			# reload and return the appropriate version					
			return self.get( visualizationId );
		except Visualization.DoesNotExist:
			raise ProcessingError(errMsg + " : Visualization with id " + str(visualizationId) + " does not exist.")
		except Dashboard.DoesNotExist:
			raise ProcessingError(errMsg + " : Dashboard with id " + str(dashboardId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignDashboard( self, visualizationId ):
		errMsg = "Failed to unassign element " + str(dashboardId) + " for Dashboard on Visualization"

		try:
			# get the Visualization from db
			visualization = self.get( visualizationId ).first()	
			
			# assign to None for unassignment
			visualization.dashboard = None			

			#save it
			visualization.save()

			# reload and return the appropriate version					
			return self.get( visualizationId );
		except Visualization.DoesNotExist:
			raise ProcessingError(errMsg + " : Visualization with id " + str(visualizationId) + " does not exist.")
		except Exception:
			return None;
		
	def assignReport( self, visualizationId, reportId ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.ReportDelegate import ReportDelegate

		errMsg = "Failed to assign element " + str(reportId) + " for Report on Visualization"

		try:
			# get the Visualization from db
			visualization = self.get( visualizationId ).first()	
			
			# get the Report from db
			report = ReportDelegate().get(reportId).first();
			
			# assign the Report		
			visualization.report = report
			
			#save it
			visualization.save()

			# reload and return the appropriate version					
			return self.get( visualizationId );
		except Visualization.DoesNotExist:
			raise ProcessingError(errMsg + " : Visualization with id " + str(visualizationId) + " does not exist.")
		except Report.DoesNotExist:
			raise ProcessingError(errMsg + " : Report with id " + str(reportId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignReport( self, visualizationId ):
		errMsg = "Failed to unassign element " + str(reportId) + " for Report on Visualization"

		try:
			# get the Visualization from db
			visualization = self.get( visualizationId ).first()	
			
			# assign to None for unassignment
			visualization.report = None			

			#save it
			visualization.save()

			# reload and return the appropriate version					
			return self.get( visualizationId );
		except Visualization.DoesNotExist:
			raise ProcessingError(errMsg + " : Visualization with id " + str(visualizationId) + " does not exist.")
		except Exception:
			return None;
		
	def addMetrics( self, visualizationId, metricsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.MetricDelegate import MetricDelegate

		errMsg = "Failed to add elements " + str(metricsIds) + " for Metrics on Visualization"

		try:
			# get the Visualization
			visualization = self.get( visualizationId ).first()
				
			# split on a comma with no spaces
			idList = metricsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Metric		
				metric = MetricDelegate().get(id).first();	
				# add the Metric
				visualization.metrics.add(metric)
				
			# save it		
			visualization.save()
			
			# reload and return the appropriate version
			return self.get( visualizationId );
		except Visualization.DoesNotExist:
			raise ProcessingError(errMsg + " : Visualization with id " + str(visualizationId) + " does not exist.")
		except Metric.DoesNotExist:
			raise ProcessingError(errMsg + " : Metric does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeMetrics( self, visualizationId, metricsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.MetricDelegate import MetricDelegate

		errMsg = "Failed to remove elements " + str(metricsIds) + " for Metrics on Visualization"

		try:
			# get the Visualization
			visualization = self.get( visualizationId ).first()
				
			# split on a comma with no spaces
			idList = metricsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Metric		
				metric = MetricDelegate().get(id).first();	
				# add the Metric
				visualization.metrics.remove(metric)
				
			# save it		
			visualization.save()
			
			# reload and return the appropriate version
			return self.get( visualizationId );
		except Visualization.DoesNotExist:
			raise ProcessingError(errMsg + " : Visualization with id " + str(visualizationId) + " does not exist.")
		except Metric.DoesNotExist:
			raise ProcessingError(errMsg + " : Metric does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addDimensions( self, visualizationId, dimensionsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DimensionDelegate import DimensionDelegate

		errMsg = "Failed to add elements " + str(dimensionsIds) + " for Dimensions on Visualization"

		try:
			# get the Visualization
			visualization = self.get( visualizationId ).first()
				
			# split on a comma with no spaces
			idList = dimensionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Dimension		
				dimension = DimensionDelegate().get(id).first();	
				# add the Dimension
				visualization.dimensions.add(dimension)
				
			# save it		
			visualization.save()
			
			# reload and return the appropriate version
			return self.get( visualizationId );
		except Visualization.DoesNotExist:
			raise ProcessingError(errMsg + " : Visualization with id " + str(visualizationId) + " does not exist.")
		except Dimension.DoesNotExist:
			raise ProcessingError(errMsg + " : Dimension does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDimensions( self, visualizationId, dimensionsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DimensionDelegate import DimensionDelegate

		errMsg = "Failed to remove elements " + str(dimensionsIds) + " for Dimensions on Visualization"

		try:
			# get the Visualization
			visualization = self.get( visualizationId ).first()
				
			# split on a comma with no spaces
			idList = dimensionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Dimension		
				dimension = DimensionDelegate().get(id).first();	
				# add the Dimension
				visualization.dimensions.remove(dimension)
				
			# save it		
			visualization.save()
			
			# reload and return the appropriate version
			return self.get( visualizationId );
		except Visualization.DoesNotExist:
			raise ProcessingError(errMsg + " : Visualization with id " + str(visualizationId) + " does not exist.")
		except Dimension.DoesNotExist:
			raise ProcessingError(errMsg + " : Dimension does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addDatasets( self, visualizationId, datasetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to add elements " + str(datasetsIds) + " for Datasets on Visualization"

		try:
			# get the Visualization
			visualization = self.get( visualizationId ).first()
				
			# split on a comma with no spaces
			idList = datasetsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataSet		
				dataSet = DataSetDelegate().get(id).first();	
				# add the DataSet
				visualization.datasets.add(dataSet)
				
			# save it		
			visualization.save()
			
			# reload and return the appropriate version
			return self.get( visualizationId );
		except Visualization.DoesNotExist:
			raise ProcessingError(errMsg + " : Visualization with id " + str(visualizationId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDatasets( self, visualizationId, datasetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to remove elements " + str(datasetsIds) + " for Datasets on Visualization"

		try:
			# get the Visualization
			visualization = self.get( visualizationId ).first()
				
			# split on a comma with no spaces
			idList = datasetsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataSet		
				dataSet = DataSetDelegate().get(id).first();	
				# add the DataSet
				visualization.datasets.remove(dataSet)
				
			# save it		
			visualization.save()
			
			# reload and return the appropriate version
			return self.get( visualizationId );
		except Visualization.DoesNotExist:
			raise ProcessingError(errMsg + " : Visualization with id " + str(visualizationId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
