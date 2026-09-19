class Visualization < ApplicationRecord
  enum ChartType: [:Table, :Bar, :Line, :Area, :Pie, :Scatter, :Heatmap, :KPI]


  composed_of :chartOptions,
    class_name: "ChartOptions",
    mapping: [
      ${$mapping}, 
      %w[chartOptions_legendPosition legendPosition]
    ]

  has_many :Dashboard, class_name: 'Dashboard'
  has_many :Report, class_name: 'Report'
  has_many :Metrics, class_name: 'Metric'
  has_many :Dimensions, class_name: 'Dimension'
  has_many :Datasets, class_name: 'DataSet'

end
