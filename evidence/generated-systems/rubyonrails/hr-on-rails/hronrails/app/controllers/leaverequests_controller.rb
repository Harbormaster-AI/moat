class LeaveRequestsController < ApplicationController
  def index
    @leaveRequests = LeaveRequest.all
  end
 
  def show
    @leaveRequest = LeaveRequest.find(params[:id])
  end
 
  def new
    @leaveRequest = LeaveRequest.new
  end
 
  def edit
    @leaveRequest = LeaveRequest.find(params[:id])
  end
 
  def create
    @leaveRequest = LeaveRequest.new(leaveRequest_params)
 
    if @leaveRequest.save
      redirect_to leaveRequests_path
    else
      render 'new'
    end
  end
 
  def update
    @leaveRequest = LeaveRequest.find(params[:id])
 
    if @leaveRequest.update(leaveRequest_params)
      redirect_to leaveRequests_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @leaveRequest = LeaveRequest.find(params[:id])
    @leaveRequest.destroy
    redirect_to leaveRequests_path
  end

 
  private
    def leaveRequest_params
      params.require(:leaveRequest).permit(:requestNumber, :startDate, :endDate, :reason, :hours, :Status)
    end
end