class TimeSeriessController < ApplicationController
  def index
    @timeSeriess = TimeSeries.all
  end
 
  def show
    @timeSeries = TimeSeries.find(params[:id])
  end
 
  def new
    @timeSeries = TimeSeries.new
  end
 
  def edit
    @timeSeries = TimeSeries.find(params[:id])
  end
 
  def create
    @timeSeries = TimeSeries.new(timeSeries_params)
 
    if @timeSeries.save
      redirect_to timeSeriess_path
    else
      render 'new'
    end
  end
 
  def update
    @timeSeries = TimeSeries.find(params[:id])
 
    if @timeSeries.update(timeSeries_params)
      redirect_to timeSeriess_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @timeSeries = TimeSeries.find(params[:id])
    @timeSeries.destroy
    redirect_to timeSeriess_path
  end

 
  private
    def timeSeries_params
      params.require(:timeSeries).permit(:name, :timezone, :Granularity)
    end
end