class ForecastLinesController < ApplicationController
  def index
    @forecastLines = ForecastLine.all
  end
 
  def show
    @forecastLine = ForecastLine.find(params[:id])
  end
 
  def new
    @forecastLine = ForecastLine.new
  end
 
  def edit
    @forecastLine = ForecastLine.find(params[:id])
  end
 
  def create
    @forecastLine = ForecastLine.new(forecastLine_params)
 
    if @forecastLine.save
      redirect_to forecastLines_path
    else
      render 'new'
    end
  end
 
  def update
    @forecastLine = ForecastLine.find(params[:id])
 
    if @forecastLine.update(forecastLine_params)
      redirect_to forecastLines_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @forecastLine = ForecastLine.find(params[:id])
    @forecastLine.destroy
    redirect_to forecastLines_path
  end

 
  private
    def forecastLine_params
      params.require(:forecastLine).permit(:period, :quantity, :confidence)
    end
end