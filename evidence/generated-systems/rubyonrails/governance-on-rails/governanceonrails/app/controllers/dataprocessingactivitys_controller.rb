class DataProcessingActivitysController < ApplicationController
  def index
    @dataProcessingActivitys = DataProcessingActivity.all
  end
 
  def show
    @dataProcessingActivity = DataProcessingActivity.find(params[:id])
  end
 
  def new
    @dataProcessingActivity = DataProcessingActivity.new
  end
 
  def edit
    @dataProcessingActivity = DataProcessingActivity.find(params[:id])
  end
 
  def create
    @dataProcessingActivity = DataProcessingActivity.new(dataProcessingActivity_params)
 
    if @dataProcessingActivity.save
      redirect_to dataProcessingActivitys_path
    else
      render 'new'
    end
  end
 
  def update
    @dataProcessingActivity = DataProcessingActivity.find(params[:id])
 
    if @dataProcessingActivity.update(dataProcessingActivity_params)
      redirect_to dataProcessingActivitys_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @dataProcessingActivity = DataProcessingActivity.find(params[:id])
    @dataProcessingActivity.destroy
    redirect_to dataProcessingActivitys_path
  end

 
  private
    def dataProcessingActivity_params
      params.require(:dataProcessingActivity).permit(:name, :purpose, :startDate, :LawfulBasis)
    end
end