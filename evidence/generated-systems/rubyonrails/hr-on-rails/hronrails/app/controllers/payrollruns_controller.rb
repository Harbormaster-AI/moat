class PayrollRunsController < ApplicationController
  def index
    @payrollRuns = PayrollRun.all
  end
 
  def show
    @payrollRun = PayrollRun.find(params[:id])
  end
 
  def new
    @payrollRun = PayrollRun.new
  end
 
  def edit
    @payrollRun = PayrollRun.find(params[:id])
  end
 
  def create
    @payrollRun = PayrollRun.new(payrollRun_params)
 
    if @payrollRun.save
      redirect_to payrollRuns_path
    else
      render 'new'
    end
  end
 
  def update
    @payrollRun = PayrollRun.find(params[:id])
 
    if @payrollRun.update(payrollRun_params)
      redirect_to payrollRuns_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @payrollRun = PayrollRun.find(params[:id])
    @payrollRun.destroy
    redirect_to payrollRuns_path
  end

 
  private
    def payrollRun_params
      params.require(:payrollRun).permit(:runNumber, :periodStart, :periodEnd, :paymentDate, :Status)
    end
end