class PayrollCalendarsController < ApplicationController
  def index
    @payrollCalendars = PayrollCalendar.all
  end
 
  def show
    @payrollCalendar = PayrollCalendar.find(params[:id])
  end
 
  def new
    @payrollCalendar = PayrollCalendar.new
  end
 
  def edit
    @payrollCalendar = PayrollCalendar.find(params[:id])
  end
 
  def create
    @payrollCalendar = PayrollCalendar.new(payrollCalendar_params)
 
    if @payrollCalendar.save
      redirect_to payrollCalendars_path
    else
      render 'new'
    end
  end
 
  def update
    @payrollCalendar = PayrollCalendar.find(params[:id])
 
    if @payrollCalendar.update(payrollCalendar_params)
      redirect_to payrollCalendars_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @payrollCalendar = PayrollCalendar.find(params[:id])
    @payrollCalendar.destroy
    redirect_to payrollCalendars_path
  end

 
  private
    def payrollCalendar_params
      params.require(:payrollCalendar).permit(:name, :country, :PayFrequency)
    end
end