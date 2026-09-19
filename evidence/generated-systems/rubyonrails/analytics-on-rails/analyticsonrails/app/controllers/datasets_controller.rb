class DataSetsController < ApplicationController
  def index
    @dataSets = DataSet.all
  end
 
  def show
    @dataSet = DataSet.find(params[:id])
  end
 
  def new
    @dataSet = DataSet.new
  end
 
  def edit
    @dataSet = DataSet.find(params[:id])
  end
 
  def create
    @dataSet = DataSet.new(dataSet_params)
 
    if @dataSet.save
      redirect_to dataSets_path
    else
      render 'new'
    end
  end
 
  def update
    @dataSet = DataSet.find(params[:id])
 
    if @dataSet.update(dataSet_params)
      redirect_to dataSets_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @dataSet = DataSet.find(params[:id])
    @dataSet.destroy
    redirect_to dataSets_path
  end

 
  private
    def dataSet_params
      params.require(:dataSet).permit(:name, :schemaVersion, :refreshSchedule, :Sensitive, :DataFormat)
    end
end